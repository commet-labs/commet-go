package commet

import (
	"context"
	"fmt"
)

type ListInvoicesParams struct {
	CustomerID     string        `json:"customer_id,omitempty"`
	Status         InvoiceStatus `json:"status,omitempty"`
	SubscriptionID string        `json:"subscription_id,omitempty"`
	Limit          *int          `json:"limit,omitempty"`
	Cursor         string        `json:"cursor,omitempty"`
}

type CreateAdjustmentParams struct {
	CustomerID  string         `json:"customer_id"`
	Amount      float64        `json:"amount"`
	Description string         `json:"description,omitempty"`
	Metadata    map[string]any `json:"metadata,omitempty"`
}

type UpdateInvoiceStatusParams struct {
	// Status must be InvoiceStatusPaid or InvoiceStatusVoid; only outstanding
	// invoices can be changed.
	Status InvoiceStatus `json:"status"`
}

type InvoicesResource struct {
	http *httpClient
}

func (r *InvoicesResource) List(ctx context.Context, params *ListInvoicesParams) (*ApiResponse[[]InvoiceListItem], error) {
	queryParams := map[string]string{}
	if params != nil {
		if params.CustomerID != "" {
			queryParams["customer_id"] = params.CustomerID
		}
		if params.Status != "" {
			queryParams["status"] = string(params.Status)
		}
		if params.SubscriptionID != "" {
			queryParams["subscription_id"] = params.SubscriptionID
		}
		if params.Limit != nil {
			queryParams["limit"] = fmt.Sprintf("%d", *params.Limit)
		}
		if params.Cursor != "" {
			queryParams["cursor"] = params.Cursor
		}
	}
	return parseResponse[[]InvoiceListItem](r.http.get(ctx, "/invoices", queryParams))
}

func (r *InvoicesResource) Get(ctx context.Context, invoiceID string) (*ApiResponse[InvoiceDetail], error) {
	return parseResponse[InvoiceDetail](r.http.get(ctx, fmt.Sprintf("/invoices/%s", invoiceID), nil))
}

func (r *InvoicesResource) CreateAdjustment(ctx context.Context, params *CreateAdjustmentParams) (*ApiResponse[CreateAdjustmentResult], error) {
	body := buildBody(map[string]any{
		"customer_id": params.CustomerID,
		"amount":      params.Amount,
		"description": params.Description,
		"metadata":    params.Metadata,
	})
	return parseResponse[CreateAdjustmentResult](r.http.post(ctx, "/invoices", body, ""))
}

func (r *InvoicesResource) GetDownloadUrl(ctx context.Context, invoiceID string) (*ApiResponse[InvoiceDownloadResult], error) {
	return parseResponse[InvoiceDownloadResult](r.http.get(ctx, fmt.Sprintf("/invoices/%s/download", invoiceID), nil))
}

func (r *InvoicesResource) Send(ctx context.Context, invoiceID string) (*ApiResponse[InvoiceSendResult], error) {
	return parseResponse[InvoiceSendResult](r.http.post(ctx, fmt.Sprintf("/invoices/%s/send", invoiceID), map[string]any{}, ""))
}

func (r *InvoicesResource) UpdateStatus(ctx context.Context, invoiceID string, params *UpdateInvoiceStatusParams) (*ApiResponse[InvoiceStatusResult], error) {
	body := buildBody(map[string]any{
		"status": string(params.Status),
	})
	return parseResponse[InvoiceStatusResult](r.http.put(ctx, fmt.Sprintf("/invoices/%s/status", invoiceID), body, ""))
}
