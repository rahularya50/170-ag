package main

import (
	"170-ag/ent/generated/migrate"
	_ "170-ag/ent/generated/runtime"
	"170-ag/resolvers"
	"170-ag/site"
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
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

	srv := site.HandlerWithViewerContext(
		handler.NewDefaultServer(resolvers.NewSchema(client)),
		client,
	)

	http.Handle("/playground", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	http.Handle("/login", site.LoginHandler(client))
	http.HandleFunc("/logout", site.HandleLogout)

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
