package db

import "examples/sample/book-manager-service/model"

type BookDB interface {
	GetBookList() ([]model.Book, error)
}
