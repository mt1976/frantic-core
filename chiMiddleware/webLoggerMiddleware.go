package chiMiddleware

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"runtime"
	"strings"
	"time"

	"github.com/go-chi/chi/v5/middleware"
)

var ErrClientAborted = fmt.Errorf("request aborted: client disconnected before response was sent")

func RequestLogger(logger *log.Logger, o *Options) func(http.Handler) http.Handler {
	if o == nil {
		o = &defaultOptions
	}
	if len(o.LogBodyContentTypes) == 0 {
		o.LogBodyContentTypes = defaultOptions.LogBodyContentTypes
	}
	if o.LogBodyMaxLen == 0 {
		o.LogBodyMaxLen = defaultOptions.LogBodyMaxLen
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), ctxKeyLogAttrs{}, &[]string{})

			logReqBody := o.LogRequestBody != nil && o.LogRequestBody(r)
			logRespBody := o.LogResponseBody != nil && o.LogResponseBody(r)

			var reqBody bytes.Buffer
			if logReqBody || o.LogExtraAttrs != nil {
				r.Body = io.NopCloser(io.TeeReader(r.Body, &reqBody))
			}

			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

			var respBody bytes.Buffer
			if o.LogResponseBody != nil && o.LogResponseBody(r) {
				ww.Tee(&respBody)
			}

			start := time.Now()

			defer func() {
				var extra []string

				if rec := recover(); rec != nil {
					// Return HTTP 500 if recover is enabled and no response status was set.
					if o.RecoverPanics && ww.Status() == 0 && r.Header.Get("Connection") != "Upgrade" {
						ww.WriteHeader(http.StatusInternalServerError)
					}

					if rec == http.ErrAbortHandler || !o.RecoverPanics {
						// Re-panic http.ErrAbortHandler unconditionally, and re-panic other errors if panic recovery is disabled.
						defer panic(rec)
					}

					extra = append(extra, fmt.Sprintf("panic=%v", rec))

					if rec != http.ErrAbortHandler {
						pc := make([]uintptr, 10)   // Capture up to 10 stack frames.
						n := runtime.Callers(3, pc) // Skip 3 frames (this middleware + runtime/panic.go).
						pc = pc[:n]

						// Process panic stack frames to print detailed information.
						frames := runtime.CallersFrames(pc)
						var stackValues []string
						for frame, more := frames.Next(); more; frame, more = frames.Next() {
							if !strings.Contains(frame.File, "runtime/panic.go") {
								stackValues = append(stackValues, fmt.Sprintf("%s:%d", frame.File, frame.Line))
							}
						}
						extra = append(extra, fmt.Sprintf("stack=[%s]", strings.Join(stackValues, ", ")))
					}
				}

				duration := time.Since(start)
				statusCode := ww.Status()
				if statusCode == 0 {
					// If the handler never calls w.WriteHeader(statusCode) explicitly,
					// Go's http package automatically sends HTTP 200 OK to the client.
					statusCode = 200
				}

				// Skip logging if the request is filtered by the Skip function.
				if o.Skip != nil && o.Skip(r, statusCode) {
					return
				}

				extra = append(extra,
					fmt.Sprintf("remote=%s", r.RemoteAddr),
					fmt.Sprintf("bytes_in=%d", r.ContentLength),
					fmt.Sprintf("bytes_out=%d", ww.BytesWritten()),
				)

				// Log selected request headers.
				for _, h := range o.LogRequestHeaders {
					if v := r.Header.Get(h); v != "" {
						extra = append(extra, fmt.Sprintf("req.%s=%s", h, v))
					}
				}

				// Log selected response headers.
				for _, h := range o.LogResponseHeaders {
					if v := ww.Header().Get(h); v != "" {
						extra = append(extra, fmt.Sprintf("resp.%s=%s", h, v))
					}
				}

				if err := ctx.Err(); errors.Is(err, context.Canceled) {
					extra = append(extra, fmt.Sprintf("error=%v", ErrClientAborted))
				}

				if logReqBody || o.LogExtraAttrs != nil {
					// Ensure the request body is fully read if the underlying HTTP handler didn't do so.

					if n, _ := io.Copy(io.Discard, r.Body); n > 0 {
						extra = append(extra, fmt.Sprintf("bytes_unread=%d", n))
					}
				}
				if logReqBody {
					extra = append(extra, fmt.Sprintf("req_body=%s", logBody(&reqBody, r.Header, o)))
				}
				if logRespBody {
					extra = append(extra, fmt.Sprintf("resp_body=%s", logBody(&respBody, ww.Header(), o)))
				}
				if o.LogExtraAttrs != nil {
					extra = append(extra, o.LogExtraAttrs(r, reqBody.String(), statusCode)...)
				}
				extra = append(extra, getAttrs(ctx)...)

				logger.Printf("%s %s => HTTP %d (%v) %s", r.Method, r.URL, statusCode, duration, strings.Join(extra, " "))
			}()

			next.ServeHTTP(ww, r.WithContext(ctx))
		})
	}
}

func logBody(body *bytes.Buffer, header http.Header, o *Options) string {
	if body.Len() == 0 {
		return ""
	}
	contentType := header.Get("Content-Type")
	for _, whitelisted := range o.LogBodyContentTypes {
		if strings.HasPrefix(contentType, whitelisted) {
			if o.LogBodyMaxLen <= 0 || o.LogBodyMaxLen >= body.Len() {
				return body.String()
			}
			return body.String()[:o.LogBodyMaxLen] + "... [trimmed]"
		}
	}
	return fmt.Sprintf("[body redacted for Content-Type: %s]", contentType)
}
