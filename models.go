package commet

import "encoding/json"

func hasJSONFields(fields map[string]json.RawMessage, names ...string) bool {
	for _, name := range names {
		if _, ok := fields[name]; !ok {
			return false
		}
	}
	return true
}

type ActiveAddon struct {
	Slug             string      `json:"slug"`
	Name             string      `json:"name"`
	BasePrice        int         `json:"base_price"`
	FeatureCode      string      `json:"feature_code"`
	FeatureName      string      `json:"feature_name"`
	FeatureType      FeatureType `json:"feature_type"`
	ConsumptionModel string      `json:"consumption_model"`
	ActivatedAt      string      `json:"activated_at"`
	Object           string      `json:"object"`
	Livemode         bool        `json:"livemode"`
}

type AddedPlanToGroup struct {
	Success  bool   `json:"success"`
	Object   string `json:"object"`
	Livemode bool   `json:"livemode"`
}

type Addon struct {
	ID               string  `json:"id"`
	Name             string  `json:"name"`
	Slug             string  `json:"slug"`
	Description      *string `json:"description"`
	BasePrice        int     `json:"base_price"`
	FeatureCode      string  `json:"feature_code"`
	FeatureName      string  `json:"feature_name"`
	CreatedAt        string  `json:"created_at"`
	UpdatedAt        string  `json:"updated_at"`
	ConsumptionModel string  `json:"consumption_model"`
	IncludedUnits    *int    `json:"included_units"`
	OverageRate      *int    `json:"overage_rate"`
	CreditCost       *int    `json:"credit_cost"`
	Object           string  `json:"object"`
	Livemode         bool    `json:"livemode"`
}

type AddonsListActiveResult struct {
	Object     string        `json:"object"`
	Data       []ActiveAddon `json:"data"`
	HasMore    bool          `json:"has_more"`
	NextCursor *string       `json:"next_cursor,omitempty"`
}

type AddonsListResult struct {
	Object     string  `json:"object"`
	Data       []Addon `json:"data"`
	HasMore    bool    `json:"has_more"`
	NextCursor *string `json:"next_cursor,omitempty"`
}

type AddPlanFeatureParamsOverage struct {
	Enabled   *bool `json:"enabled,omitempty"`
	UnitPrice *int  `json:"unit_price,omitempty"`
}

type AddPlanPriceParamsMarketPricesItem struct {
	MarketGroupID string `json:"market_group_id"`
	Currency      string `json:"currency"`
	Price         int    `json:"price"`
}

type ApiKey struct {
	ID         string  `json:"id"`
	Name       string  `json:"name"`
	Prefix     string  `json:"prefix"`
	ExpiresAt  *string `json:"expires_at"`
	LastUsedAt *string `json:"last_used_at"`
	CreatedAt  string  `json:"created_at"`
	Object     string  `json:"object"`
	Livemode   bool    `json:"livemode"`
}

type ApiKeysListResult struct {
	Object     string   `json:"object"`
	Data       []ApiKey `json:"data"`
	HasMore    bool     `json:"has_more"`
	NextCursor *string  `json:"next_cursor,omitempty"`
}

type BalanceAdjustment struct {
	Amount     int     `json:"amount"`
	NewBalance int     `json:"new_balance"`
	Reason     *string `json:"reason"`
	Object     string  `json:"object"`
	Livemode   bool    `json:"livemode"`
}

type BalanceTopup struct {
	Amount   int    `json:"amount"`
	Object   string `json:"object"`
	Livemode bool   `json:"livemode"`
}

type BatchCreateCustomersParamsCustomersItem struct {
	Email       string                                          `json:"email"`
	ID          *string                                         `json:"id,omitempty"`
	ExternalID  *string                                         `json:"external_id,omitempty"`
	FullName    *string                                         `json:"full_name,omitempty"`
	TaxDocument *string                                         `json:"tax_document,omitempty"`
	Timezone    *Timezone                                       `json:"timezone,omitempty"`
	Metadata    map[string]any                                  `json:"metadata,omitempty"`
	Address     *BatchCreateCustomersParamsCustomersItemAddress `json:"address,omitempty"`
}

type BatchCreateCustomersParamsCustomersItemAddress struct {
	Line1      string  `json:"line1"`
	Line2      *string `json:"line2,omitempty"`
	City       string  `json:"city"`
	State      *string `json:"state,omitempty"`
	PostalCode string  `json:"postal_code"`
	Country    string  `json:"country"`
	Region     *string `json:"region,omitempty"`
}

type ClaimLink struct {
	URL       string `json:"url"`
	ExpiresAt string `json:"expires_at"`
	Object    string `json:"object"`
	Livemode  bool   `json:"livemode"`
}

type CompletePayoutVerificationParamsBank struct {
	AccountNumber     string  `json:"account_number"`
	AccountHolderName string  `json:"account_holder_name"`
	RoutingNumber     *string `json:"routing_number,omitempty"`
	AccountType       *string `json:"account_type,omitempty"`
}

type CompletePayoutVerificationParamsCompany struct {
	Name           string                                                `json:"name"`
	TaxID          string                                                `json:"tax_id"`
	Address        CompletePayoutVerificationParamsCompanyAddress        `json:"address"`
	Representative CompletePayoutVerificationParamsCompanyRepresentative `json:"representative"`
}

type CompletePayoutVerificationParamsCompanyAddress struct {
	Line1      string  `json:"line1"`
	Line2      *string `json:"line2,omitempty"`
	City       string  `json:"city"`
	State      *string `json:"state,omitempty"`
	PostalCode string  `json:"postal_code"`
	Country    string  `json:"country"`
}

type CompletePayoutVerificationParamsCompanyRepresentative struct {
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Phone     *string `json:"phone,omitempty"`
	Email     *string `json:"email,omitempty"`
}

type CompletePayoutVerificationParamsIndividual struct {
	FirstName   string                                            `json:"first_name"`
	LastName    string                                            `json:"last_name"`
	Phone       string                                            `json:"phone"`
	DateOfBirth string                                            `json:"date_of_birth"`
	SsnLast4    *string                                           `json:"ssn_last4,omitempty"`
	IDNumber    *string                                           `json:"id_number,omitempty"`
	Address     CompletePayoutVerificationParamsIndividualAddress `json:"address"`
}

type CompletePayoutVerificationParamsIndividualAddress struct {
	Line1      string  `json:"line1"`
	Line2      *string `json:"line2,omitempty"`
	City       string  `json:"city"`
	State      *string `json:"state,omitempty"`
	PostalCode string  `json:"postal_code"`
	Country    string  `json:"country"`
}

type CreateCustomerParamsAddress struct {
	Line1      string  `json:"line1"`
	Line2      *string `json:"line2,omitempty"`
	City       string  `json:"city"`
	State      *string `json:"state,omitempty"`
	PostalCode string  `json:"postal_code"`
	Country    string  `json:"country"`
	Region     *string `json:"region,omitempty"`
}

type CreatedApiKey struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	APIKey    string `json:"api_key"`
	Prefix    string `json:"prefix"`
	ExpiresAt string `json:"expires_at"`
	CreatedAt string `json:"created_at"`
	Object    string `json:"object"`
	Livemode  bool   `json:"livemode"`
}

type CreatedSubscription struct {
	ID                  string                                  `json:"id"`
	CustomerID          string                                  `json:"customer_id"`
	Plan                CreatedSubscriptionPlan                 `json:"plan"`
	Name                string                                  `json:"name"`
	Description         *string                                 `json:"description"`
	Status              SubscriptionStatus                      `json:"status"`
	BillingInterval     *BillingInterval                        `json:"billing_interval"`
	TrialEndsAt         *string                                 `json:"trial_ends_at"`
	CurrentPeriod       *CreatedSubscriptionCurrentPeriod       `json:"current_period"`
	Cancellation        *CreatedSubscriptionCancellation        `json:"cancellation"`
	CancelAtPeriodEnd   bool                                    `json:"cancel_at_period_end"`
	ScheduledPlanChange *CreatedSubscriptionScheduledPlanChange `json:"scheduled_plan_change"`
	Discount            *CreatedSubscriptionDiscount            `json:"discount"`
	StartDate           string                                  `json:"start_date"`
	EndDate             *string                                 `json:"end_date"`
	BillingDayOfMonth   *int                                    `json:"billing_day_of_month"`
	NextBillingDate     *string                                 `json:"next_billing_date"`
	CheckoutURL         *string                                 `json:"checkout_url"`
	CreatedAt           string                                  `json:"created_at"`
	UpdatedAt           string                                  `json:"updated_at"`
	CheckoutProvider    *PaymentProvider                        `json:"checkout_provider"`
	PriceID             *string                                 `json:"price_id"`
	Object              string                                  `json:"object"`
	Livemode            bool                                    `json:"livemode"`
}

type CreatedSubscriptionCancellation struct {
	ScheduledAt string  `json:"scheduled_at"`
	Reason      *string `json:"reason"`
	EffectiveAt string  `json:"effective_at"`
}

type CreatedSubscriptionCurrentPeriod struct {
	Start         string  `json:"start"`
	End           string  `json:"end"`
	DaysRemaining float64 `json:"days_remaining"`
}

type CreatedSubscriptionDiscount struct {
	Type   string  `json:"type"`
	Value  float64 `json:"value"`
	Name   *string `json:"name"`
	EndsAt *string `json:"ends_at"`
}

type CreatedSubscriptionPlan struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type CreatedSubscriptionScheduledPlanChange struct {
	ChangeType         string  `json:"change_type"`
	NewPlanID          *string `json:"new_plan_id"`
	NewPlanName        *string `json:"new_plan_name"`
	NewBillingInterval *string `json:"new_billing_interval"`
	ScheduledFor       string  `json:"scheduled_for"`
}

type CreatedWebhook struct {
	ID          string   `json:"id"`
	URL         string   `json:"url"`
	Events      []string `json:"events"`
	Description *string  `json:"description"`
	IsActive    bool     `json:"is_active"`
	APIVersion  *string  `json:"api_version"`
	CreatedAt   string   `json:"created_at"`
	SecretKey   string   `json:"secret_key"`
	Object      string   `json:"object"`
	Livemode    bool     `json:"livemode"`
}

type CreateOfferParamsPhasesItem struct {
	Value any
}

func (value *CreateOfferParamsPhasesItem) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	var discriminator string
	if raw, ok := fields["type"]; ok {
		if err := json.Unmarshal(raw, &discriminator); err != nil {
			return err
		}
	}
	switch discriminator {
	case "free_trial":
		var decoded CreateOfferParamsPhasesItemVariant1
		if err := json.Unmarshal(data, &decoded); err != nil {
			return err
		}
		value.Value = decoded
		return nil
	case "percentage":
		var decoded CreateOfferParamsPhasesItemVariant2
		if err := json.Unmarshal(data, &decoded); err != nil {
			return err
		}
		value.Value = decoded
		return nil
	case "amount_off":
		var decoded CreateOfferParamsPhasesItemVariant3
		if err := json.Unmarshal(data, &decoded); err != nil {
			return err
		}
		value.Value = decoded
		return nil
	case "fixed_price":
		var decoded CreateOfferParamsPhasesItemVariant4
		if err := json.Unmarshal(data, &decoded); err != nil {
			return err
		}
		value.Value = decoded
		return nil
	}
	if hasJSONFields(fields, "type", "duration_cycles", "percentage") {
		var decoded CreateOfferParamsPhasesItemVariant2
		if err := json.Unmarshal(data, &decoded); err != nil {
			return err
		}
		value.Value = decoded
		return nil
	}
	if hasJSONFields(fields, "type", "duration_cycles", "amounts") {
		var decoded CreateOfferParamsPhasesItemVariant3
		if err := json.Unmarshal(data, &decoded); err != nil {
			return err
		}
		value.Value = decoded
		return nil
	}
	if hasJSONFields(fields, "type", "duration_cycles", "prices") {
		var decoded CreateOfferParamsPhasesItemVariant4
		if err := json.Unmarshal(data, &decoded); err != nil {
			return err
		}
		value.Value = decoded
		return nil
	}
	if hasJSONFields(fields, "type", "duration_days") {
		var decoded CreateOfferParamsPhasesItemVariant1
		if err := json.Unmarshal(data, &decoded); err != nil {
			return err
		}
		value.Value = decoded
		return nil
	}
	var decoded CreateOfferParamsPhasesItemVariant1
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	value.Value = decoded
	return nil
}

func (value CreateOfferParamsPhasesItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(value.Value)
}

func (value CreateOfferParamsPhasesItem) AsCreateOfferParamsPhasesItemVariant1() (CreateOfferParamsPhasesItemVariant1, bool) {
	decoded, ok := value.Value.(CreateOfferParamsPhasesItemVariant1)
	return decoded, ok
}

func (value CreateOfferParamsPhasesItem) AsCreateOfferParamsPhasesItemVariant2() (CreateOfferParamsPhasesItemVariant2, bool) {
	decoded, ok := value.Value.(CreateOfferParamsPhasesItemVariant2)
	return decoded, ok
}

func (value CreateOfferParamsPhasesItem) AsCreateOfferParamsPhasesItemVariant3() (CreateOfferParamsPhasesItemVariant3, bool) {
	decoded, ok := value.Value.(CreateOfferParamsPhasesItemVariant3)
	return decoded, ok
}

func (value CreateOfferParamsPhasesItem) AsCreateOfferParamsPhasesItemVariant4() (CreateOfferParamsPhasesItemVariant4, bool) {
	decoded, ok := value.Value.(CreateOfferParamsPhasesItemVariant4)
	return decoded, ok
}

type CreateOfferParamsPhasesItemVariant1 struct {
	Type         string `json:"type"`
	DurationDays int    `json:"duration_days"`
}

