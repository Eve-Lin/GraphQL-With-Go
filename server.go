package main

import (
	"log"
	"net/http"

	"github.com/graphql-go/handler"
)

func main() {

	dbInit()

	log.Println("Initializing the schema")
	SchemaInit()

	graphqlHandler := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})

	log.Println("The server is starting")
	http.Handle("/graphql", graphqlHandler)
	http.ListenAndServe(":9999", nil)
}
