package graph

import (
	"github.com/Lux00000/PostsAndComments/internal/service"
	"github.com/Lux00000/PostsAndComments/server"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	PostsService      service.Posts
	CommentsService   service.Comments
	CommentsObservers server.Observers
}
