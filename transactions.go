package commet

import (
	"context"
	"fmt"
)

type RefundTransactionParams struct {
	Amount         *int   `json:"amount,omitempty"`
	IdempotencyKey string `json:"-"`
}

type RetryTransactionParams struct {
	IdempotencyKey string `json:"-"`
}

type ListTransactionsParams struct {
	Cursor        *string            `json:"cursor,omitempty"`
	Limit         *int               `json:"limit,omitempty"`
	Status        *TransactionStatus `json:"status,omitempty"`
	CustomerEmail *string            `json:"customer_email,omitempty"`
}

type TransactionsResource struct {
	http *httpClient
}

// Refund a positive amount in the payment currency’s smallest unit, or omit amount to refund the remaining amount. Returns the refund resource with its actual status.
func (r *TransactionsResource) Refund(ctx context.Context, id string, params *RefundTransactionParams) (*Refund, error) {
	body := buildBody(map[string]any{
		"amount": params.Amount,
	})
	return parseDirectResponse[Refund](r.http.post(ctx, fmt.Sprintf("/transactions/%s/refund", id), body, params.IdempotencyKey))
}

// Retry a failed subscription renewal and return an honest retry result. The original failed transaction remains immutable.
func (r *TransactionsResource) Retry(ctx context.Context, id string, params *RetryTransactionParams) (*TransactionRetry, error) {
	return parseDirectResponse[TransactionRetry](r.http.post(ctx, fmt.Sprintf("/transactions/%s/retry", id), map[string]any{}, params.IdempotencyKey))
}

// Retrieve a single payment transaction by its public ID, including provider details.
func (r *TransactionsResource) Get(ctx context.Context, id string) (*Transaction, error) {
	return parseDirectResponse[Transaction](r.http.get(ctx, fmt.Sprintf("/transactions/%s", id), nil))
}

// List payment transactions with cursor-based pagination. Filter by status or customer email.
func (r *TransactionsResource) List(ctx context.Context, params *ListTransactionsParams) (*TransactionsListResult, error) {
	query := map[string]string{}
	if params.Cursor != nil {
		query["cursor"] = *params.Cursor
	}
	if params.Limit != nil {
		query["limit"] = fmt.Sprintf("%d", *params.Limit)
	}
	if params.Status != nil {
		query["status"] = string(*params.Status)
	}
	if params.CustomerEmail != nil {
		query["customer_email"] = *params.CustomerEmail
	}
	return parseDirectResponse[TransactionsListResult](r.http.get(ctx, "/transactions", query))
}
