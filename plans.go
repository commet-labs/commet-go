package commet

import (
	"context"
	"fmt"
)

// ListPlansParams holds parameters for listing plans.
type ListPlansParams struct {
	IncludePrivate *bool  `json:"include_private,omitempty"`
	Limit          *int   `json:"limit,omitempty"`
	Cursor         string `json:"cursor,omitempty"`
}

// PlansResource provides access to plan operations.
type PlansResource struct {
	http *httpClient
}

// List retrieves a paginated list of plans.
func (r *PlansResource) List(ctx context.Context, params *ListPlansParams) (*ApiResponse, error) {
	queryParams := map[string]string{}
	if params != nil {
		if params.IncludePrivate != nil {
			queryParams["include_private"] = fmt.Sprintf("%t", *params.IncludePrivate)
		}
		if params.Limit != nil {
			queryParams["limit"] = fmt.Sprintf("%d", *params.Limit)
		}
		if params.Cursor != "" {
			queryParams["cursor"] = params.Cursor
		}
	}
	return r.http.get(ctx, "/plans", queryParams)
}

// Get retrieves a plan by its code.
func (r *PlansResource) Get(ctx context.Context, planCode string) (*ApiResponse, error) {
	return r.http.get(ctx, fmt.Sprintf("/plans/%s", planCode), nil)
}
