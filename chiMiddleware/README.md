# chiMiddleware

`chiMiddleware` contains middleware helpers intended for use with `go-chi/chi`.

## Middleware

- `HandleHTMLMinification() func(http.Handler) http.Handler`
  - Minifies HTML responses when `Content-Type` contains `text/html`.
  - Uses `tdewolff/minify`.
- `HandleHTTPMethodConversion(next http.Handler) http.Handler`
  - Converts incoming `POST`/`GET` requests to `PUT` or `DELETE` if the form contains `_method=PUT` or `_method=DELETE`.
- `HandleBrotli(next http.Handler) http.Handler`
  - Adds Brotli compression support to responses.
- `InjectUserContext(next http.Handler, uid, username string) http.Handler`
  - Injects user ID and username into the request context.
- `InjectWorkerPoolIntoContext(next http.Handler, workerPool pond.Pool) http.Handler`
  - Injects a worker pool into the request context.

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
