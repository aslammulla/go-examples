package httpclient

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

// ClientOption defines the function type for client configuration
type ClientOption func(*Client) error

// Client wraps an http.Client with retry and TLS support
type Client struct {
	httpClient    *http.Client
	retries       int
	retryInterval time.Duration
	baseURL       string
	headers       map[string]string
}

// New creates a new Client with options
func New(opts ...ClientOption) *Client {
	c := &Client{
		httpClient: &http.Client{},
		retries:    3, // default retries
		headers:    make(map[string]string),
	}

	// Apply options
	for _, opt := range opts {
		_ = opt(c)
	}

	return c
}

// WithTimeout sets the client timeout
func WithTimeout(timeout time.Duration) ClientOption {
	return func(c *Client) error {
		c.httpClient.Timeout = timeout
		return nil
	}
}

// WithRetry sets the retry count and interval
func WithRetry(retries int, interval time.Duration) ClientOption {
	return func(c *Client) error {
		c.retries = retries
		c.retryInterval = interval
		return nil
	}
}

// WithTransport sets a custom transport
func WithTransport(transport *http.Transport) ClientOption {
	return func(c *Client) error {
		c.httpClient.Transport = transport
		return nil
	}
}

// WithBaseURL sets the base URL for all requests
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		c.baseURL = baseURL
		return nil
	}
}

// WithDefaultHeaders sets default headers for all requests
func WithDefaultHeaders(headers map[string]string) ClientOption {
	return func(c *Client) error {
		for k, v := range headers {
			c.headers[k] = v
		}
		return nil
	}
}

// WithTLSConfig sets up TLS configuration
func WithTLSConfig(certFile, keyFile, caFile string) ClientOption {
	return func(c *Client) error {
		cert, err := tls.LoadX509KeyPair(certFile, keyFile)
		if err != nil {
			return fmt.Errorf("failed to load client cert: %w", err)
		}

		caCert, err := os.ReadFile(caFile)
		if err != nil {
			return fmt.Errorf("failed to read CA file: %w", err)
		}

		caCertPool := x509.NewCertPool()
		caCertPool.AppendCertsFromPEM(caCert)

		tlsConfig := &tls.Config{
			Certificates: []tls.Certificate{cert},
			RootCAs:      caCertPool,
			MinVersion:   tls.VersionTLS12,
		}

		transport := &http.Transport{TLSClientConfig: tlsConfig}
		c.httpClient.Transport = transport
		return nil
	}
}

// NewWithTLS creates a Client that uses a custom TLS certificate
func NewWithTLS(timeout time.Duration, retries int, certFile, keyFile, caFile string) (*Client, error) {
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return nil, fmt.Errorf("failed to load client cert: %w", err)
	}

	caCert, err := os.ReadFile(caFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read CA file: %w", err)
	}

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      caCertPool,
		MinVersion:   tls.VersionTLS12,
	}

	transport := &http.Transport{TLSClientConfig: tlsConfig}

	return &Client{
		httpClient: &http.Client{Timeout: timeout, Transport: transport},
		retries:    retries,
	}, nil
}

// doRequest executes an HTTP request with retry logic
func (c *Client) doRequest(req *http.Request) (*http.Response, error) {
	// Apply default headers
	for k, v := range c.headers {
		if req.Header.Get(k) == "" { // Don't override if already set
			req.Header.Set(k, v)
		}
	}

	var resp *http.Response
	var err error

	for attempt := 0; attempt <= c.retries; attempt++ {
		resp, err = c.httpClient.Do(req)
		if err == nil && (resp.StatusCode < 500 || attempt == c.retries) {
			return resp, nil
		}
		if err != nil {
			fmt.Printf("Attempt %d failed: %v\n", attempt+1, err)
		} else {
			resp.Body.Close()
		}

		// Use the configured retry interval or default to exponential backoff
		if c.retryInterval > 0 {
			time.Sleep(c.retryInterval)
		} else {
			time.Sleep(time.Duration(attempt+1) * time.Second)
		}
	}
	return nil, fmt.Errorf("request failed after %d retries: %w", c.retries, err)
}

// Get performs a GET request with variadic options
func (c *Client) Get(url string, opts ...RequestOption) (*http.Response, error) {
	if c.baseURL != "" {
		url = c.baseURL + url
	}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	for _, opt := range opts {
		if err := opt(req); err != nil {
			return nil, err
		}
	}
	return c.doRequest(req)
}

// Post performs a POST request with variadic options
func (c *Client) Post(url string, body []byte, opts ...RequestOption) (*http.Response, error) {
	if c.baseURL != "" {
		url = c.baseURL + url
	}
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	for _, opt := range opts {
		if err := opt(req); err != nil {
			return nil, err
		}
	}
	return c.doRequest(req)
}

// ReadBody helper to read response body
func ReadBody(resp *http.Response) ([]byte, error) {
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}
