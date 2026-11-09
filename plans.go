package commet

import (
	"context"
	"fmt"
)

type PublishPlanDraftParams struct {
	PublishAs      string                              `json:"publish_as"`
	Move           *string                             `json:"move,omitempty"`
	Changes        []PublishPlanDraftParamsChangesItem `json:"changes,omitempty"`
	Preview        *bool                               `json:"preview,omitempty"`
	ExperimentCode *string                             `json:"experiment_code,omitempty"`
	IdempotencyKey string                              `json:"-"`
}

type UpdatePlanDraftParams struct {
	Prices         []UpdatePlanDraftParamsPricesItem    `json:"prices,omitempty"`
	Countries      []UpdatePlanDraftParamsCountriesItem `json:"countries,omitempty"`
	Rules          *UpdatePlanDraftParamsRules          `json:"rules,omitempty"`
	Undo           []string                             `json:"undo,omitempty"`
	IdempotencyKey string                               `json:"-"`
}

type CreatePlanDraftParams struct {
	FromVersionID  *string `json:"from_version_id,omitempty"`
	IdempotencyKey string  `json:"-"`
}

type ArchivePlanExperimentParams struct {
	Subscribers    *string                                  `json:"subscribers,omitempty"`
	Changes        []ArchivePlanExperimentParamsChangesItem `json:"changes,omitempty"`
	Preview        *bool                                    `json:"preview,omitempty"`
	IdempotencyKey string                                   `json:"-"`
}

type MakePlanExperimentCurrentParams struct {
	IdempotencyKey string `json:"-"`
}

type CreatePlanExperimentVersionParams struct {
	Prices         []CreatePlanExperimentVersionParamsPricesItem     `json:"prices,omitempty"`
	Currencies     []CreatePlanExperimentVersionParamsCurrenciesItem `json:"currencies,omitempty"`
	MaxPlanCredits *int                                              `json:"max_plan_credits,omitempty"`
	MaxPlanBalance *int                                              `json:"max_plan_balance,omitempty"`
	IdempotencyKey string                                            `json:"-"`
}

type CreatePlanExperimentParams struct {
	Prices         []CreatePlanExperimentParamsPricesItem     `json:"prices,omitempty"`
	Currencies     []CreatePlanExperimentParamsCurrenciesItem `json:"currencies,omitempty"`
	MaxPlanCredits *int                                       `json:"max_plan_credits,omitempty"`
	MaxPlanBalance *int                                       `json:"max_plan_balance,omitempty"`
	Code           string                                     `json:"code"`
	IdempotencyKey string                                     `json:"-"`
}

type UpdatePlanFeatureParams struct {
	Enabled        *bool                           `json:"enabled,omitempty"`
	IncludedAmount *int                            `json:"included_amount,omitempty"`
	Unlimited      *bool                           `json:"unlimited,omitempty"`
	Overage        *UpdatePlanFeatureParamsOverage `json:"overage,omitempty"`
	CreditsPerUnit *int                            `json:"credits_per_unit,omitempty"`
	IdempotencyKey string                          `json:"-"`
}

type AddPlanFeatureParams struct {
	FeatureID      string                       `json:"feature_id"`
	Enabled        *bool                        `json:"enabled,omitempty"`
	IncludedAmount *int                         `json:"included_amount,omitempty"`
	Unlimited      *bool                        `json:"unlimited,omitempty"`
	Overage        *AddPlanFeatureParamsOverage `json:"overage,omitempty"`
	CreditsPerUnit *int                         `json:"credits_per_unit,omitempty"`
	PricingMode    *string                      `json:"pricing_mode,omitempty"`
	Margin         *int                         `json:"margin,omitempty"`
	IdempotencyKey string                       `json:"-"`
}

type CancelPlanMoveParams struct {
	IdempotencyKey string `json:"-"`
}

type CancelPlanMoveSubscriptionParams struct {
	IdempotencyKey string `json:"-"`
}

