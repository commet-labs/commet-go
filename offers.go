package commet

import (
	"context"
	"fmt"
)

type UpdateOfferParams struct {
	Name           string                        `json:"name"`
	Purpose        string                        `json:"purpose"`
	PlanPriceIds   []string                      `json:"plan_price_ids"`
	Phases         []UpdateOfferParamsPhasesItem `json:"phases"`
	Metadata       map[string]any                `json:"metadata,omitempty"`
	StartsAt       *string                       `json:"starts_at,omitempty"`
	EndsAt         *string                       `json:"ends_at,omitempty"`
	Active         *bool                         `json:"active,omitempty"`
	IdempotencyKey string                        `json:"-"`
}

type ListOffersParams struct {
	Cursor      *string `json:"cursor,omitempty"`
	Limit       *int    `json:"limit,omitempty"`
	PlanPriceID *string `json:"plan_price_id,omitempty"`
	Purpose     *string `json:"purpose,omitempty"`
	Active      *bool   `json:"active,omitempty"`
}

type CreateOfferParams struct {
	Name           string                        `json:"name"`
	Purpose        string                        `json:"purpose"`
	PlanPriceIds   []string                      `json:"plan_price_ids"`
	Phases         []CreateOfferParamsPhasesItem `json:"phases"`
	Metadata       map[string]any                `json:"metadata,omitempty"`
	StartsAt       *string                       `json:"starts_at,omitempty"`
	EndsAt         *string                       `json:"ends_at,omitempty"`
	Active         *bool                         `json:"active,omitempty"`
	IdempotencyKey string                        `json:"-"`
}

type OffersResource struct {
	http *httpClient
}

// Retrieve a canonical offer by its public ID.
func (r *OffersResource) Get(ctx context.Context, id string) (*Offer, error) {
	return parseDirectResponse[Offer](r.http.get(ctx, fmt.Sprintf("/offers/%s", id), nil))
}

// Replace an offer's catalog definition. Existing offer applications keep their immutable accepted terms.
func (r *OffersResource) Update(ctx context.Context, id string, params *UpdateOfferParams) (*Offer, error) {
	body := buildBody(map[string]any{
		"name":           params.Name,
		"purpose":        params.Purpose,
		"plan_price_ids": params.PlanPriceIds,
		"phases":         params.Phases,
		"metadata":       params.Metadata,
		"starts_at":      params.StartsAt,
		"ends_at":        params.EndsAt,
		"active":         params.Active,
	})
	return parseDirectResponse[Offer](r.http.patch(ctx, fmt.Sprintf("/offers/%s", id), body, params.IdempotencyKey))
}

// Soft-delete an offer. Existing applications and their accepted terms remain available for billing and audit.
func (r *OffersResource) Delete(ctx context.Context, id string) (*DeletedOffer, error) {
	return parseDirectResponse[DeletedOffer](r.http.delete(ctx, fmt.Sprintf("/offers/%s", id), nil, ""))
}

// List the organization's canonical introductory and promotional offers.
func (r *OffersResource) List(ctx context.Context, params *ListOffersParams) (*OffersListResult, error) {
	query := map[string]string{}
	if params.Cursor != nil {
		query["cursor"] = *params.Cursor
	}
	if params.Limit != nil {
		query["limit"] = fmt.Sprintf("%d", *params.Limit)
	}
	if params.PlanPriceID != nil {
		query["plan_price_id"] = *params.PlanPriceID
	}
	if params.Purpose != nil {
		query["purpose"] = *params.Purpose
	}
	if params.Active != nil {
		query["active"] = fmt.Sprintf("%t", *params.Active)
	}
	return parseDirectResponse[OffersListResult](r.http.get(ctx, "/offers", query))
}

// Create a canonical offer scoped to one or more plan prices. Currency-specific phases require an explicit USD value and never fall back across currencies.
func (r *OffersResource) Create(ctx context.Context, params *CreateOfferParams) (*Offer, error) {
	body := buildBody(map[string]any{
		"name":           params.Name,
		"purpose":        params.Purpose,
		"plan_price_ids": params.PlanPriceIds,
		"phases":         params.Phases,
		"metadata":       params.Metadata,
		"starts_at":      params.StartsAt,
		"ends_at":        params.EndsAt,
		"active":         params.Active,
	})
	return parseDirectResponse[Offer](r.http.post(ctx, "/offers", body, params.IdempotencyKey))
}
