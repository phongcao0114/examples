package service

import (
	"examples/sample/book-manager-service/db"
)

type BookManagerService interface {
	GetBookList() (interface{}, error)
}

type BookManagerServiceImpl struct {
	*db.MasterDB
}

func NewBookManagerService(db *db.MasterDB) BookManagerService {
	return &BookManagerServiceImpl{db}
}
