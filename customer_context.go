package commet

import "context"

// CustomerContext provides customer-scoped API access.
// All operations are automatically scoped to the customer ID.
type CustomerContext struct {
	Features     *CustomerFeatures
	Seats        *CustomerSeats
	Usage        *CustomerUsage
	Subscription *CustomerSubscription
	Portal       *CustomerPortal
}

// CustomerFeatures wraps FeaturesResource scoped to a customer.
type CustomerFeatures struct {
	customerID string
	resource   *FeaturesResource
}

// Get retrieves a feature by code.
func (f *CustomerFeatures) Get(ctx context.Context, code string) (*ApiResponse, error) {
	return f.resource.Get(ctx, code, f.customerID)
}

// Check checks if the customer has access to a feature.
func (f *CustomerFeatures) Check(ctx context.Context, code string) (*ApiResponse, error) {
	return f.resource.Check(ctx, code, f.customerID)
}

// CanUse checks if the customer can use a feature.
func (f *CustomerFeatures) CanUse(ctx context.Context, code string) (*ApiResponse, error) {
	return f.resource.CanUse(ctx, code, f.customerID)
}

// List retrieves all features for the customer.
func (f *CustomerFeatures) List(ctx context.Context) (*ApiResponse, error) {
	return f.resource.List(ctx, f.customerID)
}

// CustomerSeats wraps SeatsResource scoped to a customer.
type CustomerSeats struct {
	customerID string
	resource   *SeatsResource
}

// Add adds seats of a given type.
func (s *CustomerSeats) Add(ctx context.Context, seatType string, count int) (*ApiResponse, error) {
	return s.resource.Add(ctx, &SeatParams{
		SeatType:   seatType,
		Count:      count,
		CustomerID: s.customerID,
	})
}

// Remove removes seats of a given type.
func (s *CustomerSeats) Remove(ctx context.Context, seatType string, count int) (*ApiResponse, error) {
	return s.resource.Remove(ctx, &SeatParams{
		SeatType:   seatType,
		Count:      count,
		CustomerID: s.customerID,
	})
}

// Set sets the seat count for a given type.
func (s *CustomerSeats) Set(ctx context.Context, seatType string, count int) (*ApiResponse, error) {
	return s.resource.Set(ctx, &SeatParams{
		SeatType:   seatType,
		Count:      count,
		CustomerID: s.customerID,
	})
}

// GetBalance retrieves the balance for a specific seat type.
func (s *CustomerSeats) GetBalance(ctx context.Context, seatType string) (*ApiResponse, error) {
	return s.resource.GetBalance(ctx, &GetSeatBalanceParams{
		SeatType:   seatType,
		CustomerID: s.customerID,
	})
}

// CustomerUsage wraps UsageResource scoped to a customer.
type CustomerUsage struct {
	customerID string
	resource   *UsageResource
}

// Track records a usage event for the customer.
func (u *CustomerUsage) Track(ctx context.Context, feature string, opts ...TrackOption) (*ApiResponse, error) {
	params := &TrackUsageParams{
		Feature:    feature,
		CustomerID: u.customerID,
	}
	for _, opt := range opts {
		opt(params)
	}
	return u.resource.Track(ctx, params)
}

// TrackOption configures optional parameters for customer-scoped usage tracking.
type TrackOption func(*TrackUsageParams)

// WithValue sets the usage value.
func WithValue(value int) TrackOption {
	return func(p *TrackUsageParams) {
		p.Value = &value
	}
}

// WithProperties sets the usage properties.
func WithProperties(properties map[string]string) TrackOption {
	return func(p *TrackUsageParams) {
		p.Properties = properties
	}
}

// CustomerSubscription wraps SubscriptionsResource scoped to a customer.
type CustomerSubscription struct {
	customerID string
	resource   *SubscriptionsResource
}

// Get retrieves the active subscription for the customer.
func (sub *CustomerSubscription) Get(ctx context.Context) (*ApiResponse, error) {
	return sub.resource.Get(ctx, sub.customerID)
}

// CustomerPortal wraps PortalResource scoped to a customer.
type CustomerPortal struct {
	customerID string
	resource   *PortalResource
}

// GetURL requests a portal access URL for the customer.
func (p *CustomerPortal) GetURL(ctx context.Context) (*ApiResponse, error) {
	return p.resource.GetURL(ctx, &GetPortalURLParams{
		CustomerID: p.customerID,
	})
}
