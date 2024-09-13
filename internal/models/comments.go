package models

import "errors"

type Comment struct {
	ID       int
	PostID   int
	ParentID *int
	Author   string
	Text     string
	Children []*Comment
}

func NewComment(postID int, parentID *int, author, text string) (*Comment, error) {
	if len(text) < 2000 {
		return &Comment{
			ID:       generateCommentID(),
			PostID:   postID,
			ParentID: parentID,
			Author:   author,
			Text:     text,
			Children: []*Comment{},
		}, nil
	} else {
		return nil, errors.New("Validation Exception")
	}

}

var commentIDCounter = 0

func generateCommentID() int {
	commentIDCounter++
	return commentIDCounter
}
