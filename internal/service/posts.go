package service

import (
	"errors"
	"fmt"
	"github.com/Lux00000/post-and-comments/internal/models"
	"github.com/Lux00000/post-and-comments/internal/storage"
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
		return models.Post{}, fmt.Errorf("error creating post: %w", err)
	}

	return newPost, nil

}

func (p PostsService) GetPostById(postId int) (models.Post, error) {

	if postId <= 0 {
		return models.Post{}, errors.New("PostId is zero")
	}

	post, err := p.pst.GetPostById(postId)
	if err != nil {
		return models.Post{}, fmt.Errorf("error getting post by id: %v", err)
	}

	return *post, nil
}

func (p PostsService) GetAllPosts(page, pageSize *int) ([]models.Post, error) {

	posts, err := p.pst.GetAllPosts(*page, *pageSize)
	if err != nil {
		return nil, fmt.Errorf("error getting all posts: %v", err)
	}

	return posts, nil
}
