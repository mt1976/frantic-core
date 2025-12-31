# Template Store DAO

Package `templateStore` provides Data Access Object (DAO) functionality for managing `TemplateStore` entities.

## Domain Information

- **Domain**: `Template` (from `Domain`)
- **Table Name**: `TemplateStore` (from `TableName = Domain + "Store"`)

## Package Variables

- `Domain = "Template"`
- `TableName = Domain + "Store"`
- `Fields` provides a structured way to reference model field names.

## Struct Definition

### TemplateStore

`TemplateStore` represents a User entity.

| Field           | Type            | Tags                                      | Description                     |
| --------------- | --------------- | ----------------------------------------- | ------------------------------- |
| `ID`            | `int`           | `storm:"id,increment=100"`                | Primary key with auto increment |
| `Key`           | `string`        | `storm:"unique"`                          | Key                             |
| `Raw`           | `string`        | `storm:"unique"`                          | Raw ID before encoding          |
| `UID`           | `string`        | `validate:"required"`                     |                                 |
| `GID`           | `string`        | `storm:"index" validate:"required"`       |                                 |
| `RealName`      | `string`        | `validate:"required,min=5"`               | This field will not be indexed  |
| `UserName`      | `string`        | `validate:"required,min=5"`               |                                 |
| `UserCode`      | `string`        | `storm:"index" validate:"required,min=5"` |                                 |
| `Email`         | `string`        |                                           |                                 |
| `Notes`         | `string`        | `validate:"max=75"`                       |                                 |
| `Active`        | `dao.StormBool` |                                           |                                 |
| `ExampleInt`    | `dao.Int`       |                                           |                                 |
| `ExampleFloat`  | `dao.Float`     |                                           |                                 |
| `ExampleBool`   | `dao.Bool`      |                                           |                                 |
| `ExampleDate`   | `time.Time`     |                                           |                                 |
| `ExampleString` | `string`        |                                           |                                 |
| `LastLogin`     | `time.Time`     |                                           |                                 |
| `LastHost`      | `string`        |                                           |                                 |
| `Audit`         | `audit.Audit`   | `csv:"-"`                                 | Audit data                      |

### Field Constants

```go
Fields.ID            = "ID"
Fields.Key           = "Key"
Fields.Raw           = "Raw"
Fields.Audit         = "Audit"
Fields.UID           = "UID"
Fields.GID           = "GID"
Fields.RealName      = "RealName"
Fields.UserName      = "UserName"
Fields.UserCode      = "UserCode"
Fields.Email         = "Email"
Fields.Notes         = "Notes"
Fields.Active        = "Active"
Fields.ExampleInt    = "ExampleInt"
Fields.ExampleFloat  = "ExampleFloat"
Fields.ExampleBool   = "ExampleBool"
Fields.ExampleDate   = "ExampleDate"
Fields.ExampleString = "ExampleString"
Fields.LastLogin     = "LastLogin"
Fields.LastHost      = "LastHost"
```

## Operations

### Initialisation / Lifecycle

| Function          | Signature                                           | Description |
| ---------------- | --------------------------------------------------- | ----------- |
| `Initialise`     | `Initialise(ctx context.Context)`                   | Sets up the database connection and prepares the DAO for operations. |
| `IsInitialised`  | `IsInitialised() bool`                              | Returns the initialisation status of the DAO. |
| `Close`          | `Close()`                                           | Terminates the connection to the database used by the DAO. |
| `Drop`           | `Drop() error`                                      | Removes the DAO's database entirely. |
| `ClearDown`      | `ClearDown(ctx context.Context) error`              | Removes all records from the DAO's database. |
| `GetDatabaseConnections` | `GetDatabaseConnections() func() ([]*database.DB, error)` | Returns a function that fetches the current database instances. |

### Create

| Function | Signature                                                                                        | Description |
| -------- | ------------------------------------------------------------------------------------------------ | ----------- |
| `New`    | `New() TemplateStore`                                                                            | Creates a new template instance. |
| `Create` | `Create(ctx context.Context, userName, uid, realName, email, gid string) (TemplateStore, error)` | Creates a new template instance in the database. |
| `Add`    | `Add(ctx context.Context) (TemplateStore, error)`                                                | (No doc comment) |
| `(*TemplateStore) Create` | `(record *TemplateStore) Create(ctx context.Context, note string) error`            | Inserts a new TemplateStore record into the database. |

### Read

| Function      | Signature                                                          | Description |
| ------------- | ------------------------------------------------------------------ | ----------- |
| `GetBy`       | `GetBy(field dao.Field, value any) (TemplateStore, error)`         | Retrieves a record by specified field and value. |
| `GetById`     | `GetById(id any) (TemplateStore, error)`                           | Retrieves a record by ID. |
| `GetByKey`    | `GetByKey(key any) (TemplateStore, error)`                         | Retrieves a record by key. |
| `GetAll`      | `GetAll() ([]TemplateStore, error)`                                | Retrieves all records. |
| `GetAllWhere` | `GetAllWhere(field dao.Field, value any) ([]TemplateStore, error)` | Retrieves all records matching criteria. |
| `Count`       | `Count() (int, error)`                                             | Returns the total number of records. |
| `CountWhere`  | `CountWhere(field dao.Field, value any) (int, error)`              | Counts records matching criteria. |

### Update

| Function           | Signature                                                                                                      | Description |
| ------------------ | -------------------------------------------------------------------------------------------------------------- | ----------- |
| `(*TemplateStore) Update`           | `(record *TemplateStore) Update(ctx context.Context, note string) error`                         | Updates the record in the database. |
| `(*TemplateStore) UpdateWithAction` | `(record *TemplateStore) UpdateWithAction(ctx context.Context, auditAction audit.Action, note string) error` | Updates with a specified audit action. |
| `(*TemplateStore) SetName`          | `(u *TemplateStore) SetName(name string) error`                                                  | (No doc comment) |

