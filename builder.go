package main //builder

import (
	"errors"
	"fmt"
	"log"
	"strings"
)

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

// Node ...
type Node struct {
	name   string
	labels []string
	props  map[string]interface{}
}

// Prop ...
type Prop struct {
	key   string
	value interface{}
}

// NewNode ...
func NewNode(name string) *Node {
	return &Node{
		name,
		[]string{},
		make(map[string]interface{}),
	}
}

// Labels ...
func (n *Node) Labels(labels ...string) *Node {
	for _, label := range labels {
		n.labels = append(n.labels, label)
	}
	return n
}

// Props ...
func (n *Node) Props(props ...Prop) *Node {
	for _, prop := range props {
		n.props[prop.key] = prop.value
	}
	return n
}

// ToString ...
func (n Node) String() string {
	var labels, props string

	if len(n.labels) > 0 {
		labels = fmt.Sprintf(":%v", strings.Join(n.labels, ":"))
	}

	if len(n.props) > 0 {
		var propsArr []string
		for key, prop := range n.props {
			switch prop.(type) {
			case string:
				propsArr = append(propsArr, fmt.Sprintf("%v: '%v'", key, prop))
			default:
				propsArr = append(propsArr, fmt.Sprintf("%v: %v", key, prop))
			}
		}
		props = fmt.Sprintf("{%v}", strings.Join(propsArr, ", "))
	}

	return fmt.Sprintf("(%v%v %v)", n.name, labels, props)
}

// Match ...
func (qb QueryBuilder) Match(matchClause string) QueryBuilder {
	return QueryBuilder{
		qb.query + `
		MATCH 
			` + matchClause,
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

// Limit ...
func (qb QueryBuilder) Limit(limitClause string) QueryBuilder {
	return QueryBuilder{
		qb.query + `
		LIMIT 
			` + limitClause,
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
	n := NewNode("w1").Labels("Person", "Wallet").Props(Prop{"name", "Theo Gauchoux"}, Prop{"age", 22})
	fmt.Println(n)

	var qb QueryBuilder
	fmt.Println(qb)

	// w1 := Node{[]string{}, map[string]interface{}{}}
	res, err := NewQueryBuilder().Match(
		`(w1:Wallet {address: {w1}})
			(w2:Wallet),
			p = (w1)-[tx:|*DEPTH*|]-(w2)`).With(
		`wp, w1, w2,
			w2.address AS recipient,
			|*WHERE*| AS tx2`,
	).Return("p, w1, w2, length(p), tx2").Limit("20").Execute()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("res: \n", res, "\n---\n ")

	res, err = NewQueryBuilder().Match("(a:Person)").Where("a.from = \"Sweden\"").Return("a").Execute()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("res: \n", res, "\n---\n ")
}