type CreatePlanMoveParams struct {
	SubscriptionIds []string                          `json:"subscription_ids,omitempty"`
	ToVersionID     string                            `json:"to_version_id"`
	Effective       string                            `json:"effective"`
	Changes         []CreatePlanMoveParamsChangesItem `json:"changes,omitempty"`
	Preview         *bool                             `json:"preview,omitempty"`
	FromVersionID   *string                           `json:"from_version_id,omitempty"`
	IdempotencyKey  string                            `json:"-"`
}

type KeepPlanSubscriberOnVersionParams struct {
	IdempotencyKey string `json:"-"`
}

type SetDefaultPlanPriceParams struct {
	IdempotencyKey string `json:"-"`
}

type UpsertRegionalPricesParams struct {
	Overrides      []UpsertRegionalPricesParamsOverridesItem `json:"overrides"`
	IdempotencyKey string                                    `json:"-"`
}

type UpdatePlanPriceParams struct {
	Price                      *int                                    `json:"price,omitempty"`
	IsDefault                  *bool                                   `json:"is_default,omitempty"`
	TrialDays                  *int                                    `json:"trial_days,omitempty"`
	TrialRequiresPaymentMethod *bool                                   `json:"trial_requires_payment_method,omitempty"`
	IncludedBalance            *int                                    `json:"included_balance,omitempty"`
	IncludedCredits            *int                                    `json:"included_credits,omitempty"`
	Metadata                   map[string]any                          `json:"metadata,omitempty"`
	MarketPrices               []UpdatePlanPriceParamsMarketPricesItem `json:"market_prices,omitempty"`
	IdempotencyKey             string                                  `json:"-"`
}

type AddPlanPriceParams struct {
	BillingInterval            string                               `json:"billing_interval"`
	Metadata                   map[string]any                       `json:"metadata,omitempty"`
	Price                      *int                                 `json:"price,omitempty"`
	TrialDays                  *int                                 `json:"trial_days,omitempty"`
	TrialRequiresPaymentMethod *bool                                `json:"trial_requires_payment_method,omitempty"`
	IsDefault                  *bool                                `json:"is_default,omitempty"`
	IncludedBalance            *int                                 `json:"included_balance,omitempty"`
	IncludedCredits            *int                                 `json:"included_credits,omitempty"`
	MarketPrices               []AddPlanPriceParamsMarketPricesItem `json:"market_prices,omitempty"`
	InheritsFromPriceID        *string                              `json:"inherits_from_price_id,omitempty"`
	IdempotencyKey             string                               `json:"-"`
}

type SetPlanRegionalPricingParams struct {
	Currency       string                                     `json:"currency"`
	ExchangeRate   float64                                    `json:"exchange_rate"`
	Prices         []SetPlanRegionalPricingParamsPricesItem   `json:"prices,omitempty"`
	Features       []SetPlanRegionalPricingParamsFeaturesItem `json:"features,omitempty"`
	IdempotencyKey string                                     `json:"-"`
}

type UpdatePlanParams struct {
	Name              *string                            `json:"name,omitempty"`
	Description       *string                            `json:"description,omitempty"`
	Metadata          map[string]any                     `json:"metadata,omitempty"`
	IsPublic          *bool                              `json:"is_public,omitempty"`
	UnusedPlanCredits *UpdatePlanParamsUnusedPlanCredits `json:"unused_plan_credits,omitempty"`
	PlanCreditLimit   *UpdatePlanParamsPlanCreditLimit   `json:"plan_credit_limit,omitempty"`
	UnusedPlanBalance *UpdatePlanParamsUnusedPlanBalance `json:"unused_plan_balance,omitempty"`
	PlanBalanceLimit  *UpdatePlanParamsPlanBalanceLimit  `json:"plan_balance_limit,omitempty"`
	BlockOnExhaustion *bool                              `json:"block_on_exhaustion,omitempty"`
	IsDefault         *bool                              `json:"is_default,omitempty"`
	IdempotencyKey    string                             `json:"-"`
}

type ComparePlanVersionsParams struct {
	Against string `json:"against"`
}

type ListPlanVersionsParams struct {
	Cursor *string `json:"cursor,omitempty"`
	Limit  *int    `json:"limit,omitempty"`
}

