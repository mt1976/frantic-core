# stringHelpers

`stringHelpers` provides common string utilities — formatting, quoting, encoding, special-character handling, camel-casing, and comparison.

## Functions

### Casing and Formatting

| Function | Description |
|---|---|
| `LowerFirst(s string) string` | Lowercases the first character |
| `CamelCase(string) string` | Converts to camelCase |
| `PadRight(s, p string, l int) string` | Pads string to length on the right |
| `PadLeft(s, p string, l int) string` | Pads string to length on the left |

### Array Conversion

| Function | Description |
|---|---|
| `ArrToString([]string) string` | Joins array elements with newlines |
| `StrArrayToString([]string) string` | Joins array elements with newlines |
| `StrArrayToStringWithSep([]string, sep string) string` | Joins array with a custom separator |

### Special Characters

| Function | Description |
|---|---|
| `RemoveSpecialChars(string) string` | Replaces non-alphanumeric chars with dashes |
| `RemoveSpecialCharacters(string) string` | Strips all non-letter/digit Unicode chars |
| `ReplaceWildcard(orig, wildcard, replacement string) string` | Replaces `{{wildcard}}` patterns in a string |
| `MakeStringStorable(string) string` | Replaces CR/LF with `{cr}`/`{lf}` for storage |
| `MakeStringDisplayable(string) string` | Reverses `MakeStringStorable` |

### Quoting and Wrapping

| Function | Description |
|---|---|
| `DQuote(s string) string` | Wraps in double quotes |
| `SQuote(s string) string` | Wraps in single quotes |
| `DBracket(s string) string` | Wraps in `[[ ]]` |
| `SBracket(s string) string` | Wraps in `[ ]` |
| `DChevrons(s string) string` | Wraps in `<< >>` |
| `SChevrons(s string) string` | Wraps in `< >` |
| `DCurlies(s string) string` | Wraps in `{{ }}` |
| `SCurlies(s string) string` | Wraps in `{ }` |
| `DParentheses(s string) string` | Wraps in `(( ))` |
| `SParentheses(s string) string` | Wraps in `( )` |

### Encoding and Comparison

| Function | Description |
|---|---|
| `Encode(string) string` | SHA3-256 hash of path-safe encoded input |
| `CompareStrings(a, b string)` | Logs a line-by-line diff between two strings |

## Example

```go
import "github.com/mt1976/frantic-core/stringHelpers"

func main() {
    fmt.Println(stringHelpers.CamelCase("hello world"))    // "helloWorld"
    fmt.Println(stringHelpers.DQuote("hello"))             // `"hello"`
    fmt.Println(stringHelpers.PadRight("hi", ".", 10))     // "hi........"
    fmt.Println(stringHelpers.RemoveSpecialChars("a@b#c")) // "a-b-c"
}
```
