package commet

import (
	"context"
	"fmt"
)

type CreateCustomerParams struct {
	Email          string            `json:"billing_email"`
	ID             string            `json:"-"`
	FullName       string            `json:"full_name,omitempty"`
	Domain         string            `json:"domain,omitempty"`
	Website        string            `json:"website,omitempty"`
	Timezone       string            `json:"timezone,omitempty"`
	Language       string            `json:"language,omitempty"`
	Industry       string            `json:"industry,omitempty"`
	Metadata       map[string]any    `json:"metadata,omitempty"`
	Address        map[string]string `json:"address,omitempty"`
	IdempotencyKey string            `json:"-"`
}

type UpdateCustomerParams struct {
	Email          string            `json:"billing_email,omitempty"`
	ExternalID     string            `json:"external_id,omitempty"`
	FullName       string            `json:"full_name,omitempty"`
	Domain         string            `json:"domain,omitempty"`
	Website        string            `json:"website,omitempty"`
	Timezone       string            `json:"timezone,omitempty"`
	Language       string            `json:"language,omitempty"`
	Industry       string            `json:"industry,omitempty"`
	Metadata       map[string]any    `json:"metadata,omitempty"`
	Address        map[string]string `json:"address,omitempty"`
	IdempotencyKey string            `json:"-"`
}

type ListCustomersParams struct {
	CustomerID string `json:"customer_id,omitempty"`
	IsActive   *bool  `json:"is_active,omitempty"`
	Search     string `json:"search,omitempty"`
	Limit      *int   `json:"limit,omitempty"`
	Cursor     string `json:"cursor,omitempty"`
}

type CustomersResource struct {
	http *httpClient
}

func (r *CustomersResource) Create(ctx context.Context, params *CreateCustomerParams) (*ApiResponse[Customer], error) {
	body := buildBody(map[string]any{
		"billing_email": params.Email,
		"external_id":   params.ID,
		"full_name":     params.FullName,
		"domain":        params.Domain,
		"website":       params.Website,
		"timezone":      params.Timezone,
		"language":      params.Language,
		"industry":      params.Industry,
		"metadata":      params.Metadata,
		"address":       params.Address,
	})
	return parseResponse[Customer](r.http.post(ctx, "/customers", body, params.IdempotencyKey))
}

func (r *CustomersResource) CreateBatch(ctx context.Context, customers []CreateCustomerParams, idempotencyKey string) (*ApiResponse[BatchResult], error) {
	mapped := make([]any, len(customers))
	for i, c := range customers {
		mapped[i] = buildBody(map[string]any{
			"billing_email": c.Email,
			"external_id":   c.ID,
			"full_name":     c.FullName,
			"domain":        c.Domain,
			"website":       c.Website,
			"timezone":      c.Timezone,
			"language":      c.Language,
			"industry":      c.Industry,
			"metadata":      c.Metadata,
			"address":       c.Address,
		})
	}
	body := map[string]any{"customers": mapped}
	return parseResponse[BatchResult](r.http.post(ctx, "/customers/batch", body, idempotencyKey))
}

func (r *CustomersResource) Get(ctx context.Context, customerID string) (*ApiResponse[Customer], error) {
	return parseResponse[Customer](r.http.get(ctx, fmt.Sprintf("/customers/%s", customerID), nil))
}

func (r *CustomersResource) Update(ctx context.Context, customerID string, params *UpdateCustomerParams) (*ApiResponse[Customer], error) {
	body := buildBody(map[string]any{
		"billing_email": params.Email,
		"external_id":   params.ExternalID,
		"full_name":     params.FullName,
		"domain":        params.Domain,
		"website":       params.Website,
		"timezone":      params.Timezone,
		"language":      params.Language,
		"industry":      params.Industry,
		"metadata":      params.Metadata,
		"address":       params.Address,
	})
	return parseResponse[Customer](r.http.put(ctx, fmt.Sprintf("/customers/%s", customerID), body, params.IdempotencyKey))
}

func (r *CustomersResource) List(ctx context.Context, params *ListCustomersParams) (*ApiResponse[[]Customer], error) {
	queryParams := map[string]string{}
	if params != nil {
		if params.CustomerID != "" {
			queryParams["customer_id"] = params.CustomerID
		}
		if params.IsActive != nil {
			queryParams["is_active"] = fmt.Sprintf("%t", *params.IsActive)
		}
		if params.Search != "" {
			queryParams["search"] = params.Search
		}
		if params.Limit != nil {
			queryParams["limit"] = fmt.Sprintf("%d", *params.Limit)
		}
		if params.Cursor != "" {
			queryParams["cursor"] = params.Cursor
		}
	}
	return parseResponse[[]Customer](r.http.get(ctx, "/customers", queryParams))
}

func (r *CustomersResource) Archive(ctx context.Context, customerID string, idempotencyKey string) (*ApiResponse[Customer], error) {
	body := map[string]any{"is_active": false}
	return parseResponse[Customer](r.http.put(ctx, fmt.Sprintf("/customers/%s", customerID), body, idempotencyKey))
}
