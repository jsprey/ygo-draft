package ygoprodeck

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"golang.org/x/time/rate"
)

// RLHTTPClient Rate Limited HTTP Client.
type RLHTTPClient struct {
	Client      HttpClient
	RateLimiter Limiter
}

// Limiter is responsible to block outgoing http requests to abide to a defined rate limit.
type Limiter interface {
	Wait(ctx context.Context) error
}

// HttpClient is responsible to retrieve files from a server via http.
type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
	Get(url string) (resp *http.Response, err error)
}

// NewDefaultRateLimitedClient creates a new rate limited client.
func NewDefaultRateLimitedClient() *RLHTTPClient {
	rl := rate.NewLimiter(rate.Every(60*time.Second), 600) // a maximum of 10 request every second

	return &RLHTTPClient{
		Client:      http.DefaultClient,
		RateLimiter: rl,
	}
}

// Do dispatches the HTTP request to the network
func (c *RLHTTPClient) Do(req *http.Request) (*http.Response, error) {
	// Comment out the below 5 lines to turn off ratelimiting
	ctx := context.Background()

	err := c.RateLimiter.Wait(ctx) // This is a blocking call. Honors the rate limit
	if err != nil {
		return nil, err
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetJsonFromTarget makes a get request to the target urls and unmarshal the body as given data type.
func (c *RLHTTPClient) GetJsonFromTarget(targetUrl string, data any) error {
	resp, err := c.Client.Get(targetUrl)
	if err != nil {
		return fmt.Errorf("failed to get [%s]: %w", targetUrl, err)
	}

	body, err := io.ReadAll(resp.Body)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)
	if err != nil {
		return fmt.Errorf("failed read body: %w", err)
	}

	err = json.Unmarshal(body, data)
	if err != nil {
		return fmt.Errorf("failed to unmarshal body: %w", err)
	}

	return nil
}
