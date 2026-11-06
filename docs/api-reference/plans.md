# Plans

API version: `2026-08-27`

## PublishDraft

`client.Plans.PublishDraft(ctx, ...)`

`POST /plans/{id}/drafts/{draftId}/publish` · operation `publish-plan-draft`

Publish a draft as the plan's current version, as a new experiment with its own code, or as the next version of the experiment it started from. A country missing a price, or removing the signup trial of the plan's group from the current version, stops the publication. With move, subscribers of earlier versions move at their renewal or now; if any of them does not fit the new version, nothing is published and the response lists each one with what would make it fit, to send back in changes. Paused subscribers stay where they are. Send preview to check without publishing.

### Parameters

- `ID` (`string`, required)
- `DraftID` (`string`, required)
- `PublishAs` (`string`, required) — current sells the draft to every new buyer.
- `Move` (`string`, optional) — What happens to subscribers of earlier versions: keep them, move them at their next renewal, or move them now (features and balance change today, the price at renewal). Nobody moves unless every one of them fits or has a change below.
- `Changes` (`[]PublishPlanDraftParamsChangesItem`, optional) — Per-subscriber changes that make each one fit, applied at its renewal.
- `Preview` (`bool`, optional) — Check everything without publishing.
- `ExperimentCode` (`string`, optional) — The experiment's own code, after the plan's code.

### Valid parameter combinations

- `PublishAs`
- `PublishAs` + `ExperimentCode`
- `PublishAs`

### Request options

- `IdempotencyKey` (`string`, optional) — Unique key used to safely retry this write for 24 hours without applying it twice.

### Returns

`PlanDraftPublication`

## GetDraft

`client.Plans.GetDraft(ctx, ...)`

`GET /plans/{id}/drafts/{draftId}` · operation `get-plan-draft`

Retrieve a draft with the version it started from and what it changes.

### Parameters

- `ID` (`string`, required)
- `DraftID` (`string`, required)

### Returns

`PlanDraft`

## UpdateDraft

`client.Plans.UpdateDraft(ctx, ...)`

`PATCH /plans/{id}/drafts/{draftId}` · operation `update-plan-draft`

Change a draft's prices per interval, trials, the interval shown first, countries with their currency and prices, and credit or balance rules, or undo single changes. Each block sent replaces that block; blocks left out stay as they are. Buyers and subscribers see nothing until the draft is published.

### Parameters

- `ID` (`string`, required)
- `DraftID` (`string`, required)
- `Prices` (`[]UpdatePlanDraftParamsPricesItem`, optional) — Every interval the draft sells. An interval the draft sold and this list leaves out is removed from the draft.
- `Countries` (`[]UpdatePlanDraftParamsCountriesItem`, optional) — Every country with its own currency or prices. A country the draft sold and this list leaves out is removed. A country missing a price blocks publishing.
- `Rules` (`UpdatePlanDraftParamsRules`, optional) — Credit or balance rules; fields left out stay as they are.
- `Undo` (`[]string`, optional) — Change keys to put back as the version the draft started from had them.

### Request options

- `IdempotencyKey` (`string`, optional) — Unique key used to safely retry this write for 24 hours without applying it twice.

### Returns

`PlanDraft`

## DeleteDraft

`client.Plans.DeleteDraft(ctx, ...)`

`DELETE /plans/{id}/drafts/{draftId}` · operation `delete-plan-draft`

Discard a draft. Nothing it changed was ever sold.

### Parameters

- `ID` (`string`, required)
- `DraftID` (`string`, required)

### Returns

`DeletedObject`

## ListDrafts

`client.Plans.ListDrafts(ctx, ...)`

`GET /plans/{id}/drafts` · operation `list-plan-drafts`

List the plan's open drafts, oldest first, each with the version it started from and what it changes. Drafts sell nothing until published.

### Parameters

- `ID` (`string`, required)

### Returns

`PlansListDraftsResult`

## CreateDraft

`client.Plans.CreateDraft(ctx, ...)`

`POST /plans/{id}/drafts` · operation `create-plan-draft`

Start a draft from a version of the plan: the current one, an experiment's version or an older one. A plan can have any number of drafts. Edit it, then publish it as the current version, as an experiment or as an experiment's next version; nothing changes for buyers or subscribers until then.

### Parameters

- `ID` (`string`, required)
- `FromVersionID` (`string`, optional) — The version to start from: the current version, an experiment's version or an older version. Defaults to the current version.

### Request options

- `IdempotencyKey` (`string`, optional) — Unique key used to safely retry this write for 24 hours without applying it twice.

