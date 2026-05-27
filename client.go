package commet

import (
	"errors"
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
	apiVersion string
	timeout    time.Duration
	retries    int
	telemetry  bool
	debug      bool
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

type Client struct {
	Customers     Customers
	Plans         Plans
	Subscriptions Subscriptions
	Usage         Usage
	Seats         Seats
	Features      Features
	Portal        Portal
	CreditPacks   CreditPacks
	Addons        Addons
	Webhooks      WebhookService
	ApiKeys       ApiKeys
	Invoices      Invoices
	Transactions  Transactions
	PromoCodes    PromoCodes
	PlanGroups    PlanGroups

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

	h := newHTTPClient(apiKey, cfg.apiVersion, cfg.timeout, cfg.retries, cfg.telemetry, cfg.debug)

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
	c.Addons = &AddonsResource{http: h}
	c.Webhooks = &WebhooksResource{http: h}
	c.ApiKeys = &ApiKeysResource{http: h}
	c.Invoices = &InvoicesResource{http: h}
	c.Transactions = &TransactionsResource{http: h}
	c.PromoCodes = &PromoCodesResource{http: h}
	c.PlanGroups = &PlanGroupsResource{http: h}

	return c, nil
}

func (c *Client) Close() {
	c.http.close()
}
