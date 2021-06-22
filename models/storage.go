package models

// IStorage is an interface to the storage
type IStorage interface {
	Add(Book)
	GetAll() Books
}
