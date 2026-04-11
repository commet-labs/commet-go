package commet

import "context"

type CustomerContext struct {
	Features     CustomerContextFeatures
	Seats        CustomerContextSeats
	Usage        CustomerContextUsage
	Subscription CustomerContextSubscription
	Portal       CustomerContextPortal
}

type CustomerFeatures struct {
	customerID string
	resource   Features
}

func (f *CustomerFeatures) Get(ctx context.Context, code string) (*ApiResponse[FeatureAccess], error) {
	return f.resource.Get(ctx, code, f.customerID)
}

func (f *CustomerFeatures) Check(ctx context.Context, code string) (*ApiResponse[CheckResult], error) {
	return f.resource.Check(ctx, code, f.customerID)
}

func (f *CustomerFeatures) CanUse(ctx context.Context, code string) (*ApiResponse[CanUseResult], error) {
	return f.resource.CanUse(ctx, code, f.customerID)
}

func (f *CustomerFeatures) List(ctx context.Context) (*ApiResponse[[]FeatureAccess], error) {
	return f.resource.List(ctx, f.customerID)
}

type CustomerSeats struct {
	customerID string
	resource   Seats
}

func (s *CustomerSeats) Add(ctx context.Context, seatType string, count int) (*ApiResponse[SeatEvent], error) {
	return s.resource.Add(ctx, &SeatParams{
		SeatType:   seatType,
		Count:      count,
		CustomerID: s.customerID,
	})
}

func (s *CustomerSeats) Remove(ctx context.Context, seatType string, count int) (*ApiResponse[SeatEvent], error) {
	return s.resource.Remove(ctx, &SeatParams{
		SeatType:   seatType,
		Count:      count,
		CustomerID: s.customerID,
	})
}

func (s *CustomerSeats) Set(ctx context.Context, seatType string, count int) (*ApiResponse[SeatEvent], error) {
	return s.resource.Set(ctx, &SeatParams{
		SeatType:   seatType,
		Count:      count,
		CustomerID: s.customerID,
	})
}

func (s *CustomerSeats) GetBalance(ctx context.Context, seatType string) (*ApiResponse[SeatBalance], error) {
	return s.resource.GetBalance(ctx, &GetSeatBalanceParams{
		SeatType:   seatType,
		CustomerID: s.customerID,
	})
}

type CustomerUsage struct {
	customerID string
	resource   Usage
}

func (u *CustomerUsage) Track(ctx context.Context, feature string, opts ...TrackOption) (*ApiResponse[UsageEvent], error) {
	params := &TrackUsageParams{
		Feature:    feature,
		CustomerID: u.customerID,
	}
	for _, opt := range opts {
		opt(params)
	}
	return u.resource.Track(ctx, params)
}

type TrackOption func(*TrackUsageParams)

func WithValue(value int) TrackOption {
	return func(p *TrackUsageParams) {
		p.Value = &value
	}
}

func WithProperties(properties map[string]string) TrackOption {
	return func(p *TrackUsageParams) {
		p.Properties = properties
	}
}

type CustomerSubscription struct {
	customerID string
	resource   Subscriptions
}

func (sub *CustomerSubscription) Get(ctx context.Context) (*ApiResponse[ActiveSubscription], error) {
	return sub.resource.Get(ctx, sub.customerID)
}

type CustomerPortal struct {
	customerID string
	resource   Portal
}

func (p *CustomerPortal) GetURL(ctx context.Context) (*ApiResponse[PortalSession], error) {
	return p.resource.GetURL(ctx, &GetPortalURLParams{
		CustomerID: p.customerID,
	})
}