type CreatePlanVersionParams struct {
	Prices         []CreatePlanVersionParamsPricesItem     `json:"prices,omitempty"`
	Currencies     []CreatePlanVersionParamsCurrenciesItem `json:"currencies,omitempty"`
	MaxPlanCredits *int                                    `json:"max_plan_credits,omitempty"`
	MaxPlanBalance *int                                    `json:"max_plan_balance,omitempty"`
	IdempotencyKey string                                  `json:"-"`
}

type SetPlanVisibilityParams struct {
	IsPublic       bool   `json:"is_public"`
	IdempotencyKey string `json:"-"`
}

type ListPlansParams struct {
	IncludePrivate *bool `json:"include_private,omitempty"`
}

type CreatePlanParams struct {
	Name              string                             `json:"name"`
	Code              string                             `json:"code"`
	Description       *string                            `json:"description,omitempty"`
	ConsumptionModel  *string                            `json:"consumption_model,omitempty"`
	IsPublic          *bool                              `json:"is_public,omitempty"`
	IsFree            *bool                              `json:"is_free,omitempty"`
	BlockOnExhaustion *bool                              `json:"block_on_exhaustion,omitempty"`
	UnusedPlanCredits *CreatePlanParamsUnusedPlanCredits `json:"unused_plan_credits,omitempty"`
	PlanCreditLimit   *CreatePlanParamsPlanCreditLimit   `json:"plan_credit_limit,omitempty"`
	UnusedPlanBalance *CreatePlanParamsUnusedPlanBalance `json:"unused_plan_balance,omitempty"`
	PlanBalanceLimit  *CreatePlanParamsPlanBalanceLimit  `json:"plan_balance_limit,omitempty"`
	PlanGroupID       *string                            `json:"plan_group_id,omitempty"`
	IsDefault         *bool                              `json:"is_default,omitempty"`
	Metadata          map[string]any                     `json:"metadata,omitempty"`
	IdempotencyKey    string                             `json:"-"`
}

type PlansResource struct {
	http *httpClient
}

// Publish a draft as the plan's current version, as a new experiment with its own code, or as the next version of the experiment it started from. A country missing a price, or removing the signup trial of the plan's group from the current version, stops the publication. With move, subscribers of earlier versions move at their renewal or now; if any of them does not fit the new version, nothing is published and the response lists each one with what would make it fit, to send back in changes. Paused subscribers stay where they are. Send preview to check without publishing.
func (r *PlansResource) PublishDraft(ctx context.Context, id string, draftID string, params *PublishPlanDraftParams) (*PlanDraftPublication, error) {
	body := buildBody(map[string]any{
		"publish_as":      params.PublishAs,
		"move":            params.Move,
		"changes":         params.Changes,
		"preview":         params.Preview,
		"experiment_code": params.ExperimentCode,
	})
	return parseDirectResponse[PlanDraftPublication](r.http.post(ctx, fmt.Sprintf("/plans/%s/drafts/%s/publish", id, draftID), body, params.IdempotencyKey))
}

// Retrieve a draft with the version it started from and what it changes.
func (r *PlansResource) GetDraft(ctx context.Context, id string, draftID string) (*PlanDraft, error) {
	return parseDirectResponse[PlanDraft](r.http.get(ctx, fmt.Sprintf("/plans/%s/drafts/%s", id, draftID), nil))
}

// Change a draft's prices per interval, trials, the interval shown first, countries with their currency and prices, and credit or balance rules, or undo single changes. Each block sent replaces that block; blocks left out stay as they are. Buyers and subscribers see nothing until the draft is published.
func (r *PlansResource) UpdateDraft(ctx context.Context, id string, draftID string, params *UpdatePlanDraftParams) (*PlanDraft, error) {
	body := buildBody(map[string]any{
		"prices":    params.Prices,
		"countries": params.Countries,
		"rules":     params.Rules,
		"undo":      params.Undo,
	})
	return parseDirectResponse[PlanDraft](r.http.patch(ctx, fmt.Sprintf("/plans/%s/drafts/%s", id, draftID), body, params.IdempotencyKey))
}

