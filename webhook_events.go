package commet

import "encoding/json"

type WebhookEventType string

const (
	EventSubscriptionCreated               WebhookEventType = "subscription.created"
	EventSubscriptionActivated             WebhookEventType = "subscription.activated"
	EventSubscriptionReactivated           WebhookEventType = "subscription.reactivated"
	EventSubscriptionCanceled              WebhookEventType = "subscription.canceled"
	EventSubscriptionUpdated               WebhookEventType = "subscription.updated"
	EventSubscriptionPlanChanged           WebhookEventType = "subscription.plan_changed"
	EventSubscriptionCancellationScheduled WebhookEventType = "subscription.cancellation_scheduled"
	EventSubscriptionCancellationRevoked   WebhookEventType = "subscription.cancellation_revoked"
	EventSubscriptionPlanChangeScheduled   WebhookEventType = "subscription.plan_change_scheduled"
	EventSubscriptionPlanChangeRevoked     WebhookEventType = "subscription.plan_change_revoked"
	EventSubscriptionPastDue               WebhookEventType = "subscription.past_due"
	EventTrialStarted                      WebhookEventType = "trial.started"
	EventTrialConverted                    WebhookEventType = "trial.converted"
	EventTrialExpired                      WebhookEventType = "trial.expired"
	EventTrialWillEnd                      WebhookEventType = "trial.will_end"
	EventTrialCheckoutReady                WebhookEventType = "trial.checkout_ready"
	EventCheckoutReady                     WebhookEventType = "checkout.ready"
	EventPaymentReceived                   WebhookEventType = "payment.received"
	EventPaymentFailed                     WebhookEventType = "payment.failed"
	EventPaymentRecovered                  WebhookEventType = "payment.recovered"
	EventPaymentRetryFailed                WebhookEventType = "payment.retry_failed"
	EventPaymentRefunded                   WebhookEventType = "payment.refunded"
	EventPaymentDisputed                   WebhookEventType = "payment.disputed"
	EventPaymentDisputeResolved            WebhookEventType = "payment.dispute_resolved"
	EventPaymentLinkCreated                WebhookEventType = "payment_link.created"
	EventPaymentLinkCompleted              WebhookEventType = "payment_link.completed"
	EventPaymentLinkFailed                 WebhookEventType = "payment_link.failed"
	EventPaymentLinkCanceled               WebhookEventType = "payment_link.canceled"
	EventInvoiceCreated                    WebhookEventType = "invoice.created"
	EventInvoiceVoided                     WebhookEventType = "invoice.voided"
	EventInvoiceOverdue                    WebhookEventType = "invoice.overdue"
	EventInvoiceUpcoming                   WebhookEventType = "invoice.upcoming"
	EventPaymentMethodAttached             WebhookEventType = "payment_method.attached"
	EventPaymentMethodUpdated              WebhookEventType = "payment_method.updated"
	EventCustomerCreated                   WebhookEventType = "customer.created"
	EventCustomerUpdated                   WebhookEventType = "customer.updated"
	EventCustomerStateChanged              WebhookEventType = "customer.state_changed"
	EventPlanGrantCreated                  WebhookEventType = "plan_grant.created"
	EventPlanGrantUpdated                  WebhookEventType = "plan_grant.updated"
	EventPlanGrantExpired                  WebhookEventType = "plan_grant.expired"
	EventPlanGrantRevoked                  WebhookEventType = "plan_grant.revoked"
	EventCreditsGranted                    WebhookEventType = "credits.granted"
	EventCreditsPurchased                  WebhookEventType = "credits.purchased"
	EventCreditsLow                        WebhookEventType = "credits.low"
	EventCreditsDepleted                   WebhookEventType = "credits.depleted"
	EventCreditsExpired                    WebhookEventType = "credits.expired"
	EventBalanceToppedUp                   WebhookEventType = "balance.topped_up"
	EventBalanceLow                        WebhookEventType = "balance.low"
	EventBalanceDepleted                   WebhookEventType = "balance.depleted"
	EventQuotaThresholdReached             WebhookEventType = "quota.threshold_reached"
	EventQuotaExceeded                     WebhookEventType = "quota.exceeded"
	EventSeatsUpdated                      WebhookEventType = "seats.updated"
	EventSeatsLimitReached                 WebhookEventType = "seats.limit_reached"
	EventAddonActivated                    WebhookEventType = "addon.activated"
	EventAddonDeactivated                  WebhookEventType = "addon.deactivated"
	EventUsageRecorded                     WebhookEventType = "usage.recorded"
	EventPayoutAvailable                   WebhookEventType = "payout.available"
	EventPayoutCreated                     WebhookEventType = "payout.created"
	EventPayoutPaid                        WebhookEventType = "payout.paid"
	EventPayoutFailed                      WebhookEventType = "payout.failed"
)

// Fired when a subscription record is created with status pending_payment. The first charge has not been confirmed yet — do NOT grant access here. Wait for subscription.activated.
type SubscriptionCreatedData struct {
	SubscriptionID string  `json:"subscriptionId"`
	CustomerID     string  `json:"customerId"`
	PlanID         string  `json:"planId"`
	PlanName       string  `json:"planName"`
	Status         string  `json:"status"`
	StartDate      string  `json:"startDate"`
	Name           *string `json:"name"`
}

// Fired once, when the subscription's first charge succeeds and it becomes active — this is where you grant access. Never re-fired on renewals; use payment.received for per-charge notifications.
type SubscriptionActivatedData struct {
	SubscriptionID     string  `json:"subscriptionId"`
	CustomerID         string  `json:"customerId"`
	Status             string  `json:"status"`
	CurrentPeriodStart string  `json:"currentPeriodStart"`
	CurrentPeriodEnd   string  `json:"currentPeriodEnd"`
	Name               *string `json:"name"`
	InvoiceID          string  `json:"invoiceId"`
	InvoiceNumber      string  `json:"invoiceNumber"`
	InvoiceTotal       float64 `json:"invoiceTotal"`
	InvoiceCurrency    string  `json:"invoiceCurrency"`
	Provider           *string `json:"provider"`
}

// Fired when a canceled subscription is reactivated and its reactivation charge succeeds. The subscription returns to active with a fresh invoice and a billing period anchored to the reactivation date. Distinct from subscription.activated (first activation) and payment.recovered (past_due recovery, which keeps the original anchor).
type SubscriptionReactivatedData struct {
	SubscriptionID     string  `json:"subscriptionId"`
	CustomerID         string  `json:"customerId"`
	Status             string  `json:"status"`
	CurrentPeriodStart string  `json:"currentPeriodStart"`
	CurrentPeriodEnd   string  `json:"currentPeriodEnd"`
	Name               *string `json:"name"`
	InvoiceID          string  `json:"invoiceId"`
	InvoiceNumber      string  `json:"invoiceNumber"`
	InvoiceTotal       float64 `json:"invoiceTotal"`
	InvoiceCurrency    string  `json:"invoiceCurrency"`
	Provider           string  `json:"provider"`
}