### Returns

`PlanDraft`

## ArchiveExperiment

`client.Plans.ArchiveExperiment(ctx, ...)`

`POST /plans/{id}/experiments/{code}/archive` · operation `archive-plan-experiment`

Stop selling the experiment. Its code sells the plan's current version from then on, so links already shared keep working. Choose what happens to its subscribers: keep them on its versions, or move them to the current version at their renewal or now. If any of them does not fit, nothing changes and the response lists each one with what would make it fit, to send back in changes.

### Parameters

- `ID` (`string`, required)
- `Code` (`string`, required)
- `Subscribers` (`string`, optional) — What happens to the experiment's subscribers: keep them on its versions, move them to the plan's current version at their next renewal, or move them now. Nobody moves, and the experiment keeps running, unless every one of them fits or has a change.
- `Changes` (`[]ArchivePlanExperimentParamsChangesItem`, optional) — Per-subscriber changes that make each one fit, applied at its renewal.
- `Preview` (`bool`, optional) — Check the archive and the move without changing anything.

### Request options

- `IdempotencyKey` (`string`, optional) — Unique key used to safely retry this write for 24 hours without applying it twice.

### Returns

`PlanExperimentArchive`

## MakeExperimentCurrent

`client.Plans.MakeExperimentCurrent(ctx, ...)`

`POST /plans/{id}/experiments/{code}/make-current` · operation `make-plan-experiment-current`

Make the experiment's latest version the plan's current version, with exactly its content and the next version number. New buyers get it; its subscribers are already on it. The experiment closes, and its code sells the plan's current version from then on. Nobody else moves.

### Parameters

- `ID` (`string`, required)
- `Code` (`string`, required)

### Request options

- `IdempotencyKey` (`string`, optional) — Unique key used to safely retry this write for 24 hours without applying it twice.

### Returns

`PlanVersion`

## GetExperiment

`client.Plans.GetExperiment(ctx, ...)`

`GET /plans/{id}/experiments/{code}` · operation `get-plan-experiment`

Get one experiment of a plan by its own code, with the version it sells and how many subscriptions are on it.

### Parameters

- `ID` (`string`, required)
- `Code` (`string`, required)

### Returns

`PlanExperiment`

## CreateExperimentVersion

`client.Plans.CreateExperimentVersion(ctx, ...)`

`POST /plans/{id}/experiments/{code}/versions` · operation `create-plan-experiment-version`

Publish the experiment's next version from its latest version with the changes given. Its code sells the new version from now on; subscribers on its earlier versions stay where they are until moved.

### Parameters

- `ID` (`string`, required)
- `Code` (`string`, required)
- `Prices` (`[]CreatePlanExperimentVersionParamsPricesItem`, optional)
- `Currencies` (`[]CreatePlanExperimentVersionParamsCurrenciesItem`, optional) — Currencies the current version does not sell, with the exchange rate (local units per USD) used for the amounts not listed, credit purchases, add-ons and overage.
- `MaxPlanCredits` (`int | null`, optional) — Credits plans that roll over unused credits: the most plan credits a subscription can hold. Null removes the cap.
- `MaxPlanBalance` (`int | null`, optional) — Balance plans that roll over unused balance: the most plan balance a subscription can hold, in USD cents. Other currencies convert it with the plan's exchange rate unless the dashboard sets a custom amount for that currency. Null removes the cap.

### Request options

- `IdempotencyKey` (`string`, optional) — Unique key used to safely retry this write for 24 hours without applying it twice.

### Returns

`PlanVersion`

## ListExperiments

`client.Plans.ListExperiments(ctx, ...)`

`GET /plans/{id}/experiments` · operation `list-plan-experiments`

List a plan's experiments, newest first, with the version each one sells and how many subscriptions are on them.

### Parameters

- `ID` (`string`, required)

### Returns

`PlansListExperimentsResult`

## CreateExperiment

`client.Plans.CreateExperiment(ctx, ...)`

`POST /plans/{id}/experiments` · operation `create-plan-experiment`

Start an experiment from the plan's current version with the changes given. It is sold only through its code (planCode `<plan code>:<code>`); new buyers of the plan keep getting the current version.

### Parameters

