package main

import (
	"database/sql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	graphql "github.com/Lux00000/post-and-comments/internal/graphql/gen"
	"github.com/Lux00000/post-and-comments/internal/graphql/resolver"
	"github.com/Lux00000/post-and-comments/internal/service"
	"github.com/Lux00000/post-and-comments/internal/storage"
	"github.com/Lux00000/post-and-comments/internal/storage/inmemory"
	"github.com/Lux00000/post-and-comments/internal/storage/postgres"
	"net/http"

	"log"
	"os"

	_ "github.com/lib/pq"
)

const defaultPort = "8080"

func main() {
	os.Setenv("STORAGE_TYPE", "DATABASE_URL")
	os.Setenv("DATABASE_URL", "postgres://user:password@db:5432/postsandcomments?sslmode=disable")
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	var postsStorage storage.Posts
	var commentStorage storage.Comments

	storageType := os.Getenv("STORAGE_TYPE")
	if storageType == "inmemory" {
		log.Println("Using in-memory storage")
		postsStorage = inmemory.NewInMemoryPost()
		commentStorage = inmemory.NewInMemoryComment()
	} else {
		dbURL := os.Getenv("DATABASE_URL")
		if dbURL == "" {
			log.Fatal("DATABASE_URL environment variable is not set")
		}
		db, err := sql.Open("postgres", dbURL)
		if err != nil {
			log.Fatalf("failed to connect to database: %v", err)
		}
		defer db.Close()
		postsStorage = postgres.NewDBPostPostgres(db)
		commentStorage = postgres.NewDBCommentPostgres(db)
	}
	var postService = service.NewPostsService(postsStorage)
	var commentService = service.NewCommentsService(commentStorage, postsStorage)
	rslvr := resolver.NewResolver(*postService, *commentService, resolver.NewCommentsObserver())
	srv := handler.NewDefaultServer(graphql.NewExecutableSchema(graphql.Config{Resolvers: rslvr}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
