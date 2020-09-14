package resolver

import (
	"examples/sample/gateway/resolver/query"
	"github.com/graphql-go/graphql"
	"log"
)

func Init() *graphql.Schema {
	schema, err := graphql.NewSchema(
		graphql.SchemaConfig{
			Query: query.Query(),
		},
	)

	if err != nil {
		log.Fatalf("schema: Init: %v", err)
		return nil
	}

	return &schema
}
