# banking

`banking` provides helpers for financial identifiers — IBAN, ISIN, LEI, and UTI parsing, validation, and formatting, plus GLEIF API lookup.

## Types

| Type | Description |
|---|---|
| `IBAN` | IBAN value object with validation |
| `ISIN` | ISIN value object with Luhn checksum validation |
| `LEI` | Legal Entity Identifier value object |
| `UTI` | ISO 23897:2020 Unique Transaction Identifier value object |
| `GLIEF` | Struct modelling the GLEIF API v1 LEI-records response |

## Functions

### IBAN

| Function | Description |
|---|---|
| `NewIBAN(iban string) (IBAN, error)` | Constructs and validates an IBAN |
| `(*IBAN).String() string` | Returns the raw IBAN string |

### ISIN

| Function | Description |
|---|---|
| `(*ISIN).IsValid() bool` | Validates checksum and country prefix |
| `(*ISIN).String() string` | Returns the raw ISIN string |
| `(*ISIN).Get() string` | Returns the ISIN value |
| `(*ISIN).Set(string) error` | Sets and validates the ISIN |
| `(*ISIN).Printable() string` | Returns human-formatted ISIN (e.g. `GB 123456789 0`) |

### LEI

| Function | Description |
|---|---|
| `NewLEI(lei string) (LEI, error)` | Constructs and validates an LEI |
| `(*LEI).Formatted() string` | Returns a segmented LEI (LOU, reserved, entity, checksum) |
| `(*LEI).String() string` | Returns the raw LEI string |

### UTI

| Function | Description |
|---|---|
| `NewISO23897UTI(generatingEntity string) (UTI, error)` | Generates an ISO 23897 UTI from a 20-char entity ID |
| `(*UTI).String() string` | Returns the raw UTI string |
| `(*UTI).Get() string` | Returns the UTI value |
| `(*UTI).Set(string) error` | Sets and validates the UTI |
| `(*UTI).IsValid() (bool, error)` | Validates UTI length (42–52 chars) |
| `(*UTI).IsEmpty() bool` | Checks if the UTI is empty |
| `(*UTI).Formatted() string` | Returns human-formatted UTI |

### GLEIF Lookup

| Function | Description |
|---|---|
| `Lookup_LEI(inISIN string) (string, error)` | Looks up the LEI for an ISIN via the GLEIF API |

## Example

```go
import "github.com/mt1976/frantic-core/banking"

func main() {
    iban, err := banking.NewIBAN("GB82WEST12345698765432")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(iban.String())

    lei, err := banking.NewLEI("529900T8BM49AURSDO55")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(lei.Formatted())
}
```
