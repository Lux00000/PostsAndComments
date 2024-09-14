package inmemory

import (
	"errors"
	"github.com/Lux00000/PostsAndComments/internal/models"
	"sync"
)

type PostsInMemory struct {
	postCounter int
	posts       []models.Post
	mu          sync.RWMutex
}

func NewPostsInMemory(count int) *PostsInMemory {
	return &PostsInMemory{
		postCounter: 0,
		posts:       make([]models.Post, 0),
	}
}

func (s *PostsInMemory) CreatePost(post *models.Post) (models.Post, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.postCounter++

	post.ID = s.postCounter

	s.posts = append(s.posts, *post)

	return *post, nil
}

func (s *PostsInMemory) GetAllPosts(limit, offset int) ([]models.Post, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if offset > s.postCounter {
		return nil, errors.New("offset > postCounter")
	}

	if offset+limit > s.postCounter || limit == -1 {
		return s.posts[offset:], nil
	}

	return s.posts[offset : offset+limit], nil
}

func (s *PostsInMemory) GetPostById(id int) (models.Post, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if id > s.postCounter || id <= 0 {
		return models.Post{}, errors.New("post not found")
	}

	return s.posts[id-1], nil
}
