package models

type Comment struct {
	ID     int
	Author string
	Test   string
}

func NewComment(author, text string) *Comment {
	return &Comment{
		ID:     gennerateCommentID(),
		Author: author,
		Test:   text,
	}
}

var commentIDCounter = 0

func gennerateCommentID() int {
	commentIDCounter++
	return commentIDCounter
}
