package models

type Comment struct {
	ID              int
	PostID          int
	ParentCommentID *string
	AuthorId        int
	Text            string
}
