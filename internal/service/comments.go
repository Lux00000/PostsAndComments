package service

import (
	"errors"
	"fmt"
	"github.com/Lux00000/post-and-comments/internal/models"
	"github.com/Lux00000/post-and-comments/internal/storage"
)

type CommentsService struct {
	comments storage.Comments
	post     storage.Posts
}

func NewCommentsService(com storage.Comments, posts storage.Posts) *CommentsService {
	return &CommentsService{comments: com, post: posts}
}

func (c *CommentsService) CreateComment(comment models.Comment) (models.Comment, error) {
	if comment.AuthorId == 0 {
		return models.Comment{}, fmt.Errorf("AuthorId is zero")
	}

	if len(comment.Text) >= 2000 {
		return models.Comment{}, errors.New("Content is too long")
	}

	if comment.PostID <= 0 {
		return models.Comment{}, errors.New("PostId is less or equal than zero")
	}

	post, err := c.post.GetPostById(comment.PostID)
	if err != nil {
		return models.Comment{}, fmt.Errorf("CommentService error GetPostById: %w", err)
		// TODO : Везде пробрасывать ошибку наверх, как в примере выше.
	}

	if !post.AllowComments {
		return models.Comment{}, errors.New("Comments are closed")

	}

	newComment, err := c.comments.CreateComment(comment)
	if err != nil {
		return models.Comment{}, fmt.Errorf("CommentService error CreateComment: %w", err)
	}

	return newComment, nil
}

func (c CommentsService) GetCommentsByPost(postId int, page *int, pageSize *int) ([]*models.Comment, error) {

	comments, err := c.comments.GetCommentsByPost(postId, page, pageSize)
	if err != nil {
		return nil, fmt.Errorf("CommentService error GetChildrenOfComment: %w", err)
	}

	return comments, nil
}

func (c CommentsService) GetChildrenOfComment(commentId int) ([]*models.Comment, error) {
	if commentId <= 0 {
		return nil, errors.New("CommentId is zero")
	}

	comments, err := c.comments.GetChildrenOfComment(commentId)
	if err != nil {
		return nil, errors.New("CommentService error GetChildrenOfComment")
	}
	return comments, nil
}
