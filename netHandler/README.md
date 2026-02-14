# netHandler

`netHandler` provides network host reachability checking via ping.

## Functions

| Function | Description |
|---|---|
| `CheckHostAvailability(addr string) (bool, error)` | Pings a host (platform-aware) and returns whether it is reachable |

## Example

```go
import "github.com/mt1976/frantic-core/netHandler"

func main() {
    available, err := netHandler.CheckHostAvailability("192.168.1.1")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Host available:", available)
}
```
