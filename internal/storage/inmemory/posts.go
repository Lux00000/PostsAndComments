package inmemory

import (
	"errors"
	"fmt"
	"github.com/Lux00000/post-and-comments/internal/models"
	"sync"
)

type InMemoryPost struct {
	posts map[int]*models.Post
	mu    sync.RWMutex
}

func NewInMemoryPost() *InMemoryPost {
	return &InMemoryPost{
		posts: make(map[int]*models.Post, 0),
	}
}

func (s *InMemoryPost) CreatePost(post models.Post) (models.Post, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	post.ID = len(s.posts) + 1
	s.posts[post.ID] = &post
	return post, nil
}

func (s *InMemoryPost) GetAllPosts(limit, offset int) ([]models.Post, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if offset > len(s.posts) {
		return nil, errors.New("offset out of range")
	}
	if offset < 0 || limit < 0 {
		return nil, errors.New("limit out of range")
	}
	posts := make([]models.Post, 0, len(s.posts))
	for _, post := range s.posts {
		posts = append(posts, *post)
	}
	return posts, nil
}

func (s *InMemoryPost) GetPostById(id int) (*models.Post, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	post, ok := s.posts[id]
	if !ok {
		return nil, fmt.Errorf("post not found")
	}
	return post, nil
}
