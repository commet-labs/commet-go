package commet

import (
	"context"
	"fmt"
)

type UpdateOfferParams struct {
	Name           string                        `json:"name"`
	Phases         []UpdateOfferParamsPhasesItem `json:"phases"`
	Metadata       map[string]any                `json:"metadata,omitempty"`
	StartsAt       *string                       `json:"starts_at,omitempty"`
	EndsAt         *string                       `json:"ends_at,omitempty"`
	Active         *bool                         `json:"active,omitempty"`
	IdempotencyKey string                        `json:"-"`
}

type ListOffersParams struct {
	Cursor *string `json:"cursor,omitempty"`
	Limit  *int    `json:"limit,omitempty"`
	Active *bool   `json:"active,omitempty"`
}

type CreateOfferParams struct {
	Name           string                        `json:"name"`
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

// Retrieve reusable offer terms by public ID.
func (r *OffersResource) Get(ctx context.Context, id string) (*Offer, error) {
	return parseDirectResponse[Offer](r.http.get(ctx, fmt.Sprintf("/offers/%s", id), nil))
}

// Replace reusable offer terms. Existing applications keep their immutable accepted terms.
func (r *OffersResource) Update(ctx context.Context, id string, params *UpdateOfferParams) (*Offer, error) {
	body := buildBody(map[string]any{
		"name":      params.Name,
		"phases":    params.Phases,
		"metadata":  params.Metadata,
		"starts_at": params.StartsAt,
		"ends_at":   params.EndsAt,
		"active":    params.Active,
	})
	return parseDirectResponse[Offer](r.http.patch(ctx, fmt.Sprintf("/offers/%s", id), body, params.IdempotencyKey))
}

// Soft-delete an Offer. Existing applications and their accepted terms remain available for billing and audit.
func (r *OffersResource) Delete(ctx context.Context, id string) (*DeletedOffer, error) {
	return parseDirectResponse[DeletedOffer](r.http.delete(ctx, fmt.Sprintf("/offers/%s", id), nil, ""))
}

// List reusable offer terms. Offers are independent from plans, prices, eligibility, and distribution channels.
func (r *OffersResource) List(ctx context.Context, params *ListOffersParams) (*OffersListResult, error) {
	query := map[string]string{}
	if params.Cursor != nil {
		query["cursor"] = *params.Cursor
	}
	if params.Limit != nil {
		query["limit"] = fmt.Sprintf("%d", *params.Limit)
	}
	if params.Active != nil {
		query["active"] = fmt.Sprintf("%t", *params.Active)
	}
	return parseDirectResponse[OffersListResult](r.http.get(ctx, "/offers", query))
}

// Create reusable offer terms without assigning a plan, price, eligibility rule, or distribution channel.
func (r *OffersResource) Create(ctx context.Context, params *CreateOfferParams) (*Offer, error) {
	body := buildBody(map[string]any{
		"name":      params.Name,
		"phases":    params.Phases,
		"metadata":  params.Metadata,
		"starts_at": params.StartsAt,
		"ends_at":   params.EndsAt,
		"active":    params.Active,
	})
	return parseDirectResponse[Offer](r.http.post(ctx, "/offers", body, params.IdempotencyKey))
}
