package commet

import (
	"context"
	"fmt"
)

type RevokeCustomerCreditParams struct {
	IdempotencyKey string `json:"-"`
}

type CreateCustomerCreditParams struct {
	Amount         int     `json:"amount"`
	Currency       string  `json:"currency"`
	Reason         string  `json:"reason"`
	ExpiresAt      *string `json:"expires_at,omitempty"`
	IdempotencyKey string  `json:"-"`
}

type RevokePlanGrantParams struct {
	Reason         string `json:"reason"`
	IdempotencyKey string `json:"-"`
}

type UpdatePlanGrantParams struct {
	Reason         string  `json:"reason"`
	Duration       string  `json:"duration"`
	DurationCycles *int    `json:"duration_cycles,omitempty"`
	ExpiresAt      *string `json:"expires_at,omitempty"`
	IdempotencyKey string  `json:"-"`
}

type CreatePlanGrantParams struct {
	SubscriptionID string  `json:"subscription_id"`
	PlanID         string  `json:"plan_id"`
	Reason         string  `json:"reason"`
	Duration       string  `json:"duration"`
	DurationCycles *int    `json:"duration_cycles,omitempty"`
	ExpiresAt      *string `json:"expires_at,omitempty"`
	IdempotencyKey string  `json:"-"`
}

