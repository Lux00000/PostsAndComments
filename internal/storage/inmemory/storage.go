package inmemory

import (
	"errors"
	"github.com/Lux00000/PostsAndComments/internal/models"
	"sync"
)

type Storage struct {
	counter  int
	posts    map[int]*models.Post
	comments map[int]*models.Comment
	mu       sync.RWMutex
}

func NewStorage() *Storage {
	return &Storage{
		counter:  0,
		posts:    make(map[int]*models.Post, 0),
		comments: make(map[int]*models.Comment),
	}
}

func (s *Storage) CreatePost(post *models.Post) (models.Post, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.counter++

	post.ID = s.counter

	s.posts[post.ID] = post

	return *post, nil
}

func (s *Storage) GetAllPosts(limit, offset int) ([]models.Post, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if offset > s.counter {
		return nil, errors.New("offset > postCounter")
	}

	var posts []models.Post
	for i := offset; i <= s.counter && (limit == -1 || i < offset+limit); i++ {
		if post, ok := s.posts[i]; ok {
			posts = append(posts, *post)
		}
	}

	return posts, nil
}

func (s *Storage) GetPostById(id int) (models.Post, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if post, ok := s.posts[id]; ok {
		return *post, nil
	}

	return models.Post{}, errors.New("post not found")
}
