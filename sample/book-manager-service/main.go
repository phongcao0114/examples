package main

import (
	"examples/sample/book-manager-service/dbmysql"
	"examples/sample/book-manager-service/handler"
	"examples/sample/book-manager-service/service"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := dbmysql.Init()
	service := service.NewBookManagerService(&db)
	routersInit := handler.InitRouter(service)
	server := &http.Server{
		Addr:    ":8080",
		Handler: routersInit,
	}

	server.ListenAndServe()
}
