package main //builder

import (
	"errors"
	"fmt"
	"log"

	"github.com/rkrasiuk/cypher-artisan/node"
)

// Read query structure
// [MATCH WHERE]
// [OPTIONAL MATCH WHERE]
// [WITH [ORDER BY] [SKIP] [LIMIT]]
// RETURN [ORDER BY] [SKIP] [LIMIT]

// Stringer ...
type Stringer string

func (s Stringer) String() string { return string(s) }

// NewQueryBuilder ...
func NewQueryBuilder() QueryBuilder {
	return QueryBuilder{}
}

// QueryBuilder ...
type QueryBuilder struct {
	query string
}

// Builder ...
type Builder interface {
}

// Match ...
func (qb QueryBuilder) Match(patterns ...string) QueryBuilder {
	query := qb.query + `
		MATCH 
			`

	for _, pattern := range patterns {
		query += pattern + `,
			`
	}
	return QueryBuilder{
		query,
	}
}

// Where ...
func (qb QueryBuilder) Where(whereClause string) QueryBuilder {
	return QueryBuilder{
		qb.query + `
		WHERE 
			` + whereClause,
	}
}

// With ...
func (qb QueryBuilder) With(withClause string) QueryBuilder {
	return QueryBuilder{
		qb.query + `
		WITH 
			` + withClause,
	}
}

// Return ...
func (qb QueryBuilder) Return(returnClause string) QueryBuilder {
	return QueryBuilder{
		qb.query + `
		RETURN 
			` + returnClause,
	}
}

// OrderBy ...
func (qb QueryBuilder) OrderBy(orderByClause string) QueryBuilder {
	return QueryBuilder{
		qb.query + `
		ORDER BY 
		` + orderByClause,
	}
}

// Limit ...
func (qb QueryBuilder) Limit(limitClause string) QueryBuilder {
	return QueryBuilder{
		qb.query + `
		LIMIT ` + limitClause,
	}
}

// Execute ...
func (qb QueryBuilder) Execute() (string, error) {
	success := true

	if success {
		return qb.query, nil
	}
	return "", errors.New("Failed to execute query")
}

func main() {
	n := node.NewNode("w1").Labels("Person", "Wallet").Props(node.Prop{"name", "Theo Gauchoux"}, node.Prop{"age", 22})
	fmt.Println(n)

	// var qb QueryBuilder

	res, err := NewQueryBuilder().
		Match(
			node.NewNode("w1").Labels("Wallet").Props(node.Prop{"address", "{w1}"}).String(),
			node.NewNode("w2").Labels("Wallet").String(),
			"p = (w1)-[tx:|*DEPTH*|]-(w2)",
		).
		With(
			`wp, w1, w2,
			w2.address AS recipient,
			|*WHERE*| AS tx2`,
		).
		Return("p, w1, w2, length(p), tx2").Limit("20").Execute()

	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("res: \n", res, "\n---\n ")

	res, err = NewQueryBuilder().
		Match("(a:Person)").
		Where(`a.from = "Sweden"`).
		Return("a").Execute()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("res: \n", res, "\n---\n ")

	res, err = NewQueryBuilder().
		Match(
			node.NewNode("w1").Labels("Wallet").String(),
			node.NewNode("w2").Labels("Wallet").String(),
			"p = shortestPath((w1)-[*..]-(w2))",
		).
		Where("w1.address = {w1} AND w2.address IN {w2}").
		Return("p, length(p)").Execute()

	fmt.Println(res)
}
