# application

`application` provides OS and environment utilities for host/platform detection and system identity.

## Functions

| Function | Description |
|---|---|
| `OS() string` | Returns `"Windows"` or `"*nix"` based on the runtime platform |
| `IsRunningOnWindows() bool` | Returns true when running on Windows |
| `RunningInDockerContainer() bool` | *(Deprecated)* Delegates to `dockerHelpers.IsDockerContainer()` |
| `HostName() string` | Returns the lowercase hostname of the machine |
| `HostIP() string` | Returns the first non-loopback IPv4 address |
| `SystemIdentity() string` | Returns a cleaned, lowercase host identifier (alphanumeric only) |

## Constants

- `WINDOWS` — `"Windows"`
- `NIX` — `"*nix"`

## Example

```go
import "github.com/mt1976/frantic-core/application"

func main() {
    fmt.Println("OS:", application.OS())
    fmt.Println("Host:", application.HostName())
    fmt.Println("IP:", application.HostIP())
    fmt.Println("Identity:", application.SystemIdentity())
}
```
