package graph

import (
	"fmt"
	"github.com/Lux00000/PostsAndComments/graph/interfaces"
	"github.com/Lux00000/PostsAndComments/graph/model"
	"sync"
)

type inMemoryPostService struct {
	posts map[int]*model.Post
	mu    sync.RWMutex
}

func NewInMemoryPostService() interfaces.PostService {
	return &inMemoryPostService{
		posts: make(map[int]*model.Post),
	}
}

func (s *inMemoryPostService) CreatePost(post *model.Post) (*model.Post, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	post.ID = len(s.posts) + 1
	s.posts[post.ID] = post
	return post, nil
}

func (s *inMemoryPostService) GetAllPosts(page, pageSize int) ([]*model.Post, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	posts := make([]*model.Post, 0, len(s.posts))
	for _, post := range s.posts {
		posts = append(posts, post)
	}
	return posts, nil
}

func (s *inMemoryPostService) GetPostByID(id int) (*model.Post, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	post, ok := s.posts[id]
	if !ok {
		return nil, fmt.Errorf("post not found")
	}
	return post, nil
}
