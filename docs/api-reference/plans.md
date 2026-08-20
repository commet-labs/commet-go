# Plans

API version: `2026-07-31`

## UpdateFeature

`client.Plans.UpdateFeature(ctx, ...)`

`PATCH /plans/{id}/features/{featureId}` · operation `update-plan-feature`

Update limits, overage, or enabled status of a feature on a plan.

### Parameters

- `ID` (`string`, required)
- `FeatureID` (`string`, required)
- `Enabled` (`bool`, optional)
- `IncludedAmount` (`int`, optional)
- `Unlimited` (`bool`, optional)
- `Overage` (`UpdatePlanFeatureParamsOverage`, optional)
- `CreditsPerUnit` (`int | null`, optional)

### Request options

- `IdempotencyKey` (`string`, optional) — Unique key used to safely retry this write for 24 hours without applying it twice.

### Returns

`PlanFeature`

## RemoveFeature

`client.Plans.RemoveFeature(ctx, ...)`

`DELETE /plans/{id}/features/{featureId}` · operation `remove-plan-feature`

Detach a feature from a plan.

### Parameters

- `ID` (`string`, required)
- `FeatureID` (`string`, required)

### Returns

`RemovedPlanFeature`

## AddFeature

`client.Plans.AddFeature(ctx, ...)`

`POST /plans/{id}/features` · operation `add-plan-feature`

Attach a feature to a plan with limits, overage, and credits configuration.

### Parameters

- `ID` (`string`, required)
- `FeatureID` (`string`, required)
- `Enabled` (`bool`, optional)
- `IncludedAmount` (`int`, optional)
- `Unlimited` (`bool`, optional)
- `Overage` (`AddPlanFeatureParamsOverage`, optional)
- `CreditsPerUnit` (`int | null`, optional)
- `PricingMode` (`string`, optional)
- `Margin` (`int | null`, optional)

### Request options

- `IdempotencyKey` (`string`, optional) — Unique key used to safely retry this write for 24 hours without applying it twice.

### Returns

`PlanFeature`

## SetDefaultPrice

`client.Plans.SetDefaultPrice(ctx, ...)`

`PUT /plans/{id}/prices/{priceId}/default` · operation `set-default-plan-price`

Set a specific price as the default and return the updated plan price.

### Parameters

- `ID` (`string`, required)
- `PriceID` (`string`, required)

### Request options

- `IdempotencyKey` (`string`, optional) — Unique key used to safely retry this write for 24 hours without applying it twice.

### Returns

`PlanPrice`

## SetRegionalPrices

`client.Plans.SetRegionalPrices(ctx, ...)`

`PUT /plans/{id}/prices/{priceId}/regional` · operation `upsert-regional-prices`

Create or update regional currency price overrides for a plan price.

### Parameters

- `ID` (`string`, required)
- `PriceID` (`string`, required)
- `Overrides` (`[]UpsertRegionalPricesParamsOverridesItem`, required)

### Request options

- `IdempotencyKey` (`string`, optional) — Unique key used to safely retry this write for 24 hours without applying it twice.

### Returns

`PlanRegionalPricing`

## DeleteRegionalPrices

`client.Plans.DeleteRegionalPrices(ctx, ...)`

`DELETE /plans/{id}/prices/{priceId}/regional` · operation `delete-regional-prices`

Remove all regional currency overrides for a plan price. The request is rejected while billable subscriptions depend on an override.

### Parameters

- `ID` (`string`, required)
- `PriceID` (`string`, required)

### Returns

`DeletedPlanRegionalPricing`

## UpdatePrice

`client.Plans.UpdatePrice(ctx, ...)`

`PATCH /plans/{id}/prices/{priceId}` · operation `update-plan-price`

Update a base price or market price variant. Removing a base market override is rejected while a variant depends on it. Offer terms are managed through Offers.

### Parameters

- `ID` (`string`, required)
- `PriceID` (`string`, required)
- `Price` (`int`, optional)
- `IsDefault` (`bool`, optional)
- `TrialDays` (`int`, optional)
- `IncludedBalance` (`int | null`, optional)
- `IncludedCredits` (`int | null`, optional)
- `Metadata` (`map[string]any`, optional) — Metadata keys to merge into the existing price metadata.
- `MarketPrices` (`[]UpdatePlanPriceParamsMarketPricesItem`, optional)

### Request options

- `IdempotencyKey` (`string`, optional) — Unique key used to safely retry this write for 24 hours without applying it twice.

### Returns

`PlanPrice`

## DeletePrice

