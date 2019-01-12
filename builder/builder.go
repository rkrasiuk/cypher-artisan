package builder

import (
	"fmt"
	"strconv"
	"strings"
)

// QueryBuilder ...
type QueryBuilder struct {
	query string
}

// NewQueryBuilder ...
func NewQueryBuilder() QueryBuilder {
	return QueryBuilder{}
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
func (qb QueryBuilder) With(withClauses ...string) QueryBuilder {
	return QueryBuilder{
		qb.query + `
		WITH 
			` + strings.Join(withClauses, ", "),
	}
}

// Return ...
func (qb QueryBuilder) Return(returnClauses ...string) QueryBuilder {
	return QueryBuilder{
		qb.query + `
		RETURN 
			` + strings.Join(returnClauses, ", "),
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

// OrderByDesc ...
func (qb QueryBuilder) OrderByDesc(orderByDescClause string) QueryBuilder {
	return QueryBuilder{
		qb.query + `
		ORDER BY 
		` + orderByDescClause + ` DESC`,
	}
}

// Limit ...
func (qb QueryBuilder) Limit(limit int) QueryBuilder {
	return QueryBuilder{
		qb.query + `
		LIMIT	` + strconv.Itoa(limit),
	}
}

// Execute ...
func (qb QueryBuilder) Execute() string {
	return qb.query
}

// As ...
func As(initial, alias string) string {
	return fmt.Sprintf("%v AS %v", initial, alias)
}

// Assign ...
func Assign(name, pattern string) string {
	return fmt.Sprintf("%v = %v", name, pattern)
}