// Fired when a subscription is actually terminated. A scheduled cancellation fires it at the end of the billing period; immediate cancellations, full refunds (cancelReason refund), and exhausted dunning retries (cancelReason dunning_exhausted) fire it right away. The status is now canceled and access should be revoked. This event is NOT fired when cancellation is scheduled — that triggers subscription.updated instead. See the cancellation lifecycle below.
type SubscriptionCanceledData struct {
	SubscriptionID string  `json:"subscriptionId"`
	CustomerID     string  `json:"customerId"`
	Status         string  `json:"status"`
	CanceledAt     string  `json:"canceledAt"`
	CancelReason   *string `json:"cancelReason"`
	EndDate        string  `json:"endDate"`
}

// Fired when subscription details change. The most common trigger is scheduling a cancellation — when a customer cancels, the status stays "active" until the billing period ends, but canceledAt and endDate are set immediately. Use this event to show "your subscription will end on {endDate}" in your UI. Access should NOT be revoked here — wait for subscription.canceled.
type SubscriptionUpdatedData struct {
	SubscriptionID string  `json:"subscriptionId"`
	CustomerID     string  `json:"customerId"`
	Status         string  `json:"status"`
	CanceledAt     string  `json:"canceledAt"`
	CancelReason   *string `json:"cancelReason"`
	EndDate        string  `json:"endDate"`
}

// Fired when a subscription changes from one plan to another, including upgrades, downgrades, and billing interval changes. Access does not change on this event — the subscription stays active.
type SubscriptionPlanChangedData struct {
	SubscriptionID  string         `json:"subscriptionId"`
	CustomerID      string         `json:"customerId"`
	PreviousPlan    WebhookPlanRef `json:"previousPlan"`
	CurrentPlan     WebhookPlanRef `json:"currentPlan"`
	BillingInterval *string        `json:"billingInterval"`
	Credit          float64        `json:"credit"`
	Charge          float64        `json:"charge"`
	TotalCharged    float64        `json:"totalCharged"`
}

// Fired when a cancellation is scheduled for the end of the billing period. The subscription stays active until effectiveAt — do NOT revoke access here. subscription.updated also fires for backward compatibility.
type SubscriptionCancellationScheduledData struct {
	SubscriptionID string  `json:"subscriptionId"`
	CustomerID     string  `json:"customerId"`
	Status         string  `json:"status"`
	CanceledAt     string  `json:"canceledAt"`
	CancelReason   *string `json:"cancelReason"`
	EffectiveAt    string  `json:"effectiveAt"`
}

// Fired when a scheduled cancellation is reverted before it executes. The subscription continues on its current plan and billing period as if it had never been canceled.
type SubscriptionCancellationRevokedData struct {
	SubscriptionID   string `json:"subscriptionId"`
	CustomerID       string `json:"customerId"`
	Status           string `json:"status"`
	CurrentPeriodEnd string `json:"currentPeriodEnd"`
}

// Fired when a plan change (downgrade or shorter interval) is scheduled for the end of the billing period. The subscription stays on the current plan until effectiveAt, when subscription.plan_changed fires.
type SubscriptionPlanChangeScheduledData struct {
	SubscriptionID           string         `json:"subscriptionId"`
	CustomerID               string         `json:"customerId"`
	Status                   string         `json:"status"`
	CurrentPlan              WebhookPlanRef `json:"currentPlan"`
	ScheduledPlan            WebhookPlanRef `json:"scheduledPlan"`
	BillingInterval          *string        `json:"billingInterval"`
	ScheduledBillingInterval *string        `json:"scheduledBillingInterval"`
	EffectiveAt              string         `json:"effectiveAt"`
}

// Fired when a scheduled plan change is replaced by a different one before it executes. The replacement also fires subscription.plan_change_scheduled with the new target plan.
type SubscriptionPlanChangeRevokedData struct {
	SubscriptionID         string         `json:"subscriptionId"`
	CustomerID             string         `json:"customerId"`
	Status                 string         `json:"status"`
	CurrentPlan            WebhookPlanRef `json:"currentPlan"`
	RevokedPlan            WebhookPlanRef `json:"revokedPlan"`
	BillingInterval        *string        `json:"billingInterval"`
	RevokedBillingInterval *string        `json:"revokedBillingInterval"`
}

// Fired when a recurring payment fails on a previously paid subscription and its status becomes past_due. past_due is a grace window, not a cutoff: usage and seats keep working while new purchases are blocked, and dunning retries the charge — use this to notify the customer and recover the payment.
type SubscriptionPastDueData struct {
	SubscriptionID string `json:"subscriptionId"`
	CustomerID     string `json:"customerId"`
	Status         string `json:"status"`
	InvoiceID      string `json:"invoiceId"`
	InvoiceNumber  string `json:"invoiceNumber"`
}

// Fired when a subscription enters its trial period after checkout. Grant access here — trialing subscriptions have full access until trialEndsAt.
type TrialStartedData struct {
	SubscriptionID string `json:"subscriptionId"`
	CustomerID     string `json:"customerId"`
	Status         string `json:"status"`
	PlanID         string `json:"planId"`
	PlanName       string `json:"planName"`
	TrialEndsAt    string `json:"trialEndsAt"`
}

// Fired when a trialing customer converts to a paid subscription before the trial ends — today this happens when they change plan during the trial, which charges the full new plan price immediately. Trials that simply run out fire trial.expired instead.
type TrialConvertedData struct {
	SubscriptionID string `json:"subscriptionId"`
	CustomerID     string `json:"customerId"`
	Status         string `json:"status"`
	PlanID         string `json:"planId"`
	PlanName       string `json:"planName"`
}

// Fired when a trial period runs out and the billing cycle activates the subscription. The first regular invoice is generated right after — this is the natural trial-to-paid transition.
type TrialExpiredData struct {
	SubscriptionID string `json:"subscriptionId"`
	CustomerID     string `json:"customerId"`
	Status         string `json:"status"`
	PlanID         string `json:"planId"`
	PlanName       string `json:"planName"`
	TrialEndsAt    string `json:"trialEndsAt"`
}

// Predictive event fired once, 3 days before a trial ends. Use it to remind the customer that billing starts soon. Emitted by a daily scan with a deterministic idempotency key, so it never fires twice for the same trial end date.
type TrialWillEndData struct {
	SubscriptionID string `json:"subscriptionId"`
	CustomerID     string `json:"customerId"`
	Status         string `json:"status"`
	PlanID         string `json:"planId"`
	PlanName       string `json:"planName"`
	TrialEndsAt    string `json:"trialEndsAt"`
}

// Fired when a trial checkout link is ready to share with the customer. Completing this checkout saves a payment method and starts the trial (trial.started) — the customer is not charged until the trial ends.
type TrialCheckoutReadyData struct {
	SubscriptionID string  `json:"subscriptionId"`
	CustomerID     string  `json:"customerId"`
	PlanName       string  `json:"planName"`
	TrialDays      float64 `json:"trialDays"`
	CheckoutURL    string  `json:"checkoutUrl"`
}