// Discard a draft. Nothing it changed was ever sold.
func (r *PlansResource) DeleteDraft(ctx context.Context, id string, draftID string) (*DeletedObject, error) {
	return parseDirectResponse[DeletedObject](r.http.delete(ctx, fmt.Sprintf("/plans/%s/drafts/%s", id, draftID), nil, ""))
}

// List the plan's open drafts, oldest first, each with the version it started from and what it changes. Drafts sell nothing until published.
func (r *PlansResource) ListDrafts(ctx context.Context, id string) (*PlansListDraftsResult, error) {
	return parseDirectResponse[PlansListDraftsResult](r.http.get(ctx, fmt.Sprintf("/plans/%s/drafts", id), nil))
}

// Start a draft from a version of the plan: the current one, an experiment's version or an older one. A plan can have any number of drafts. Edit it, then publish it as the current version, as an experiment or as an experiment's next version; nothing changes for buyers or subscribers until then.
func (r *PlansResource) CreateDraft(ctx context.Context, id string, params *CreatePlanDraftParams) (*PlanDraft, error) {
	body := buildBody(map[string]any{
		"from_version_id": params.FromVersionID,
	})
	return parseDirectResponse[PlanDraft](r.http.post(ctx, fmt.Sprintf("/plans/%s/drafts", id), body, params.IdempotencyKey))
}

// Stop selling the experiment. Its code sells the plan's current version from then on, so links already shared keep working. Choose what happens to its subscribers: keep them on its versions, or move them to the current version at their renewal or now. If any of them does not fit, nothing changes and the response lists each one with what would make it fit, to send back in changes.
func (r *PlansResource) ArchiveExperiment(ctx context.Context, id string, code string, params *ArchivePlanExperimentParams) (*PlanExperimentArchive, error) {
	body := buildBody(map[string]any{
		"subscribers": params.Subscribers,
		"changes":     params.Changes,
		"preview":     params.Preview,
	})
	return parseDirectResponse[PlanExperimentArchive](r.http.post(ctx, fmt.Sprintf("/plans/%s/experiments/%s/archive", id, code), body, params.IdempotencyKey))
}

// Make the experiment's latest version the plan's current version, with exactly its content and the next version number. New buyers get it; its subscribers are already on it. The experiment closes, and its code sells the plan's current version from then on. Nobody else moves.
func (r *PlansResource) MakeExperimentCurrent(ctx context.Context, id string, code string, params *MakePlanExperimentCurrentParams) (*PlanVersion, error) {
	return parseDirectResponse[PlanVersion](r.http.post(ctx, fmt.Sprintf("/plans/%s/experiments/%s/make-current", id, code), map[string]any{}, params.IdempotencyKey))
}

// Get one experiment of a plan by its own code, with the version it sells and how many subscriptions are on it.
func (r *PlansResource) GetExperiment(ctx context.Context, id string, code string) (*PlanExperiment, error) {
	return parseDirectResponse[PlanExperiment](r.http.get(ctx, fmt.Sprintf("/plans/%s/experiments/%s", id, code), nil))
}

// Publish the experiment's next version from its latest version with the changes given. Its code sells the new version from now on; subscribers on its earlier versions stay where they are until moved.
func (r *PlansResource) CreateExperimentVersion(ctx context.Context, id string, code string, params *CreatePlanExperimentVersionParams) (*PlanVersion, error) {
	body := buildBody(map[string]any{
		"prices":           params.Prices,
		"currencies":       params.Currencies,
		"max_plan_credits": params.MaxPlanCredits,
		"max_plan_balance": params.MaxPlanBalance,
	})
	return parseDirectResponse[PlanVersion](r.http.post(ctx, fmt.Sprintf("/plans/%s/experiments/%s/versions", id, code), body, params.IdempotencyKey))
}

// List a plan's experiments, newest first, with the version each one sells and how many subscriptions are on them.
func (r *PlansResource) ListExperiments(ctx context.Context, id string) (*PlansListExperimentsResult, error) {
	return parseDirectResponse[PlansListExperimentsResult](r.http.get(ctx, fmt.Sprintf("/plans/%s/experiments", id), nil))
}

