package commet

type FeatureType string

const (
	FeatureTypeBoolean FeatureType = "boolean"
	FeatureTypeUsage   FeatureType = "usage"
	FeatureTypeSeats   FeatureType = "seats"
	FeatureTypeQuota   FeatureType = "quota"
)

type BillingInterval string

const (
	BillingIntervalWeekly    BillingInterval = "weekly"
	BillingIntervalMonthly   BillingInterval = "monthly"
	BillingIntervalQuarterly BillingInterval = "quarterly"
	BillingIntervalYearly    BillingInterval = "yearly"
	BillingIntervalOneTime   BillingInterval = "one_time"
)

type SubscriptionStatus string

const (
	SubscriptionStatusDraft          SubscriptionStatus = "draft"
	SubscriptionStatusPendingPayment SubscriptionStatus = "pending_payment"
	SubscriptionStatusTrialing       SubscriptionStatus = "trialing"
	SubscriptionStatusActive         SubscriptionStatus = "active"
	SubscriptionStatusPaused         SubscriptionStatus = "paused"
	SubscriptionStatusPastDue        SubscriptionStatus = "past_due"
	SubscriptionStatusCanceled       SubscriptionStatus = "canceled"
	SubscriptionStatusExpired        SubscriptionStatus = "expired"
)

type ConsumptionModel string

const (
	ConsumptionModelMetered ConsumptionModel = "metered"
	ConsumptionModelCredits ConsumptionModel = "credits"
	ConsumptionModelBalance ConsumptionModel = "balance"
)

type DiscountType string

const (
	DiscountTypePercentage DiscountType = "percentage"
	DiscountTypeAmount     DiscountType = "amount"
)

type SeatEventType string

const (
	SeatEventTypeAdd    SeatEventType = "add"
	SeatEventTypeRemove SeatEventType = "remove"
	SeatEventTypeSet    SeatEventType = "set"
)

type OverageModel string

const (
	OverageModelPerUnit OverageModel = "per_unit"
)

type Currency string

const (
	CurrencyUSD Currency = "USD"
	CurrencyEUR Currency = "EUR"
	CurrencyGBP Currency = "GBP"
	CurrencyCAD Currency = "CAD"
	CurrencyAUD Currency = "AUD"
	CurrencyJPY Currency = "JPY"
	CurrencyARS Currency = "ARS"
	CurrencyBRL Currency = "BRL"
	CurrencyMXN Currency = "MXN"
	CurrencyCLP Currency = "CLP"
)

type Customer struct {
	ID             string         `json:"id"`
	Object         string         `json:"object"`
	Livemode       bool           `json:"livemode"`
	OrganizationID string         `json:"organization_id"`
	FullName       string         `json:"full_name,omitempty"`
	Domain         string         `json:"domain,omitempty"`
	Website        string         `json:"website,omitempty"`
	BillingEmail   string         `json:"billing_email"`
	Timezone       string         `json:"timezone,omitempty"`
	Language       string         `json:"language,omitempty"`
	Industry       string         `json:"industry,omitempty"`
	EmployeeCount  string         `json:"employee_count,omitempty"`
	Metadata       map[string]any `json:"metadata,omitempty"`
	CreatedAt      string         `json:"created_at"`
	UpdatedAt      string         `json:"updated_at"`
}

type BatchResult struct {
	Successful []Customer        `json:"successful"`
	Failed     []BatchFailure    `json:"failed"`
}

type BatchFailure struct {
	Index int    `json:"index"`
	Error string `json:"error"`
	Data  any    `json:"data"`
}

type Subscription struct {
	ID                 string             `json:"id"`
	Object             string             `json:"object"`
	Livemode           bool               `json:"livemode"`
	CustomerID         string             `json:"customer_id"`
	PlanID             string             `json:"plan_id"`
	PlanName           string             `json:"plan_name"`
	Name               string             `json:"name"`
	Description        string             `json:"description,omitempty"`
	Status             SubscriptionStatus `json:"status"`
	BillingInterval    BillingInterval    `json:"billing_interval,omitempty"`
	TrialEndsAt        string             `json:"trial_ends_at,omitempty"`
	StartDate          string             `json:"start_date"`
	EndDate            string             `json:"end_date,omitempty"`
	CurrentPeriodStart string             `json:"current_period_start,omitempty"`
	CurrentPeriodEnd   string             `json:"current_period_end,omitempty"`
	BillingDayOfMonth  int                `json:"billing_day_of_month"`
	CheckoutURL        string             `json:"checkout_url,omitempty"`
	CreatedAt          string             `json:"created_at"`
	UpdatedAt          string             `json:"updated_at"`
}

