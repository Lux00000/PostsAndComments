package interfaces

import (
	"github.com/Lux00000/PostsAndComments/internal/models"
)

type PostService interface {
	CreatePost(post *models.Post) (*models.Post, error)
	GetAllPosts(page, pageSize int) ([]*models.Post, error)
	GetPostByID(id int) (*models.Post, error)
}