type CreateOfferParamsPhasesItemVariant2 struct {
	Type           string `json:"type"`
	DurationCycles int    `json:"duration_cycles"`
	Percentage     int    `json:"percentage"`
}

type CreateOfferParamsPhasesItemVariant3 struct {
	Type           string                                           `json:"type"`
	DurationCycles int                                              `json:"duration_cycles"`
	Amounts        []CreateOfferParamsPhasesItemVariant3AmountsItem `json:"amounts"`
}

type CreateOfferParamsPhasesItemVariant3AmountsItem struct {
	Currency string `json:"currency"`
	Amount   int    `json:"amount"`
}

type CreateOfferParamsPhasesItemVariant4 struct {
	Type           string                                          `json:"type"`
	DurationCycles int                                             `json:"duration_cycles"`
	Prices         []CreateOfferParamsPhasesItemVariant4PricesItem `json:"prices"`
}

type CreateOfferParamsPhasesItemVariant4PricesItem struct {
	Currency string `json:"currency"`
	Amount   int    `json:"amount"`
}

type CreditGrant struct {
	Credits  int    `json:"credits"`
	Object   string `json:"object"`
	Livemode bool   `json:"livemode"`
}

type CreditPack struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
	Credits     int     `json:"credits"`
	Price       int     `json:"price"`
	IsActive    bool    `json:"is_active"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
	Object      string  `json:"object"`
	Livemode    bool    `json:"livemode"`
}

type CreditPackListItem struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
	Credits     int     `json:"credits"`
	Price       int     `json:"price"`
	Currency    string  `json:"currency"`
	Object      string  `json:"object"`
	Livemode    bool    `json:"livemode"`
}

type CreditPacksListResult struct {
	Object     string               `json:"object"`
	Data       []CreditPackListItem `json:"data"`
	HasMore    bool                 `json:"has_more"`
	NextCursor *string              `json:"next_cursor,omitempty"`
}

type Customer struct {
	ID           string         `json:"id"`
	ExternalID   *string        `json:"external_id"`
	FullName     *string        `json:"full_name"`
	Email        string         `json:"email"`
	TaxDocument  *string        `json:"tax_document"`
	DocumentType *string        `json:"document_type"`
	Timezone     *string        `json:"timezone"`
	Metadata     map[string]any `json:"metadata"`
	CreatedAt    string         `json:"created_at"`
	UpdatedAt    string         `json:"updated_at"`
	Object       string         `json:"object"`
	Livemode     bool           `json:"livemode"`
}

type CustomerBatch struct {
	Successful []CustomerBatchSuccessfulItem `json:"successful"`
	Failed     []CustomerBatchFailedItem     `json:"failed"`
	Object     string                        `json:"object"`
	Livemode   bool                          `json:"livemode"`
}

type CustomerBatchFailedItem struct {
	Index int                         `json:"index"`
	Error string                      `json:"error"`
	Data  CustomerBatchFailedItemData `json:"data"`
}

type CustomerBatchFailedItemData struct {
	ID          *string                             `json:"id,omitempty"`
	ExternalID  *string                             `json:"external_id,omitempty"`
	Email       string                              `json:"email"`
	FullName    *string                             `json:"full_name,omitempty"`
	TaxDocument *string                             `json:"tax_document,omitempty"`
	Timezone    *string                             `json:"timezone,omitempty"`
	Metadata    map[string]any                      `json:"metadata,omitempty"`
	Address     *CustomerBatchFailedItemDataAddress `json:"address,omitempty"`
}

type CustomerBatchFailedItemDataAddress struct {
	Line1      string  `json:"line1"`
	Line2      *string `json:"line2,omitempty"`
	City       string  `json:"city"`
	State      *string `json:"state,omitempty"`
	PostalCode string  `json:"postal_code"`
	Country    string  `json:"country"`
	Region     *string `json:"region,omitempty"`
}

type CustomerBatchSuccessfulItem struct {
	ID         string  `json:"id"`
	ExternalID *string `json:"external_id"`
	Email      string  `json:"email"`
}

type CustomersListResult struct {
	Object     string     `json:"object"`
	Data       []Customer `json:"data"`
	HasMore    bool       `json:"has_more"`
	NextCursor *string    `json:"next_cursor,omitempty"`
}

type DeletedObject struct {
	ID       string `json:"id"`
	Deleted  any    `json:"deleted"`
	Object   string `json:"object"`
	Livemode bool   `json:"livemode"`
}

type DeletedOffer struct {
	Deleted  any    `json:"deleted"`
	Object   string `json:"object"`
	Livemode bool   `json:"livemode"`
}

type DeletedPlanRegionalPricing struct {
	Deleted  any    `json:"deleted"`
	Object   string `json:"object"`
	Livemode bool   `json:"livemode"`
}

type DeletedSubscriptionAddon struct {
	ID            string  `json:"id"`
	Status        string  `json:"status"`
	DeactivatedAt *string `json:"deactivated_at"`
	Object        string  `json:"object"`
	Livemode      bool    `json:"livemode"`
}

type Feature struct {
	ID          string      `json:"id"`
	Name        string      `json:"name"`
	Code        string      `json:"code"`
	Type        FeatureType `json:"type"`
	Description *string     `json:"description"`
	UnitName    *string     `json:"unit_name"`
	CreatedAt   string      `json:"created_at"`
	UpdatedAt   string      `json:"updated_at"`
	Object      string      `json:"object"`
	Livemode    bool        `json:"livemode"`
}

type FeatureAccess struct {
	Value any
}

func (value *FeatureAccess) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	var discriminator string
	if raw, ok := fields["type"]; ok {
		if err := json.Unmarshal(raw, &discriminator); err != nil {
			return err
		}
	}
	switch discriminator {
	case "boolean":
		var decoded FeatureAccessVariant1
		if err := json.Unmarshal(data, &decoded); err != nil {
			return err
		}
		value.Value = decoded
		return nil
	case "usage":
		var decoded FeatureAccessVariant2
		if err := json.Unmarshal(data, &decoded); err != nil {
			return err
		}
		value.Value = decoded
		return nil
	case "seats":
		var decoded FeatureAccessVariant3
		if err := json.Unmarshal(data, &decoded); err != nil {
			return err
		}
		value.Value = decoded
		return nil
	case "quota":
		var decoded FeatureAccessVariant4
		if err := json.Unmarshal(data, &decoded); err != nil {
			return err
		}
		value.Value = decoded
		return nil
	}
	if hasJSONFields(fields, "code", "name", "unit_name", "allowed", "type", "enabled", "object", "livemode") {
		var decoded FeatureAccessVariant1
		if err := json.Unmarshal(data, &decoded); err != nil {
			return err
		}
		value.Value = decoded
		return nil
	}
	if hasJSONFields(fields, "code", "name", "unit_name", "allowed", "type", "consumption", "object", "livemode") {
		var decoded FeatureAccessVariant2
		if err := json.Unmarshal(data, &decoded); err != nil {
			return err
		}
		value.Value = decoded
		return nil
	}
	if hasJSONFields(fields, "code", "name", "unit_name", "allowed", "type", "usage", "object", "livemode") {
		var decoded FeatureAccessVariant3
		if err := json.Unmarshal(data, &decoded); err != nil {
			return err
		}
		value.Value = decoded
		return nil
	}
	if hasJSONFields(fields, "code", "name", "unit_name", "allowed", "type", "usage", "object", "livemode") {
		var decoded FeatureAccessVariant4
		if err := json.Unmarshal(data, &decoded); err != nil {
			return err
		}
		value.Value = decoded
		return nil
	}
	var decoded FeatureAccessVariant1
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	value.Value = decoded
	return nil
}

func (value FeatureAccess) MarshalJSON() ([]byte, error) {
	return json.Marshal(value.Value)
}

func (value FeatureAccess) AsFeatureAccessVariant1() (FeatureAccessVariant1, bool) {
	decoded, ok := value.Value.(FeatureAccessVariant1)
	return decoded, ok
}

func (value FeatureAccess) AsFeatureAccessVariant2() (FeatureAccessVariant2, bool) {
	decoded, ok := value.Value.(FeatureAccessVariant2)
	return decoded, ok
}

func (value FeatureAccess) AsFeatureAccessVariant3() (FeatureAccessVariant3, bool) {
	decoded, ok := value.Value.(FeatureAccessVariant3)
	return decoded, ok
}

func (value FeatureAccess) AsFeatureAccessVariant4() (FeatureAccessVariant4, bool) {
	decoded, ok := value.Value.(FeatureAccessVariant4)
	return decoded, ok
}

type FeatureAccessListResult struct {
	Object     string          `json:"object"`
	Data       []FeatureAccess `json:"data"`
	HasMore    bool            `json:"has_more"`
	NextCursor *string         `json:"next_cursor,omitempty"`
}

type FeatureAccessVariant1 struct {
	Code     string  `json:"code"`
	Name     string  `json:"name"`
	UnitName *string `json:"unit_name"`
	Allowed  bool    `json:"allowed"`
	Type     string  `json:"type"`
	Enabled  bool    `json:"enabled"`
	Object   string  `json:"object"`
	Livemode bool    `json:"livemode"`
}

type FeatureAccessVariant2 struct {
	Code        string                           `json:"code"`
	Name        string                           `json:"name"`
	UnitName    *string                          `json:"unit_name"`
	Allowed     bool                             `json:"allowed"`
	Type        string                           `json:"type"`
	Consumption FeatureAccessVariant2Consumption `json:"consumption"`
	Object      string                           `json:"object"`
	Livemode    bool                             `json:"livemode"`
}

type FeatureAccessVariant2Consumption struct {
	Value any
}

func (value *FeatureAccessVariant2Consumption) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	var discriminator string
	if raw, ok := fields["model"]; ok {
		if err := json.Unmarshal(raw, &discriminator); err != nil {
			return err
		}
	}
	switch discriminator {
	case "metered":
		var decoded FeatureAccessVariant2ConsumptionVariant1
		if err := json.Unmarshal(data, &decoded); err != nil {
			return err
		}
		value.Value = decoded
		return nil
	case "credits":
		var decoded FeatureAccessVariant2ConsumptionVariant2
		if err := json.Unmarshal(data, &decoded); err != nil {
			return err
		}
		value.Value = decoded
		return nil
	case "balance":
		var decoded FeatureAccessVariant2ConsumptionVariant3
		if err := json.Unmarshal(data, &decoded); err != nil {
			return err
		}
		value.Value = decoded
		return nil
	}
	if hasJSONFields(fields, "model", "period", "units_used", "included_units", "unlimited", "overage") {
		var decoded FeatureAccessVariant2ConsumptionVariant1
		if err := json.Unmarshal(data, &decoded); err != nil {
			return err
		}
		value.Value = decoded
		return nil
	}
	if hasJSONFields(fields, "model", "period", "units_used", "credits_per_unit", "credits_consumed", "available_units") {
		var decoded FeatureAccessVariant2ConsumptionVariant2
		if err := json.Unmarshal(data, &decoded); err != nil {
			return err
		}
		value.Value = decoded
		return nil
	}
	if hasJSONFields(fields, "model", "period", "units_used", "spent") {
		var decoded FeatureAccessVariant2ConsumptionVariant3
		if err := json.Unmarshal(data, &decoded); err != nil {
			return err
		}
		value.Value = decoded
		return nil
	}
	var decoded FeatureAccessVariant2ConsumptionVariant1
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	value.Value = decoded
	return nil
}

func (value FeatureAccessVariant2Consumption) MarshalJSON() ([]byte, error) {
	return json.Marshal(value.Value)
}

func (value FeatureAccessVariant2Consumption) AsFeatureAccessVariant2ConsumptionVariant1() (FeatureAccessVariant2ConsumptionVariant1, bool) {
	decoded, ok := value.Value.(FeatureAccessVariant2ConsumptionVariant1)
	return decoded, ok
}

func (value FeatureAccessVariant2Consumption) AsFeatureAccessVariant2ConsumptionVariant2() (FeatureAccessVariant2ConsumptionVariant2, bool) {
	decoded, ok := value.Value.(FeatureAccessVariant2ConsumptionVariant2)
	return decoded, ok
}

func (value FeatureAccessVariant2Consumption) AsFeatureAccessVariant2ConsumptionVariant3() (FeatureAccessVariant2ConsumptionVariant3, bool) {
	decoded, ok := value.Value.(FeatureAccessVariant2ConsumptionVariant3)
	return decoded, ok
}

type FeatureAccessVariant2ConsumptionVariant1 struct {
	Model          string                                          `json:"model"`
	Period         FeatureAccessVariant2ConsumptionVariant1Period  `json:"period"`
	UnitsUsed      float64                                         `json:"units_used"`
	IncludedUnits  float64                                         `json:"included_units"`
	RemainingUnits *float64                                        `json:"remaining_units,omitempty"`
	Unlimited      bool                                            `json:"unlimited"`
	Overage        FeatureAccessVariant2ConsumptionVariant1Overage `json:"overage"`
}

type FeatureAccessVariant2ConsumptionVariant1Overage struct {
	Enabled   bool                                                      `json:"enabled"`
	Units     float64                                                   `json:"units"`
	UnitPrice *FeatureAccessVariant2ConsumptionVariant1OverageUnitPrice `json:"unit_price,omitempty"`
}

type FeatureAccessVariant2ConsumptionVariant1OverageUnitPrice struct {
	Amount   int    `json:"amount"`
	Currency string `json:"currency"`
	Scale    any    `json:"scale"`
}

type FeatureAccessVariant2ConsumptionVariant1Period struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

type FeatureAccessVariant2ConsumptionVariant2 struct {
	Model           string                                         `json:"model"`
	Period          FeatureAccessVariant2ConsumptionVariant2Period `json:"period"`
	UnitsUsed       float64                                        `json:"units_used"`
	CreditsPerUnit  int                                            `json:"credits_per_unit"`
	CreditsConsumed float64                                        `json:"credits_consumed"`
	AvailableUnits  int                                            `json:"available_units"`
}

type FeatureAccessVariant2ConsumptionVariant2Period struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

type FeatureAccessVariant2ConsumptionVariant3 struct {
	Model          string                                             `json:"model"`
	Period         FeatureAccessVariant2ConsumptionVariant3Period     `json:"period"`
	UnitsUsed      float64                                            `json:"units_used"`
	Spent          FeatureAccessVariant2ConsumptionVariant3Spent      `json:"spent"`
	AvailableUnits *int                                               `json:"available_units,omitempty"`
	UnitPrice      *FeatureAccessVariant2ConsumptionVariant3UnitPrice `json:"unit_price,omitempty"`
}

type FeatureAccessVariant2ConsumptionVariant3Period struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

type FeatureAccessVariant2ConsumptionVariant3Spent struct {
	Amount   int    `json:"amount"`
	Currency string `json:"currency"`
}

type FeatureAccessVariant2ConsumptionVariant3UnitPrice struct {
	Amount   int    `json:"amount"`
	Currency string `json:"currency"`
	Scale    any    `json:"scale"`
}

type FeatureAccessVariant3 struct {
	Code     string                     `json:"code"`
	Name     string                     `json:"name"`
	UnitName *string                    `json:"unit_name"`
	Allowed  bool                       `json:"allowed"`
	Type     string                     `json:"type"`
	Usage    FeatureAccessVariant3Usage `json:"usage"`
	Object   string                     `json:"object"`
	Livemode bool                       `json:"livemode"`
}

type FeatureAccessVariant3Usage struct {
	Period         FeatureAccessVariant3UsagePeriod  `json:"period"`
	UnitsUsed      float64                           `json:"units_used"`
	IncludedUnits  float64                           `json:"included_units"`
	RemainingUnits *float64                          `json:"remaining_units,omitempty"`
	Unlimited      bool                              `json:"unlimited"`
	Overage        FeatureAccessVariant3UsageOverage `json:"overage"`
}

type FeatureAccessVariant3UsageOverage struct {
	Enabled   bool                                        `json:"enabled"`
	Units     float64                                     `json:"units"`
	UnitPrice *FeatureAccessVariant3UsageOverageUnitPrice `json:"unit_price,omitempty"`
}

type FeatureAccessVariant3UsageOverageUnitPrice struct {
	Amount   int    `json:"amount"`
	Currency string `json:"currency"`
	Scale    any    `json:"scale"`
}

type FeatureAccessVariant3UsagePeriod struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

type FeatureAccessVariant4 struct {
	Code     string                     `json:"code"`
	Name     string                     `json:"name"`
	UnitName *string                    `json:"unit_name"`
	Allowed  bool                       `json:"allowed"`
	Type     string                     `json:"type"`
	Usage    FeatureAccessVariant4Usage `json:"usage"`
	Object   string                     `json:"object"`
	Livemode bool                       `json:"livemode"`
}

type FeatureAccessVariant4Usage struct {
	Period         FeatureAccessVariant4UsagePeriod  `json:"period"`
	UnitsUsed      float64                           `json:"units_used"`
	IncludedUnits  float64                           `json:"included_units"`
	RemainingUnits *float64                          `json:"remaining_units,omitempty"`
	Unlimited      bool                              `json:"unlimited"`
	Overage        FeatureAccessVariant4UsageOverage `json:"overage"`
	BilledUnits    float64                           `json:"billed_units"`
}

type FeatureAccessVariant4UsageOverage struct {
	Enabled   bool                                        `json:"enabled"`
	Units     float64                                     `json:"units"`
	UnitPrice *FeatureAccessVariant4UsageOverageUnitPrice `json:"unit_price,omitempty"`
}

type FeatureAccessVariant4UsageOverageUnitPrice struct {
	Amount   int    `json:"amount"`
	Currency string `json:"currency"`
	Scale    any    `json:"scale"`
}

type FeatureAccessVariant4UsagePeriod struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

type FeaturesListResult struct {
	Object     string    `json:"object"`
	Data       []Feature `json:"data"`
	HasMore    bool      `json:"has_more"`
	NextCursor *string   `json:"next_cursor,omitempty"`
}

type Invoice struct {
	ID             string                 `json:"id"`
	CustomerID     string                 `json:"customer_id"`
	SubscriptionID *string                `json:"subscription_id"`
	InvoiceNumber  string                 `json:"invoice_number"`
	Status         string                 `json:"status"`
	InvoiceType    InvoiceType            `json:"invoice_type"`
	Currency       string                 `json:"currency"`
	Subtotal       int                    `json:"subtotal"`
	DiscountAmount int                    `json:"discount_amount"`
	TaxAmount      int                    `json:"tax_amount"`
	Total          int                    `json:"total"`
	PeriodStart    string                 `json:"period_start"`
	PeriodEnd      string                 `json:"period_end"`
	IssueDate      string                 `json:"issue_date"`
	DueDate        string                 `json:"due_date"`
	Memo           *string                `json:"memo"`
	Metadata       map[string]any         `json:"metadata"`
	CreatedAt      string                 `json:"created_at"`
	UpdatedAt      string                 `json:"updated_at"`
	CreditApplied  int                    `json:"credit_applied"`
	PlanName       *string                `json:"plan_name"`
	PONumber       *string                `json:"po_number"`
	Reference      *string                `json:"reference"`
	LineItems      []InvoiceLineItemsItem `json:"line_items"`
	Object         string                 `json:"object"`
	Livemode       bool                   `json:"livemode"`
}

type InvoiceDownload struct {
	URL       string `json:"url"`
	ExpiresAt string `json:"expires_at"`
	Object    string `json:"object"`
	Livemode  bool   `json:"livemode"`
}

type InvoiceLineItemsItem struct {
	LineType       string  `json:"line_type"`
	FeatureName    *string `json:"feature_name"`
	Description    string  `json:"description"`
	Quantity       int     `json:"quantity"`
	UnitAmount     int     `json:"unit_amount"`
	Amount         int     `json:"amount"`
	IncludedAmount *int    `json:"included_amount"`
	UsedAmount     *int    `json:"used_amount"`
	OverageAmount  *int    `json:"overage_amount"`
	DiscountType   *string `json:"discount_type"`
	DiscountValue  *int    `json:"discount_value"`
	DiscountName   *string `json:"discount_name"`
	ChargeType     string  `json:"charge_type"`
}

type InvoiceListItem struct {
	ID             string         `json:"id"`
	CustomerID     string         `json:"customer_id"`
	SubscriptionID *string        `json:"subscription_id"`
	InvoiceNumber  string         `json:"invoice_number"`
	Status         string         `json:"status"`
	InvoiceType    InvoiceType    `json:"invoice_type"`
	Currency       string         `json:"currency"`
	Subtotal       int            `json:"subtotal"`
	DiscountAmount int            `json:"discount_amount"`
	TaxAmount      int            `json:"tax_amount"`
	Total          int            `json:"total"`
	PeriodStart    string         `json:"period_start"`
	PeriodEnd      string         `json:"period_end"`
	IssueDate      string         `json:"issue_date"`
	DueDate        string         `json:"due_date"`
	Memo           *string        `json:"memo"`
	Metadata       map[string]any `json:"metadata"`
	CreatedAt      string         `json:"created_at"`
	UpdatedAt      string         `json:"updated_at"`
	Object         string         `json:"object"`
	Livemode       bool           `json:"livemode"`
}

type InvoicesListResult struct {
	Object     string            `json:"object"`
	Data       []InvoiceListItem `json:"data"`
	HasMore    bool              `json:"has_more"`
	NextCursor *string           `json:"next_cursor,omitempty"`
}

type MarketGroup struct {
	ID           string         `json:"id"`
	Name         string         `json:"name"`
	CountryCodes []string       `json:"country_codes"`
	Metadata     map[string]any `json:"metadata"`
	CreatedAt    string         `json:"created_at"`
	UpdatedAt    string         `json:"updated_at"`
	Object       string         `json:"object"`
	Livemode     bool           `json:"livemode"`
}

type Offer struct {
	ID           string            `json:"id"`
	Name         string            `json:"name"`
	Purpose      string            `json:"purpose"`
	PlanPriceIds []string          `json:"plan_price_ids"`
	Phases       []OfferPhasesItem `json:"phases"`
	Metadata     map[string]any    `json:"metadata"`
	StartsAt     *string           `json:"starts_at"`
	EndsAt       *string           `json:"ends_at"`
	Active       bool              `json:"active"`
	CreatedAt    string            `json:"created_at"`
	UpdatedAt    string            `json:"updated_at"`
	Object       string            `json:"object"`
	Livemode     bool              `json:"livemode"`
}

type OfferPhasesItem struct {
	Value any
}

func (value *OfferPhasesItem) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	var discriminator string
	if raw, ok := fields["type"]; ok {
		if err := json.Unmarshal(raw, &discriminator); err != nil {
			return err
		}
	}
	switch discriminator {
	case "free_trial":
		var decoded OfferPhasesItemVariant1
		if err := json.Unmarshal(data, &decoded); err != nil {
			return err
		}
		value.Value = decoded
		return nil
	case "percentage":
		var decoded OfferPhasesItemVariant2
		if err := json.Unmarshal(data, &decoded); err != nil {
			return err
		}
		value.Value = decoded
		return nil
	case "amount_off":
		var decoded OfferPhasesItemVariant3
		if err := json.Unmarshal(data, &decoded); err != nil {
			return err
		}
		value.Value = decoded
		return nil
	case "fixed_price":
		var decoded OfferPhasesItemVariant4
		if err := json.Unmarshal(data, &decoded); err != nil {
			return err
		}
		value.Value = decoded
		return nil
	}
	if hasJSONFields(fields, "type", "duration_cycles", "percentage") {
		var decoded OfferPhasesItemVariant2
		if err := json.Unmarshal(data, &decoded); err != nil {
			return err
		}
		value.Value = decoded
		return nil
	}
	if hasJSONFields(fields, "type", "duration_cycles", "amounts") {
		var decoded OfferPhasesItemVariant3
		if err := json.Unmarshal(data, &decoded); err != nil {
			return err
		}
		value.Value = decoded
		return nil
	}
	if hasJSONFields(fields, "type", "duration_cycles", "prices") {
		var decoded OfferPhasesItemVariant4
		if err := json.Unmarshal(data, &decoded); err != nil {
			return err
		}
		value.Value = decoded
		return nil
	}
	if hasJSONFields(fields, "type", "duration_days") {
		var decoded OfferPhasesItemVariant1
		if err := json.Unmarshal(data, &decoded); err != nil {
			return err
		}
		value.Value = decoded
		return nil
	}
	var decoded OfferPhasesItemVariant1
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	value.Value = decoded
	return nil
}

func (value OfferPhasesItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(value.Value)
}

func (value OfferPhasesItem) AsOfferPhasesItemVariant1() (OfferPhasesItemVariant1, bool) {
	decoded, ok := value.Value.(OfferPhasesItemVariant1)
	return decoded, ok
}

func (value OfferPhasesItem) AsOfferPhasesItemVariant2() (OfferPhasesItemVariant2, bool) {
	decoded, ok := value.Value.(OfferPhasesItemVariant2)
	return decoded, ok
}

func (value OfferPhasesItem) AsOfferPhasesItemVariant3() (OfferPhasesItemVariant3, bool) {
	decoded, ok := value.Value.(OfferPhasesItemVariant3)
	return decoded, ok
}

func (value OfferPhasesItem) AsOfferPhasesItemVariant4() (OfferPhasesItemVariant4, bool) {
	decoded, ok := value.Value.(OfferPhasesItemVariant4)
	return decoded, ok
}

type OfferPhasesItemVariant1 struct {
	Type         string `json:"type"`
	DurationDays int    `json:"duration_days"`
}

type OfferPhasesItemVariant2 struct {
	Type           string `json:"type"`
	DurationCycles int    `json:"duration_cycles"`
	Percentage     int    `json:"percentage"`
}

type OfferPhasesItemVariant3 struct {
	Type           string                               `json:"type"`
	DurationCycles int                                  `json:"duration_cycles"`
	Amounts        []OfferPhasesItemVariant3AmountsItem `json:"amounts"`
}

type OfferPhasesItemVariant3AmountsItem struct {
	Currency string `json:"currency"`
	Amount   int    `json:"amount"`
}

type OfferPhasesItemVariant4 struct {
	Type           string                              `json:"type"`
	DurationCycles int                                 `json:"duration_cycles"`
	Prices         []OfferPhasesItemVariant4PricesItem `json:"prices"`
}

type OfferPhasesItemVariant4PricesItem struct {
	Currency string `json:"currency"`
	Amount   int    `json:"amount"`
}

type OffersListResult struct {
	Object     string  `json:"object"`
	Data       []Offer `json:"data"`
	HasMore    bool    `json:"has_more"`
	NextCursor *string `json:"next_cursor,omitempty"`
}

type Payment struct {
	ID             string         `json:"id"`
	CustomerID     *string        `json:"customer_id"`
	Kind           string         `json:"kind"`
	Status         string         `json:"status"`
	Provider       string         `json:"provider"`
	AmountSubtotal int            `json:"amount_subtotal"`
	TaxAmount      int            `json:"tax_amount"`
	AmountTotal    int            `json:"amount_total"`
	Currency       string         `json:"currency"`
	Description    string         `json:"description"`
	Metadata       map[string]any `json:"metadata"`
	URL            *string        `json:"url"`
	ExpiresAt      *string        `json:"expires_at"`
	CreatedAt      string         `json:"created_at"`
	UpdatedAt      string         `json:"updated_at"`
	Object         string         `json:"object"`
	Livemode       bool           `json:"livemode"`
}

type PaymentMethodUpdateCheckout struct {
	CheckoutURL string `json:"checkout_url"`
	Object      string `json:"object"`
	Livemode    bool   `json:"livemode"`
}

type PaymentsListResult struct {
	Object     string    `json:"object"`
	Data       []Payment `json:"data"`
	HasMore    bool      `json:"has_more"`
	NextCursor *string   `json:"next_cursor,omitempty"`
}

type Payout struct {
	ID                 string  `json:"id"`
	Status             string  `json:"status"`
	Amount             int     `json:"amount"`
	Fee                int     `json:"fee"`
	NetAmount          int     `json:"net_amount"`
	Currency           string  `json:"currency"`
	Description        *string `json:"description"`
	ProviderTransferID string  `json:"provider_transfer_id"`
	CreatedAt          string  `json:"created_at"`
	Object             string  `json:"object"`
	Livemode           bool    `json:"livemode"`
}

type PayoutBankAccount struct {
	ID                        string  `json:"id"`
	ProviderExternalAccountID *string `json:"provider_external_account_id"`
	HolderName                string  `json:"holder_name"`
	Last4                     string  `json:"last4"`
	BankName                  *string `json:"bank_name"`
	Country                   string  `json:"country"`
	Currency                  string  `json:"currency"`
	AccountType               *string `json:"account_type"`
	IsDefault                 bool    `json:"is_default"`
	Status                    string  `json:"status"`
	CreatedAt                 string  `json:"created_at"`
	Object                    string  `json:"object"`
	Livemode                  bool    `json:"livemode"`
}

type PayoutVerification struct {
	Value any
}

func (value *PayoutVerification) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	var discriminator string
	if raw, ok := fields["outcome"]; ok {
		if err := json.Unmarshal(raw, &discriminator); err != nil {
			return err
		}
	}
	switch discriminator {
	case "existing":
		var decoded PayoutVerificationVariant1
		if err := json.Unmarshal(data, &decoded); err != nil {
			return err
		}
		value.Value = decoded
		return nil
	case "created":
		var decoded PayoutVerificationVariant2
		if err := json.Unmarshal(data, &decoded); err != nil {
			return err
		}
		value.Value = decoded
		return nil
	}
	if hasJSONFields(fields, "provider_account_id", "status", "transfers_enabled", "outcome", "business_type", "country", "object", "livemode") {
		var decoded PayoutVerificationVariant2
		if err := json.Unmarshal(data, &decoded); err != nil {
			return err
		}
		value.Value = decoded
		return nil
	}
	if hasJSONFields(fields, "provider_account_id", "status", "transfers_enabled", "outcome", "object", "livemode") {
		var decoded PayoutVerificationVariant1
		if err := json.Unmarshal(data, &decoded); err != nil {
			return err
		}
		value.Value = decoded
		return nil
	}
	var decoded PayoutVerificationVariant1
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	value.Value = decoded
	return nil
}

func (value PayoutVerification) MarshalJSON() ([]byte, error) {
	return json.Marshal(value.Value)
}

func (value PayoutVerification) AsPayoutVerificationVariant1() (PayoutVerificationVariant1, bool) {
	decoded, ok := value.Value.(PayoutVerificationVariant1)
	return decoded, ok
}

func (value PayoutVerification) AsPayoutVerificationVariant2() (PayoutVerificationVariant2, bool) {
	decoded, ok := value.Value.(PayoutVerificationVariant2)
	return decoded, ok
}

type PayoutVerificationVariant1 struct {
	ProviderAccountID string `json:"provider_account_id"`
	Status            string `json:"status"`
	TransfersEnabled  bool   `json:"transfers_enabled"`
	Outcome           string `json:"outcome"`
	Object            string `json:"object"`
	Livemode          bool   `json:"livemode"`
}

type PayoutVerificationVariant2 struct {
	ProviderAccountID string `json:"provider_account_id"`
	Status            string `json:"status"`
	TransfersEnabled  bool   `json:"transfers_enabled"`
	Outcome           string `json:"outcome"`
	BusinessType      string `json:"business_type"`
	Country           string `json:"country"`
	Object            string `json:"object"`
	Livemode          bool   `json:"livemode"`
}

type Plan struct {
	ID                string                  `json:"id"`
	Name              string                  `json:"name"`
	Code              string                  `json:"code"`
	Description       *string                 `json:"description"`
	ConsumptionModel  *ConsumptionModel       `json:"consumption_model"`
	IsPublic          bool                    `json:"is_public"`
	IsDefault         bool                    `json:"is_default"`
	IsFree            bool                    `json:"is_free"`
	BlockOnExhaustion *bool                   `json:"block_on_exhaustion"`
	SortOrder         int                     `json:"sort_order"`
	PlanGroupID       *string                 `json:"plan_group_id"`
	Metadata          map[string]any          `json:"metadata"`
	CreatedAt         string                  `json:"created_at"`
	UpdatedAt         string                  `json:"updated_at"`
	Features          []PlanFeaturesItem      `json:"features"`
	Prices            []PlanPricesItem        `json:"prices"`
	ExchangeRates     []PlanExchangeRatesItem `json:"exchange_rates"`
	Object            string                  `json:"object"`
	Livemode          bool                    `json:"livemode"`
}

type PlanChange struct {
	Value any
}

func (value *PlanChange) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	var discriminator string
	if raw, ok := fields["outcome"]; ok {
		if err := json.Unmarshal(raw, &discriminator); err != nil {
			return err
		}
	}
	switch discriminator {
	case "requires_checkout":
		var decoded PlanChangeVariant1
		if err := json.Unmarshal(data, &decoded); err != nil {
			return err
		}
		value.Value = decoded
		return nil
	case "scheduled":
		var decoded PlanChangeVariant2
		if err := json.Unmarshal(data, &decoded); err != nil {
			return err
		}
		value.Value = decoded
		return nil
	case "completed":
		var decoded PlanChangeVariant3
		if err := json.Unmarshal(data, &decoded); err != nil {
			return err
		}
		value.Value = decoded
		return nil
	}
	if hasJSONFields(fields, "outcome", "id", "scheduled", "customer_id", "previous_plan", "current_plan", "billing_interval", "billing", "object", "livemode") {
		var decoded PlanChangeVariant3
		if err := json.Unmarshal(data, &decoded); err != nil {
			return err
		}
		value.Value = decoded
		return nil
	}
	if hasJSONFields(fields, "outcome", "id", "scheduled", "scheduled_for", "change_type", "customer_id", "object", "livemode") {
		var decoded PlanChangeVariant2
		if err := json.Unmarshal(data, &decoded); err != nil {
			return err
		}
		value.Value = decoded
		return nil
	}
	if hasJSONFields(fields, "outcome", "requires_checkout", "checkout_url", "object", "livemode") {
		var decoded PlanChangeVariant1
		if err := json.Unmarshal(data, &decoded); err != nil {
			return err
		}
		value.Value = decoded
		return nil
	}
	var decoded PlanChangeVariant1
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	value.Value = decoded
	return nil
}

func (value PlanChange) MarshalJSON() ([]byte, error) {
	return json.Marshal(value.Value)
}

func (value PlanChange) AsPlanChangeVariant1() (PlanChangeVariant1, bool) {
	decoded, ok := value.Value.(PlanChangeVariant1)
	return decoded, ok
}

func (value PlanChange) AsPlanChangeVariant2() (PlanChangeVariant2, bool) {
	decoded, ok := value.Value.(PlanChangeVariant2)
	return decoded, ok
}

func (value PlanChange) AsPlanChangeVariant3() (PlanChangeVariant3, bool) {
	decoded, ok := value.Value.(PlanChangeVariant3)
	return decoded, ok
}

type PlanChangeVariant1 struct {
	Outcome          string                              `json:"outcome"`
	RequiresCheckout any                                 `json:"requires_checkout"`
	CheckoutURL      string                              `json:"checkout_url"`
	OfferApplication *PlanChangeVariant1OfferApplication `json:"offer_application,omitempty"`
	Object           string                              `json:"object"`
	Livemode         bool                                `json:"livemode"`
}

type PlanChangeVariant1OfferApplication struct {
	ID             string                                         `json:"id"`
	OfferID        string                                         `json:"offer_id"`
	Name           string                                         `json:"name"`
	Currency       string                                         `json:"currency"`
	Subtotal       int                                            `json:"subtotal"`
	DiscountAmount int                                            `json:"discount_amount"`
	Total          int                                            `json:"total"`
	Phases         []PlanChangeVariant1OfferApplicationPhasesItem `json:"phases"`
}

type PlanChangeVariant1OfferApplicationPhasesItem struct {
	Value any
}

func (value *PlanChangeVariant1OfferApplicationPhasesItem) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	var discriminator string
	if raw, ok := fields["type"]; ok {
		if err := json.Unmarshal(raw, &discriminator); err != nil {
			return err
		}
	}
	switch discriminator {
	case "percentage":
		var decoded PlanChangeVariant1OfferApplicationPhasesItemVariant1
		if err := json.Unmarshal(data, &decoded); err != nil {
			return err
		}
		value.Value = decoded
		return nil
	case "amount_off":
		var decoded PlanChangeVariant1OfferApplicationPhasesItemVariant2
		if err := json.Unmarshal(data, &decoded); err != nil {
			return err
		}
		value.Value = decoded
		return nil
	case "fixed_price":
		var decoded PlanChangeVariant1OfferApplicationPhasesItemVariant3
		if err := json.Unmarshal(data, &decoded); err != nil {
			return err
		}
		value.Value = decoded
		return nil
	}
	if hasJSONFields(fields, "type", "duration_cycles", "starts_at", "ends_at", "percentage") {
		var decoded PlanChangeVariant1OfferApplicationPhasesItemVariant1
		if err := json.Unmarshal(data, &decoded); err != nil {
			return err
		}
		value.Value = decoded
		return nil
	}
	if hasJSONFields(fields, "type", "duration_cycles", "starts_at", "ends_at", "amount") {
		var decoded PlanChangeVariant1OfferApplicationPhasesItemVariant2
		if err := json.Unmarshal(data, &decoded); err != nil {
			return err
		}
		value.Value = decoded
		return nil
	}
	if hasJSONFields(fields, "type", "duration_cycles", "starts_at", "ends_at", "price") {
		var decoded PlanChangeVariant1OfferApplicationPhasesItemVariant3
		if err := json.Unmarshal(data, &decoded); err != nil {
			return err
		}
		value.Value = decoded
		return nil
	}
	var decoded PlanChangeVariant1OfferApplicationPhasesItemVariant1
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	value.Value = decoded
	return nil
}

func (value PlanChangeVariant1OfferApplicationPhasesItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(value.Value)
}

func (value PlanChangeVariant1OfferApplicationPhasesItem) AsPlanChangeVariant1OfferApplicationPhasesItemVariant1() (PlanChangeVariant1OfferApplicationPhasesItemVariant1, bool) {
	decoded, ok := value.Value.(PlanChangeVariant1OfferApplicationPhasesItemVariant1)
	return decoded, ok
}

func (value PlanChangeVariant1OfferApplicationPhasesItem) AsPlanChangeVariant1OfferApplicationPhasesItemVariant2() (PlanChangeVariant1OfferApplicationPhasesItemVariant2, bool) {
	decoded, ok := value.Value.(PlanChangeVariant1OfferApplicationPhasesItemVariant2)
	return decoded, ok
}

func (value PlanChangeVariant1OfferApplicationPhasesItem) AsPlanChangeVariant1OfferApplicationPhasesItemVariant3() (PlanChangeVariant1OfferApplicationPhasesItemVariant3, bool) {
	decoded, ok := value.Value.(PlanChangeVariant1OfferApplicationPhasesItemVariant3)
	return decoded, ok
}

type PlanChangeVariant1OfferApplicationPhasesItemVariant1 struct {
	Type           string  `json:"type"`
	DurationCycles int     `json:"duration_cycles"`
	StartsAt       *string `json:"starts_at"`
	EndsAt         *string `json:"ends_at"`
	Percentage     int     `json:"percentage"`
}

type PlanChangeVariant1OfferApplicationPhasesItemVariant2 struct {
	Type           string  `json:"type"`
	DurationCycles int     `json:"duration_cycles"`
	StartsAt       *string `json:"starts_at"`
	EndsAt         *string `json:"ends_at"`
	Amount         int     `json:"amount"`
}

type PlanChangeVariant1OfferApplicationPhasesItemVariant3 struct {
	Type           string  `json:"type"`
	DurationCycles int     `json:"duration_cycles"`
	StartsAt       *string `json:"starts_at"`
	EndsAt         *string `json:"ends_at"`
	Price          int     `json:"price"`
}

type PlanChangeVariant2 struct {
	Outcome            string                              `json:"outcome"`
	ID                 string                              `json:"id"`
	Scheduled          any                                 `json:"scheduled"`
	ScheduledFor       string                              `json:"scheduled_for"`
	ChangeType         string                              `json:"change_type"`
	CustomerID         string                              `json:"customer_id"`
	NewPlanID          *string                             `json:"new_plan_id,omitempty"`
	NewPlanName        *string                             `json:"new_plan_name,omitempty"`
	NewBillingInterval *string                             `json:"new_billing_interval,omitempty"`
	SeatLimitWarning   *PlanChangeVariant2SeatLimitWarning `json:"seat_limit_warning,omitempty"`
	Object             string                              `json:"object"`
	Livemode           bool                                `json:"livemode"`
}

type PlanChangeVariant2SeatLimitWarning struct {
	FeatureCode   string `json:"feature_code"`
	FeatureName   string `json:"feature_name"`
	CurrentSeats  int    `json:"current_seats"`
	Included      int    `json:"included"`
	NewPlanName   string `json:"new_plan_name"`
	EffectiveDate string `json:"effective_date"`
}

type PlanChangeVariant3 struct {
	Outcome          string                              `json:"outcome"`
	ID               string                              `json:"id"`
	Scheduled        any                                 `json:"scheduled"`
	CustomerID       string                              `json:"customer_id"`
	PreviousPlan     PlanChangeVariant3PreviousPlan      `json:"previous_plan"`
	CurrentPlan      PlanChangeVariant3CurrentPlan       `json:"current_plan"`
	BillingInterval  string                              `json:"billing_interval"`
	Billing          PlanChangeVariant3Billing           `json:"billing"`
	InvoiceID        *string                             `json:"invoice_id,omitempty"`
	OfferApplication *PlanChangeVariant3OfferApplication `json:"offer_application,omitempty"`
	Object           string                              `json:"object"`
	Livemode         bool                                `json:"livemode"`
}

type PlanChangeVariant3Billing struct {
	Credit                 int `json:"credit"`
	CreditsApplied         int `json:"credits_applied"`
	Charge                 int `json:"charge"`
	TaxAmount              int `json:"tax_amount"`
	NetAmount              int `json:"net_amount"`
	TotalCharged           int `json:"total_charged"`
	RemainingCreditBalance int `json:"remaining_credit_balance"`
}

type PlanChangeVariant3CurrentPlan struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type PlanChangeVariant3OfferApplication struct {
	ID             string                                         `json:"id"`
	OfferID        string                                         `json:"offer_id"`
	Name           string                                         `json:"name"`
	Currency       string                                         `json:"currency"`
	Subtotal       int                                            `json:"subtotal"`
	DiscountAmount int                                            `json:"discount_amount"`
	Total          int                                            `json:"total"`
	Phases         []PlanChangeVariant3OfferApplicationPhasesItem `json:"phases"`
}

type PlanChangeVariant3OfferApplicationPhasesItem struct {
	Value any
}

func (value *PlanChangeVariant3OfferApplicationPhasesItem) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	var discriminator string
	if raw, ok := fields["type"]; ok {
		if err := json.Unmarshal(raw, &discriminator); err != nil {
			return err
		}
	}
	switch discriminator {
	case "percentage":
		var decoded PlanChangeVariant3OfferApplicationPhasesItemVariant1
		if err := json.Unmarshal(data, &decoded); err != nil {
			return err
		}
		value.Value = decoded
		return nil
	case "amount_off":
		var decoded PlanChangeVariant3OfferApplicationPhasesItemVariant2
		if err := json.Unmarshal(data, &decoded); err != nil {
			return err
		}
		value.Value = decoded
		return nil
	case "fixed_price":
		var decoded PlanChangeVariant3OfferApplicationPhasesItemVariant3
		if err := json.Unmarshal(data, &decoded); err != nil {
			return err
		}
		value.Value = decoded
		return nil
	}
	if hasJSONFields(fields, "type", "duration_cycles", "starts_at", "ends_at", "percentage") {
		var decoded PlanChangeVariant3OfferApplicationPhasesItemVariant1
		if err := json.Unmarshal(data, &decoded); err != nil {
			return err
		}
		value.Value = decoded
		return nil
	}
	if hasJSONFields(fields, "type", "duration_cycles", "starts_at", "ends_at", "amount") {
		var decoded PlanChangeVariant3OfferApplicationPhasesItemVariant2
		if err := json.Unmarshal(data, &decoded); err != nil {
			return err
		}
		value.Value = decoded
		return nil
	}
	if hasJSONFields(fields, "type", "duration_cycles", "starts_at", "ends_at", "price") {
		var decoded PlanChangeVariant3OfferApplicationPhasesItemVariant3
		if err := json.Unmarshal(data, &decoded); err != nil {
			return err
		}
		value.Value = decoded
		return nil
	}
	var decoded PlanChangeVariant3OfferApplicationPhasesItemVariant1
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	value.Value = decoded
	return nil
}

func (value PlanChangeVariant3OfferApplicationPhasesItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(value.Value)
}

func (value PlanChangeVariant3OfferApplicationPhasesItem) AsPlanChangeVariant3OfferApplicationPhasesItemVariant1() (PlanChangeVariant3OfferApplicationPhasesItemVariant1, bool) {
	decoded, ok := value.Value.(PlanChangeVariant3OfferApplicationPhasesItemVariant1)
	return decoded, ok
}

func (value PlanChangeVariant3OfferApplicationPhasesItem) AsPlanChangeVariant3OfferApplicationPhasesItemVariant2() (PlanChangeVariant3OfferApplicationPhasesItemVariant2, bool) {
	decoded, ok := value.Value.(PlanChangeVariant3OfferApplicationPhasesItemVariant2)
	return decoded, ok
}

func (value PlanChangeVariant3OfferApplicationPhasesItem) AsPlanChangeVariant3OfferApplicationPhasesItemVariant3() (PlanChangeVariant3OfferApplicationPhasesItemVariant3, bool) {
	decoded, ok := value.Value.(PlanChangeVariant3OfferApplicationPhasesItemVariant3)
	return decoded, ok
}

type PlanChangeVariant3OfferApplicationPhasesItemVariant1 struct {
	Type           string  `json:"type"`
	DurationCycles int     `json:"duration_cycles"`
	StartsAt       *string `json:"starts_at"`
	EndsAt         *string `json:"ends_at"`
	Percentage     int     `json:"percentage"`
}

type PlanChangeVariant3OfferApplicationPhasesItemVariant2 struct {
	Type           string  `json:"type"`
	DurationCycles int     `json:"duration_cycles"`
	StartsAt       *string `json:"starts_at"`
	EndsAt         *string `json:"ends_at"`
	Amount         int     `json:"amount"`
}

type PlanChangeVariant3OfferApplicationPhasesItemVariant3 struct {
	Type           string  `json:"type"`
	DurationCycles int     `json:"duration_cycles"`
	StartsAt       *string `json:"starts_at"`
	EndsAt         *string `json:"ends_at"`
	Price          int     `json:"price"`
}

type PlanChangeVariant3PreviousPlan struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type PlanExchangeRatesItem struct {
	Currency     string  `json:"currency"`
	ExchangeRate float64 `json:"exchange_rate"`
}

type PlanFeature struct {
	PlanID         string             `json:"plan_id"`
	FeatureID      string             `json:"feature_id"`
	Enabled        bool               `json:"enabled"`
	IncludedAmount int                `json:"included_amount"`
	Unlimited      bool               `json:"unlimited"`
	Overage        PlanFeatureOverage `json:"overage"`
	CreditsPerUnit *int               `json:"credits_per_unit"`
	PricingMode    string             `json:"pricing_mode"`
	Margin         *int               `json:"margin"`
	Object         string             `json:"object"`
	Livemode       bool               `json:"livemode"`
}

type PlanFeatureOverage struct {
	Enabled   bool `json:"enabled"`
	UnitPrice int  `json:"unit_price"`
}

type PlanFeaturesItem struct {
	Code           string                               `json:"code"`
	Name           string                               `json:"name"`
	Type           FeatureType                          `json:"type"`
	UnitName       *string                              `json:"unit_name"`
	Enabled        bool                                 `json:"enabled"`
	IncludedAmount *int                                 `json:"included_amount"`
	Unlimited      bool                                 `json:"unlimited"`
	Overage        *PlanFeaturesItemOverage             `json:"overage"`
	RegionalPrices []PlanFeaturesItemRegionalPricesItem `json:"regional_prices"`
}

type PlanFeaturesItemOverage struct {
	Enabled   bool    `json:"enabled"`
	Model     *string `json:"model"`
	UnitPrice *int    `json:"unit_price"`
}

type PlanFeaturesItemRegionalPricesItem struct {
	Currency         string `json:"currency"`
	OverageUnitPrice *int   `json:"overage_unit_price"`
	AutoSynced       bool   `json:"auto_synced"`
}

type PlanGroup struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
	IsPublic    bool    `json:"is_public"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
	Object      string  `json:"object"`
	Livemode    bool    `json:"livemode"`
}

