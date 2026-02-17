package chiMiddleware

import (
	"context"
	"fmt"
)

type ctxKeyLogAttrs struct{}

func (c *ctxKeyLogAttrs) String() string {
	return "httplog attrs context"
}

// SetAttrs sets string attributes on the request log.
func SetAttrs(ctx context.Context, attrs ...string) {
	if ptr, ok := ctx.Value(ctxKeyLogAttrs{}).(*[]string); ok && ptr != nil {
		*ptr = append(*ptr, attrs...)
	}
}

func getAttrs(ctx context.Context) []string {
	if ptr, ok := ctx.Value(ctxKeyLogAttrs{}).(*[]string); ok && ptr != nil {
		return *ptr
	}

	return nil
}

// SetError sets the error attribute on the request log.
func SetError(ctx context.Context, err error) error {
	if err != nil {
		SetAttrs(ctx, fmt.Sprintf("error=%v", err))
	}

	return err
}
