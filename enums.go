package commet

type SubscriptionStatus string

const (
	SubscriptionStatusDraft          SubscriptionStatus = "draft"
	SubscriptionStatusPendingPayment SubscriptionStatus = "pending_payment"
	SubscriptionStatusTrialing       SubscriptionStatus = "trialing"
	SubscriptionStatusActive         SubscriptionStatus = "active"
	SubscriptionStatusPastDue        SubscriptionStatus = "past_due"
	SubscriptionStatusCanceled       SubscriptionStatus = "canceled"
)

type BillingInterval string

const (
	BillingIntervalWeekly    BillingInterval = "weekly"
	BillingIntervalMonthly   BillingInterval = "monthly"
	BillingIntervalQuarterly BillingInterval = "quarterly"
	BillingIntervalYearly    BillingInterval = "yearly"
	BillingIntervalOneTime   BillingInterval = "one_time"
)

type ConsumptionModel string

const (
	ConsumptionModelMetered ConsumptionModel = "metered"
	ConsumptionModelCredits ConsumptionModel = "credits"
	ConsumptionModelBalance ConsumptionModel = "balance"
)

type InvoiceType string

const (
	InvoiceTypeRecurring       InvoiceType = "recurring"
	InvoiceTypeOverage         InvoiceType = "overage"
	InvoiceTypePlanChange      InvoiceType = "plan_change"
	InvoiceTypeAdjustment      InvoiceType = "adjustment"
	InvoiceTypeCreditPurchase  InvoiceType = "credit_purchase"
	InvoiceTypeBalanceTopup    InvoiceType = "balance_topup"
	InvoiceTypeAddonActivation InvoiceType = "addon_activation"
	InvoiceTypeOneTimePayment  InvoiceType = "one_time_payment"
	InvoiceTypeReactivation    InvoiceType = "reactivation"
)

type TransactionStatus string

const (
	TransactionStatusPending   TransactionStatus = "pending"
	TransactionStatusSucceeded TransactionStatus = "succeeded"
	TransactionStatusFailed    TransactionStatus = "failed"
	TransactionStatusRefunded  TransactionStatus = "refunded"
	TransactionStatusDisputed  TransactionStatus = "disputed"
)

type FeatureType string

const (
	FeatureTypeBoolean FeatureType = "boolean"
	FeatureTypeUsage   FeatureType = "usage"
	FeatureTypeSeats   FeatureType = "seats"
	FeatureTypeQuota   FeatureType = "quota"
)

type DiscountType string

const (
	DiscountTypePercentage DiscountType = "percentage"
	DiscountTypeAmount     DiscountType = "amount"
)

type Timezone string

const (
	TimezoneUTC                Timezone = "UTC"
	TimezoneAmericaNewYork     Timezone = "America/New_York"
	TimezoneAmericaChicago     Timezone = "America/Chicago"
	TimezoneAmericaDenver      Timezone = "America/Denver"
	TimezoneAmericaLosAngeles  Timezone = "America/Los_Angeles"
	TimezoneAmericaSaoPaulo    Timezone = "America/Sao_Paulo"
	TimezoneAmericaMexicoCity  Timezone = "America/Mexico_City"
	TimezoneAmericaBuenosAires Timezone = "America/Buenos_Aires"
	TimezoneAmericaSantiago    Timezone = "America/Santiago"
	TimezoneAmericaBogota      Timezone = "America/Bogota"
	TimezoneAmericaLima        Timezone = "America/Lima"
	TimezoneAmericaAsuncion    Timezone = "America/Asuncion"
	TimezoneEuropeLondon       Timezone = "Europe/London"
	TimezoneEuropeParis        Timezone = "Europe/Paris"
	TimezoneEuropeBerlin       Timezone = "Europe/Berlin"
	TimezoneEuropeMadrid       Timezone = "Europe/Madrid"
	TimezoneAsiaTokyo          Timezone = "Asia/Tokyo"
	TimezoneAsiaShanghai       Timezone = "Asia/Shanghai"
	TimezoneAsiaSingapore      Timezone = "Asia/Singapore"
	TimezoneAsiaDubai          Timezone = "Asia/Dubai"
	TimezoneAustraliaSydney    Timezone = "Australia/Sydney"
)
