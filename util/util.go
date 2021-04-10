package util

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

func deleteKey(keyToDelete string, val reflect.Value) reflect.Value {
	// Indirect through pointers and interfaces
	for val.Kind() == reflect.Ptr || val.Kind() == reflect.Interface {
		val = val.Elem()
	}

	switch val.Kind() {
	case reflect.Array, reflect.Slice:
		for i := 0; i < val.Len(); i++ {
			deleteKey(keyToDelete, val.Index(i))
		}
	case reflect.Map:
		for _, k := range val.MapKeys() {
			if k.String() == keyToDelete {
				delete(val.Interface().(map[string]interface{}), k.String())
				continue
			}
			deleteKey(keyToDelete, val.MapIndex(k))
		}
	default:
		// Do we need that case?
	}
	return val
}

func PrintDocumentWithoutLoc(document interface{}) {
	var p2 map[string]interface{}
	j, _ := json.MarshalIndent(document, "", " ")
	json.Unmarshal(j, &p2)
	deleteKey("Loc", reflect.ValueOf(p2))
	jj, _ := json.MarshalIndent(p2, "", " ")
	fmt.Println(string(jj))
}

type ExecuteInput struct {
	Driver     neo4j.Driver
	Cypher     string
	Params     map[string]interface{}
	AccessMode neo4j.AccessMode
}

func Execute(input ExecuteInput) interface{} {
	session := input.Driver.NewSession(neo4j.SessionConfig{AccessMode: input.AccessMode})
	defer session.Close()

	res, err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		records, err := tx.Run(
			input.Cypher,
			input.Params,
		)

		if err != nil {
			return nil, err
		}

		var results []interface{}

		for records.Next() {
			record := records.Record()
			results = append(results, record.Values[0])
		}

		return results, nil
	})

	if err != nil {
		fmt.Println("failed to execute")
		panic(err)
	}

	return res
}