// Fired when a checkout link for a subscription's first invoice is ready to share with the customer. Commet also emails the link — use this event to deliver it through your own channels.
type CheckoutReadyData struct {
	SubscriptionID  string  `json:"subscriptionId"`
	CustomerID      string  `json:"customerId"`
	InvoiceID       string  `json:"invoiceId"`
	InvoiceNumber   string  `json:"invoiceNumber"`
	InvoiceTotal    float64 `json:"invoiceTotal"`
	InvoiceCurrency string  `json:"invoiceCurrency"`
	CheckoutURL     string  `json:"checkoutUrl"`
}

// Fired every time a payment settles successfully — the first payment and every renewal alike. subscription.activated fires alongside it only on the first one.
type PaymentReceivedData struct {
	InvoiceID            string   `json:"invoiceId"`
	InvoiceNumber        string   `json:"invoiceNumber"`
	InvoiceTotal         float64  `json:"invoiceTotal"`
	CustomerID           string   `json:"customerId"`
	SubscriptionID       *string  `json:"subscriptionId"`
	PaymentTransactionID *string  `json:"paymentTransactionId"`
	Provider             *string  `json:"provider"`
	GrossAmount          *float64 `json:"grossAmount"`
	Currency             *string  `json:"currency"`
	OrgNetAmount         *float64 `json:"orgNetAmount"`
	CustomerEmail        *string  `json:"customerEmail"`
	PaidAt               string   `json:"paidAt"`
}

// Fired when a recurring charge fails. This event is for recurring charge failures only — card declines during initial checkout do not trigger this event.
type PaymentFailedData struct {
	InvoiceID      string  `json:"invoiceId"`
	InvoiceNumber  string  `json:"invoiceNumber"`
	CustomerID     string  `json:"customerId"`
	SubscriptionID *string `json:"subscriptionId"`
	Provider       string  `json:"provider"`
	FailureCode    string  `json:"failureCode"`
	FailureMessage string  `json:"failureMessage"`
	RecoveryURL    *string `json:"recoveryUrl"`
}

// Fired when an outstanding invoice that previously failed is successfully paid — automatically on retry or by the customer through the portal. The subscription returns to active at the same time; use this event to close the dunning flow you opened on payment.failed.
type PaymentRecoveredData struct {
	InvoiceID      string  `json:"invoiceId"`
	InvoiceNumber  string  `json:"invoiceNumber"`
	InvoiceTotal   float64 `json:"invoiceTotal"`
	CustomerID     string  `json:"customerId"`
	SubscriptionID *string `json:"subscriptionId"`
	Provider       *string `json:"provider"`
}

// Fired when all dunning retries are exhausted and the subscription is canceled. This is the terminal event of the dunning flow — payment.recovered will not follow. Revoke access when you receive this.
type PaymentRetryFailedData struct {
	InvoiceID      string `json:"invoiceId"`
	InvoiceNumber  string `json:"invoiceNumber"`
	CustomerID     string `json:"customerId"`
	SubscriptionID string `json:"subscriptionId"`
	Provider       string `json:"provider"`
	Reason         string `json:"reason"`
}

// Fired when a payment is refunded, fully or partially. A full refund of a subscription invoice also cancels the subscription immediately (subscription.canceled fires with reason refund); partial refunds leave the subscription untouched.
type PaymentRefundedData struct {
	PaymentTransactionID string  `json:"paymentTransactionId"`
	Provider             string  `json:"provider"`
	PaymentLinkID        *string `json:"paymentLinkId"`
	InvoiceID            *string `json:"invoiceId"`
	InvoiceNumber        *string `json:"invoiceNumber"`
	CustomerID           *string `json:"customerId"`
	SubscriptionID       *string `json:"subscriptionId"`
	RefundAmount         float64 `json:"refundAmount"`
	Currency             string  `json:"currency"`
}

// Fired when a cardholder opens a dispute (chargeback) against a payment. The disputed amount is frozen from your payout balance while the dispute is open; Commet, as the Merchant of Record, handles the resolution process. payment.dispute_resolved fires with the outcome.
type PaymentDisputedData struct {
	PaymentTransactionID string  `json:"paymentTransactionId"`
	Provider             string  `json:"provider"`
	PaymentLinkID        *string `json:"paymentLinkId"`
	InvoiceID            *string `json:"invoiceId"`
	InvoiceNumber        *string `json:"invoiceNumber"`
	CustomerID           *string `json:"customerId"`
	SubscriptionID       *string `json:"subscriptionId"`
	DisputeAmount        float64 `json:"disputeAmount"`
	Currency             string  `json:"currency"`
	DisputeReason        *string `json:"disputeReason"`
}

// Fired when a dispute is closed. Carries the same identifiers as payment.disputed plus the outcome: won restores the frozen amount to your balance, lost keeps the chargeback deducted.
type PaymentDisputeResolvedData struct {
	PaymentTransactionID string  `json:"paymentTransactionId"`
	Provider             string  `json:"provider"`
	PaymentLinkID        *string `json:"paymentLinkId"`
	InvoiceID            *string `json:"invoiceId"`
	InvoiceNumber        *string `json:"invoiceNumber"`
	CustomerID           *string `json:"customerId"`
	SubscriptionID       *string `json:"subscriptionId"`
	DisputeAmount        float64 `json:"disputeAmount"`
	Currency             string  `json:"currency"`
	DisputeReason        *string `json:"disputeReason"`
	Outcome              string  `json:"outcome"`
}

// Fired when a payment link is created. The link is pending — the customer has not paid yet. Do NOT fulfill here; wait for payment_link.completed.
type PaymentLinkCreatedData struct {
	PaymentID   string  `json:"paymentId"`
	Status      string  `json:"status"`
	Amount      float64 `json:"amount"`
	Currency    string  `json:"currency"`
	Description string  `json:"description"`
	CustomerID  *string `json:"customerId"`
}

// Fired when a payment link is paid. The charge settled and a one-time invoice was generated. Fulfill the purchase on this event.
type PaymentLinkCompletedData struct {
	PaymentID            string  `json:"paymentId"`
	Status               string  `json:"status"`
	Amount               float64 `json:"amount"`
	Currency             string  `json:"currency"`
	Description          string  `json:"description"`
	CustomerID           *string `json:"customerId"`
	InvoiceID            string  `json:"invoiceId"`
	InvoiceNumber        string  `json:"invoiceNumber"`
	PaymentTransactionID *string `json:"paymentTransactionId"`
}

// Fired when a payment link charge attempt is declined. The link stays open and can be paid again — a failed link is retryable.
type PaymentLinkFailedData struct {
	PaymentID      string  `json:"paymentId"`
	Status         string  `json:"status"`
	Amount         float64 `json:"amount"`
	Currency       string  `json:"currency"`
	Description    string  `json:"description"`
	CustomerID     *string `json:"customerId"`
	FailureCode    string  `json:"failureCode"`
	FailureMessage string  `json:"failureMessage"`
}

