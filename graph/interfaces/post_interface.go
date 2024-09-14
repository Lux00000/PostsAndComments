package interfaces

import "github.com/Lux00000/PostsAndComments/graph/model"

type PostService interface {
	CreatePost(post *model.Post) (*model.Post, error)
	GetAllPosts(page, pageSize int) ([]*model.Post, error)
	GetPostByID(id int) (*model.Post, error)
}
