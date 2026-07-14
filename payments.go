package commet

import (
	"context"
	"fmt"
)

type ListPaymentsParams struct {
	CustomerID *string `json:"customer_id,omitempty"`
	Cursor     *string `json:"cursor,omitempty"`
	Limit      *int    `json:"limit,omitempty"`
}

type CreatePaymentParams struct {
	Amount         int               `json:"amount"`
	Currency       string            `json:"currency"`
	CustomerID     *string           `json:"customer_id,omitempty"`
	Description    string            `json:"description"`
	SuccessURL     *string           `json:"success_url,omitempty"`
	Metadata       map[string]string `json:"metadata,omitempty"`
	IdempotencyKey string            `json:"-"`
}

type ChargePaymentParams struct {
	CustomerID     string            `json:"customer_id"`
	Amount         int               `json:"amount"`
	Currency       string            `json:"currency"`
	Description    string            `json:"description"`
	Metadata       map[string]string `json:"metadata,omitempty"`
	IdempotencyKey string            `json:"-"`
}

type CancelPaymentParams struct {
	IdempotencyKey string `json:"-"`
}

type PaymentsResource struct {
	http *httpClient
}

// List payments with cursor-based pagination. Filter by customer.
func (r *PaymentsResource) List(ctx context.Context, params *ListPaymentsParams) (*ApiResponse[[]Payment], error) {
	query := map[string]string{}
	if params.CustomerID != nil {
		query["customer_id"] = *params.CustomerID
	}
	if params.Cursor != nil {
		query["cursor"] = *params.Cursor
	}
	if params.Limit != nil {
		query["limit"] = fmt.Sprintf("%d", *params.Limit)
	}
	return parseResponse[[]Payment](r.http.get(ctx, "/payments", query))
}

// Create a hosted payment link. Returns a url the customer opens to pay with any card. Calculates tax, generates an invoice, and vaults the payment method on confirmation. No subscription or plan required.
func (r *PaymentsResource) Create(ctx context.Context, params *CreatePaymentParams) (*ApiResponse[Payment], error) {
	body := buildBody(map[string]any{
		"amount":      params.Amount,
		"currency":    params.Currency,
		"customer_id": params.CustomerID,
		"description": params.Description,
		"success_url": params.SuccessURL,
		"metadata":    params.Metadata,
	})
	return parseResponse[Payment](r.http.post(ctx, "/payments", body, params.IdempotencyKey))
}

// Charge a customer's vaulted payment method off-session. Calculates tax, generates an invoice, and sends a receipt. Requires the customer to have a subscription in active, trialing, or past_due state.
func (r *PaymentsResource) Charge(ctx context.Context, params *ChargePaymentParams) (*ApiResponse[Payment], error) {
	body := buildBody(map[string]any{
		"customer_id": params.CustomerID,
		"amount":      params.Amount,
		"currency":    params.Currency,
		"description": params.Description,
		"metadata":    params.Metadata,
	})
	return parseResponse[Payment](r.http.post(ctx, "/payments/charge", body, params.IdempotencyKey))
}

// Retrieve a payment by its public ID.
func (r *PaymentsResource) Get(ctx context.Context, id string) (*ApiResponse[Payment], error) {
	return parseResponse[Payment](r.http.get(ctx, fmt.Sprintf("/payments/%s", id), nil))
}

// Cancel a pending payment link so it can no longer be paid. Only a link that has not been paid or started processing can be canceled; canceling an already canceled link is a no-op. Charges cannot be canceled.
func (r *PaymentsResource) Cancel(ctx context.Context, id string, params *CancelPaymentParams) (*ApiResponse[Payment], error) {
	return parseResponse[Payment](r.http.post(ctx, fmt.Sprintf("/payments/%s/cancel", id), map[string]any{}, params.IdempotencyKey))
}
