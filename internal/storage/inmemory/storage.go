package inmemory

import (
	"errors"
	"github.com/Lux00000/PostsAndComments/internal/models"
	"sync"
)

type Storage struct {
	posts    map[int]*models.Post
	comments map[int]*models.Comment
	mu       sync.RWMutex
}

func NewStorage() *Storage {
	return &Storage{
		posts:    make(map[int]*models.Post),
		comments: make(map[int]*models.Comment),
	}
}

func (s *Storage) AddPost(post *models.Post) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.posts[post.ID] = post
}

func (s *Storage) GetPost(id int) (*models.Post, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	post, exists := s.posts[id]
	return post, exists
}

func (s *Storage) GetAllPosts() []*models.Post {
	s.mu.RLock()
	defer s.mu.RUnlock()
	posts := make([]*models.Post, 0, len(s.posts))

	for _, post := range s.posts {
		posts = append(posts, post)
	}
	return posts
}

func (s *Storage) AddComment(comment *models.Comment) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	targetPost, exists := s.posts[comment.PostID]
	if targetPost.AllowComments == false {
		return errors.New("post disabled comments")
	}
	s.comments[comment.ID] = comment

	if comment.ParentID != nil {
		parentComment, exists := s.comments[*comment.ParentID]
		if exists {
			parentComment.Children = append(parentComment.Children, comment)
		}

	} else if exists {
		targetPost.Comments = append(targetPost.Comments, comment)

	} else {
		return errors.New("wrong comment data")
	}

	return nil

}
