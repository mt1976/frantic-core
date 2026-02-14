# mockData

`mockData` provides deterministic mock reference datasets for testing — countries, currencies, rate ladders, titles, genders, and biological sexes.

## Types

| Type | Description |
|---|---|
| `Country` | Country info: IBANLength, Currency, ISOCode, ISOCode3 |
| `Currency` | Currency info: Code, SpotDays, Name, Character, DPS, QuoteDPS, Type, MajorUnit, MinorUnit, and more |
| `BiologicalSex` | Biological sex entry: Name, Description |
| `Gender` | Gender entry: Name, Description |
| `Title` | Title string value |
| `Rung` | Rate ladder rung: Code, Name, Alternative, Index |

## Constants

| Constant | Description |
|---|---|
| `Fiat` | Fiat currency type |
| `Crypto` | Cryptocurrency type |
| `Metals` | Precious metals type |
| `Testing` | Testing/mock currency type |

## Variables

| Variable | Description |
|---|---|
| `Countries` | Map of country data keyed by ISO code |
| `Currencies` | Map of currency data keyed by currency code |
| `BiologicalSexes` | Map of biological sex entries |
| `Genders` | Map of gender entries |
| `Titles` | Map of title entries |
| `Ladder` | Map of rate ladder rungs |
| `LadderSize` | Number of rungs in the rate ladder |

## Functions

### Countries

| Function | Description |
|---|---|
| `GetCountryInfo(code string) (Country, error)` | Lookup by 2 or 3-letter ISO code |

### Currencies

| Function | Description |
|---|---|
| `GetCurrency(code string) (Currency, error)` | Lookup by currency code |
| `(*Currency).Age() int` | Returns years since the currency's introduction |

### Biology

| Function | Description |
|---|---|
| `GetBiologyList() []string` | List all biology keys |
| `GetBiologyInfo(string) BiologicalSex` | Lookup a biology entry |
| `IsValidBiology(string) bool` | Validate a biology key |

### Genders

| Function | Description |
|---|---|
| `GetGenderList() []string` | List all gender keys |
| `IsValidGender(string) bool` | Validate a gender key |

### Titles

| Function | Description |
|---|---|
| `GetList() []string` | List all title keys |
| `IsValidTitle(string) bool` | Validate a title key |

### Rate Ladder

| Function | Description |
|---|---|
| `GetRateLadderList() []string` | List all tenor codes |
| `IsValidPeriod(string) bool` | Validate a tenor code |
| `GetRateLadderByIndex(int) Rung` | Get a ladder rung by index |
| `LadderToString(map[string]Rung) string` | Serialize the ladder to a string |
| `GetTenorInfo(string) (Rung, error)` | Lookup tenor rung info |

## Example

```go
import "github.com/mt1976/frantic-core/mockData"

func main() {
    gbp, _ := mockData.GetCurrency("GBP")
    fmt.Println(gbp.Name, gbp.Character) // "Pound Sterling" "£"

    uk, _ := mockData.GetCountryInfo("GB")
    fmt.Println(uk.Currency) // "GBP"
}
```
