package main //artisan

import (
	"fmt"

	"github.com/neo4j/neo4j-go-driver/neo4j"
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

/////

func main() {
	n := art.NewNode("w1").
		Labels("Person", "Wallet").
		Props(Prop{"name", "Theo Gauchoux"}, Prop{"age", 22})
	fmt.Println(n)

	// var qb QueryBuilder

	res := QueryBuilder().
		Match(
			Node("w1").Labels("Wallet").String(),
			Node("w2").Labels("Wallet").String(),
			"p = (w1)-[tx:|*DEPTH*|]-(w2)",
		).
		With(
			`wp, w1, w2,
			w2.address AS recipient,
			|*WHERE*| AS tx2`,
		).
		Return("p, w1, w2, length(p), tx2").Limit(20).Execute()

	fmt.Println("res: \n", res, "\n---\n ")

	res = QueryBuilder().
		Match("(a:Person)").
		Where(`a.from = "Sweden"`).
		Return("a").Execute()
	fmt.Println("res: \n", res, "\n---\n ")

	res = QueryBuilder().
		Match(
			Node("w1").Labels("Wallet").String(),
			Node("w2").Labels("Wallet").String(),
			"p = shortestPath((w1)-[*..]-(w2))",
		).
		Where("w1.address = {w1} AND w2.address IN {w2}").
		Return("p, length(p)").Execute()

	fmt.Println(res)

	edge := Edge("").Props(Prop{"key", "value"}).Path("")
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

	q := QueryBuilder().Match(
		Node("people").Labels("Person").String(),
	).Return("people.name").Limit(10).Execute()
	fmt.Println("---\n", q)
	runAndPrint(session, q)

	q = QueryBuilder().
		Match(
			Node("nineties").Labels("Movie").String(),
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
