# htmlHelpers

`htmlHelpers` provides small helpers for HTML value conversion and URL-safe encoding/decoding.

## Functions

| Function | Description |
|---|---|
| `ValueToInt(s string) int` | Converts an HTML form value to an int |
| `ValueToBool(s string) bool` | Returns true if the value is `"on"` |
| `BoolToValue(b bool) string` | Returns `"on"` or `"off"` |
| `ToPathSafe(s string) (string, error)` | URL-escapes then base64-encodes a string for safe use in URL paths |
| `FromPathSafe(s string) (string, error)` | Reverses `ToPathSafe` |

## Example

```go
import "github.com/mt1976/frantic-core/htmlHelpers"

func main() {
    safe, _ := htmlHelpers.ToPathSafe("Hello World/123")
    original, _ := htmlHelpers.FromPathSafe(safe)
    fmt.Println(original) // "Hello World/123"

    fmt.Println(htmlHelpers.ValueToBool("on"))  // true
    fmt.Println(htmlHelpers.BoolToValue(false))  // "off"
}