// Fired when a pending payment link is canceled before being paid. A canceled link can no longer be paid.
type PaymentLinkCanceledData struct {
	PaymentID   string  `json:"paymentId"`
	Status      string  `json:"status"`
	Amount      float64 `json:"amount"`
	Currency    string  `json:"currency"`
	Description string  `json:"description"`
	CustomerID  *string `json:"customerId"`
}

// Fired when a new invoice is generated for a subscription, typically at the start of a billing period.
type InvoiceCreatedData struct {
	InvoiceID      string  `json:"invoiceId"`
	InvoiceNumber  string  `json:"invoiceNumber"`
	InvoiceStatus  string  `json:"invoiceStatus"`
	PeriodStart    string  `json:"periodStart"`
	PeriodEnd      string  `json:"periodEnd"`
	IssueDate      string  `json:"issueDate"`
	DueDate        string  `json:"dueDate"`
	Currency       string  `json:"currency"`
	Subtotal       float64 `json:"subtotal"`
	Total          float64 `json:"total"`
	CustomerID     string  `json:"customerId"`
	SubscriptionID *string `json:"subscriptionId"`
}

// Fired when an invoice is voided — nullified before collection, either manually or automatically when its subscription is canceled. Voiding is terminal: a void invoice is never retried or collected.
type InvoiceVoidedData struct {
	InvoiceID      string  `json:"invoiceId"`
	InvoiceNumber  string  `json:"invoiceNumber"`
	InvoiceStatus  string  `json:"invoiceStatus"`
	PeriodStart    string  `json:"periodStart"`
	PeriodEnd      string  `json:"periodEnd"`
	IssueDate      string  `json:"issueDate"`
	DueDate        string  `json:"dueDate"`
	Currency       string  `json:"currency"`
	Subtotal       float64 `json:"subtotal"`
	Total          float64 `json:"total"`
	CustomerID     string  `json:"customerId"`
	SubscriptionID *string `json:"subscriptionId"`
}

// Fired once when an outstanding invoice passes its due date without payment. The invoice keeps its outstanding status — overdue is a fact about the due date, not a new status. Use it to start your own dunning flow.
type InvoiceOverdueData struct {
	InvoiceID      string  `json:"invoiceId"`
	InvoiceNumber  string  `json:"invoiceNumber"`
	InvoiceStatus  string  `json:"invoiceStatus"`
	PeriodStart    string  `json:"periodStart"`
	PeriodEnd      string  `json:"periodEnd"`
	IssueDate      string  `json:"issueDate"`
	DueDate        string  `json:"dueDate"`
	Currency       string  `json:"currency"`
	Subtotal       float64 `json:"subtotal"`
	Total          float64 `json:"total"`
	CustomerID     string  `json:"customerId"`
	SubscriptionID *string `json:"subscriptionId"`
}

// Predictive event fired once, 3 days before an active subscription renews. Use it to notify the customer before they are charged. Carries no amount — usage-based charges are only final at renewal, when invoice.created delivers the actual invoice.
type InvoiceUpcomingData struct {
	SubscriptionID   string  `json:"subscriptionId"`
	CustomerID       string  `json:"customerId"`
	Status           string  `json:"status"`
	PlanID           string  `json:"planId"`
	PlanName         string  `json:"planName"`
	BillingInterval  *string `json:"billingInterval"`
	CurrentPeriodEnd string  `json:"currentPeriodEnd"`
}

// Fired when Commet records a payment method for a subscription: after a paid checkout, when a trial starts with a card on file, or when a zero-total checkout completes. The card object carries display metadata only — full numbers never leave the payment provider.
type PaymentMethodAttachedData struct {
	SubscriptionID string           `json:"subscriptionId"`
	CustomerID     string           `json:"customerId"`
	Card           *WebhookCardInfo `json:"card"`
}

// Fired when a customer replaces their default payment method through the customer portal. The new method applies to all of the customer's subscriptions. A payment method update is also a strong recovery signal for past-due subscriptions.
type PaymentMethodUpdatedData struct {
	CustomerID string           `json:"customerId"`
	Card       *WebhookCardInfo `json:"card"`
}

// Fired when a customer is created, via the API (including batch create), SDK, or dashboard. The payload is the customer resource exactly as GET /customers returns it.
type CustomerCreatedData struct {
	ID           string         `json:"id"`
	ExternalID   *string        `json:"externalId"`
	FullName     *string        `json:"fullName"`
	Email        string         `json:"email"`
	TaxDocument  *string        `json:"taxDocument"`
	DocumentType *string        `json:"documentType"`
	Timezone     *string        `json:"timezone"`
	Metadata     map[string]any `json:"metadata"`
	CreatedAt    string         `json:"createdAt"`
	UpdatedAt    string         `json:"updatedAt"`
}

// Fired when a customer's details change (email, name, timezone, externalId, or metadata). Carries the same customer resource shape as customer.created with the current values.
type CustomerUpdatedData struct {
	ID           string         `json:"id"`
	ExternalID   *string        `json:"externalId"`
	FullName     *string        `json:"fullName"`
	Email        string         `json:"email"`
	TaxDocument  *string        `json:"taxDocument"`
	DocumentType *string        `json:"documentType"`
	Timezone     *string        `json:"timezone"`
	Metadata     map[string]any `json:"metadata"`
	CreatedAt    string         `json:"createdAt"`
	UpdatedAt    string         `json:"updatedAt"`
}

// Aggregate entitlement event answering one question: what can this customer access right now? Fired on every entitlement transition (subscription lifecycle, plan changes, trials, past due, scheduled cancellations) with the customer's CURRENT subscription, plan, features, seats, and credits or balance. Handle this single event to keep access in sync instead of wiring every lifecycle event.
type CustomerStateChangedData struct {
	CustomerID       string                 `json:"customerId"`
	Trigger          string                 `json:"trigger"`
	Status           string                 `json:"status"`
	SubscriptionID   *string                `json:"subscriptionId"`
	Plan             *WebhookPlanRef        `json:"plan"`
	BillingInterval  *string                `json:"billingInterval"`
	ConsumptionModel *string                `json:"consumptionModel"`
	Features         []any                  `json:"features"`
	Seats            []WebhookSeatSummary   `json:"seats"`
	Credits          *WebhookCreditsBalance `json:"credits"`
	Balance          *WebhookBalance        `json:"balance"`
}

// Fired after a plan grant is durably created. The payload is the grant snapshot at creation.
type PlanGrantCreatedData struct {
	ID                  string                          `json:"id"`
	CustomerID          string                          `json:"customerId"`
	SubscriptionID      string                          `json:"subscriptionId"`
	BasePlanID          string                          `json:"basePlanId"`
	TargetPlanID        string                          `json:"targetPlanId"`
	TargetPlanReleaseID string                          `json:"targetPlanReleaseId"`
	Status              string                          `json:"status"`
	Duration            string                          `json:"duration"`
	DurationCycles      *int                            `json:"durationCycles"`
	StartsAt            string                          `json:"startsAt"`
	ExpiresAt           *string                         `json:"expiresAt"`
	Reason              string                          `json:"reason"`
	Source              string                          `json:"source"`
	RevokedAt           *string                         `json:"revokedAt"`
	CreatedAt           string                          `json:"createdAt"`
	UpdatedAt           string                          `json:"updatedAt"`
	Events              []WebhookPlanGrantTimelineEvent `json:"events"`
}

