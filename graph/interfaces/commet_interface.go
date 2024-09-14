package interfaces

import "github.com/Lux00000/PostsAndComments/graph/model"

type CommentService interface {
	CreateComment(comment *model.Comment) (*model.Comment, error)
	CommentsSubscription(postID string) (<-chan *model.Comment, error)
}
