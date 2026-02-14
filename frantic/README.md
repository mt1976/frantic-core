# frantic

`frantic` provides identity utilities and validation for official project/component names in the frantic ecosystem.

## Types

| Type | Description |
|---|---|
| `Identity` | Represents a two-part project identity (e.g. `"frantic-core"`) |

## Functions

| Function | Description |
|---|---|
| `New(name string) (Identity, error)` | Parses and validates a `"prefix-suffix"` identity string |
| `(*Identity).String() string` | Returns the full identity name |
| `(*Identity).Name() string` | Returns the full identity name |
| `(*Identity).Prefix() string` | Returns the prefix part |
| `(*Identity).Suffix() string` | Returns the suffix part |
| `(*Identity).IsOfficial() error` | Validates against known official origins |
| `ValidateIdentityOrigin(name string) error` | Checks if a name is a recognized project origin |

## Example

```go
import "github.com/mt1976/frantic-core/frantic"

func main() {
    id, err := frantic.New("frantic-core")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(id.Prefix()) // "frantic"
    fmt.Println(id.Suffix()) // "core"
    fmt.Println(id.IsOfficial()) // nil (valid)
}
```
