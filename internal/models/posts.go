package models

type Post struct {
	ID            int
	Title         string
	Content       string
	Author        string
	Comments      []*Comment
	AllowComments bool
}

func NewPost(title, content, author string, allowComments bool) *Post {
	return &Post{
		ID:            generatePostID(),
		Title:         title,
		Content:       content,
		Author:        author,
		AllowComments: allowComments,
		Comments:      []*Comment{},
	}
}

var postIDCounter = 0

func generatePostID() int {
	postIDCounter++
	return postIDCounter
}
