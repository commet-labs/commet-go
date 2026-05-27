package commet

import (
	"context"
	"fmt"
)

type ListPlanGroupsParams struct {
	Limit  *int   `json:"limit,omitempty"`
	Cursor string `json:"cursor,omitempty"`
}

type CreatePlanGroupParams struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	IsPublic    *bool  `json:"is_public,omitempty"`
}

type UpdatePlanGroupParams struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	IsPublic    *bool  `json:"is_public,omitempty"`
}

type AddPlanToGroupParams struct {
	PlanID    string `json:"plan_id"`
	SortOrder *int   `json:"sort_order,omitempty"`
}

type ReorderPlansParams struct {
	PlanIDs []string `json:"plan_ids"`
}

type PlanGroupsResource struct {
	http *httpClient
}

func (r *PlanGroupsResource) List(ctx context.Context, params *ListPlanGroupsParams) (*ApiResponse[[]PlanGroup], error) {
	queryParams := map[string]string{}
	if params != nil {
		if params.Limit != nil {
			queryParams["limit"] = fmt.Sprintf("%d", *params.Limit)
		}
		if params.Cursor != "" {
			queryParams["cursor"] = params.Cursor
		}
	}
	return parseResponse[[]PlanGroup](r.http.get(ctx, "/plan-groups", queryParams))
}

func (r *PlanGroupsResource) Get(ctx context.Context, planGroupID string) (*ApiResponse[PlanGroupDetail], error) {
	return parseResponse[PlanGroupDetail](r.http.get(ctx, fmt.Sprintf("/plan-groups/%s", planGroupID), nil))
}

func (r *PlanGroupsResource) Create(ctx context.Context, params *CreatePlanGroupParams) (*ApiResponse[PlanGroup], error) {
	body := buildBody(map[string]any{
		"name":        params.Name,
		"description": params.Description,
		"is_public":   params.IsPublic,
	})
	return parseResponse[PlanGroup](r.http.post(ctx, "/plan-groups", body, ""))
}

func (r *PlanGroupsResource) Update(ctx context.Context, planGroupID string, params *UpdatePlanGroupParams) (*ApiResponse[PlanGroup], error) {
	body := buildBody(map[string]any{
		"name":        params.Name,
		"description": params.Description,
		"is_public":   params.IsPublic,
	})
	return parseResponse[PlanGroup](r.http.put(ctx, fmt.Sprintf("/plan-groups/%s", planGroupID), body, ""))
}

func (r *PlanGroupsResource) Delete(ctx context.Context, planGroupID string) (*ApiResponse[struct{}], error) {
	return parseResponse[struct{}](r.http.delete(ctx, fmt.Sprintf("/plan-groups/%s", planGroupID), nil, ""))
}

func (r *PlanGroupsResource) AddPlan(ctx context.Context, planGroupID string, params *AddPlanToGroupParams) (*ApiResponse[PlanGroupDetail], error) {
	body := buildBody(map[string]any{
		"plan_id":    params.PlanID,
		"sort_order": params.SortOrder,
	})
	return parseResponse[PlanGroupDetail](r.http.post(ctx, fmt.Sprintf("/plan-groups/%s/plans", planGroupID), body, ""))
}

func (r *PlanGroupsResource) RemovePlan(ctx context.Context, planGroupID string, planID string) (*ApiResponse[struct{}], error) {
	return parseResponse[struct{}](r.http.delete(ctx, fmt.Sprintf("/plan-groups/%s/plans/%s", planGroupID, planID), nil, ""))
}

func (r *PlanGroupsResource) ReorderPlans(ctx context.Context, planGroupID string, params *ReorderPlansParams) (*ApiResponse[PlanGroupDetail], error) {
	body := buildBody(map[string]any{
		"plan_ids": params.PlanIDs,
	})
	return parseResponse[PlanGroupDetail](r.http.put(ctx, fmt.Sprintf("/plan-groups/%s/plans/reorder", planGroupID), body, ""))
}
