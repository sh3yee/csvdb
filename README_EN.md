# csvdb

[English](README.md) | [中文](README_CN.md)

A Go library to operate CSV files like a database.

## Installation

```bash
go get github.com/sh3yee/csvdb
```

## Import Paths

This project is mainly used via subpackages:

```go
import (
    "github.com/sh3yee/csvdb/column"
    "github.com/sh3yee/csvdb/row"
    "github.com/sh3yee/csvdb/query"
)
```

Note: The root path `github.com/sh3yee/csvdb` is mainly for module indexing and documentation entry. Import subpackages for actual features.

If you use gopkg.in:

```bash
go get gopkg.in/sh3yee/csvdb.v1
```

## Quick Start

```go
package main

import (
    "github.com/sh3yee/csvdb/column"
    "github.com/sh3yee/csvdb/row"
)

func main() {
    // Column operations
    col := column.New("users.csv")
    col.Add("email")
    col.Alter("name", "username")
    col.DeleteByName("age")

    // Row operations
    r := row.New("users.csv")
    r.Add([]string{"1", "tom", "tom@example.com"})
    r.Update(0, []string{"1", "tommy", "tommy@example.com"})
    r.Delete(0)
}
```

## API

### Column

#### Add columns

| Method | Description |
|------|------|
| `Add(field string) error` | Add an empty column at the end |
| `AddAt(field string, index int) error` | Add an empty column at a specific index |
| `AddWithDefault(field, defaultValue string) error` | Add a column with default value at the end |
| `AddAtWithDefault(field, defaultValue string, index int) error` | Add a column with default value at a specific index |

#### Alter columns

| Method | Description |
|------|------|
| `Alter(oldName, newName string) error` | Rename a column by name |
| `AlterByIndex(index int, newName string) error` | Rename a column by index |

#### Delete columns

| Method | Description |
|------|------|
| `DeleteByName(name string) error` | Delete a column by name |
| `DeleteByIndex(index int) error` | Delete a column by index |

### Row

#### Add rows

| Method | Description |
|------|------|
| `Add(values []string) error` | Add a row to the end |
| `AddAt(values []string, index int) error` | Add a row at a specific index |

#### Update rows

| Method | Description |
|------|------|
| `Update(index int, values []string) error` | Update row at index |
| `UpdateBy(column, value string, newValues []string) error` | Update rows that match condition |

#### Delete rows

| Method | Description |
|------|------|
| `Delete(index int) error` | Delete row at index |
| `DeleteBy(column, value string) error` | Delete rows that match condition |

#### Query rows

| Method | Description |
|------|------|
| `Get(index int) ([]string, error)` | Get row at index |
| `GetBy(column, value string) ([][]string, error)` | Get rows that match condition |
| `GetAll() ([][]string, error)` | Get all rows |

### Query

#### Conditional queries

| Method | Description |
|------|------|
| `Find(cond Condition) ([][]string, error)` | Single-condition query |
| `FindAll(conds ...Condition) *Result` | Multi-condition query, returns chainable Result |
| `FindIn(column string, values []string) ([][]string, error)` | IN query |
| `FindNotIn(column string, values []string) ([][]string, error)` | NOT IN query |

#### Condition struct

```go
type Condition struct {
    Column string
    Op     string // "=", "!=", ">", "<", ">=", "<=", "like"
    Value  string
}
```

#### Chainable Result methods

| Method | Description |
|------|------|
| `Select(columns ...string) *Result` | Select specific columns |
| `OrderBy(column, order string) *Result` | Sort by one column, `order` is `asc` or `desc` |
| `ThenBy(column, order string) *Result` | Add additional sorting columns |
| `Limit(n int) *Result` | Limit result size |
| `Offset(n int) *Result` | Skip first N rows |
| `Get() ([][]string, error)` | Get all results |
| `First() ([]string, error)` | Get first row |
| `Count() (int, error)` | Count rows |
| `Exists() (bool, error)` | Check if records exist |

#### Query examples

```go
q := query.New("users.csv")

// Single-condition query
rows, _ := q.Find(query.Condition{Column: "age", Op: ">", Value: "18"})

// Multi-condition query + select + pagination
rows, _ := q.FindAll(
    query.Condition{Column: "age", Op: ">=", Value: "18"},
    query.Condition{Column: "city", Op: "=", Value: "Beijing"},
).Select("name", "email").Limit(10).Offset(5).Get()

// IN query
rows, _ := q.FindIn("id", []string{"1", "2", "3"})

// LIKE query
rows, _ := q.Find(query.Condition{Column: "name", Op: "like", Value: "tom%"})

// Sort
rows, _ := q.FindAll().OrderBy("age", "desc").Get()

// Multi-column sort
rows, _ := q.FindAll().OrderBy("age", "asc").ThenBy("name", "desc").Get()

// Query + sort + pagination
rows, _ := q.FindAll(
    query.Condition{Column: "status", Op: "=", Value: "active"},
).OrderBy("created_at", "desc").Limit(10).Get()

// Check existence
exists, _ := q.Find(query.Condition{Column: "id", Op: "=", Value: "1"}).Exists()
```

#### Aggregation

| Method | Description |
|------|------|
| `Count() (int, error)` | Count rows |
| `Sum(column string) (float64, error)` | Sum values in a column |
| `Avg(column string) (float64, error)` | Average values in a column |
| `Min(column string) (string, error)` | Minimum value (string comparison) |
| `Max(column string) (string, error)` | Maximum value (string comparison) |
| `MinFloat(column string) (float64, error)` | Minimum value (numeric comparison) |
| `MaxFloat(column string) (float64, error)` | Maximum value (numeric comparison) |

#### Aggregation examples

```go
q := query.New("users.csv")

// Count
count, _ := q.FindAll().Count()

// Sum
total, _ := q.FindAll().Sum("amount")

// Average
avg, _ := q.FindAll().Avg("age")

// Min/Max (string comparison)
minName, _ := q.FindAll().Min("name")
maxName, _ := q.FindAll().Max("name")

// Min/Max (numeric comparison)
minAge, _ := q.FindAll().MinFloat("age")
maxAge, _ := q.FindAll().MaxFloat("age")

// Aggregation after filtering
total, _ := q.FindAll(
    query.Condition{Column: "status", Op: "=", Value: "active"},
).Sum("amount")
```

### Errors

| Error | Description |
|------|------|
| `ErrColumnNotFound` | Column not found |
| `ErrRowNotFound` | Row not found |
| `ErrIndexOutOfRange` | Index out of range |

## Roadmap

| Module | Feature | Status |
|------|------|------|
| Row | Row operations (CRUD) | ✅ Done |
| Query | Conditional queries and filtering | ✅ Done |
| Sort | Sorting features | ✅ Done |
| Aggregate | Aggregation (COUNT/SUM/AVG/MIN/MAX) | ✅ Done |

## License

[MIT](LICENSE)
