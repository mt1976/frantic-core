package chiMiddleware

import (
	"bytes"
	"io"
	"net/http"
	"strings"

	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/html"
)

// responseWriterWrapper captures the response body
type responseWriterWrapper struct {
	http.ResponseWriter
	status int
	buf    *bytes.Buffer
}

func (rw *responseWriterWrapper) WriteHeader(code int) {
	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
}

func (rw *responseWriterWrapper) Write(b []byte) (int, error) {
	return rw.buf.Write(b)
}

// HandleHTMLMinification returns a chi middleware that minifies HTML responses
func HandleHTMLMinification() func(http.Handler) http.Handler {
	m := minify.New()
	m.Add("text/html", &html.Minifier{
		KeepSpecialComments: true,
	})

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Capture the response
			buf := &bytes.Buffer{}
			rw := &responseWriterWrapper{
				ResponseWriter: w,
				buf:            buf,
				status:         http.StatusOK,
			}

			next.ServeHTTP(rw, r)

			contentType := w.Header().Get("Content-Type")
			if strings.Contains(contentType, "text/html") {
				// Minify HTML content
				minified, err := m.String("text/html", buf.String())
				if err != nil {
					// fallback to unminified content
					io.WriteString(w, buf.String())
					return
				}
				w.Header().Set("Content-Length", "")
				w.WriteHeader(rw.status)
				io.WriteString(w, minified)
				// godump.Dump(minified)
			} else {
				// non-HTML response
				w.WriteHeader(rw.status)
				io.Copy(w, buf)
				// godump.Dump(buf.String())
			}
		})
	}
}
