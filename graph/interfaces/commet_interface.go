package interfaces

import (
	"github.com/Lux00000/PostsAndComments/internal/models"
)

type CommentService interface {
	CreateComment(comment *models.Comment) (*models.Comment, error)
	CommentsSubscription(postID string) (<-chan *models.Comment, error)
}
