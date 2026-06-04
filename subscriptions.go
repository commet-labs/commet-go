package commet

import (
	"context"
	"fmt"
)

type CustomIntroOffer struct {
	DiscountType   string `json:"discount_type"`
	DiscountValue  int    `json:"discount_value"`
	DurationCycles int    `json:"duration_cycles"`
}

type CreateSubscriptionParams struct {
	CustomerID       string            `json:"customer_id,omitempty"`
	PlanCode         string            `json:"plan_code,omitempty"`
	PlanID           string            `json:"plan_id,omitempty"`
	BillingInterval  string            `json:"billing_interval,omitempty"`
	InitialSeats     map[string]int    `json:"initial_seats,omitempty"`
	SkipTrial        *bool             `json:"skip_trial,omitempty"`
	Name             string            `json:"name,omitempty"`
	StartDate        string            `json:"start_date,omitempty"`
	SuccessURL       string            `json:"success_url,omitempty"`
	CustomIntroOffer *CustomIntroOffer `json:"custom_intro_offer,omitempty"`
	IdempotencyKey   string            `json:"-"`
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
	SuccessURL         string `json:"success_url,omitempty"`
	IdempotencyKey     string `json:"-"`
}

type ChangePlanPlanInfo struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price,omitempty"`
}

type ChangePlanBilling struct {
	Credit                 float64 `json:"credit"`
	CreditsApplied         float64 `json:"credits_applied"`
	Charge                 float64 `json:"charge"`
	TaxAmount              float64 `json:"tax_amount"`
	NetAmount              float64 `json:"net_amount"`
	TotalCharged           float64 `json:"total_charged"`
	RemainingCreditBalance float64 `json:"remaining_credit_balance"`
}

type ChangePlanResult struct {
	ID                 string              `json:"id"`
	Scheduled          bool                `json:"scheduled"`
	CustomerID         string              `json:"customer_id,omitempty"`
	PreviousPlan       *ChangePlanPlanInfo `json:"previous_plan,omitempty"`
	CurrentPlan        *ChangePlanPlanInfo `json:"current_plan,omitempty"`
	BillingInterval    string              `json:"billing_interval,omitempty"`
	Billing            *ChangePlanBilling  `json:"billing,omitempty"`
	InvoiceID          string              `json:"invoice_id,omitempty"`
	ScheduledFor       string              `json:"scheduled_for,omitempty"`
	ChangeType         string              `json:"change_type,omitempty"`
	EventID            string              `json:"event_id,omitempty"`
	NewPlanID          string              `json:"new_plan_id,omitempty"`
	NewPlanName        string              `json:"new_plan_name,omitempty"`
	NewBillingInterval string              `json:"new_billing_interval,omitempty"`
	RequiresCheckout   bool                `json:"requires_checkout,omitempty"`
	CheckoutURL        string              `json:"checkout_url,omitempty"`
}

type ListSubscriptionsParams struct {
	CustomerID string             `json:"customer_id,omitempty"`
	Status     SubscriptionStatus `json:"status,omitempty"`
	Limit      *int               `json:"limit,omitempty"`
	Cursor     string             `json:"cursor,omitempty"`
}

type PreviewChangeParams struct {
	PlanID          string `json:"plan_id,omitempty"`
	BillingInterval string `json:"billing_interval,omitempty"`
}

type ActivateAddonParams struct {
	AddonID string `json:"addon_id"`
}

type DeactivateAddonParams struct {
	AddonID string `json:"-"`
}

type AdjustBalanceParams struct {
	Amount float64 `json:"amount"`
	Reason string  `json:"reason,omitempty"`
	Type   string  `json:"type,omitempty"`
}

type TopupBalanceParams struct {
	Amount float64 `json:"amount"`
}

type PurchaseCreditsParams struct {
	CreditPackID string `json:"credit_pack_id"`
}

type SubscriptionsResource struct {
	http *httpClient
}

