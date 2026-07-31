package commet

type generatedResources struct {
	Addons        *AddonsResource
	ApiKeys       *ApiKeysResource
	CreditPacks   *CreditPacksResource
	Customers     *CustomersResource
	FeatureAccess *FeatureAccessResource
	Features      *FeaturesResource
	Invoices      *InvoicesResource
	Markets       *MarketsResource
	Offers        *OffersResource
	Payments      *PaymentsResource
	Payouts       *PayoutsResource
	PlanGroups    *PlanGroupsResource
	Plans         *PlansResource
	Portal        *PortalResource
	PromoCodes    *PromoCodesResource
	Provisioning  *ProvisioningResource
	Quota         *QuotaResource
	Seats         *SeatsResource
	Subscriptions *SubscriptionsResource
	TestClock     *TestClockResource
	Transactions  *TransactionsResource
	Usage         *UsageResource
}

func (r *generatedResources) wireResources(h *httpClient) {
	r.Addons = &AddonsResource{http: h}
	r.ApiKeys = &ApiKeysResource{http: h}
	r.CreditPacks = &CreditPacksResource{http: h}
	r.Customers = &CustomersResource{http: h}
	r.FeatureAccess = &FeatureAccessResource{http: h}
	r.Features = &FeaturesResource{http: h}
	r.Invoices = &InvoicesResource{http: h}
	r.Markets = &MarketsResource{http: h}
	r.Offers = &OffersResource{http: h}
	r.Payments = &PaymentsResource{http: h}
	r.Payouts = &PayoutsResource{http: h}
	r.PlanGroups = &PlanGroupsResource{http: h}
	r.Plans = &PlansResource{http: h}
	r.Portal = &PortalResource{http: h}
	r.PromoCodes = &PromoCodesResource{http: h}
	r.Provisioning = &ProvisioningResource{http: h}
	r.Quota = &QuotaResource{http: h}
	r.Seats = &SeatsResource{http: h}
	r.Subscriptions = &SubscriptionsResource{http: h}
	r.TestClock = &TestClockResource{http: h}
	r.Transactions = &TransactionsResource{http: h}
	r.Usage = &UsageResource{http: h}
}
