# financial

`financial` provides financial date/tenor calculations, settlement date logic, amount formatting, and rate ladder generation.

## Types

| Type | Description |
|---|---|
| `FinDate` | Represents a dated rate-ladder rung with Code, Name, Date, Index, Simple, External, and Human fields |
| `Tenor` | Financial tenor value object (e.g. `"1M"`, `"SP"`, `"3Y"`) with validation |

## Functions

### Tenor

| Function | Description |
|---|---|
| `NewTenor(term string) (Tenor, error)` | Creates and validates a Tenor |
| `(*Tenor).String() string` | Returns the tenor string |
| `(*Tenor).Set(term string) (*Tenor, error)` | Sets and validates the tenor |

### Date Calculations

| Function | Description |
|---|---|
| `GetDateFromTenor(tenor, tradeDate string, ccy ...string) (time.Time, error)` | Calculates the settlement date for a given tenor and currencies |
| `GetTenorFromDate(inDate, baseDate time.Time, ccy ...string) (Tenor, error)` | Reverse-maps a date to its closest tenor |
| `GetLadder(pivotDate time.Time, ccy ...string) ([]FinDate, int, error)` | Generates a full rate ladder of dated tenors |
| `GetSpotDate(time.Time) time.Time` | Returns T+2 adjusted for weekends |
| `GetTenorDate(time.Time, monthStr string) time.Time` | Returns T+N months adjusted for weekends |
| `GetFirstDayOfYear(time.Time) time.Time` | Returns Jan 2 adjusted for weekends |
| `SettlementDate(major, minor string, pivotDate time.Time) (time.Time, error)` | Calculates settlement date for a currency pair |
| `SettlementDateVia(major, minor string, pivotDate time.Time, via string) (time.Time, error)` | Calculates settlement date for a cross-currency pair |

### Amount Formatting

| Function | Description |
|---|---|
| `AbbrToInt(str string) int` | Converts abbreviated amounts (e.g. `"5M"`) to integers |
| `FormatAmount(float64, ccy string) string` | Formats amount using currency-specific symbol/DPS |
| `FormatAmountFullDPS(amount, ccy string) string` | Formats amount to 7 decimal places |
| `FormatAmountToDPS(amount, ccy, prec string) string` | Formats amount to specified decimal places |

## Example

```go
import "github.com/mt1976/frantic-core/financial"

func main() {
    tenor, _ := financial.NewTenor("3M")
    date, _ := financial.GetDateFromTenor("3M", "2024-01-15", "GBP", "USD")
    fmt.Println(tenor.String(), "settles on", date)

    formatted := financial.FormatAmount(1234567.89, "GBP")
    fmt.Println(formatted) // £1,234,567.89
}
```
