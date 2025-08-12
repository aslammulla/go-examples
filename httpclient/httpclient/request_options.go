package httpclient

import (
	"context"
	"net/http"
)

// RequestOption is a function type that modifies an HTTP request
type RequestOption func(*http.Request) error

// WithHeader adds a header to the request
func WithHeader(key, value string) RequestOption {
	return func(req *http.Request) error {
		req.Header.Set(key, value)
		return nil
	}
}

// WithQueryParam adds a query parameter to the request
func WithQueryParam(key, value string) RequestOption {
	return func(req *http.Request) error {
		q := req.URL.Query()
		q.Add(key, value)
		req.URL.RawQuery = q.Encode()
		return nil
	}
}

// WithContext sets a context for the request
func WithContext(ctx context.Context) RequestOption {
	return func(req *http.Request) error {
		*req = *req.WithContext(ctx)
		return nil
	}
}

// WithBasicAuth adds basic auth to the request
func WithBasicAuth(username, password string) RequestOption {
	return func(req *http.Request) error {
		req.SetBasicAuth(username, password)
		return nil
	}
}
