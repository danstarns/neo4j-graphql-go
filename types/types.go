package types

import "github.com/neo4j/neo4j-go-driver/v4/neo4j"

type Constructor struct {
	TypeDefs string
	Driver   neo4j.Driver
}