type UpdateCustomerParams struct {
	Email          *string                      `json:"email,omitempty"`
	FullName       *string                      `json:"full_name,omitempty"`
	TaxDocument    *string                      `json:"tax_document,omitempty"`
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

type ListCustomersParams struct {
	Cursor     *string `json:"cursor,omitempty"`
	Limit      *int    `json:"limit,omitempty"`
	ExternalID *string `json:"external_id,omitempty"`
}

type CreateCustomerParams struct {
	ID             *string                      `json:"id,omitempty"`
	ExternalID     *string                      `json:"external_id,omitempty"`
	FullName       *string                      `json:"full_name,omitempty"`
	TaxDocument    *string                      `json:"tax_document,omitempty"`
	Address        *CreateCustomerParamsAddress `json:"address,omitempty"`
	AddressID      *string                      `json:"address_id,omitempty"`
	Email          string                       `json:"email"`
	Timezone       *Timezone                    `json:"timezone,omitempty"`
	Metadata       map[string]any               `json:"metadata,omitempty"`
	IdempotencyKey string                       `json:"-"`
}

type CustomersResource struct {
	http *httpClient
}

// Revoke the unallocated remainder of a customer credit grant. Applied invoice history is unchanged.
func (r *CustomersResource) RevokeCredit(ctx context.Context, id string, creditID string, params *RevokeCustomerCreditParams) (*CustomerCreditRevocation, error) {
	return parseDirectResponse[CustomerCreditRevocation](r.http.post(ctx, fmt.Sprintf("/customers/%s/credits/%s/revoke", id, creditID), map[string]any{}, params.IdempotencyKey))
}

// List currency-specific invoice credit grants and their remaining balances for a customer.
func (r *CustomersResource) ListCredits(ctx context.Context, id string) (*CustomersListCreditsResult, error) {
	return parseDirectResponse[CustomersListCreditsResult](r.http.get(ctx, fmt.Sprintf("/customers/%s/credits", id), nil))
}

// Grant monetary credit in one currency. Credit is applied FIFO before tax to eligible recurring invoices.
func (r *CustomersResource) CreateCredit(ctx context.Context, id string, params *CreateCustomerCreditParams) (*CustomerCredit, error) {
	body := buildBody(map[string]any{
		"amount":     params.Amount,
		"currency":   params.Currency,
		"reason":     params.Reason,
		"expires_at": params.ExpiresAt,
	})
	return parseDirectResponse[CustomerCredit](r.http.post(ctx, fmt.Sprintf("/customers/%s/credits", id), body, params.IdempotencyKey))
}

// End expanded access immediately and restore the base plan's limits. The subscription, billing cycle, invoices, and payment state remain unchanged.
func (r *CustomersResource) RevokePlanGrant(ctx context.Context, id string, grantID string, params *RevokePlanGrantParams) (*PlanGrant, error) {
	body := buildBody(map[string]any{
		"reason": params.Reason,
	})
	return parseDirectResponse[PlanGrant](r.http.post(ctx, fmt.Sprintf("/customers/%s/plan-grants/%s/revoke", id, grantID), body, params.IdempotencyKey))
}

// Keep the overlay for a number of the subscription's existing billing cycles, set an exact deadline, or leave it active until revoked. The billing anchor is never reset.
func (r *CustomersResource) UpdatePlanGrant(ctx context.Context, id string, grantID string, params *UpdatePlanGrantParams) (*PlanGrant, error) {
	body := buildBody(map[string]any{
		"reason":          params.Reason,
		"duration":        params.Duration,
		"duration_cycles": params.DurationCycles,
		"expires_at":      params.ExpiresAt,
	})
	return parseDirectResponse[PlanGrant](r.http.patch(ctx, fmt.Sprintf("/customers/%s/plan-grants/%s", id, grantID), body, params.IdempotencyKey))
}

// List the independent audit timeline for paid-plan access granted without checkout or payment credentials.
func (r *CustomersResource) ListPlanGrants(ctx context.Context, id string) (*CustomersListPlanGrantsResult, error) {
	return parseDirectResponse[CustomersListPlanGrantsResult](r.http.get(ctx, fmt.Sprintf("/customers/%s/plan-grants", id), nil))
}

// Temporarily expand an active subscription's feature access using a higher plan in the same plan group. Billing, prices, periods, invoices, and the base subscription remain unchanged.
func (r *CustomersResource) CreatePlanGrant(ctx context.Context, id string, params *CreatePlanGrantParams) (*PlanGrant, error) {
	body := buildBody(map[string]any{
		"subscription_id": params.SubscriptionID,
		"plan_id":         params.PlanID,
		"reason":          params.Reason,
		"duration":        params.Duration,
		"duration_cycles": params.DurationCycles,
		"expires_at":      params.ExpiresAt,
	})
	return parseDirectResponse[PlanGrant](r.http.post(ctx, fmt.Sprintf("/customers/%s/plan-grants", id), body, params.IdempotencyKey))
}

// Retrieve a customer by their public ID, including subscription status and metadata.
func (r *CustomersResource) Get(ctx context.Context, id string) (*Customer, error) {
	return parseDirectResponse[Customer](r.http.get(ctx, fmt.Sprintf("/customers/%s", id), nil))
}

// Update a customer's name, external ID, or metadata.
func (r *CustomersResource) Update(ctx context.Context, id string, params *UpdateCustomerParams) (*Customer, error) {
	body := buildBody(map[string]any{
		"email":        params.Email,
		"full_name":    params.FullName,
		"tax_document": params.TaxDocument,
		"external_id":  params.ExternalID,
		"timezone":     params.Timezone,
		"metadata":     params.Metadata,
		"address":      params.Address,
	})
	return parseDirectResponse[Customer](r.http.patch(ctx, fmt.Sprintf("/customers/%s", id), body, params.IdempotencyKey))
}

// Create up to 100 customers in a single request.
func (r *CustomersResource) CreateBatch(ctx context.Context, params *BatchCreateCustomersParams) (*CustomerBatch, error) {
	body := buildBody(map[string]any{
		"customers": params.Customers,
	})
	return parseDirectResponse[CustomerBatch](r.http.post(ctx, "/customers/batch", body, params.IdempotencyKey))
}

// List customers with cursor-based pagination.
func (r *CustomersResource) List(ctx context.Context, params *ListCustomersParams) (*CustomersListResult, error) {
	query := map[string]string{}
	if params.Cursor != nil {
		query["cursor"] = *params.Cursor
	}
	if params.Limit != nil {
		query["limit"] = fmt.Sprintf("%d", *params.Limit)
	}
	if params.ExternalID != nil {
		query["external_id"] = *params.ExternalID
	}
	return parseDirectResponse[CustomersListResult](r.http.get(ctx, "/customers", query))
}

// Create a new customer. Idempotent when customerId is provided.
func (r *CustomersResource) Create(ctx context.Context, params *CreateCustomerParams) (*Customer, error) {
	body := buildBody(map[string]any{
		"id":           params.ID,
		"external_id":  params.ExternalID,
		"full_name":    params.FullName,
		"tax_document": params.TaxDocument,
		"address":      params.Address,
		"address_id":   params.AddressID,
		"email":        params.Email,
		"timezone":     params.Timezone,
		"metadata":     params.Metadata,
	})
	return parseDirectResponse[Customer](r.http.post(ctx, "/customers", body, params.IdempotencyKey))
}