// Start an experiment from the plan's current version with the changes given. It is sold only through its code (planCode `<plan code>:<code>`); new buyers of the plan keep getting the current version.
func (r *PlansResource) CreateExperiment(ctx context.Context, id string, params *CreatePlanExperimentParams) (*PlanExperiment, error) {
	body := buildBody(map[string]any{
		"prices":           params.Prices,
		"currencies":       params.Currencies,
		"max_plan_credits": params.MaxPlanCredits,
		"max_plan_balance": params.MaxPlanBalance,
		"code":             params.Code,
	})
	return parseDirectResponse[PlanExperiment](r.http.post(ctx, fmt.Sprintf("/plans/%s/experiments", id), body, params.IdempotencyKey))
}

// Update limits, overage, or enabled status of a feature on a plan.
func (r *PlansResource) UpdateFeature(ctx context.Context, id string, featureID string, params *UpdatePlanFeatureParams) (*PlanFeature, error) {
	body := buildBody(map[string]any{
		"enabled":          params.Enabled,
		"included_amount":  params.IncludedAmount,
		"unlimited":        params.Unlimited,
		"overage":          params.Overage,
		"credits_per_unit": params.CreditsPerUnit,
	})
	return parseDirectResponse[PlanFeature](r.http.patch(ctx, fmt.Sprintf("/plans/%s/features/%s", id, featureID), body, params.IdempotencyKey))
}

// Detach a feature from a plan.
func (r *PlansResource) RemoveFeature(ctx context.Context, id string, featureID string) (*RemovedPlanFeature, error) {
	return parseDirectResponse[RemovedPlanFeature](r.http.delete(ctx, fmt.Sprintf("/plans/%s/features/%s", id, featureID), nil, ""))
}

// Attach a feature to a plan with limits, overage, and credits configuration.
func (r *PlansResource) AddFeature(ctx context.Context, id string, params *AddPlanFeatureParams) (*PlanFeature, error) {
	body := buildBody(map[string]any{
		"feature_id":       params.FeatureID,
		"enabled":          params.Enabled,
		"included_amount":  params.IncludedAmount,
		"unlimited":        params.Unlimited,
		"overage":          params.Overage,
		"credits_per_unit": params.CreditsPerUnit,
		"pricing_mode":     params.PricingMode,
		"margin":           params.Margin,
	})
	return parseDirectResponse[PlanFeature](r.http.post(ctx, fmt.Sprintf("/plans/%s/features", id), body, params.IdempotencyKey))
}

// Cancel a move for every subscriber who has not moved yet. They stay on their version; those already moved stay moved.
func (r *PlansResource) CancelMove(ctx context.Context, id string, moveID string, params *CancelPlanMoveParams) (*PlanMove, error) {
	return parseDirectResponse[PlanMove](r.http.post(ctx, fmt.Sprintf("/plans/%s/moves/%s/cancel", id, moveID), map[string]any{}, params.IdempotencyKey))
}

// Read a move of subscribers between plan versions with each subscriber's status: moving, moved, needs a change, left out, canceled or kept on its version.
func (r *PlansResource) GetMove(ctx context.Context, id string, moveID string) (*PlanMove, error) {
	return parseDirectResponse[PlanMove](r.http.get(ctx, fmt.Sprintf("/plans/%s/moves/%s", id, moveID), nil))
}

// Cancel a move for one subscriber who has not moved yet. It stays on its version; the rest of the move continues.
func (r *PlansResource) CancelMoveSubscription(ctx context.Context, id string, moveID string, subscriptionID string, params *CancelPlanMoveSubscriptionParams) (*PlanMove, error) {
	return parseDirectResponse[PlanMove](r.http.post(ctx, fmt.Sprintf("/plans/%s/moves/%s/subscriptions/%s/cancel", id, moveID, subscriptionID), map[string]any{}, params.IdempotencyKey))
}