type PlanGroupDetail struct {
	ID          string                     `json:"id"`
	Name        string                     `json:"name"`
	Description *string                    `json:"description"`
	IsPublic    bool                       `json:"is_public"`
	CreatedAt   string                     `json:"created_at"`
	UpdatedAt   string                     `json:"updated_at"`
	Plans       []PlanGroupDetailPlansItem `json:"plans"`
	Object      string                     `json:"object"`
	Livemode    bool                       `json:"livemode"`
}

type PlanGroupDetailPlansItem struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	SortOrder int    `json:"sort_order"`
}

type PlanGroupsListResult struct {
	Object     string      `json:"object"`
	Data       []PlanGroup `json:"data"`
	HasMore    bool        `json:"has_more"`
	NextCursor *string     `json:"next_cursor,omitempty"`
}

type PlanPrice struct {
	ID                  string                      `json:"id"`
	PlanID              string                      `json:"plan_id"`
	BillingInterval     BillingInterval             `json:"billing_interval"`
	Price               int                         `json:"price"`
	IsDefault           bool                        `json:"is_default"`
	TrialDays           int                         `json:"trial_days"`
	IncludedBalance     *int                        `json:"included_balance"`
	IncludedCredits     *int                        `json:"included_credits"`
	OfferID             *string                     `json:"offer_id"`
	InheritsFromPriceID *string                     `json:"inherits_from_price_id"`
	Metadata            map[string]any              `json:"metadata"`
	MarketPrices        []PlanPriceMarketPricesItem `json:"market_prices"`
	CreatedAt           string                      `json:"created_at"`
	UpdatedAt           string                      `json:"updated_at"`
	Object              string                      `json:"object"`
	Livemode            bool                        `json:"livemode"`
}

