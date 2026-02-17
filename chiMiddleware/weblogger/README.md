# webLogger

> HTTP request logging middleware for Go, using the standard `log` package

`webLogger` is a lightweight HTTP request logging middleware for Go web applications. It accepts a standard `*log.Logger` (e.g. `logHandler.Web`) and logs each request with method, path, status code, duration, and optional extras.

## Features

- **Standard Logging**: Uses Go's standard `log` package — pass any `*log.Logger`
- **Panic Recovery**: Recovers panics with stack traces and HTTP 500 responses
- **Body Logging**: Conditional request/response body capture with content-type filtering
- **Custom Attributes**: Add string attributes from handlers and middlewares via `SetAttrs`
- **Request Filtering**: Skip logging for specific requests via `Skip` function
- **`curl` Generation**: Generate `curl` commands for debugging

## Usage

```go
package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mt1976/frantic-core/chiMiddleware/weblogger"
	"github.com/mt1976/frantic-core/logHandler"
)

func main() {
	r := chi.NewRouter()

	// Request logger using logHandler.Web
	r.Use(webLogger.RequestLogger(logHandler.Web, &webLogger.Options{
		RecoverPanics: true,
		Skip: func(req *http.Request, respStatus int) bool {
			return respStatus == 404 || respStatus == 405
		},
		LogRequestHeaders:  []string{"Origin"},
		LogResponseHeaders: []string{},
	}))

	// Set request log attributes from within middleware.
	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			webLogger.SetAttrs(r.Context(), "user=user1")
			next.ServeHTTP(w, r)
		})
	})

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world \n"))
	})

	http.ListenAndServe("localhost:8000", r)
}
```

## License
[MIT license](./LICENSE)