// List a plan's latest moves of subscribers between versions, newest first, including the change of a single subscription, with how many are moving, moved, need a change, were left out, canceled or kept on their version.
func (r *PlansResource) ListMoves(ctx context.Context, id string) (*PlansListMovesResult, error) {
	return parseDirectResponse[PlansListMovesResult](r.http.get(ctx, fmt.Sprintf("/plans/%s/moves", id), nil))
}

// Move a list of subscriptions, or every subscription of another version, to a newer version of the plan, at their next renewal or now. Nobody moves unless every one of them fits or has a change in changes; otherwise the response lists each one with what would make it fit. Only active, trialing and past due subscriptions move. Paused ones, and listed subscriptions in another status or already on the destination, are left out. Send preview to check without moving anyone.
func (r *PlansResource) CreateMove(ctx context.Context, id string, params *CreatePlanMoveParams) (*PlanMove, error) {
	body := buildBody(map[string]any{
		"subscription_ids": params.SubscriptionIds,
		"to_version_id":    params.ToVersionID,
		"effective":        params.Effective,
		"changes":          params.Changes,
		"preview":          params.Preview,
		"from_version_id":  params.FromVersionID,
	})
	return parseDirectResponse[PlanMove](r.http.post(ctx, fmt.Sprintf("/plans/%s/moves", id), body, params.IdempotencyKey))
}

// Leave a subscriber that needs a change on its version. It stops appearing among the subscribers that need a change; move it later like any other subscriber.
func (r *PlansResource) KeepSubscriberOnVersion(ctx context.Context, id string, subscriptionID string, params *KeepPlanSubscriberOnVersionParams) (*KeptOnVersion, error) {
	return parseDirectResponse[KeptOnVersion](r.http.post(ctx, fmt.Sprintf("/plans/%s/needs-a-change/%s/keep", id, subscriptionID), map[string]any{}, params.IdempotencyKey))
}

// List the plan's subscribers whose move could not complete, with the version they stay on, the one they could not move to and why. Move them again with a change, or keep them on their version so they leave this list.
func (r *PlansResource) ListSubscribersNeedingAChange(ctx context.Context, id string) (*PlansListSubscribersNeedingAChangeResult, error) {
	return parseDirectResponse[PlansListSubscribersNeedingAChangeResult](r.http.get(ctx, fmt.Sprintf("/plans/%s/needs-a-change", id), nil))
}

// Set a specific price as the default and return the updated plan price.
func (r *PlansResource) SetDefaultPrice(ctx context.Context, id string, priceID string, params *SetDefaultPlanPriceParams) (*PlanPrice, error) {
	return parseDirectResponse[PlanPrice](r.http.put(ctx, fmt.Sprintf("/plans/%s/prices/%s/default", id, priceID), map[string]any{}, params.IdempotencyKey))
}

// Create or update regional currency price overrides for a plan price.
func (r *PlansResource) SetRegionalPrices(ctx context.Context, id string, priceID string, params *UpsertRegionalPricesParams) (*PlanRegionalPricing, error) {
	body := buildBody(map[string]any{
		"overrides": params.Overrides,
	})
	return parseDirectResponse[PlanRegionalPricing](r.http.put(ctx, fmt.Sprintf("/plans/%s/prices/%s/regional", id, priceID), body, params.IdempotencyKey))
}

// Remove all regional currency overrides for a plan price. The request is rejected while billable subscriptions depend on an override.
func (r *PlansResource) DeleteRegionalPrices(ctx context.Context, id string, priceID string) (*DeletedPlanRegionalPricing, error) {
	return parseDirectResponse[DeletedPlanRegionalPricing](r.http.delete(ctx, fmt.Sprintf("/plans/%s/prices/%s/regional", id, priceID), nil, ""))
}

// Update a base price or market price variant. Removing a base market override is rejected while a variant depends on it. Offer terms are managed through Offers.
func (r *PlansResource) UpdatePrice(ctx context.Context, id string, priceID string, params *UpdatePlanPriceParams) (*PlanPrice, error) {
	body := buildBody(map[string]any{
		"price":                         params.Price,
		"is_default":                    params.IsDefault,
		"trial_days":                    params.TrialDays,
		"trial_requires_payment_method": params.TrialRequiresPaymentMethod,
		"included_balance":              params.IncludedBalance,
		"included_credits":              params.IncludedCredits,
		"metadata":                      params.Metadata,
		"market_prices":                 params.MarketPrices,
	})
	return parseDirectResponse[PlanPrice](r.http.patch(ctx, fmt.Sprintf("/plans/%s/prices/%s", id, priceID), body, params.IdempotencyKey))
}

