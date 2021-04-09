package schema

import (
	"fmt"

	"github.com/danstarns/neo4j-graphql-go/types"

	"github.com/graphql-go/graphql"
)

func MakeAugmentedSchema(input types.Constructor) graphql.Schema {
	fields := graphql.Fields{
		"hello": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "world", nil
			},
		},
	}

	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}

	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}

	graphQLSchema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		fmt.Println("failed to build schema")
		panic(err)
	}

	return graphQLSchema
}
