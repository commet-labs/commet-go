package commet

import (
	"context"
	"fmt"
)

type ListTransactionsParams struct {
	Status        TransactionStatus `json:"status,omitempty"`
	CustomerEmail string            `json:"customer_email,omitempty"`
	Limit         *int              `json:"limit,omitempty"`
	Cursor        string            `json:"cursor,omitempty"`
}

type TransactionsResource struct {
	http *httpClient
}

func (r *TransactionsResource) List(ctx context.Context, params *ListTransactionsParams) (*ApiResponse[[]TransactionListItem], error) {
	queryParams := map[string]string{}
	if params != nil {
		if params.Status != "" {
			queryParams["status"] = string(params.Status)
		}
		if params.CustomerEmail != "" {
			queryParams["customer_email"] = params.CustomerEmail
		}
		if params.Limit != nil {
			queryParams["limit"] = fmt.Sprintf("%d", *params.Limit)
		}
		if params.Cursor != "" {
			queryParams["cursor"] = params.Cursor
		}
	}
	return parseResponse[[]TransactionListItem](r.http.get(ctx, "/transactions", queryParams))
}

func (r *TransactionsResource) Get(ctx context.Context, transactionID string) (*ApiResponse[TransactionDetail], error) {
	return parseResponse[TransactionDetail](r.http.get(ctx, fmt.Sprintf("/transactions/%s", transactionID), nil))
}

func (r *TransactionsResource) Refund(ctx context.Context, transactionID string) (*ApiResponse[TransactionRefundResult], error) {
	return parseResponse[TransactionRefundResult](r.http.post(ctx, fmt.Sprintf("/transactions/%s/refund", transactionID), map[string]any{}, ""))
}

func (r *TransactionsResource) Retry(ctx context.Context, transactionID string) (*ApiResponse[TransactionRetryResult], error) {
	return parseResponse[TransactionRetryResult](r.http.post(ctx, fmt.Sprintf("/transactions/%s/retry", transactionID), map[string]any{}, ""))
}
