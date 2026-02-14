package chiMiddleware

import (
	"net/http"

	"github.com/mt1976/frantic-core/contextHandler"
	"github.com/mt1976/frantic-core/logHandler"
)

func InjectUserContext(next http.Handler, uid string, username string) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			logHandler.Trace.Println("Injecting user context into request")
			ctx := r.Context()
			ctx = contextHandler.SetSession_UserKey(ctx, uid)
			ctx = contextHandler.SetSession_UserCode(ctx, username)
			logHandler.Trace.Printf("User Context Injected: %v=%v", contextHandler.GetSession_UserKey(ctx), contextHandler.GetSession_UserCode(ctx))
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
}
