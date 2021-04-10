package translate

import (
	"fmt"
	"strings"

	"github.com/danstarns/neo4j-graphql-go/node"
	"github.com/danstarns/neo4j-graphql-go/util"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
)

type TranslateReadInput struct {
	Node          node.Node
	ResolveParams graphql.ResolveParams
}

type projectionInput struct {
	node          node.Node
	varName       string
	rootSelection *ast.Field
	params        *interface{}
}

func createProjection(input projectionInput) string {
	var strs []string

	for _, selection := range input.rootSelection.GetSelectionSet().Selections {

		f := selection.(*ast.Field)

		strs = append(strs, fmt.Sprintf("%s:%s", f.Name.Value, f.Name.Value))
	}

	joined := strings.Join(strs, ", ")

	return fmt.Sprintf("{ %s }", joined)
}

func TranslateRead(input TranslateReadInput) (string, interface{}) {
	var params interface{}

	rootSelection := input.ResolveParams.Info.FieldASTs[0]

	util.PrintDocumentWithoutLoc(rootSelection)

	varName := "this"
	var strs []string
	strs = append(strs, fmt.Sprintf("MATCH (%s:%s)", varName, input.Node.Name))

	pInput := projectionInput{node: input.Node, varName: varName, rootSelection: rootSelection, params: &params}

	proj := createProjection(pInput)

	strs = append(strs, fmt.Sprintf("RETURN %s %s AS %s", varName, proj, varName))

	return strings.Join(strs, "\n"), params
}
