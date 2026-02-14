# emailHandler

`emailHandler` provides email composition and sending via SMTP using `gomail`.

## Variables

| Variable | Description |
|---|---|
| `EMAIL_From` | Default sender address |
| `EMAIL_Footer` | Default email footer text |
| `Emailer` | `*gomail.Dialer` instance for sending |

## Constants

| Constant | Description |
|---|---|
| `DATEMSG` | Date format string for email timestamps |
| `ES_FROM`, `ES_TO`, `ES_CC`, `ES_SUBJECT`, `ES_TYPE` | Email field identifiers |

## Functions

| Function | Description |
|---|---|
| `Email_init() *gomail.Dialer` | Initializes the email dialer from `commonConfig` settings |
| `SendEmail(to, name, subject, body string)` | Sends an HTML email to the given recipient |

## Example

```go
import "github.com/mt1976/frantic-core/emailHandler"

func main() {
    emailHandler.Email_init()
    emailHandler.SendEmail("user@example.com", "User", "Subject", "<p>Hello!</p>")
}
```
