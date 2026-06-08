package commet

import (
	"context"
	"testing"
)

// TestAddPriceSendsNestedIntroOfferWithEnumAndCamelKeys covers a nested typed
// struct (AddPlanPriceParamsIntroOffer) carrying a typed enum pointer
// (*DiscountType). The enum must serialize as its wire string ("percentage") and
// the nested keys must be camelCase. BillingInterval (top-level enum) likewise.
func TestAddPriceSendsNestedIntroOfferWithEnumAndCamelKeys(t *testing.T) {
	client, captured := newWireServer(t, 200, `{"data":{"id":"price_1","plan_id":"plan_1","billing_interval":"monthly","price":1000,"is_default":true,"trial_days":0}}`)

	_, err := client.Plans.AddPrice(context.Background(), "plan_1", &AddPlanPriceParams{
		BillingInterval: BillingIntervalMonthly,
		Price:           1000,
		TrialDays:       intPtr(14),
		IntroOffer: &AddPlanPriceParamsIntroOffer{
			Enabled:        true,
			DiscountType:   discountTypePtr(DiscountTypePercentage),
			DiscountValue:  intPtr(2000),
			DurationCycles: intPtr(3),
		},
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

	offer, ok := body["introOffer"].(map[string]any)
	if !ok {
		t.Fatalf("introOffer = %v, want nested object (raw: %s)", body["introOffer"], captured.Body)
	}
	if offer["discountType"] != "percentage" {
		t.Errorf("introOffer.discountType = %v, want percentage (typed enum did not serialize as string) (raw: %s)", offer["discountType"], captured.Body)
	}
	if offer["durationCycles"] != float64(3) {
		t.Errorf("introOffer.durationCycles = %v, want 3 (raw: %s)", offer["durationCycles"], captured.Body)
	}
	if _, exists := offer["discount_type"]; exists {
		t.Errorf("nested snake_case introOffer.discount_type leaked: %s", captured.Body)
	}
	if _, exists := offer["duration_cycles"]; exists {
		t.Errorf("nested snake_case introOffer.duration_cycles leaked: %s", captured.Body)
	}
}

// TestSetRegionalPricesSendsTypedArrayBodyAsCamel verifies a body whose top-level
// field is a slice of typed structs ([]UpsertRegionalPricesParamsOverridesItem)
// camelizes each array element's keys, and the optional *int IncludedBalance is
// omitted on the element where it is nil but present where set.
func TestSetRegionalPricesSendsTypedArrayBodyAsCamel(t *testing.T) {
	client, captured := newWireServer(t, 200, `{"data":{"price_id":"price_1","overrides":[],"object":"plan_regional_pricing","livemode":false}}`)

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

// TestSetRegionalPricingSendsExchangeRateAndNestedArrays covers the richest plan
// request: a float exchange rate plus three typed-struct arrays (prices, features,
// intro offers). Asserts the float survives, array element enum (DiscountType)
// serializes, and snake keys do not leak at the array-element level.
func TestSetRegionalPricingSendsExchangeRateAndNestedArrays(t *testing.T) {
	client, captured := newWireServer(t, 200, `{"data":{"plan_id":"plan_1","currency":"eur","exchange_rate":0.92,"prices_configured":1,"features_configured":0,"object":"plan_regional_pricing_result","livemode":false}}`)

	_, err := client.Plans.SetRegionalPricing(context.Background(), "plan_1", &SetPlanRegionalPricingParams{
		Currency:     "eur",
		ExchangeRate: 0.92,
		Prices: []SetPlanRegionalPricingParamsPricesItem{
			{PriceID: "price_1", Price: 920, IncludedBalance: intPtr(1000)},
		},
		IntroOffers: []SetPlanRegionalPricingParamsIntroOffersItem{
			{PriceID: "price_1", DiscountType: DiscountTypeAmount, DiscountValue: 500, DurationCycles: 2},
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

	introOffers, ok := body["introOffers"].([]any)
	if !ok || len(introOffers) != 1 {
		t.Fatalf("introOffers = %v, want 1-element array (raw: %s)", body["introOffers"], captured.Body)
	}
	offer := introOffers[0].(map[string]any)
	if offer["priceId"] != "price_1" {
		t.Errorf("introOffers[0].priceId = %v, want price_1 (raw: %s)", offer["priceId"], captured.Body)
	}
	if offer["discountType"] != "amount" {
		t.Errorf("introOffers[0].discountType = %v, want amount (enum in array did not serialize) (raw: %s)", offer["discountType"], captured.Body)
	}
	if _, exists := offer["discount_type"]; exists {
		t.Errorf("snake_case discount_type leaked inside array element: %s", captured.Body)
	}
}

// TestGetPlanUnmarshalsNestedPricesFeaturesAndEnums verifies a deeply-nested
// camelCase response decodes into the typed Plan: nested prices (with intro offer
// and regional prices), nested features (with typed FeatureType enum and regional
// prices), the *ConsumptionModel enum pointer, and nullable pointers.
func TestGetPlanUnmarshalsNestedPricesFeaturesAndEnums(t *testing.T) {
	client, _ := newWireServer(t, 200, `{"data":{
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
			{"billingInterval":"monthly","price":1000,"isDefault":true,"trialDays":14,"includedBalance":null,"includedCredits":null,
				"introOffer":{"enabled":true,"discountType":"percentage","discountValue":2000,"durationCycles":3},
				"regionalPrices":[{"currency":"eur","price":920,"includedBalance":null,"autoSynced":true}]}
		],
		"features":[
			{"code":"seats","name":"Seats","type":"seats","unitName":null,"enabled":true,"includedAmount":5,"unlimited":false,
				"overage":{"enabled":true,"model":"per_unit","unitPrice":100},
				"regionalPrices":[{"currency":"eur","overageUnitPrice":92,"autoSynced":true}]}
		]
	}}`)

	resp, err := client.Plans.Get(context.Background(), "pro")
	if err != nil {
		t.Fatalf("Get: %v", err)
	}
	plan := resp.Data

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
	if price.IntroOffer == nil {
		t.Fatal("Prices[0].IntroOffer = nil, want decoded nested object")
	}
	if price.IntroOffer.DiscountType == nil || *price.IntroOffer.DiscountType != DiscountTypePercentage {
		t.Errorf("Prices[0].IntroOffer.DiscountType = %v, want percentage enum", price.IntroOffer.DiscountType)
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

func discountTypePtr(d DiscountType) *DiscountType { return &d }
