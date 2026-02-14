# tuiInputHelper

`tuiInputHelper` provides terminal UI input helpers — prompting for text and password input with optional mandatory enforcement.

## Functions

| Function | Description |
|---|---|
| `GetUserInput(question string, inputWidth int, inputMandatory bool, isPassword bool) string` | Prompts the user in the terminal for input; masks characters for password entry; repeats the prompt if the field is mandatory and left empty |

## Example

```go
import "github.com/mt1976/frantic-core/tuiInputHelper"

func main() {
    name := tuiInputHelper.GetUserInput("Enter your name", 40, true, false)
    fmt.Println("Hello,", name)

    password := tuiInputHelper.GetUserInput("Enter password", 40, true, true)
    _ = password
}
```
