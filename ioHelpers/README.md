# ioHelpers

`ioHelpers` provides file I/O utilities — read, write, copy, backup, dump, and directory management.

## Functions

### File Operations

| Function | Description |
|---|---|
| `Read(fileName, path string) (string, error)` | Reads file contents as a string |
| `Write(fileName, path, content string) (bool, error)` | Writes content to a file |
| `WriteData(fileName, path, content string) int` | Writes data to a file (returns 1 on success) |
| `Copy(fileName, fromPath, toPath string) bool` | Copies a file between paths |
| `CopyFile(src, dst string) error` | Copies a file preserving permissions |
| `Touch(filename string) bool` | Returns true if the file exists |

### Directory Operations

| Function | Description |
|---|---|
| `MkDir(path string) error` | Creates a single directory |
| `MkdirAll(path string) error` | Recursively creates directories |
| `Dir(path string) ([]string, error)` | Lists subdirectory names |
| `GetFolders(path string) ([]string, error)` | Lists subdirectory names |
| `Empty(dir string) error` | Deletes all files in a directory |
| `DeleteFolder(path string) error` | Removes a directory tree |

### Database & Debugging

| Function | Description |
|---|---|
| `GetDBFileName(name string) string` | Constructs a database file path |
| `Dump(tableName, where, action, recordID string, data any)` | JSON-dumps a record to a timestamped file |
| `Backup(table, location string)` | Copies a DB file to a backup location |

## Example

```go
import "github.com/mt1976/frantic-core/ioHelpers"

func main() {
    content, err := ioHelpers.Read("config.json", "./data")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(content)

    _, err = ioHelpers.Write("output.txt", "./data", "hello world")
    if err != nil {
        log.Fatal(err)
    }
}
```
