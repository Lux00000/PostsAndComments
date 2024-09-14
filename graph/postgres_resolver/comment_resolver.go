package postgres_resolver

import (
	"database/sql"
	"github.com/Lux00000/PostsAndComments/graph/interfaces"
	"github.com/Lux00000/PostsAndComments/graph/model"
)

type dbCommentService struct {
	db *sql.DB
}

func NewDBCommentService(db *sql.DB) interfaces.CommentService {
	return &dbCommentService{db: db}
}

func (s *dbCommentService) CreateComment(comment *model.Comment) (*model.Comment, error) {
	query := `INSERT INTO comments (post_id, parent_comment_id, author_id, text) VALUES ($1, $2, $3, $4) RETURNING id`
	err := s.db.QueryRow(query, comment.PostID, comment.ParentCommentID, comment.AuthorID, comment.Text).Scan(&comment.ID)
	if err != nil {
		return nil, err
	}
	return comment, nil
}

func (s *dbCommentService) CommentsSubscription(postID string) (<-chan *model.Comment, error) {
	// Реализация подписки для базы данных
	return make(chan *model.Comment), nil
}
