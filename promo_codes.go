package commet

import (
	"context"
	"fmt"
)

type ListPromoCodesParams struct {
	Limit  *int   `json:"limit,omitempty"`
	Cursor string `json:"cursor,omitempty"`
}

type CreatePromoCodeParams struct {
	Code           string   `json:"code"`
	DiscountType   string   `json:"discount_type"`
	DiscountValue  float64  `json:"discount_value"`
	DurationCycles *int     `json:"duration_cycles,omitempty"`
	MaxRedemptions *int     `json:"max_redemptions,omitempty"`
	ExpiresAt      string   `json:"expires_at,omitempty"`
	PlanIDs        []string `json:"plan_ids,omitempty"`
}

type UpdatePromoCodeParams struct {
	MaxRedemptions *int     `json:"max_redemptions,omitempty"`
	ExpiresAt      string   `json:"expires_at,omitempty"`
	Active         *bool    `json:"active,omitempty"`
	PlanIDs        []string `json:"plan_ids,omitempty"`
}

type PromoCodesResource struct {
	http *httpClient
}

func (r *PromoCodesResource) List(ctx context.Context, params *ListPromoCodesParams) (*ApiResponse[[]PromoCode], error) {
	queryParams := map[string]string{}
	if params != nil {
		if params.Limit != nil {
			queryParams["limit"] = fmt.Sprintf("%d", *params.Limit)
		}
		if params.Cursor != "" {
			queryParams["cursor"] = params.Cursor
		}
	}
	return parseResponse[[]PromoCode](r.http.get(ctx, "/promo-codes", queryParams))
}

func (r *PromoCodesResource) Get(ctx context.Context, promoCodeID string) (*ApiResponse[PromoCodeDetail], error) {
	return parseResponse[PromoCodeDetail](r.http.get(ctx, fmt.Sprintf("/promo-codes/%s", promoCodeID), nil))
}

func (r *PromoCodesResource) Create(ctx context.Context, params *CreatePromoCodeParams) (*ApiResponse[PromoCode], error) {
	body := buildBody(map[string]any{
		"code":            params.Code,
		"discount_type":   params.DiscountType,
		"discount_value":  params.DiscountValue,
		"duration_cycles": params.DurationCycles,
		"max_redemptions": params.MaxRedemptions,
		"expires_at":      params.ExpiresAt,
		"plan_ids":        params.PlanIDs,
	})
	return parseResponse[PromoCode](r.http.post(ctx, "/promo-codes", body, ""))
}

func (r *PromoCodesResource) Update(ctx context.Context, promoCodeID string, params *UpdatePromoCodeParams) (*ApiResponse[PromoCodeDetail], error) {
	body := buildBody(map[string]any{
		"max_redemptions": params.MaxRedemptions,
		"expires_at":      params.ExpiresAt,
		"active":          params.Active,
		"plan_ids":        params.PlanIDs,
	})
	return parseResponse[PromoCodeDetail](r.http.put(ctx, fmt.Sprintf("/promo-codes/%s", promoCodeID), body, ""))
}