type CreatedSubscription struct {
	ID                      string             `json:"id"`
	Object                  string             `json:"object"`
	Livemode                bool               `json:"livemode"`
	CustomerID              string             `json:"customer_id"`
	PlanID                  string             `json:"plan_id"`
	PlanName                string             `json:"plan_name"`
	Name                    string             `json:"name"`
	Status                  SubscriptionStatus `json:"status"`
	BillingInterval         BillingInterval    `json:"billing_interval,omitempty"`
	TrialEndsAt             string             `json:"trial_ends_at,omitempty"`
	StartDate               string             `json:"start_date"`
	EndDate                 string             `json:"end_date,omitempty"`
	CurrentPeriodStart      string             `json:"current_period_start,omitempty"`
	CurrentPeriodEnd        string             `json:"current_period_end,omitempty"`
	BillingDayOfMonth       int                `json:"billing_day_of_month"`
	CheckoutURL             string             `json:"checkout_url,omitempty"`
	CreatedAt               string             `json:"created_at"`
	UpdatedAt               string             `json:"updated_at"`
	IntroOfferEndsAt        string             `json:"intro_offer_ends_at,omitempty"`
	IntroOfferDiscountType  DiscountType       `json:"intro_offer_discount_type,omitempty"`
	IntroOfferDiscountValue *float64           `json:"intro_offer_discount_value,omitempty"`
}

type ActiveSubscription struct {
	ID                string                `json:"id"`
	Object            string                `json:"object"`
	Livemode          bool                  `json:"livemode"`
	CustomerID        string                `json:"customer_id"`
	Plan              SubscriptionPlan      `json:"plan"`
	Name              string                `json:"name"`
	Description       string                `json:"description,omitempty"`
	Status            SubscriptionStatus    `json:"status"`
	ConsumptionModel  ConsumptionModel      `json:"consumption_model"`
	TrialEndsAt       string                `json:"trial_ends_at,omitempty"`
	CurrentPeriod     SubscriptionPeriod    `json:"current_period"`
	Features          []FeatureSummary      `json:"features"`
	Credits           *CreditsSummary       `json:"credits,omitempty"`
	Balance           *BalanceSummary       `json:"balance,omitempty"`
	Cancellation      *CancellationSummary  `json:"cancellation,omitempty"`
	Discount          *DiscountSummary      `json:"discount,omitempty"`
	StartDate         string                `json:"start_date"`
	EndDate           string                `json:"end_date,omitempty"`
	BillingDayOfMonth int                   `json:"billing_day_of_month"`
	NextBillingDate   string                `json:"next_billing_date"`
	CheckoutURL       string                `json:"checkout_url,omitempty"`
	CreatedAt         string                `json:"created_at"`
	UpdatedAt         string                `json:"updated_at"`
}

type CreditsSummary struct {
	Remaining int `json:"remaining"`
	Included  int `json:"included"`
	Purchased int `json:"purchased"`
}

type BalanceSummary struct {
	Remaining int      `json:"remaining"`
	Included  int      `json:"included"`
	Currency  Currency `json:"currency"`
}

type SubscriptionPlan struct {
	ID              string          `json:"id"`
	Name            string          `json:"name"`
	BasePrice       float64         `json:"base_price"`
	BillingInterval BillingInterval `json:"billing_interval,omitempty"`
}

type SubscriptionPeriod struct {
	Start         string `json:"start"`
	End           string `json:"end"`
	DaysRemaining int    `json:"days_remaining"`
}

type FeatureSummary struct {
	Code    string        `json:"code"`
	Name    string        `json:"name"`
	Type    FeatureType   `json:"type"`
	Enabled *bool         `json:"enabled,omitempty"`
	Usage   *FeatureUsage `json:"usage,omitempty"`
}

