package dbmysql

import (
	"context"
	"database/sql"
	"examples/sample/book-manager-service/db"
	"examples/sample/book-manager-service/model"
	"log"
	"time"
)

type BookDBImpl struct {
	db *sql.DB
}

func initBookDBImpl() db.BookDB {
	db, err := sql.Open("mysql", "root:12345678x@X@/book")
	if err != nil {
		panic(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	err = db.PingContext(ctx)
	if err != nil {
		log.Fatal(err)
	}

	instance := new(BookDBImpl)
	instance.db = db
	return instance
}

func (b BookDBImpl) GetBookList() ([]model.Book, error) {
	var bookList []model.Book
	stmt, err := b.db.Prepare("SELECT * FROM book")
	if err != nil {
		return []model.Book{}, err
	}

	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	resp, err := stmt.QueryContext(ctx)
	if err != nil {
		return []model.Book{}, err
	}

	for resp.Next() {
		var book model.Book
		if err = resp.Scan(&book.ID, &book.Name, &book.Author); err != nil {
			return []model.Book{}, err
		}
		bookList = append(bookList, book)
	}
	return bookList, nil
}
