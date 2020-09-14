package handler

import (
	"examples/sample/book-manager-service/service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter(s service.BookManagerService) *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	apiv1 := r.Group("/api/v1")
	apiv1.GET("/books", func(c *gin.Context) {
		resp, err := s.GetBookList()
		if err != nil {
			c.JSON(http.StatusNotFound, err)
		} else {
			fmt.Printf("%+v\n", resp)
			c.JSON(http.StatusOK, resp)
		}
	})

	return r
}