`client.Plans.DeletePrice(ctx, ...)`

`DELETE /plans/{id}/prices/{priceId}` · operation `delete-plan-price`

Archive a price for new subscriptions. Existing subscriptions that selected it continue using its current catalog value.

### Parameters

- `ID` (`string`, required)
- `PriceID` (`string`, required)

### Returns

`DeletedObject`

## AddPrice

`client.Plans.AddPrice(ctx, ...)`

`POST /plans/{id}/prices` · operation `add-plan-price`

Add a base price or a selectable market price variant. Variants inherit their base price outside the markets they override. Configure introductory and promotional benefits through Offers.

### Parameters

- `ID` (`string`, required)
- `BillingInterval` (`string`, required)
- `Metadata` (`map[string]any`, optional)
- `Price` (`int`, optional)
- `TrialDays` (`int`, optional)
- `IsDefault` (`bool`, optional)
- `IncludedBalance` (`int | null`, optional)
- `IncludedCredits` (`int | null`, optional)
- `MarketPrices` (`[]AddPlanPriceParamsMarketPricesItem`, optional)
- `InheritsFromPriceID` (`string`, optional)

### Request options

- `IdempotencyKey` (`string`, optional) — Unique key used to safely retry this write for 24 hours without applying it twice.

### Returns

`PlanPrice`

## SetRegionalPricing

`client.Plans.SetRegionalPricing(ctx, ...)`

`PUT /plans/{id}/regional` · operation `set-plan-regional-pricing`

Configure regional prices and feature overage values for one currency. Currency-specific offer terms are managed through Offers.

### Parameters

- `ID` (`string`, required)
- `Currency` (`string`, required)
- `ExchangeRate` (`float64`, required)
- `Prices` (`[]SetPlanRegionalPricingParamsPricesItem`, optional)
- `Features` (`[]SetPlanRegionalPricingParamsFeaturesItem`, optional)

### Request options

- `IdempotencyKey` (`string`, optional) — Unique key used to safely retry this write for 24 hours without applying it twice.

### Returns

`PlanRegionalPricingResult`

## Get

`client.Plans.Get(ctx, ...)`

`GET /plans/{id}` · operation `get-plan`

Get a plan with public price IDs and their automatic introductory offer IDs.

### Parameters

- `ID` (`string`, required)

### Returns

`Plan`

## Update

`client.Plans.Update(ctx, ...)`

`PATCH /plans/{id}` · operation `update-plan`

Update a plan's name, description, visibility, or metadata.

### Parameters

- `ID` (`string`, required)
- `Name` (`string`, optional)
- `Description` (`string | null`, optional)
- `Metadata` (`map[string]any`, optional)
- `IsPublic` (`bool`, optional)

### Request options

- `IdempotencyKey` (`string`, optional) — Unique key used to safely retry this write for 24 hours without applying it twice.

### Returns

`Plan`

## Delete

`client.Plans.Delete(ctx, ...)`

`DELETE /plans/{id}` · operation `delete-plan`

Soft-delete a plan.

### Parameters

- `ID` (`string`, required)

### Returns

`DeletedObject`

## SetVisibility

`client.Plans.SetVisibility(ctx, ...)`

`PUT /plans/{id}/visibility` · operation `set-plan-visibility`

Set a plan's public visibility and return the updated plan.

### Parameters

- `ID` (`string`, required)
- `IsPublic` (`bool`, required)

### Request options

- `IdempotencyKey` (`string`, optional) — Unique key used to safely retry this write for 24 hours without applying it twice.

### Returns

`Plan`

## List

`client.Plans.List(ctx, ...)`

`GET /plans` · operation `list-plans`

List plans with public price IDs and their automatic introductory offer IDs.

### Parameters

- `IncludePrivate` (`bool`, optional)

### Returns

`PlansListResult`

## Create

`client.Plans.Create(ctx, ...)`

`POST /plans` · operation `create-plan`

Create a new plan with optional consumption model, visibility, and plan group assignment.

### Parameters

- `Name` (`string`, required)
- `Code` (`string`, required)
- `Description` (`string`, optional)
- `ConsumptionModel` (`string`, optional)
- `IsPublic` (`bool`, optional)
- `IsFree` (`bool`, optional)
- `BlockOnExhaustion` (`bool`, optional)
- `PlanGroupID` (`string`, optional)
- `Metadata` (`map[string]any`, optional)

### Request options

- `IdempotencyKey` (`string`, optional) — Unique key used to safely retry this write for 24 hours without applying it twice.

### Returns

`Plan`
