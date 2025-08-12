package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"http-client-example/httpclient"
)

func main() {
	// Create a custom transport
	transport := &http.Transport{
		MaxIdleConns:        100,
		MaxIdleConnsPerHost: 100,
		IdleConnTimeout:     90 * time.Second,
	}

	// Create client with various options
	client := httpclient.New(
		httpclient.WithTimeout(10*time.Second),
		httpclient.WithRetry(3, time.Second),
		httpclient.WithTransport(transport),
		httpclient.WithBaseURL("https://jsonplaceholder.typicode.com"),
		httpclient.WithDefaultHeaders(map[string]string{
			"User-Agent":   "CustomHTTPClient/1.0",
			"Content-Type": "application/json",
		}),
	)

	// Context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// GET request example with query parameters
	resp, err := client.Get("/posts",
		httpclient.WithContext(ctx),
		httpclient.WithHeader("Accept", "application/json"),
		httpclient.WithQueryParam("userId", "1"),
	)
	if err != nil {
		log.Fatalf("GET request failed: %v", err)
	}
	body, _ := httpclient.ReadBody(resp)
	fmt.Println("GET Response:", string(body))

	// POST request example with all options
	postBody := []byte(`{"title":"foo","body":"bar","userId":1}`)
	resp, err = client.Post("/posts",
		postBody,
		httpclient.WithContext(ctx),
		httpclient.WithHeader("Content-Type", "application/json"),
		httpclient.WithBasicAuth("user", "password"),
	)
	if err != nil {
		log.Fatalf("POST request failed: %v", err)
	}
	body, _ = httpclient.ReadBody(resp)
	fmt.Println("POST Response:", string(body))

	// Example with TLS (commented out as it needs certificates)
	/*
		clientWithTLS := httpclient.New(
			httpclient.WithTimeout(10*time.Second),
			httpclient.WithTLSConfig(
				"path/to/cert.pem",
				"path/to/key.pem",
				"path/to/ca.pem",
			),
		)
	*/
}
