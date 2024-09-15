package service

import (
	"errors"
	"github.com/Lux00000/PostsAndComments/internal/models"
	"github.com/Lux00000/PostsAndComments/internal/storage"
)

type CommentsService struct {
	com  storage.Comments
	post PostGet
}
type PostGet interface {
	GetPostById(id int) (models.Post, error)
}

func NewCommentsService(com storage.Comments, get PostGet) *CommentsService {
	return &CommentsService{com: com, post: get}
}

func (c *CommentsService) CreateComment(comment models.Comment) (models.Comment, error) {
	if comment.AuthorId == 0 {
		return models.Comment{}, errors.New("AuthorId is zero")
	}

	if len(comment.Text) >= 2000 {
		return models.Comment{}, errors.New("Content is too long")
	}

	if comment.PostID <= 0 {
		return models.Comment{}, errors.New("PostId is zero")
	}

	post, err := c.post.GetPostById(comment.PostID)
	if err != nil {
		return models.Comment{}, errors.New("CommentService error GetPostById")
	}

	if !post.AllowComments {
		return models.Comment{}, errors.New("Comments are closed")

	}

	newComment, err := c.com.CreateComment(comment)
	if err != nil {
		return models.Comment{}, errors.New("CommentService error CreateComment")
	}

	return newComment, nil
}

func (c CommentsService) GetCommentsByPost(postId int, page *int, pageSize *int) ([]*models.Comment, error) {

	if postId <= 0 {
		return nil, errors.New("PostId is zero")
	}

	if *page < 0 {
		return nil, errors.New("Page is zero")
	}

	if *pageSize < 0 {
		return nil, errors.New("PageSize is zero")
	}

	offset := (*page - 1) * *pageSize
	limit := *pageSize

	comments, err := c.com.GetCommentsByPost(postId, limit, offset)
	if err != nil {
		return nil, errors.New("CommentService error GetCommentsByPost")
	}

	return comments, nil
}

func (c CommentsService) GetChildrenOfComment(commentId int) ([]*models.Comment, error) {

	if commentId <= 0 {
		return nil, errors.New("CommentId is zero")
	}

	comments, err := c.com.GetChildrenOfComment(commentId)
	if err != nil {
		return nil, errors.New("CommentService error GetChildrenOfComment")
	}
	return comments, nil

}