- `ID` (`string`, required)
- `Prices` (`[]CreatePlanExperimentParamsPricesItem`, optional)
- `Currencies` (`[]CreatePlanExperimentParamsCurrenciesItem`, optional) — Currencies the current version does not sell, with the exchange rate (local units per USD) used for the amounts not listed, credit purchases, add-ons and overage.
- `MaxPlanCredits` (`int | null`, optional) — Credits plans that roll over unused credits: the most plan credits a subscription can hold. Null removes the cap.
- `MaxPlanBalance` (`int | null`, optional) — Balance plans that roll over unused balance: the most plan balance a subscription can hold, in USD cents. Other currencies convert it with the plan's exchange rate unless the dashboard sets a custom amount for that currency. Null removes the cap.
- `Code` (`string`, required) — The experiment's own code: lowercase letters, numbers, hyphens and underscores. It is sold as planCode `<plan code>:<code>` and is never reused in the plan.

### Request options

- `IdempotencyKey` (`string`, optional) — Unique key used to safely retry this write for 24 hours without applying it twice.

### Returns

`PlanExperiment`

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

## CancelMove

`client.Plans.CancelMove(ctx, ...)`

`POST /plans/{id}/moves/{moveId}/cancel` · operation `cancel-plan-move`

Cancel a move for every subscriber who has not moved yet. They stay on their version; those already moved stay moved.

### Parameters

- `ID` (`string`, required)
- `MoveID` (`string`, required)

### Request options

- `IdempotencyKey` (`string`, optional) — Unique key used to safely retry this write for 24 hours without applying it twice.

### Returns

`PlanMove`

## GetMove

`client.Plans.GetMove(ctx, ...)`

`GET /plans/{id}/moves/{moveId}` · operation `get-plan-move`

Read a move of subscribers between plan versions with each subscriber's status: moving, moved, needs a change, left out, canceled or kept on its version.

### Parameters

- `ID` (`string`, required)
- `MoveID` (`string`, required)

### Returns

`PlanMove`

## CancelMoveSubscription

`client.Plans.CancelMoveSubscription(ctx, ...)`

`POST /plans/{id}/moves/{moveId}/subscriptions/{subscriptionId}/cancel` · operation `cancel-plan-move-subscription`

Cancel a move for one subscriber who has not moved yet. It stays on its version; the rest of the move continues.

### Parameters

- `ID` (`string`, required)
- `MoveID` (`string`, required)
- `SubscriptionID` (`string`, required)

### Request options

- `IdempotencyKey` (`string`, optional) — Unique key used to safely retry this write for 24 hours without applying it twice.

### Returns

`PlanMove`

## ListMoves

`client.Plans.ListMoves(ctx, ...)`

`GET /plans/{id}/moves` · operation `list-plan-moves`

List a plan's latest moves of subscribers between versions, newest first, with how many are moving, moved, need a change, were left out, canceled or kept on their version.

### Parameters

- `ID` (`string`, required)

### Returns

`PlansListMovesResult`

## CreateMove

`client.Plans.CreateMove(ctx, ...)`

`POST /plans/{id}/moves` · operation `create-plan-move`

Move a list of subscriptions, or every subscription of another version, to a newer version of the plan, at their next renewal or now. Nobody moves unless every one of them fits or has a change in changes; otherwise the response lists each one with what would make it fit. Only active, trialing and past due subscriptions move. Paused ones, and listed subscriptions in another status or already on the destination, are left out. Send preview to check without moving anyone.

### Parameters

- `ID` (`string`, required)
- `SubscriptionIds` (`[]string`, optional) — Public IDs of the subscriptions to move.
- `ToVersionID` (`string`, required) — The version subscribers move to.
- `Effective` (`string`, required) — now changes features and the included amount today and the price at renewal; period_end moves each subscriber at its next renewal.
- `Changes` (`[]CreatePlanMoveParamsChangesItem`, optional) — Per-subscriber changes that make each one fit, applied at its renewal.
- `Preview` (`bool`, optional) — Check every subscriber without moving anyone.
- `FromVersionID` (`string`, optional) — Move every active, trialing and past due subscription on this version of the plan.

### Valid parameter combinations

- `SubscriptionIds` + `ToVersionID` + `Effective`
- `FromVersionID` + `ToVersionID` + `Effective`

### Request options

- `IdempotencyKey` (`string`, optional) — Unique key used to safely retry this write for 24 hours without applying it twice.

### Returns

`PlanMove`

## KeepSubscriberOnVersion

`client.Plans.KeepSubscriberOnVersion(ctx, ...)`

`POST /plans/{id}/needs-a-change/{subscriptionId}/keep` · operation `keep-plan-subscriber-on-version`

Leave a subscriber that needs a change on its version. It stops appearing among the subscribers that need a change; move it later like any other subscriber.

### Parameters

- `ID` (`string`, required)
- `SubscriptionID` (`string`, required)

### Request options

