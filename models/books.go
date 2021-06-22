package models

// Book is a type that represents a book
type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	ISBN   string `json:"isbn"`
}

// Books is a type that represents a bunch of books
type Books []Book
