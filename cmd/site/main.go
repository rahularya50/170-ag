package main

import (
	"170-ag/ent/generated/migrate"
	_ "170-ag/ent/generated/runtime"
	"170-ag/resolvers"
	"170-ag/site"
	"170-ag/site/web"
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	client, err := site.GetEntClient()
	if err != nil {
		log.Fatal("opening ent client", err)
	}
	defer client.Close()

	if err := client.Schema.Create(
		context.Background(),
		migrate.WithGlobalUniqueID(true),
	); err != nil {
		log.Fatal("running schema migration", err)
	}

	gqlServer := handler.NewDefaultServer(resolvers.NewSchema(client))
	gqlServer.SetErrorPresenter(func(ctx context.Context, err error) *gqlerror.Error {
		return gqlerror.Errorf("something went wrong")
	})
	gqlServer.Use(extension.FixedComplexityLimit(250))

	srv := web.HandlerWithViewerContext(gqlServer, client)
	srv = site.HandleWithEntClient(srv, client)

	http.Handle("/playground", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	http.Handle("/login", web.LoginHandler(client))
	http.HandleFunc("/logout", web.HandleLogout)

	log.Printf("connect to http://localhost:%s/playground for GraphQL playground", port)

	http.Handle("/static/", http.FileServer(http.Dir("frontend/build/")))

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		http.ServeFile(rw, r, "frontend/build/index.html")
	})

	httpSrv := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         ":" + port,
	}
	log.Fatal(httpSrv.ListenAndServe())
}
