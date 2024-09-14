package models

type Post struct {
	ID            int
	Title         string
	Content       string
	AuthorId      int
	AllowComments bool
}
