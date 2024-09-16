package service

import (
	"github.com/Lux00000/post-and-comments/internal/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateComment(t *testing.T) {
	mockComments := new(MockComments)
	mockPosts := new(MockPosts)
	service := NewCommentsService(mockComments, mockPosts)

	comment := models.Comment{
		AuthorId: 1,
		PostID:   1,
		Text:     "Test comment",
	}

	post := &models.Post{
		ID:            1,
		AllowComments: true,
	}

	mockPosts.On("GetPostById", 1).Return(post, nil)
	mockComments.On("CreateComment", comment).Return(comment, nil)

	result, err := service.CreateComment(comment)
	assert.NoError(t, err)
	assert.Equal(t, comment, result)

	mockPosts.AssertExpectations(t)
	mockComments.AssertExpectations(t)
}

func TestGetCommentsByPost(t *testing.T) {
	mockComments := new(MockComments)
	mockPosts := new(MockPosts)
	service := NewCommentsService(mockComments, mockPosts)

	comments := []*models.Comment{
		{ID: 1, PostID: 1, Text: "Comment 1"},
		{ID: 2, PostID: 1, Text: "Comment 2"},
	}

	page := 1
	pageSize := 10

	mockComments.On("GetCommentsByPost", 1, &page, &pageSize).Return(comments, nil)

	result, err := service.GetCommentsByPost(1, &page, &pageSize)
	assert.NoError(t, err)
	assert.Equal(t, comments, result)

	mockComments.AssertExpectations(t)
}

func TestGetChildrenOfComment(t *testing.T) {
	mockComments := new(MockComments)
	mockPosts := new(MockPosts)
	service := NewCommentsService(mockComments, mockPosts)

	children := []*models.Comment{
		{ID: 2, ParentCommentID: stringPtr("1"), Text: "Child 1"},
		{ID: 3, ParentCommentID: stringPtr("1"), Text: "Child 2"},
	}

	mockComments.On("GetChildrenOfComment", 1).Return(children, nil)

	result, err := service.GetChildrenOfComment(1)
	assert.NoError(t, err)
	assert.Equal(t, children, result)

	mockComments.AssertExpectations(t)
}

func stringPtr(s string) *string {
	return &s
}
