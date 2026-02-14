# colours

`colours` exposes ANSI colour escape sequences for terminal text styling.

## Variables

| Variable | Description |
|---|---|
| `Reset` | Resets terminal styling |
| `Red` | Red text |
| `Green` | Green text |
| `Yellow` | Yellow text |
| `Blue` | Blue text |
| `Magenta` | Magenta text |
| `Cyan` | Cyan text |
| `Gray` | Gray text |
| `White` | White text |

## Example

```go
import "github.com/mt1976/frantic-core/colours"

func main() {
    fmt.Println(colours.Green + "Success!" + colours.Reset)
    fmt.Println(colours.Red + "Error!" + colours.Reset)
}
```
