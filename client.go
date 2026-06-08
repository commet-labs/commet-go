package commet

import (
	"errors"
	"net/http"
	"strings"
	"time"
)

type RequestOptions struct {
	ApiVersion     string
	IdempotencyKey string
	Timeout        time.Duration
}

type Option func(*clientConfig)

type clientConfig struct {
	apiVersion       string
	timeout          time.Duration
	retries          int
	telemetry        bool
	debug            bool
	customHTTPClient *http.Client
}

func WithTimeout(timeout time.Duration) Option {
	return func(c *clientConfig) {
		c.timeout = timeout
	}
}

func WithRetries(retries int) Option {
	return func(c *clientConfig) {
		c.retries = retries
	}
}

func WithTelemetry(enabled bool) Option {
	return func(c *clientConfig) {
		c.telemetry = enabled
	}
}

func WithApiVersion(version string) Option {
	return func(c *clientConfig) {
		c.apiVersion = version
	}
}

func WithDebug(enabled bool) Option {
	return func(c *clientConfig) {
		c.debug = enabled
	}
}

// WithHTTPClient sets a custom *http.Client (e.g. for custom transports, proxies,
// or intercepting requests in tests). If unset, a default client is used.
func WithHTTPClient(client *http.Client) Option {
	return func(c *clientConfig) {
		c.customHTTPClient = client
	}
}

type Client struct {
	generatedResources

	// Preserved hand-written resources (not in the generated registry).
	Usage    *UsageResource
	Webhooks *WebhooksResource

	http *httpClient
}

func New(apiKey string, opts ...Option) (*Client, error) {
	if apiKey == "" {
		return nil, errors.New("commet: API key is required")
	}

	if !strings.HasPrefix(apiKey, "ck_") {
		return nil, errors.New("commet: invalid API key format, expected format: ck_xxx...")
	}

	cfg := &clientConfig{
		timeout:   30 * time.Second,
		retries:   3,
		telemetry: true,
	}

	for _, opt := range opts {
		opt(cfg)
	}

	h := newHTTPClient(apiKey, cfg.apiVersion, cfg.timeout, cfg.retries, cfg.telemetry, cfg.debug, cfg.customHTTPClient)

	c := &Client{
		http: h,
	}

	c.wireResources(h)

	c.Usage = &UsageResource{http: h}
	c.Webhooks = &WebhooksResource{http: h}

	return c, nil
}

func (c *Client) Close() {
	c.http.close()
}
