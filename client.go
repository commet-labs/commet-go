package commet

import (
	"errors"
	"strings"
	"time"
)

// Environment represents the target API environment.
type Environment string

const (
	Sandbox    Environment = "sandbox"
	Production Environment = "production"
)

// Option configures the Commet client.
type Option func(*clientConfig)

type clientConfig struct {
	environment Environment
	timeout     time.Duration
	retries     int
}

// WithEnvironment sets the API environment. Defaults to Sandbox.
func WithEnvironment(env Environment) Option {
	return func(c *clientConfig) {
		c.environment = env
	}
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

// Client is the Commet SDK client.
type Client struct {
	Customers     *CustomersResource
	Plans         *PlansResource
	Subscriptions *SubscriptionsResource
	Usage         *UsageResource
	Seats         *SeatsResource
	Features      *FeaturesResource
	Portal        *PortalResource
	CreditPacks   *CreditPacksResource
	Webhooks      *Webhooks

	http        *httpClient
	environment Environment
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
		environment: Sandbox,
		timeout:     30 * time.Second,
		retries:     3,
	}

	for _, opt := range opts {
		opt(cfg)
	}

	if _, ok := baseURLs[cfg.environment]; !ok {
		return nil, errors.New("commet: invalid environment, must be 'sandbox' or 'production'")
	}

	h := newHTTPClient(apiKey, cfg.environment, cfg.timeout, cfg.retries)

	c := &Client{
		http:        h,
		environment: cfg.environment,
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

// Environment returns the configured environment.
func (c *Client) Environment() Environment {
	return c.environment
}

// IsSandbox returns true if the client is configured for sandbox.
func (c *Client) IsSandbox() bool {
	return c.environment == Sandbox
}

// IsProduction returns true if the client is configured for production.
func (c *Client) IsProduction() bool {
	return c.environment == Production
}
