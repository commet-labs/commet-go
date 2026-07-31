# Commet Go SDK

Billing and usage tracking for SaaS applications.

## Installation

```bash
go get github.com/commet-labs/commet-go/v9
```

## Quick start

```go
package main

import (
	"context"
	"log"

	commet "github.com/commet-labs/commet-go/v9"
)

func main() {
	client, err := commet.New("ck_xxx")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	ctx := context.Background()

	// Create a customer
	customer, err := client.Customers.Create(ctx, &commet.CreateCustomerParams{
		Email: "user@example.com",
	})
	if err != nil {
		log.Fatal(err)
	}

	// Create a subscription
	planCode := "pro"
	client.Subscriptions.Create(ctx, &commet.CreateSubscriptionParams{
		CustomerID: customer.ID,
		PlanCode:   &planCode,
	})

	// Track usage
	value := 1.0
	client.Usage.Track(ctx, &commet.TrackUsageParams{
		FeatureCode: "api_calls",
		CustomerID:  customer.ID,
		Value:       &value,
	})

	// Track AI token usage
	model := "claude-sonnet-4-20250514"
	inputTokens := int64(1000)
	outputTokens := int64(500)
	client.Usage.Track(ctx, &commet.TrackUsageParams{
		FeatureCode:  "ai_generation",
		CustomerID:   customer.ID,
		Model:        &model,
		InputTokens:  &inputTokens,
		OutputTokens: &outputTokens,
	})
}
```

## Offers and pricing Markets

SDK v9 exposes independent Offers, top-level Markets, and selectable `PriceID` variants:

```go
market, err := client.Markets.Create(ctx, &commet.CreateMarketParams{
	Name:         "Argentina",
	CountryCodes: []string{"AR"},
})

offer, err := client.Offers.Create(ctx, &commet.CreateOfferParams{
	Name: "30-day trial",
	Phases: []commet.CreateOfferParamsPhasesItem{{
		Value: commet.CreateOfferParamsPhasesItemVariant1{
			Type:         "free_trial",
			DurationDays: 30,
		},
	}},
})
```

Promo Codes reference compatible Offers. Omitting `PriceID` during subscription creation keeps normal default-price resolution.

## Quota

Track a durable integer balance (e.g. projects, tasks) that carries across billing periods:

```go
// Add to the balance (Count defaults to 1)
customerID := customer.ID
count := 5
client.Quota.Add(ctx, &commet.AddQuotaParams{
	CustomerID:  &customerID,
	FeatureCode: "tasks",
	Count:       &count,
})

// Set the balance to an exact value
client.Quota.Set(ctx, &commet.SetQuotaParams{
	CustomerID:  &customerID,
	FeatureCode: "tasks",
	Count:       10,
})

// Remove from the balance (Count defaults to 1)
client.Quota.Remove(ctx, &commet.RemoveQuotaParams{
	CustomerID:  &customerID,
	FeatureCode: "tasks",
})

// Read the current allowance (held vs included, remaining)
client.Quota.Get(ctx, &commet.GetQuotaAllowanceParams{
	CustomerID:  customer.ID,
	FeatureCode: "tasks",
})

// Read every quota allowance for a customer
client.Quota.GetAll(ctx, &commet.GetAllQuotaAllowancesParams{
	CustomerID: customer.ID,
})
```

## Webhook verification

```go
webhooks := &commet.WebhooksResource{}

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
