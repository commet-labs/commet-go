package commet

import (
	"context"
	"fmt"
)

// CreateCustomerParams holds parameters for creating a customer.
type CreateCustomerParams struct {
	Email          string         `json:"billing_email"`
	ID             string         `json:"-"`
	FullName       string         `json:"full_name,omitempty"`
	Domain         string         `json:"domain,omitempty"`
	Website        string         `json:"website,omitempty"`
	Timezone       string         `json:"timezone,omitempty"`
	Language       string         `json:"language,omitempty"`
	Industry       string         `json:"industry,omitempty"`
	Metadata       map[string]any `json:"metadata,omitempty"`
	Address        map[string]string `json:"address,omitempty"`
	IdempotencyKey string         `json:"-"`
}

// UpdateCustomerParams holds parameters for updating a customer.
type UpdateCustomerParams struct {
	Email      string            `json:"billing_email,omitempty"`
	ExternalID string            `json:"external_id,omitempty"`
	FullName   string            `json:"full_name,omitempty"`
	Domain     string            `json:"domain,omitempty"`
	Website    string            `json:"website,omitempty"`
	Timezone   string            `json:"timezone,omitempty"`
	Language   string            `json:"language,omitempty"`
	Industry   string            `json:"industry,omitempty"`
	Metadata   map[string]any    `json:"metadata,omitempty"`
	Address    map[string]string `json:"address,omitempty"`
	IdempotencyKey string        `json:"-"`
}

// ListCustomersParams holds parameters for listing customers.
type ListCustomersParams struct {
	CustomerID string `json:"customer_id,omitempty"`
	IsActive   *bool  `json:"is_active,omitempty"`
	Search     string `json:"search,omitempty"`
	Limit      *int   `json:"limit,omitempty"`
	Cursor     string `json:"cursor,omitempty"`
}

// CustomersResource provides access to customer operations.
type CustomersResource struct {
	http *httpClient
}

// Create creates a new customer.
func (r *CustomersResource) Create(ctx context.Context, params *CreateCustomerParams) (*ApiResponse, error) {
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
	return r.http.post(ctx, "/customers", body, params.IdempotencyKey)
}

// CreateBatch creates multiple customers in a single request.
func (r *CustomersResource) CreateBatch(ctx context.Context, customers []CreateCustomerParams, idempotencyKey string) (*ApiResponse, error) {
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
	return r.http.post(ctx, "/customers/batch", body, idempotencyKey)
}

// Get retrieves a customer by ID.
func (r *CustomersResource) Get(ctx context.Context, customerID string) (*ApiResponse, error) {
	return r.http.get(ctx, fmt.Sprintf("/customers/%s", customerID), nil)
}

// Update updates an existing customer.
func (r *CustomersResource) Update(ctx context.Context, customerID string, params *UpdateCustomerParams) (*ApiResponse, error) {
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
	return r.http.put(ctx, fmt.Sprintf("/customers/%s", customerID), body, params.IdempotencyKey)
}

// List retrieves a paginated list of customers.
func (r *CustomersResource) List(ctx context.Context, params *ListCustomersParams) (*ApiResponse, error) {
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
	return r.http.get(ctx, "/customers", queryParams)
}

// Archive marks a customer as inactive.
func (r *CustomersResource) Archive(ctx context.Context, customerID string, idempotencyKey string) (*ApiResponse, error) {
	body := map[string]any{"is_active": false}
	return r.http.put(ctx, fmt.Sprintf("/customers/%s", customerID), body, idempotencyKey)
}