// Fired after a plan grant duration or deadline is durably changed. The payload is the grant snapshot at that update.
type PlanGrantUpdatedData struct {
	ID                  string                          `json:"id"`
	CustomerID          string                          `json:"customerId"`
	SubscriptionID      string                          `json:"subscriptionId"`
	BasePlanID          string                          `json:"basePlanId"`
	TargetPlanID        string                          `json:"targetPlanId"`
	TargetPlanReleaseID string                          `json:"targetPlanReleaseId"`
	Status              string                          `json:"status"`
	Duration            string                          `json:"duration"`
	DurationCycles      *int                            `json:"durationCycles"`
	StartsAt            string                          `json:"startsAt"`
	ExpiresAt           *string                         `json:"expiresAt"`
	Reason              string                          `json:"reason"`
	Source              string                          `json:"source"`
	RevokedAt           *string                         `json:"revokedAt"`
	CreatedAt           string                          `json:"createdAt"`
	UpdatedAt           string                          `json:"updatedAt"`
	Events              []WebhookPlanGrantTimelineEvent `json:"events"`
}

// Fired after a plan grant is durably expired, whether discovered automatically or while replacing an expired grant.
type PlanGrantExpiredData struct {
	ID                  string                          `json:"id"`
	CustomerID          string                          `json:"customerId"`
	SubscriptionID      string                          `json:"subscriptionId"`
	BasePlanID          string                          `json:"basePlanId"`
	TargetPlanID        string                          `json:"targetPlanId"`
	TargetPlanReleaseID string                          `json:"targetPlanReleaseId"`
	Status              string                          `json:"status"`
	Duration            string                          `json:"duration"`
	DurationCycles      *int                            `json:"durationCycles"`
	StartsAt            string                          `json:"startsAt"`
	ExpiresAt           *string                         `json:"expiresAt"`
	Reason              string                          `json:"reason"`
	Source              string                          `json:"source"`
	RevokedAt           *string                         `json:"revokedAt"`
	CreatedAt           string                          `json:"createdAt"`
	UpdatedAt           string                          `json:"updatedAt"`
	Events              []WebhookPlanGrantTimelineEvent `json:"events"`
}

// Fired after an active plan grant is durably revoked. The payload is the grant snapshot at revocation.
type PlanGrantRevokedData struct {
	ID                  string                          `json:"id"`
	CustomerID          string                          `json:"customerId"`
	SubscriptionID      string                          `json:"subscriptionId"`
	BasePlanID          string                          `json:"basePlanId"`
	TargetPlanID        string                          `json:"targetPlanId"`
	TargetPlanReleaseID string                          `json:"targetPlanReleaseId"`
	Status              string                          `json:"status"`
	Duration            string                          `json:"duration"`
	DurationCycles      *int                            `json:"durationCycles"`
	StartsAt            string                          `json:"startsAt"`
	ExpiresAt           *string                         `json:"expiresAt"`
	Reason              string                          `json:"reason"`
	Source              string                          `json:"source"`
	RevokedAt           *string                         `json:"revokedAt"`
	CreatedAt           string                          `json:"createdAt"`
	UpdatedAt           string                          `json:"updatedAt"`
	Events              []WebhookPlanGrantTimelineEvent `json:"events"`
}

// Fired when non-purchase credits are granted to a subscription: plan-included credits at the start of each billing period, or a manual adjustment from the dashboard. Credit pack purchases fire credits.purchased instead.
type CreditsGrantedData struct {
	SubscriptionID string  `json:"subscriptionId"`
	CustomerID     string  `json:"customerId"`
	Credits        float64 `json:"credits"`
	Reason         string  `json:"reason"`
}

// Fired when a customer buys a credit pack through the customer portal and the payment succeeds. Purchased credits never expire — unlike plan credits, they survive period resets. Plan-included credit grants fire credits.granted instead.
type CreditsPurchasedData struct {
	SubscriptionID string  `json:"subscriptionId"`
	CustomerID     string  `json:"customerId"`
	InvoiceID      string  `json:"invoiceId"`
	InvoiceNumber  string  `json:"invoiceNumber"`
	CreditPackName string  `json:"creditPackName"`
	Credits        float64 `json:"credits"`
}

// Fired when a subscription's remaining credits cross below 10% of the credits granted for the current period. Emitted once per billing period, when the crossing happens.
type CreditsLowData struct {
	SubscriptionID   string  `json:"subscriptionId"`
	CustomerID       string  `json:"customerId"`
	RemainingCredits float64 `json:"remainingCredits"`
	ThresholdCredits float64 `json:"thresholdCredits"`
	PeriodCredits    float64 `json:"periodCredits"`
}

// Fired when a subscription's credits hit zero. Usage requests that need more credits than remain are rejected from this point. Also fires customer.state_changed with trigger credits_depleted.
type CreditsDepletedData struct {
	SubscriptionID   string  `json:"subscriptionId"`
	CustomerID       string  `json:"customerId"`
	RemainingCredits float64 `json:"remainingCredits"`
}

// Fired at the period reset when unused plan credits from the previous period are discarded. Plan credits expire at period end; purchased credits never expire and are not affected.
type CreditsExpiredData struct {
	SubscriptionID string  `json:"subscriptionId"`
	CustomerID     string  `json:"customerId"`
	ExpiredCredits float64 `json:"expiredCredits"`
}

// Fired when a customer on a balance plan tops up their prepaid balance through the customer portal and the payment succeeds.
type BalanceToppedUpData struct {
	SubscriptionID string  `json:"subscriptionId"`
	CustomerID     string  `json:"customerId"`
	InvoiceID      string  `json:"invoiceId"`
	InvoiceNumber  string  `json:"invoiceNumber"`
	Amount         float64 `json:"amount"`
	Currency       string  `json:"currency"`
}

// Fired when a subscription's prepaid balance crosses below 10% of its last refill (period reset, top-up, or manual adjustment). Emitted once per crossing.
type BalanceLowData struct {
	SubscriptionID   string  `json:"subscriptionId"`
	CustomerID       string  `json:"customerId"`
	CurrentBalance   float64 `json:"currentBalance"`
	ThresholdBalance float64 `json:"thresholdBalance"`
	Currency         string  `json:"currency"`
}

// Fired when a subscription's prepaid balance crosses to zero or below. With block-on-exhaustion plans further usage is rejected; otherwise the balance can go negative. Also fires customer.state_changed with trigger balance_depleted.
type BalanceDepletedData struct {
	SubscriptionID string  `json:"subscriptionId"`
	CustomerID     string  `json:"customerId"`
	CurrentBalance float64 `json:"currentBalance"`
	Currency       string  `json:"currency"`
}

