package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/Lux00000/post-and-comments/internal/models"
)

type dbCommentPostgres struct {
	db *sql.DB
}

func NewDBCommentPostgres(db *sql.DB) *dbCommentPostgres {
	return &dbCommentPostgres{db: db}
}

func (s *dbCommentPostgres) CreateComment(comment models.Comment) (models.Comment, error) {
	query := `INSERT INTO comments (post_id, parent_comment_id, author_id, text) VALUES ($1, $2, $3, $4) RETURNING id`
	err := s.db.QueryRow(query, comment.PostID, comment.ParentCommentID, comment.AuthorId, comment.Text).Scan(&comment.ID)
	if err != nil {
		return comment, err
	}
	return comment, nil
}

func (s *dbCommentPostgres) GetCommentsByPost(postId int, page *int, pageSize *int) ([]*models.Comment, error) {
	if page == nil || pageSize == nil {
		return nil, errors.New("page and pageSize must be provided")
	}

	offset := (*page - 1) * *pageSize
	limit := *pageSize

	query := `
        SELECT id, post_id, parent_comment_id, author_id, text
        FROM comments
        WHERE post_id = $1
        ORDER BY id
        LIMIT $2 OFFSET $3
    `

	rows, err := s.db.Query(query, postId, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to query comments: %v", err)
	}
	defer rows.Close()

	var comments []*models.Comment
	for rows.Next() {
		var comment models.Comment
		if err := rows.Scan(&comment.ID, &comment.PostID, &comment.ParentCommentID, &comment.AuthorId, &comment.Text); err != nil {
			return nil, fmt.Errorf("failed to scan comment row: %v", err)
		}
		comments = append(comments, &comment)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error occurred during iteration: %v", err)
	}

	return comments, nil
}

func (s *dbCommentPostgres) GetChildrenOfComment(commentId int) ([]*models.Comment, error) {
	query := `SELECT * FROM comments WHERE parent_comment_id = $1`
	rows, err := s.db.Query(query, commentId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []*models.Comment
	for rows.Next() {
		var comment models.Comment
		if err := rows.Scan(&comment.ID, &comment.PostID, &comment.ParentCommentID, &comment.AuthorId, &comment.Text); err != nil {
			return nil, err
		}
		comments = append(comments, &comment)
	}
	return comments, nil
}