func (r *SubscriptionsResource) Create(ctx context.Context, params *CreateSubscriptionParams) (*ApiResponse[Subscription], error) {
	fields := map[string]any{
		"customer_id":      params.CustomerID,
		"plan_code":        params.PlanCode,
		"plan_id":          params.PlanID,
		"billing_interval": params.BillingInterval,
		"initial_seats":    params.InitialSeats,
		"skip_trial":       params.SkipTrial,
		"name":             params.Name,
		"start_date":       params.StartDate,
		"success_url":      params.SuccessURL,
	}
	if params.CustomIntroOffer != nil {
		fields["custom_intro_offer"] = map[string]any{
			"discount_type":   params.CustomIntroOffer.DiscountType,
			"discount_value":  params.CustomIntroOffer.DiscountValue,
			"duration_cycles": params.CustomIntroOffer.DurationCycles,
		}
	}
	body := buildBody(fields)
	return parseResponse[Subscription](r.http.post(ctx, "/subscriptions", body, params.IdempotencyKey))
}

func (r *SubscriptionsResource) GetActive(ctx context.Context, customerID string) (*ApiResponse[ActiveSubscription], error) {
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
		"success_url":          params.SuccessURL,
	})
	return parseResponse[ChangePlanResult](r.http.post(ctx, fmt.Sprintf("/subscriptions/%s/change-plan", params.SubscriptionID), body, params.IdempotencyKey))
}

func (r *SubscriptionsResource) List(ctx context.Context, params *ListSubscriptionsParams) (*ApiResponse[[]SubscriptionListItem], error) {
	queryParams := map[string]string{}
	if params != nil {
		if params.CustomerID != "" {
			queryParams["customer_id"] = params.CustomerID
		}
		if params.Status != "" {
			queryParams["status"] = string(params.Status)
		}
		if params.Limit != nil {
			queryParams["limit"] = fmt.Sprintf("%d", *params.Limit)
		}
		if params.Cursor != "" {
			queryParams["cursor"] = params.Cursor
		}
	}
	return parseResponse[[]SubscriptionListItem](r.http.get(ctx, "/subscriptions", queryParams))
}

func (r *SubscriptionsResource) PreviewChange(ctx context.Context, subscriptionID string, params *PreviewChangeParams) (*ApiResponse[PreviewChangeResult], error) {
	body := buildBody(map[string]any{
		"plan_id":          params.PlanID,
		"billing_interval": params.BillingInterval,
	})
	return parseResponse[PreviewChangeResult](r.http.post(ctx, fmt.Sprintf("/subscriptions/%s/preview-change", subscriptionID), body, ""))
}

func (r *SubscriptionsResource) ActivateAddon(ctx context.Context, subscriptionID string, params *ActivateAddonParams) (*ApiResponse[ActivateAddonResult], error) {
	body := buildBody(map[string]any{
		"addon_id": params.AddonID,
	})
	return parseResponse[ActivateAddonResult](r.http.post(ctx, fmt.Sprintf("/subscriptions/%s/addons", subscriptionID), body, ""))
}

func (r *SubscriptionsResource) DeactivateAddon(ctx context.Context, subscriptionID string, addonID string) (*ApiResponse[DeactivateAddonResult], error) {
	return parseResponse[DeactivateAddonResult](r.http.delete(ctx, fmt.Sprintf("/subscriptions/%s/addons/%s", subscriptionID, addonID), nil, ""))
}

func (r *SubscriptionsResource) AdjustBalance(ctx context.Context, subscriptionID string, params *AdjustBalanceParams) (*ApiResponse[AdjustBalanceResult], error) {
	body := buildBody(map[string]any{
		"amount": params.Amount,
		"reason": params.Reason,
		"type":   params.Type,
	})
	return parseResponse[AdjustBalanceResult](r.http.post(ctx, fmt.Sprintf("/subscriptions/%s/balance/adjust", subscriptionID), body, ""))
}

func (r *SubscriptionsResource) TopupBalance(ctx context.Context, subscriptionID string, params *TopupBalanceParams) (*ApiResponse[TopupBalanceResult], error) {
	body := buildBody(map[string]any{
		"amount": params.Amount,
	})
	return parseResponse[TopupBalanceResult](r.http.post(ctx, fmt.Sprintf("/subscriptions/%s/balance/topup", subscriptionID), body, ""))
}

func (r *SubscriptionsResource) PurchaseCredits(ctx context.Context, subscriptionID string, params *PurchaseCreditsParams) (*ApiResponse[PurchaseCreditsResult], error) {
	body := buildBody(map[string]any{
		"credit_pack_id": params.CreditPackID,
	})
	return parseResponse[PurchaseCreditsResult](r.http.post(ctx, fmt.Sprintf("/subscriptions/%s/credits", subscriptionID), body, ""))
}
