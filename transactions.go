package commet

import (
	"context"
	"fmt"
)

type ListTransactionsParams struct {
	Status        *TransactionStatus `json:"status,omitempty"`
	CustomerEmail *string            `json:"customer_email,omitempty"`
	Limit         *int               `json:"limit,omitempty"`
	Cursor        *string            `json:"cursor,omitempty"`
}

type RefundTransactionParams struct {
	IdempotencyKey string `json:"-"`
}

type RetryTransactionParams struct {
	IdempotencyKey string `json:"-"`
}

type TransactionsResource struct {
	http *httpClient
}

// List payment transactions with cursor-based pagination. Filter by status or customer email.
func (r *TransactionsResource) List(ctx context.Context, params *ListTransactionsParams) (*ApiResponse[[]Transaction], error) {
	query := map[string]string{}
	if params.Status != nil {
		query["status"] = string(*params.Status)
	}
	if params.CustomerEmail != nil {
		query["customer_email"] = *params.CustomerEmail
	}
	if params.Limit != nil {
		query["limit"] = fmt.Sprintf("%d", *params.Limit)
	}
	if params.Cursor != nil {
		query["cursor"] = *params.Cursor
	}
	return parseResponse[[]Transaction](r.http.get(ctx, "/transactions", query))
}

// Retrieve a single payment transaction by its public ID, including provider details.
func (r *TransactionsResource) Get(ctx context.Context, id string) (*ApiResponse[Transaction], error) {
	return parseResponse[Transaction](r.http.get(ctx, fmt.Sprintf("/transactions/%s", id), nil))
}

// Issue a full refund for a payment transaction.
func (r *TransactionsResource) Refund(ctx context.Context, id string, params *RefundTransactionParams) (*ApiResponse[TransactionRefund], error) {
	return parseResponse[TransactionRefund](r.http.post(ctx, fmt.Sprintf("/transactions/%s/refund", id), map[string]any{}, params.IdempotencyKey))
}

// Retry a failed payment transaction. Creates a new invoice and initiates a new payment attempt.
func (r *TransactionsResource) Retry(ctx context.Context, id string, params *RetryTransactionParams) (*ApiResponse[TransactionRetry], error) {
	return parseResponse[TransactionRetry](r.http.post(ctx, fmt.Sprintf("/transactions/%s/retry", id), map[string]any{}, params.IdempotencyKey))
}
