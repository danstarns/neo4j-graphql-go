package schema

import (
	"encoding/json"
	"fmt"

	"github.com/danstarns/neo4j-graphql-go/node"
	"github.com/danstarns/neo4j-graphql-go/translate"
	"github.com/danstarns/neo4j-graphql-go/types"
	"github.com/danstarns/neo4j-graphql-go/util"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
	"github.com/graphql-go/graphql/language/parser"
)

func MakeAugmentedSchema(input types.Constructor) graphql.Schema {
	document, err := parser.Parse(parser.ParseParams{Source: input.TypeDefs})
	if err != nil {
		fmt.Println("Cannot parse typeDefs")
		panic(err)
	}

	var nodes []node.Node

	for _, value := range document.Definitions {
		if value.GetKind() == "ObjectDefinition" {
			var p2 ast.ObjectDefinition
			j, _ := json.Marshal(value)
			json.Unmarshal(j, &p2)
			nodes = append(nodes, node.NewNode(p2))
		}
	}

	queryFields := graphql.Fields{}
	var types []graphql.Type

	for _, n := range nodes {
		var fields = graphql.Fields{}
		for _, prim := range n.PrimitiveFields {
			f := graphql.Field{Type: graphql.String}
			fields[prim.FieldName] = &f
		}

		o := graphql.NewObject(graphql.ObjectConfig{Name: n.Name, Fields: fields})
		types = append(types, o)

		queryFields[n.Name] = &graphql.Field{
			Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull((o)))),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				readInput := translate.TranslateReadInput{
					Node:          n,
					ResolveParams: p,
				}

				cypher, params := translate.TranslateRead(readInput)

				result := util.Execute(util.ExecuteInput{
					Driver:     input.Driver,
					AccessMode: neo4j.AccessModeRead,
					Cypher:     cypher,
					Params:     params,
				})

				return result, nil
			},
		}

	}

	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: queryFields}

	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery), Types: types}

	graphQLSchema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		fmt.Println("failed to build schema")
		panic(err)
	}

	return graphQLSchema
}
