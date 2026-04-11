package commet

import (
	"context"
	"fmt"
)

type ListPlansParams struct {
	IncludePrivate *bool  `json:"include_private,omitempty"`
	Limit          *int   `json:"limit,omitempty"`
	Cursor         string `json:"cursor,omitempty"`
}

type PlansResource struct {
	http *httpClient
}

func (r *PlansResource) List(ctx context.Context, params *ListPlansParams) (*ApiResponse[[]Plan], error) {
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
	return parseResponse[[]Plan](r.http.get(ctx, "/plans", queryParams))
}

func (r *PlansResource) Get(ctx context.Context, planCode string) (*ApiResponse[PlanDetail], error) {
	return parseResponse[PlanDetail](r.http.get(ctx, fmt.Sprintf("/plans/%s", planCode), nil))
}
