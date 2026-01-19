# contextHandler

`contextHandler` manages request-scoped context values (typically “session” related) and provides helpers to get and set those values.

Keys are derived from `commonConfig` session key names, and an internal identifier is created via `idHelpers.Encode`.

## Getters

- `GetSession_ID(ctx context.Context) string`
- `GetSession_UserKey(ctx context.Context) string`
- `GetSession_UserCode(ctx context.Context) string`
- `GetSession_Token(ctx context.Context) any`
- `GetSession_Expiry(ctx context.Context) time.Time`
- `GetSession_Locale(ctx context.Context) string`
- `GetSession_Theme(ctx context.Context) string`
- `GetSession_Timezone(ctx context.Context) string`
- `GetSession_UserRole(ctx context.Context) string`

## Setters

- `SetSession_ID(ctx context.Context, sessionID string) context.Context`
- `SetSession_UserKey(ctx context.Context, userKey string) context.Context`
- `SetSession_UserCode(ctx context.Context, userCode string) context.Context`
- `SetSession_Token(ctx context.Context, token any) context.Context`
- `SetSession_Expiry(ctx context.Context, expiry time.Time) context.Context`
- `SetSession_Locale(ctx context.Context, locale string) context.Context`
- `SetSession_Theme(ctx context.Context, theme string) context.Context`
- `SetSession_Timezone(ctx context.Context, timezone string) context.Context`
- `SetSession_UserRole(ctx context.Context, role string) context.Context`

## Debugging

- `Debug(ctx context.Context, name string)`

## Example

```go
import (
    "context"
    "time"

    "github.com/mt1976/frantic-core/contextHandler"
)

func example() {
    ctx := context.Background()
    ctx = contextHandler.SetSession_ID(ctx, "abc")
    ctx = contextHandler.SetSession_Expiry(ctx, time.Now().Add(30*time.Minute))

    _ = contextHandler.GetSession_ID(ctx)
}
```
