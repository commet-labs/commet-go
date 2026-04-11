package commet

import "context"

type Customers interface {
	Create(ctx context.Context, params *CreateCustomerParams) (*ApiResponse[Customer], error)
	CreateBatch(ctx context.Context, customers []CreateCustomerParams, idempotencyKey string) (*ApiResponse[BatchResult], error)
	Get(ctx context.Context, customerID string) (*ApiResponse[Customer], error)
	Update(ctx context.Context, customerID string, params *UpdateCustomerParams) (*ApiResponse[Customer], error)
	List(ctx context.Context, params *ListCustomersParams) (*ApiResponse[[]Customer], error)
	Archive(ctx context.Context, customerID string, idempotencyKey string) (*ApiResponse[Customer], error)
}

type Subscriptions interface {
	Create(ctx context.Context, params *CreateSubscriptionParams) (*ApiResponse[Subscription], error)
	Get(ctx context.Context, customerID string) (*ApiResponse[ActiveSubscription], error)
	Cancel(ctx context.Context, subscriptionID string, params *CancelSubscriptionParams) (*ApiResponse[Subscription], error)
}

type Usage interface {
	Track(ctx context.Context, params *TrackUsageParams) (*ApiResponse[UsageEvent], error)
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
	Check(ctx context.Context, code string, customerID string) (*ApiResponse[CheckResult], error)
	CanUse(ctx context.Context, code string, customerID string) (*ApiResponse[CanUseResult], error)
	List(ctx context.Context, customerID string) (*ApiResponse[[]FeatureAccess], error)
}

type Plans interface {
	List(ctx context.Context, params *ListPlansParams) (*ApiResponse[[]Plan], error)
	Get(ctx context.Context, planCode string) (*ApiResponse[PlanDetail], error)
}

type Portal interface {
	GetURL(ctx context.Context, params *GetPortalURLParams) (*ApiResponse[PortalSession], error)
}

type CreditPacks interface {
	List(ctx context.Context) (*ApiResponse[[]CreditPack], error)
}

type WebhookVerifier interface {
	Verify(payload string, signature string, secret string) bool
	VerifyAndParse(rawBody string, signature string, secret string) (map[string]any, error)
}

type CustomerContextFeatures interface {
	Get(ctx context.Context, code string) (*ApiResponse[FeatureAccess], error)
	Check(ctx context.Context, code string) (*ApiResponse[CheckResult], error)
	CanUse(ctx context.Context, code string) (*ApiResponse[CanUseResult], error)
	List(ctx context.Context) (*ApiResponse[[]FeatureAccess], error)
}

type CustomerContextSeats interface {
	Add(ctx context.Context, seatType string, count int) (*ApiResponse[SeatEvent], error)
	Remove(ctx context.Context, seatType string, count int) (*ApiResponse[SeatEvent], error)
	Set(ctx context.Context, seatType string, count int) (*ApiResponse[SeatEvent], error)
	GetBalance(ctx context.Context, seatType string) (*ApiResponse[SeatBalance], error)
}

type CustomerContextUsage interface {
	Track(ctx context.Context, feature string, opts ...TrackOption) (*ApiResponse[UsageEvent], error)
}

type CustomerContextSubscription interface {
	Get(ctx context.Context) (*ApiResponse[ActiveSubscription], error)
}

type CustomerContextPortal interface {
	GetURL(ctx context.Context) (*ApiResponse[PortalSession], error)
}
