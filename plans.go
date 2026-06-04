package commet

import (
	"context"
	"fmt"
)

type ListPlansParams struct {
	IncludePrivate *bool  `json:"include_private,omitempty"`
	Limit          *int   `json:"limit,omitempty"`
	Cursor         string `json:"cursor,omitempty"`
}

type CreatePlanParams struct {
	Name              string         `json:"name"`
	Code              string         `json:"code"`
	Description       string         `json:"description,omitempty"`
	ConsumptionModel  string         `json:"consumption_model,omitempty"`
	IsPublic          *bool          `json:"is_public,omitempty"`
	IsFree            *bool          `json:"is_free,omitempty"`
	BlockOnExhaustion *bool          `json:"block_on_exhaustion,omitempty"`
	PlanGroupID       string         `json:"plan_group_id,omitempty"`
	Metadata          map[string]any `json:"metadata,omitempty"`
}

type UpdatePlanParams struct {
	Name        string         `json:"name,omitempty"`
	Description string         `json:"description,omitempty"`
	Metadata    map[string]any `json:"metadata,omitempty"`
	IsPublic    *bool          `json:"is_public,omitempty"`
}

type SetVisibilityParams struct {
	IsPublic bool `json:"is_public"`
}

type AddPlanFeatureParams struct {
	FeatureID        string   `json:"feature_id"`
	Enabled          *bool    `json:"enabled,omitempty"`
	IncludedAmount   *int     `json:"included_amount,omitempty"`
	Unlimited        *bool    `json:"unlimited,omitempty"`
	OverageEnabled   *bool    `json:"overage_enabled,omitempty"`
	CreditsPerUnit   *int     `json:"credits_per_unit,omitempty"`
	PricingMode      string   `json:"pricing_mode,omitempty"`
	OverageUnitPrice *float64 `json:"overage_unit_price,omitempty"`
	Margin           *float64 `json:"margin,omitempty"`
}

type UpdatePlanFeatureParams struct {
	Enabled          *bool    `json:"enabled,omitempty"`
	IncludedAmount   *int     `json:"included_amount,omitempty"`
	Unlimited        *bool    `json:"unlimited,omitempty"`
	OverageEnabled   *bool    `json:"overage_enabled,omitempty"`
	CreditsPerUnit   *int     `json:"credits_per_unit,omitempty"`
	PricingMode      string   `json:"pricing_mode,omitempty"`
	OverageUnitPrice *float64 `json:"overage_unit_price,omitempty"`
	Margin           *float64 `json:"margin,omitempty"`
}

type AddPlanPriceParams struct {
	BillingInterval          BillingInterval `json:"billing_interval"`
	Price                    int             `json:"price"`
	TrialDays                *int            `json:"trial_days,omitempty"`
	IsDefault                *bool           `json:"is_default,omitempty"`
	IncludedBalance          *int            `json:"included_balance,omitempty"`
	IncludedCredits          *int            `json:"included_credits,omitempty"`
	IntroOfferEnabled        *bool           `json:"intro_offer_enabled,omitempty"`
	IntroOfferDiscountType   string          `json:"intro_offer_discount_type,omitempty"`
	IntroOfferDiscountValue  *float64        `json:"intro_offer_discount_value,omitempty"`
	IntroOfferDurationCycles *int            `json:"intro_offer_duration_cycles,omitempty"`
}

type UpdatePlanPriceParams struct {
	Price                    *int     `json:"price,omitempty"`
	IsDefault                *bool    `json:"is_default,omitempty"`
	TrialDays                *int     `json:"trial_days,omitempty"`
	IncludedBalance          *int     `json:"included_balance,omitempty"`
	IncludedCredits          *int     `json:"included_credits,omitempty"`
	IntroOfferEnabled        *bool    `json:"intro_offer_enabled,omitempty"`
	IntroOfferDiscountType   *string  `json:"intro_offer_discount_type,omitempty"`
	IntroOfferDiscountValue  *float64 `json:"intro_offer_discount_value,omitempty"`
	IntroOfferDurationCycles *int     `json:"intro_offer_duration_cycles,omitempty"`
}

type SetRegionalPricesParams struct {
	Overrides []RegionalPriceOverride `json:"overrides"`
}

type PlansResource struct {
	http *httpClient
}

func (r *PlansResource) List(ctx context.Context, params *ListPlansParams) (*ApiResponse[[]Plan], error) {
	queryParams := map[string]string{}
	if params != nil {
		if params.IncludePrivate != nil {
			queryParams["include_private"] = fmt.Sprintf("%t", *params.IncludePrivate)
		}
		if params.Limit != nil {
			queryParams["limit"] = fmt.Sprintf("%d", *params.Limit)
		}
		if params.Cursor != "" {
			queryParams["cursor"] = params.Cursor
		}
	}
	return parseResponse[[]Plan](r.http.get(ctx, "/plans", queryParams))
}

func (r *PlansResource) Get(ctx context.Context, planID string) (*ApiResponse[PlanDetail], error) {
	return parseResponse[PlanDetail](r.http.get(ctx, fmt.Sprintf("/plans/%s", planID), nil))
}

func (r *PlansResource) Create(ctx context.Context, params *CreatePlanParams) (*ApiResponse[PlanManage], error) {
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
	return parseResponse[PlanManage](r.http.post(ctx, "/plans/manage", body, ""))
}

func (r *PlansResource) Update(ctx context.Context, planID string, params *UpdatePlanParams) (*ApiResponse[PlanManage], error) {
	body := buildBody(map[string]any{
		"name":        params.Name,
		"description": params.Description,
		"metadata":    params.Metadata,
		"is_public":   params.IsPublic,
	})
	return parseResponse[PlanManage](r.http.put(ctx, fmt.Sprintf("/plans/%s/manage", planID), body, ""))
}

