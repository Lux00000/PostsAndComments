package graph

import (
	"database/sql"
	graph "github.com/Lux00000/PostsAndComments/graph/inmemory_resolver"
	"github.com/Lux00000/PostsAndComments/graph/interfaces"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	PostService    interfaces.PostService
	CommentService interfaces.CommentService
}

func NewResolver(db *sql.DB) *Resolver {
	var postService interfaces.PostService
	var commentService interfaces.CommentService

	if db != nil {
		postService = NewDBPostService(db)
		commentService = NewDBCommentService(db)
	} else {
		postService = graph.NewInMemoryPostService()
		commentService = graph.NewInMemoryCommentService()
	}

	return &Resolver{
		PostService:    postService,
		CommentService: commentService,
	}
}
