package main

import (
	"170-ag/ent/generated/migrate"
	_ "170-ag/ent/generated/runtime"
	"170-ag/proto/schemas"
	"170-ag/resolvers"
	"170-ag/site"
	"170-ag/site/grading"
	"170-ag/site/project"
	"170-ag/site/web"
	"context"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"google.golang.org/grpc"
)

const defaultPort = "8080"
const defaultGrpcPort = "8081"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	grpcPort := os.Getenv("GRPC_PORT")
	if grpcPort == "" {
		grpcPort = defaultGrpcPort
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

	gscopeToken := os.Getenv("GRADESCOPE_TOKEN")
	if gscopeToken == "" {
		panic("gradescope API token must be provided")
	}

	projectServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			project.AuthProjectRequest(client, gscopeToken),
			grading.InterceptErrors,
		),
	)
	grpcPortInt, err := strconv.Atoi(grpcPort)
	if err != nil {
		panic(err)
	}
	schemas.RegisterProjectScoresServer(projectServer, &project.ProjectScoresServer{Client: client})
	go site.ServeOnPort(projectServer, grpcPortInt)

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

	scoreboardToken := os.Getenv("SCOREBOARD_TOKEN")
	if scoreboardToken == "" {
		panic("scoreboard API token must be provided")
	}

	if os.Getenv("ENV") == "dev" {
		http.Handle(
			"/scoreboard",
			project.ScoreboardHandler(client),
		)
	} else {
		http.HandleFunc(
			"/scoreboard",
			project.HandlerCheckingAuthorizationToken(
				project.ScoreboardHandler(client), scoreboardToken,
			),
		)
	}

	httpSrv := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         ":" + port,
	}
	log.Fatal(httpSrv.ListenAndServe())
}
