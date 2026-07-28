package commet

import (
	"context"
	"fmt"
)

type UpdatePlanFeatureParams struct {
	Enabled        *bool                           `json:"enabled,omitempty"`
	IncludedAmount *int                            `json:"included_amount,omitempty"`
	Unlimited      *bool                           `json:"unlimited,omitempty"`
	Overage        *UpdatePlanFeatureParamsOverage `json:"overage,omitempty"`
	CreditsPerUnit *int                            `json:"credits_per_unit,omitempty"`
	IdempotencyKey string                          `json:"-"`
}

type AddPlanFeatureParams struct {
	FeatureID      string                       `json:"feature_id"`
	Enabled        *bool                        `json:"enabled,omitempty"`
	IncludedAmount *int                         `json:"included_amount,omitempty"`
	Unlimited      *bool                        `json:"unlimited,omitempty"`
	Overage        *AddPlanFeatureParamsOverage `json:"overage,omitempty"`
	CreditsPerUnit *int                         `json:"credits_per_unit,omitempty"`
	PricingMode    *string                      `json:"pricing_mode,omitempty"`
	Margin         *int                         `json:"margin,omitempty"`
	IdempotencyKey string                       `json:"-"`
}

type SetDefaultPlanPriceParams struct {
	IdempotencyKey string `json:"-"`
}

type UpsertRegionalPricesParams struct {
	Overrides      []UpsertRegionalPricesParamsOverridesItem `json:"overrides"`
	IdempotencyKey string                                    `json:"-"`
}

type UpdatePlanPriceParams struct {
	Price           *int                                    `json:"price,omitempty"`
	IsDefault       *bool                                   `json:"is_default,omitempty"`
	TrialDays       *int                                    `json:"trial_days,omitempty"`
	IncludedBalance *int                                    `json:"included_balance,omitempty"`
	IncludedCredits *int                                    `json:"included_credits,omitempty"`
	Metadata        map[string]any                          `json:"metadata,omitempty"`
	MarketPrices    []UpdatePlanPriceParamsMarketPricesItem `json:"market_prices,omitempty"`
	IdempotencyKey  string                                  `json:"-"`
}

type AddPlanPriceParams struct {
	BillingInterval     string                               `json:"billing_interval"`
	Metadata            map[string]any                       `json:"metadata,omitempty"`
	Price               *int                                 `json:"price,omitempty"`
	TrialDays           *int                                 `json:"trial_days,omitempty"`
	IsDefault           *bool                                `json:"is_default,omitempty"`
	IncludedBalance     *int                                 `json:"included_balance,omitempty"`
	IncludedCredits     *int                                 `json:"included_credits,omitempty"`
	MarketPrices        []AddPlanPriceParamsMarketPricesItem `json:"market_prices,omitempty"`
	InheritsFromPriceID *string                              `json:"inherits_from_price_id,omitempty"`
	IdempotencyKey      string                               `json:"-"`
}

type SetPlanRegionalPricingParams struct {
	Currency       string                                     `json:"currency"`
	ExchangeRate   float64                                    `json:"exchange_rate"`
	Prices         []SetPlanRegionalPricingParamsPricesItem   `json:"prices,omitempty"`
	Features       []SetPlanRegionalPricingParamsFeaturesItem `json:"features,omitempty"`
	IdempotencyKey string                                     `json:"-"`
}

type UpdatePlanParams struct {
	Name           *string        `json:"name,omitempty"`
	Description    *string        `json:"description,omitempty"`
	Metadata       map[string]any `json:"metadata,omitempty"`
	IsPublic       *bool          `json:"is_public,omitempty"`
	IdempotencyKey string         `json:"-"`
}

type SetPlanVisibilityParams struct {
	IsPublic       bool   `json:"is_public"`
	IdempotencyKey string `json:"-"`
}

type ListPlansParams struct {
	IncludePrivate *bool `json:"include_private,omitempty"`
}

type CreatePlanParams struct {
	Name              string         `json:"name"`
	Code              string         `json:"code"`
	Description       *string        `json:"description,omitempty"`
	ConsumptionModel  *string        `json:"consumption_model,omitempty"`
	IsPublic          *bool          `json:"is_public,omitempty"`
	IsFree            *bool          `json:"is_free,omitempty"`
	BlockOnExhaustion *bool          `json:"block_on_exhaustion,omitempty"`
	PlanGroupID       *string        `json:"plan_group_id,omitempty"`
	Metadata          map[string]any `json:"metadata,omitempty"`
	IdempotencyKey    string         `json:"-"`
}

