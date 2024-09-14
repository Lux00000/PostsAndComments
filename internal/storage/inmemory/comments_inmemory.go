package inmemory

import (
	"errors"
	"github.com/Lux00000/PostsAndComments/internal/models"
	"sync"
)

type inMemoryComment struct {
	comments map[int]*models.Comment
	mu       sync.RWMutex
}

func NewInMemoryComment() *inMemoryComment {
	return &inMemoryComment{
		comments: make(map[int]*models.Comment, 0),
	}
}

func (s *inMemoryComment) CreateComment(comment *models.Comment) (*models.Comment, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	comment.ID = len(s.comments) + 1
	s.comments[comment.ID] = comment
	return comment, nil
}

func (s *inMemoryComment) GetCommentsByPost(posId, limit, offset int) ([]*models.Comment, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	comments := make([]*models.Comment, 0)
	for _, comment := range s.comments {
		if comment.PostID == posId {
			comments = append(comments, comment)
		}
	}
	if offset > len(comments) {
		return nil, errors.New("offset > len(comments)")
	}
	if offset < 0 || limit < 0 {
		return nil, errors.New("limit or offset < 0")
	}
	return comments[offset:limit], nil
}

func (s *inMemoryComment) GetChildrenComment(commentId int) ([]*models.Comment, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if commentId > len(s.comments) {
		return nil, errors.New("commentId > len(s.comments)")
	}
	comments := make([]*models.Comment, 0)
	for _, comment := range s.comments {
		if comment.ID == commentId {
			comments = append(comments, comment)
		}
	}
	return comments, nil
}
