package commet

import (
	"context"
	"fmt"
)

type ListPromoCodesParams struct {
	Limit  *int    `json:"limit,omitempty"`
	Cursor *string `json:"cursor,omitempty"`
}

type CreatePromoCodeParams struct {
	Code            string           `json:"code"`
	DiscountType    DiscountType     `json:"discount_type"`
	DiscountValue   int              `json:"discount_value"`
	DurationCycles  *int             `json:"duration_cycles,omitempty"`
	BillingInterval *BillingInterval `json:"billing_interval,omitempty"`
	MaxRedemptions  *int             `json:"max_redemptions,omitempty"`
	ExpiresAt       *string          `json:"expires_at,omitempty"`
	PlanIds         []string         `json:"plan_ids,omitempty"`
	IdempotencyKey  string           `json:"-"`
}

type UpdatePromoCodeParams struct {
	BillingInterval *BillingInterval `json:"billing_interval,omitempty"`
	MaxRedemptions  *int             `json:"max_redemptions,omitempty"`
	ExpiresAt       *string          `json:"expires_at,omitempty"`
	Active          *bool            `json:"active,omitempty"`
	PlanIds         []string         `json:"plan_ids,omitempty"`
	IdempotencyKey  string           `json:"-"`
}

type PromoCodesResource struct {
	http *httpClient
}

// List promo codes with cursor-based pagination.
func (r *PromoCodesResource) List(ctx context.Context, params *ListPromoCodesParams) (*ApiResponse[[]PromoCode], error) {
	query := map[string]string{}
	if params.Limit != nil {
		query["limit"] = fmt.Sprintf("%d", *params.Limit)
	}
	if params.Cursor != nil {
		query["cursor"] = *params.Cursor
	}
	return parseResponse[[]PromoCode](r.http.get(ctx, "/promo-codes", query))
}

// Retrieve a promo code by its public ID.
func (r *PromoCodesResource) Get(ctx context.Context, id string) (*ApiResponse[PromoCode], error) {
	return parseResponse[PromoCode](r.http.get(ctx, fmt.Sprintf("/promo-codes/%s", id), nil))
}

// Create a new promo code. Optionally restrict it to specific plans and a billing interval.
//
// **100% discounts are not supported.** Percentage codes must be strictly less than 100% (`discountValue` < 10000 basis points). For full waivers, use an introductory offer on the plan instead. At checkout, any code — percentage or fixed amount — that would reduce the total below the currency's minimum charge ($0.50 USD equivalent) is silently dropped.
func (r *PromoCodesResource) Create(ctx context.Context, params *CreatePromoCodeParams) (*ApiResponse[PromoCode], error) {
	body := buildBody(map[string]any{
		"code":             params.Code,
		"discount_type":    params.DiscountType,
		"discount_value":   params.DiscountValue,
		"duration_cycles":  params.DurationCycles,
		"billing_interval": params.BillingInterval,
		"max_redemptions":  params.MaxRedemptions,
		"expires_at":       params.ExpiresAt,
		"plan_ids":         params.PlanIds,
	})
	return parseResponse[PromoCode](r.http.post(ctx, "/promo-codes", body, params.IdempotencyKey))
}

// Update a promo code's billing interval, redemption limits, expiration, active status, or plan restrictions.
func (r *PromoCodesResource) Update(ctx context.Context, id string, params *UpdatePromoCodeParams) (*ApiResponse[PromoCode], error) {
	body := buildBody(map[string]any{
		"billing_interval": params.BillingInterval,
		"max_redemptions":  params.MaxRedemptions,
		"expires_at":       params.ExpiresAt,
		"active":           params.Active,
		"plan_ids":         params.PlanIds,
	})
	return parseResponse[PromoCode](r.http.put(ctx, fmt.Sprintf("/promo-codes/%s", id), body, params.IdempotencyKey))
}
