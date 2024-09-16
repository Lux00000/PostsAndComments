package postgres

import (
	"database/sql"
	"fmt"
	"github.com/Lux00000/post-and-comments/internal/models"
)

type dbPostPostgres struct {
	db *sql.DB
}

func NewDBPostPostgres(db *sql.DB) *dbPostPostgres {
	return &dbPostPostgres{db: db}
}

func (s *dbPostPostgres) CreatePost(post models.Post) (models.Post, error) {
	query := `INSERT INTO Posts (title, content, author_id, allow_comments) 
			  VALUES ($1, $2, $3, $4) RETURNING id`

	if post.Title == "" || post.Content == "" || post.AuthorId == 0 {
		return post, fmt.Errorf("failed to create post: missing required fields")
	}

	err := s.db.QueryRow(query, post.Title, post.Content, post.AuthorId, post.AllowComments).Scan(&post.ID)
	if err != nil {
		return post, fmt.Errorf("failed to create post: %v", err)
	}

	return post, nil
}

func (s *dbPostPostgres) GetAllPosts(page, pageSize int) ([]models.Post, error) {
	offset := (page - 1) * pageSize
	limit := pageSize
	query := `SELECT id, title, content, author_id, allow_comments FROM posts ORDER BY id LIMIT $1 OFFSET $2`
	rows, err := s.db.Query(query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to query posts: %v", err)
	}
	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var post models.Post
		if err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.AuthorId, &post.AllowComments); err != nil {
			return nil, fmt.Errorf("failed to scan post row: %v", err)
		}
		posts = append(posts, post)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error occurred during iteration: %v", err)
	}

	return posts, nil
}

func (s *dbPostPostgres) GetPostById(id int) (*models.Post, error) {
	query := `SELECT * FROM Posts WHERE id = $1`
	var post models.Post
	err := s.db.QueryRow(query, id).Scan(&post.ID, &post.Title, &post.Content, &post.AuthorId, &post.AllowComments)
	if err != nil {
		return nil, err
	}
	return &post, nil
}
