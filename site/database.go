package site

import (
	ent "170-ag/ent/generated"
	"context"
	"net/http"
	"os"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

func GetEntClient() (*ent.Client, error) {
	env := os.Getenv("ENV")
	if env == "dev" {
		return ent.Open(
			"sqlite3",
			"file:app.db?&_fk=1",
		)
	} else {
		// This connects to the Cloud SQL Proxy on the same Pod, which has an encrypted connection to the actual db
		// This is why it is OK for SSL to be disabled!
		return ent.Open(
			"postgres",
			"host=localhost port=5432 user=autograder-web@formidable-gate-337712.iam dbname=autograder sslmode=disable",
		)
	}
}

type clientKey struct{}

func ContextWithEntClient(ctx context.Context, client *ent.Client) context.Context {
	return context.WithValue(ctx, clientKey{}, client)
}

func EntClientFromContext(ctx context.Context) *ent.Client {
	return ctx.Value(clientKey{}).(*ent.Client)
}

func HandleWithEntClient(h http.Handler, client *ent.Client) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		r = r.WithContext(ContextWithEntClient(r.Context(), client))
		h.ServeHTTP(rw, r)
	})
}
