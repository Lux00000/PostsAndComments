package postgres

import (
	"database/sql"
	"github.com/Lux00000/PostsAndComments/internal/models"
)

type dbPostPostgres struct {
	db *sql.DB
}

func NewDBPostPostgres(db *sql.DB) *dbPostPostgres {
	return &dbPostPostgres{db: db}
}

func (s *dbPostPostgres) CreatePost(post *models.Post) (*models.Post, error) {
	query := `INSERT INTO Posts (title, content, author_id, allow_comments) VALUES ($1, $2, $3, $4) RETURNING id`
	err := s.db.QueryRow(query, post.Title, post.Content, post.AuthorId, post.AllowComments).Scan(&post.ID)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (s *dbPostPostgres) GetAllPosts(page, pageSize int) ([]*models.Post, error) {
	offset := (page - 1) * pageSize
	query := `SELECT * FROM Posts ORDER BY id OFFSET $1 LIMIT $2`
	rows, err := s.db.Query(query, offset, pageSize)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []*models.Post
	for rows.Next() {
		var post models.Post
		if err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.AuthorId, &post.AllowComments); err != nil {
			return nil, err
		}
		posts = append(posts, &post)
	}
	return posts, nil
}

func (s *dbPostPostgres) GetPostByID(id int) (*models.Post, error) {
	query := `SELECT * FROM Posts WHERE id = $1`
	var post models.Post
	err := s.db.QueryRow(query, id).Scan(&post.ID, &post.Title, &post.Content, &post.AuthorId, &post.AllowComments)
	if err != nil {
		return nil, err
	}
	return &post, nil
}
