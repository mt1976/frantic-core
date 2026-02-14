package chiMiddleware

import (
	"net/http"

	"github.com/alitto/pond/v2"
	"github.com/mt1976/frantic-core/contextHandler"
	"github.com/mt1976/frantic-core/logHandler"
)

func InjectWorkerPoolIntoContext(next http.Handler, workerPool pond.Pool) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			logHandler.Trace.Println("Adding worker pool to context")
			ctx := r.Context()
			ctx = contextHandler.AddWorkerPoolToContext(ctx, workerPool)
			logHandler.Trace.Printf("Worker Pool added to context")
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
}
