package commet

import (
	"context"
	"fmt"
)

type ListSubscriptionsParams struct {
	CustomerID *string             `json:"customer_id,omitempty"`
	Status     *SubscriptionStatus `json:"status,omitempty"`
}

type CreateSubscriptionParams struct {
	PlanID          *string                             `json:"plan_id,omitempty"`
	PlanCode        *string                             `json:"plan_code,omitempty"`
	CustomerID      string                              `json:"customer_id"`
	BillingInterval *BillingInterval                    `json:"billing_interval,omitempty"`
	InitialSeats    map[string]int                      `json:"initial_seats,omitempty"`
	SkipTrial       *bool                               `json:"skip_trial,omitempty"`
	IntroOffer      *CreateSubscriptionParamsIntroOffer `json:"intro_offer,omitempty"`
	Name            *string                             `json:"name,omitempty"`
	StartDate       *string                             `json:"start_date,omitempty"`
	SuccessURL      *string                             `json:"success_url,omitempty"`
	IdempotencyKey  string                              `json:"-"`
}

type GetActiveSubscriptionParams struct {
	CustomerID string `json:"customer_id"`
}

type CancelSubscriptionParams struct {
	Reason         *string `json:"reason,omitempty"`
	Immediate      *bool   `json:"immediate,omitempty"`
	IdempotencyKey string  `json:"-"`
}

type UncancelSubscriptionParams struct {
	IdempotencyKey string `json:"-"`
}

type ChangePlanParams struct {
	NewPlanID          *string `json:"new_plan_id,omitempty"`
	NewBillingInterval *string `json:"new_billing_interval,omitempty"`
	SuccessURL         *string `json:"success_url,omitempty"`
	IdempotencyKey     string  `json:"-"`
}

type PreviewChangePlanParams struct {
	PlanID          string           `json:"plan_id"`
	BillingInterval *BillingInterval `json:"billing_interval,omitempty"`
	IdempotencyKey  string           `json:"-"`
}

type ActivateAddonParams struct {
	AddonID        string `json:"addon_id"`
	IdempotencyKey string `json:"-"`
}

type AdjustBalanceParams struct {
	Amount         int     `json:"amount"`
	Reason         *string `json:"reason,omitempty"`
	Type           *string `json:"type,omitempty"`
	IdempotencyKey string  `json:"-"`
}

type TopupBalanceParams struct {
	Amount         int    `json:"amount"`
	IdempotencyKey string `json:"-"`
}

type PurchaseCreditsParams struct {
	CreditPackID   string `json:"credit_pack_id"`
	IdempotencyKey string `json:"-"`
}

type SubscriptionsResource struct {
	http *httpClient
}

// List all subscriptions. Filter by customer ID or status.
func (r *SubscriptionsResource) List(ctx context.Context, params *ListSubscriptionsParams) (*ApiResponse[[]Subscription], error) {
	query := map[string]string{}
	if params.CustomerID != nil {
		query["customer_id"] = *params.CustomerID
	}
	if params.Status != nil {
		query["status"] = string(*params.Status)
	}
	return parseResponse[[]Subscription](r.http.get(ctx, "/subscriptions", query))
}

// Create a subscription for a customer. Requires planId or planCode plus customerId.
func (r *SubscriptionsResource) Create(ctx context.Context, params *CreateSubscriptionParams) (*ApiResponse[Subscription], error) {
	body := buildBody(map[string]any{
		"plan_id":          params.PlanID,
		"plan_code":        params.PlanCode,
		"customer_id":      params.CustomerID,
		"billing_interval": params.BillingInterval,
		"initial_seats":    params.InitialSeats,
		"skip_trial":       params.SkipTrial,
		"intro_offer":      params.IntroOffer,
		"name":             params.Name,
		"start_date":       params.StartDate,
		"success_url":      params.SuccessURL,
	})
	return parseResponse[Subscription](r.http.post(ctx, "/subscriptions", body, params.IdempotencyKey))
}

// Get a subscription by its public ID, regardless of status (including pending_payment and past_due).
func (r *SubscriptionsResource) Get(ctx context.Context, id string) (*ApiResponse[Subscription], error) {
	return parseResponse[Subscription](r.http.get(ctx, fmt.Sprintf("/subscriptions/%s", id), nil))
}

// Get the active subscription for a customer. Returns null if none.
func (r *SubscriptionsResource) GetActive(ctx context.Context, params *GetActiveSubscriptionParams) (*ApiResponse[SubscriptionsGetActiveResult], error) {
	query := map[string]string{}
	if params.CustomerID != "" {
		query["customer_id"] = params.CustomerID
	}
	return parseResponse[SubscriptionsGetActiveResult](r.http.get(ctx, "/subscriptions/active", query))
}

