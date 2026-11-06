# Plans

API version: `2026-08-27`

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
- `TrialRequiresPaymentMethod` (`bool`, optional) — Set to false so the price's free trial starts without a payment method. Requires trialDays on the price.
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
- `TrialRequiresPaymentMethod` (`bool`, optional) — Set to false so the price's free trial starts without a payment method. Requires trialDays on the price.
- `IsDefault` (`bool`, optional)
- `IncludedBalance` (`int | null`, optional)
- `IncludedCredits` (`int | null`, optional)
- `MarketPrices` (`[]AddPlanPriceParamsMarketPricesItem`, optional)
- `InheritsFromPriceID` (`string`, optional)

### Valid parameter combinations

- `BillingInterval` + `Price`
- `BillingInterval` + `InheritsFromPriceID` + `MarketPrices`

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

Update a plan's name, description, visibility, default status, metadata, overage behavior, what happens to its unused plan credits or plan balance, or its plan credit or plan balance limit.

### Parameters

- `ID` (`string`, required)
- `Name` (`string`, optional)
- `Description` (`string | null`, optional)
- `Metadata` (`map[string]any`, optional)
- `IsPublic` (`bool`, optional)
- `UnusedPlanCredits` (`UpdatePlanParamsUnusedPlanCredits`, optional) — Credits plans only. What happens to unused plan credits at each reset. Existing subscriptions apply it from their next reset.
- `PlanCreditLimit` (`UpdatePlanParamsPlanCreditLimit | null`, optional) — Credits plans only. The most plan credits a subscription spends per day or per week; null removes the limit. Raising or removing it applies to existing subscriptions at once; lowering it, adding it or changing its interval applies from their next period.
- `UnusedPlanBalance` (`UpdatePlanParamsUnusedPlanBalance`, optional) — Balance plans only. What happens to unused plan balance at each reset. Existing subscriptions apply it from their next reset.
- `PlanBalanceLimit` (`UpdatePlanParamsPlanBalanceLimit | null`, optional) — Balance plans only. The most plan balance a subscription spends per day or per week, in USD cents; null removes the limit. Raising or removing it applies to existing subscriptions at once; lowering it, adding it or changing its interval applies from their next period.
- `BlockOnExhaustion` (`bool`, optional) — Credits and balance plans only. false lets usage continue past the plan and purchased amounts and bills the missing amount as overage at the next renewal; a credits plan bills it at the price per credit and needs one. Trials and free plans always stop, so free plans accept only true. Existing subscriptions apply it from their next reset.
- `IsDefault` (`bool`, optional) — Make this plan the default plan of its plan group: customers created with that planGroupId start on it, and a paid plan of the group that ends from the customer portal, a revoked payment authorization or exhausted dunning moves to it instead of canceling. Requires a free plan in a plan group; replaces the group's previous default plan. false unmarks it.

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

Create a new plan with optional consumption model, visibility, plan group assignment, overage behavior, and the rollover and spend limit of its plan credits or plan balance.

### Parameters

- `Name` (`string`, required)
- `Code` (`string`, required)
- `Description` (`string`, optional)
- `ConsumptionModel` (`string`, optional)
- `IsPublic` (`bool`, optional)
- `IsFree` (`bool`, optional)
- `BlockOnExhaustion` (`bool`, optional) — Credits and balance plans. Omitted, credits plans and free plans stop usage when the plan and purchased amounts run out, and paid balance plans bill the missing balance as overage. A new credits plan always stops usage: it allows overage only with a price per credit, which is set after creation, and then `PATCH /v1/plans/{id}` turns overage on.
- `UnusedPlanCredits` (`CreatePlanParamsUnusedPlanCredits`, optional) — Credits plans only. What happens to unused plan credits at each reset; omitted means they expire.
- `PlanCreditLimit` (`CreatePlanParamsPlanCreditLimit | null`, optional) — Credits plans only. The most plan credits a subscription spends per day or per week; omitted or null means no limit.
- `UnusedPlanBalance` (`CreatePlanParamsUnusedPlanBalance`, optional) — Balance plans only. What happens to unused plan balance at each reset; omitted means it expires.
- `PlanBalanceLimit` (`CreatePlanParamsPlanBalanceLimit | null`, optional) — Balance plans only. The most plan balance a subscription spends per day or per week, in USD cents; omitted or null means no limit.
- `PlanGroupID` (`string`, optional)
- `IsDefault` (`bool`, optional) — Make this plan the default plan of its plan group: customers created with that planGroupId start on it, and a paid plan of the group that ends from the customer portal, a revoked payment authorization or exhausted dunning moves to it instead of canceling. Requires a free plan with a planGroupId; replaces the group's previous default plan.
- `Metadata` (`map[string]any`, optional)

### Request options

- `IdempotencyKey` (`string`, optional) — Unique key used to safely retry this write for 24 hours without applying it twice.

### Returns

`Plan`