type PlansResource struct {
	http *httpClient
}

// Update limits, overage, or enabled status of a feature on a plan.
func (r *PlansResource) UpdateFeature(ctx context.Context, id string, featureID string, params *UpdatePlanFeatureParams) (*PlanFeature, error) {
	body := buildBody(map[string]any{
		"enabled":          params.Enabled,
		"included_amount":  params.IncludedAmount,
		"unlimited":        params.Unlimited,
		"overage":          params.Overage,
		"credits_per_unit": params.CreditsPerUnit,
	})
	return parseDirectResponse[PlanFeature](r.http.patch(ctx, fmt.Sprintf("/plans/%s/features/%s", id, featureID), body, params.IdempotencyKey))
}

// Detach a feature from a plan.
func (r *PlansResource) RemoveFeature(ctx context.Context, id string, featureID string) (*RemovedPlanFeature, error) {
	return parseDirectResponse[RemovedPlanFeature](r.http.delete(ctx, fmt.Sprintf("/plans/%s/features/%s", id, featureID), nil, ""))
}

// Attach a feature to a plan with limits, overage, and credits configuration.
func (r *PlansResource) AddFeature(ctx context.Context, id string, params *AddPlanFeatureParams) (*PlanFeature, error) {
	body := buildBody(map[string]any{
		"feature_id":       params.FeatureID,
		"enabled":          params.Enabled,
		"included_amount":  params.IncludedAmount,
		"unlimited":        params.Unlimited,
		"overage":          params.Overage,
		"credits_per_unit": params.CreditsPerUnit,
		"pricing_mode":     params.PricingMode,
		"margin":           params.Margin,
	})
	return parseDirectResponse[PlanFeature](r.http.post(ctx, fmt.Sprintf("/plans/%s/features", id), body, params.IdempotencyKey))
}

// Set a specific price as the default and return the updated plan price.
func (r *PlansResource) SetDefaultPrice(ctx context.Context, id string, priceID string, params *SetDefaultPlanPriceParams) (*PlanPrice, error) {
	return parseDirectResponse[PlanPrice](r.http.put(ctx, fmt.Sprintf("/plans/%s/prices/%s/default", id, priceID), map[string]any{}, params.IdempotencyKey))
}

// Create or update regional currency price overrides for a plan price.
func (r *PlansResource) SetRegionalPrices(ctx context.Context, id string, priceID string, params *UpsertRegionalPricesParams) (*PlanRegionalPricing, error) {
	body := buildBody(map[string]any{
		"overrides": params.Overrides,
	})
	return parseDirectResponse[PlanRegionalPricing](r.http.put(ctx, fmt.Sprintf("/plans/%s/prices/%s/regional", id, priceID), body, params.IdempotencyKey))
}

// Remove all regional currency overrides for a plan price. The request is rejected while billable subscriptions depend on an override.
func (r *PlansResource) DeleteRegionalPrices(ctx context.Context, id string, priceID string) (*DeletedPlanRegionalPricing, error) {
	return parseDirectResponse[DeletedPlanRegionalPricing](r.http.delete(ctx, fmt.Sprintf("/plans/%s/prices/%s/regional", id, priceID), nil, ""))
}

// Update a base price or market price variant. Removing a base market override is rejected while a variant depends on it. Offer terms are managed through Offers.
func (r *PlansResource) UpdatePrice(ctx context.Context, id string, priceID string, params *UpdatePlanPriceParams) (*PlanPrice, error) {
	body := buildBody(map[string]any{
		"price":            params.Price,
		"is_default":       params.IsDefault,
		"trial_days":       params.TrialDays,
		"included_balance": params.IncludedBalance,
		"included_credits": params.IncludedCredits,
		"metadata":         params.Metadata,
		"market_prices":    params.MarketPrices,
	})
	return parseDirectResponse[PlanPrice](r.http.patch(ctx, fmt.Sprintf("/plans/%s/prices/%s", id, priceID), body, params.IdempotencyKey))
}

