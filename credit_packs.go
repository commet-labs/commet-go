package commet

import (
	"context"
	"fmt"
)

type UpdateCreditPackParams struct {
	Name           *string `json:"name,omitempty"`
	Description    *string `json:"description,omitempty"`
	Credits        *int    `json:"credits,omitempty"`
	Price          *int    `json:"price,omitempty"`
	IsActive       *bool   `json:"is_active,omitempty"`
	IdempotencyKey string  `json:"-"`
}

type CreateCreditPackParams struct {
	Name           string  `json:"name"`
	Description    *string `json:"description,omitempty"`
	Credits        int     `json:"credits"`
	Price          int     `json:"price"`
	IsActive       *bool   `json:"is_active,omitempty"`
	IdempotencyKey string  `json:"-"`
}

type CreditPacksResource struct {
	http *httpClient
}

// Update a credit pack's name, description, credits, price, or active status.
func (r *CreditPacksResource) Update(ctx context.Context, id string, params *UpdateCreditPackParams) (*CreditPack, error) {
	body := buildBody(map[string]any{
		"name":        params.Name,
		"description": params.Description,
		"credits":     params.Credits,
		"price":       params.Price,
		"is_active":   params.IsActive,
	})
	return parseDirectResponse[CreditPack](r.http.patch(ctx, fmt.Sprintf("/credit-packs/%s", id), body, params.IdempotencyKey))
}

// Soft-delete a credit pack.
func (r *CreditPacksResource) Delete(ctx context.Context, id string) (*DeletedObject, error) {
	return parseDirectResponse[DeletedObject](r.http.delete(ctx, fmt.Sprintf("/credit-packs/%s", id), nil, ""))
}

// List all active credit packs.
func (r *CreditPacksResource) List(ctx context.Context) (*CreditPacksListResult, error) {
	return parseDirectResponse[CreditPacksListResult](r.http.get(ctx, "/credit-packs", nil))
}

// Create a new credit pack.
func (r *CreditPacksResource) Create(ctx context.Context, params *CreateCreditPackParams) (*CreditPack, error) {
	body := buildBody(map[string]any{
		"name":        params.Name,
		"description": params.Description,
		"credits":     params.Credits,
		"price":       params.Price,
		"is_active":   params.IsActive,
	})
	return parseDirectResponse[CreditPack](r.http.post(ctx, "/credit-packs", body, params.IdempotencyKey))
}
