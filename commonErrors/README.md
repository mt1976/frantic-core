# commonErrors

`commonErrors` declares shared sentinel errors and helper functions to wrap errors with consistent, contextual messages.

It’s used across the codebase so callers can:

- compare root causes with `errors.Is(err, commonErrors.ErrNotFound)` (etc.)
- wrap errors with useful context (`table`, `field`, `operation`, etc.)

## Examples of sentinel errors

- `ErrNotFound`
- `ErrValidationFailed`
- `ErrInvalidField`
- `ErrInvalidType`
- `ErrDAONotInitialised`

Cache-related:

- `ErrCacheNotEnabled`
- `ErrCacheRecordNotFound`
- `ErrCacheMultipleRecordsFound`

## Wrapper helpers

This package includes many wrapper functions such as:

- `ErrNotFoundWrapper(table string, err error) error`
- `ErrGetWrapper(table, field string, value any, readErr error) error`
- `ErrDAOInitialisationWrapper(table string, initErr error) error`
- `ErrCacheRecordNotFoundWrapper(table string, key any) error`

## Example

```go
import (
    "errors"

    ce "github.com/mt1976/frantic-core/commonErrors"
)

func example(err error) bool {
    return errors.Is(err, ce.ErrNotFound)
}
```