func (r *PlansResource) Delete(ctx context.Context, planID string) (*ApiResponse[DeleteResult], error) {
	return parseResponse[DeleteResult](r.http.delete(ctx, fmt.Sprintf("/plans/%s/manage", planID), nil, ""))
}

func (r *PlansResource) SetVisibility(ctx context.Context, planID string, params *SetVisibilityParams) (*ApiResponse[PlanManage], error) {
	body := buildBody(map[string]any{
		"is_public": params.IsPublic,
	})
	return parseResponse[PlanManage](r.http.put(ctx, fmt.Sprintf("/plans/%s/visibility", planID), body, ""))
}

func (r *PlansResource) AddFeature(ctx context.Context, planID string, params *AddPlanFeatureParams) (*ApiResponse[PlanFeatureManage], error) {
	body := buildBody(map[string]any{
		"feature_id":         params.FeatureID,
		"enabled":            params.Enabled,
		"included_amount":    params.IncludedAmount,
		"unlimited":          params.Unlimited,
		"overage_enabled":    params.OverageEnabled,
		"credits_per_unit":   params.CreditsPerUnit,
		"pricing_mode":       params.PricingMode,
		"overage_unit_price": params.OverageUnitPrice,
		"margin":             params.Margin,
	})
	return parseResponse[PlanFeatureManage](r.http.post(ctx, fmt.Sprintf("/plans/%s/features", planID), body, ""))
}

func (r *PlansResource) UpdateFeature(ctx context.Context, planID string, featureID string, params *UpdatePlanFeatureParams) (*ApiResponse[PlanFeatureManage], error) {
	body := buildBody(map[string]any{
		"enabled":            params.Enabled,
		"included_amount":    params.IncludedAmount,
		"unlimited":          params.Unlimited,
		"overage_enabled":    params.OverageEnabled,
		"credits_per_unit":   params.CreditsPerUnit,
		"pricing_mode":       params.PricingMode,
		"overage_unit_price": params.OverageUnitPrice,
		"margin":             params.Margin,
	})
	return parseResponse[PlanFeatureManage](r.http.put(ctx, fmt.Sprintf("/plans/%s/features/%s", planID, featureID), body, ""))
}

func (r *PlansResource) RemoveFeature(ctx context.Context, planID string, featureID string) (*ApiResponse[RemoveResult], error) {
	return parseResponse[RemoveResult](r.http.delete(ctx, fmt.Sprintf("/plans/%s/features/%s", planID, featureID), nil, ""))
}

func (r *PlansResource) AddPrice(ctx context.Context, planID string, params *AddPlanPriceParams) (*ApiResponse[PlanPriceManage], error) {
	body := buildBody(map[string]any{
		"billing_interval":            params.BillingInterval,
		"price":                       params.Price,
		"trial_days":                  params.TrialDays,
		"is_default":                  params.IsDefault,
		"included_balance":            params.IncludedBalance,
		"included_credits":            params.IncludedCredits,
		"intro_offer_enabled":         params.IntroOfferEnabled,
		"intro_offer_discount_type":   params.IntroOfferDiscountType,
		"intro_offer_discount_value":  params.IntroOfferDiscountValue,
		"intro_offer_duration_cycles": params.IntroOfferDurationCycles,
	})
	return parseResponse[PlanPriceManage](r.http.post(ctx, fmt.Sprintf("/plans/%s/prices", planID), body, ""))
}

func (r *PlansResource) UpdatePrice(ctx context.Context, planID string, priceID string, params *UpdatePlanPriceParams) (*ApiResponse[PlanPriceManage], error) {
	body := buildBody(map[string]any{
		"price":                       params.Price,
		"is_default":                  params.IsDefault,
		"trial_days":                  params.TrialDays,
		"included_balance":            params.IncludedBalance,
		"included_credits":            params.IncludedCredits,
		"intro_offer_enabled":         params.IntroOfferEnabled,
		"intro_offer_discount_type":   params.IntroOfferDiscountType,
		"intro_offer_discount_value":  params.IntroOfferDiscountValue,
		"intro_offer_duration_cycles": params.IntroOfferDurationCycles,
	})
	return parseResponse[PlanPriceManage](r.http.put(ctx, fmt.Sprintf("/plans/%s/prices/%s", planID, priceID), body, ""))
}

func (r *PlansResource) DeletePrice(ctx context.Context, planID string, priceID string) (*ApiResponse[DeleteResult], error) {
	return parseResponse[DeleteResult](r.http.delete(ctx, fmt.Sprintf("/plans/%s/prices/%s", planID, priceID), nil, ""))
}

func (r *PlansResource) SetDefaultPrice(ctx context.Context, planID string, priceID string) (*ApiResponse[PlanPriceManage], error) {
	return parseResponse[PlanPriceManage](r.http.put(ctx, fmt.Sprintf("/plans/%s/prices/%s/default", planID, priceID), map[string]any{}, ""))
}

func (r *PlansResource) SetRegionalPrices(ctx context.Context, planID string, priceID string, params *SetRegionalPricesParams) (*ApiResponse[RegionalPriceResult], error) {
	body := buildBody(map[string]any{
		"overrides": params.Overrides,
	})
	return parseResponse[RegionalPriceResult](r.http.put(ctx, fmt.Sprintf("/plans/%s/prices/%s/regional", planID, priceID), body, ""))
}

func (r *PlansResource) DeleteRegionalPrices(ctx context.Context, planID string, priceID string) (*ApiResponse[DeleteResult], error) {
	return parseResponse[DeleteResult](r.http.delete(ctx, fmt.Sprintf("/plans/%s/prices/%s/regional", planID, priceID), nil, ""))
}
