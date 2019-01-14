# cypher-artisan

## Query Sample

```
MATCH (m:Movie)<-[:RATED]-(u:User)
WHERE m.title CONTAINS "Matrix"
WITH m.title AS movie, COUNT(*) AS reviews
RETURN movie, reviews
ORDER BY reviews DESC
LIMIT 5;
```

```go
    QueryBuilder().
		Match(
			Edge("").Labels("RATED").Path(IncomingPath).Relationship(
				Node("m").Labels("Movie").String(),
				Node("u").Labels("User").String(),
			),
		).
		Where(`m.title CONTAINS "Matrix"`).
		With(As("m.title", "movie"), As("COUNT(*)", "reviews")).
		Return("movie", "reviews").
		OrderByDesc("reviews").
		Limit(5).
        Execute()
```

## Roadmap

### Query Clauses List
        
    + Match
    + Where
    + With
    + Return
    + OrderBy
    + Limit
    + OrderByDesc
    + As
    + Assign (`=`)

    - OptionalMatch
    - Skip
    - Create
    - CreateUnique
    - Merge
    - Set
    - Delete
    - Remove
    - ForEach
    - ReturnDistinct
    - Union
    - On
    - DetachDelete
    - Call
    - Yield
    - Unwind
    - Case (When/Then/Else) End
    - CreateIndexOn
    - UsingIndexOn
    - DropIndexOn
    - CreateConstraintOn
    - DropConstraintOn
    - Assert
    - AssertIsUnique

### Operators
### Predicates/Conditions
### Performance
    - Profile
    - Explain


## Cypher RefCard

### Read Query Structure
```
[MATCH WHERE]
[OPTIONAL MATCH WHERE]
[WITH [ORDER BY] [SKIP] [LIMIT]]
RETURN [ORDER BY] [SKIP] [LIMIT]
```

### Write-Only Query Structure
```
(CREATE [UNIQUE] | MERGE)*
[SET|DELETE|REMOVE|FOREACH]*
[RETURN [ORDER BY] [SKIP] [LIMIT]]
```

### Read-Write Query Structure
```
[MATCH WHERE]
[OPTIONAL MATCH WHERE]
[WITH [ORDER BY] [SKIP] [LIMIT]]
(CREATE [UNIQUE] | MERGE)*
[SET|DELETE|REMOVE|FOREACH]*
[RETURN [ORDER BY] [SKIP] [LIMIT]]
```