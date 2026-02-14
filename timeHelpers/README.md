# timeHelpers

`timeHelpers` provides timezone inference from locale region codes.

## Functions

| Function | Description |
|---|---|
| `InferTimezoneFromLocale(locale string) (string, error)` | Maps a locale string (e.g. `"en_GB"`) to an IANA timezone name |

## Example

```go
import "github.com/mt1976/frantic-core/timeHelpers"

func main() {
    tz, err := timeHelpers.InferTimezoneFromLocale("en_GB")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Timezone:", tz) // "Europe/London"
}
```