// Fired when a metered feature's usage crosses 80% of its included quantity for the current period. Emitted once per feature per billing period, when the crossing happens.
type QuotaThresholdReachedData struct {
	SubscriptionID string  `json:"subscriptionId"`
	CustomerID     string  `json:"customerId"`
	FeatureCode    string  `json:"featureCode"`
	CurrentUsage   float64 `json:"currentUsage"`
	IncludedAmount float64 `json:"includedAmount"`
	PeriodStart    string  `json:"periodStart"`
}

// Fired when a metered feature passes its included quantity. With overage enabled it means overage billing began; with overage disabled it means the hard limit was hit and further usage is rejected (this case also fires customer.state_changed with trigger quota_exceeded). Emitted once per feature per billing period.
type QuotaExceededData struct {
	SubscriptionID string  `json:"subscriptionId"`
	CustomerID     string  `json:"customerId"`
	FeatureCode    string  `json:"featureCode"`
	CurrentUsage   float64 `json:"currentUsage"`
	IncludedAmount float64 `json:"includedAmount"`
	OverageEnabled bool    `json:"overageEnabled"`
	PeriodStart    string  `json:"periodStart"`
}

// Fired when a customer's seat count changes for a seats-type feature — via the SDK seats endpoints or the dashboard. Also fires customer.state_changed with trigger seats_updated.
type SeatsUpdatedData struct {
	CustomerID     string  `json:"customerId"`
	SubscriptionID *string `json:"subscriptionId"`
	FeatureCode    string  `json:"featureCode"`
	PreviousSeats  float64 `json:"previousSeats"`
	CurrentSeats   float64 `json:"currentSeats"`
}

// Fired when a seat change reaches or passes the included seat limit of the customer's plan. Emitted once per crossing — only when the count moves from below the limit to at or above it.
type SeatsLimitReachedData struct {
	CustomerID     string  `json:"customerId"`
	SubscriptionID string  `json:"subscriptionId"`
	FeatureCode    string  `json:"featureCode"`
	CurrentSeats   float64 `json:"currentSeats"`
	IncludedSeats  float64 `json:"includedSeats"`
}

// Fired when an add-on is activated on a subscription — via the API or a customer portal purchase. The prorated activation charge, if any, has already succeeded. Also fires customer.state_changed with trigger addon_activated.
type AddonActivatedData struct {
	SubscriptionID string          `json:"subscriptionId"`
	CustomerID     string          `json:"customerId"`
	Addon          WebhookAddonRef `json:"addon"`
	FeatureCode    string          `json:"featureCode"`
	ProratedPrice  float64         `json:"proratedPrice"`
	Currency       string          `json:"currency"`
}

// Fired when an active add-on is deactivated from a subscription. Also fires customer.state_changed with trigger addon_deactivated.
type AddonDeactivatedData struct {
	SubscriptionID string          `json:"subscriptionId"`
	CustomerID     string          `json:"customerId"`
	Addon          WebhookAddonRef `json:"addon"`
	FeatureCode    string          `json:"featureCode"`
}

// Fired for every processed usage event. HIGH VOLUME: this fires once per tracked event, so it is excluded from family select-all in the dashboard — subscribe to it explicitly and make sure your endpoint can absorb your own ingest rate.
type UsageRecordedData struct {
	SubscriptionID string  `json:"subscriptionId"`
	CustomerID     string  `json:"customerId"`
	UsageEventID   string  `json:"usageEventId"`
	FeatureCode    string  `json:"featureCode"`
	Value          float64 `json:"value"`
	Ts             string  `json:"ts"`
}

// Organization-level event about YOUR money as the merchant. Fired when payment funds the provider was holding become available to pay out to your bank.
type PayoutAvailableData struct {
	AvailableAmount float64 `json:"availableAmount"`
	Currency        string  `json:"currency"`
}

// Fired when a payout of your available balance is requested and the transfer toward your bank is initiated. The lifecycle continues with payout.paid or payout.failed.
type PayoutCreatedData struct {
	PayoutID        string          `json:"payoutId"`
	Amount          float64         `json:"amount"`
	Fee             float64         `json:"fee"`
	NetAmount       float64         `json:"netAmount"`
	Currency        string          `json:"currency"`
	Status          string          `json:"status"`
	DestinationBank *WebhookBankRef `json:"destinationBank"`
	CreatedAt       string          `json:"createdAt"`
}

// Fired when the bank settlement of a payout completes — the moment the money actually reaches your bank account, confirmed by the payment provider. Fires exactly once per payout.
type PayoutPaidData struct {
	PayoutID        string          `json:"payoutId"`
	Amount          float64         `json:"amount"`
	Fee             float64         `json:"fee"`
	NetAmount       float64         `json:"netAmount"`
	Currency        string          `json:"currency"`
	Status          string          `json:"status"`
	DestinationBank *WebhookBankRef `json:"destinationBank"`
	PaidAt          *string         `json:"paidAt"`
}

// Fired when the provider reports a payout could not be completed — most commonly a bank rejection (closed account, invalid details). The funds return to your available balance.
type PayoutFailedData struct {
	PayoutID        string          `json:"payoutId"`
	Amount          float64         `json:"amount"`
	Fee             float64         `json:"fee"`
	NetAmount       float64         `json:"netAmount"`
	Currency        string          `json:"currency"`
	Status          string          `json:"status"`
	DestinationBank *WebhookBankRef `json:"destinationBank"`
	FailedAt        *string         `json:"failedAt"`
	FailureCode     *string         `json:"failureCode"`
	FailureMessage  *string         `json:"failureMessage"`
}

type WebhookEvent struct {
	Event          WebhookEventType `json:"event"`
	Timestamp      string           `json:"timestamp"`
	OrganizationID string           `json:"organizationId"`
	Mode           string           `json:"mode"`
	APIVersion     string           `json:"apiVersion"`
	Data           json.RawMessage  `json:"data"`
}

func (e *WebhookEvent) AsSubscriptionCreated() (*SubscriptionCreatedData, error) {
	var d SubscriptionCreatedData
	if err := json.Unmarshal(e.Data, &d); err != nil {
		return nil, err
	}
	return &d, nil
}

func (e *WebhookEvent) AsSubscriptionActivated() (*SubscriptionActivatedData, error) {
	var d SubscriptionActivatedData
	if err := json.Unmarshal(e.Data, &d); err != nil {
		return nil, err
	}
	return &d, nil
}

func (e *WebhookEvent) AsSubscriptionReactivated() (*SubscriptionReactivatedData, error) {
	var d SubscriptionReactivatedData
	if err := json.Unmarshal(e.Data, &d); err != nil {
		return nil, err
	}
	return &d, nil
}

func (e *WebhookEvent) AsSubscriptionCanceled() (*SubscriptionCanceledData, error) {
	var d SubscriptionCanceledData
	if err := json.Unmarshal(e.Data, &d); err != nil {
		return nil, err
	}
	return &d, nil
}

func (e *WebhookEvent) AsSubscriptionUpdated() (*SubscriptionUpdatedData, error) {
	var d SubscriptionUpdatedData
	if err := json.Unmarshal(e.Data, &d); err != nil {
		return nil, err
	}
	return &d, nil
}

