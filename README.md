# neo4j-graphql-go

Exploratory Neo4j GraphQL Golang Implementation

## Installation

```
$ go get \
	github.com/neo4j/neo4j-go-driver/v4/neo4j \
	github.com/danstarns/neo4j-graphql-go \
	github.com/graphql-go/handler
```

## Quick Start

```go
package main

import (
	"fmt"
	"net/http"

	neo4jGraphQL "github.com/danstarns/neo4j-graphql-go"
	neo4jGraphQLTypes "github.com/danstarns/neo4j-graphql-go/types"
	"github.com/graphql-go/handler"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

func main() {
	driver, _ := neo4j.NewDriver(
		"bolt://localhost:7687",
		neo4j.BasicAuth("admin", "password", "")
	)

	typeDefs := `
		type Movie {
			id: ID
			title: String
			imdbRating: Int
		}
	`

	neoSchema := neo4jGraphQL.NewSchema(neo4jGraphQLTypes.Constructor{TypeDefs: typeDefs, Driver: driver})

	handler := handler.New(&handler.Config{
		Schema:   &neoSchema.Schema,
		Pretty:   true,
		GraphiQL: true,
	})

	http.Handle("/graphql", handler)
	http.ListenAndServe(":8080", nil)

	fmt.Println("http://localhost:8080")
}
```