### Delete

| Function      | Signature                                                                      | Description |
| ------------- | ------------------------------------------------------------------------------ | ----------- |
| `Delete`      | `Delete(ctx context.Context, id int, note string) error`                       | Deletes a record by ID. |
| `DeleteBy`    | `DeleteBy(ctx context.Context, field dao.Field, value any, note string) error` | Deletes a record by specified field and value. |
| `DeleteByKey` | `DeleteByKey(ctx context.Context, key string, note string) error`              | Deletes a record by key. |

### Validation / Utility

| Function | Signature                                           | Description |
| -------- | --------------------------------------------------- | ----------- |
| `(*TemplateStore) Validate` | `(record *TemplateStore) Validate() error`                    | Checks if the record is valid. |
| `(*TemplateStore) Clone`    | `(record *TemplateStore) Clone(ctx context.Context) (TemplateStore, error)` | Clones the current record in the database. |
| `(*TemplateStore) Spew`     | `(record *TemplateStore) Spew()`                               | Outputs the record contents to the Info log. |
| `Login`                     | `Login(ctx context.Context)`                                   | (No doc comment) |

### Lookup

| Function           | Signature                                          | Description |
| ------------------ | -------------------------------------------------- | ----------- |
| `GetDefaultLookup` | `GetDefaultLookup() (lookup.Lookup, error)`        | Builds a default lookup using `Key` and `Raw`. |
| `GetLookup`        | `GetLookup(field, value dao.Field) (lookup.Lookup, error)` | Builds a lookup for specified fields. |

### Import / Export

| Function | Signature | Description |
| -------- | --------- | ----------- |
| `ExportAllAsCSV`  | `ExportAllAsCSV() error`              | Exports all records as a CSV file. |
| `ExportAllAsJSON` | `ExportAllAsJSON(message string)`     | Exports all records as JSON files. |
| `ImportAllFromCSV` | `ImportAllFromCSV() error`           | (No doc comment) |
| `(*TemplateStore) ExportRecordAsCSV`  | `(record *TemplateStore) ExportRecordAsCSV(name string) error` | Exports a single record as CSV. |
| `(*TemplateStore) ExportRecordAsJSON` | `(record *TemplateStore) ExportRecordAsJSON(name string)`      | Exports a single record as JSON. |

### Worker / Cache

| Function  | Signature                                   | Description |
| --------- | ------------------------------------------- | ----------- |
| `PreLoad` | `PreLoad(ctx context.Context) error`        | Preloads the cache from the database. |
| `Worker`  | `Worker(j jobs.Job, db *database.DB)`       | Job scheduled to run at a predefined interval. |

## Usage Examples

> Note: call `Initialise(ctx)` once during application startup (before CRUD operations), and `Close()` during shutdown.

### Initialise and create a record

```go
package main

import (
    "context"
    "log"

    "github.com/mt1976/frantic-core/dao/test/templateStore"
)

func main() {
    ctx := context.Background()

    templateStore.Initialise(ctx)
    defer templateStore.Close()

    rec, err := templateStore.Create(ctx, "jdoe", "1001", "John Doe", "jdoe@example.com", "group-1")
    if err != nil {
        log.Fatal(err)
    }

    log.Printf("created ID=%d Key=%s", rec.ID, rec.Key)
}
```

### Read, update, validate, and delete

```go
package main

import (
    "context"
    "log"

    "github.com/mt1976/frantic-core/dao/test/templateStore"
)

func main() {
    ctx := context.Background()

    templateStore.Initialise(ctx)
    defer templateStore.Close()

    // Fetch by a field (recommended)
    rec, err := templateStore.GetBy(templateStore.Fields.UserCode, "1001:jdoe")
    if err != nil {
        log.Fatal(err)
    }

    // Update
    rec.Email = "john.doe@newdomain.test"
    if err := rec.Validate(); err != nil {
        log.Fatal(err)
    }
    if err := rec.Update(ctx, "updated email"); err != nil {
        log.Fatal(err)
    }

    // Delete by ID
    if err := templateStore.Delete(ctx, rec.ID, "cleanup"); err != nil {
        log.Fatal(err)
    }
}
```

### Lookup and export

```go
package main

import (
    "context"
    "log"

    "github.com/mt1976/frantic-core/dao/test/templateStore"
)

func main() {
    ctx := context.Background()

    templateStore.Initialise(ctx)
    defer templateStore.Close()

    // Build a default lookup (Key -> Raw)
    lu, err := templateStore.GetDefaultLookup()
    if err != nil {
        log.Fatal(err)
    }
    _ = lu // use lookup

    // Export all records
    if err := templateStore.ExportAllAsCSV(); err != nil {
        log.Fatal(err)
    }
    templateStore.ExportAllAsJSON("nightly backup")
}
```

### Cache preload

```go
package main

import (
    "context"
    "log"

    "github.com/mt1976/frantic-core/dao/test/templateStore"
)

func main() {
    ctx := context.Background()

    templateStore.Initialise(ctx)
    defer templateStore.Close()

    if err := templateStore.PreLoad(ctx); err != nil {
        log.Fatal(err)
    }
}
```

## Package Files

- `templateStore.go`
- `templateStoreCache.go`
- `templateStoreDB.go`
- `templateStoreHelpers.go`
- `templateStoreImpex.go`
- `templateStoreInternals.go`
- `templateStoreModel.go`
- `templateStoreNew.go`
- `templateStoreWorker.go`
