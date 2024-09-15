package service

import (
	"errors"
	"github.com/Lux00000/PostsAndComments/internal/models"
	"github.com/Lux00000/PostsAndComments/internal/storage"
)

type PostsService struct {
	pst storage.Posts
}

func NewPostsService(pst storage.Posts) *PostsService {
	return &PostsService{pst: pst}
}

func (p PostsService) CreatePost(post models.Post) (models.Post, error) {

	if post.AuthorId == 0 {
		return models.Post{}, errors.New("AuthorId is zero")
	}

	newPost, err := p.pst.CreatePost(post)
	if err != nil {
		return models.Post{}, err
	}

	return newPost, nil

}

func (p PostsService) GetPostById(postId int) (models.Post, error) {

	if postId <= 0 {
		return models.Post{}, errors.New("PostId is zero")
	}

	post, err := p.pst.GetPostById(postId)
	if err != nil {
		return models.Post{}, err
	}

	return post, nil
}

func (p PostsService) GetAllPosts(page, pageSize *int) ([]models.Post, error) {

	if page != nil && *page < 0 {
		return nil, errors.New("page is negative")

	}

	if pageSize != nil && *pageSize < 0 {
		return nil, errors.New("pageSize is negative")
	}
	offset := (*page - 1) * *pageSize
	limit := *pageSize

	posts, err := p.pst.GetAllPosts(limit, offset)
	if err != nil {
		return nil, errors.New("GetAllPostsService error")
	}

	return posts, nil
}
