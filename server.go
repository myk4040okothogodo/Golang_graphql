package main

import (
	"log"
	"net/http"
	"os"
   database "github.com/myk4040okothogodo/hackernews/internal/pkg/db/migrations/mysql"
  "github.com/go-chi/chi/v5"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/myk4040okothogodo/hackernews/graph"
	"github.com/myk4040okothogodo/hackernews/graph/generated"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}


  router := chi.NewRouter()
  database.InitDB()
  database.Migrate()

  server := handler.NewDefaultServer(hackernews.NewExecutableSchema(hackernews.Config{Resolvers: &hackernews.Resolver{}, &graph.Resolver{}}))
	//srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", server)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