func (e *WebhookEvent) AsSubscriptionPlanChanged() (*SubscriptionPlanChangedData, error) {
	var d SubscriptionPlanChangedData
	if err := json.Unmarshal(e.Data, &d); err != nil {
		return nil, err
	}
	return &d, nil
}

func (e *WebhookEvent) AsSubscriptionCancellationScheduled() (*SubscriptionCancellationScheduledData, error) {
	var d SubscriptionCancellationScheduledData
	if err := json.Unmarshal(e.Data, &d); err != nil {
		return nil, err
	}
	return &d, nil
}

func (e *WebhookEvent) AsSubscriptionCancellationRevoked() (*SubscriptionCancellationRevokedData, error) {
	var d SubscriptionCancellationRevokedData
	if err := json.Unmarshal(e.Data, &d); err != nil {
		return nil, err
	}
	return &d, nil
}

func (e *WebhookEvent) AsSubscriptionPlanChangeScheduled() (*SubscriptionPlanChangeScheduledData, error) {
	var d SubscriptionPlanChangeScheduledData
	if err := json.Unmarshal(e.Data, &d); err != nil {
		return nil, err
	}
	return &d, nil
}

func (e *WebhookEvent) AsSubscriptionPlanChangeRevoked() (*SubscriptionPlanChangeRevokedData, error) {
	var d SubscriptionPlanChangeRevokedData
	if err := json.Unmarshal(e.Data, &d); err != nil {
		return nil, err
	}
	return &d, nil
}

func (e *WebhookEvent) AsSubscriptionPastDue() (*SubscriptionPastDueData, error) {
	var d SubscriptionPastDueData
	if err := json.Unmarshal(e.Data, &d); err != nil {
		return nil, err
	}
	return &d, nil
}

func (e *WebhookEvent) AsTrialStarted() (*TrialStartedData, error) {
	var d TrialStartedData
	if err := json.Unmarshal(e.Data, &d); err != nil {
		return nil, err
	}
	return &d, nil
}

func (e *WebhookEvent) AsTrialConverted() (*TrialConvertedData, error) {
	var d TrialConvertedData
	if err := json.Unmarshal(e.Data, &d); err != nil {
		return nil, err
	}
	return &d, nil
}

func (e *WebhookEvent) AsTrialExpired() (*TrialExpiredData, error) {
	var d TrialExpiredData
	if err := json.Unmarshal(e.Data, &d); err != nil {
		return nil, err
	}
	return &d, nil
}

func (e *WebhookEvent) AsTrialWillEnd() (*TrialWillEndData, error) {
	var d TrialWillEndData
	if err := json.Unmarshal(e.Data, &d); err != nil {
		return nil, err
	}
	return &d, nil
}

func (e *WebhookEvent) AsTrialCheckoutReady() (*TrialCheckoutReadyData, error) {
	var d TrialCheckoutReadyData
	if err := json.Unmarshal(e.Data, &d); err != nil {
		return nil, err
	}
	return &d, nil
}

func (e *WebhookEvent) AsCheckoutReady() (*CheckoutReadyData, error) {
	var d CheckoutReadyData
	if err := json.Unmarshal(e.Data, &d); err != nil {
		return nil, err
	}
	return &d, nil
}

func (e *WebhookEvent) AsPaymentReceived() (*PaymentReceivedData, error) {
	var d PaymentReceivedData
	if err := json.Unmarshal(e.Data, &d); err != nil {
		return nil, err
	}
	return &d, nil
}

func (e *WebhookEvent) AsPaymentFailed() (*PaymentFailedData, error) {
	var d PaymentFailedData
	if err := json.Unmarshal(e.Data, &d); err != nil {
		return nil, err
	}
	return &d, nil
}

func (e *WebhookEvent) AsPaymentRecovered() (*PaymentRecoveredData, error) {
	var d PaymentRecoveredData
	if err := json.Unmarshal(e.Data, &d); err != nil {
		return nil, err
	}
	return &d, nil
}

func (e *WebhookEvent) AsPaymentRetryFailed() (*PaymentRetryFailedData, error) {
	var d PaymentRetryFailedData
	if err := json.Unmarshal(e.Data, &d); err != nil {
		return nil, err
	}
	return &d, nil
}

func (e *WebhookEvent) AsPaymentRefunded() (*PaymentRefundedData, error) {
	var d PaymentRefundedData
	if err := json.Unmarshal(e.Data, &d); err != nil {
		return nil, err
	}
	return &d, nil
}

func (e *WebhookEvent) AsPaymentDisputed() (*PaymentDisputedData, error) {
	var d PaymentDisputedData
	if err := json.Unmarshal(e.Data, &d); err != nil {
		return nil, err
	}
	return &d, nil
}

func (e *WebhookEvent) AsPaymentDisputeResolved() (*PaymentDisputeResolvedData, error) {
	var d PaymentDisputeResolvedData
	if err := json.Unmarshal(e.Data, &d); err != nil {
		return nil, err
	}
	return &d, nil
}

func (e *WebhookEvent) AsPaymentLinkCreated() (*PaymentLinkCreatedData, error) {
	var d PaymentLinkCreatedData
	if err := json.Unmarshal(e.Data, &d); err != nil {
		return nil, err
	}
	return &d, nil
}

func (e *WebhookEvent) AsPaymentLinkCompleted() (*PaymentLinkCompletedData, error) {
	var d PaymentLinkCompletedData
	if err := json.Unmarshal(e.Data, &d); err != nil {
		return nil, err
	}
	return &d, nil
}

func (e *WebhookEvent) AsPaymentLinkFailed() (*PaymentLinkFailedData, error) {
	var d PaymentLinkFailedData
	if err := json.Unmarshal(e.Data, &d); err != nil {
		return nil, err
	}
	return &d, nil
}

func (e *WebhookEvent) AsPaymentLinkCanceled() (*PaymentLinkCanceledData, error) {
	var d PaymentLinkCanceledData
	if err := json.Unmarshal(e.Data, &d); err != nil {
		return nil, err
	}
	return &d, nil
}

func (e *WebhookEvent) AsInvoiceCreated() (*InvoiceCreatedData, error) {
	var d InvoiceCreatedData
	if err := json.Unmarshal(e.Data, &d); err != nil {
		return nil, err
	}
	return &d, nil
}

func (e *WebhookEvent) AsInvoiceVoided() (*InvoiceVoidedData, error) {
	var d InvoiceVoidedData
	if err := json.Unmarshal(e.Data, &d); err != nil {
		return nil, err
	}
	return &d, nil
}

func (e *WebhookEvent) AsInvoiceOverdue() (*InvoiceOverdueData, error) {
	var d InvoiceOverdueData
	if err := json.Unmarshal(e.Data, &d); err != nil {
		return nil, err
	}
	return &d, nil
}

func (e *WebhookEvent) AsInvoiceUpcoming() (*InvoiceUpcomingData, error) {
	var d InvoiceUpcomingData
	if err := json.Unmarshal(e.Data, &d); err != nil {
		return nil, err
	}
	return &d, nil
}

