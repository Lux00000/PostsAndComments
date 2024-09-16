package service

import (
	"github.com/Lux00000/post-and-comments/internal/models"
	"github.com/stretchr/testify/mock"
)

// MockComments - mock для storage.Comments
type MockComments struct {
	mock.Mock
}

func (m *MockComments) CreateComment(comment models.Comment) (models.Comment, error) {
	args := m.Called(comment)
	return args.Get(0).(models.Comment), args.Error(1)
}

func (m *MockComments) GetCommentsByPost(postId int, page *int, pageSize *int) ([]*models.Comment, error) {
	args := m.Called(postId, page, pageSize)
	return args.Get(0).([]*models.Comment), args.Error(1)
}

func (m *MockComments) GetChildrenOfComment(commentId int) ([]*models.Comment, error) {
	args := m.Called(commentId)
	return args.Get(0).([]*models.Comment), args.Error(1)
}

type MockPosts struct {
	mock.Mock
}

func (m *MockPosts) GetPostById(id int) (*models.Post, error) {
	args := m.Called(id)
	return args.Get(0).(*models.Post), args.Error(1)
}

func (m *MockPosts) CreatePost(post models.Post) (models.Post, error) {
	args := m.Called(post)
	return args.Get(0).(models.Post), args.Error(1)
}

func (m *MockPosts) GetAllPosts(limit int, offset int) ([]models.Post, error) {
	args := m.Called(limit, offset)
	return args.Get(0).([]models.Post), args.Error(1)
}