type PlanPriceMarketPricesItem struct {
	MarketGroupID string `json:"market_group_id"`
	Currency      string `json:"currency"`
	Price         int    `json:"price"`
}

type PlanPricesItem struct {
	ID                  string                             `json:"id"`
	BillingInterval     BillingInterval                    `json:"billing_interval"`
	Price               int                                `json:"price"`
	IsDefault           bool                               `json:"is_default"`
	TrialDays           int                                `json:"trial_days"`
	IncludedBalance     *int                               `json:"included_balance"`
	IncludedCredits     *int                               `json:"included_credits"`
	OfferID             *string                            `json:"offer_id"`
	InheritsFromPriceID *string                            `json:"inherits_from_price_id"`
	Metadata            map[string]any                     `json:"metadata"`
	MarketPrices        []PlanPricesItemMarketPricesItem   `json:"market_prices"`
	RegionalPrices      []PlanPricesItemRegionalPricesItem `json:"regional_prices"`
}

type PlanPricesItemMarketPricesItem struct {
	MarketGroupID string `json:"market_group_id"`
	Currency      string `json:"currency"`
	Price         int    `json:"price"`
}

type PlanPricesItemRegionalPricesItem struct {
	Currency        string `json:"currency"`
	Price           int    `json:"price"`
	IncludedBalance *int   `json:"included_balance"`
	AutoSynced      bool   `json:"auto_synced"`
}

