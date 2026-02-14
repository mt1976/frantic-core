# notificationHandler

`notificationHandler` provides push notification delivery via the Pushover API.

## Functions

| Function | Description |
|---|---|
| `Send(inMessage, inTitle string, key int) error` | Sends a Pushover notification with the given title and message, with an optional callback URL |

Configuration (user key, API token) is read from `commonConfig`.

## Example

```go
import "github.com/mt1976/frantic-core/notificationHandler"

func main() {
    err := notificationHandler.Send("Deployment complete", "App Notification", 0)
    if err != nil {
        log.Fatal(err)
    }
}
```
