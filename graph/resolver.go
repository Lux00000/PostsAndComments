package graph

import (
	"database/sql"
	graph "github.com/Lux00000/PostsAndComments/graph/inmemory_resolver"
	"github.com/Lux00000/PostsAndComments/graph/interfaces"
	"github.com/Lux00000/PostsAndComments/graph/postgres_resolver"
)

type Resolver struct {
	PostService    interfaces.PostService
	CommentService interfaces.CommentService
}

func NewResolver(db *sql.DB) *Resolver {
	var postService interfaces.PostService
	var commentService interfaces.CommentService

	if db != nil {
		postService = postgres_resolver.NewDBPostService(db)
		commentService = postgres_resolver.NewDBCommentService(db)
	} else {
		postService = graph.NewInMemoryPostService()
		commentService = graph.NewInMemoryCommentService()
	}

	return &Resolver{
		PostService:    postService,
		CommentService: commentService,
	}
}
