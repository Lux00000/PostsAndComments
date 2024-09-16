package storage

import "github.com/Lux00000/post-and-comments/internal/models"

type Posts interface {
	CreatePost(post models.Post) (models.Post, error)
	GetPostById(id int) (*models.Post, error)
	GetAllPosts(limit, offset int) ([]models.Post, error)
}

type Comments interface {
	CreateComment(comment models.Comment) (models.Comment, error)
	GetCommentsByPost(postId int, page *int, pageSize *int) ([]*models.Comment, error)
	GetChildrenOfComment(commentId int) ([]*models.Comment, error)
}
