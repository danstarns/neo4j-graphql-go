package node

import (
	"github.com/graphql-go/graphql/language/ast"
)

type PrimitiveField struct {
	FieldName string
	Kind      string
}

type Node struct {
	Name            string
	Kind            string
	PrimitiveFields []PrimitiveField
}

func NewNode(definition ast.ObjectDefinition) Node {
	name := definition.Name.Value
	kind := definition.GetKind()
	var primitiveFields []PrimitiveField

	for _, f := range definition.Fields {
		primitive := PrimitiveField{FieldName: f.Name.Value, Kind: f.GetKind()}
		primitiveFields = append(primitiveFields, primitive)
	}

	return Node{Name: name, PrimitiveFields: primitiveFields, Kind: kind}
}
