package inmemory

import (
	"errors"
	"github.com/Lux00000/post-and-comments/internal/models"
	"sync"
)

type InMemoryComment struct {
	comments map[int]*models.Comment
	mu       sync.RWMutex
}

func NewInMemoryComment() *InMemoryComment {
	return &InMemoryComment{
		comments: make(map[int]*models.Comment, 0),
	}
}

func (s *InMemoryComment) CreateComment(comment models.Comment) (models.Comment, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	comment.ID = len(s.comments) + 1
	s.comments[comment.ID] = &comment
	return comment, nil
}

func (s *InMemoryComment) GetCommentsByPost(postId int, page *int, pageSize *int) ([]*models.Comment, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	comments := make([]*models.Comment, 0)
	for _, comment := range s.comments {
		if comment.PostID == postId {
			comments = append(comments, comment)
		}
	}

	if page == nil || pageSize == nil {
		return comments, nil
	}

	offset := (*page - 1) * *pageSize
	limit := *pageSize

	if offset >= len(comments) {
		return []*models.Comment{}, nil
	}

	if offset < 0 || limit < 0 {
		return nil, errors.New("limit or offset < 0")
	}

	if offset+limit > len(comments) {
		limit = len(comments) - offset
	}

	return comments[offset : offset+limit], nil
}

func (s *InMemoryComment) GetChildrenOfComment(commentId int) ([]*models.Comment, error) {
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
