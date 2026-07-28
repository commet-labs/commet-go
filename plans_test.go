package commet

import (
	"context"
	"testing"
)

func TestAddPriceSendsCamelCaseRequest(t *testing.T) {
	client, captured := newWireServer(t, 200, `{"id":"price_1","planId":"plan_1","billingInterval":"monthly","price":1000,"isDefault":true,"trialDays":0}`)

	_, err := client.Plans.AddPrice(context.Background(), "plan_1", &AddPlanPriceParams{
		BillingInterval: "monthly",
		Price:           intPtr(1000),
		TrialDays:       intPtr(14),
	})
	if err != nil {
		t.Fatalf("AddPrice: %v", err)
	}

	body := decodeBody(t, captured.Body)

	if body["billingInterval"] != "monthly" {
		t.Errorf("billingInterval = %v, want monthly (enum->wire string failed) (raw: %s)", body["billingInterval"], captured.Body)
	}
	if body["trialDays"] != float64(14) {
		t.Errorf("trialDays = %v, want 14 (raw: %s)", body["trialDays"], captured.Body)
	}
	if _, exists := body["trial_days"]; exists {
		t.Errorf("snake_case trial_days leaked: %s", captured.Body)
	}
}

// TestSetRegionalPricesSendsTypedArrayBodyAsCamel verifies a body whose top-level
// field is a slice of typed structs ([]UpsertRegionalPricesParamsOverridesItem)
// camelizes each array element's keys, and the optional *int IncludedBalance is
// omitted on the element where it is nil but present where set.
func TestSetRegionalPricesSendsTypedArrayBodyAsCamel(t *testing.T) {
	client, captured := newWireServer(t, 200, `{"priceId":"price_1","overrides":[],"object":"plan_regional_pricing","livemode":false}`)

	_, err := client.Plans.SetRegionalPrices(context.Background(), "plan_1", "price_1", &UpsertRegionalPricesParams{
		Overrides: []UpsertRegionalPricesParamsOverridesItem{
			{Currency: "eur", Price: 900, IncludedBalance: intPtr(500)},
			{Currency: "gbp", Price: 800}, // IncludedBalance nil
		},
	})
	if err != nil {
		t.Fatalf("SetRegionalPrices: %v", err)
	}

	body := decodeBody(t, captured.Body)
	overrides, ok := body["overrides"].([]any)
	if !ok {
		t.Fatalf("overrides = %v, want array (raw: %s)", body["overrides"], captured.Body)
	}
	if len(overrides) != 2 {
		t.Fatalf("len(overrides) = %d, want 2", len(overrides))
	}

	first := overrides[0].(map[string]any)
	if first["currency"] != "eur" || first["price"] != float64(900) {
		t.Errorf("overrides[0] = %v, want currency=eur price=900 (raw: %s)", first, captured.Body)
	}
	if first["includedBalance"] != float64(500) {
		t.Errorf("overrides[0].includedBalance = %v, want 500 (raw: %s)", first["includedBalance"], captured.Body)
	}
	if _, exists := first["included_balance"]; exists {
		t.Errorf("nested snake_case included_balance leaked in array element: %s", captured.Body)
	}

	second := overrides[1].(map[string]any)
	if _, exists := second["includedBalance"]; exists {
		t.Errorf("overrides[1].includedBalance should be omitted when nil, got %v (raw: %s)", second["includedBalance"], captured.Body)
	}
}

// TestSetRegionalPricingSendsExchangeRateAndNestedArrays covers a float exchange
// rate plus typed price arrays.
func TestSetRegionalPricingSendsExchangeRateAndNestedArrays(t *testing.T) {
	client, captured := newWireServer(t, 200, `{"planId":"plan_1","currency":"eur","exchangeRate":0.92,"pricesConfigured":1,"featuresConfigured":0,"object":"plan_regional_pricing_result","livemode":false}`)

	_, err := client.Plans.SetRegionalPricing(context.Background(), "plan_1", &SetPlanRegionalPricingParams{
		Currency:     "eur",
		ExchangeRate: 0.92,
		Prices: []SetPlanRegionalPricingParamsPricesItem{
			{PriceID: "price_1", Price: 920, IncludedBalance: intPtr(1000)},
		},
	})
	if err != nil {
		t.Fatalf("SetRegionalPricing: %v", err)
	}

	body := decodeBody(t, captured.Body)
	if body["exchangeRate"] != 0.92 {
		t.Errorf("exchangeRate = %v, want 0.92 (raw: %s)", body["exchangeRate"], captured.Body)
	}
	if _, exists := body["exchange_rate"]; exists {
		t.Errorf("snake_case exchange_rate leaked: %s", captured.Body)
	}
}

