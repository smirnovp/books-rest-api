package models

// Book is a type that represents a book
type Book struct {
	ID     int64  `json:"id" form:"id" xml:"id" binding:"-"`
	Title  string `json:"title" form:"title" xml:"title" binding:"required"`
	Author string `json:"author" form:"author" xml:"author" binding:"required"`
	ISBN   string `json:"isbn" form:"isbn" xml:"isbn" binding:"-"`
}

// Books is a type that represents a bunch of books
type Books []Book
