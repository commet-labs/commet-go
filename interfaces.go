package commet

import "context"

type Customers interface {
	Create(ctx context.Context, params *CreateCustomerParams) (*ApiResponse[Customer], error)
	CreateBatch(ctx context.Context, customers []CreateCustomerParams, idempotencyKey string) (*ApiResponse[BatchResult], error)
	Get(ctx context.Context, customerID string) (*ApiResponse[Customer], error)
	Update(ctx context.Context, customerID string, params *UpdateCustomerParams) (*ApiResponse[Customer], error)
	List(ctx context.Context, params *ListCustomersParams) (*ApiResponse[[]Customer], error)
}

type Subscriptions interface {
	Create(ctx context.Context, params *CreateSubscriptionParams) (*ApiResponse[Subscription], error)
	GetActive(ctx context.Context, customerID string) (*ApiResponse[ActiveSubscription], error)
	Cancel(ctx context.Context, subscriptionID string, params *CancelSubscriptionParams) (*ApiResponse[Subscription], error)
	Uncancel(ctx context.Context, subscriptionID string) (*ApiResponse[Subscription], error)
	ChangePlan(ctx context.Context, params *ChangePlanParams) (*ApiResponse[ChangePlanResult], error)
	List(ctx context.Context, params *ListSubscriptionsParams) (*ApiResponse[[]SubscriptionListItem], error)
	PreviewChange(ctx context.Context, subscriptionID string, params *PreviewChangeParams) (*ApiResponse[PreviewChangeResult], error)
	ActivateAddon(ctx context.Context, subscriptionID string, params *ActivateAddonParams) (*ApiResponse[ActivateAddonResult], error)
	DeactivateAddon(ctx context.Context, subscriptionID string, addonID string) (*ApiResponse[DeactivateAddonResult], error)
	AdjustBalance(ctx context.Context, subscriptionID string, params *AdjustBalanceParams) (*ApiResponse[AdjustBalanceResult], error)
	TopupBalance(ctx context.Context, subscriptionID string, params *TopupBalanceParams) (*ApiResponse[TopupBalanceResult], error)
	PurchaseCredits(ctx context.Context, subscriptionID string, params *PurchaseCreditsParams) (*ApiResponse[PurchaseCreditsResult], error)
}

type Usage interface {
	Track(ctx context.Context, params *TrackUsageParams) (*ApiResponse[UsageEvent], error)
	TrackModelTokens(ctx context.Context, params *TrackModelTokensParams) (*ApiResponse[UsageEvent], error)
	Check(ctx context.Context, params *CheckUsageParams) (*ApiResponse[UsageCheckResult], error)
}

type Seats interface {
	Add(ctx context.Context, params *SeatParams) (*ApiResponse[SeatEvent], error)
	Remove(ctx context.Context, params *SeatParams) (*ApiResponse[SeatEvent], error)
	Set(ctx context.Context, params *SeatParams) (*ApiResponse[SeatEvent], error)
	SetAll(ctx context.Context, params *SetAllSeatsParams) (*ApiResponse[[]SeatEvent], error)
	GetBalance(ctx context.Context, params *GetSeatBalanceParams) (*ApiResponse[SeatBalance], error)
	GetAllBalances(ctx context.Context, params *GetAllSeatBalancesParams) (*ApiResponse[map[string]SeatBalance], error)
}

type Features interface {
	Get(ctx context.Context, code string, customerID string) (*ApiResponse[FeatureAccess], error)
	CanUse(ctx context.Context, code string, customerID string) (*ApiResponse[CanUseResult], error)
	List(ctx context.Context, customerID string) (*ApiResponse[[]FeatureAccess], error)
	Create(ctx context.Context, params *CreateFeatureParams) (*ApiResponse[Feature], error)
	Update(ctx context.Context, code string, params *UpdateFeatureParams) (*ApiResponse[Feature], error)
	Delete(ctx context.Context, code string) (*ApiResponse[DeleteResult], error)
}

type Plans interface {
	List(ctx context.Context, params *ListPlansParams) (*ApiResponse[[]Plan], error)
	Get(ctx context.Context, planID string) (*ApiResponse[PlanDetail], error)
	Create(ctx context.Context, params *CreatePlanParams) (*ApiResponse[PlanManage], error)
	Update(ctx context.Context, planID string, params *UpdatePlanParams) (*ApiResponse[PlanManage], error)
	Delete(ctx context.Context, planID string) (*ApiResponse[DeleteResult], error)
	SetVisibility(ctx context.Context, planID string, params *SetVisibilityParams) (*ApiResponse[PlanManage], error)
	AddFeature(ctx context.Context, planID string, params *AddPlanFeatureParams) (*ApiResponse[PlanFeatureManage], error)
	UpdateFeature(ctx context.Context, planID string, featureID string, params *UpdatePlanFeatureParams) (*ApiResponse[PlanFeatureManage], error)
	RemoveFeature(ctx context.Context, planID string, featureID string) (*ApiResponse[RemoveResult], error)
	AddPrice(ctx context.Context, planID string, params *AddPlanPriceParams) (*ApiResponse[PlanPriceManage], error)
	UpdatePrice(ctx context.Context, planID string, priceID string, params *UpdatePlanPriceParams) (*ApiResponse[PlanPriceManage], error)
	DeletePrice(ctx context.Context, planID string, priceID string) (*ApiResponse[DeleteResult], error)
	SetDefaultPrice(ctx context.Context, planID string, priceID string) (*ApiResponse[PlanPriceManage], error)
	SetRegionalPrices(ctx context.Context, planID string, priceID string, params *SetRegionalPricesParams) (*ApiResponse[RegionalPriceResult], error)
	DeleteRegionalPrices(ctx context.Context, planID string, priceID string) (*ApiResponse[DeleteResult], error)
}

