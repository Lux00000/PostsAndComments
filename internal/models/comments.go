package models

type Comment struct {
	ID              int
	PostID          int
	ParentCommentID *int
	AuthorId        int
	Text            string
}
