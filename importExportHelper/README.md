# importExportHelper

`importExportHelper` contains helpers for importing and exporting data in CSV and JSON formats.

It’s designed to work with the project’s standard file locations (via `paths`) and logging conventions.

## Key functions

- `ExportCSV[T any](exportName string, exportList []T, idField entities.Field) error`
- `ExportJSON[T any](exportName string, exportList []T, idField entities.Field) error`
- `ImportCSV[T any](importName string, entryTypeToInsert T, importProcessor func(*T) (string, error)) error`

## Behaviour notes

- CSV delimiter defaults to `FIELDSEPARATOR` (currently `|`).
- `ExportCSV` writes to the defaults folder (`paths.Defaults()`), and appends a generated `# ...` metadata line at the end of the file.
- `ExportJSON` writes one JSON file per record into the dumps folder (`paths.Dumps()`).
- Naming uses a KSUID-based prefix (via `idHelpers.GetUUID()`), and attempts to include the record’s ID field.

## Example

```go
import (
    "github.com/mt1976/frantic-core/dao/entities"
    "github.com/mt1976/frantic-core/importExportHelper"
)

type User struct {
    ID   int
    Name string
}

func exportUsers(users []User) error {
    return importExportHelper.ExportCSV("users", users, entities.Field("ID"))
}

func importUsers() error {
    return importExportHelper.ImportCSV("users", User{}, func(u *User) (string, error) {
        // create/update in DB here
        return u.Name, nil
    })
}
```
