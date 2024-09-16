package resolver

import (
	"github.com/Lux00000/post-and-comments/internal/service"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	PostsService      service.PostsService
	CommentsService   service.CommentsService
	CommentsObservers *CommentsObservers
}

func NewResolver(postsService service.PostsService, commentsService service.CommentsService, observers *CommentsObservers) *Resolver {
	return &Resolver{
		PostsService:      postsService,
		CommentsService:   commentsService,
		CommentsObservers: observers,
	}
}