type PlanRegionalPricing struct {
	PriceID   string                             `json:"price_id"`
	Overrides []PlanRegionalPricingOverridesItem `json:"overrides"`
	Object    string                             `json:"object"`
	Livemode  bool                               `json:"livemode"`
}

type PlanRegionalPricingOverridesItem struct {
	Currency        string `json:"currency"`
	Price           int    `json:"price"`
	IncludedBalance *int   `json:"included_balance,omitempty"`
}

type PlanRegionalPricingResult struct {
	PlanID             string  `json:"plan_id"`
	Currency           string  `json:"currency"`
	ExchangeRate       float64 `json:"exchange_rate"`
	PricesConfigured   int     `json:"prices_configured"`
	FeaturesConfigured int     `json:"features_configured"`
	Object             string  `json:"object"`
	Livemode           bool    `json:"livemode"`
}

type PlansListResult struct {
	Object     string  `json:"object"`
	Data       []Plan  `json:"data"`
	HasMore    bool    `json:"has_more"`
	NextCursor *string `json:"next_cursor,omitempty"`
}

type PortalAccess struct {
	PortalURL string `json:"portal_url"`
	Object    string `json:"object"`
	Livemode  bool   `json:"livemode"`
}

type PreviewChange struct {
	Currency          string                         `json:"currency"`
	CurrentPlanCredit int                            `json:"current_plan_credit"`
	NewPlanCharge     int                            `json:"new_plan_charge"`
	EstimatedTotal    int                            `json:"estimated_total"`
	EffectiveDate     string                         `json:"effective_date"`
	DaysRemaining     int                            `json:"days_remaining"`
	TotalDays         int                            `json:"total_days"`
	IsUpgrade         bool                           `json:"is_upgrade"`
	OfferApplication  *PreviewChangeOfferApplication `json:"offer_application,omitempty"`
	Object            string                         `json:"object"`
	Livemode          bool                           `json:"livemode"`
}

type PreviewChangeOfferApplication struct {
	ID             string                                    `json:"id"`
	OfferID        string                                    `json:"offer_id"`
	Name           string                                    `json:"name"`
	Currency       string                                    `json:"currency"`
	Subtotal       int                                       `json:"subtotal"`
	DiscountAmount int                                       `json:"discount_amount"`
	Total          int                                       `json:"total"`
	Phases         []PreviewChangeOfferApplicationPhasesItem `json:"phases"`
}

type PreviewChangeOfferApplicationPhasesItem struct {
	Value any
}

