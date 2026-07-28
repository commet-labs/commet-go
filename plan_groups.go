package commet

import (
	"context"
	"fmt"
)

type ReorderPlansInGroupParams struct {
	PlanIds        []string `json:"plan_ids"`
	IdempotencyKey string   `json:"-"`
}

type AddPlanToGroupParams struct {
	PlanID         string `json:"plan_id"`
	SortOrder      *int   `json:"sort_order,omitempty"`
	IdempotencyKey string `json:"-"`
}

type UpdatePlanGroupParams struct {
	Name           *string `json:"name,omitempty"`
	Description    *string `json:"description,omitempty"`
	IsPublic       *bool   `json:"is_public,omitempty"`
	IdempotencyKey string  `json:"-"`
}

type ListPlanGroupsParams struct {
	Cursor *string `json:"cursor,omitempty"`
	Limit  *int    `json:"limit,omitempty"`
}

type CreatePlanGroupParams struct {
	Name           string  `json:"name"`
	Description    *string `json:"description,omitempty"`
	IsPublic       *bool   `json:"is_public,omitempty"`
	IdempotencyKey string  `json:"-"`
}

type PlanGroupsResource struct {
	http *httpClient
}

// Remove a plan from a plan group.
func (r *PlanGroupsResource) RemovePlan(ctx context.Context, id string, planID string) (*RemovedPlanFromGroup, error) {
	return parseDirectResponse[RemovedPlanFromGroup](r.http.delete(ctx, fmt.Sprintf("/plan-groups/%s/plans/%s", id, planID), nil, ""))
}

// Set the display order of plans within a group. All plan IDs in the group must be provided.
func (r *PlanGroupsResource) ReorderPlans(ctx context.Context, id string, params *ReorderPlansInGroupParams) (*ReorderedPlans, error) {
	body := buildBody(map[string]any{
		"plan_ids": params.PlanIds,
	})
	return parseDirectResponse[ReorderedPlans](r.http.put(ctx, fmt.Sprintf("/plan-groups/%s/plans/reorder", id), body, params.IdempotencyKey))
}

// Add an existing plan to a plan group with optional sort order.
func (r *PlanGroupsResource) AddPlan(ctx context.Context, id string, params *AddPlanToGroupParams) (*AddedPlanToGroup, error) {
	body := buildBody(map[string]any{
		"plan_id":    params.PlanID,
		"sort_order": params.SortOrder,
	})
	return parseDirectResponse[AddedPlanToGroup](r.http.post(ctx, fmt.Sprintf("/plan-groups/%s/plans", id), body, params.IdempotencyKey))
}

// Retrieve a plan group by ID, including its plans ordered by sortOrder.
func (r *PlanGroupsResource) Get(ctx context.Context, id string) (*PlanGroupDetail, error) {
	return parseDirectResponse[PlanGroupDetail](r.http.get(ctx, fmt.Sprintf("/plan-groups/%s", id), nil))
}

// Update a plan group's name, description, or visibility.
func (r *PlanGroupsResource) Update(ctx context.Context, id string, params *UpdatePlanGroupParams) (*PlanGroup, error) {
	body := buildBody(map[string]any{
		"name":        params.Name,
		"description": params.Description,
		"is_public":   params.IsPublic,
	})
	return parseDirectResponse[PlanGroup](r.http.patch(ctx, fmt.Sprintf("/plan-groups/%s", id), body, params.IdempotencyKey))
}

// Delete a plan group. Plans in the group are unlinked, not deleted.
func (r *PlanGroupsResource) Delete(ctx context.Context, id string) (*DeletedObject, error) {
	return parseDirectResponse[DeletedObject](r.http.delete(ctx, fmt.Sprintf("/plan-groups/%s", id), nil, ""))
}

// List plan groups with cursor-based pagination.
func (r *PlanGroupsResource) List(ctx context.Context, params *ListPlanGroupsParams) (*PlanGroupsListResult, error) {
	query := map[string]string{}
	if params.Cursor != nil {
		query["cursor"] = *params.Cursor
	}
	if params.Limit != nil {
		query["limit"] = fmt.Sprintf("%d", *params.Limit)
	}
	return parseDirectResponse[PlanGroupsListResult](r.http.get(ctx, "/plan-groups", query))
}

// Create a new plan group for organizing plans.
func (r *PlanGroupsResource) Create(ctx context.Context, params *CreatePlanGroupParams) (*PlanGroup, error) {
	body := buildBody(map[string]any{
		"name":        params.Name,
		"description": params.Description,
		"is_public":   params.IsPublic,
	})
	return parseDirectResponse[PlanGroup](r.http.post(ctx, "/plan-groups", body, params.IdempotencyKey))
}
