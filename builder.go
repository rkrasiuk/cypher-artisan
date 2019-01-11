package main //builder

import (
	"fmt"
	"strconv"

	"github.com/neo4j/neo4j-go-driver/neo4j"
	art "github.com/rkrasiuk/cypher-artisan/ascii-art"
)

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

	for i, pattern := range patterns {
		query += pattern
		if i != len(patterns)-1 {
			query += `,`
		}
		query += `
		`
	}
	return QueryBuilder{
		query,
	}
}

// Where ...
// WHERE is always part of a MATCH, OPTIONAL MATCH, WITH or START clause
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
func (qb QueryBuilder) Limit(limit int) QueryBuilder {
	return QueryBuilder{
		qb.query + `
		LIMIT ` + strconv.Itoa(limit),
	}
}

// Execute ...
func (qb QueryBuilder) Execute() string {
	return qb.query
}

func main() {
	n := art.NewNode("w1").
		Labels("Person", "Wallet").
		Props(art.Prop{"name", "Theo Gauchoux"}, art.Prop{"age", 22})
	fmt.Println(n)

	// var qb QueryBuilder

	res := NewQueryBuilder().
		Match(
			art.NewNode("w1").Labels("Wallet").String(),
			art.NewNode("w2").Labels("Wallet").String(),
			"p = (w1)-[tx:|*DEPTH*|]-(w2)",
		).
		With(
			`wp, w1, w2,
			w2.address AS recipient,
			|*WHERE*| AS tx2`,
		).
		Return("p, w1, w2, length(p), tx2").Limit(20).Execute()

	fmt.Println("res: \n", res, "\n---\n ")

	res = NewQueryBuilder().
		Match("(a:Person)").
		Where(`a.from = "Sweden"`).
		Return("a").Execute()
	fmt.Println("res: \n", res, "\n---\n ")

	res = NewQueryBuilder().
		Match(
			art.NewNode("w1").Labels("Wallet").String(),
			art.NewNode("w2").Labels("Wallet").String(),
			"p = shortestPath((w1)-[*..]-(w2))",
		).
		Where("w1.address = {w1} AND w2.address IN {w2}").
		Return("p, length(p)").Execute()

	fmt.Println(res)

	edge := art.NewEdge("").Props(art.Prop{"key", "value"}).Path("")
	fmt.Println("edge:", edge)

	var (
		driver  neo4j.Driver
		session neo4j.Session
		err     error
	)

	if driver, err = neo4j.NewDriver(
		"bolt://localhost:7687",
		neo4j.BasicAuth("neo4j", "12345678", ""), // super creative ;)
	); err != nil {
		panic(err)
	}
	defer driver.Close()

	if session, err = driver.Session(neo4j.AccessModeWrite); err != nil {
		panic(err)
	}
	defer session.Close()

	q := NewQueryBuilder().Match(
		art.NewNode("people").Labels("Person").String(),
	).Return("people.name").Limit(10).Execute()
	fmt.Println("---\n", q)
	runAndPrint(session, q)

	q = NewQueryBuilder().
		Match(
			art.NewNode("nineties").Labels("Movie").String(),
		).
		Where("nineties.released >= 1990 AND nineties.released < 2000").
		Return("nineties.title").
		Execute()
	runAndPrint(session, q)
}

func runAndPrint(session neo4j.Session, q string) {
	result, err := session.Run(q, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println("RESULTS")
	for result.Next() {
		fmt.Println(result.Record().GetByIndex(0))
	}
	if err := result.Err(); err != nil {
		panic(err)
	}
}
