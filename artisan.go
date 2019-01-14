package artisan

import (
	art "github.com/rkrasiuk/cypher-artisan/ascii-art"
	"github.com/rkrasiuk/cypher-artisan/builder"
)

const (
	// PlainPath represents `--` path
	PlainPath = art.Plain

	// OutgoingPath represents `-->` path
	OutgoingPath = art.Outgoing

	// IncomingPath represents `<--` path
	IncomingPath = art.Incoming

	// BidirectionalPath represents `<-->` path
	BidirectionalPath = art.Bidirectional
)

// Prop ...
type Prop = art.Prop

// Node return a pointer to new Node `()`
// In case no name needed to be provided pass an empty string
func Node(name string) *art.Node {
	return art.NewNode(name)
}

// Edge return a pointer to new Edge `-[]-`
// In case no name needed to be provided pass an empty string
func Edge(name string) *art.Edge {
	return art.NewEdge(name)
}

// QueryBuilder represents a new QueryBuilder for Cypher
func QueryBuilder() builder.QueryBuilder {
	return builder.NewQueryBuilder()
}

// As ...
func As(initial, alias string) string {
	return builder.As(initial, alias)
}

// Assign ...
func Assign(name, pattern string) string {
	return builder.Assign(name, pattern)
}
