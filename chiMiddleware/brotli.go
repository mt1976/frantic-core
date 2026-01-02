package chiMiddleware

import (
	"io"
	"net/http"

	brotli_enc "github.com/andybalholm/brotli"
	"github.com/go-chi/chi/middleware"
)

func HandleBrotli(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if the request accepts Brotli encoding
		compressor := middleware.NewCompressor(5, "/*")
		compressor.SetEncoder("br", func(w io.Writer, level int) io.Writer {

			writer := brotli_enc.NewWriter(w)
			defer writer.Close()

			return writer
			//	params := brotli_enc.
			//		params.SetQuality(level)
			///return brotli_enc.Encoder(w, w, "", false)
		})
	})
}
