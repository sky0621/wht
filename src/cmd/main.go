package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	app "github.com/sky0621/wht"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

// see https://cloud.google.com/run/docs/quickstarts/build-and-deploy?hl=ja
func main() {
	log.Print("helloworld3: starting server...")

	srv := handler.NewDefaultServer(app.NewExecutableSchema(app.Config{Resolvers: &app.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("helloworld: listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
