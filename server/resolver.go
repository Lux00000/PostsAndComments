package server

import (
	"context"
	"database/sql"
	"github.com/Lux00000/PostsAndComments/internal/models"
	"github.com/Lux00000/PostsAndComments/internal/storage/inmemory"
	"github.com/Lux00000/PostsAndComments/internal/storage/postgres"
)

type Resolver struct {
	postRepo    PostRepository
	commentRepo CommentRepository
	CommentsObservers Observers
}
}

type PostRepository interface {
	CreatePost(post *models.Post) (*models.Post, error)
	GetAllPosts(limit, offset int) ([]*models.Post, error)
	GetPostByID(id int) (*models.Post, error)
}

type CommentRepository interface {
	CreateComment(comment *models.Comment) (*models.Comment, error)
	GetCommentsByPost(postId, limit, offset int) ([]*models.Comment, error)
	GetChildrenComment(commentId int) ([]*models.Comment, error)
}

func NewResolver(db *sql.DB) *Resolver {
	var postRepo PostRepository
	var commentRepo CommentRepository

	if db == nil {
		postRepo = inmemory.NewInMemoryPost()
		commentRepo = inmemory.NewInMemoryComment()
	} else {
		postRepo = postgres.NewDBPostPostgres(db)
		commentRepo = postgres.NewDBCommentPostgres(db)
	}

	return &Resolver{
		postRepo:    postRepo,
		commentRepo: commentRepo,
	}
}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreatePost(ctx context.Context, input models.PostInput) (*models.Post, error) {
	post := &models.Post{
		Title:         input.Title,
		Content:       input.Content,
		AuthorId:      input.AuthorID,
		AllowComments: input.AllowComments,
	}
	return r.postRepo.CreatePost(post)
}

func (r *mutationResolver) CreateComment(ctx context.Context, input models.CommentInput) (*models.Comment, error) {
	comment := &models.Comment{
		PostID:          input.PostID,
		ParentCommentID: input.ParentCommentID,
		AuthorId:        input.AuthorID,
		Text:            input.Text,
	}
	return r.commentRepo.CreateComment(comment)
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Posts(ctx context.Context, limit *int, offset *int) ([]*models.Post, error) {
	limitVal := 10
	offsetVal := 0
	if limit != nil {
		limitVal = *limit
	}
	if offset != nil {
		offsetVal = *offset
	}
	return r.postRepo.GetPosts(limitVal, offsetVal)
}

func (r *queryResolver) CommentsByPost(ctx context.Context, postID int, limit *int, offset *int) ([]*models.Comment, error) {
	limitVal := 10
	offsetVal := 0
	if limit != nil {
		limitVal = *limit
	}
	if offset != nil {
		offsetVal = *offset
	}
	return r.commentRepo.GetCommentsByPost(postID, limitVal, offsetVal)
}

func (r *queryResolver) ChildrenComments(ctx context.Context, commentID int) ([]*models.Comment, error) {
	return r.commentRepo.GetChildrenComment(commentID)
}