type Portal interface {
	GetURL(ctx context.Context, params *GetPortalURLParams) (*ApiResponse[PortalSession], error)
}

type CreditPacks interface {
	List(ctx context.Context) (*ApiResponse[[]CreditPack], error)
	Create(ctx context.Context, params *CreateCreditPackParams) (*ApiResponse[CreditPackDetail], error)
	Update(ctx context.Context, creditPackID string, params *UpdateCreditPackParams) (*ApiResponse[CreditPackDetail], error)
	Delete(ctx context.Context, creditPackID string) (*ApiResponse[DeleteResult], error)
}

type Addons interface {
	ListActive(ctx context.Context, customerID string) (*ApiResponse[[]ActiveAddon], error)
	List(ctx context.Context, params *ListAddonsParams) (*ApiResponse[[]Addon], error)
	Get(ctx context.Context, addonID string) (*ApiResponse[Addon], error)
	Create(ctx context.Context, params *CreateAddonParams) (*ApiResponse[Addon], error)
	Update(ctx context.Context, addonID string, params *UpdateAddonParams) (*ApiResponse[Addon], error)
	Delete(ctx context.Context, addonID string) (*ApiResponse[DeleteResult], error)
}

type WebhookService interface {
	Verify(payload string, signature string, secret string) bool
	VerifyAndParse(rawBody string, signature string, secret string) (map[string]any, error)
	List(ctx context.Context, params *ListWebhooksParams) (*ApiResponse[[]WebhookEndpoint], error)
	Create(ctx context.Context, params *CreateWebhookParams) (*ApiResponse[WebhookEndpointCreated], error)
	Delete(ctx context.Context, webhookID string) (*ApiResponse[DeleteResult], error)
	Test(ctx context.Context, webhookID string) (*ApiResponse[WebhookTestResult], error)
}

type ApiKeys interface {
	List(ctx context.Context, params *ListApiKeysParams) (*ApiResponse[[]ApiKey], error)
	Create(ctx context.Context, params *CreateApiKeyParams) (*ApiResponse[ApiKeyCreated], error)
	Delete(ctx context.Context, apiKeyID string) (*ApiResponse[struct{}], error)
}

type Invoices interface {
	List(ctx context.Context, params *ListInvoicesParams) (*ApiResponse[[]InvoiceListItem], error)
	Get(ctx context.Context, invoiceID string) (*ApiResponse[InvoiceDetail], error)
	CreateAdjustment(ctx context.Context, params *CreateAdjustmentParams) (*ApiResponse[CreateAdjustmentResult], error)
	GetDownloadUrl(ctx context.Context, invoiceID string) (*ApiResponse[InvoiceDownloadResult], error)
	Send(ctx context.Context, invoiceID string) (*ApiResponse[InvoiceSendResult], error)
	UpdateStatus(ctx context.Context, invoiceID string, params *UpdateInvoiceStatusParams) (*ApiResponse[InvoiceStatusResult], error)
}

type Transactions interface {
	List(ctx context.Context, params *ListTransactionsParams) (*ApiResponse[[]TransactionListItem], error)
	Get(ctx context.Context, transactionID string) (*ApiResponse[TransactionDetail], error)
	Refund(ctx context.Context, transactionID string) (*ApiResponse[TransactionRefundResult], error)
	Retry(ctx context.Context, transactionID string) (*ApiResponse[TransactionRetryResult], error)
}

type PromoCodes interface {
	List(ctx context.Context, params *ListPromoCodesParams) (*ApiResponse[[]PromoCode], error)
	Get(ctx context.Context, promoCodeID string) (*ApiResponse[PromoCodeDetail], error)
	Create(ctx context.Context, params *CreatePromoCodeParams) (*ApiResponse[PromoCode], error)
	Update(ctx context.Context, promoCodeID string, params *UpdatePromoCodeParams) (*ApiResponse[PromoCodeDetail], error)
}

type PlanGroups interface {
	List(ctx context.Context, params *ListPlanGroupsParams) (*ApiResponse[[]PlanGroup], error)
	Get(ctx context.Context, planGroupID string) (*ApiResponse[PlanGroupDetail], error)
	Create(ctx context.Context, params *CreatePlanGroupParams) (*ApiResponse[PlanGroup], error)
	Update(ctx context.Context, planGroupID string, params *UpdatePlanGroupParams) (*ApiResponse[PlanGroup], error)
	Delete(ctx context.Context, planGroupID string) (*ApiResponse[struct{}], error)
	AddPlan(ctx context.Context, planGroupID string, params *AddPlanToGroupParams) (*ApiResponse[PlanGroupDetail], error)
	RemovePlan(ctx context.Context, planGroupID string, planID string) (*ApiResponse[struct{}], error)
	ReorderPlans(ctx context.Context, planGroupID string, params *ReorderPlansParams) (*ApiResponse[PlanGroupDetail], error)
}
