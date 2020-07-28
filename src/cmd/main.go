package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"

	"github.com/sky0621/wht/adapter/controller"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

// see https://cloud.google.com/run/docs/quickstarts/build-and-deploy?hl=ja
func main() {
	log.Print("wht: starting server...")

	/*
	 * DB connect setting
	 */
	var dbURI string
	if os.Getenv("WHT_ENV") == "local" {
		var (
			dbPass = os.Getenv("WHT_DB_PASS")
		)
		dbURI = fmt.Sprintf("host=localhost port=11111 dbname=whtdb user=postgres password=%s sslmode=disable", dbPass)
	} else {
		var (
			dbHost = os.Getenv("WHT_DB_HOST")
			dbUser = os.Getenv("WHT_DB_USER")
			dbPass = os.Getenv("WHT_DB_PASS")
			dbName = os.Getenv("WHT_DB_NAME")
		)
		dbURI = fmt.Sprintf("host=/cloudsql/%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbUser, dbPass, dbName)
	}

	// dbPool is the pool of database connections.
	dbPool, err := sqlx.Open("postgres", dbURI)
	if err != nil {
		log.Println(err)
		return
	}

	if err := dbPool.Ping(); err != nil {
		log.Println(err)
		return
	}
	defer func() {
		if dbPool != nil {
			if err := dbPool.Close(); err != nil {
				log.Println(err)
			}
		}
	}()

	/*
	 * Web server setting
	 */
	srv := handler.NewDefaultServer(controller.NewExecutableSchema(controller.Config{Resolvers: controller.NewResolver(dbPool)}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("wht: listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
