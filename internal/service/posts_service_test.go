package service

import (
	"github.com/Lux00000/post-and-comments/internal/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreatePost(t *testing.T) {
	mockPosts := new(MockPosts)
	service := NewPostsService(mockPosts)

	post := models.Post{
		AuthorId: 1,
		Title:    "Test Post",
		Content:  "This is a test post.",
	}

	mockPosts.On("CreatePost", post).Return(post, nil)

	result, err := service.CreatePost(post)
	assert.NoError(t, err)
	assert.Equal(t, post, result)

	mockPosts.AssertExpectations(t)
}

func TestGetPostById(t *testing.T) {
	mockPosts := new(MockPosts)
	service := NewPostsService(mockPosts)

	post := &models.Post{
		ID:       1,
		AuthorId: 1,
		Title:    "Test Post",
		Content:  "This is a test post.",
	}

	mockPosts.On("GetPostById", 1).Return(post, nil)

	result, err := service.GetPostById(1)
	assert.NoError(t, err)
	assert.Equal(t, *post, result)

	mockPosts.AssertExpectations(t)
}

func TestGetAllPosts(t *testing.T) {
	mockPosts := new(MockPosts)
	service := NewPostsService(mockPosts)

	posts := []models.Post{
		{ID: 1, AuthorId: 1, Title: "Post 1", Content: "Content 1"},
		{ID: 2, AuthorId: 2, Title: "Post 2", Content: "Content 2"},
	}

	page := 1
	pageSize := 10

	mockPosts.On("GetAllPosts", page, pageSize).Return(posts, nil)

	result, err := service.GetAllPosts(&page, &pageSize)

	assert.NoError(t, err)
	assert.Equal(t, posts, result)

	mockPosts.AssertExpectations(t)
}