type FeatureUsage struct {
	Current          int      `json:"current"`
	Included         int      `json:"included"`
	Overage          int      `json:"overage"`
	OverageUnitPrice *float64 `json:"overage_unit_price,omitempty"`
}

type Plan struct {
	ID          string        `json:"id"`
	Object      string        `json:"object"`
	Livemode    bool          `json:"livemode"`
	Code        string        `json:"code"`
	Name        string        `json:"name"`
	Description string        `json:"description,omitempty"`
	IsPublic    bool          `json:"is_public"`
	IsFree      bool          `json:"is_free"`
	IsDefault   bool          `json:"is_default"`
	SortOrder   int           `json:"sort_order"`
	Prices      []PlanPrice   `json:"prices"`
	Features    []PlanFeature `json:"features"`
	CreatedAt   string        `json:"created_at"`
}

type PlanPrice struct {
	BillingInterval BillingInterval `json:"billing_interval"`
	Price           int             `json:"price"`
	IsDefault       bool            `json:"is_default"`
	TrialDays       int             `json:"trial_days"`
}

type PlanFeature struct {
	Code             string      `json:"code"`
	Name             string      `json:"name"`
	Type             FeatureType `json:"type"`
	UnitName         string      `json:"unit_name,omitempty"`
	Enabled          *bool       `json:"enabled,omitempty"`
	IncludedAmount   *int        `json:"included_amount,omitempty"`
	Unlimited        *bool       `json:"unlimited,omitempty"`
	OverageEnabled   *bool       `json:"overage_enabled,omitempty"`
	OverageUnitPrice *float64    `json:"overage_unit_price,omitempty"`
}

type PlanDetail struct {
	ID          string              `json:"id"`
	Object      string              `json:"object"`
	Livemode    bool                `json:"livemode"`
	Code        string              `json:"code"`
	Name        string              `json:"name"`
	Description string              `json:"description,omitempty"`
	IsPublic    bool                `json:"is_public"`
	IsDefault   bool                `json:"is_default"`
	SortOrder   int                 `json:"sort_order"`
	Prices      []PlanDetailPrice   `json:"prices"`
	Features    []PlanDetailFeature `json:"features"`
	CreatedAt   string              `json:"created_at"`
	UpdatedAt   string              `json:"updated_at"`
}

type PlanDetailPrice struct {
	BillingInterval BillingInterval `json:"billing_interval"`
	Price           int             `json:"price"`
	IsDefault       bool            `json:"is_default"`
	TrialDays       int             `json:"trial_days"`
	IntroOffer      *PlanIntroOffer `json:"intro_offer,omitempty"`
}

type PlanIntroOffer struct {
	Enabled        bool         `json:"enabled"`
	DiscountType   DiscountType `json:"discount_type,omitempty"`
	DiscountValue  *float64     `json:"discount_value,omitempty"`
	DurationCycles *int     `json:"duration_cycles,omitempty"`
}

type PlanDetailFeature struct {
	Code             string         `json:"code"`
	Name             string         `json:"name"`
	Type             FeatureType    `json:"type"`
	UnitName         string         `json:"unit_name,omitempty"`
	Enabled          *bool          `json:"enabled,omitempty"`
	IncludedAmount   *int           `json:"included_amount,omitempty"`
	Unlimited        *bool          `json:"unlimited,omitempty"`
	OverageEnabled   *bool          `json:"overage_enabled,omitempty"`
	OverageUnitPrice *float64       `json:"overage_unit_price,omitempty"`
	Overage          *OverageConfig `json:"overage,omitempty"`
}

type OverageConfig struct {
	Enabled   bool         `json:"enabled"`
	Model     OverageModel `json:"model,omitempty"`
	UnitPrice *float64     `json:"unit_price,omitempty"`
}

type FeatureAccess struct {
	Object           string      `json:"object"`
	Livemode         bool        `json:"livemode"`
	Code             string      `json:"code"`
	Name             string      `json:"name"`
	Type             FeatureType `json:"type"`
	Allowed          bool        `json:"allowed"`
	Enabled          *bool    `json:"enabled,omitempty"`
	Current          *int     `json:"current,omitempty"`
	Included         *int     `json:"included,omitempty"`
	Remaining        *int     `json:"remaining,omitempty"`
	Overage          *int     `json:"overage,omitempty"`
	OverageUnitPrice *float64 `json:"overage_unit_price,omitempty"`
	Unlimited        *bool    `json:"unlimited,omitempty"`
	OverageEnabled   *bool    `json:"overage_enabled,omitempty"`
}