func (value *PreviewChangeOfferApplicationPhasesItem) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	var discriminator string
	if raw, ok := fields["type"]; ok {
		if err := json.Unmarshal(raw, &discriminator); err != nil {
			return err
		}
	}
	switch discriminator {
	case "percentage":
		var decoded PreviewChangeOfferApplicationPhasesItemVariant1
		if err := json.Unmarshal(data, &decoded); err != nil {
			return err
		}
		value.Value = decoded
		return nil
	case "amount_off":
		var decoded PreviewChangeOfferApplicationPhasesItemVariant2
		if err := json.Unmarshal(data, &decoded); err != nil {
			return err
		}
		value.Value = decoded
		return nil
	case "fixed_price":
		var decoded PreviewChangeOfferApplicationPhasesItemVariant3
		if err := json.Unmarshal(data, &decoded); err != nil {
			return err
		}
		value.Value = decoded
		return nil
	}
	if hasJSONFields(fields, "type", "duration_cycles", "starts_at", "ends_at", "percentage") {
		var decoded PreviewChangeOfferApplicationPhasesItemVariant1
		if err := json.Unmarshal(data, &decoded); err != nil {
			return err
		}
		value.Value = decoded
		return nil
	}
	if hasJSONFields(fields, "type", "duration_cycles", "starts_at", "ends_at", "amount") {
		var decoded PreviewChangeOfferApplicationPhasesItemVariant2
		if err := json.Unmarshal(data, &decoded); err != nil {
			return err
		}
		value.Value = decoded
		return nil
	}
	if hasJSONFields(fields, "type", "duration_cycles", "starts_at", "ends_at", "price") {
		var decoded PreviewChangeOfferApplicationPhasesItemVariant3
		if err := json.Unmarshal(data, &decoded); err != nil {
			return err
		}
		value.Value = decoded
		return nil
	}
	var decoded PreviewChangeOfferApplicationPhasesItemVariant1
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	value.Value = decoded
	return nil
}

func (value PreviewChangeOfferApplicationPhasesItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(value.Value)
}

func (value PreviewChangeOfferApplicationPhasesItem) AsPreviewChangeOfferApplicationPhasesItemVariant1() (PreviewChangeOfferApplicationPhasesItemVariant1, bool) {
	decoded, ok := value.Value.(PreviewChangeOfferApplicationPhasesItemVariant1)
	return decoded, ok
}

func (value PreviewChangeOfferApplicationPhasesItem) AsPreviewChangeOfferApplicationPhasesItemVariant2() (PreviewChangeOfferApplicationPhasesItemVariant2, bool) {
	decoded, ok := value.Value.(PreviewChangeOfferApplicationPhasesItemVariant2)
	return decoded, ok
}

func (value PreviewChangeOfferApplicationPhasesItem) AsPreviewChangeOfferApplicationPhasesItemVariant3() (PreviewChangeOfferApplicationPhasesItemVariant3, bool) {
	decoded, ok := value.Value.(PreviewChangeOfferApplicationPhasesItemVariant3)
	return decoded, ok
}

type PreviewChangeOfferApplicationPhasesItemVariant1 struct {
	Type           string  `json:"type"`
	DurationCycles int     `json:"duration_cycles"`
	StartsAt       *string `json:"starts_at"`
	EndsAt         *string `json:"ends_at"`
	Percentage     int     `json:"percentage"`
}

type PreviewChangeOfferApplicationPhasesItemVariant2 struct {
	Type           string  `json:"type"`
	DurationCycles int     `json:"duration_cycles"`
	StartsAt       *string `json:"starts_at"`
	EndsAt         *string `json:"ends_at"`
	Amount         int     `json:"amount"`
}

type PreviewChangeOfferApplicationPhasesItemVariant3 struct {
	Type           string  `json:"type"`
	DurationCycles int     `json:"duration_cycles"`
	StartsAt       *string `json:"starts_at"`
	EndsAt         *string `json:"ends_at"`
	Price          int     `json:"price"`
}

type PricingListMarketGroupsResult struct {
	Object     string        `json:"object"`
	Data       []MarketGroup `json:"data"`
	HasMore    bool          `json:"has_more"`
	NextCursor *string       `json:"next_cursor,omitempty"`
}

type PromoCode struct {
	ID              string           `json:"id"`
	Code            string           `json:"code"`
	OfferID         string           `json:"offer_id"`
	BillingInterval *BillingInterval `json:"billing_interval"`
	MaxRedemptions  *int             `json:"max_redemptions"`
	ExpiresAt       *string          `json:"expires_at"`
	IsActive        bool             `json:"is_active"`
	RedemptionCount int              `json:"redemption_count"`
	CreatedAt       string           `json:"created_at"`
	UpdatedAt       string           `json:"updated_at"`
	Object          string           `json:"object"`
	Livemode        bool             `json:"livemode"`
}

type PromoCodesListResult struct {
	Object     string      `json:"object"`
	Data       []PromoCode `json:"data"`
	HasMore    bool        `json:"has_more"`
	NextCursor *string     `json:"next_cursor,omitempty"`
}

type QuotaGetAllResult struct {
	Object     string       `json:"object"`
	Data       []UsageQuota `json:"data"`
	HasMore    bool         `json:"has_more"`
	NextCursor *string      `json:"next_cursor,omitempty"`
}

type ReactivatedSubscription struct {
	SubscriptionID   string                                   `json:"subscription_id"`
	InvoiceID        string                                   `json:"invoice_id"`
	Status           string                                   `json:"status"`
	OfferApplication *ReactivatedSubscriptionOfferApplication `json:"offer_application,omitempty"`
	Object           string                                   `json:"object"`
	Livemode         bool                                     `json:"livemode"`
}

type ReactivatedSubscriptionOfferApplication struct {
	ID             string                                              `json:"id"`
	OfferID        string                                              `json:"offer_id"`
	Name           string                                              `json:"name"`
	Currency       string                                              `json:"currency"`
	Subtotal       int                                                 `json:"subtotal"`
	DiscountAmount int                                                 `json:"discount_amount"`
	Total          int                                                 `json:"total"`
	Phases         []ReactivatedSubscriptionOfferApplicationPhasesItem `json:"phases"`
}

type ReactivatedSubscriptionOfferApplicationPhasesItem struct {
	Value any
}

func (value *ReactivatedSubscriptionOfferApplicationPhasesItem) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	var discriminator string
	if raw, ok := fields["type"]; ok {
		if err := json.Unmarshal(raw, &discriminator); err != nil {
			return err
		}
	}
	switch discriminator {
	case "percentage":
		var decoded ReactivatedSubscriptionOfferApplicationPhasesItemVariant1
		if err := json.Unmarshal(data, &decoded); err != nil {
			return err
		}
		value.Value = decoded
		return nil
	case "amount_off":
		var decoded ReactivatedSubscriptionOfferApplicationPhasesItemVariant2
		if err := json.Unmarshal(data, &decoded); err != nil {
			return err
		}
		value.Value = decoded
		return nil
	case "fixed_price":
		var decoded ReactivatedSubscriptionOfferApplicationPhasesItemVariant3
		if err := json.Unmarshal(data, &decoded); err != nil {
			return err
		}
		value.Value = decoded
		return nil
	}
	if hasJSONFields(fields, "type", "duration_cycles", "starts_at", "ends_at", "percentage") {
		var decoded ReactivatedSubscriptionOfferApplicationPhasesItemVariant1
		if err := json.Unmarshal(data, &decoded); err != nil {
			return err
		}
		value.Value = decoded
		return nil
	}
	if hasJSONFields(fields, "type", "duration_cycles", "starts_at", "ends_at", "amount") {
		var decoded ReactivatedSubscriptionOfferApplicationPhasesItemVariant2
		if err := json.Unmarshal(data, &decoded); err != nil {
			return err
		}
		value.Value = decoded
		return nil
	}
	if hasJSONFields(fields, "type", "duration_cycles", "starts_at", "ends_at", "price") {
		var decoded ReactivatedSubscriptionOfferApplicationPhasesItemVariant3
		if err := json.Unmarshal(data, &decoded); err != nil {
			return err
		}
		value.Value = decoded
		return nil
	}
	var decoded ReactivatedSubscriptionOfferApplicationPhasesItemVariant1
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	value.Value = decoded
	return nil
}

func (value ReactivatedSubscriptionOfferApplicationPhasesItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(value.Value)
}

func (value ReactivatedSubscriptionOfferApplicationPhasesItem) AsReactivatedSubscriptionOfferApplicationPhasesItemVariant1() (ReactivatedSubscriptionOfferApplicationPhasesItemVariant1, bool) {
	decoded, ok := value.Value.(ReactivatedSubscriptionOfferApplicationPhasesItemVariant1)
	return decoded, ok
}

func (value ReactivatedSubscriptionOfferApplicationPhasesItem) AsReactivatedSubscriptionOfferApplicationPhasesItemVariant2() (ReactivatedSubscriptionOfferApplicationPhasesItemVariant2, bool) {
	decoded, ok := value.Value.(ReactivatedSubscriptionOfferApplicationPhasesItemVariant2)
	return decoded, ok
}

func (value ReactivatedSubscriptionOfferApplicationPhasesItem) AsReactivatedSubscriptionOfferApplicationPhasesItemVariant3() (ReactivatedSubscriptionOfferApplicationPhasesItemVariant3, bool) {
	decoded, ok := value.Value.(ReactivatedSubscriptionOfferApplicationPhasesItemVariant3)
	return decoded, ok
}

type ReactivatedSubscriptionOfferApplicationPhasesItemVariant1 struct {
	Type           string  `json:"type"`
	DurationCycles int     `json:"duration_cycles"`
	StartsAt       *string `json:"starts_at"`
	EndsAt         *string `json:"ends_at"`
	Percentage     int     `json:"percentage"`
}

type ReactivatedSubscriptionOfferApplicationPhasesItemVariant2 struct {
	Type           string  `json:"type"`
	DurationCycles int     `json:"duration_cycles"`
	StartsAt       *string `json:"starts_at"`
	EndsAt         *string `json:"ends_at"`
	Amount         int     `json:"amount"`
}

type ReactivatedSubscriptionOfferApplicationPhasesItemVariant3 struct {
	Type           string  `json:"type"`
	DurationCycles int     `json:"duration_cycles"`
	StartsAt       *string `json:"starts_at"`
	EndsAt         *string `json:"ends_at"`
	Price          int     `json:"price"`
}

type RecoveryLink struct {
	URL      string `json:"url"`
	Token    string `json:"token"`
	Object   string `json:"object"`
	Livemode bool   `json:"livemode"`
}

type Refund struct {
	ID            string  `json:"id"`
	TransactionID string  `json:"transaction_id"`
	Amount        int     `json:"amount"`
	Currency      string  `json:"currency"`
	ChargeID      *string `json:"charge_id"`
	Status        string  `json:"status"`
	Reason        *string `json:"reason"`
	Object        string  `json:"object"`
	Livemode      bool    `json:"livemode"`
}

type RemovedPlanFeature struct {
	ID       string `json:"id"`
	Removed  any    `json:"removed"`
	Object   string `json:"object"`
	Livemode bool   `json:"livemode"`
}

type RemovedPlanFromGroup struct {
	ID       string `json:"id"`
	Removed  bool   `json:"removed"`
	Object   string `json:"object"`
	Livemode bool   `json:"livemode"`
}

type ReorderedPlans struct {
	Reordered bool   `json:"reordered"`
	Object    string `json:"object"`
	Livemode  bool   `json:"livemode"`
}

type SeatBalance struct {
	Current  int    `json:"current"`
	AsOf     string `json:"as_of"`
	Object   string `json:"object"`
	Livemode bool   `json:"livemode"`
}

type SeatBalanceCollection struct {
	Balances map[string]SeatBalanceCollectionBalancesValue `json:"balances"`
	Object   string                                        `json:"object"`
	Livemode bool                                          `json:"livemode"`
}

type SeatBalanceCollectionBalancesValue struct {
	Current int    `json:"current"`
	AsOf    string `json:"as_of"`
}

type SeatEvent struct {
	ID              string `json:"id"`
	CustomerID      string `json:"customer_id"`
	FeatureCode     string `json:"feature_code"`
	PreviousBalance int    `json:"previous_balance"`
	NewBalance      int    `json:"new_balance"`
	Ts              string `json:"ts"`
	CreatedAt       string `json:"created_at"`
	Object          string `json:"object"`
	Livemode        bool   `json:"livemode"`
}

type SeatsSetAllResult struct {
	Object     string      `json:"object"`
	Data       []SeatEvent `json:"data"`
	HasMore    bool        `json:"has_more"`
	NextCursor *string     `json:"next_cursor,omitempty"`
}

type SentInvoice struct {
	Sent     bool   `json:"sent"`
	SentAt   string `json:"sent_at"`
	Object   string `json:"object"`
	Livemode bool   `json:"livemode"`
}

type SetPlanRegionalPricingParamsFeaturesItem struct {
	FeatureID        string `json:"feature_id"`
	OverageUnitPrice int    `json:"overage_unit_price"`
}

type SetPlanRegionalPricingParamsPricesItem struct {
	PriceID         string `json:"price_id"`
	Price           int    `json:"price"`
	IncludedBalance *int   `json:"included_balance,omitempty"`
}

