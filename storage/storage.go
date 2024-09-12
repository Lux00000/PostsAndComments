package storage

import "github.com/Lux00000/PostsAndComments/models"

type Storage struct {
	posts []*models.Post
}

func NewStorage() *Storage {
	return &Storage{
		posts: []*models.Post{},
	}
}

func (s *Storage) AddPost(post *models.Post) {
	s.posts = append(s.posts, post)
}

// AddCommentToPost добавляет комментарий к посту.
func (s *Storage) AddCommentToPost(postID int, comment *models.Comment) {
	for _, post := range s.posts {
		if post.ID == postID {
			post.Comments = append(post.Comments, *comment)
			break
		}
	}
}
