package postgres_resolver

import (
	"database/sql"
	"github.com/Lux00000/PostsAndComments/graph/interfaces"
	"github.com/Lux00000/PostsAndComments/graph/model"
)

type dbPostService struct {
	db *sql.DB
}

func NewDBPostService(db *sql.DB) interfaces.PostService {
	return &dbPostService{db: db}
}

func (s *dbPostService) CreatePost(post *model.Post) (*model.Post, error) {
	query := `INSERT INTO posts (title, content, author_id, allow_comments) VALUES ($1, $2, $3, $4) RETURNING id`
	err := s.db.QueryRow(query, post.Title, post.Content, post.AuthorID, post.AllowComments).Scan(&post.ID)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (s *dbPostService) GetAllPosts(page, pageSize int) ([]*model.Post, error) {
	offset := (page - 1) * pageSize
	query := `SELECT id, title, content, author_id, allow_comments FROM posts ORDER BY id OFFSET $1 LIMIT $2`
	rows, err := s.db.Query(query, offset, pageSize)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []*model.Post
	for rows.Next() {
		var post model.Post
		if err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.AuthorID, &post.AllowComments); err != nil {
			return nil, err
		}
		posts = append(posts, &post)
	}
	return posts, nil
}

func (s *dbPostService) GetPostByID(id int) (*model.Post, error) {
	query := `SELECT id, title, content, author_id, allow_comments FROM posts WHERE id = $1`
	var post model.Post
	err := s.db.QueryRow(query, id).Scan(&post.ID, &post.Title, &post.Content, &post.AuthorID, &post.AllowComments)
	if err != nil {
		return nil, err
	}
	return &post, nil
}
