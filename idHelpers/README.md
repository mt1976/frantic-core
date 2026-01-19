# idHelpers

`idHelpers` provides helpers for generating, formatting, and normalising IDs.

It includes both hashing helpers (for stable/derived IDs) and KSUID-based helpers (for globally unique, time-sortable IDs).

## Key functions

- `Encode(string) string`
  - Produces a SHA3-256 hex digest of a path-safe version of the input.
- `Decode(string) string`
  - Reverses path-safe encoding via `htmlHelpers.FromPathSafe`.
- `GetUUIDv2() string`
  - Returns a KSUID string.
- `GetUUIDv2WithPayload(payload string) (string, error)`
  - Returns a KSUID containing up to 16 bytes of payload.
- `GetUUIDv2Payload(uuid string) string`
  - Extracts a KSUID payload.
- `InspectUUIDv2(uuid string) string`
  - Human-readable KSUID summary.
- `BuildCompositeID(part ...interface{}) (compositeID string, encodedCompositeID string, err error)`
  - Builds a delimiter-joined composite ID plus a hashed version.
- `SanitizeID(id string) string`
  - Normalises identifiers (camel-case, removes special chars, strips separators).

## Example

```go
import "github.com/mt1976/frantic-core/idHelpers"

func example() {
    raw := "Customer 123"
    stable := idHelpers.Encode(raw)

    uuid := idHelpers.GetUUIDv2()

    composite, compositeHash, _ := idHelpers.BuildCompositeID("customer", 123)
    _, _, _ = composite, compositeHash, stable
    _ = uuid
}
```