// TestGetPlanUnmarshalsNestedPricesFeaturesAndEnums verifies a deeply-nested
// camelCase response decodes into the typed Plan: nested prices (with offer IDs
// and regional prices), nested features (with typed FeatureType enum and regional
// prices), the *ConsumptionModel enum pointer, and nullable pointers.
func TestGetPlanUnmarshalsNestedPricesFeaturesAndEnums(t *testing.T) {
	client, _ := newWireServer(t, 200, `{
		"id":"plan_1",
		"name":"Pro",
		"code":"pro",
		"description":null,
		"consumptionModel":"metered",
		"isPublic":true,
		"isDefault":false,
		"isFree":false,
		"blockOnExhaustion":null,
		"sortOrder":1,
		"planGroupId":null,
		"metadata":{},
		"createdAt":"2026-01-01",
		"updatedAt":"2026-01-01",
		"object":"plan",
		"livemode":true,
		"prices":[
			{"id":"price_1","billingInterval":"monthly","price":1000,"isDefault":true,"trialDays":14,"includedBalance":null,"includedCredits":null,
				"offerId":"offer_1",
				"regionalPrices":[{"currency":"eur","price":920,"includedBalance":null,"autoSynced":true}]}
		],
		"features":[
			{"code":"seats","name":"Seats","type":"seats","unitName":null,"enabled":true,"includedAmount":5,"unlimited":false,
				"overage":{"enabled":true,"model":"per_unit","unitPrice":100},
				"regionalPrices":[{"currency":"eur","overageUnitPrice":92,"autoSynced":true}]}
		]
	}`)

	resp, err := client.Plans.Get(context.Background(), "pro")
	if err != nil {
		t.Fatalf("Get: %v", err)
	}
	plan := resp

	if plan.Description != nil {
		t.Errorf("Description = %v, want nil for wire null", *plan.Description)
	}
	if plan.ConsumptionModel == nil || *plan.ConsumptionModel != ConsumptionModelMetered {
		t.Errorf("ConsumptionModel = %v, want metered enum", plan.ConsumptionModel)
	}
	if plan.BlockOnExhaustion != nil {
		t.Errorf("BlockOnExhaustion = %v, want nil for wire null", *plan.BlockOnExhaustion)
	}

	if len(plan.Prices) != 1 {
		t.Fatalf("len(Prices) = %d, want 1", len(plan.Prices))
	}
	price := plan.Prices[0]
	if price.BillingInterval != BillingIntervalMonthly {
		t.Errorf("Prices[0].BillingInterval = %q, want monthly enum", price.BillingInterval)
	}
	if price.IncludedBalance != nil {
		t.Errorf("Prices[0].IncludedBalance = %v, want nil for wire null", *price.IncludedBalance)
	}
	if price.OfferID == nil || *price.OfferID != "offer_1" {
		t.Errorf("Prices[0].OfferID = %v, want offer_1", price.OfferID)
	}
	if len(price.RegionalPrices) != 1 || price.RegionalPrices[0].Currency != "eur" {
		t.Errorf("Prices[0].RegionalPrices = %v, want one eur entry", price.RegionalPrices)
	}

	if len(plan.Features) != 1 {
		t.Fatalf("len(Features) = %d, want 1", len(plan.Features))
	}
	feature := plan.Features[0]
	if feature.Type != FeatureTypeSeats {
		t.Errorf("Features[0].Type = %q, want seats enum", feature.Type)
	}
	if feature.IncludedAmount == nil || *feature.IncludedAmount != 5 {
		t.Errorf("Features[0].IncludedAmount = %v, want 5", feature.IncludedAmount)
	}
	if feature.Overage == nil || feature.Overage.UnitPrice == nil || *feature.Overage.UnitPrice != 100 {
		t.Errorf("Features[0].Overage.UnitPrice = %v, want 100", feature.Overage)
	}
	if len(feature.RegionalPrices) != 1 || feature.RegionalPrices[0].OverageUnitPrice == nil || *feature.RegionalPrices[0].OverageUnitPrice != 92 {
		t.Errorf("Features[0].RegionalPrices = %v, want one entry overageUnitPrice=92", feature.RegionalPrices)
	}
}