// Cancel immediately or at period end.
func (r *SubscriptionsResource) Cancel(ctx context.Context, id string, params *CancelSubscriptionParams) (*ApiResponse[CanceledSubscription], error) {
	body := buildBody(map[string]any{
		"reason":    params.Reason,
		"immediate": params.Immediate,
	})
	return parseResponse[CanceledSubscription](r.http.post(ctx, fmt.Sprintf("/subscriptions/%s/cancel", id), body, params.IdempotencyKey))
}

// Revert a scheduled cancellation. Only works when canceledAt is set but status is not yet 'canceled'.
func (r *SubscriptionsResource) Uncancel(ctx context.Context, id string, params *UncancelSubscriptionParams) (*ApiResponse[UncanceledSubscription], error) {
	return parseResponse[UncanceledSubscription](r.http.post(ctx, fmt.Sprintf("/subscriptions/%s/uncancel", id), map[string]any{}, params.IdempotencyKey))
}

// Upgrade, downgrade, or change billing interval.
func (r *SubscriptionsResource) ChangePlan(ctx context.Context, id string, params *ChangePlanParams) (*ApiResponse[PlanChange], error) {
	body := buildBody(map[string]any{
		"new_plan_id":          params.NewPlanID,
		"new_billing_interval": params.NewBillingInterval,
		"success_url":          params.SuccessURL,
	})
	return parseResponse[PlanChange](r.http.post(ctx, fmt.Sprintf("/subscriptions/%s/change-plan", id), body, params.IdempotencyKey))
}

// Preview proration details for an immediate plan change (an upgrade or a longer interval) without applying it. Returns credit, charge, and net amount. Downgrades — a cheaper plan in the same group, or a shorter interval — are scheduled for the end of the current period instead of being prorated, so they return a 400 with code `plan_change_scheduled`; apply those via the change-plan endpoint.
func (r *SubscriptionsResource) PreviewChange(ctx context.Context, id string, params *PreviewChangePlanParams) (*ApiResponse[PreviewChange], error) {
	body := buildBody(map[string]any{
		"plan_id":          params.PlanID,
		"billing_interval": params.BillingInterval,
	})
	return parseResponse[PreviewChange](r.http.post(ctx, fmt.Sprintf("/subscriptions/%s/preview-change", id), body, params.IdempotencyKey))
}

// Activate an add-on on a subscription. Charges a prorated amount for the current billing period.
func (r *SubscriptionsResource) ActivateAddon(ctx context.Context, id string, params *ActivateAddonParams) (*ApiResponse[SubscriptionAddon], error) {
	body := buildBody(map[string]any{
		"addon_id": params.AddonID,
	})
	return parseResponse[SubscriptionAddon](r.http.post(ctx, fmt.Sprintf("/subscriptions/%s/addons", id), body, params.IdempotencyKey))
}

// Deactivate an add-on from a subscription.
func (r *SubscriptionsResource) DeactivateAddon(ctx context.Context, id string, addonID string) (*ApiResponse[DeletedSubscriptionAddon], error) {
	return parseResponse[DeletedSubscriptionAddon](r.http.delete(ctx, fmt.Sprintf("/subscriptions/%s/addons/%s", id, addonID), nil, ""))
}

// Adjust a subscription's balance or credits by a signed amount. Positive adds, negative subtracts.
func (r *SubscriptionsResource) AdjustBalance(ctx context.Context, id string, params *AdjustBalanceParams) (*ApiResponse[BalanceAdjustment], error) {
	body := buildBody(map[string]any{
		"amount": params.Amount,
		"reason": params.Reason,
		"type":   params.Type,
	})
	return parseResponse[BalanceAdjustment](r.http.post(ctx, fmt.Sprintf("/subscriptions/%s/balance/adjust", id), body, params.IdempotencyKey))
}

// Top up a subscription's balance. Charges the customer's payment method for the specified amount.
func (r *SubscriptionsResource) TopupBalance(ctx context.Context, id string, params *TopupBalanceParams) (*ApiResponse[BalanceTopup], error) {
	body := buildBody(map[string]any{
		"amount": params.Amount,
	})
	return parseResponse[BalanceTopup](r.http.post(ctx, fmt.Sprintf("/subscriptions/%s/balance/topup", id), body, params.IdempotencyKey))
}

// Purchase a credit pack for a subscription. Charges the customer and adds credits to their balance.
func (r *SubscriptionsResource) PurchaseCredits(ctx context.Context, id string, params *PurchaseCreditsParams) (*ApiResponse[CreditGrant], error) {
	body := buildBody(map[string]any{
		"credit_pack_id": params.CreditPackID,
	})
	return parseResponse[CreditGrant](r.http.post(ctx, fmt.Sprintf("/subscriptions/%s/credits", id), body, params.IdempotencyKey))
}
