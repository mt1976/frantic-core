# paths

`paths` provides centralized filesystem path helpers and constants for the application's directory structure.

## Types

| Type | Description |
|---|---|
| `FileSystemPath` | Wraps a path string with `String()` and `Is()` methods |

## Functions

| Function | Description |
|---|---|
| `Application() FileSystemPath` | Current working directory (application root) |
| `Data() FileSystemPath` | Path to the `data/` directory |
| `Config() FileSystemPath` | Path to the config directory |
| `Defaults() FileSystemPath` | Path to the defaults directory |
| `Database() FileSystemPath` | Path to the database directory |
| `Logs() FileSystemPath` | Path to the logs directory |
| `Backups() FileSystemPath` | Path to the backups directory |
| `Dumps() FileSystemPath` | Path to the dumps directory |
| `Res() FileSystemPath` | Path to the resources directory |
| `Images() FileSystemPath` | Path to image resources |
| `HTML() FileSystemPath` | Path to HTML templates |
| `HTMLTemplates() FileSystemPath` | Path to HTML templates (alias) |
| `HTMLPage(name string) string` | Full path to a named HTML page |
| `HTMLTemplate() string` | Path to the main templates file |
| `Seperator() string` | OS path separator |

## Example

```go
import "github.com/mt1976/frantic-core/paths"

func main() {
    fmt.Println("App root:", paths.Application().String())
    fmt.Println("Config:", paths.Config().String())
    fmt.Println("Database:", paths.Database().String())
    fmt.Println("Logs:", paths.Logs().String())
}
```
