package storage

import "github.com/Lux00000/PostsAndComments/internal/models"

type StorageHandler struct {
	Posts
	Comments
}

type Posts interface {
	CreatePost(post models.Post) (models.Post, error)
	GetPostById(id int) (models.Post, error)
	GetAllPosts(limit, offset int) ([]models.Post, error)
}

type Comments interface {
	CreateComment(comment models.Comment) (models.Comment, error)
	GetCommentsByPost(postId, limit, offset int) ([]*models.Comment, error)
	GetChildrenOfComment(commentId int) ([]*models.Comment, error)
}
