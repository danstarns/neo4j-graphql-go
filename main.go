package neo4jGraphQL

import (
	"github.com/danstarns/neo4j-graphql-go/schema"
	"github.com/danstarns/neo4j-graphql-go/types"
	"github.com/graphql-go/graphql"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type NeoSchema struct {
	Schema graphql.Schema
	Driver neo4j.Driver
}

func NewSchema(input types.Constructor) *NeoSchema {
	graphQLSchema := schema.MakeAugmentedSchema(input)

	return &NeoSchema{Schema: graphQLSchema}
}
