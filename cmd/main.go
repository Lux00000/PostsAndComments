package main

import (
	"database/sql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Lux00000/PostsAndComments/graph"
	"github.com/Lux00000/PostsAndComments/server"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	var db *sql.DB
	var err error

	storageType := os.Getenv("STORAGE_TYPE")
	if storageType == "inmemory" {
		log.Println("Using in-memory storage")
		db = nil
	} else {
		dbURL := os.Getenv("DATABASE_URL")
		if dbURL == "" {
			log.Fatal("DATABASE_URL environment variable is not set")
		}

		db, err = sql.Open("postgres", dbURL)
		if err != nil {
			log.Fatalf("failed to connect to database: %v", err)
		}
		defer db.Close()
	}

	resolver := server.NewResolver(db)
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
