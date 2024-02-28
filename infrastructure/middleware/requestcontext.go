package middleware

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

var requestContextKey = contextKey("requestContext")

// RequestContext is the struct representation for an http request's details
type RequestContext struct {
	RequestID string `json:"request_id"`
	Method    string `json:"method"`
	URI       string `json:"uri"`
	UserAgent string `json:"user_agent"`
	Trace     string `json:"trace"`
}

// ToRequestInfo returns a map to represent the request context's information
func (r RequestContext) ToRequestInfo() map[string]string {
	return map[string]string{
		"request": r.RequestID,
		"route":   r.Method + " " + r.URI,
		"agent":   r.UserAgent,
		"trace":   r.Trace,
	}
}

// NewRequestContext creates a RequestContext from an http.Request
func NewRequestContext(r *http.Request) RequestContext {
	result := RequestContext{
		RequestID: uuid.NewString(),
		Method:    r.Method,
		URI:       r.RequestURI,
		UserAgent: r.UserAgent(),
	}
	// // Retrieve trace ID from the span context
	// if traceID, ok := tracing.GetID(r.Context()); ok {
	// 	result.Trace = traceID
	// }
	return result
}

// GetRequestContext returns the request context associated with the provided context
func GetRequestContext(ctx context.Context) RequestContext {
	if v, ok := ctx.Value(requestContextKey).(RequestContext); ok {
		return v
	}
	return RequestContext{}
}

// SetRequestContext sets the RequestContext into the provided context
func SetRequestContext(ctx context.Context, value RequestContext) context.Context {
	return context.WithValue(ctx, requestContextKey, value)
}
