package site

import (
	ent "170-ag/ent/generated"
	"170-ag/ent/generated/privacy"
	"170-ag/ent/generated/user"
	"context"
	"crypto/hmac"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	"google.golang.org/api/idtoken"
)

const clientId = "576862966452-u8f9dgug7pod1thuvaclp397h8ngm0aq.apps.googleusercontent.com"

type key int

var userEmailKey key = 0
var viewerKey key = 1

func newLoginContext(ctx context.Context, email string) context.Context {
	return context.WithValue(ctx, userEmailKey, email)
}

func EmailFromContext(ctx context.Context) (string, bool) {
	email, ok := ctx.Value(userEmailKey).(string)
	return email, ok
}

func ContextWithViewer(ctx context.Context, req *http.Request, client *ent.Client) context.Context {
	session, _ := getSessionStore().Get(req, "login-session")
	raw_user_id, exists := session.Values["user_id"]
	if !exists {
		return ctx
	}
	user_id, ok := raw_user_id.(int)
	if !ok {
		return ctx
	}
	// override checks so we can initially populate viewerContext
	allow_ctx := privacy.DecisionContext(ctx, privacy.Allow)
	viewer, err := client.User.Get(allow_ctx, user_id)
	if err != nil {
		return ctx
	}
	return context.WithValue(ctx, viewerKey, viewer)
}

func ViewerFromContext(ctx context.Context) (*ent.User, bool) {
	viewer, ok := ctx.Value(viewerKey).(*ent.User)
	return viewer, ok
}

func HandlerWithViewerContext(h http.Handler, client *ent.Client) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		r = r.WithContext(ContextWithViewer(r.Context(), r, client))
		h.ServeHTTP(rw, r)
	})
}

type loginHandler struct {
	client *ent.Client
}

func getSessionStore() *sessions.CookieStore {
	session_key, exists := os.LookupEnv("SESSION_KEY")
	if !exists {
		panic("Session key envvar not set!")
	}
	return sessions.NewCookieStore([]byte(session_key))
}

func LoginHandler(client *ent.Client) *loginHandler {
	return &loginHandler{client: client}
}

func (handler *loginHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	cookieCsrfToken, err := req.Cookie("g_csrf_token")
	if err != nil {
		http.Error(resp, "Missing CSRF cookie", http.StatusUnauthorized)
		return
	}
	csrfToken := req.PostFormValue("g_csrf_token")
	if !hmac.Equal([]byte(cookieCsrfToken.Value), []byte(csrfToken)) {
		http.Error(resp, "CSRF violation", http.StatusUnauthorized)
		return
	}
	credential := req.FormValue("credential")
	payload, err := idtoken.Validate(req.Context(), credential, clientId)
	if err != nil {
		http.Error(resp, "Invalid credentials", http.StatusUnauthorized)
		return
	}
	if !payload.Claims["email_verified"].(bool) {
		http.Error(resp, "Unverified email", http.StatusUnauthorized)
		return
	}
	email := payload.Claims["email"].(string)
	email_ctx := newLoginContext(req.Context(), email)

	raw_name := payload.Claims["name"]
	var name *string
	if raw_name == nil {
		name = nil
	} else {
		n := raw_name.(string)
		name = &n
	}

	user_id := handler.client.User.
		Create().
		SetEmail(email).
		SetNillableName(name).
		OnConflictColumns(user.FieldEmail).
		UpdateNewValues().
		IDX(email_ctx)

	// silently ignore session decode failures
	store := getSessionStore()
	session, _ := store.Get(req, "login-session")
	session.Values["user_id"] = user_id
	err = session.Save(req, resp)
	if err != nil {
		panic("failed to update session store")
	}

	_, err = resp.Write([]byte(fmt.Sprintf("Logged in with email %s and name %s", email, *name)))
	if err != nil {
		panic("failed to write back to client")
	}
}
