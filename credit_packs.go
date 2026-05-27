package commet

import (
	"context"
	"fmt"
)

type CreateCreditPackParams struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Credits     int    `json:"credits"`
	Price       int    `json:"price"`
	IsActive    *bool  `json:"is_active,omitempty"`
}

type UpdateCreditPackParams struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Credits     *int   `json:"credits,omitempty"`
	Price       *int   `json:"price,omitempty"`
	IsActive    *bool  `json:"is_active,omitempty"`
}

type CreditPacksResource struct {
	http *httpClient
}

func (r *CreditPacksResource) List(ctx context.Context) (*ApiResponse[[]CreditPack], error) {
	return parseResponse[[]CreditPack](r.http.get(ctx, "/credit-packs", nil))
}

func (r *CreditPacksResource) Create(ctx context.Context, params *CreateCreditPackParams) (*ApiResponse[CreditPackDetail], error) {
	body := buildBody(map[string]any{
		"name":        params.Name,
		"description": params.Description,
		"credits":     params.Credits,
		"price":       params.Price,
		"is_active":   params.IsActive,
	})
	return parseResponse[CreditPackDetail](r.http.post(ctx, "/credit-packs/manage", body, ""))
}

func (r *CreditPacksResource) Update(ctx context.Context, creditPackID string, params *UpdateCreditPackParams) (*ApiResponse[CreditPackDetail], error) {
	body := buildBody(map[string]any{
		"name":        params.Name,
		"description": params.Description,
		"credits":     params.Credits,
		"price":       params.Price,
		"is_active":   params.IsActive,
	})
	return parseResponse[CreditPackDetail](r.http.put(ctx, fmt.Sprintf("/credit-packs/%s", creditPackID), body, ""))
}

func (r *CreditPacksResource) Delete(ctx context.Context, creditPackID string) (*ApiResponse[DeleteResult], error) {
	return parseResponse[DeleteResult](r.http.delete(ctx, fmt.Sprintf("/credit-packs/%s", creditPackID), nil, ""))
}
