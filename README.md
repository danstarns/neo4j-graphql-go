# neo4j-graphql-go

Exploratory Neo4j GraphQL Golang Implementation

## Installation

```
$ go get github.com/neo4j/neo4j-go-driver/neo4j github.com/danstarns/neo4j-graphql-go github.com/graphql-go/graphql
```

## Quick Start

```go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	neo4jGraphQL "github.com/danstarns/neo4j-graphql-go"
	"github.com/graphql-go/graphql"
	"github.com/neo4j/neo4j-go-driver/neo4j"
)

type postData struct {
	Query     string                 `json:"query"`
	Operation string                 `json:"operation"`
	Variables map[string]interface{} `json:"variables"`
}

func main() {
	driver, _ := neo4j.NewDriver("bolt://localhost:7687", neo4j.BasicAuth("admin", "password", ""))

	typeDefs := `
		type Movie {
			id: ID
			title: String
			imdbRating: Int
		}
	`

	neoSchema := neo4jGraphQL.NewSchema(&neo4jGraphQL.Constructor{typeDefs, driver})

	http.HandleFunc("/graphql", func(w http.ResponseWriter, req *http.Request) {
		var p postData
		if err := json.NewDecoder(req.Body).Decode(&p); err != nil {
			w.WriteHeader(400)
			return
		}

		result := graphql.Do(graphql.Params{
			Context:        req.Context(),
			Schema:         neoSchema.Schema,
			RequestString:  p.Query,
			VariableValues: p.Variables,
			OperationName:  p.Operation,
		})

		if err := json.NewEncoder(w).Encode(result); err != nil {
			fmt.Printf("could not write result to response: %s", err)
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
```
