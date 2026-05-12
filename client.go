package commet

import (
	"errors"
	"strings"
	"time"
)

// Option configures the Commet client.
type Option func(*clientConfig)

type clientConfig struct {
	timeout   time.Duration
	retries   int
	telemetry bool
}

// WithTimeout sets the HTTP request timeout. Defaults to 30 seconds.
func WithTimeout(timeout time.Duration) Option {
	return func(c *clientConfig) {
		c.timeout = timeout
	}
}

// WithRetries sets the maximum number of retries for failed requests. Defaults to 3.
func WithRetries(retries int) Option {
	return func(c *clientConfig) {
		c.retries = retries
	}
}

// WithTelemetry enables or disables client telemetry headers. Defaults to true.
func WithTelemetry(enabled bool) Option {
	return func(c *clientConfig) {
		c.telemetry = enabled
	}
}

// Client is the Commet SDK client.
type Client struct {
	Customers     Customers
	Plans         Plans
	Subscriptions Subscriptions
	Usage         Usage
	Seats         Seats
	Features      Features
	Portal        Portal
	CreditPacks   CreditPacks
	Webhooks      WebhookVerifier

	http *httpClient
}

// New creates a new Commet client with the given API key and options.
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

	h := newHTTPClient(apiKey, cfg.timeout, cfg.retries, cfg.telemetry)

	c := &Client{
		http: h,
	}

	c.Customers = &CustomersResource{http: h}
	c.Plans = &PlansResource{http: h}
	c.Subscriptions = &SubscriptionsResource{http: h}
	c.Usage = &UsageResource{http: h}
	c.Seats = &SeatsResource{http: h}
	c.Features = &FeaturesResource{http: h}
	c.Portal = &PortalResource{http: h}
	c.CreditPacks = &CreditPacksResource{http: h}
	c.Webhooks = &Webhooks{}

	return c, nil
}

// Close releases resources held by the client.
func (c *Client) Close() {
	c.http.close()
}

// Customer returns a customer-scoped context for cleaner API usage.
func (c *Client) Customer(customerID string) *CustomerContext {
	return &CustomerContext{
		Features:     &CustomerFeatures{customerID: customerID, resource: c.Features},
		Seats:        &CustomerSeats{customerID: customerID, resource: c.Seats},
		Usage:        &CustomerUsage{customerID: customerID, resource: c.Usage},
		Subscription: &CustomerSubscription{customerID: customerID, resource: c.Subscriptions},
		Portal:       &CustomerPortal{customerID: customerID, resource: c.Portal},
	}
}