- `IdempotencyKey` (`string`, optional) — Unique key used to safely retry this write for 24 hours without applying it twice.

### Returns

`KeptOnVersion`

## ListSubscribersNeedingAChange

`client.Plans.ListSubscribersNeedingAChange(ctx, ...)`

`GET /plans/{id}/needs-a-change` · operation `list-plan-subscribers-needing-a-change`

List the plan's subscribers whose move could not complete, with the version they stay on, the one they could not move to and why. Move them again with a change, or keep them on their version so they leave this list.

### Parameters

- `ID` (`string`, required)

### Returns

`PlansListSubscribersNeedingAChangeResult`

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

Archive a price for new subscriptions. This publishes a new version of the plan; existing subscriptions that selected the price keep paying it on their version until moved.

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
- `UnusedPlanCredits` (`UpdatePlanParamsUnusedPlanCredits`, optional) — Credits plans only. What happens to unused plan credits at each reset. Changing it publishes a new version of the plan; each subscription stays on its version until moved.
- `PlanCreditLimit` (`UpdatePlanParamsPlanCreditLimit | null`, optional) — Credits plans only. The most plan credits a subscription spends per day or per week; null removes the limit. Changing it publishes a new version of the plan; each subscription stays on its version until moved.
- `UnusedPlanBalance` (`UpdatePlanParamsUnusedPlanBalance`, optional) — Balance plans only. What happens to unused plan balance at each reset. Changing it publishes a new version of the plan; each subscription stays on its version until moved.
- `PlanBalanceLimit` (`UpdatePlanParamsPlanBalanceLimit | null`, optional) — Balance plans only. The most plan balance a subscription spends per day or per week, in USD cents; null removes the limit. Changing it publishes a new version of the plan; each subscription stays on its version until moved.
- `BlockOnExhaustion` (`bool`, optional) — Credits and balance plans only. false lets usage continue past the plan and purchased amounts and bills the missing amount as overage at the next renewal; a credits plan bills it at the price per credit and needs one. Trials and free plans always stop, so free plans accept only true. Changing it publishes a new version of the plan; each subscription stays on its version until moved.
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

## CompareVersions

`client.Plans.CompareVersions(ctx, ...)`

`GET /plans/{id}/versions/{versionId}/changes` · operation `compare-plan-versions`

List what changes from one version of a plan to another: prices, trials, features, rules, add-ons, credit packs and countries.

### Parameters

- `ID` (`string`, required)
- `VersionID` (`string`, required)
- `Against` (`string`, required)

### Returns

`PlanVersionComparison`

## GetVersion

`client.Plans.GetVersion(ctx, ...)`

`GET /plans/{id}/versions/{versionId}` · operation `get-plan-version`

Get one published version of a plan with its prices and how many subscriptions are on it.

### Parameters

- `ID` (`string`, required)
- `VersionID` (`string`, required)

### Returns

`PlanVersion`

## ListVersions

`client.Plans.ListVersions(ctx, ...)`

`GET /plans/{id}/versions` · operation `list-plan-versions`

List a plan's published versions, newest first, with their prices and how many subscriptions are on each. Versions published before Commet recorded every term are not listed.

### Parameters

- `ID` (`string`, required)
- `Cursor` (`string`, optional)
- `Limit` (`int`, optional)

### Returns

`PlansListVersionsResult`

## CreateVersion

`client.Plans.CreateVersion(ctx, ...)`

`POST /plans/{id}/versions` · operation `create-plan-version`

Publish a new current version of the plan from its current version, changing prices per interval, included credits or balance, per-currency amounts and the rollover cap. New buyers get it; every subscriber stays on its version until moved. To try changes without selling them to everyone, create an experiment instead.

### Parameters

- `ID` (`string`, required)
- `Prices` (`[]CreatePlanVersionParamsPricesItem`, optional)
- `Currencies` (`[]CreatePlanVersionParamsCurrenciesItem`, optional) — Currencies the current version does not sell, with the exchange rate (local units per USD) used for the amounts not listed, credit purchases, add-ons and overage.
- `MaxPlanCredits` (`int | null`, optional) — Credits plans that roll over unused credits: the most plan credits a subscription can hold. Null removes the cap.
- `MaxPlanBalance` (`int | null`, optional) — Balance plans that roll over unused balance: the most plan balance a subscription can hold, in USD cents. Other currencies convert it with the plan's exchange rate unless the dashboard sets a custom amount for that currency. Null removes the cap.

### Request options

- `IdempotencyKey` (`string`, optional) — Unique key used to safely retry this write for 24 hours without applying it twice.

### Returns

`PlanVersion`

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