type Subscription struct {
	ID                  string                           `json:"id"`
	CustomerID          string                           `json:"customer_id"`
	Plan                SubscriptionPlan                 `json:"plan"`
	Name                string                           `json:"name"`
	Description         *string                          `json:"description"`
	Status              SubscriptionStatus               `json:"status"`
	BillingInterval     *BillingInterval                 `json:"billing_interval"`
	TrialEndsAt         *string                          `json:"trial_ends_at"`
	CurrentPeriod       *SubscriptionCurrentPeriod       `json:"current_period"`
	Cancellation        *SubscriptionCancellation        `json:"cancellation"`
	CancelAtPeriodEnd   bool                             `json:"cancel_at_period_end"`
	ScheduledPlanChange *SubscriptionScheduledPlanChange `json:"scheduled_plan_change"`
	Discount            *SubscriptionDiscount            `json:"discount"`
	StartDate           string                           `json:"start_date"`
	EndDate             *string                          `json:"end_date"`
	BillingDayOfMonth   *int                             `json:"billing_day_of_month"`
	NextBillingDate     *string                          `json:"next_billing_date"`
	CheckoutURL         *string                          `json:"checkout_url"`
	CreatedAt           string                           `json:"created_at"`
	UpdatedAt           string                           `json:"updated_at"`
	ConsumptionModel    *ConsumptionModel                `json:"consumption_model"`
	Features            []SubscriptionFeaturesItem       `json:"features"`
	Credits             *SubscriptionCredits             `json:"credits"`
	Balance             *SubscriptionBalance             `json:"balance"`
	PriceID             *string                          `json:"price_id"`
	Object              string                           `json:"object"`
	Livemode            bool                             `json:"livemode"`
}

type SubscriptionAddon struct {
	AddonID        string `json:"addon_id"`
	Status         string `json:"status"`
	ProratedCharge int    `json:"prorated_charge"`
	Object         string `json:"object"`
	Livemode       bool   `json:"livemode"`
}

type SubscriptionBalance struct {
	Remaining float64 `json:"remaining"`
	Included  float64 `json:"included"`
	Currency  string  `json:"currency"`
}

type SubscriptionCancellation struct {
	ScheduledAt string  `json:"scheduled_at"`
	Reason      *string `json:"reason"`
	EffectiveAt string  `json:"effective_at"`
}

type SubscriptionCredits struct {
	Remaining float64 `json:"remaining"`
	Included  float64 `json:"included"`
	Purchased float64 `json:"purchased"`
}

type SubscriptionCurrentPeriod struct {
	Start         string  `json:"start"`
	End           string  `json:"end"`
	DaysRemaining float64 `json:"days_remaining"`
}

type SubscriptionDiscount struct {
	Type   string  `json:"type"`
	Value  float64 `json:"value"`
	Name   *string `json:"name"`
	EndsAt *string `json:"ends_at"`
}

type SubscriptionFeaturesItem struct {
	Value any
}

func (value *SubscriptionFeaturesItem) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	var discriminator string
	if raw, ok := fields["type"]; ok {
		if err := json.Unmarshal(raw, &discriminator); err != nil {
			return err
		}
	}
	switch discriminator {
	case "boolean":
		var decoded SubscriptionFeaturesItemVariant1
		if err := json.Unmarshal(data, &decoded); err != nil {
			return err
		}
		value.Value = decoded
		return nil
	case "usage":
		var decoded SubscriptionFeaturesItemVariant2
		if err := json.Unmarshal(data, &decoded); err != nil {
			return err
		}
		value.Value = decoded
		return nil
	case "seats":
		var decoded SubscriptionFeaturesItemVariant3
		if err := json.Unmarshal(data, &decoded); err != nil {
			return err
		}
		value.Value = decoded
		return nil
	case "quota":
		var decoded SubscriptionFeaturesItemVariant4
		if err := json.Unmarshal(data, &decoded); err != nil {
			return err
		}
		value.Value = decoded
		return nil
	}
	if hasJSONFields(fields, "code", "name", "type", "enabled") {
		var decoded SubscriptionFeaturesItemVariant1
		if err := json.Unmarshal(data, &decoded); err != nil {
			return err
		}
		value.Value = decoded
		return nil
	}
	if hasJSONFields(fields, "code", "name", "type", "usage") {
		var decoded SubscriptionFeaturesItemVariant3
		if err := json.Unmarshal(data, &decoded); err != nil {
			return err
		}
		value.Value = decoded
		return nil
	}
	if hasJSONFields(fields, "code", "name", "type") {
		var decoded SubscriptionFeaturesItemVariant2
		if err := json.Unmarshal(data, &decoded); err != nil {
			return err
		}
		value.Value = decoded
		return nil
	}
	if hasJSONFields(fields, "code", "name", "type") {
		var decoded SubscriptionFeaturesItemVariant4
		if err := json.Unmarshal(data, &decoded); err != nil {
			return err
		}
		value.Value = decoded
		return nil
	}
	var decoded SubscriptionFeaturesItemVariant1
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	value.Value = decoded
	return nil
}

func (value SubscriptionFeaturesItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(value.Value)
}

func (value SubscriptionFeaturesItem) AsSubscriptionFeaturesItemVariant1() (SubscriptionFeaturesItemVariant1, bool) {
	decoded, ok := value.Value.(SubscriptionFeaturesItemVariant1)
	return decoded, ok
}

func (value SubscriptionFeaturesItem) AsSubscriptionFeaturesItemVariant2() (SubscriptionFeaturesItemVariant2, bool) {
	decoded, ok := value.Value.(SubscriptionFeaturesItemVariant2)
	return decoded, ok
}

func (value SubscriptionFeaturesItem) AsSubscriptionFeaturesItemVariant3() (SubscriptionFeaturesItemVariant3, bool) {
	decoded, ok := value.Value.(SubscriptionFeaturesItemVariant3)
	return decoded, ok
}

func (value SubscriptionFeaturesItem) AsSubscriptionFeaturesItemVariant4() (SubscriptionFeaturesItemVariant4, bool) {
	decoded, ok := value.Value.(SubscriptionFeaturesItemVariant4)
	return decoded, ok
}

type SubscriptionFeaturesItemVariant1 struct {
	Code    string `json:"code"`
	Name    string `json:"name"`
	Type    string `json:"type"`
	Enabled bool   `json:"enabled"`
}

type SubscriptionFeaturesItemVariant2 struct {
	Code  string                                 `json:"code"`
	Name  string                                 `json:"name"`
	Type  string                                 `json:"type"`
	Usage *SubscriptionFeaturesItemVariant2Usage `json:"usage,omitempty"`
}

type SubscriptionFeaturesItemVariant2Usage struct {
	Current          float64  `json:"current"`
	Included         float64  `json:"included"`
	OverageQuantity  float64  `json:"overage_quantity"`
	OverageUnitPrice *float64 `json:"overage_unit_price,omitempty"`
}

type SubscriptionFeaturesItemVariant3 struct {
	Code  string                                `json:"code"`
	Name  string                                `json:"name"`
	Type  string                                `json:"type"`
	Usage SubscriptionFeaturesItemVariant3Usage `json:"usage"`
}

type SubscriptionFeaturesItemVariant3Usage struct {
	Current          float64  `json:"current"`
	Included         float64  `json:"included"`
	OverageQuantity  float64  `json:"overage_quantity"`
	OverageUnitPrice *float64 `json:"overage_unit_price,omitempty"`
}

