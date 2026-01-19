# dao/database generics helpers

This folder contains the Storm-backed database layer used across the project.

The file database_generics.go adds **typed** (generic) helper functions that reduce DAO boilerplate by avoiding `any`/`[]any` results and repeated type assertions.

## Why these are functions (not methods)

These helpers are package-level functions because Go does not allow type parameters on methods of a non-generic type (e.g. `func (db *DB) Get[T any](...)`).

So usage is:

- `database.GetTyped[T](db, field, value)`
- `database.GetAllTyped[T](db, ...)`
- `database.GetAllWhereTyped[T](db, field, value)`

## Requirements / constraints

- `db` must be a valid `*database.DB` (typically returned from `database.Connect(...)`).
- `T` must be a **non-pointer struct type** (e.g. `MyRecord`, not `*MyRecord`).
  - Storm expects a pointer destination internally, but the *type parameter* should be the struct.
- `field` is a `fields.Field` and must exist on `T` for the `*Where*` helpers.

## Functions

### `GetTyped[T any](db *DB, field fields.Field, value any) (T, error)`

Fetches a single record where `field == value`.

- Returns `error` when not found (Storm’s `ErrNotFound` or wrapped upstream).

Example:

```go
import (
    "github.com/mt1976/frantic-core/dao/database"
    "github.com/mt1976/frantic-core/dao/fields"
)

type User struct {
    ID   int    `storm:"id,increment"`
    Key  string `storm:"unique"`
    Name string
}

var UserFields = struct {
    ID  fields.Field
    Key fields.Field
}{
    ID:  "ID",
    Key: "Key",
}

func loadUser(db *database.DB, key string) (User, error) {
    return database.GetTyped[User](db, UserFields.Key, key)
}
```

### `GetAllTyped[T any](db *DB, options ...func(*index.Options)) ([]T, error)`

Fetches all records of type `T`.

- The optional `options` are passed through to Storm’s `All` (ordering, limits, etc.).

Example:

```go
import (
    "github.com/asdine/storm/v3/index"
    "github.com/mt1976/frantic-core/dao/database"
)

func loadAllUsers(db *database.DB) ([]User, error) {
    // No options
    return database.GetAllTyped[User](db)

    // With options (example only — depends on Storm’s index.Options)
    // return database.GetAllTyped[User](db, func(o *index.Options) {
    //     o.Skip = 0
    //     o.Limit = 100
    // })
}
```

### `GetAllWhereTyped[T any](db *DB, field fields.Field, value any) ([]T, error)`

Fetches all records of type `T` where `field == value`.

- Validates that the field exists on `T` and that `value` has a suitable type.
- If Storm returns `ErrNotFound`, this helper returns an **empty slice** (`[]T{}`) and `nil` error.

Example:

```go
import (
    "github.com/mt1976/frantic-core/dao/database"
)

type Order struct {
    ID     int    `storm:"id,increment"`
    Status string `storm:"index"`
}

var OrderFields = struct {
    Status fields.Field
}{
    Status: "Status",
}

func loadOpenOrders(db *database.DB) ([]Order, error) {
    return database.GetAllWhereTyped[Order](db, OrderFields.Status, "OPEN")
}
```

## Common pitfalls

- **Using `*T` instead of `T`:**
  - `database.GetTyped[*User](...)` is invalid by design and will return an invalid-type error.
- **Wrong field name:**
  - `GetAllWhereTyped` checks the field exists on the struct and errors early.
- **Wrong value type:**
  - If `field` is an `int` field but you pass a `string`, it will error.

## Where this is used

A concrete example DAO that uses these helpers is the TemplateStoreV2 implementation:

- `GetBy` uses `database.GetTyped[TemplateStore](...)`
- `GetAll` uses `database.GetAllTyped[TemplateStore](...)`
- `GetAllWhere` uses `database.GetAllWhereTyped[TemplateStore](...)`

See:

- dao/test/templateStoreV2/templateStoreV2.go