type CanUseResult struct {
	Allowed       bool   `json:"allowed"`
	WillBeCharged bool   `json:"will_be_charged"`
	Reason        string `json:"reason,omitempty"`
}

type SeatEvent struct {
	ID              string        `json:"id"`
	Object          string        `json:"object"`
	Livemode        bool          `json:"livemode"`
	OrganizationID  string        `json:"organization_id"`
	CustomerID      string        `json:"customer_id"`
	FeatureCode     string        `json:"feature_code"`
	EventType       SeatEventType `json:"event_type"`
	Quantity        int           `json:"quantity"`
	PreviousBalance *int          `json:"previous_balance,omitempty"`
	NewBalance      int           `json:"new_balance"`
	Ts              string        `json:"ts"`
	CreatedAt       string        `json:"created_at"`
}

type SeatBalance struct {
	Current int    `json:"current"`
	AsOf    string `json:"as_of"`
}

type QuotaEvent struct {
	ID              string `json:"id"`
	CustomerID      string `json:"customerId"`
	FeatureCode     string `json:"featureCode"`
	PreviousBalance int    `json:"previousBalance"`
	NewBalance      int    `json:"newBalance"`
	Ts              string `json:"ts"`
	CreatedAt       string `json:"createdAt"`
}

type QuotaAllowance struct {
	FeatureCode    string  `json:"featureCode"`
	Current        int     `json:"current"`
	Included       int     `json:"included"`
	Remaining      *int    `json:"remaining"`
	BilledQuantity *int    `json:"billedQuantity,omitempty"`
	Unlimited      bool    `json:"unlimited"`
	OverageEnabled bool    `json:"overageEnabled"`
	AsOf           *string `json:"asOf"`
}

type UsageEvent struct {
	ID              string                `json:"id"`
	Object          string                `json:"object"`
	Livemode        bool                  `json:"livemode"`
	OrganizationID  string                `json:"organization_id"`
	CustomerID      string                `json:"customer_id"`
	Feature         string                `json:"feature"`
	IdempotencyKey  string                `json:"idempotency_key,omitempty"`
	Ts              string                `json:"ts"`
	Properties      []UsageEventProperty  `json:"properties,omitempty"`
	CreatedAt       string                `json:"created_at"`
}

type UsageEventProperty struct {
	ID           string `json:"id"`
	UsageEventID string `json:"usage_event_id"`
	Property     string `json:"property"`
	Value        string `json:"value"`
	CreatedAt    string `json:"created_at"`
}

type PortalSession struct {
	Success   bool   `json:"success"`
	Message   string `json:"message"`
	PortalURL string `json:"portal_url"`
}

type CreditPack struct {
	ID          string   `json:"id"`
	Object      string   `json:"object"`
	Livemode    bool     `json:"livemode"`
	Name        string   `json:"name"`
	Description string   `json:"description,omitempty"`
	Credits     int      `json:"credits"`
	Price       int      `json:"price"`
	Currency    Currency `json:"currency"`
}