func (e *WebhookEvent) AsPaymentMethodAttached() (*PaymentMethodAttachedData, error) {
	var d PaymentMethodAttachedData
	if err := json.Unmarshal(e.Data, &d); err != nil {
		return nil, err
	}
	return &d, nil
}

func (e *WebhookEvent) AsPaymentMethodUpdated() (*PaymentMethodUpdatedData, error) {
	var d PaymentMethodUpdatedData
	if err := json.Unmarshal(e.Data, &d); err != nil {
		return nil, err
	}
	return &d, nil
}

func (e *WebhookEvent) AsCustomerCreated() (*CustomerCreatedData, error) {
	var d CustomerCreatedData
	if err := json.Unmarshal(e.Data, &d); err != nil {
		return nil, err
	}
	return &d, nil
}

func (e *WebhookEvent) AsCustomerUpdated() (*CustomerUpdatedData, error) {
	var d CustomerUpdatedData
	if err := json.Unmarshal(e.Data, &d); err != nil {
		return nil, err
	}
	return &d, nil
}

func (e *WebhookEvent) AsCustomerStateChanged() (*CustomerStateChangedData, error) {
	var d CustomerStateChangedData
	if err := json.Unmarshal(e.Data, &d); err != nil {
		return nil, err
	}
	return &d, nil
}

func (e *WebhookEvent) AsPlanGrantCreated() (*PlanGrantCreatedData, error) {
	var d PlanGrantCreatedData
	if err := json.Unmarshal(e.Data, &d); err != nil {
		return nil, err
	}
	return &d, nil
}

func (e *WebhookEvent) AsPlanGrantUpdated() (*PlanGrantUpdatedData, error) {
	var d PlanGrantUpdatedData
	if err := json.Unmarshal(e.Data, &d); err != nil {
		return nil, err
	}
	return &d, nil
}

func (e *WebhookEvent) AsPlanGrantExpired() (*PlanGrantExpiredData, error) {
	var d PlanGrantExpiredData
	if err := json.Unmarshal(e.Data, &d); err != nil {
		return nil, err
	}
	return &d, nil
}

func (e *WebhookEvent) AsPlanGrantRevoked() (*PlanGrantRevokedData, error) {
	var d PlanGrantRevokedData
	if err := json.Unmarshal(e.Data, &d); err != nil {
		return nil, err
	}
	return &d, nil
}

func (e *WebhookEvent) AsCreditsGranted() (*CreditsGrantedData, error) {
	var d CreditsGrantedData
	if err := json.Unmarshal(e.Data, &d); err != nil {
		return nil, err
	}
	return &d, nil
}

func (e *WebhookEvent) AsCreditsPurchased() (*CreditsPurchasedData, error) {
	var d CreditsPurchasedData
	if err := json.Unmarshal(e.Data, &d); err != nil {
		return nil, err
	}
	return &d, nil
}

func (e *WebhookEvent) AsCreditsLow() (*CreditsLowData, error) {
	var d CreditsLowData
	if err := json.Unmarshal(e.Data, &d); err != nil {
		return nil, err
	}
	return &d, nil
}

func (e *WebhookEvent) AsCreditsDepleted() (*CreditsDepletedData, error) {
	var d CreditsDepletedData
	if err := json.Unmarshal(e.Data, &d); err != nil {
		return nil, err
	}
	return &d, nil
}

func (e *WebhookEvent) AsCreditsExpired() (*CreditsExpiredData, error) {
	var d CreditsExpiredData
	if err := json.Unmarshal(e.Data, &d); err != nil {
		return nil, err
	}
	return &d, nil
}

func (e *WebhookEvent) AsBalanceToppedUp() (*BalanceToppedUpData, error) {
	var d BalanceToppedUpData
	if err := json.Unmarshal(e.Data, &d); err != nil {
		return nil, err
	}
	return &d, nil
}

func (e *WebhookEvent) AsBalanceLow() (*BalanceLowData, error) {
	var d BalanceLowData
	if err := json.Unmarshal(e.Data, &d); err != nil {
		return nil, err
	}
	return &d, nil
}

func (e *WebhookEvent) AsBalanceDepleted() (*BalanceDepletedData, error) {
	var d BalanceDepletedData
	if err := json.Unmarshal(e.Data, &d); err != nil {
		return nil, err
	}
	return &d, nil
}

func (e *WebhookEvent) AsQuotaThresholdReached() (*QuotaThresholdReachedData, error) {
	var d QuotaThresholdReachedData
	if err := json.Unmarshal(e.Data, &d); err != nil {
		return nil, err
	}
	return &d, nil
}

func (e *WebhookEvent) AsQuotaExceeded() (*QuotaExceededData, error) {
	var d QuotaExceededData
	if err := json.Unmarshal(e.Data, &d); err != nil {
		return nil, err
	}
	return &d, nil
}

func (e *WebhookEvent) AsSeatsUpdated() (*SeatsUpdatedData, error) {
	var d SeatsUpdatedData
	if err := json.Unmarshal(e.Data, &d); err != nil {
		return nil, err
	}
	return &d, nil
}

func (e *WebhookEvent) AsSeatsLimitReached() (*SeatsLimitReachedData, error) {
	var d SeatsLimitReachedData
	if err := json.Unmarshal(e.Data, &d); err != nil {
		return nil, err
	}
	return &d, nil
}

func (e *WebhookEvent) AsAddonActivated() (*AddonActivatedData, error) {
	var d AddonActivatedData
	if err := json.Unmarshal(e.Data, &d); err != nil {
		return nil, err
	}
	return &d, nil
}

func (e *WebhookEvent) AsAddonDeactivated() (*AddonDeactivatedData, error) {
	var d AddonDeactivatedData
	if err := json.Unmarshal(e.Data, &d); err != nil {
		return nil, err
	}
	return &d, nil
}

func (e *WebhookEvent) AsUsageRecorded() (*UsageRecordedData, error) {
	var d UsageRecordedData
	if err := json.Unmarshal(e.Data, &d); err != nil {
		return nil, err
	}
	return &d, nil
}

func (e *WebhookEvent) AsPayoutAvailable() (*PayoutAvailableData, error) {
	var d PayoutAvailableData
	if err := json.Unmarshal(e.Data, &d); err != nil {
		return nil, err
	}
	return &d, nil
}

func (e *WebhookEvent) AsPayoutCreated() (*PayoutCreatedData, error) {
	var d PayoutCreatedData
	if err := json.Unmarshal(e.Data, &d); err != nil {
		return nil, err
	}
	return &d, nil
}

func (e *WebhookEvent) AsPayoutPaid() (*PayoutPaidData, error) {
	var d PayoutPaidData
	if err := json.Unmarshal(e.Data, &d); err != nil {
		return nil, err
	}
	return &d, nil
}

func (e *WebhookEvent) AsPayoutFailed() (*PayoutFailedData, error) {
	var d PayoutFailedData
	if err := json.Unmarshal(e.Data, &d); err != nil {
		return nil, err
	}
	return &d, nil
}
