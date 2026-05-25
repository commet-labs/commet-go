package commet

import (
	"context"
	"fmt"
)

type CreateSubscriptionParams struct {
	CustomerID      string         `json:"customer_id,omitempty"`
	PlanCode        string         `json:"plan_code,omitempty"`
	PlanID          string         `json:"plan_id,omitempty"`
	BillingInterval string         `json:"billing_interval,omitempty"`
	InitialSeats    map[string]int `json:"initial_seats,omitempty"`
	SkipTrial       *bool          `json:"skip_trial,omitempty"`
	Name            string         `json:"name,omitempty"`
	StartDate       string         `json:"start_date,omitempty"`
	SuccessURL      string         `json:"success_url,omitempty"`
	IdempotencyKey  string         `json:"-"`
}

type CancelSubscriptionParams struct {
	Reason         string `json:"reason,omitempty"`
	Immediate      *bool  `json:"immediate,omitempty"`
	IdempotencyKey string `json:"-"`
}

type ChangePlanParams struct {
	SubscriptionID     string `json:"-"`
	NewPlanID          string `json:"new_plan_id,omitempty"`
	NewBillingInterval string `json:"new_billing_interval,omitempty"`
	IdempotencyKey     string `json:"-"`
}

type ChangePlanPlanInfo struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price,omitempty"`
}

type ChangePlanBilling struct {
	Credit                float64 `json:"credit"`
	CreditsApplied        float64 `json:"credits_applied"`
	Charge                float64 `json:"charge"`
	TaxAmount             float64 `json:"tax_amount"`
	NetAmount             float64 `json:"net_amount"`
	TotalCharged          float64 `json:"total_charged"`
	RemainingCreditBalance float64 `json:"remaining_credit_balance"`
}

type ChangePlanResult struct {
	ID               string              `json:"id"`
	Scheduled        bool                `json:"scheduled"`
	CustomerID       string              `json:"customer_id,omitempty"`
	PreviousPlan     *ChangePlanPlanInfo `json:"previous_plan,omitempty"`
	CurrentPlan      *ChangePlanPlanInfo `json:"current_plan,omitempty"`
	BillingInterval  string              `json:"billing_interval,omitempty"`
	Billing          *ChangePlanBilling  `json:"billing,omitempty"`
	InvoiceID        string              `json:"invoice_id,omitempty"`
	ScheduledFor     string              `json:"scheduled_for,omitempty"`
	ChangeType       string              `json:"change_type,omitempty"`
	EventID          string              `json:"event_id,omitempty"`
	NewPlanID        string              `json:"new_plan_id,omitempty"`
	NewPlanName      string              `json:"new_plan_name,omitempty"`
	NewBillingInterval string            `json:"new_billing_interval,omitempty"`
	RequiresCheckout bool                `json:"requires_checkout,omitempty"`
	CheckoutURL      string              `json:"checkout_url,omitempty"`
}

type SubscriptionsResource struct {
	http *httpClient
}

func (r *SubscriptionsResource) Create(ctx context.Context, params *CreateSubscriptionParams) (*ApiResponse[Subscription], error) {
	body := buildBody(map[string]any{
		"customer_id":      params.CustomerID,
		"plan_code":        params.PlanCode,
		"plan_id":          params.PlanID,
		"billing_interval": params.BillingInterval,
		"initial_seats":    params.InitialSeats,
		"skip_trial":       params.SkipTrial,
		"name":             params.Name,
		"start_date":       params.StartDate,
		"success_url":      params.SuccessURL,
	})
	return parseResponse[Subscription](r.http.post(ctx, "/subscriptions", body, params.IdempotencyKey))
}

func (r *SubscriptionsResource) Get(ctx context.Context, customerID string) (*ApiResponse[ActiveSubscription], error) {
	return parseResponse[ActiveSubscription](r.http.get(ctx, "/subscriptions/active", map[string]string{"customer_id": customerID}))
}

func (r *SubscriptionsResource) Cancel(ctx context.Context, subscriptionID string, params *CancelSubscriptionParams) (*ApiResponse[Subscription], error) {
	body := buildBody(map[string]any{
		"reason":    params.Reason,
		"immediate": params.Immediate,
	})
	return parseResponse[Subscription](r.http.post(ctx, fmt.Sprintf("/subscriptions/%s/cancel", subscriptionID), body, params.IdempotencyKey))
}

func (r *SubscriptionsResource) Uncancel(ctx context.Context, subscriptionID string) (*ApiResponse[Subscription], error) {
	return parseResponse[Subscription](r.http.post(ctx, fmt.Sprintf("/subscriptions/%s/uncancel", subscriptionID), map[string]any{}, ""))
}

func (r *SubscriptionsResource) ChangePlan(ctx context.Context, params *ChangePlanParams) (*ApiResponse[ChangePlanResult], error) {
	body := buildBody(map[string]any{
		"new_plan_id":          params.NewPlanID,
		"new_billing_interval": params.NewBillingInterval,
	})
	return parseResponse[ChangePlanResult](r.http.post(ctx, fmt.Sprintf("/subscriptions/%s/change-plan", params.SubscriptionID), body, params.IdempotencyKey))
}