type SubscriptionFeaturesItemVariant4 struct {
	Code string `json:"code"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type SubscriptionPlan struct {
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	BasePrice float64 `json:"base_price"`
}

type SubscriptionScheduledPlanChange struct {
	ChangeType         string  `json:"change_type"`
	NewPlanID          *string `json:"new_plan_id"`
	NewPlanName        *string `json:"new_plan_name"`
	NewBillingInterval *string `json:"new_billing_interval"`
	ScheduledFor       string  `json:"scheduled_for"`
}

type SubscriptionsListResult struct {
	Object     string                `json:"object"`
	Data       []SubscriptionSummary `json:"data"`
	HasMore    bool                  `json:"has_more"`
	NextCursor *string               `json:"next_cursor,omitempty"`
}

type SubscriptionSummary struct {
	ID                  string                                  `json:"id"`
	CustomerID          string                                  `json:"customer_id"`
	Plan                SubscriptionSummaryPlan                 `json:"plan"`
	Name                string                                  `json:"name"`
	Description         *string                                 `json:"description"`
	Status              SubscriptionStatus                      `json:"status"`
	BillingInterval     *BillingInterval                        `json:"billing_interval"`
	TrialEndsAt         *string                                 `json:"trial_ends_at"`
	CurrentPeriod       *SubscriptionSummaryCurrentPeriod       `json:"current_period"`
	Cancellation        *SubscriptionSummaryCancellation        `json:"cancellation"`
	CancelAtPeriodEnd   bool                                    `json:"cancel_at_period_end"`
	ScheduledPlanChange *SubscriptionSummaryScheduledPlanChange `json:"scheduled_plan_change"`
	Discount            *SubscriptionSummaryDiscount            `json:"discount"`
	StartDate           string                                  `json:"start_date"`
	EndDate             *string                                 `json:"end_date"`
	BillingDayOfMonth   *int                                    `json:"billing_day_of_month"`
	NextBillingDate     *string                                 `json:"next_billing_date"`
	CheckoutURL         *string                                 `json:"checkout_url"`
	CreatedAt           string                                  `json:"created_at"`
	UpdatedAt           string                                  `json:"updated_at"`
	PriceID             *string                                 `json:"price_id"`
	Object              string                                  `json:"object"`
	Livemode            bool                                    `json:"livemode"`
}

type SubscriptionSummaryCancellation struct {
	ScheduledAt string  `json:"scheduled_at"`
	Reason      *string `json:"reason"`
	EffectiveAt string  `json:"effective_at"`
}

type SubscriptionSummaryCurrentPeriod struct {
	Start         string  `json:"start"`
	End           string  `json:"end"`
	DaysRemaining float64 `json:"days_remaining"`
}

type SubscriptionSummaryDiscount struct {
	Type   string  `json:"type"`
	Value  float64 `json:"value"`
	Name   *string `json:"name"`
	EndsAt *string `json:"ends_at"`
}

type SubscriptionSummaryPlan struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type SubscriptionSummaryScheduledPlanChange struct {
	ChangeType         string  `json:"change_type"`
	NewPlanID          *string `json:"new_plan_id"`
	NewPlanName        *string `json:"new_plan_name"`
	NewBillingInterval *string `json:"new_billing_interval"`
	ScheduledFor       string  `json:"scheduled_for"`
}

type TestClock struct {
	SimulatedTime *string `json:"simulated_time"`
	IsActive      bool    `json:"is_active"`
	Now           string  `json:"now"`
	Object        string  `json:"object"`
	Livemode      bool    `json:"livemode"`
}

type TestClockBilling struct {
	CustomersFound int    `json:"customers_found"`
	Enqueued       int    `json:"enqueued"`
	Failed         int    `json:"failed"`
	DunningRetried int    `json:"dunning_retried"`
	DunningFailed  int    `json:"dunning_failed"`
	Object         string `json:"object"`
	Livemode       bool   `json:"livemode"`
}

type TrackUsageParamsPropertiesItem struct {
	Property string `json:"property"`
	Value    string `json:"value"`
}

type Transaction struct {
	ID                string            `json:"id"`
	InvoiceID         *string           `json:"invoice_id"`
	GrossAmount       *int              `json:"gross_amount"`
	Subtotal          *int              `json:"subtotal"`
	TaxAmount         *int              `json:"tax_amount"`
	PresentmentAmount *int              `json:"presentment_amount"`
	Currency          string            `json:"currency"`
	Provider          PaymentProvider   `json:"provider"`
	Status            TransactionStatus `json:"status"`
	CustomerEmail     *string           `json:"customer_email"`
	CustomerName      *string           `json:"customer_name"`
	PaidAt            *string           `json:"paid_at"`
	CreatedAt         string            `json:"created_at"`
	UpdatedAt         string            `json:"updated_at"`
	AvailableAt       *string           `json:"available_at"`
	Object            string            `json:"object"`
	Livemode          bool              `json:"livemode"`
}

type TransactionListItem struct {
	ID                string            `json:"id"`
	InvoiceID         *string           `json:"invoice_id"`
	GrossAmount       *int              `json:"gross_amount"`
	Subtotal          *int              `json:"subtotal"`
	TaxAmount         *int              `json:"tax_amount"`
	PresentmentAmount *int              `json:"presentment_amount"`
	Currency          string            `json:"currency"`
	Provider          PaymentProvider   `json:"provider"`
	Status            TransactionStatus `json:"status"`
	CustomerEmail     *string           `json:"customer_email"`
	CustomerName      *string           `json:"customer_name"`
	PaidAt            *string           `json:"paid_at"`
	CreatedAt         string            `json:"created_at"`
	UpdatedAt         string            `json:"updated_at"`
	Object            string            `json:"object"`
	Livemode          bool              `json:"livemode"`
}

type TransactionRetry struct {
	OriginalTransactionID string `json:"original_transaction_id"`
	InvoiceID             string `json:"invoice_id"`
	Status                string `json:"status"`
	Object                string `json:"object"`
	Livemode              bool   `json:"livemode"`
}

type TransactionsListResult struct {
	Object     string                `json:"object"`
	Data       []TransactionListItem `json:"data"`
	HasMore    bool                  `json:"has_more"`
	NextCursor *string               `json:"next_cursor,omitempty"`
}

type UpdateCustomerParamsAddress struct {
	Line1      string  `json:"line1"`
	Line2      *string `json:"line2,omitempty"`
	City       string  `json:"city"`
	State      *string `json:"state,omitempty"`
	PostalCode string  `json:"postal_code"`
	Country    string  `json:"country"`
	Region     *string `json:"region,omitempty"`
}

type UpdateOfferParamsPhasesItem struct {
	Value any
}

func (value *UpdateOfferParamsPhasesItem) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	var discriminator string
	if raw, ok := fields["type"]; ok {
		if err := json.Unmarshal(raw, &discriminator); err != nil {
			return err
		}
	}
	switch discriminator {
	case "free_trial":
		var decoded UpdateOfferParamsPhasesItemVariant1
		if err := json.Unmarshal(data, &decoded); err != nil {
			return err
		}
		value.Value = decoded
		return nil
	case "percentage":
		var decoded UpdateOfferParamsPhasesItemVariant2
		if err := json.Unmarshal(data, &decoded); err != nil {
			return err
		}
		value.Value = decoded
		return nil
	case "amount_off":
		var decoded UpdateOfferParamsPhasesItemVariant3
		if err := json.Unmarshal(data, &decoded); err != nil {
			return err
		}
		value.Value = decoded
		return nil
	case "fixed_price":
		var decoded UpdateOfferParamsPhasesItemVariant4
		if err := json.Unmarshal(data, &decoded); err != nil {
			return err
		}
		value.Value = decoded
		return nil
	}
	if hasJSONFields(fields, "type", "duration_cycles", "percentage") {
		var decoded UpdateOfferParamsPhasesItemVariant2
		if err := json.Unmarshal(data, &decoded); err != nil {
			return err
		}
		value.Value = decoded
		return nil
	}
	if hasJSONFields(fields, "type", "duration_cycles", "amounts") {
		var decoded UpdateOfferParamsPhasesItemVariant3
		if err := json.Unmarshal(data, &decoded); err != nil {
			return err
		}
		value.Value = decoded
		return nil
	}
	if hasJSONFields(fields, "type", "duration_cycles", "prices") {
		var decoded UpdateOfferParamsPhasesItemVariant4
		if err := json.Unmarshal(data, &decoded); err != nil {
			return err
		}
		value.Value = decoded
		return nil
	}
	if hasJSONFields(fields, "type", "duration_days") {
		var decoded UpdateOfferParamsPhasesItemVariant1
		if err := json.Unmarshal(data, &decoded); err != nil {
			return err
		}
		value.Value = decoded
		return nil
	}
	var decoded UpdateOfferParamsPhasesItemVariant1
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	value.Value = decoded
	return nil
}

func (value UpdateOfferParamsPhasesItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(value.Value)
}

func (value UpdateOfferParamsPhasesItem) AsUpdateOfferParamsPhasesItemVariant1() (UpdateOfferParamsPhasesItemVariant1, bool) {
	decoded, ok := value.Value.(UpdateOfferParamsPhasesItemVariant1)
	return decoded, ok
}

func (value UpdateOfferParamsPhasesItem) AsUpdateOfferParamsPhasesItemVariant2() (UpdateOfferParamsPhasesItemVariant2, bool) {
	decoded, ok := value.Value.(UpdateOfferParamsPhasesItemVariant2)
	return decoded, ok
}

func (value UpdateOfferParamsPhasesItem) AsUpdateOfferParamsPhasesItemVariant3() (UpdateOfferParamsPhasesItemVariant3, bool) {
	decoded, ok := value.Value.(UpdateOfferParamsPhasesItemVariant3)
	return decoded, ok
}

func (value UpdateOfferParamsPhasesItem) AsUpdateOfferParamsPhasesItemVariant4() (UpdateOfferParamsPhasesItemVariant4, bool) {
	decoded, ok := value.Value.(UpdateOfferParamsPhasesItemVariant4)
	return decoded, ok
}

type UpdateOfferParamsPhasesItemVariant1 struct {
	Type         string `json:"type"`
	DurationDays int    `json:"duration_days"`
}

type UpdateOfferParamsPhasesItemVariant2 struct {
	Type           string `json:"type"`
	DurationCycles int    `json:"duration_cycles"`
	Percentage     int    `json:"percentage"`
}

type UpdateOfferParamsPhasesItemVariant3 struct {
	Type           string                                           `json:"type"`
	DurationCycles int                                              `json:"duration_cycles"`
	Amounts        []UpdateOfferParamsPhasesItemVariant3AmountsItem `json:"amounts"`
}

type UpdateOfferParamsPhasesItemVariant3AmountsItem struct {
	Currency string `json:"currency"`
	Amount   int    `json:"amount"`
}

type UpdateOfferParamsPhasesItemVariant4 struct {
	Type           string                                          `json:"type"`
	DurationCycles int                                             `json:"duration_cycles"`
	Prices         []UpdateOfferParamsPhasesItemVariant4PricesItem `json:"prices"`
}

type UpdateOfferParamsPhasesItemVariant4PricesItem struct {
	Currency string `json:"currency"`
	Amount   int    `json:"amount"`
}

type UpdatePlanFeatureParamsOverage struct {
	Enabled   *bool `json:"enabled,omitempty"`
	UnitPrice *int  `json:"unit_price,omitempty"`
}

type UpdatePlanPriceParamsMarketPricesItem struct {
	MarketGroupID string `json:"market_group_id"`
	Currency      string `json:"currency"`
	Price         int    `json:"price"`
}

type UpsertRegionalPricesParamsOverridesItem struct {
	Currency        string `json:"currency"`
	Price           int    `json:"price"`
	IncludedBalance *int   `json:"included_balance,omitempty"`
}

type UsageAdjustment struct {
	ID            string  `json:"id"`
	Value         int     `json:"value"`
	PreviousValue int     `json:"previous_value"`
	Adjustment    int     `json:"adjustment"`
	CustomerID    string  `json:"customer_id"`
	Reason        *string `json:"reason"`
	Ts            string  `json:"ts"`
	CreatedAt     string  `json:"created_at"`
	FeatureCode   string  `json:"feature_code"`
	Object        string  `json:"object"`
	Livemode      bool    `json:"livemode"`
}

type UsageCheck struct {
	Value any
}

func (value *UsageCheck) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	var discriminator string
	if raw, ok := fields["consumption_model"]; ok {
		if err := json.Unmarshal(raw, &discriminator); err != nil {
			return err
		}
	}
	switch discriminator {
	case "metered":
		var decoded UsageCheckVariant1
		if err := json.Unmarshal(data, &decoded); err != nil {
			return err
		}
		value.Value = decoded
		return nil
	case "credits":
		var decoded UsageCheckVariant2
		if err := json.Unmarshal(data, &decoded); err != nil {
			return err
		}
		value.Value = decoded
		return nil
	case "balance":
		var decoded UsageCheckVariant3
		if err := json.Unmarshal(data, &decoded); err != nil {
			return err
		}
		value.Value = decoded
		return nil
	}
	if hasJSONFields(fields, "allowed", "subscription_status", "feature_code", "quantity", "consumption_model", "current", "remaining", "unlimited", "included", "overage_enabled", "overage_unit_price", "object", "livemode") {
		var decoded UsageCheckVariant1
		if err := json.Unmarshal(data, &decoded); err != nil {
			return err
		}
		value.Value = decoded
		return nil
	}
	if hasJSONFields(fields, "allowed", "subscription_status", "feature_code", "quantity", "consumption_model", "credits_per_unit", "estimated_credits", "plan_credits", "purchased_credits", "total_credits", "object", "livemode") {
		var decoded UsageCheckVariant2
		if err := json.Unmarshal(data, &decoded); err != nil {
			return err
		}
		value.Value = decoded
		return nil
	}
	if hasJSONFields(fields, "allowed", "subscription_status", "feature_code", "quantity", "consumption_model", "unit_price", "estimated_amount", "current_balance", "block_on_exhaustion", "currency", "object", "livemode") {
		var decoded UsageCheckVariant3
		if err := json.Unmarshal(data, &decoded); err != nil {
			return err
		}
		value.Value = decoded
		return nil
	}
	var decoded UsageCheckVariant1
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	value.Value = decoded
	return nil
}

func (value UsageCheck) MarshalJSON() ([]byte, error) {
	return json.Marshal(value.Value)
}

func (value UsageCheck) AsUsageCheckVariant1() (UsageCheckVariant1, bool) {
	decoded, ok := value.Value.(UsageCheckVariant1)
	return decoded, ok
}

func (value UsageCheck) AsUsageCheckVariant2() (UsageCheckVariant2, bool) {
	decoded, ok := value.Value.(UsageCheckVariant2)
	return decoded, ok
}

func (value UsageCheck) AsUsageCheckVariant3() (UsageCheckVariant3, bool) {
	decoded, ok := value.Value.(UsageCheckVariant3)
	return decoded, ok
}

type UsageCheckVariant1 struct {
	Allowed            bool     `json:"allowed"`
	SubscriptionStatus string   `json:"subscription_status"`
	FeatureCode        string   `json:"feature_code"`
	Quantity           int      `json:"quantity"`
	Reason             *string  `json:"reason,omitempty"`
	Message            *string  `json:"message,omitempty"`
	ConsumptionModel   string   `json:"consumption_model"`
	Current            float64  `json:"current"`
	Remaining          float64  `json:"remaining"`
	Unlimited          bool     `json:"unlimited"`
	Included           float64  `json:"included"`
	OverageEnabled     bool     `json:"overage_enabled"`
	OverageUnitPrice   *float64 `json:"overage_unit_price"`
	Object             string   `json:"object"`
	Livemode           bool     `json:"livemode"`
}

type UsageCheckVariant2 struct {
	Allowed            bool    `json:"allowed"`
	SubscriptionStatus string  `json:"subscription_status"`
	FeatureCode        string  `json:"feature_code"`
	Quantity           int     `json:"quantity"`
	Reason             *string `json:"reason,omitempty"`
	Message            *string `json:"message,omitempty"`
	ConsumptionModel   string  `json:"consumption_model"`
	CreditsPerUnit     int     `json:"credits_per_unit"`
	EstimatedCredits   int     `json:"estimated_credits"`
	PlanCredits        int     `json:"plan_credits"`
	PurchasedCredits   int     `json:"purchased_credits"`
	TotalCredits       int     `json:"total_credits"`
	Object             string  `json:"object"`
	Livemode           bool    `json:"livemode"`
}

type UsageCheckVariant3 struct {
	Allowed            bool    `json:"allowed"`
	SubscriptionStatus string  `json:"subscription_status"`
	FeatureCode        string  `json:"feature_code"`
	Quantity           int     `json:"quantity"`
	Reason             *string `json:"reason,omitempty"`
	Message            *string `json:"message,omitempty"`
	ConsumptionModel   string  `json:"consumption_model"`
	UnitPrice          float64 `json:"unit_price"`
	EstimatedAmount    float64 `json:"estimated_amount"`
	CurrentBalance     float64 `json:"current_balance"`
	BlockOnExhaustion  bool    `json:"block_on_exhaustion"`
	Currency           string  `json:"currency"`
	Object             string  `json:"object"`
	Livemode           bool    `json:"livemode"`
}

type UsageEvent struct {
	ID          string                     `json:"id"`
	FeatureCode string                     `json:"feature_code"`
	Value       float64                    `json:"value"`
	CustomerID  string                     `json:"customer_id"`
	EventID     *string                    `json:"event_id"`
	Ts          string                     `json:"ts"`
	CreatedAt   string                     `json:"created_at"`
	Properties  []UsageEventPropertiesItem `json:"properties"`
	Consumption *UsageEventConsumption     `json:"consumption,omitempty"`
	Object      string                     `json:"object"`
	Livemode    bool                       `json:"livemode"`
}

type UsageEventConsumption struct {
	Model     string  `json:"model"`
	Deducted  float64 `json:"deducted"`
	Remaining float64 `json:"remaining"`
	Blocked   bool    `json:"blocked"`
}

type UsageEventPropertiesItem struct {
	Property string `json:"property"`
	Value    string `json:"value"`
}

type UsageQuota struct {
	FeatureCode    string   `json:"feature_code"`
	Current        float64  `json:"current"`
	Included       float64  `json:"included"`
	Remaining      *float64 `json:"remaining"`
	BilledQuantity float64  `json:"billed_quantity"`
	Unlimited      bool     `json:"unlimited"`
	OverageEnabled bool     `json:"overage_enabled"`
	AsOf           *string  `json:"as_of"`
	Object         string   `json:"object"`
	Livemode       bool     `json:"livemode"`
}

type UsageQuotaEvent struct {
	ID              string `json:"id"`
	CustomerID      string `json:"customer_id"`
	FeatureCode     string `json:"feature_code"`
	PreviousBalance int    `json:"previous_balance"`
	NewBalance      int    `json:"new_balance"`
	Ts              string `json:"ts"`
	CreatedAt       string `json:"created_at"`
	Object          string `json:"object"`
	Livemode        bool   `json:"livemode"`
}

type Webhook struct {
	ID          string   `json:"id"`
	URL         string   `json:"url"`
	Events      []string `json:"events"`
	Description *string  `json:"description"`
	IsActive    bool     `json:"is_active"`
	APIVersion  *string  `json:"api_version"`
	CreatedAt   string   `json:"created_at"`
	Object      string   `json:"object"`
	Livemode    bool     `json:"livemode"`
}

type WebhookAddonRef struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type WebhookBalance struct {
	CurrentBalance float64 `json:"current_balance"`
}

type WebhookBankRef struct {
	BankName *string `json:"bank_name"`
	Last4    string  `json:"last4"`
}

type WebhookCardInfo struct {
	Brand    string  `json:"brand"`
	Last4    string  `json:"last4"`
	ExpMonth float64 `json:"exp_month"`
	ExpYear  float64 `json:"exp_year"`
}

type WebhookCreditsBalance struct {
	PlanCredits      float64 `json:"plan_credits"`
	PurchasedCredits float64 `json:"purchased_credits"`
	TotalCredits     float64 `json:"total_credits"`
}

type WebhookPlanRef struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type WebhookSeatSummary struct {
	Code      string   `json:"code"`
	Current   *float64 `json:"current"`
	Included  *float64 `json:"included"`
	Remaining *float64 `json:"remaining"`
	Unlimited *bool    `json:"unlimited"`
}

type WebhooksListResult struct {
	Object     string    `json:"object"`
	Data       []Webhook `json:"data"`
	HasMore    bool      `json:"has_more"`
	NextCursor *string   `json:"next_cursor,omitempty"`
}

type WebhookTest struct {
	Success     bool   `json:"success"`
	DeliveryID  string `json:"delivery_id"`
	DeliveredAt string `json:"delivered_at"`
	Object      string `json:"object"`
	Livemode    bool   `json:"livemode"`
}