// Archive a price for new subscriptions. This publishes a new version of the plan; existing subscriptions that selected the price keep paying it on their version until moved.
func (r *PlansResource) DeletePrice(ctx context.Context, id string, priceID string) (*DeletedObject, error) {
	return parseDirectResponse[DeletedObject](r.http.delete(ctx, fmt.Sprintf("/plans/%s/prices/%s", id, priceID), nil, ""))
}

// Add a base price or a selectable market price variant. Variants inherit their base price outside the markets they override. Configure introductory and promotional benefits through Offers.
func (r *PlansResource) AddPrice(ctx context.Context, id string, params *AddPlanPriceParams) (*PlanPrice, error) {
	body := buildBody(map[string]any{
		"billing_interval":              params.BillingInterval,
		"metadata":                      params.Metadata,
		"price":                         params.Price,
		"trial_days":                    params.TrialDays,
		"trial_requires_payment_method": params.TrialRequiresPaymentMethod,
		"is_default":                    params.IsDefault,
		"included_balance":              params.IncludedBalance,
		"included_credits":              params.IncludedCredits,
		"market_prices":                 params.MarketPrices,
		"inherits_from_price_id":        params.InheritsFromPriceID,
	})
	return parseDirectResponse[PlanPrice](r.http.post(ctx, fmt.Sprintf("/plans/%s/prices", id), body, params.IdempotencyKey))
}

// Configure regional prices and feature overage values for one currency. Currency-specific offer terms are managed through Offers.
func (r *PlansResource) SetRegionalPricing(ctx context.Context, id string, params *SetPlanRegionalPricingParams) (*PlanRegionalPricingResult, error) {
	body := buildBody(map[string]any{
		"currency":      params.Currency,
		"exchange_rate": params.ExchangeRate,
		"prices":        params.Prices,
		"features":      params.Features,
	})
	return parseDirectResponse[PlanRegionalPricingResult](r.http.put(ctx, fmt.Sprintf("/plans/%s/regional", id), body, params.IdempotencyKey))
}

// Get a plan with public price IDs and their automatic introductory offer IDs.
func (r *PlansResource) Get(ctx context.Context, id string) (*Plan, error) {
	return parseDirectResponse[Plan](r.http.get(ctx, fmt.Sprintf("/plans/%s", id), nil))
}

// Update a plan's name, description, visibility, default status, metadata, overage behavior, what happens to its unused plan credits or plan balance, or its plan credit or plan balance limit.
func (r *PlansResource) Update(ctx context.Context, id string, params *UpdatePlanParams) (*Plan, error) {
	body := buildBody(map[string]any{
		"name":                params.Name,
		"description":         params.Description,
		"metadata":            params.Metadata,
		"is_public":           params.IsPublic,
		"unused_plan_credits": params.UnusedPlanCredits,
		"plan_credit_limit":   params.PlanCreditLimit,
		"unused_plan_balance": params.UnusedPlanBalance,
		"plan_balance_limit":  params.PlanBalanceLimit,
		"block_on_exhaustion": params.BlockOnExhaustion,
		"is_default":          params.IsDefault,
	})
	return parseDirectResponse[Plan](r.http.patch(ctx, fmt.Sprintf("/plans/%s", id), body, params.IdempotencyKey))
}

// Soft-delete a plan.
func (r *PlansResource) Delete(ctx context.Context, id string) (*DeletedObject, error) {
	return parseDirectResponse[DeletedObject](r.http.delete(ctx, fmt.Sprintf("/plans/%s", id), nil, ""))
}

