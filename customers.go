package commet

import (
	"context"
	"fmt"
)

type ListCustomersParams struct {
	ExternalID *string `json:"external_id,omitempty"`
	Limit      *int    `json:"limit,omitempty"`
	Cursor     *string `json:"cursor,omitempty"`
}

type CreateCustomerParams struct {
	ID             *string                      `json:"id,omitempty"`
	ExternalID     *string                      `json:"external_id,omitempty"`
	FullName       *string                      `json:"full_name,omitempty"`
	Address        *CreateCustomerParamsAddress `json:"address,omitempty"`
	AddressID      *string                      `json:"address_id,omitempty"`
	Email          string                       `json:"email"`
	Timezone       *Timezone                    `json:"timezone,omitempty"`
	Metadata       map[string]any               `json:"metadata,omitempty"`
	IdempotencyKey string                       `json:"-"`
}

type UpdateCustomerParams struct {
	Email          *string                      `json:"email,omitempty"`
	FullName       *string                      `json:"full_name,omitempty"`
	ExternalID     *string                      `json:"external_id,omitempty"`
	Timezone       *Timezone                    `json:"timezone,omitempty"`
	Metadata       map[string]any               `json:"metadata,omitempty"`
	Address        *UpdateCustomerParamsAddress `json:"address,omitempty"`
	IdempotencyKey string                       `json:"-"`
}

type BatchCreateCustomersParams struct {
	Customers      []BatchCreateCustomersParamsCustomersItem `json:"customers"`
	IdempotencyKey string                                    `json:"-"`
}

type CustomersResource struct {
	http *httpClient
}

// List customers with cursor-based pagination.
func (r *CustomersResource) List(ctx context.Context, params *ListCustomersParams) (*ApiResponse[[]Customer], error) {
	query := map[string]string{}
	if params.ExternalID != nil {
		query["external_id"] = *params.ExternalID
	}
	if params.Limit != nil {
		query["limit"] = fmt.Sprintf("%d", *params.Limit)
	}
	if params.Cursor != nil {
		query["cursor"] = *params.Cursor
	}
	return parseResponse[[]Customer](r.http.get(ctx, "/customers", query))
}

// Create a new customer. Idempotent when customerId is provided.
func (r *CustomersResource) Create(ctx context.Context, params *CreateCustomerParams) (*ApiResponse[Customer], error) {
	body := buildBody(map[string]any{
		"id":          params.ID,
		"external_id": params.ExternalID,
		"full_name":   params.FullName,
		"address":     params.Address,
		"address_id":  params.AddressID,
		"email":       params.Email,
		"timezone":    params.Timezone,
		"metadata":    params.Metadata,
	})
	return parseResponse[Customer](r.http.post(ctx, "/customers", body, params.IdempotencyKey))
}

// Retrieve a customer by their public ID, including subscription status and metadata.
func (r *CustomersResource) Get(ctx context.Context, id string) (*ApiResponse[Customer], error) {
	return parseResponse[Customer](r.http.get(ctx, fmt.Sprintf("/customers/%s", id), nil))
}

// Update a customer's name, external ID, or metadata.
func (r *CustomersResource) Update(ctx context.Context, id string, params *UpdateCustomerParams) (*ApiResponse[Customer], error) {
	body := buildBody(map[string]any{
		"email":       params.Email,
		"full_name":   params.FullName,
		"external_id": params.ExternalID,
		"timezone":    params.Timezone,
		"metadata":    params.Metadata,
		"address":     params.Address,
	})
	return parseResponse[Customer](r.http.put(ctx, fmt.Sprintf("/customers/%s", id), body, params.IdempotencyKey))
}

// Create up to 100 customers in a single request.
func (r *CustomersResource) CreateBatch(ctx context.Context, params *BatchCreateCustomersParams) (*ApiResponse[CustomerBatch], error) {
	body := buildBody(map[string]any{
		"customers": params.Customers,
	})
	return parseResponse[CustomerBatch](r.http.post(ctx, "/customers/batch", body, params.IdempotencyKey))
}
