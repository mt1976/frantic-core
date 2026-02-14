# dateHelpers

`dateHelpers` provides comprehensive date/time formatting, parsing, arithmetic, comparison, and business-day adjustment utilities.

## Types

| Type | Description |
|---|---|
| `DateFormat` | Struct holding all Go time layout strings (External, DMY, Internal, Detail, YMD, Calendar, BackupDate, BackupFolder, HTMLInput, DMY4, MYH) |

## Variables

| Variable | Description |
|---|---|
| `Format` | The active `DateFormat` instance, populated from configuration |

## Functions

### Formatting and Parsing

Paired format/parse functions for each date layout:

| Format | Parse |
|---|---|
| `FormatAudit(time.Time) string` | `ParseAudit(string) (time.Time, error)` |
| `FormatDMY(time.Time) string` | `ParseDMY(string) (time.Time, error)` |
| `FormatYMD(time.Time) string` | `ParseYMD(string) (time.Time, error)` |
| `FormatCalendar(time.Time) string` | `ParseCalendar(string) (time.Time, error)` |
| `FormatHuman(time.Time) string` | `ParseHuman(string) (time.Time, error)` |
| `FormatInternal(time.Time) string` | `ParseInternal(string) (time.Time, error)` |
| `FormatHTMLInput(time.Time) string` | `ParseHTMLInput(string) (time.Time, error)` |
| `FormatDMY4(time.Time) string` | `ParseDMY4(string) (time.Time, error)` |
| `FormatDetail(time.Time) string` | `ParseDetail(string) (time.Time, error)` |
| `FormatExternal(time.Time) string` | `ParseExternal(string) (time.Time, error)` |
| `FormatBackupDate(time.Time) string` | `ParseBackupDate(string) (time.Time, error)` |
| `FormatBackupFolder(time.Time) string` | `ParseBackupFolder(string) (time.Time, error)` |
| `FormatHumanFromString(string) string` | `ParseHumanFromString(string) (time.Time, error)` |

### Day Info

| Function | Description |
|---|---|
| `Today() time.Time` | Returns today's date |
| `Tomorrow() time.Time` | Returns tomorrow's date |
| `Yesterday() time.Time` | Returns yesterday's date |
| `StartOfDay(time.Time) time.Time` | Returns midnight of the given day |
| `EndOfDay(time.Time) time.Time` | Returns 23:59:59 of the given day |
| `IsWorkingDay(time.Time) bool` | Returns true if Monday–Friday |
| `NextWorkingDay(time.Time) time.Time` | Returns the next business day |
| `PreviousWorkingDay(time.Time) time.Time` | Returns the previous business day |
| `AdjustToNextWorkingday(time.Time) time.Time` | Adjusts to next working day if on a weekend |
| `AdjustToPreviousWorkingday(time.Time) time.Time` | Adjusts to previous working day if on a weekend |
| `DifferenceInDays(a, b time.Time) int` | Returns the number of days between two dates |
| `OrdinaliseDay(int) string` | Returns ordinal suffix (e.g. "1st", "2nd") |
| `HumanDate(time.Time) string` | Returns a friendly date string |

### Comparisons

| Function | Description |
|---|---|
| `IsSameDay(a, b time.Time) bool` | True if same calendar day |
| `IsTomorrow(time.Time) bool` | True if the date is tomorrow |
| `IsYesterday(time.Time) bool` | True if the date is yesterday |
| `IsToday(time.Time) bool` | True if the date is today |
| `IsBefore / IsAfter(a, b time.Time) bool` | Date ordering |
| `IsBeforeOrEqualTo / IsAfterOrEqualTo(a, b) bool` | Inclusive date ordering |
| `IsBetweenDates(d, start, end time.Time) bool` | True if within range |
| `IsWorkingDateBetween(d, start, end) bool` | True if within range and a working day |
| `IsWeekend / IsWeekday(time.Time) bool` | Day-of-week checks |
| `IsMonday … IsSunday(time.Time) bool` | Individual day checks |
| `IsAfterDays / IsBeforeDays(d, n) bool` | True if date is after/before N days from now |
| `IsInXDays / IsNotInXDays(d, n) bool` | True if exactly/not exactly N days from now |

### Arithmetic

| Function | Description |
|---|---|
| `AddDays / SubtractDays(t, n) time.Time` | Add/subtract calendar days |
| `AddWorkingDays / SubtractWorkingDays(t, n) time.Time` | Add/subtract business days |
| `AddWeeks / SubtractWeeks(t, n) time.Time` | Add/subtract weeks |
| `AddWorkingWeeks / SubtractWorkingWeeks(t, n) time.Time` | Add/subtract working weeks |
| `AddMonths / SubtractMonths(t, n) time.Time` | Add/subtract months |
| `AddWorkingMonths / SubtractWorkingMonths(t, n) time.Time` | Add/subtract working months |
| `AddYears / SubtractYears(t, n) time.Time` | Add/subtract years |
| `AddWorkingYears / SubtractWorkingYears(t, n) time.Time` | Add/subtract working years |
| `AddMonthAndAdjust(t, n) time.Time` | Add months with end-of-month adjustment |

## Example

```go
import "github.com/mt1976/frantic-core/dateHelpers"

func main() {
    today := dateHelpers.Today()
    tomorrow := dateHelpers.Tomorrow()
    fmt.Println("Today:", dateHelpers.FormatHuman(today))
    fmt.Println("Tomorrow:", dateHelpers.FormatHuman(tomorrow))
    fmt.Println("Is working day:", dateHelpers.IsWorkingDay(today))

    nextMonth := dateHelpers.AddWorkingMonths(today, 1)
    fmt.Println("Next working month:", dateHelpers.FormatHuman(nextMonth))
}
```
