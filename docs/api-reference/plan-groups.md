# Plan Groups

API version: `2026-07-31`

## RemovePlan

`client.PlanGroups.RemovePlan(ctx, ...)`

`DELETE /plan-groups/{id}/plans/{planId}` · operation `remove-plan-from-group`

Remove a plan from a plan group.

### Parameters

- `ID` (`string`, required)
- `PlanID` (`string`, required)

### Returns

`RemovedPlanFromGroup`

## ReorderPlans

`client.PlanGroups.ReorderPlans(ctx, ...)`

`PUT /plan-groups/{id}/plans/reorder` · operation `reorder-plans-in-group`

Set the display order of plans within a group. All plan IDs in the group must be provided.

### Parameters

- `ID` (`string`, required)
- `PlanIds` (`[]string`, required)

### Request options

- `IdempotencyKey` (`string`, optional) — Unique key used to safely retry this write for 24 hours without applying it twice.

### Returns

`ReorderedPlans`

## AddPlan

`client.PlanGroups.AddPlan(ctx, ...)`

`POST /plan-groups/{id}/plans` · operation `add-plan-to-group`

Add an existing plan to a plan group with optional sort order.

### Parameters

- `ID` (`string`, required)
- `PlanID` (`string`, required)
- `SortOrder` (`int`, optional)

### Request options

- `IdempotencyKey` (`string`, optional) — Unique key used to safely retry this write for 24 hours without applying it twice.

### Returns

`AddedPlanToGroup`

## Get

`client.PlanGroups.Get(ctx, ...)`

`GET /plan-groups/{id}` · operation `get-plan-group`

Retrieve a plan group by ID, including its plans ordered by sortOrder.

### Parameters

- `ID` (`string`, required)

### Returns

`PlanGroupDetail`

## Update

`client.PlanGroups.Update(ctx, ...)`

`PATCH /plan-groups/{id}` · operation `update-plan-group`

Update a plan group's name, description, or visibility.

### Parameters

- `ID` (`string`, required)
- `Name` (`string`, optional)
- `Description` (`string | null`, optional)
- `IsPublic` (`bool`, optional)

### Request options

- `IdempotencyKey` (`string`, optional) — Unique key used to safely retry this write for 24 hours without applying it twice.

### Returns

`PlanGroup`

## Delete

`client.PlanGroups.Delete(ctx, ...)`

`DELETE /plan-groups/{id}` · operation `delete-plan-group`

Delete a plan group. Plans in the group are unlinked, not deleted.

### Parameters

- `ID` (`string`, required)

### Returns

`DeletedObject`

## List

`client.PlanGroups.List(ctx, ...)`

`GET /plan-groups` · operation `list-plan-groups`

List plan groups with cursor-based pagination.

### Parameters

- `Cursor` (`string`, optional)
- `Limit` (`int`, optional)

### Returns

`PlanGroupsListResult`

## Create

`client.PlanGroups.Create(ctx, ...)`

`POST /plan-groups` · operation `create-plan-group`

Create a new plan group for organizing plans.

### Parameters

- `Name` (`string`, required)
- `Description` (`string`, optional)
- `IsPublic` (`bool`, optional)

### Request options

- `IdempotencyKey` (`string`, optional) — Unique key used to safely retry this write for 24 hours without applying it twice.

### Returns

`PlanGroup`
