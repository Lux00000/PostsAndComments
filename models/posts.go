package models

type Post struct {
	ID       int
	Title    string
	Content  string
	Comments []Comment
}

func NewPost(title, content string) *Post {
	return &Post{
		ID:       generateID(),
		Title:    title,
		Content:  content,
		Comments: []Comment{},
	}
}

var postIDCounter = 0

func generateID() int {
	postIDCounter++
	return postIDCounter
}