// Archive a price for new subscriptions. Existing subscriptions that selected it continue using its current catalog value.
func (r *PlansResource) DeletePrice(ctx context.Context, id string, priceID string) (*DeletedObject, error) {
	return parseDirectResponse[DeletedObject](r.http.delete(ctx, fmt.Sprintf("/plans/%s/prices/%s", id, priceID), nil, ""))
}

// Add a base price or a selectable market price variant. Variants inherit their base price outside the markets they override. Configure introductory and promotional benefits through Offers.
func (r *PlansResource) AddPrice(ctx context.Context, id string, params *AddPlanPriceParams) (*PlanPrice, error) {
	body := buildBody(map[string]any{
		"billing_interval":       params.BillingInterval,
		"metadata":               params.Metadata,
		"price":                  params.Price,
		"trial_days":             params.TrialDays,
		"is_default":             params.IsDefault,
		"included_balance":       params.IncludedBalance,
		"included_credits":       params.IncludedCredits,
		"market_prices":          params.MarketPrices,
		"inherits_from_price_id": params.InheritsFromPriceID,
	})
	return parseDirectResponse[PlanPrice](r.http.post(ctx, fmt.Sprintf("/plans/%s/prices", id), body, params.IdempotencyKey))
}

// Configure regional prices and feature overage values for one currency. Currency-specific offer terms are managed through Offers.
func (r *PlansResource) SetRegionalPricing(ctx context.Context, id string, params *SetPlanRegionalPricingParams) (*PlanRegionalPricingResult, error) {
	body := buildBody(map[string]any{
		"currency":      params.Currency,
		"exchange_rate": params.ExchangeRate,
		"prices":        params.Prices,
		"features":      params.Features,
	})
	return parseDirectResponse[PlanRegionalPricingResult](r.http.put(ctx, fmt.Sprintf("/plans/%s/regional", id), body, params.IdempotencyKey))
}

// Get a plan with public price IDs and their automatic introductory offer IDs.
func (r *PlansResource) Get(ctx context.Context, id string) (*Plan, error) {
	return parseDirectResponse[Plan](r.http.get(ctx, fmt.Sprintf("/plans/%s", id), nil))
}

// Update a plan's name, description, visibility, or metadata.
func (r *PlansResource) Update(ctx context.Context, id string, params *UpdatePlanParams) (*Plan, error) {
	body := buildBody(map[string]any{
		"name":        params.Name,
		"description": params.Description,
		"metadata":    params.Metadata,
		"is_public":   params.IsPublic,
	})
	return parseDirectResponse[Plan](r.http.patch(ctx, fmt.Sprintf("/plans/%s", id), body, params.IdempotencyKey))
}

// Soft-delete a plan.
func (r *PlansResource) Delete(ctx context.Context, id string) (*DeletedObject, error) {
	return parseDirectResponse[DeletedObject](r.http.delete(ctx, fmt.Sprintf("/plans/%s", id), nil, ""))
}

// Set a plan's public visibility and return the updated plan.
func (r *PlansResource) SetVisibility(ctx context.Context, id string, params *SetPlanVisibilityParams) (*Plan, error) {
	body := buildBody(map[string]any{
		"is_public": params.IsPublic,
	})
	return parseDirectResponse[Plan](r.http.put(ctx, fmt.Sprintf("/plans/%s/visibility", id), body, params.IdempotencyKey))
}

// List plans with public price IDs and their automatic introductory offer IDs.
func (r *PlansResource) List(ctx context.Context, params *ListPlansParams) (*PlansListResult, error) {
	query := map[string]string{}
	if params.IncludePrivate != nil {
		query["include_private"] = fmt.Sprintf("%t", *params.IncludePrivate)
	}
	return parseDirectResponse[PlansListResult](r.http.get(ctx, "/plans", query))
}

// Create a new plan with optional consumption model, visibility, and plan group assignment.
func (r *PlansResource) Create(ctx context.Context, params *CreatePlanParams) (*Plan, error) {
	body := buildBody(map[string]any{
		"name":                params.Name,
		"code":                params.Code,
		"description":         params.Description,
		"consumption_model":   params.ConsumptionModel,
		"is_public":           params.IsPublic,
		"is_free":             params.IsFree,
		"block_on_exhaustion": params.BlockOnExhaustion,
		"plan_group_id":       params.PlanGroupID,
		"metadata":            params.Metadata,
	})
	return parseDirectResponse[Plan](r.http.post(ctx, "/plans", body, params.IdempotencyKey))
}
