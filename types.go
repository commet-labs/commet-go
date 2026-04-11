package commet

type Customer struct {
	ID             string         `json:"id"`
	OrganizationID string         `json:"organization_id"`
	ExternalID     string         `json:"external_id,omitempty"`
	FullName       string         `json:"full_name,omitempty"`
	Domain         string         `json:"domain,omitempty"`
	Website        string         `json:"website,omitempty"`
	BillingEmail   string         `json:"billing_email"`
	Timezone       string         `json:"timezone,omitempty"`
	Language       string         `json:"language,omitempty"`
	Industry       string         `json:"industry,omitempty"`
	EmployeeCount  string         `json:"employee_count,omitempty"`
	Metadata       map[string]any `json:"metadata,omitempty"`
	IsActive       bool           `json:"is_active"`
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
	ID                     string `json:"id"`
	CustomerID             string `json:"customer_id"`
	PlanID                 string `json:"plan_id"`
	PlanName               string `json:"plan_name"`
	Name                   string `json:"name"`
	Description            string `json:"description,omitempty"`
	Status                 string `json:"status"`
	BillingInterval        string `json:"billing_interval,omitempty"`
	TrialEndsAt            string `json:"trial_ends_at,omitempty"`
	StartDate              string `json:"start_date"`
	EndDate                string `json:"end_date,omitempty"`
	CurrentPeriodStart     string `json:"current_period_start,omitempty"`
	CurrentPeriodEnd       string `json:"current_period_end,omitempty"`
	BillingDayOfMonth      int    `json:"billing_day_of_month"`
	CheckoutURL            string `json:"checkout_url,omitempty"`
	CreatedAt              string `json:"created_at"`
	UpdatedAt              string `json:"updated_at"`
	IntroOfferEndsAt       string `json:"intro_offer_ends_at,omitempty"`
	IntroOfferDiscountType string `json:"intro_offer_discount_type,omitempty"`
	IntroOfferDiscountValue *float64 `json:"intro_offer_discount_value,omitempty"`
}

type ActiveSubscription struct {
	ID             string              `json:"id"`
	CustomerID     string              `json:"customer_id"`
	Plan           SubscriptionPlan    `json:"plan"`
	Name           string              `json:"name"`
	Description    string              `json:"description,omitempty"`
	Status         string              `json:"status"`
	TrialEndsAt    string              `json:"trial_ends_at,omitempty"`
	CurrentPeriod  SubscriptionPeriod  `json:"current_period"`
	Features       []FeatureSummary    `json:"features"`
	StartDate      string              `json:"start_date"`
	EndDate        string              `json:"end_date,omitempty"`
	BillingDayOfMonth int              `json:"billing_day_of_month"`
	NextBillingDate string             `json:"next_billing_date"`
	CheckoutURL    string              `json:"checkout_url,omitempty"`
	CreatedAt      string              `json:"created_at"`
	UpdatedAt      string              `json:"updated_at"`
}

type SubscriptionPlan struct {
	ID              string  `json:"id"`
	Name            string  `json:"name"`
	BasePrice       float64 `json:"base_price"`
	BillingInterval string  `json:"billing_interval,omitempty"`
}

type SubscriptionPeriod struct {
	Start         string `json:"start"`
	End           string `json:"end"`
	DaysRemaining int    `json:"days_remaining"`
}

type FeatureSummary struct {
	Code    string        `json:"code"`
	Name    string        `json:"name"`
	Type    string        `json:"type"`
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
	BillingInterval string `json:"billing_interval"`
	Price           int    `json:"price"`
	IsDefault       bool   `json:"is_default"`
	TrialDays       int    `json:"trial_days"`
}

type PlanFeature struct {
	Code             string   `json:"code"`
	Name             string   `json:"name"`
	Type             string   `json:"type"`
	UnitName         string   `json:"unit_name,omitempty"`
	Enabled          *bool    `json:"enabled,omitempty"`
	IncludedAmount   *int     `json:"included_amount,omitempty"`
	Unlimited        *bool    `json:"unlimited,omitempty"`
	OverageEnabled   *bool    `json:"overage_enabled,omitempty"`
	OverageUnitPrice *float64 `json:"overage_unit_price,omitempty"`
}

type PlanDetail struct {
	ID          string              `json:"id"`
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
	BillingInterval string          `json:"billing_interval"`
	Price           int             `json:"price"`
	IsDefault       bool            `json:"is_default"`
	TrialDays       int             `json:"trial_days"`
	IntroOffer      *PlanIntroOffer `json:"intro_offer,omitempty"`
}

type PlanIntroOffer struct {
	Enabled        bool     `json:"enabled"`
	DiscountType   string   `json:"discount_type,omitempty"`
	DiscountValue  *float64 `json:"discount_value,omitempty"`
	DurationCycles *int     `json:"duration_cycles,omitempty"`
}

type PlanDetailFeature struct {
	Code             string          `json:"code"`
	Name             string          `json:"name"`
	Type             string          `json:"type"`
	UnitName         string          `json:"unit_name,omitempty"`
	Enabled          *bool           `json:"enabled,omitempty"`
	IncludedAmount   *int            `json:"included_amount,omitempty"`
	Unlimited        *bool           `json:"unlimited,omitempty"`
	OverageEnabled   *bool           `json:"overage_enabled,omitempty"`
	OverageUnitPrice *float64        `json:"overage_unit_price,omitempty"`
	Overage          *OverageConfig  `json:"overage,omitempty"`
}

type OverageConfig struct {
	Enabled   bool     `json:"enabled"`
	Model     string   `json:"model,omitempty"`
	UnitPrice *float64 `json:"unit_price,omitempty"`
}

type FeatureAccess struct {
	Code             string   `json:"code"`
	Name             string   `json:"name"`
	Type             string   `json:"type"`
	Allowed          bool     `json:"allowed"`
	Enabled          *bool    `json:"enabled,omitempty"`
	Current          *int     `json:"current,omitempty"`
	Included         *int     `json:"included,omitempty"`
	Remaining        *int     `json:"remaining,omitempty"`
	Overage          *int     `json:"overage,omitempty"`
	OverageUnitPrice *float64 `json:"overage_unit_price,omitempty"`
	Unlimited        *bool    `json:"unlimited,omitempty"`
	OverageEnabled   *bool    `json:"overage_enabled,omitempty"`
}

type CheckResult struct {
	Allowed bool `json:"allowed"`
}

type CanUseResult struct {
	Allowed       bool   `json:"allowed"`
	WillBeCharged bool   `json:"will_be_charged"`
	Reason        string `json:"reason,omitempty"`
}

type SeatEvent struct {
	ID              string `json:"id"`
	OrganizationID  string `json:"organization_id"`
	CustomerID      string `json:"customer_id"`
	SeatType        string `json:"seat_type"`
	EventType       string `json:"event_type"`
	Quantity        int    `json:"quantity"`
	PreviousBalance *int   `json:"previous_balance,omitempty"`
	NewBalance      int    `json:"new_balance"`
	Ts              string `json:"ts"`
	CreatedAt       string `json:"created_at"`
}

type SeatBalance struct {
	Current int    `json:"current"`
	AsOf    string `json:"as_of"`
}

type UsageEvent struct {
	ID              string                `json:"id"`
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
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Credits     int    `json:"credits"`
	Price       int    `json:"price"`
	Currency    string `json:"currency"`
}
