package chiMiddleware

import (
	"net/http"
)

type Options struct {
	// RecoverPanics recovers from panics occurring in the underlying HTTP handlers
	// and middlewares and returns HTTP 500 unless response status was already set.
	//
	// NOTE: Panics are logged as errors automatically, regardless of this setting.
	RecoverPanics bool

	// Skip is an optional predicate function that determines whether to skip
	// recording logs for a given request.
	//
	// If nil, all requests are recorded.
	// If provided, requests where Skip returns true will not be recorded.
	Skip func(req *http.Request, respStatus int) bool

	// LogRequestHeaders is a list of headers to be logged.
	// If not provided, the default is ["Content-Type", "Origin"].
	//
	// WARNING: Do not leak any request headers with sensitive information.
	LogRequestHeaders []string

	// LogRequestBody is an optional predicate function that controls logging of request body.
	//
	// If the function returns true, the request body will be logged.
	// If false, no request body will be logged.
	//
	// WARNING: Do not leak any request bodies with sensitive information.
	LogRequestBody func(req *http.Request) bool

	// LogResponseHeaders controls a list of headers to be logged.
	//
	// If not provided, there are no default headers.
	LogResponseHeaders []string

	// LogResponseBody is an optional predicate function that controls logging of response body.
	//
	// If the function returns true, the response body will be logged.
	// If false, no response body will be logged.
	//
	// WARNING: Do not leak any response bodies with sensitive information.
	LogResponseBody func(req *http.Request) bool

	// LogBodyContentTypes defines a list of body Content-Types that are safe to be logged
	// with LogRequestBody or LogResponseBody options.
	//
	// If not provided, the default is ["application/json", "application/xml", "text/plain", "text/csv", "application/x-www-form-urlencoded", ""].
	LogBodyContentTypes []string

	// LogBodyMaxLen defines the maximum length of the body to be logged.
	//
	// If not provided, the default is 1024 bytes. Set to -1 to log the full body.
	LogBodyMaxLen int

	// LogExtraAttrs is an optional function that lets you add extra string attributes
	// to the request log.
	//
	// WARNING: Be careful not to leak any sensitive information in the logs.
	LogExtraAttrs func(req *http.Request, reqBody string, respStatus int) []string
}

var defaultOptions = Options{
	RecoverPanics:       true,
	LogRequestHeaders:   []string{"Content-Type", "Origin"},
	LogResponseHeaders:  []string{"Content-Type"},
	LogBodyContentTypes: []string{"application/json", "application/xml", "text/plain", "text/csv", "application/x-www-form-urlencoded", ""},
	LogBodyMaxLen:       1024,
}
