# mathHelpers

`mathHelpers` provides small numeric helper functions — random numbers, min/max, and coin toss.

## Functions

| Function | Description |
|---|---|
| `RandomInt(max int) int` | Returns a random int in `[0, max)` |
| `RandomBetween(min, max int) int` | Returns a random int in `[min, max)` |
| `CoinToss() bool` | Returns a random boolean |
| `Max(x, y int) int` | Returns the larger of two ints |
| `Min(x, y int) int` | Returns the smaller of two ints |

## Example

```go
import "github.com/mt1976/frantic-core/mathHelpers"

func main() {
    n := mathHelpers.RandomBetween(1, 100)
    fmt.Println("Random number:", n)
    fmt.Println("Coin toss:", mathHelpers.CoinToss())
    fmt.Println("Max:", mathHelpers.Max(10, 20))
}
```
