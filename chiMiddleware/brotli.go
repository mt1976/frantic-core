package chiMiddleware

import (
	"io"
	"net/http"

	"github.com/go-chi/chi/middleware"
	brotli_enc "gopkg.in/kothar/brotli-go.v0/enc"
)

func HandleBrotli(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if the request accepts Brotli encoding
		compressor := middleware.NewCompressor(5, "/*")
		compressor.SetEncoder("br", func(w io.Writer, level int) io.Writer {
			params := brotli_enc.NewBrotliParams()
			params.SetQuality(level)
			return brotli_enc.NewBrotliWriter(params, w)
		})
	})
}
