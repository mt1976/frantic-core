# chiMiddleware

`chiMiddleware` contains middleware helpers intended for use with `go-chi/chi`.

## Middleware

- `HandleHTMLMinification() func(http.Handler) http.Handler`
  - Minifies HTML responses when `Content-Type` contains `text/html`.
  - Uses `tdewolff/minify`.
- `HandleHTTPMethodConversion(next http.Handler) http.Handler`
  - Converts incoming `POST`/`GET` requests to `PUT` or `DELETE` if the form contains `_method=PUT` or `_method=DELETE`.
- `HandleBrotli(next http.Handler) http.Handler`
  - Brotli support. (Currently incomplete: it sets up an encoder but does not call `next`.)

## Example (chi)

```go
import (
    "net/http"

    "github.com/go-chi/chi/v5"
    "github.com/mt1976/frantic-core/chiMiddleware"
)

func router() http.Handler {
    r := chi.NewRouter()

    r.Use(chiMiddleware.HandleHTMLMinification())
    r.Use(chiMiddleware.HandleHTTPMethodConversion)

    return r
}
```
