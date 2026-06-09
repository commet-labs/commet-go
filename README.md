# Commet Go SDK

Billing and usage tracking for SaaS applications.

## Installation

```bash
go get github.com/commet-labs/commet-go/v6
```

## Quick start

```go
package main

import (
	"context"
	"log"

	commet "github.com/commet-labs/commet-go/v6"
)

func main() {
	client, err := commet.New("ck_xxx")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	ctx := context.Background()

	// Create a customer
	client.Customers.Create(ctx, &commet.CreateCustomerParams{
		Email:      "user@example.com",
		ExternalID: "user_123",
	})

	// Create a subscription
	client.Subscriptions.Create(ctx, &commet.CreateSubscriptionParams{
		ExternalID: "user_123",
		PlanCode:   "pro",
	})

	// Track usage
	client.Usage.Track(ctx, &commet.TrackUsageParams{
		Feature:    "api_calls",
		ExternalID: "user_123",
	})

	// Track AI token usage
	inputTokens := 1000
	outputTokens := 500
	client.Usage.Track(ctx, &commet.TrackUsageParams{
		Feature:      "ai_generation",
		ExternalID:   "user_123",
		Model:        "claude-sonnet-4-20250514",
		InputTokens:  &inputTokens,
		OutputTokens: &outputTokens,
	})
}
```

## Quota

Track a durable integer balance (e.g. projects, tasks) that carries across billing periods:

```go
// Add to the balance (Count defaults to 1)
client.Quota.Add(ctx, &commet.QuotaParams{
	CustomerID:  "user_123",
	FeatureCode: "tasks",
	Count:       5,
})

// Set the balance to an exact value
client.Quota.Set(ctx, &commet.QuotaParams{
	CustomerID:  "user_123",
	FeatureCode: "tasks",
	Count:       10,
})

// Remove from the balance (Count defaults to 1)
client.Quota.Remove(ctx, &commet.QuotaParams{
	CustomerID:  "user_123",
	FeatureCode: "tasks",
})

// Read the current allowance (held vs included, remaining)
client.Quota.Get(ctx, &commet.GetQuotaParams{
	CustomerID:  "user_123",
	FeatureCode: "tasks",
})

// Read every quota allowance for a customer
client.Quota.GetAll(ctx, &commet.GetAllQuotaParams{
	CustomerID: "user_123",
})
```

## Customer context

Scope all operations to a customer to avoid repeating `ExternalID`:

```go
customer := client.Customer("user_123")

customer.Usage.Track(ctx, "api_calls")
customer.Features.Check(ctx, "custom_branding")
customer.Seats.Add(ctx, "editor", 3)
customer.Portal.GetURL(ctx)
```

## Webhook verification

```go
webhooks := &commet.Webhooks{}

payload, err := webhooks.VerifyAndParse(
	requestBody,
	request.Header.Get("x-commet-signature"),
	"whsec_xxx",
)
if err != nil {
	log.Fatal("Invalid webhook signature")
}

if payload["event"] == "subscription.activated" {
	// handle activation
}
```

## License

MIT
