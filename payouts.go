package commet

import "context"

type AddPayoutBankAccountParams struct {
	AccountNumber     string  `json:"account_number"`
	AccountHolderName string  `json:"account_holder_name"`
	RoutingNumber     *string `json:"routing_number,omitempty"`
	AccountType       *string `json:"account_type,omitempty"`
	SetDefault        *bool   `json:"set_default,omitempty"`
	IdempotencyKey    string  `json:"-"`
}

type RequestPayoutParams struct {
	Amount         int     `json:"amount"`
	Description    *string `json:"description,omitempty"`
	IdempotencyKey string  `json:"-"`
}

type PayoutsResource struct {
	http *httpClient
}

// Add an additional destination bank account to the organization's existing payout account. Country and currency are resolved from the organization. The full account number is never returned — only `last4`.
func (r *PayoutsResource) AddBankAccount(ctx context.Context, params *AddPayoutBankAccountParams) (*PayoutBankAccount, error) {
	body := buildBody(map[string]any{
		"account_number":      params.AccountNumber,
		"account_holder_name": params.AccountHolderName,
		"routing_number":      params.RoutingNumber,
		"account_type":        params.AccountType,
		"set_default":         params.SetDefault,
	})
	return parseDirectResponse[PayoutBankAccount](r.http.post(ctx, "/payouts/bank-accounts", body, params.IdempotencyKey))
}

// Withdraw available balance to the organization's verified payout account. `amount` is in cents (USD, minimum 1000 = $10). The payout is created in `pending` and settles to `paid` asynchronously as provider webhooks arrive.
func (r *PayoutsResource) Request(ctx context.Context, params *RequestPayoutParams) (*Payout, error) {
	body := buildBody(map[string]any{
		"amount":      params.Amount,
		"description": params.Description,
	})
	return parseDirectResponse[Payout](r.http.post(ctx, "/payouts", body, params.IdempotencyKey))
}

// Deprecated. Complete business and identity verification in the Commet dashboard. This endpoint no longer accepts or processes KYC data.
// Deprecated.
func (r *PayoutsResource) CompleteVerification(ctx context.Context) (*struct{}, error) {
	return parseDirectResponse[struct{}](r.http.post(ctx, "/payouts/verification", map[string]any{}, ""))
}
