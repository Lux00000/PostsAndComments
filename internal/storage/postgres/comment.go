package postgres

import (
	"database/sql"
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
	offset := 0
	limit := -1
	if page != nil && pageSize != nil {
		offset = (*page - 1) * *pageSize
		limit = *pageSize
	}

	query := `SELECT * FROM comments WHERE post_id = $1 ORDER BY id OFFSET $2`
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
		if err := rows.Scan(&comment.ID, &comment.PostID, &comment.ParentCommentID, &comment.AuthorId, &comment.Text); err != nil {
			return nil, err
		}
		comments = append(comments, &comment)
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
