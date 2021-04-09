package neo4jGraphQL

import (
	"github.com/danstarns/neo4j-graphql-go/schema"
	"github.com/danstarns/neo4j-graphql-go/types"
	"github.com/graphql-go/graphql"
)

type NeoSchema struct {
	Schema graphql.Schema
}

func NewSchema(input types.Constructor) *NeoSchema {
	graphQLSchema := schema.MakeAugmentedSchema(input)

	return &NeoSchema{Schema: graphQLSchema}
}
