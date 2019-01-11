package builder

import (
	"strconv"
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
