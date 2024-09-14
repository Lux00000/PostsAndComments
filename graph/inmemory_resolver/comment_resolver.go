package graph

import (
	"github.com/Lux00000/PostsAndComments/graph/interfaces"
	"github.com/Lux00000/PostsAndComments/graph/model"
	"sync"
)

type inMemoryCommentService struct {
	comments map[int]*model.Comment
	mu       sync.RWMutex
}

func NewInMemoryCommentService() interfaces.CommentService {
	return &inMemoryCommentService{
		comments: make(map[int]*model.Comment),
	}
}

func (s *inMemoryCommentService) CreateComment(comment *model.Comment) (*model.Comment, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	comment.ID = len(s.comments) + 1
	s.comments[comment.ID] = comment
	return comment, nil
}

func (s *inMemoryCommentService) CommentsSubscription(postID string) (<-chan *model.Comment, error) {

	return make(chan *model.Comment), nil
}
