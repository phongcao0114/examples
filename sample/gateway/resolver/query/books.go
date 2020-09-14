package query

import (
	"encoding/json"
	"examples/sample/book-manager-service/model"
	"examples/sample/gateway/datatype"
	"fmt"
	"github.com/graphql-go/graphql"
	"net/http"
)

var Books = &graphql.Field{
	Type: datatype.ListBook,
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		resp, err := http.Get("http://localhost:8080/api/v1/books")
		if err != nil {
			return []model.Book{}, err
		}
		defer resp.Body.Close()

		var books []model.Book
		json.NewDecoder(resp.Body).Decode(&books)
		fmt.Println(books)
		return books, nil
	},
}
