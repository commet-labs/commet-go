package commet

import (
	"context"
	"fmt"
)

type DownloadInvoiceParams struct {
	IdempotencyKey string `json:"-"`
}

type SendInvoiceParams struct {
	IdempotencyKey string `json:"-"`
}

type UpdateInvoiceStatusParams struct {
	Status         string `json:"status"`
	IdempotencyKey string `json:"-"`
}

type ListInvoicesParams struct {
	Cursor         *string `json:"cursor,omitempty"`
	Limit          *int    `json:"limit,omitempty"`
	CustomerID     *string `json:"customer_id,omitempty"`
	Status         *string `json:"status,omitempty"`
	SubscriptionID *string `json:"subscription_id,omitempty"`
}

type CreateAdjustmentInvoiceParams struct {
	CustomerID     string         `json:"customer_id"`
	Amount         int            `json:"amount"`
	Description    string         `json:"description"`
	Metadata       map[string]any `json:"metadata,omitempty"`
	IdempotencyKey string         `json:"-"`
}

type InvoicesResource struct {
	http *httpClient
}

// Generate a signed URL to download the invoice as a PDF. The URL expires after 7 days.
func (r *InvoicesResource) GetDownloadURL(ctx context.Context, id string, params *DownloadInvoiceParams) (*InvoiceDownload, error) {
	return parseDirectResponse[InvoiceDownload](r.http.post(ctx, fmt.Sprintf("/invoices/%s/download-links", id), map[string]any{}, params.IdempotencyKey))
}

// Retrieve a single invoice by its public ID, including line items.
func (r *InvoicesResource) Get(ctx context.Context, id string) (*Invoice, error) {
	return parseDirectResponse[Invoice](r.http.get(ctx, fmt.Sprintf("/invoices/%s", id), nil))
}

// Send the invoice to the customer via email.
func (r *InvoicesResource) Send(ctx context.Context, id string, params *SendInvoiceParams) (*SentInvoice, error) {
	return parseDirectResponse[SentInvoice](r.http.post(ctx, fmt.Sprintf("/invoices/%s/send", id), map[string]any{}, params.IdempotencyKey))
}

// Mark an outstanding invoice as "paid" or "void" and return the updated invoice. Cannot change the status of already paid or voided invoices.
func (r *InvoicesResource) UpdateStatus(ctx context.Context, id string, params *UpdateInvoiceStatusParams) (*Invoice, error) {
	body := buildBody(map[string]any{
		"status": params.Status,
	})
	return parseDirectResponse[Invoice](r.http.patch(ctx, fmt.Sprintf("/invoices/%s/status", id), body, params.IdempotencyKey))
}

// List invoices with cursor-based pagination. Filter by customer, status, or subscription.
func (r *InvoicesResource) List(ctx context.Context, params *ListInvoicesParams) (*InvoicesListResult, error) {
	query := map[string]string{}
	if params.Cursor != nil {
		query["cursor"] = *params.Cursor
	}
	if params.Limit != nil {
		query["limit"] = fmt.Sprintf("%d", *params.Limit)
	}
	if params.CustomerID != nil {
		query["customer_id"] = *params.CustomerID
	}
	if params.Status != nil {
		query["status"] = *params.Status
	}
	if params.SubscriptionID != nil {
		query["subscription_id"] = *params.SubscriptionID
	}
	return parseDirectResponse[InvoicesListResult](r.http.get(ctx, "/invoices", query))
}

// Create a one-off adjustment invoice and return the created invoice. Use a negative amount for a credit.
func (r *InvoicesResource) CreateAdjustment(ctx context.Context, params *CreateAdjustmentInvoiceParams) (*Invoice, error) {
	body := buildBody(map[string]any{
		"customer_id": params.CustomerID,
		"amount":      params.Amount,
		"description": params.Description,
		"metadata":    params.Metadata,
	})
	return parseDirectResponse[Invoice](r.http.post(ctx, "/invoices", body, params.IdempotencyKey))
}
