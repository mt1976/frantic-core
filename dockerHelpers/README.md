# dockerHelpers

`dockerHelpers` provides utilities for detecting Docker environments and deploying startup payloads.

## Functions

| Function | Description |
|---|---|
| `IsDockerContainer() bool` | Returns true if `/.dockerenv` exists (i.e. running inside a Docker container) |
| `DeployDefaultsPayload() error` | Copies non-TOML files from `startupPayload/` to `./data/defaults/` |

## Example

```go
import "github.com/mt1976/frantic-core/dockerHelpers"

func main() {
    if dockerHelpers.IsDockerContainer() {
        fmt.Println("Running in Docker")
        if err := dockerHelpers.DeployDefaultsPayload(); err != nil {
            log.Fatal(err)
        }
    }
}
```
