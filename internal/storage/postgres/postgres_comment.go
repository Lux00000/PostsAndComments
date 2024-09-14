package postgres

import (
	"database/sql"
	"github.com/Lux00000/PostsAndComments/internal/models"
)

type dbCommentPostgres struct {
	db *sql.DB
}

func NewDBCommentPostgres(db *sql.DB) *dbCommentPostgres {
	return &dbCommentPostgres{db: db}
}

func (s *dbCommentPostgres) CreateComment(comment *models.Comment) (*models.Comment, error) {
	query := `INSERT INTO comments (post_id, parent_comment_id, author_id, text) VALUES ($1, $2, $3, $4) RETURNING id`
	err := s.db.QueryRow(query, comment.PostID, comment.ParentCommentID, comment.AuthorId, comment.Text).Scan(&comment.ID)
	if err != nil {
		return nil, err
	}
	return comment, nil
}

func (s *dbCommentPostgres) GetCommentsByPost(postId, limit, offset int) ([]*models.Comment, error) {
	query := `
		SELECT * 
		FROM Comments 
		WHERE post_id = $1 AND parent_comment_id IS NULL 
		ORDER BY id 
		OFFSET $2`

	args := []interface{}{postId, offset}

	if limit >= 0 {
		query += " LIMIT $3"
		args = append(args, limit)
	}

	rows, err := s.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []*models.Comment

	for rows.Next() {
		var comment models.Comment
		err := rows.Scan(&comment.ID, &comment.PostID, &comment.ParentCommentID, &comment.AuthorId, &comment.Text)
		if err != nil {
			return nil, err
		}
		comments = append(comments, &comment)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return comments, nil
}

func (s *dbCommentPostgres) GetChildrenComment(commentId int) ([]*models.Comment, error) {
	query := `SELECT * FROM comments WHERE parent_comment_id = $1`

	args := []interface{}{commentId}

	rows, err := s.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []*models.Comment

	for rows.Next() {
		var comment models.Comment
		err := rows.Scan(&comment.ID, &comment.PostID, &comment.ParentCommentID, &comment.AuthorId, &comment.Text)
		if err != nil {
			return nil, err
		}
		comments = append(comments, &comment)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return comments, nil
}