type CreditPackDetail struct {
	ID          string `json:"id"`
	Object      string `json:"object"`
	Livemode    bool   `json:"livemode"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Credits     int    `json:"credits"`
	Price       int    `json:"price"`
	IsActive    bool   `json:"is_active"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type CancellationSummary struct {
	ScheduledAt string `json:"scheduled_at"`
	Reason      string `json:"reason"`
	EffectiveAt string `json:"effective_at"`
}

type DiscountSummary struct {
	Type  string  `json:"type"`
	Value float64 `json:"value"`
	Name  string  `json:"name"`
	EndsAt string `json:"ends_at"`
}

type SubscriptionListItem struct {
	ID                string             `json:"id"`
	Object            string             `json:"object"`
	Livemode          bool               `json:"livemode"`
	CustomerID        string             `json:"customer_id"`
	PlanID            string             `json:"plan_id"`
	PlanName          string             `json:"plan_name"`
	Name              string             `json:"name"`
	Status            SubscriptionStatus `json:"status"`
	StartDate         string             `json:"start_date"`
	EndDate           string             `json:"end_date,omitempty"`
	BillingDayOfMonth int                `json:"billing_day_of_month"`
	CreatedAt         string             `json:"created_at"`
	UpdatedAt         string             `json:"updated_at"`
}

type PreviewChangeResult struct {
	CurrentPlanCredit float64 `json:"current_plan_credit"`
	NewPlanCharge     float64 `json:"new_plan_charge"`
	EstimatedTotal    float64 `json:"estimated_total"`
	EffectiveDate     string  `json:"effective_date"`
	DaysRemaining     int     `json:"days_remaining"`
	TotalDays         int     `json:"total_days"`
	IsUpgrade         bool    `json:"is_upgrade"`
}

type ActivateAddonResult struct {
	AddonID        string  `json:"addon_id"`
	Status         string  `json:"status"`
	ProratedCharge float64 `json:"prorated_charge"`
}

type DeactivateAddonResult struct {
	ID            string `json:"id"`
	Status        string `json:"status"`
	DeactivatedAt string `json:"deactivated_at"`
}

type AdjustBalanceResult struct {
	Amount     float64 `json:"amount"`
	NewBalance float64 `json:"new_balance"`
	Reason     string  `json:"reason"`
}

type TopupBalanceResult struct {
	Amount float64 `json:"amount"`
}

type PurchaseCreditsResult struct {
	Credits int `json:"credits"`
}

type ApiKey struct {
	ID         string `json:"id"`
	Object     string `json:"object"`
	Livemode   bool   `json:"livemode"`
	Name       string `json:"name"`
	Prefix     string `json:"prefix"`
	ExpiresAt  string `json:"expires_at,omitempty"`
	LastUsedAt string `json:"last_used_at,omitempty"`
	CreatedAt  string `json:"created_at"`
}

type ApiKeyCreated struct {
	ApiKey
	ApiKeyValue string `json:"api_key"`
}

type InvoiceLineItem struct {
	LineType        string  `json:"line_type"`
	FeatureName     string  `json:"feature_name"`
	Description     string  `json:"description"`
	Quantity        int     `json:"quantity"`
	UnitAmount      float64 `json:"unit_amount"`
	Amount          float64 `json:"amount"`
	IncludedAmount  *int    `json:"included_amount,omitempty"`
	UsedAmount      *int    `json:"used_amount,omitempty"`
	OverageAmount   *int    `json:"overage_amount,omitempty"`
	DiscountType    string  `json:"discount_type,omitempty"`
	DiscountValue   *float64 `json:"discount_value,omitempty"`
	DiscountName    string  `json:"discount_name,omitempty"`
	ChargeType      string  `json:"charge_type,omitempty"`
}

type InvoiceListItem struct {
	ID              string         `json:"id"`
	Object          string         `json:"object"`
	Livemode        bool           `json:"livemode"`
	CustomerID      string         `json:"customer_id"`
	SubscriptionID  string         `json:"subscription_id,omitempty"`
	InvoiceNumber   string         `json:"invoice_number"`
	Status          string         `json:"status"`
	InvoiceType     string         `json:"invoice_type"`
	Currency        string         `json:"currency"`
	Subtotal        float64        `json:"subtotal"`
	DiscountAmount  float64        `json:"discount_amount"`
	TaxAmount       float64        `json:"tax_amount"`
	Total           float64        `json:"total"`
	PeriodStart     string         `json:"period_start,omitempty"`
	PeriodEnd       string         `json:"period_end,omitempty"`
	IssueDate       string         `json:"issue_date"`
	DueDate         string         `json:"due_date,omitempty"`
	Memo            string         `json:"memo,omitempty"`
	Metadata        map[string]any `json:"metadata,omitempty"`
	CreatedAt       string         `json:"created_at"`
	UpdatedAt       string         `json:"updated_at"`
}

type InvoiceDetail struct {
	InvoiceListItem
	CreditApplied float64           `json:"credit_applied"`
	PlanName      string            `json:"plan_name,omitempty"`
	PoNumber      string            `json:"po_number,omitempty"`
	Reference     string            `json:"reference,omitempty"`
	LineItems     []InvoiceLineItem `json:"line_items"`
}

type InvoiceDownloadResult struct {
	URL       string `json:"url"`
	ExpiresAt string `json:"expires_at"`
}

type InvoiceSendResult struct {
	Sent   bool   `json:"sent"`
	SentAt string `json:"sent_at"`
}

type InvoiceStatusResult struct {
	ID        string `json:"id"`
	Status    string `json:"status"`
	UpdatedAt string `json:"updated_at"`
}

type CreateAdjustmentResult struct {
	ID            string         `json:"id"`
	Object        string         `json:"object"`
	Livemode      bool           `json:"livemode"`
	CustomerID    string         `json:"customer_id"`
	InvoiceNumber string         `json:"invoice_number"`
	Status        string         `json:"status"`
	InvoiceType   string         `json:"invoice_type"`
	Currency      string         `json:"currency"`
	Subtotal      float64        `json:"subtotal"`
	TaxAmount     float64        `json:"tax_amount"`
	Total         float64        `json:"total"`
	IssueDate     string         `json:"issue_date"`
	DueDate       string         `json:"due_date,omitempty"`
	Memo          string         `json:"memo,omitempty"`
	Metadata      map[string]any `json:"metadata,omitempty"`
	CreatedAt     string         `json:"created_at"`
	UpdatedAt     string         `json:"updated_at"`
}

type TransactionListItem struct {
	ID            string  `json:"id"`
	Object        string  `json:"object"`
	Livemode      bool    `json:"livemode"`
	InvoiceID     string  `json:"invoice_id"`
	GrossAmount   float64 `json:"gross_amount"`
	Subtotal      float64 `json:"subtotal"`
	TaxAmount     float64 `json:"tax_amount"`
	Currency      string  `json:"currency"`
	Status        string  `json:"status"`
	CustomerEmail string  `json:"customer_email"`
	CustomerName  string  `json:"customer_name,omitempty"`
	PaidAt        string  `json:"paid_at,omitempty"`
	CreatedAt     string  `json:"created_at"`
	UpdatedAt     string  `json:"updated_at"`
}

type TransactionDetail struct {
	TransactionListItem
	AvailableAt string `json:"available_at,omitempty"`
}

type TransactionRefundResult struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}

