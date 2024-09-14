package inmemory

import (
	"errors"
	"github.com/Lux00000/PostsAndComments/internal/models"
	"sync"
)

type CommentsInMemory struct {
	commCounter int
	comments    []models.Comment
	mu          sync.RWMutex
}

func NewCommentsInMemory(count int) *CommentsInMemory {
	return &CommentsInMemory{
		commCounter: 0,
		comments:    make([]models.Comment, 0),
	}
}

func (c *CommentsInMemory) CreateComment(comment models.Comment) (models.Comment, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.commCounter++

	comment.ID = c.commCounter

	c.comments = append(c.comments, comment)

	return comment, nil

}

func (c *CommentsInMemory) GetCommentsByPost(postId, limit, offset int) ([]*models.Comment, error) {

	c.mu.RLock()
	defer c.mu.RUnlock()

	var res []*models.Comment

	for _, comment := range c.comments {
		if comment.ParentCommentID == nil && comment.PostID == postId {
			com := comment
			res = append(res, &com)
		}
	}

	if offset > len(res) {
		return nil, errors.New("offset out of range")
	}

	if offset+limit > len(res) || limit == -1 {
		return res[offset:], nil
	}

	return res[offset : offset+limit], nil
}

func (c *CommentsInMemory) GetChildrenOfComment(commentId int) ([]*models.Comment, error) {

	c.mu.RLock()
	defer c.mu.RUnlock()

	if commentId > c.commCounter {
		return nil, nil
	}

	var res []*models.Comment

	for _, comment := range c.comments {
		if comment.ParentCommentID != nil && *comment.ParentCommentID == commentId {
			com := comment
			res = append(res, &com)
		}
	}

	return res, nil
}
