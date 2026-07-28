package commet

import (
	"context"
	"fmt"
)

type UpdatePromoCodeParams struct {
	BillingInterval *string  `json:"billing_interval,omitempty"`
	MaxRedemptions  *int     `json:"max_redemptions,omitempty"`
	ExpiresAt       *string  `json:"expires_at,omitempty"`
	Active          *bool    `json:"active,omitempty"`
	PlanIds         []string `json:"plan_ids,omitempty"`
	IdempotencyKey  string   `json:"-"`
}

type ListPromoCodesParams struct {
	Cursor *string `json:"cursor,omitempty"`
	Limit  *int    `json:"limit,omitempty"`
}

type CreatePromoCodeParams struct {
	Code            string   `json:"code"`
	OfferID         string   `json:"offer_id"`
	BillingInterval *string  `json:"billing_interval,omitempty"`
	MaxRedemptions  *int     `json:"max_redemptions,omitempty"`
	ExpiresAt       *string  `json:"expires_at,omitempty"`
	PlanIds         []string `json:"plan_ids,omitempty"`
	IdempotencyKey  string   `json:"-"`
}

type PromoCodesResource struct {
	http *httpClient
}

// Retrieve a promo code by its public ID.
func (r *PromoCodesResource) Get(ctx context.Context, id string) (*PromoCode, error) {
	return parseDirectResponse[PromoCode](r.http.get(ctx, fmt.Sprintf("/promo-codes/%s", id), nil))
}

// Update a promo code's billing interval, redemption limits, expiration, active status, or plan restrictions.
func (r *PromoCodesResource) Update(ctx context.Context, id string, params *UpdatePromoCodeParams) (*PromoCode, error) {
	body := buildBody(map[string]any{
		"billing_interval": params.BillingInterval,
		"max_redemptions":  params.MaxRedemptions,
		"expires_at":       params.ExpiresAt,
		"active":           params.Active,
		"plan_ids":         params.PlanIds,
	})
	return parseDirectResponse[PromoCode](r.http.patch(ctx, fmt.Sprintf("/promo-codes/%s", id), body, params.IdempotencyKey))
}

// List promo codes with cursor-based pagination.
func (r *PromoCodesResource) List(ctx context.Context, params *ListPromoCodesParams) (*PromoCodesListResult, error) {
	query := map[string]string{}
	if params.Cursor != nil {
		query["cursor"] = *params.Cursor
	}
	if params.Limit != nil {
		query["limit"] = fmt.Sprintf("%d", *params.Limit)
	}
	return parseDirectResponse[PromoCodesListResult](r.http.get(ctx, "/promo-codes", query))
}

// Create a distribution code for an existing promotional offer. Offer economics remain owned by the referenced Offer.
func (r *PromoCodesResource) Create(ctx context.Context, params *CreatePromoCodeParams) (*PromoCode, error) {
	body := buildBody(map[string]any{
		"code":             params.Code,
		"offer_id":         params.OfferID,
		"billing_interval": params.BillingInterval,
		"max_redemptions":  params.MaxRedemptions,
		"expires_at":       params.ExpiresAt,
		"plan_ids":         params.PlanIds,
	})
	return parseDirectResponse[PromoCode](r.http.post(ctx, "/promo-codes", body, params.IdempotencyKey))
}