type TransactionRetryResult struct {
	ID                 string `json:"id"`
	Status             string `json:"status"`
	RetryInvoiceNumber string `json:"retry_invoice_number"`
}

type PromoCode struct {
	ID              string  `json:"id"`
	Object          string  `json:"object"`
	Livemode        bool    `json:"livemode"`
	Code            string  `json:"code"`
	DiscountType    string  `json:"discount_type"`
	DiscountValue   float64 `json:"discount_value"`
	DurationCycles  *int    `json:"duration_cycles,omitempty"`
	MaxRedemptions  *int    `json:"max_redemptions,omitempty"`
	ExpiresAt       string  `json:"expires_at,omitempty"`
	Active          bool    `json:"active"`
	RedemptionCount int     `json:"redemption_count"`
	CreatedAt       string  `json:"created_at"`
}

type PromoCodeDetail struct {
	PromoCode
	UpdatedAt string `json:"updated_at"`
}

type PlanGroup struct {
	ID          string `json:"id"`
	Object      string `json:"object"`
	Livemode    bool   `json:"livemode"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	IsPublic    bool   `json:"is_public"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type PlanGroupPlan struct {
	ID        string `json:"id"`
	Code      string `json:"code"`
	Name      string `json:"name"`
	SortOrder int    `json:"sort_order"`
}

type PlanGroupDetail struct {
	PlanGroup
	Plans []PlanGroupPlan `json:"plans"`
}

type PlanManage struct {
	ID                string         `json:"id"`
	Object            string         `json:"object"`
	Livemode          bool           `json:"livemode"`
	Name              string         `json:"name"`
	Code              string         `json:"code"`
	Description       string         `json:"description,omitempty"`
	ConsumptionModel  string         `json:"consumption_model,omitempty"`
	IsPublic          bool           `json:"is_public"`
	IsDefault         bool           `json:"is_default"`
	IsFree            bool           `json:"is_free"`
	BlockOnExhaustion bool           `json:"block_on_exhaustion"`
	SortOrder         int            `json:"sort_order"`
	PlanGroupID       string         `json:"plan_group_id,omitempty"`
	Metadata          map[string]any `json:"metadata,omitempty"`
	CreatedAt         string         `json:"created_at"`
	UpdatedAt         string         `json:"updated_at"`
}

type PlanFeatureManage struct {
	PlanID           string   `json:"plan_id"`
	FeatureID        string   `json:"feature_id"`
	Enabled          bool     `json:"enabled"`
	IncludedAmount   *int     `json:"included_amount,omitempty"`
	Unlimited        bool     `json:"unlimited"`
	OverageEnabled   bool     `json:"overage_enabled"`
	CreditsPerUnit   *int     `json:"credits_per_unit,omitempty"`
	PricingMode      string   `json:"pricing_mode"`
	OverageUnitPrice *float64 `json:"overage_unit_price,omitempty"`
	Margin           *float64 `json:"margin,omitempty"`
}

type PlanPriceManage struct {
	ID                       string          `json:"id"`
	Object                   string          `json:"object"`
	Livemode                 bool            `json:"livemode"`
	PlanID                   string          `json:"plan_id"`
	BillingInterval          BillingInterval `json:"billing_interval"`
	Price                    int             `json:"price"`
	IsDefault                bool            `json:"is_default"`
	TrialDays                int             `json:"trial_days"`
	IncludedBalance          *int            `json:"included_balance,omitempty"`
	IncludedCredits          *int            `json:"included_credits,omitempty"`
	IntroOfferEnabled        bool            `json:"intro_offer_enabled"`
	IntroOfferDiscountType   string          `json:"intro_offer_discount_type,omitempty"`
	IntroOfferDiscountValue  *float64        `json:"intro_offer_discount_value,omitempty"`
	IntroOfferDurationCycles *int            `json:"intro_offer_duration_cycles,omitempty"`
	CreatedAt                string          `json:"created_at"`
	UpdatedAt                string          `json:"updated_at"`
}

type RegionalPriceResult struct {
	PriceID   string `json:"price_id"`
	Overrides []RegionalPriceOverride `json:"overrides"`
}

type RegionalPriceOverride struct {
	Currency string `json:"currency"`
	Price    int    `json:"price"`
}

type DeleteResult struct {
	ID      string `json:"id"`
	Deleted bool   `json:"deleted"`
}

type RemoveResult struct {
	ID      string `json:"id"`
	Removed bool   `json:"removed"`
}

type Feature struct {
	ID          string      `json:"id"`
	Object      string      `json:"object"`
	Livemode    bool        `json:"livemode"`
	Name        string      `json:"name"`
	Code        string      `json:"code"`
	Type        FeatureType `json:"type"`
	Description string      `json:"description,omitempty"`
	UnitName    string      `json:"unit_name,omitempty"`
	CreatedAt   string      `json:"created_at"`
	UpdatedAt   string      `json:"updated_at"`
}

type Addon struct {
	ID               string   `json:"id"`
	Object           string   `json:"object"`
	Livemode         bool     `json:"livemode"`
	Name             string   `json:"name"`
	Slug             string   `json:"slug"`
	Description      string   `json:"description,omitempty"`
	BasePrice        int      `json:"base_price"`
	FeatureCode      string   `json:"feature_code"`
	FeatureName      string   `json:"feature_name"`
	ConsumptionModel string   `json:"consumption_model"`
	IncludedUnits    *int     `json:"included_units,omitempty"`
	OverageRate      *float64 `json:"overage_rate,omitempty"`
	CreditCost       *int     `json:"credit_cost,omitempty"`
	CreatedAt        string   `json:"created_at"`
	UpdatedAt        string   `json:"updated_at"`
}


type WebhookEndpoint struct {
	ID          string   `json:"id"`
	Object      string   `json:"object"`
	Livemode    bool     `json:"livemode"`
	URL         string   `json:"url"`
	Events      []string `json:"events"`
	Description string   `json:"description,omitempty"`
	IsActive    bool     `json:"is_active"`
	CreatedAt   string   `json:"created_at"`
}

type WebhookEndpointCreated struct {
	WebhookEndpoint
	SecretKey string `json:"secret_key"`
}

type WebhookTestResult struct {
	Success     bool   `json:"success"`
	DeliveredAt string `json:"delivered_at"`
}
