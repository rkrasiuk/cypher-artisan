# cypher-artisan

## Roadmap

### Query Clauses List
        
    + Match
    + Where
    + With
    + Return
    + OrderBy
    + Limit

    - As
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
    - OrderByDesc
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