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
	var (
		dbUser = os.Getenv("WHT_DB_USER")
		dbPass = os.Getenv("WHT_DB_PASS")
		//dbHost = os.Getenv("WHT_DB_HOST")
		//dbPort = os.Getenv("WHT_DB_PORT")
		instanceConnectionName = os.Getenv("WHT_DB_INSTANCE_CONNECTION_NAME")
		dbName                 = os.Getenv("WHT_DB_NAME")
	)

	var dbURI string
	dbURI = fmt.Sprintf("%s:%s@unix(/cloudsql/%s)/%s", dbUser, dbPass, instanceConnectionName, dbName)
	//dbURI = fmt.Sprintf("host=%s user=%s password=%s port=%s database=%s", dbHost, dbUser, dbPass, dbPort, dbName)
	//dbURI = "host=localhost port=15432 dbname=wolf-db user=postgres password=xxxx sslmode=disable"

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

	log.Printf("helloworld: listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