// List what changes from one version of a plan to another: prices, trials, features, rules, add-ons, credit packs and countries.
func (r *PlansResource) CompareVersions(ctx context.Context, id string, versionID string, params *ComparePlanVersionsParams) (*PlanVersionComparison, error) {
	query := map[string]string{}
	if params.Against != "" {
		query["against"] = params.Against
	}
	return parseDirectResponse[PlanVersionComparison](r.http.get(ctx, fmt.Sprintf("/plans/%s/versions/%s/changes", id, versionID), query))
}

// Get one published version of a plan with its prices and how many subscriptions are on it.
func (r *PlansResource) GetVersion(ctx context.Context, id string, versionID string) (*PlanVersion, error) {
	return parseDirectResponse[PlanVersion](r.http.get(ctx, fmt.Sprintf("/plans/%s/versions/%s", id, versionID), nil))
}

// List a plan's published versions, newest first, with their prices and how many subscriptions are on each. Versions published before Commet recorded every term are not listed.
func (r *PlansResource) ListVersions(ctx context.Context, id string, params *ListPlanVersionsParams) (*PlansListVersionsResult, error) {
	query := map[string]string{}
	if params.Cursor != nil {
		query["cursor"] = *params.Cursor
	}
	if params.Limit != nil {
		query["limit"] = fmt.Sprintf("%d", *params.Limit)
	}
	return parseDirectResponse[PlansListVersionsResult](r.http.get(ctx, fmt.Sprintf("/plans/%s/versions", id), query))
}

// Publish a new current version of the plan from its current version, changing prices per interval, included credits or balance, per-currency amounts and the rollover cap. New buyers get it; every subscriber stays on its version until moved. To try changes without selling them to everyone, create an experiment instead.
func (r *PlansResource) CreateVersion(ctx context.Context, id string, params *CreatePlanVersionParams) (*PlanVersion, error) {
	body := buildBody(map[string]any{
		"prices":           params.Prices,
		"currencies":       params.Currencies,
		"max_plan_credits": params.MaxPlanCredits,
		"max_plan_balance": params.MaxPlanBalance,
	})
	return parseDirectResponse[PlanVersion](r.http.post(ctx, fmt.Sprintf("/plans/%s/versions", id), body, params.IdempotencyKey))
}

// Set a plan's public visibility and return the updated plan.
func (r *PlansResource) SetVisibility(ctx context.Context, id string, params *SetPlanVisibilityParams) (*Plan, error) {
	body := buildBody(map[string]any{
		"is_public": params.IsPublic,
	})
	return parseDirectResponse[Plan](r.http.put(ctx, fmt.Sprintf("/plans/%s/visibility", id), body, params.IdempotencyKey))
}

// List plans with public price IDs and their automatic introductory offer IDs.
func (r *PlansResource) List(ctx context.Context, params *ListPlansParams) (*PlansListResult, error) {
	query := map[string]string{}
	if params.IncludePrivate != nil {
		query["include_private"] = fmt.Sprintf("%t", *params.IncludePrivate)
	}
	return parseDirectResponse[PlansListResult](r.http.get(ctx, "/plans", query))
}

// Create a new plan with optional consumption model, visibility, plan group assignment, overage behavior, and the rollover and spend limit of its plan credits or plan balance.
func (r *PlansResource) Create(ctx context.Context, params *CreatePlanParams) (*Plan, error) {
	body := buildBody(map[string]any{
		"name":                params.Name,
		"code":                params.Code,
		"description":         params.Description,
		"consumption_model":   params.ConsumptionModel,
		"is_public":           params.IsPublic,
		"is_free":             params.IsFree,
		"block_on_exhaustion": params.BlockOnExhaustion,
		"unused_plan_credits": params.UnusedPlanCredits,
		"plan_credit_limit":   params.PlanCreditLimit,
		"unused_plan_balance": params.UnusedPlanBalance,
		"plan_balance_limit":  params.PlanBalanceLimit,
		"plan_group_id":       params.PlanGroupID,
		"is_default":          params.IsDefault,
		"metadata":            params.Metadata,
	})
	return parseDirectResponse[Plan](r.http.post(ctx, "/plans", body, params.IdempotencyKey))
}
