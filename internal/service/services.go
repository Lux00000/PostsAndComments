package service

import (
	"github.com/Lux00000/PostsAndComments/internal/models"
	"github.com/Lux00000/PostsAndComments/internal/storage"
)

type Services struct {
	Posts
	Comments
}

func NewServices(storage *storage.StorageHandler) *Services {
	return &Services{
		Posts:    NewPostsService(storage.Posts),
		Comments: NewCommentsService(storage.Comments, storage.Posts),
	}
}

type Posts interface {
	CreatePost(post models.Post) (models.Post, error)
	GetPostById(id int) (models.Post, error)
	GetAllPosts(page, pageSize *int) ([]models.Post, error)
}

type Comments interface {
	CreateComment(comment models.Comment) (models.Comment, error)
	GetCommentsByPost(postId int, page *int, pageSize *int) ([]*models.Comment, error)
	GetChildrenOfComment(commentId int) ([]*models.Comment, error)
}
