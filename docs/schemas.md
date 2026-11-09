# Schemas

Generated from Commet API version `2026-08-27`.

## Enums

### BillingInterval

- `"weekly"`
- `"monthly"`
- `"quarterly"`
- `"semiannual"`
- `"yearly"`
- `"one_time"`

### ConsumptionModel

- `"metered"`
- `"credits"`
- `"balance"`

### FeatureType

- `"boolean"`
- `"usage"`
- `"seats"`
- `"quota"`

### InvoiceType

- `"recurring"`
- `"overage"`
- `"plan_change"`
- `"adjustment"`
- `"credit_purchase"`
- `"balance_topup"`
- `"addon_activation"`
- `"one_time_payment"`
- `"reactivation"`
- `"resume"`
- `"prepayment"`

### PaymentMethod

- `"card"`
- `"oxxo"`
- `"mercado_pago"`
- `"google_pay"`
- `"smartpix"`

### PaymentProvider

- `"stripe"`
- `"commet"`
- `"dlocal"`

### SubPaymentMethod

- `"credit_card"`
- `"debit_card"`
- `"prepaid_card"`
- `"bank_transfer"`
- `"account_money"`

### SubscriptionStatus

- `"draft"`
- `"pending_payment"`
- `"trialing"`
- `"active"`
- `"past_due"`
- `"paused"`
- `"canceled"`

### Timezone

- `"UTC"`
- `"America/New_York"`
- `"America/Chicago"`
- `"America/Denver"`
- `"America/Los_Angeles"`
- `"America/Sao_Paulo"`
- `"America/Mexico_City"`
- `"America/Buenos_Aires"`
- `"America/Santiago"`
- `"America/Bogota"`
- `"America/Lima"`
- `"America/Asuncion"`
- `"Europe/London"`
- `"Europe/Paris"`
- `"Europe/Berlin"`
- `"Europe/Madrid"`
- `"Asia/Tokyo"`
- `"Asia/Shanghai"`
- `"Asia/Singapore"`
- `"Asia/Dubai"`
- `"Australia/Sydney"`

### TransactionStatus

- `"pending"`
- `"succeeded"`
- `"failed"`
- `"refunded"`
- `"disputed"`

## Models

### ActiveAddon

- `Slug` (`string`, required)
- `Name` (`string`, required)
- `BasePrice` (`int`, required)
- `FeatureCode` (`string`, required)
- `FeatureName` (`string`, required)
- `FeatureType` (`FeatureType`, required)
- `ConsumptionModel` (`string`, required)
- `ActivatedAt` (`string`, required)
- `Object` (`string`, required)
- `Livemode` (`bool`, required)

### AddedPlanToGroup

- `Success` (`bool`, required)
- `Object` (`string`, required)
- `Livemode` (`bool`, required)

### Addon

- `ID` (`string`, required)
- `Name` (`string`, required)
- `Slug` (`string`, required)
- `Description` (`string | null`, required)
- `BasePrice` (`int`, required)
- `FeatureCode` (`string`, required)
- `FeatureName` (`string`, required)
- `CreatedAt` (`string`, required)
- `UpdatedAt` (`string`, required)
- `ConsumptionModel` (`string`, required)
- `IncludedUnits` (`int | null`, required)
- `OverageRate` (`int | null`, required)
- `CreditCost` (`int | null`, required)
- `PlanIds` (`[]string`, required) — Plans that sell the add-on.
- `Object` (`string`, required)
- `Livemode` (`bool`, required)

### AddonsListActiveResult

- `Object` (`string`, required)
- `Data` (`[]ActiveAddon`, required)
- `HasMore` (`bool`, required)
- `NextCursor` (`string`, optional)

### AddonsListResult

- `Object` (`string`, required)
- `Data` (`[]Addon`, required)
- `HasMore` (`bool`, required)
- `NextCursor` (`string`, optional)

### AddPlanFeatureParamsOverage

- `Enabled` (`bool`, optional)
- `UnitPrice` (`int`, optional)

### AddPlanPriceParamsMarketPricesItem

- `MarketGroupID` (`string`, required) — Public ID of a reusable pricing market group.
- `Currency` (`string`, required) — Presentment currency configured for this plan and market.
- `Price` (`int`, required) — Market price in the currency's minor unit.

### ApiKey

- `ID` (`string`, required)
- `Name` (`string`, required)
- `Prefix` (`string`, required)
- `ExpiresAt` (`string | null`, required)
- `LastUsedAt` (`string | null`, required)
- `CreatedAt` (`string`, required)
- `Object` (`string`, required)
- `Livemode` (`bool`, required)

### ApiKeysListResult

- `Object` (`string`, required)
- `Data` (`[]ApiKey`, required)
- `HasMore` (`bool`, required)
- `NextCursor` (`string`, optional)

### ArchivePlanExperimentParamsChangesItem

- `SubscriptionID` (`string`, required)
- `KeepOnVersion` (`any`, optional) — Leave this subscriber on its version.
- `BillingInterval` (`string`, optional)
- `SeatLimit` (`any`, optional) — Lower its seats to the version's limit at renewal.
- `RemoveAddonIds` (`[]string`, optional) — Add-ons it stops paying at renewal.

### BalanceAdjustment

- `Amount` (`int`, required)
- `NewBalance` (`int`, required)
- `Reason` (`string | null`, required)
- `Object` (`string`, required)
- `Livemode` (`bool`, required)

### BalanceTopup

- `Amount` (`int`, required)
- `Object` (`string`, required)
- `Livemode` (`bool`, required)

### BatchCreateCustomersParamsCustomersItem

- `Email` (`string`, required)
- `ID` (`string`, optional)
- `ExternalID` (`string`, optional)
- `FullName` (`string`, optional)
- `TaxDocument` (`string`, optional)
- `Timezone` (`Timezone`, optional)
- `Metadata` (`map[string]any`, optional)
- `Address` (`BatchCreateCustomersParamsCustomersItemAddress`, optional)

### BatchCreateCustomersParamsCustomersItemAddress

- `Line1` (`string`, required)
- `Line2` (`string`, optional)
- `City` (`string`, required)
- `State` (`string`, optional)
- `PostalCode` (`string`, required)
- `Country` (`string`, required)
- `Region` (`string`, optional)

### BusinessTaxId

- `Country` (`string`, required)
- `Type` (`string`, required)
- `Value` (`string`, required)

### ClaimLink

- `URL` (`string`, required)
- `ExpiresAt` (`string`, required)
- `Object` (`string`, required)
- `Livemode` (`bool`, required)

### CreateApiKeyParamsPermissions

- `Customer` (`[]string`, optional)
- `Subscription` (`[]string`, optional)
- `Invoice` (`[]string`, optional)
- `Usage` (`[]string`, optional)
- `Seat` (`[]string`, optional)
- `Plan` (`[]string`, optional)
- `PlanGroup` (`[]string`, optional)
- `Feature` (`[]string`, optional)
- `Addon` (`[]string`, optional)
- `CreditPack` (`[]string`, optional)
- `Offer` (`[]string`, optional)
- `PromoCode` (`[]string`, optional)
- `MarketGroup` (`[]string`, optional)
- `Payment` (`[]string`, optional)
- `Transaction` (`[]string`, optional)
- `Payout` (`[]string`, optional)
- `TestClock` (`[]string`, optional)
- `Organization` (`[]string`, optional)
- `APIKey` (`[]string`, optional)

### CreateCustomerParamsAddress

- `Line1` (`string`, required)
- `Line2` (`string`, optional)
- `City` (`string`, required)
- `State` (`string`, optional)
- `PostalCode` (`string`, required)
- `Country` (`string`, required)
- `Region` (`string`, optional)

### CreatedApiKey

- `ID` (`string`, required)
- `Name` (`string`, required)
- `APIKey` (`string`, required)
- `Prefix` (`string`, required)
- `ExpiresAt` (`string`, required)
- `CreatedAt` (`string`, required)
- `Object` (`string`, required)
- `Livemode` (`bool`, required)

### CreatedSubscription

- `ID` (`string`, required)
- `CustomerID` (`string`, required)
- `Plan` (`CreatedSubscriptionPlan`, required)
- `Name` (`string`, required)
- `Description` (`string | null`, required)
- `Status` (`SubscriptionStatus`, required)
- `BillingInterval` (`BillingInterval | null`, required)
- `TrialEndsAt` (`string | null`, required)
- `CurrentPeriod` (`CreatedSubscriptionCurrentPeriod | null`, required)
- `Cancellation` (`CreatedSubscriptionCancellation | null`, required)
- `CancelAtPeriodEnd` (`bool`, required)
- `ScheduledPlanChange` (`CreatedSubscriptionScheduledPlanChange | null`, required)
- `StartDate` (`string`, required)
- `EndDate` (`string | null`, required)
- `BillingDayOfMonth` (`int | null`, required)
- `NextBillingDate` (`string | null`, required)
- `CheckoutURL` (`string | null`, required)
- `HasPaymentMethod` (`bool`, required) — Whether Commet can charge this subscription on its own at its next charge (the trial end or the renewal). A trial that started without a payment method reports false until the customer adds one for this subscription; a payment method saved for another subscription does not count. A subscription on a free plan reports false.
- `CreatedAt` (`string`, required)
- `UpdatedAt` (`string`, required)
- `PlanVersion` (`CreatedSubscriptionPlanVersion | null`, required) — The plan version whose terms the subscription has now. It keeps paying and receiving everything from this version until it is moved.
- `OfferApplications` (`[]SubscriptionOfferApplication`, required)
- `Pause` (`CreatedSubscriptionPause | null`, required)
- `CheckoutProvider` (`PaymentProvider | null`, required) — Payment provider resolved for this checkout when the subscription response was created. This is an informational snapshot and may differ when the checkout is loaded if its country or the organization's routing changes.
- `PriceID` (`string | null`, required)
- `Object` (`string`, required)
- `Livemode` (`bool`, required)

### CreatedSubscriptionCancellation

- `ScheduledAt` (`string`, required)
- `Reason` (`string | null`, required)
- `EffectiveAt` (`string`, required)

### CreatedSubscriptionCurrentPeriod

- `Start` (`string`, required)
- `End` (`string`, required)
- `DaysRemaining` (`float64`, required)

### CreatedSubscriptionPause

Variants:

- `CreatedSubscriptionPauseVariant1`
- `CreatedSubscriptionPauseVariant2`

Discriminator: `Status`

- `"scheduled"` → `CreatedSubscriptionPauseVariant1`
- `"active"` → `CreatedSubscriptionPauseVariant2`

### CreatedSubscriptionPauseVariant1

- `Status` (`string`, required)
- `Mode` (`string`, required)
- `RequestedAt` (`string`, required)
- `EffectiveAt` (`string`, required)
- `ResumeAt` (`string | null`, required)

### CreatedSubscriptionPauseVariant2

- `Status` (`string`, required)
- `Mode` (`string`, required)
- `RequestedAt` (`string`, required)
- `EffectiveAt` (`string`, required)
- `ResumeAt` (`string | null`, required)

### CreatedSubscriptionPlan

- `ID` (`string`, required)
- `Name` (`string`, required)

### CreatedSubscriptionPlanVersion

- `ID` (`string`, required) — Public plan version ID.
- `Label` (`string`, required) — Main versions are 1, 2, 3; an experiment's versions carry its own code, like price-test 2.

### CreatedSubscriptionScheduledPlanChange

- `ChangeType` (`string`, required)
- `NewPlanID` (`string | null`, required)
- `NewPlanName` (`string | null`, required)
- `NewBillingInterval` (`string | null`, required)
- `ScheduledFor` (`string`, required)

### CreatedWebhook

- `ID` (`string`, required)
- `URL` (`string`, required)
- `Events` (`[]string`, required)
- `Description` (`string | null`, required)
- `IsActive` (`bool`, required)
- `APIVersion` (`string | null`, required)
- `CreatedAt` (`string`, required)
- `SecretKey` (`string`, required)
- `Object` (`string`, required)
- `Livemode` (`bool`, required)

### CreateOfferParamsPhasesItem

Variants:

- `CreateOfferParamsPhasesItemVariant1`
- `CreateOfferParamsPhasesItemVariant2`
- `CreateOfferParamsPhasesItemVariant3`

Discriminator: `Type`

- `"percentage"` → `CreateOfferParamsPhasesItemVariant1`
- `"amount_off"` → `CreateOfferParamsPhasesItemVariant2`
- `"fixed_price"` → `CreateOfferParamsPhasesItemVariant3`

### CreateOfferParamsPhasesItemVariant1

- `Type` (`string`, required)
- `DurationCycles` (`int | null`, required)
- `DurationInterval` (`string | null`, optional) — Unit the phase duration is counted in. Only a fixed-price phase may set it, because its amount is declared rather than derived from the plan. Defaults to the plan's own billing interval.
- `Percentage` (`int`, required) — Discount in basis points. 5000 means 50%.

### CreateOfferParamsPhasesItemVariant2

- `Type` (`string`, required)
- `DurationCycles` (`int | null`, required)
- `DurationInterval` (`string | null`, optional) — Unit the phase duration is counted in. Only a fixed-price phase may set it, because its amount is declared rather than derived from the plan. Defaults to the plan's own billing interval.
- `Amounts` (`[]CreateOfferParamsPhasesItemVariant2AmountsItem`, required)

### CreateOfferParamsPhasesItemVariant2AmountsItem

- `Currency` (`string`, required)
- `Amount` (`int`, required) — Amount in the currency's minor unit (for example, cents for USD).

### CreateOfferParamsPhasesItemVariant3

- `Type` (`string`, required)
- `DurationCycles` (`int | null`, required)
- `DurationInterval` (`string | null`, optional) — Unit the phase duration is counted in. Only a fixed-price phase may set it, because its amount is declared rather than derived from the plan. Defaults to the plan's own billing interval.
- `Prices` (`[]CreateOfferParamsPhasesItemVariant3PricesItem`, required)

### CreateOfferParamsPhasesItemVariant3PricesItem

- `Currency` (`string`, required)
- `Amount` (`int`, required) — Amount in the currency's minor unit (for example, cents for USD).

### CreatePlanExperimentParamsCurrenciesItem

- `Currency` (`string`, required)
- `ExchangeRate` (`float64`, required)

### CreatePlanExperimentParamsPricesItem

- `BillingInterval` (`string`, required) — An interval the plan already sells. Prices not listed are copied unchanged from the current version.
- `Price` (`int`, required) — USD price in the currency's minor unit (cents).
- `IncludedCredits` (`int | null`, optional)
- `IncludedBalance` (`int | null`, optional)
- `RegionalPrices` (`[]CreatePlanExperimentParamsPricesItemRegionalPricesItem`, optional) — Per-currency amounts. A currency the current version does not sell also needs its exchangeRate in currencies; its amounts not listed here are derived from the USD price.

### CreatePlanExperimentParamsPricesItemRegionalPricesItem

- `Currency` (`string`, required)
- `Price` (`int`, required)
- `IncludedBalance` (`int | null`, optional)

### CreatePlanExperimentVersionParamsCurrenciesItem

- `Currency` (`string`, required)
- `ExchangeRate` (`float64`, required)

### CreatePlanExperimentVersionParamsPricesItem

- `BillingInterval` (`string`, required) — An interval the plan already sells. Prices not listed are copied unchanged from the current version.
- `Price` (`int`, required) — USD price in the currency's minor unit (cents).
- `IncludedCredits` (`int | null`, optional)
- `IncludedBalance` (`int | null`, optional)
- `RegionalPrices` (`[]CreatePlanExperimentVersionParamsPricesItemRegionalPricesItem`, optional) — Per-currency amounts. A currency the current version does not sell also needs its exchangeRate in currencies; its amounts not listed here are derived from the USD price.

### CreatePlanExperimentVersionParamsPricesItemRegionalPricesItem

- `Currency` (`string`, required)
- `Price` (`int`, required)
- `IncludedBalance` (`int | null`, optional)

### CreatePlanMoveParamsChangesItem

- `SubscriptionID` (`string`, required)
- `KeepOnVersion` (`any`, optional) — Leave this subscriber on its version.
- `BillingInterval` (`string`, optional)
- `SeatLimit` (`any`, optional) — Lower its seats to the version's limit at renewal.
- `RemoveAddonIds` (`[]string`, optional) — Add-ons it stops paying at renewal.

### CreatePlanParamsPlanBalanceLimit

- `Amount` (`int`, required) — The most plan balance a subscription spends in one window, in USD cents. Other currencies convert it with the plan's exchange rate unless the dashboard sets a custom amount for that currency.
- `Interval` (`string`, required) — The window length: `day` is 24 hours and `week` is 7 days, starting at the time of day of the subscription's billing anchor.

### CreatePlanParamsPlanCreditLimit

- `Credits` (`int`, required) — The most plan credits a subscription spends in one window.
- `Interval` (`string`, required) — The window length: `day` is 24 hours and `week` is 7 days, starting at the time of day of the subscription's billing anchor.

### CreatePlanParamsUnusedPlanBalance

Variants:

- `CreatePlanParamsUnusedPlanBalanceVariant1`
- `CreatePlanParamsUnusedPlanBalanceVariant2`

Discriminator: `Type`

- `"expire"` → `CreatePlanParamsUnusedPlanBalanceVariant1`
- `"roll_over"` → `CreatePlanParamsUnusedPlanBalanceVariant2`

### CreatePlanParamsUnusedPlanBalanceVariant1

- `Type` (`string`, required)

### CreatePlanParamsUnusedPlanBalanceVariant2

- `Type` (`string`, required)
- `MaxPlanBalance` (`int | null`, required) — The most plan balance a subscription holds after a reset, in USD cents, including the balance that reset grants; balance above it expires. At least the included balance of every price. Other currencies convert it with the plan's exchange rate unless the dashboard sets a custom amount for that currency. null means no limit other than the largest storable amount, 2,147,483,647 cents.

### CreatePlanParamsUnusedPlanCredits

Variants:

- `CreatePlanParamsUnusedPlanCreditsVariant1`
- `CreatePlanParamsUnusedPlanCreditsVariant2`

Discriminator: `Type`

- `"expire"` → `CreatePlanParamsUnusedPlanCreditsVariant1`
- `"roll_over"` → `CreatePlanParamsUnusedPlanCreditsVariant2`

### CreatePlanParamsUnusedPlanCreditsVariant1

- `Type` (`string`, required)

### CreatePlanParamsUnusedPlanCreditsVariant2

- `Type` (`string`, required)
- `MaxPlanCredits` (`int | null`, required) — The most plan credits a subscription holds after a reset, including the credits that reset grants; credits above it expire. At least the included credits of every price. null means no limit other than the largest storable balance, 2,147,483,647 plan credits.

### CreatePlanVersionParamsCurrenciesItem

- `Currency` (`string`, required)
- `ExchangeRate` (`float64`, required)

### CreatePlanVersionParamsPricesItem

- `BillingInterval` (`string`, required) — An interval the plan already sells. Prices not listed are copied unchanged from the current version.
- `Price` (`int`, required) — USD price in the currency's minor unit (cents).
- `IncludedCredits` (`int | null`, optional)
- `IncludedBalance` (`int | null`, optional)
- `RegionalPrices` (`[]CreatePlanVersionParamsPricesItemRegionalPricesItem`, optional) — Per-currency amounts. A currency the current version does not sell also needs its exchangeRate in currencies; its amounts not listed here are derived from the USD price.

### CreatePlanVersionParamsPricesItemRegionalPricesItem

- `Currency` (`string`, required)
- `Price` (`int`, required)
- `IncludedBalance` (`int | null`, optional)

### CreditGrant

- `Credits` (`int`, required)
- `Object` (`string`, required)
- `Livemode` (`bool`, required)

### CreditPack

- `ID` (`string`, required)
- `Name` (`string`, required)
- `Description` (`string | null`, required)
- `Credits` (`int`, required)
- `Price` (`int`, required)
- `PlanIds` (`[]string`, required) — Plans that sell the credit pack.
- `IsActive` (`bool`, required)
- `CreatedAt` (`string`, required)
- `UpdatedAt` (`string`, required)
- `Object` (`string`, required)
- `Livemode` (`bool`, required)

### CreditPackListItem

- `ID` (`string`, required)
- `Name` (`string`, required)
- `Description` (`string | null`, required)
- `Credits` (`int`, required)
- `Price` (`int`, required)
- `PlanIds` (`[]string`, required) — Plans that sell the credit pack.
- `Currency` (`string`, required)
- `Object` (`string`, required)
- `Livemode` (`bool`, required)

### CreditPacksListResult

- `Object` (`string`, required)
- `Data` (`[]CreditPackListItem`, required)
- `HasMore` (`bool`, required)
- `NextCursor` (`string`, optional)

### Customer

- `ID` (`string`, required)
- `ExternalID` (`string | null`, required)
- `FullName` (`string | null`, required)
- `Email` (`string`, required)
- `TaxDocument` (`string | null`, required)
- `DocumentType` (`string | null`, required)
- `BusinessTaxID` (`BusinessTaxId | null`, required)
- `Timezone` (`string | null`, required)
- `Metadata` (`map[string]any | null`, required)
- `CreatedAt` (`string`, required)
- `UpdatedAt` (`string`, required)
- `Object` (`string`, required)
- `Livemode` (`bool`, required)

### CustomerBatch

- `Successful` (`[]CustomerBatchSuccessfulItem`, required)
- `Failed` (`[]CustomerBatchFailedItem`, required)
- `Object` (`string`, required)
- `Livemode` (`bool`, required)

### CustomerBatchFailedItem

- `Index` (`int`, required)
- `Error` (`string`, required)
- `Data` (`CustomerBatchFailedItemData`, required)

### CustomerBatchFailedItemData

- `ID` (`string`, optional)
- `ExternalID` (`string`, optional)
- `Email` (`string`, required)
- `FullName` (`string | null`, optional)
- `TaxDocument` (`string | null`, optional)
- `Timezone` (`string`, optional)
- `Metadata` (`map[string]any | null`, optional)
- `Address` (`CustomerBatchFailedItemDataAddress`, optional)

### CustomerBatchFailedItemDataAddress

- `Line1` (`string`, required)
- `Line2` (`string`, optional)
- `City` (`string`, required)
- `State` (`string`, optional)
- `PostalCode` (`string`, required)
- `Country` (`string`, required)
- `Region` (`string`, optional)

### CustomerBatchSuccessfulItem

- `ID` (`string`, required)
- `ExternalID` (`string | null`, required)
- `Email` (`string`, required)

### CustomerCredit

- `ID` (`string`, required)
- `Amount` (`int`, required) — Original grant amount in the currency's smallest unit.
- `AppliedAmount` (`int`, required)
- `ReversedAmount` (`int`, required)
- `RevokedAmount` (`int`, required)
- `RemainingAmount` (`int`, required)
- `Currency` (`string`, required)
- `Reason` (`string`, required)
- `Source` (`string`, required)
- `ExpiresAt` (`string | null`, required)
- `CreatedAt` (`string`, required)
- `Object` (`string`, required)
- `Livemode` (`bool`, required)

### CustomerCreditRevocation

- `ID` (`string`, required)
- `RemainingAmount` (`int`, required)
- `RevokedAmount` (`int`, required)
- `Currency` (`string`, required)
- `Object` (`string`, required)
- `Livemode` (`bool`, required)

### CustomersListCreditsResult

- `Object` (`string`, required)
- `Data` (`[]CustomerCredit`, required)
- `HasMore` (`bool`, required)
- `NextCursor` (`string`, optional)

### CustomersListPlanGrantsResult

- `Object` (`string`, required)
- `Data` (`[]PlanGrant`, required)
- `HasMore` (`bool`, required)
- `NextCursor` (`string`, optional)

### CustomersListResult

- `Object` (`string`, required)
- `Data` (`[]Customer`, required)
- `HasMore` (`bool`, required)
- `NextCursor` (`string`, optional)

### DeletedObject

- `ID` (`string`, required)
- `Deleted` (`any`, required)
- `Object` (`string`, required)
- `Livemode` (`bool`, required)

### DeletedOffer

- `Deleted` (`any`, required)
- `Object` (`string`, required)
- `Livemode` (`bool`, required)

### DeletedPlanRegionalPricing

- `Deleted` (`any`, required)
- `Object` (`string`, required)
- `Livemode` (`bool`, required)

### DeletedSubscriptionAddon

- `ID` (`string`, required)
- `Status` (`string`, required)
- `DeactivatedAt` (`string | null`, required)
- `Object` (`string`, required)
- `Livemode` (`bool`, required)

### Feature

- `ID` (`string`, required)
- `Name` (`string`, required)
- `Code` (`string`, required)
- `Type` (`FeatureType`, required)
- `Description` (`string | null`, required)
- `UnitName` (`string | null`, required)
- `CreatedAt` (`string`, required)
- `UpdatedAt` (`string`, required)
- `Object` (`string`, required)
- `Livemode` (`bool`, required)

### FeatureAccess

Variants:

- `FeatureAccessVariant1`
- `FeatureAccessVariant2`
- `FeatureAccessVariant3`
- `FeatureAccessVariant4`

Discriminator: `Type`

- `"boolean"` → `FeatureAccessVariant1`
- `"usage"` → `FeatureAccessVariant2`
- `"seats"` → `FeatureAccessVariant3`
- `"quota"` → `FeatureAccessVariant4`

### FeatureAccessListResult

- `Object` (`string`, required)
- `Data` (`[]FeatureAccess`, required)
- `HasMore` (`bool`, required)
- `NextCursor` (`string`, optional)

### FeatureAccessVariant1

- `Code` (`string`, required) — Unique feature code.
- `Name` (`string`, required) — Display name of the feature.
- `UnitName` (`string | null`, required) — Display name for one product unit, or null when not applicable.
- `Allowed` (`bool`, required) — Whether the customer can currently access or consume the feature.
- `Type` (`string`, required)
- `Enabled` (`bool`, required) — Whether the feature is enabled.
- `BaseAccess` (`FeatureAccessVariant1BaseAccess | null`, optional)
- `Object` (`string`, required)
- `Livemode` (`bool`, required)

### FeatureAccessVariant1BaseAccess

- `Enabled` (`bool`, required)

### FeatureAccessVariant2

- `Code` (`string`, required) — Unique feature code.
- `Name` (`string`, required) — Display name of the feature.
- `UnitName` (`string | null`, required) — Display name for one product unit, or null when not applicable.
- `Allowed` (`bool`, required) — Whether the customer can currently access or consume the feature.
- `Type` (`string`, required)
- `Consumption` (`FeatureAccessVariant2Consumption`, required)
- `BaseAccess` (`FeatureAccessVariant2BaseAccess | null`, optional)
- `Object` (`string`, required)
- `Livemode` (`bool`, required)

### FeatureAccessVariant2BaseAccess

- `IncludedUnits` (`float64`, required)
- `Unlimited` (`bool`, required)

### FeatureAccessVariant2Consumption

Variants:

- `FeatureAccessVariant2ConsumptionVariant1`
- `FeatureAccessVariant2ConsumptionVariant2`
- `FeatureAccessVariant2ConsumptionVariant3`

Discriminator: `Model`

- `"metered"` → `FeatureAccessVariant2ConsumptionVariant1`
- `"credits"` → `FeatureAccessVariant2ConsumptionVariant2`
- `"balance"` → `FeatureAccessVariant2ConsumptionVariant3`

### FeatureAccessVariant2ConsumptionVariant1

- `Model` (`string`, required) — Usage is measured against an included allowance and overage.
- `Period` (`FeatureAccessVariant2ConsumptionVariant1Period`, required) — Time range used to calculate this feature's consumption.
- `UnitsUsed` (`float64`, required) — Product units recorded during the period.
- `IncludedUnits` (`float64`, required) — Product units included in the subscription for the period.
- `RemainingUnits` (`float64`, optional) — Included units not yet consumed. Absent when usage is unlimited.
- `Unlimited` (`bool`, required) — Whether the feature has no usage limit.
- `Overage` (`FeatureAccessVariant2ConsumptionVariant1Overage`, required)

### FeatureAccessVariant2ConsumptionVariant1Overage

- `Enabled` (`bool`, required) — Whether usage above the included amount is allowed and billed.
- `Units` (`float64`, required) — Units consumed above the included amount.
- `UnitPrice` (`FeatureAccessVariant2ConsumptionVariant1OverageUnitPrice`, optional) — Price for one additional product unit.

### FeatureAccessVariant2ConsumptionVariant1OverageUnitPrice

- `Amount` (`int`, required) — Integer rate amount. Divide by scale to obtain the price.
- `Currency` (`string`, required) — Lowercase ISO 4217 currency code.
- `Scale` (`any`, required) — Divide amount by scale to obtain the major-unit price.

### FeatureAccessVariant2ConsumptionVariant1Period

- `Start` (`string`, required) — Inclusive usage period start.
- `End` (`string`, required) — Exclusive usage period end.

### FeatureAccessVariant2ConsumptionVariant2

- `Model` (`string`, required) — Product usage consumes credits from a shared pool.
- `Period` (`FeatureAccessVariant2ConsumptionVariant2Period`, required) — Time range used to calculate this feature's consumption.
- `UnitsUsed` (`float64`, required) — Product units recorded during the period.
- `CreditsPerUnit` (`int`, required) — Credits deducted for each product unit.
- `CreditsConsumed` (`float64`, required) — Actual credits deducted by this feature during the period.
- `AvailableUnits` (`int`, required) — Additional product units available from the current shared credit pool at this feature's conversion rate.

### FeatureAccessVariant2ConsumptionVariant2Period

- `Start` (`string`, required) — Inclusive usage period start.
- `End` (`string`, required) — Exclusive usage period end.

### FeatureAccessVariant2ConsumptionVariant3

- `Model` (`string`, required) — Product usage deducts money from a shared balance.
- `Period` (`FeatureAccessVariant2ConsumptionVariant3Period`, required) — Time range used to calculate this feature's consumption.
- `UnitsUsed` (`float64`, required) — Product units recorded during the period.
- `Spent` (`FeatureAccessVariant2ConsumptionVariant3Spent`, required) — Actual money deducted for this feature during the period.
- `AvailableUnits` (`int`, optional) — Estimated additional units available from the current shared balance at this feature's fixed price. Absent for dynamic pricing.
- `UnitPrice` (`FeatureAccessVariant2ConsumptionVariant3UnitPrice`, optional) — Price for one additional product unit.

### FeatureAccessVariant2ConsumptionVariant3Period

- `Start` (`string`, required) — Inclusive usage period start.
- `End` (`string`, required) — Exclusive usage period end.

### FeatureAccessVariant2ConsumptionVariant3Spent

- `Amount` (`int`, required) — Amount in the currency's smallest unit.
- `Currency` (`string`, required) — Lowercase ISO 4217 currency code.

### FeatureAccessVariant2ConsumptionVariant3UnitPrice

- `Amount` (`int`, required) — Integer rate amount. Divide by scale to obtain the price.
- `Currency` (`string`, required) — Lowercase ISO 4217 currency code.
- `Scale` (`any`, required) — Divide amount by scale to obtain the major-unit price.

### FeatureAccessVariant3

- `Code` (`string`, required) — Unique feature code.
- `Name` (`string`, required) — Display name of the feature.
- `UnitName` (`string | null`, required) — Display name for one product unit, or null when not applicable.
- `Allowed` (`bool`, required) — Whether the customer can currently access or consume the feature.
- `Type` (`string`, required)
- `Usage` (`FeatureAccessVariant3Usage`, required)
- `BaseAccess` (`FeatureAccessVariant3BaseAccess | null`, optional)
- `Object` (`string`, required)
- `Livemode` (`bool`, required)

### FeatureAccessVariant3BaseAccess

- `IncludedUnits` (`float64`, required)
- `Unlimited` (`bool`, required)

### FeatureAccessVariant3Usage

- `Period` (`FeatureAccessVariant3UsagePeriod`, required) — Time range used to calculate this feature's consumption.
- `UnitsUsed` (`float64`, required) — Current units assigned or in use.
- `IncludedUnits` (`float64`, required) — Units included in the subscription for the period.
- `RemainingUnits` (`float64`, optional) — Included units still available. Absent when usage is unlimited.
- `Unlimited` (`bool`, required) — Whether the feature has no usage limit.
- `Overage` (`FeatureAccessVariant3UsageOverage`, required)

### FeatureAccessVariant3UsageOverage

- `Enabled` (`bool`, required) — Whether usage above the included amount is allowed and billed.
- `Units` (`float64`, required) — Units consumed above the included amount.
- `UnitPrice` (`FeatureAccessVariant3UsageOverageUnitPrice`, optional) — Price for one additional product unit.

### FeatureAccessVariant3UsageOverageUnitPrice

- `Amount` (`int`, required) — Integer rate amount. Divide by scale to obtain the price.
- `Currency` (`string`, required) — Lowercase ISO 4217 currency code.
- `Scale` (`any`, required) — Divide amount by scale to obtain the major-unit price.

### FeatureAccessVariant3UsagePeriod

- `Start` (`string`, required) — Inclusive usage period start.
- `End` (`string`, required) — Exclusive usage period end.

### FeatureAccessVariant4

- `Code` (`string`, required) — Unique feature code.
- `Name` (`string`, required) — Display name of the feature.
- `UnitName` (`string | null`, required) — Display name for one product unit, or null when not applicable.
- `Allowed` (`bool`, required) — Whether the customer can currently access or consume the feature.
- `Type` (`string`, required)
- `Usage` (`FeatureAccessVariant4Usage`, required)
- `BaseAccess` (`FeatureAccessVariant4BaseAccess | null`, optional)
- `Object` (`string`, required)
- `Livemode` (`bool`, required)

### FeatureAccessVariant4BaseAccess

- `IncludedUnits` (`float64`, required)
- `Unlimited` (`bool`, required)

### FeatureAccessVariant4Usage

- `Period` (`FeatureAccessVariant4UsagePeriod`, required) — Time range used to calculate this feature's consumption.
- `UnitsUsed` (`float64`, required) — Current units assigned or in use.
- `IncludedUnits` (`float64`, required) — Units included in the subscription for the period.
- `RemainingUnits` (`float64`, optional) — Included units still available. Absent when usage is unlimited.
- `Unlimited` (`bool`, required) — Whether the feature has no usage limit.
- `Overage` (`FeatureAccessVariant4UsageOverage`, required)
- `BilledUnits` (`float64`, required) — Highest quota reached during the period and used for billing.

### FeatureAccessVariant4UsageOverage

- `Enabled` (`bool`, required) — Whether usage above the included amount is allowed and billed.
- `Units` (`float64`, required) — Units consumed above the included amount.
- `UnitPrice` (`FeatureAccessVariant4UsageOverageUnitPrice`, optional) — Price for one additional product unit.

### FeatureAccessVariant4UsageOverageUnitPrice

- `Amount` (`int`, required) — Integer rate amount. Divide by scale to obtain the price.
- `Currency` (`string`, required) — Lowercase ISO 4217 currency code.
- `Scale` (`any`, required) — Divide amount by scale to obtain the major-unit price.

### FeatureAccessVariant4UsagePeriod

- `Start` (`string`, required) — Inclusive usage period start.
- `End` (`string`, required) — Exclusive usage period end.

### FeaturesListResult

- `Object` (`string`, required)
- `Data` (`[]Feature`, required)
- `HasMore` (`bool`, required)
- `NextCursor` (`string`, optional)

### Invoice

- `ID` (`string`, required)
- `CustomerID` (`string`, required)
- `SubscriptionID` (`string | null`, required)
- `InvoiceNumber` (`string`, required)
- `Status` (`string`, required)
- `InvoiceType` (`InvoiceType`, required)
- `Currency` (`string`, required)
- `Subtotal` (`int`, required)
- `DiscountAmount` (`int`, required)
- `TaxAmount` (`int`, required)
- `Total` (`int`, required)
- `PeriodStart` (`string`, required)
- `PeriodEnd` (`string`, required)
- `IssueDate` (`string`, required)
- `DueDate` (`string`, required)
- `Memo` (`string | null`, required)
- `Metadata` (`map[string]any`, required)
- `CreatedAt` (`string`, required)
- `UpdatedAt` (`string`, required)
- `CreditApplied` (`int`, required)
- `PlanName` (`string | null`, required)
- `PoNumber` (`string | null`, required)
- `Reference` (`string | null`, required)
- `LineItems` (`[]InvoiceLineItemsItem`, required)
- `Object` (`string`, required)
- `Livemode` (`bool`, required)

### InvoiceDownload

- `URL` (`string`, required)
- `ExpiresAt` (`string`, required)
- `Object` (`string`, required)
- `Livemode` (`bool`, required)

### InvoiceLineItemsItem

- `LineType` (`string`, required)
- `FeatureName` (`string | null`, required)
- `Description` (`string`, required)
- `Quantity` (`int`, required)
- `UnitAmount` (`int`, required)
- `Amount` (`int`, required)
- `IncludedAmount` (`int | null`, required)
- `UsedAmount` (`int | null`, required)
- `OverageAmount` (`int | null`, required)
- `DiscountType` (`string | null`, required)
- `DiscountValue` (`int | null`, required)
- `DiscountName` (`string | null`, required)
- `ChargeType` (`string`, required)

### InvoiceListItem

- `ID` (`string`, required)
- `CustomerID` (`string`, required)
- `SubscriptionID` (`string | null`, required)
- `InvoiceNumber` (`string`, required)
- `Status` (`string`, required)
- `InvoiceType` (`InvoiceType`, required)
- `Currency` (`string`, required)
- `Subtotal` (`int`, required)
- `DiscountAmount` (`int`, required)
- `TaxAmount` (`int`, required)
- `Total` (`int`, required)
- `PeriodStart` (`string`, required)
- `PeriodEnd` (`string`, required)
- `IssueDate` (`string`, required)
- `DueDate` (`string`, required)
- `Memo` (`string | null`, required)
- `Metadata` (`map[string]any`, required)
- `CreatedAt` (`string`, required)
- `UpdatedAt` (`string`, required)
- `Object` (`string`, required)
- `Livemode` (`bool`, required)

### InvoicesListResult

- `Object` (`string`, required)
- `Data` (`[]InvoiceListItem`, required)
- `HasMore` (`bool`, required)
- `NextCursor` (`string`, optional)

### KeptOnVersion

- `SubscriptionID` (`string`, required)
- `KeptOnVersion` (`any`, required)
- `Object` (`string`, required)
- `Livemode` (`bool`, required)

### ManagedBy

Variants:

- `string`
- `any`

### Market

- `ID` (`string`, required)
- `Name` (`string`, required)
- `CountryCodes` (`[]string`, required)
- `Metadata` (`map[string]any`, required)
- `CreatedAt` (`string`, required)
- `UpdatedAt` (`string`, required)
- `Object` (`string`, required)
- `Livemode` (`bool`, required)

### MarketsListResult

- `Object` (`string`, required)
- `Data` (`[]Market`, required)
- `HasMore` (`bool`, required)
- `NextCursor` (`string`, optional)

### Offer

- `ID` (`string`, required)
- `Name` (`string`, required)
- `Phases` (`[]OfferPhasesItem`, required)
- `Metadata` (`map[string]any`, required)
- `StartsAt` (`string | null`, required)
- `EndsAt` (`string | null`, required)
- `Active` (`bool`, required)
- `ManagedBy` (`ManagedBy`, required) — Integration that manages this record, or null when the merchant does. While an integration manages it, its terms change only in that integration and edits answer 409 managed_by_integration.
- `CreatedAt` (`string`, required)
- `UpdatedAt` (`string`, required)
- `Object` (`string`, required)
- `Livemode` (`bool`, required)

### OfferPhasesItem

Variants:

- `OfferPhasesItemVariant1`
- `OfferPhasesItemVariant2`
- `OfferPhasesItemVariant3`

Discriminator: `Type`

- `"percentage"` → `OfferPhasesItemVariant1`
- `"amount_off"` → `OfferPhasesItemVariant2`
- `"fixed_price"` → `OfferPhasesItemVariant3`

### OfferPhasesItemVariant1

- `Type` (`string`, required)
- `DurationCycles` (`int | null`, required)
- `DurationInterval` (`string | null`, required) — Unit the phase duration is counted in. Only a fixed-price phase may set it, because its amount is declared rather than derived from the plan. Defaults to the plan's own billing interval.
- `Percentage` (`int`, required) — Discount in basis points. 5000 means 50%.

### OfferPhasesItemVariant2

- `Type` (`string`, required)
- `DurationCycles` (`int | null`, required)
- `DurationInterval` (`string | null`, required) — Unit the phase duration is counted in. Only a fixed-price phase may set it, because its amount is declared rather than derived from the plan. Defaults to the plan's own billing interval.
- `Amounts` (`[]OfferPhasesItemVariant2AmountsItem`, required)

### OfferPhasesItemVariant2AmountsItem

- `Currency` (`string`, required)
- `Amount` (`int`, required) — Amount in the currency's minor unit (for example, cents for USD).

### OfferPhasesItemVariant3

- `Type` (`string`, required)
- `DurationCycles` (`int | null`, required)
- `DurationInterval` (`string | null`, required) — Unit the phase duration is counted in. Only a fixed-price phase may set it, because its amount is declared rather than derived from the plan. Defaults to the plan's own billing interval.
- `Prices` (`[]OfferPhasesItemVariant3PricesItem`, required)

### OfferPhasesItemVariant3PricesItem

- `Currency` (`string`, required)
- `Amount` (`int`, required) — Amount in the currency's minor unit (for example, cents for USD).

### OffersListResult

- `Object` (`string`, required)
- `Data` (`[]Offer`, required)
- `HasMore` (`bool`, required)
- `NextCursor` (`string`, optional)

### Payment

- `PaymentContext` (`PaymentPaymentContext | null`, required) — Charge context captured for new payments. Null for historical payments with no captured context.
- `ID` (`string`, required)
- `CustomerID` (`string | null`, required)
- `Kind` (`string`, required)
- `Status` (`string`, required)
- `Provider` (`string`, required)
- `AmountSubtotal` (`int`, required)
- `TaxAmount` (`int`, required)
- `AmountTotal` (`int`, required)
- `Currency` (`string`, required)
- `Description` (`string`, required)
- `Metadata` (`map[string]any | null`, required)
- `URL` (`string | null`, required)
- `ExpiresAt` (`string | null`, required)
- `CreatedAt` (`string`, required)
- `UpdatedAt` (`string`, required)
- `Object` (`string`, required)
- `Livemode` (`bool`, required)

### PaymentMethodUpdateCheckout

- `CheckoutURL` (`string`, required)
- `Object` (`string`, required)
- `Livemode` (`bool`, required)

### PaymentPaymentContext

- `Reason` (`string`, required) — The original reason for the charge. Recovery never replaces this reason.
- `PaymentLinkID` (`string | null`, required) — The public payment link ID, independently of the reason, or null when no payment link originated the charge.
- `Recovery` (`PaymentPaymentContextRecovery | null`, required)

### PaymentPaymentContextRecovery

Variants:

- `PaymentPaymentContextRecoveryVariant1`
- `PaymentPaymentContextRecoveryVariant2`

Discriminator: `Type`

- `"payment_recovery"` → `PaymentPaymentContextRecoveryVariant1`
- `"dunning_retry"` → `PaymentPaymentContextRecoveryVariant2`

### PaymentPaymentContextRecoveryVariant1

- `Type` (`string`, required)

### PaymentPaymentContextRecoveryVariant2

- `Type` (`string`, required)
- `Attempt` (`int`, required) — Current retry, starting at 1. The original decline is not a retry.
- `MaxAttempts` (`int`, required) — Total retries applicable to this charge's dunning schedule.

### PaymentsListResult

- `Object` (`string`, required)
- `Data` (`[]Payment`, required)
- `HasMore` (`bool`, required)
- `NextCursor` (`string`, optional)

### Payout

- `ID` (`string`, required)
- `Status` (`string`, required)
- `Amount` (`int`, required)
- `Fee` (`int`, required)
- `NetAmount` (`int`, required)
- `Currency` (`string`, required)
- `Description` (`string | null`, required)
- `ProviderTransferID` (`string`, required)
- `CreatedAt` (`string`, required)
- `Object` (`string`, required)
- `Livemode` (`bool`, required)

### PayoutBankAccount

- `ID` (`string`, required)
- `ProviderExternalAccountID` (`string | null`, required)
- `HolderName` (`string`, required)
- `Last4` (`string`, required)
- `BankName` (`string | null`, required)
- `Country` (`string`, required)
- `Currency` (`string`, required)
- `AccountType` (`string | null`, required)
- `IsDefault` (`bool`, required)
- `Status` (`string`, required)
- `CreatedAt` (`string`, required)
- `Object` (`string`, required)
- `Livemode` (`bool`, required)

### Plan

- `ID` (`string`, required)
- `Name` (`string`, required)
- `Code` (`string`, required)
- `Description` (`string | null`, required)
- `ConsumptionModel` (`ConsumptionModel | null`, required)
- `IsPublic` (`bool`, required)
- `IsDefault` (`bool`, required)
- `IsFree` (`bool`, required)
- `BlockOnExhaustion` (`bool | null`, required) — Credits and balance plans: true stops usage when the plan and purchased amounts run out; false lets usage continue and bills the missing amount as overage at the next renewal (credits at the price per credit). Trials and free plans always stop.
- `UnusedPlanCredits` (`PlanUnusedPlanCredits | null`, required) — Credits plans only; null for other consumption models. What happens to unused plan credits at each reset. Purchased credits never expire and never count toward maxPlanCredits.
- `PlanCreditLimit` (`PlanPlanCreditLimit | null`, required) — Credits plans only. The most plan credits a subscription spends per day or per week; null means no limit.
- `UnusedPlanBalance` (`PlanUnusedPlanBalance | null`, required) — Balance plans only; null for other consumption models. What happens to unused plan balance at each reset. Purchased balance never expires and never counts toward maxPlanBalance.
- `PlanBalanceLimit` (`PlanPlanBalanceLimit | null`, required) — Balance plans only. The most plan balance a subscription spends per day or per week, in USD cents; null means no limit.
- `SortOrder` (`int`, required)
- `PlanGroupID` (`string | null`, required)
- `Metadata` (`map[string]any | null`, required)
- `CreatedAt` (`string`, required)
- `UpdatedAt` (`string`, required)
- `Features` (`[]PlanFeaturesItem`, required)
- `Prices` (`[]PlanPricesItem`, required)
- `ExchangeRates` (`[]PlanExchangeRatesItem`, required)
- `Countries` (`[]PlanCountriesItem`, required) — Countries with their own currency or prices in the current version. Buyers from any other country pay the USD prices.
- `Object` (`string`, required)
- `Livemode` (`bool`, required)

### PlanChange

Variants:

- `PlanChangeVariant1`
- `PlanChangeVariant2`
- `PlanChangeVariant3`

Discriminator: `Outcome`

- `"requires_checkout"` → `PlanChangeVariant1`
- `"scheduled"` → `PlanChangeVariant2`
- `"completed"` → `PlanChangeVariant3`

### PlanChangeVariant1

- `Outcome` (`string`, required)
- `RequiresCheckout` (`any`, required)
- `CheckoutURL` (`string`, required)
- `OfferApplication` (`PlanChangeVariant1OfferApplication`, optional)
- `Object` (`string`, required)
- `Livemode` (`bool`, required)

### PlanChangeVariant1OfferApplication

- `ID` (`string`, required)
- `OfferID` (`string`, required)
- `Name` (`string`, required)
- `Currency` (`string`, required)
- `Subtotal` (`int`, required) — Subtotal in the currency's minor unit.
- `DiscountAmount` (`int`, required) — Discount in the currency's minor unit.
- `Total` (`int`, required) — Total in the currency's minor unit.
- `Phases` (`[]PlanChangeVariant1OfferApplicationPhasesItem`, required)
- `AppliesTo` (`PlanChangeVariant1OfferApplicationAppliesTo`, required)

### PlanChangeVariant1OfferApplicationAppliesTo

Variants:

- `PlanChangeVariant1OfferApplicationAppliesToVariant1`
- `PlanChangeVariant1OfferApplicationAppliesToVariant2`
- `PlanChangeVariant1OfferApplicationAppliesToVariant3`

Discriminator: `Type`

- `"plan_price"` → `PlanChangeVariant1OfferApplicationAppliesToVariant1`
- `"addon"` → `PlanChangeVariant1OfferApplicationAppliesToVariant2`
- `"credit_pack"` → `PlanChangeVariant1OfferApplicationAppliesToVariant3`

### PlanChangeVariant1OfferApplicationAppliesToVariant1

- `Type` (`string`, required)
- `ID` (`string`, required)

### PlanChangeVariant1OfferApplicationAppliesToVariant2

- `Type` (`string`, required)
- `ID` (`string`, required)

### PlanChangeVariant1OfferApplicationAppliesToVariant3

- `Type` (`string`, required)
- `ID` (`string`, required)

### PlanChangeVariant1OfferApplicationPhasesItem

Variants:

- `PlanChangeVariant1OfferApplicationPhasesItemVariant1`
- `PlanChangeVariant1OfferApplicationPhasesItemVariant2`
- `PlanChangeVariant1OfferApplicationPhasesItemVariant3`
- `PlanChangeVariant1OfferApplicationPhasesItemVariant4`

Discriminator: `Type`

- `"free_trial"` → `PlanChangeVariant1OfferApplicationPhasesItemVariant1`
- `"percentage"` → `PlanChangeVariant1OfferApplicationPhasesItemVariant2`
- `"amount_off"` → `PlanChangeVariant1OfferApplicationPhasesItemVariant3`
- `"fixed_price"` → `PlanChangeVariant1OfferApplicationPhasesItemVariant4`

### PlanChangeVariant1OfferApplicationPhasesItemVariant1

- `Type` (`string`, required)
- `DurationDays` (`int`, required)
- `StartsAt` (`string | null`, required)
- `EndsAt` (`string | null`, required)

### PlanChangeVariant1OfferApplicationPhasesItemVariant2

- `Type` (`string`, required)
- `DurationCycles` (`int | null`, required)
- `DurationInterval` (`string | null`, required)
- `StartsAt` (`string | null`, required)
- `EndsAt` (`string | null`, required)
- `Percentage` (`int`, required) — Discount in basis points. 5000 means 50%.

### PlanChangeVariant1OfferApplicationPhasesItemVariant3

- `Type` (`string`, required)
- `DurationCycles` (`int | null`, required)
- `DurationInterval` (`string | null`, required)
- `StartsAt` (`string | null`, required)
- `EndsAt` (`string | null`, required)
- `Amount` (`int`, required) — Discount in the application currency's minor unit.

### PlanChangeVariant1OfferApplicationPhasesItemVariant4

- `Type` (`string`, required)
- `DurationCycles` (`int | null`, required)
- `DurationInterval` (`string | null`, required)
- `StartsAt` (`string | null`, required)
- `EndsAt` (`string | null`, required)
- `Price` (`int`, required) — Fixed price in the application currency's minor unit.

### PlanChangeVariant2

- `Outcome` (`string`, required)
- `ID` (`string`, required)
- `Scheduled` (`any`, required)
- `ScheduledFor` (`string`, required)
- `ChangeType` (`string`, required)
- `CustomerID` (`string`, required)
- `NewPlanID` (`string`, optional)
- `NewPlanName` (`string`, optional)
- `NewBillingInterval` (`string`, optional)
- `SeatLimitWarning` (`PlanChangeVariant2SeatLimitWarning`, optional)
- `Object` (`string`, required)
- `Livemode` (`bool`, required)

### PlanChangeVariant2SeatLimitWarning

- `FeatureCode` (`string`, required)
- `FeatureName` (`string`, required)
- `CurrentSeats` (`int`, required)
- `Included` (`int`, required)
- `NewPlanName` (`string`, required)
- `EffectiveDate` (`string`, required)

### PlanChangeVariant3

- `Outcome` (`string`, required)
- `ID` (`string`, required)
- `Scheduled` (`any`, required)
- `CustomerID` (`string`, required)
- `PreviousPlan` (`PlanChangeVariant3PreviousPlan`, required)
- `CurrentPlan` (`PlanChangeVariant3CurrentPlan`, required)
- `BillingInterval` (`string`, required)
- `Billing` (`PlanChangeVariant3Billing`, required)
- `InvoiceID` (`string`, optional)
- `OfferApplication` (`PlanChangeVariant3OfferApplication`, optional)
- `Object` (`string`, required)
- `Livemode` (`bool`, required)

### PlanChangeVariant3Billing

- `Credit` (`int`, required)
- `CreditsApplied` (`int`, required)
- `Charge` (`int`, required)
- `TaxAmount` (`int`, required)
- `NetAmount` (`int`, required)
- `TotalCharged` (`int`, required)
- `RemainingCreditBalance` (`int`, required)

### PlanChangeVariant3CurrentPlan

- `ID` (`string`, required)
- `Name` (`string`, required)
- `Price` (`int`, required)

### PlanChangeVariant3OfferApplication

- `ID` (`string`, required)
- `OfferID` (`string`, required)
- `Name` (`string`, required)
- `Currency` (`string`, required)
- `Subtotal` (`int`, required) — Subtotal in the currency's minor unit.
- `DiscountAmount` (`int`, required) — Discount in the currency's minor unit.
- `Total` (`int`, required) — Total in the currency's minor unit.
- `Phases` (`[]PlanChangeVariant3OfferApplicationPhasesItem`, required)
- `AppliesTo` (`PlanChangeVariant3OfferApplicationAppliesTo`, required)

### PlanChangeVariant3OfferApplicationAppliesTo

Variants:

- `PlanChangeVariant3OfferApplicationAppliesToVariant1`
- `PlanChangeVariant3OfferApplicationAppliesToVariant2`
- `PlanChangeVariant3OfferApplicationAppliesToVariant3`

Discriminator: `Type`

- `"plan_price"` → `PlanChangeVariant3OfferApplicationAppliesToVariant1`
- `"addon"` → `PlanChangeVariant3OfferApplicationAppliesToVariant2`
- `"credit_pack"` → `PlanChangeVariant3OfferApplicationAppliesToVariant3`

### PlanChangeVariant3OfferApplicationAppliesToVariant1

- `Type` (`string`, required)
- `ID` (`string`, required)

### PlanChangeVariant3OfferApplicationAppliesToVariant2

- `Type` (`string`, required)
- `ID` (`string`, required)

### PlanChangeVariant3OfferApplicationAppliesToVariant3

- `Type` (`string`, required)
- `ID` (`string`, required)

### PlanChangeVariant3OfferApplicationPhasesItem

Variants:

- `PlanChangeVariant3OfferApplicationPhasesItemVariant1`
- `PlanChangeVariant3OfferApplicationPhasesItemVariant2`
- `PlanChangeVariant3OfferApplicationPhasesItemVariant3`
- `PlanChangeVariant3OfferApplicationPhasesItemVariant4`

Discriminator: `Type`

- `"free_trial"` → `PlanChangeVariant3OfferApplicationPhasesItemVariant1`
- `"percentage"` → `PlanChangeVariant3OfferApplicationPhasesItemVariant2`
- `"amount_off"` → `PlanChangeVariant3OfferApplicationPhasesItemVariant3`
- `"fixed_price"` → `PlanChangeVariant3OfferApplicationPhasesItemVariant4`

### PlanChangeVariant3OfferApplicationPhasesItemVariant1

- `Type` (`string`, required)
- `DurationDays` (`int`, required)
- `StartsAt` (`string | null`, required)
- `EndsAt` (`string | null`, required)

### PlanChangeVariant3OfferApplicationPhasesItemVariant2

- `Type` (`string`, required)
- `DurationCycles` (`int | null`, required)
- `DurationInterval` (`string | null`, required)
- `StartsAt` (`string | null`, required)
- `EndsAt` (`string | null`, required)
- `Percentage` (`int`, required) — Discount in basis points. 5000 means 50%.

### PlanChangeVariant3OfferApplicationPhasesItemVariant3

- `Type` (`string`, required)
- `DurationCycles` (`int | null`, required)
- `DurationInterval` (`string | null`, required)
- `StartsAt` (`string | null`, required)
- `EndsAt` (`string | null`, required)
- `Amount` (`int`, required) — Discount in the application currency's minor unit.

### PlanChangeVariant3OfferApplicationPhasesItemVariant4

- `Type` (`string`, required)
- `DurationCycles` (`int | null`, required)
- `DurationInterval` (`string | null`, required)
- `StartsAt` (`string | null`, required)
- `EndsAt` (`string | null`, required)
- `Price` (`int`, required) — Fixed price in the application currency's minor unit.

### PlanChangeVariant3PreviousPlan

- `ID` (`string`, required)
- `Name` (`string`, required)

### PlanCountriesItem

- `Country` (`string`, required) — ISO country code.
- `Currency` (`string`, required) — The currency buyers from the country pay, USD included.
- `Prices` (`[]PlanCountriesItemPricesItem`, required)
- `Addons` (`[]PlanCountriesItemAddonsItem`, required) — Every add-on the version sells, priced for the country.
- `CreditPacks` (`[]PlanCountriesItemCreditPacksItem`, required) — Every credit pack the version sells, priced for the country.

### PlanCountriesItemAddonsItem

- `AddonID` (`string`, required)
- `Price` (`int`, required) — What buyers from the country pay for the add-on, in the currency's minor unit.
- `Manual` (`bool`, required) — True when set by hand; false when it follows the add-on's USD price at the exchange rate.

### PlanCountriesItemCreditPacksItem

- `CreditPackID` (`string`, required)
- `Price` (`int`, required) — What buyers from the country pay for the pack, in the currency's minor unit.
- `Manual` (`bool`, required) — True when set by hand; false when it follows the pack's USD price at the exchange rate.

### PlanCountriesItemPricesItem

- `BillingInterval` (`BillingInterval`, required)
- `Price` (`int`, required) — The country's price in the currency's minor unit.
- `Manual` (`bool`, required) — True when set by hand; false when it follows the USD base at the exchange rate.

### PlanDraft

- `ID` (`string`, required) — Public plan draft ID.
- `FromVersionID` (`string`, required) — The version the draft started from.
- `Changes` (`[]PlanDraftChangesItem`, required) — What the draft changes against the version it started from.
- `CreatedAt` (`string`, required)
- `UpdatedAt` (`string`, required)
- `Object` (`string`, required)
- `Livemode` (`bool`, required)

### PlanDraftChangesItem

- `Key` (`string`, required) — Identifies the change; send it to undo just this change.
- `Section` (`string`, required) — Where the change is: Price, Trial, a rules section, Add-ons, or a country.
- `Country` (`string | null`, required) — ISO country code when the change belongs to one country.
- `Label` (`string`, required) — What changed, such as Monthly price.
- `Before` (`string | null`, required) — The value before, formatted.
- `After` (`string | null`, required) — The value now, formatted.

### PlanDraftPublication

- `Preview` (`bool`, required)
- `Version` (`PlanDraftPublicationVersion | null`, required) — The published version. Null for a preview.
- `MoveID` (`string | null`, required) — The move of subscribers, if any.
- `Subscribers` (`PlanDraftPublicationSubscribers | null`, required)
- `Object` (`string`, required)
- `Livemode` (`bool`, required)

### PlanDraftPublicationSubscribers

- `Moving` (`int`, required)
- `KeptOnVersion` (`int`, required)
- `NeedsAChange` (`int`, required) — Left on their version because no change makes them fit; they show among the subscribers that need a change.
- `LeftOut` (`[]PlanDraftPublicationSubscribersLeftOutItem`, required)

### PlanDraftPublicationSubscribersLeftOutItem

- `SubscriptionID` (`string`, required)
- `Code` (`string`, required)
- `Message` (`string`, required)

### PlanDraftPublicationVersion

- `ID` (`string`, required) — Public plan version ID.
- `Label` (`string`, required) — The plan's versions are numbered 1, 2, 3; an experiment's own versions read with its code, such as price-test 2.
- `Number` (`int | null`, required) — The plan's version number. Null for a version that only belongs to an experiment; an experiment version made current takes the next number.
- `Experiment` (`PlanDraftPublicationVersionExperiment | null`, required) — The experiment this version belongs to, if any.
- `IsCurrent` (`bool`, required) — Whether new buyers get this version.
- `ParentVersionID` (`string | null`, required) — The version this one was created from.
- `SubscriptionCount` (`int`, required) — Active, trialing, past due and paused subscriptions on it.
- `MaxPlanCredits` (`int | null`, required)
- `MaxPlanBalance` (`int | null`, required)
- `Prices` (`[]PlanDraftPublicationVersionPricesItem`, required)
- `CreatedAt` (`string`, required)

### PlanDraftPublicationVersionExperiment

- `Code` (`string`, required) — The experiment's code, as sold with planCode.
- `Version` (`int`, required) — Its number within the experiment.

### PlanDraftPublicationVersionPricesItem

- `ID` (`string`, required) — Public plan price ID.
- `BillingInterval` (`string`, required)
- `Price` (`int`, required)
- `IncludedCredits` (`int | null`, required)
- `IncludedBalance` (`int | null`, required)
- `RegionalPrices` (`[]PlanDraftPublicationVersionPricesItemRegionalPricesItem`, required)

### PlanDraftPublicationVersionPricesItemRegionalPricesItem

- `Currency` (`string`, required)
- `Price` (`int`, required)
- `IncludedBalance` (`int | null`, required)

### PlanExchangeRatesItem

- `Currency` (`string`, required)
- `ExchangeRate` (`float64`, required)

### PlanExperiment

- `ID` (`string`, required) — Public plan experiment ID.
- `Code` (`string`, required) — The code to sell it with as planCode, such as pro:price-test. It always sells the experiment's latest version while running, and the plan's current version once closed.
- `Status` (`string`, required) — running sells its latest version; won became the plan's current version; archived stopped selling.
- `StartedFromVersionID` (`string`, required) — The plan's version that was current when it started.
- `LatestVersionID` (`string`, required) — Its latest version.
- `SellsVersionID` (`string`, required) — The version its code sells to new buyers today.
- `SubscriptionCount` (`int`, required) — Active, trialing, past due and paused subscriptions on its versions.
- `CreatedAt` (`string`, required)
- `ClosedAt` (`string | null`, required)
- `Object` (`string`, required)
- `Livemode` (`bool`, required)

### PlanExperimentArchive

- `Experiment` (`PlanExperimentArchiveExperiment`, required)
- `MoveID` (`string | null`, required) — The move of its subscribers, or null when they stay.
- `Object` (`string`, required)
- `Livemode` (`bool`, required)

### PlanExperimentArchiveExperiment

- `ID` (`string`, required) — Public plan experiment ID.
- `Code` (`string`, required) — The code to sell it with as planCode, such as pro:price-test. It always sells the experiment's latest version while running, and the plan's current version once closed.
- `Status` (`string`, required) — running sells its latest version; won became the plan's current version; archived stopped selling.
- `StartedFromVersionID` (`string`, required) — The plan's version that was current when it started.
- `LatestVersionID` (`string`, required) — Its latest version.
- `SellsVersionID` (`string`, required) — The version its code sells to new buyers today.
- `SubscriptionCount` (`int`, required) — Active, trialing, past due and paused subscriptions on its versions.
- `CreatedAt` (`string`, required)
- `ClosedAt` (`string | null`, required)

### PlanFeature

- `PlanID` (`string`, required)
- `FeatureID` (`string`, required)
- `Enabled` (`bool`, required)
- `IncludedAmount` (`int`, required)
- `Unlimited` (`bool`, required)
- `Overage` (`PlanFeatureOverage`, required)
- `CreditsPerUnit` (`int | null`, required)
- `PricingMode` (`string`, required)
- `Margin` (`int | null`, required)
- `Object` (`string`, required)
- `Livemode` (`bool`, required)

### PlanFeatureOverage

- `Enabled` (`bool`, required)
- `UnitPrice` (`int`, required)

### PlanFeaturesItem

- `Code` (`string`, required)
- `Name` (`string`, required)
- `Type` (`FeatureType`, required)
- `UnitName` (`string | null`, required)
- `Enabled` (`bool`, required)
- `IncludedAmount` (`int | null`, required)
- `Unlimited` (`bool`, required)
- `Overage` (`PlanFeaturesItemOverage | null`, required)
- `RegionalPrices` (`[]PlanFeaturesItemRegionalPricesItem`, required)

### PlanFeaturesItemOverage

- `Enabled` (`bool`, required)
- `Model` (`string | null`, required)
- `UnitPrice` (`int | null`, required)

### PlanFeaturesItemRegionalPricesItem

- `Currency` (`string`, required)
- `OverageUnitPrice` (`int | null`, required)
- `AutoSynced` (`bool`, required)

### PlanGrant

- `ID` (`string`, required)
- `CustomerID` (`string`, required)
- `SubscriptionID` (`string`, required)
- `BasePlanID` (`string`, required)
- `PlanID` (`string`, required)
- `PlanReleaseID` (`string`, required)
- `Status` (`string`, required)
- `Duration` (`string`, required)
- `DurationCycles` (`int | null`, required)
- `StartsAt` (`string`, required)
- `ExpiresAt` (`string | null`, required)
- `Reason` (`string`, required)
- `Source` (`string`, required)
- `RevokedAt` (`string | null`, required)
- `CreatedAt` (`string`, required)
- `UpdatedAt` (`string`, required)
- `Events` (`[]PlanGrantEventsItem`, required)
- `Object` (`string`, required)
- `Livemode` (`bool`, required)

### PlanGrantEventsItem

- `ID` (`string`, required)
- `Type` (`string`, required)
- `Reason` (`string`, required)
- `Source` (`string`, required)
- `PreviousExpiresAt` (`string | null`, required)
- `ExpiresAt` (`string | null`, required)
- `Duration` (`string | null`, required)
- `DurationCycles` (`int | null`, required)
- `RequestedExpiresAt` (`string | null`, required)
- `CreatedAt` (`string`, required)

### PlanGroup

- `ID` (`string`, required)
- `Name` (`string`, required)
- `Description` (`string | null`, required)
- `IsPublic` (`bool`, required)
- `CreatedAt` (`string`, required)
- `UpdatedAt` (`string`, required)
- `Object` (`string`, required)
- `Livemode` (`bool`, required)

### PlanGroupDetail

- `ID` (`string`, required)
- `Name` (`string`, required)
- `Description` (`string | null`, required)
- `IsPublic` (`bool`, required)
- `CreatedAt` (`string`, required)
- `UpdatedAt` (`string`, required)
- `Plans` (`[]PlanGroupDetailPlansItem`, required)
- `Object` (`string`, required)
- `Livemode` (`bool`, required)

### PlanGroupDetailPlansItem

- `ID` (`string`, required)
- `Name` (`string`, required)
- `SortOrder` (`int`, required)

### PlanGroupsListResult

- `Object` (`string`, required)
- `Data` (`[]PlanGroup`, required)
- `HasMore` (`bool`, required)
- `NextCursor` (`string`, optional)

### PlanMove

- `ID` (`string | null`, required) — Public move ID. Null for a preview, which records nothing.
- `ToVersionID` (`string`, required)
- `FromVersionIds` (`[]string`, required) — The versions its subscribers are moving from.
- `Effective` (`string`, required)
- `Preview` (`bool`, required)
- `Status` (`string`, required) — checked is a preview where everyone fits; moving while subscribers wait for their move; completed once every move ran; canceled after it was canceled.
- `Counts` (`PlanMoveCounts`, required)
- `Subscriptions` (`[]PlanMoveSubscriptionsItem`, required)
- `CreatedAt` (`string | null`, required)
- `CompletedAt` (`string | null`, required)
- `CanceledAt` (`string | null`, required)
- `Object` (`string`, required)
- `Livemode` (`bool`, required)

### PlanMoveCounts

- `Moving` (`int`, required)
- `Moved` (`int`, required)
- `NeedsAChange` (`int`, required)
- `LeftOut` (`int`, required)
- `Canceled` (`int`, required)
- `KeptOnVersion` (`int`, required)

### PlanMoveSubscriptionsItem

- `SubscriptionID` (`string`, required)
- `Status` (`string`, required) — moving waits for its renewal or is being moved; moved is on the destination; needs_a_change could not move and keeps its version until it is resolved or kept; left_out was not moved because it was not active, trialing or past due, such as a paused subscription (invalid_subscription_status), or because a listed subscription was already on the destination (no_change_requested); canceled was canceled before it moved; kept_on_version was left on its version by the merchant.
- `Code` (`string | null`, required)
- `Message` (`string | null`, required)

### PlanPlanBalanceLimit

- `Amount` (`int`, required) — The most plan balance a subscription spends in one window, in USD cents. Other currencies convert it with the plan's exchange rate unless the dashboard sets a custom amount for that currency.
- `Interval` (`string`, required) — The window length: `day` is 24 hours and `week` is 7 days, starting at the time of day of the subscription's billing anchor.

### PlanPlanCreditLimit

- `Credits` (`int`, required) — The most plan credits a subscription spends in one window.
- `Interval` (`string`, required) — The window length: `day` is 24 hours and `week` is 7 days, starting at the time of day of the subscription's billing anchor.

### PlanPrice

- `ID` (`string`, required) — Public plan price ID.
- `PlanID` (`string`, required)
- `BillingInterval` (`BillingInterval`, required)
- `Price` (`int`, required) — Price in the currency's minor unit (for example, cents for USD).
- `IsDefault` (`bool`, required)
- `TrialDays` (`int`, required)
- `TrialRequiresPaymentMethod` (`bool | null`, required) — Whether the price's free trial needs a payment method to start. When false, a new subscription starts trialing right away without checkout; without a payment method added by the trial end, the customer moves to the plan group's default plan, or the subscription is canceled when there is none. Null when the price has no trial.
- `IncludedBalance` (`int | null`, required)
- `IncludedCredits` (`int | null`, required)
- `OfferID` (`string | null`, required) — Automatic introductory offer for this price.
- `InheritsFromPriceID` (`string | null`, required) — Public base price ID for a market price variant, or null for a base price.
- `Metadata` (`map[string]any`, required) — Application metadata. Variant display names may use metadata.name.
- `MarketPrices` (`[]PlanPriceMarketPricesItem`, required) — Country-market overrides. Variants inherit their base price for every market not listed.
- `CreatedAt` (`string`, required)
- `UpdatedAt` (`string`, required)
- `Object` (`string`, required)
- `Livemode` (`bool`, required)

### PlanPriceMarketPricesItem

- `MarketGroupID` (`string`, required) — Public pricing market group ID.
- `Currency` (`string`, required) — Presentment currency for this market.
- `Price` (`int`, required) — Market price in the currency's minor unit.

### PlanPricesItem

- `ID` (`string`, required) — Public plan price ID.
- `BillingInterval` (`BillingInterval`, required)
- `Price` (`int`, required) — Price in the currency's minor unit (for example, cents for USD).
- `IsDefault` (`bool`, required)
- `TrialDays` (`int`, required)
- `TrialRequiresPaymentMethod` (`bool | null`, required) — Whether the price's free trial needs a payment method to start. When false, a new subscription starts trialing right away without checkout; without a payment method added by the trial end, the customer moves to the plan group's default plan, or the subscription is canceled when there is none. Null when the price has no trial.
- `IncludedBalance` (`int | null`, required)
- `IncludedCredits` (`int | null`, required)
- `OfferID` (`string | null`, required) — Automatic introductory offer for this price. Pass a Promotional Offer ID when creating a subscription to override it.
- `InheritsFromPriceID` (`string | null`, required) — Public base price ID for a market price variant, or null for a base price.
- `Metadata` (`map[string]any`, required) — Application metadata. Variant display names may use metadata.name.
- `MarketPrices` (`[]PlanPricesItemMarketPricesItem`, required) — Country-market overrides. An empty array means currency pricing and then the global USD price remain the fallback.
- `RegionalPrices` (`[]PlanPricesItemRegionalPricesItem`, required)

### PlanPricesItemMarketPricesItem

- `MarketGroupID` (`string`, required) — Public pricing market group ID.
- `Currency` (`string`, required) — Presentment currency for this market.
- `Price` (`int`, required) — Market price in the currency's minor unit.

### PlanPricesItemRegionalPricesItem

- `Currency` (`string`, required)
- `Price` (`int`, required)
- `IncludedBalance` (`int | null`, required)
- `AutoSynced` (`bool`, required)

### PlanRegionalPricing

- `PriceID` (`string`, required)
- `Overrides` (`[]PlanRegionalPricingOverridesItem`, required)
- `Object` (`string`, required)
- `Livemode` (`bool`, required)

### PlanRegionalPricingOverridesItem

- `Currency` (`string`, required)
- `Price` (`int`, required)
- `IncludedBalance` (`int`, optional)

### PlanRegionalPricingResult

- `PlanID` (`string`, required)
- `Currency` (`string`, required)
- `ExchangeRate` (`float64`, required)
- `PricesConfigured` (`int`, required)
- `FeaturesConfigured` (`int`, required)
- `Object` (`string`, required)
- `Livemode` (`bool`, required)

### PlansListDraftsResult

- `Object` (`string`, required)
- `Data` (`[]PlanDraft`, required)
- `HasMore` (`bool`, required)
- `NextCursor` (`string`, optional)

### PlansListExperimentsResult

- `Object` (`string`, required)
- `Data` (`[]PlanExperiment`, required)
- `HasMore` (`bool`, required)
- `NextCursor` (`string`, optional)

### PlansListMovesResult

- `Object` (`string`, required)
- `Data` (`[]PlanMove`, required)
- `HasMore` (`bool`, required)
- `NextCursor` (`string`, optional)

### PlansListResult

- `Object` (`string`, required)
- `Data` (`[]Plan`, required)
- `HasMore` (`bool`, required)
- `NextCursor` (`string`, optional)

### PlansListSubscribersNeedingAChangeResult

- `Object` (`string`, required)
- `Data` (`[]SubscriberNeedingAChange`, required)
- `HasMore` (`bool`, required)
- `NextCursor` (`string`, optional)

### PlansListVersionsResult

- `Object` (`string`, required)
- `Data` (`[]PlanVersion`, required)
- `HasMore` (`bool`, required)
- `NextCursor` (`string`, optional)

### PlanUnusedPlanBalance

Variants:

- `PlanUnusedPlanBalanceVariant1`
- `PlanUnusedPlanBalanceVariant2`

Discriminator: `Type`

- `"expire"` → `PlanUnusedPlanBalanceVariant1`
- `"roll_over"` → `PlanUnusedPlanBalanceVariant2`

### PlanUnusedPlanBalanceVariant1

- `Type` (`string`, required)

### PlanUnusedPlanBalanceVariant2

- `Type` (`string`, required)
- `MaxPlanBalance` (`int | null`, required) — The most plan balance a subscription holds after a reset, in USD cents, including the balance that reset grants; balance above it expires. At least the included balance of every price. Other currencies convert it with the plan's exchange rate unless the dashboard sets a custom amount for that currency. null means no limit other than the largest storable amount, 2,147,483,647 cents.

### PlanUnusedPlanCredits

Variants:

- `PlanUnusedPlanCreditsVariant1`
- `PlanUnusedPlanCreditsVariant2`

Discriminator: `Type`

- `"expire"` → `PlanUnusedPlanCreditsVariant1`
- `"roll_over"` → `PlanUnusedPlanCreditsVariant2`

### PlanUnusedPlanCreditsVariant1

- `Type` (`string`, required)

### PlanUnusedPlanCreditsVariant2

- `Type` (`string`, required)
- `MaxPlanCredits` (`int | null`, required) — The most plan credits a subscription holds after a reset, including the credits that reset grants; credits above it expire. At least the included credits of every price. null means no limit other than the largest storable balance, 2,147,483,647 plan credits.

### PlanVersion

- `ID` (`string`, required) — Public plan version ID.
- `Label` (`string`, required) — The plan's versions are numbered 1, 2, 3; an experiment's own versions read with its code, such as price-test 2.
- `Number` (`int | null`, required) — The plan's version number. Null for a version that only belongs to an experiment; an experiment version made current takes the next number.
- `Experiment` (`PlanVersionExperiment | null`, required) — The experiment this version belongs to, if any.
- `IsCurrent` (`bool`, required) — Whether new buyers get this version.
- `ParentVersionID` (`string | null`, required) — The version this one was created from.
- `SubscriptionCount` (`int`, required) — Active, trialing, past due and paused subscriptions on it.
- `MaxPlanCredits` (`int | null`, required)
- `MaxPlanBalance` (`int | null`, required)
- `Prices` (`[]PlanVersionPricesItem`, required)
- `CreatedAt` (`string`, required)
- `Object` (`string`, required)
- `Livemode` (`bool`, required)

### PlanVersionComparison

- `VersionID` (`string`, required) — The version the changes start from.
- `AgainstVersionID` (`string`, required) — The version they end at.
- `Changes` (`[]PlanVersionComparisonChangesItem`, required) — Every term that differs, the plan's first and then each country.
- `Object` (`string`, required)
- `Livemode` (`bool`, required)

### PlanVersionComparisonChangesItem

- `Key` (`string`, required) — Identifies the change; send it to undo just this change.
- `Section` (`string`, required) — Where the change is: Price, Trial, a rules section, Add-ons, or a country.
- `Country` (`string | null`, required) — ISO country code when the change belongs to one country.
- `Label` (`string`, required) — What changed, such as Monthly price.
- `Before` (`string | null`, required) — The value before, formatted.
- `After` (`string | null`, required) — The value now, formatted.

### PlanVersionExperiment

- `Code` (`string`, required) — The experiment's code, as sold with planCode.
- `Version` (`int`, required) — Its number within the experiment.

### PlanVersionPricesItem

- `ID` (`string`, required) — Public plan price ID.
- `BillingInterval` (`string`, required)
- `Price` (`int`, required)
- `IncludedCredits` (`int | null`, required)
- `IncludedBalance` (`int | null`, required)
- `RegionalPrices` (`[]PlanVersionPricesItemRegionalPricesItem`, required)

### PlanVersionPricesItemRegionalPricesItem

- `Currency` (`string`, required)
- `Price` (`int`, required)
- `IncludedBalance` (`int | null`, required)

### PortalAccess

- `PortalURL` (`string`, required)
- `Object` (`string`, required)
- `Livemode` (`bool`, required)

### PrepaidPeriodsCheckout

- `CheckoutURL` (`string`, required)
- `Object` (`string`, required)
- `Livemode` (`bool`, required)

### PreviewChange

- `Currency` (`string`, required)
- `CurrentPlanCredit` (`int`, required)
- `NewPlanCharge` (`int`, required)
- `EstimatedTotal` (`int`, required)
- `EffectiveDate` (`string`, required)
- `DaysRemaining` (`int`, required)
- `TotalDays` (`int`, required)
- `IsUpgrade` (`bool`, required)
- `DeactivatedAddons` (`[]PreviewChangeDeactivatedAddonsItem`, required) — Active add-ons that the version the change moves the subscription to does not sell. The change deactivates them when it applies.
- `OfferApplication` (`PreviewChangeOfferApplication`, optional)
- `Object` (`string`, required)
- `Livemode` (`bool`, required)

### PreviewChangeDeactivatedAddonsItem

- `AddonID` (`string`, required)
- `Name` (`string`, required)

### PreviewChangeOfferApplication

- `ID` (`string`, required)
- `OfferID` (`string`, required)
- `Name` (`string`, required)
- `Currency` (`string`, required)
- `Subtotal` (`int`, required) — Subtotal in the currency's minor unit.
- `DiscountAmount` (`int`, required) — Discount in the currency's minor unit.
- `Total` (`int`, required) — Total in the currency's minor unit.
- `Phases` (`[]PreviewChangeOfferApplicationPhasesItem`, required)
- `AppliesTo` (`PreviewChangeOfferApplicationAppliesTo`, required)

### PreviewChangeOfferApplicationAppliesTo

Variants:

- `PreviewChangeOfferApplicationAppliesToVariant1`
- `PreviewChangeOfferApplicationAppliesToVariant2`
- `PreviewChangeOfferApplicationAppliesToVariant3`

Discriminator: `Type`

- `"plan_price"` → `PreviewChangeOfferApplicationAppliesToVariant1`
- `"addon"` → `PreviewChangeOfferApplicationAppliesToVariant2`
- `"credit_pack"` → `PreviewChangeOfferApplicationAppliesToVariant3`

### PreviewChangeOfferApplicationAppliesToVariant1

- `Type` (`string`, required)
- `ID` (`string`, required)

### PreviewChangeOfferApplicationAppliesToVariant2

- `Type` (`string`, required)
- `ID` (`string`, required)

### PreviewChangeOfferApplicationAppliesToVariant3

- `Type` (`string`, required)
- `ID` (`string`, required)

### PreviewChangeOfferApplicationPhasesItem

Variants:

- `PreviewChangeOfferApplicationPhasesItemVariant1`
- `PreviewChangeOfferApplicationPhasesItemVariant2`
- `PreviewChangeOfferApplicationPhasesItemVariant3`
- `PreviewChangeOfferApplicationPhasesItemVariant4`

Discriminator: `Type`

- `"free_trial"` → `PreviewChangeOfferApplicationPhasesItemVariant1`
- `"percentage"` → `PreviewChangeOfferApplicationPhasesItemVariant2`
- `"amount_off"` → `PreviewChangeOfferApplicationPhasesItemVariant3`
- `"fixed_price"` → `PreviewChangeOfferApplicationPhasesItemVariant4`

### PreviewChangeOfferApplicationPhasesItemVariant1

- `Type` (`string`, required)
- `DurationDays` (`int`, required)
- `StartsAt` (`string | null`, required)
- `EndsAt` (`string | null`, required)

### PreviewChangeOfferApplicationPhasesItemVariant2

- `Type` (`string`, required)
- `DurationCycles` (`int | null`, required)
- `DurationInterval` (`string | null`, required)
- `StartsAt` (`string | null`, required)
- `EndsAt` (`string | null`, required)
- `Percentage` (`int`, required) — Discount in basis points. 5000 means 50%.

### PreviewChangeOfferApplicationPhasesItemVariant3

- `Type` (`string`, required)
- `DurationCycles` (`int | null`, required)
- `DurationInterval` (`string | null`, required)
- `StartsAt` (`string | null`, required)
- `EndsAt` (`string | null`, required)
- `Amount` (`int`, required) — Discount in the application currency's minor unit.

### PreviewChangeOfferApplicationPhasesItemVariant4

- `Type` (`string`, required)
- `DurationCycles` (`int | null`, required)
- `DurationInterval` (`string | null`, required)
- `StartsAt` (`string | null`, required)
- `EndsAt` (`string | null`, required)
- `Price` (`int`, required) — Fixed price in the application currency's minor unit.

### PromoCode

- `ID` (`string`, required)
- `Code` (`string`, required)
- `OfferID` (`string`, required)
- `BillingInterval` (`BillingInterval | null`, required)
- `MaxRedemptions` (`int | null`, required)
- `ExpiresAt` (`string | null`, required)
- `IsActive` (`bool`, required)
- `RedemptionCount` (`int`, required)
- `FirstPurchaseOnly` (`bool`, required) — Whether the code only applies to a customer's first purchase: one with no paid invoice above zero and no paid subscription.
- `ManagedBy` (`ManagedBy`, required) — Integration that manages this record, or null when the merchant does. While an integration manages it, its terms change only in that integration and edits answer 409 managed_by_integration.
- `CreatedAt` (`string`, required)
- `UpdatedAt` (`string`, required)
- `Object` (`string`, required)
- `Livemode` (`bool`, required)

### PromoCodesListResult

- `Object` (`string`, required)
- `Data` (`[]PromoCode`, required)
- `HasMore` (`bool`, required)
- `NextCursor` (`string`, optional)

### PublishPlanDraftParamsChangesItem

- `SubscriptionID` (`string`, required)
- `KeepOnVersion` (`any`, optional) — Leave this subscriber on its version.
- `BillingInterval` (`string`, optional)
- `SeatLimit` (`any`, optional) — Lower its seats to the version's limit at renewal.
- `RemoveAddonIds` (`[]string`, optional) — Add-ons it stops paying at renewal.

### QuotaGetAllResult

- `Object` (`string`, required)
- `Data` (`[]UsageQuota`, required)
- `HasMore` (`bool`, required)
- `NextCursor` (`string`, optional)

### ReactivatedSubscription

- `SubscriptionID` (`string`, required)
- `InvoiceID` (`string`, required)
- `Status` (`string`, required)
- `OfferApplication` (`ReactivatedSubscriptionOfferApplication`, optional)
- `Object` (`string`, required)
- `Livemode` (`bool`, required)

### ReactivatedSubscriptionOfferApplication

- `ID` (`string`, required)
- `OfferID` (`string`, required)
- `Name` (`string`, required)
- `Currency` (`string`, required)
- `Subtotal` (`int`, required) — Subtotal in the currency's minor unit.
- `DiscountAmount` (`int`, required) — Discount in the currency's minor unit.
- `Total` (`int`, required) — Total in the currency's minor unit.
- `Phases` (`[]ReactivatedSubscriptionOfferApplicationPhasesItem`, required)
- `AppliesTo` (`ReactivatedSubscriptionOfferApplicationAppliesTo`, required)

### ReactivatedSubscriptionOfferApplicationAppliesTo

Variants:

- `ReactivatedSubscriptionOfferApplicationAppliesToVariant1`
- `ReactivatedSubscriptionOfferApplicationAppliesToVariant2`
- `ReactivatedSubscriptionOfferApplicationAppliesToVariant3`

Discriminator: `Type`

- `"plan_price"` → `ReactivatedSubscriptionOfferApplicationAppliesToVariant1`
- `"addon"` → `ReactivatedSubscriptionOfferApplicationAppliesToVariant2`
- `"credit_pack"` → `ReactivatedSubscriptionOfferApplicationAppliesToVariant3`

### ReactivatedSubscriptionOfferApplicationAppliesToVariant1

- `Type` (`string`, required)
- `ID` (`string`, required)

### ReactivatedSubscriptionOfferApplicationAppliesToVariant2

- `Type` (`string`, required)
- `ID` (`string`, required)

### ReactivatedSubscriptionOfferApplicationAppliesToVariant3

- `Type` (`string`, required)
- `ID` (`string`, required)

### ReactivatedSubscriptionOfferApplicationPhasesItem

Variants:

- `ReactivatedSubscriptionOfferApplicationPhasesItemVariant1`
- `ReactivatedSubscriptionOfferApplicationPhasesItemVariant2`
- `ReactivatedSubscriptionOfferApplicationPhasesItemVariant3`
- `ReactivatedSubscriptionOfferApplicationPhasesItemVariant4`

Discriminator: `Type`

- `"free_trial"` → `ReactivatedSubscriptionOfferApplicationPhasesItemVariant1`
- `"percentage"` → `ReactivatedSubscriptionOfferApplicationPhasesItemVariant2`
- `"amount_off"` → `ReactivatedSubscriptionOfferApplicationPhasesItemVariant3`
- `"fixed_price"` → `ReactivatedSubscriptionOfferApplicationPhasesItemVariant4`

### ReactivatedSubscriptionOfferApplicationPhasesItemVariant1

- `Type` (`string`, required)
- `DurationDays` (`int`, required)
- `StartsAt` (`string | null`, required)
- `EndsAt` (`string | null`, required)

### ReactivatedSubscriptionOfferApplicationPhasesItemVariant2

- `Type` (`string`, required)
- `DurationCycles` (`int | null`, required)
- `DurationInterval` (`string | null`, required)
- `StartsAt` (`string | null`, required)
- `EndsAt` (`string | null`, required)
- `Percentage` (`int`, required) — Discount in basis points. 5000 means 50%.

### ReactivatedSubscriptionOfferApplicationPhasesItemVariant3

- `Type` (`string`, required)
- `DurationCycles` (`int | null`, required)
- `DurationInterval` (`string | null`, required)
- `StartsAt` (`string | null`, required)
- `EndsAt` (`string | null`, required)
- `Amount` (`int`, required) — Discount in the application currency's minor unit.

### ReactivatedSubscriptionOfferApplicationPhasesItemVariant4

- `Type` (`string`, required)
- `DurationCycles` (`int | null`, required)
- `DurationInterval` (`string | null`, required)
- `StartsAt` (`string | null`, required)
- `EndsAt` (`string | null`, required)
- `Price` (`int`, required) — Fixed price in the application currency's minor unit.

### RecoveryLink

- `URL` (`string`, required)
- `Token` (`string`, required)
- `Object` (`string`, required)
- `Livemode` (`bool`, required)

### Refund

- `ID` (`string`, required)
- `TransactionID` (`string`, required)
- `Amount` (`int`, required)
- `Currency` (`string`, required)
- `ChargeID` (`string | null`, required)
- `Status` (`string`, required)
- `Reason` (`string | null`, required)
- `Object` (`string`, required)
- `Livemode` (`bool`, required)

### RemovedPlanFeature

- `ID` (`string`, required)
- `Removed` (`any`, required)
- `Object` (`string`, required)
- `Livemode` (`bool`, required)

### RemovedPlanFromGroup

- `ID` (`string`, required)
- `Removed` (`bool`, required)
- `Object` (`string`, required)
- `Livemode` (`bool`, required)

### ReorderedPlans

- `Reordered` (`bool`, required)
- `Object` (`string`, required)
- `Livemode` (`bool`, required)

### SeatBalance

- `Current` (`int`, required)
- `AsOf` (`string`, required)
- `Object` (`string`, required)
- `Livemode` (`bool`, required)

### SeatBalanceCollection

- `Balances` (`map[string]SeatBalanceCollectionBalancesValue`, required)
- `Object` (`string`, required)
- `Livemode` (`bool`, required)

### SeatBalanceCollectionBalancesValue

- `Current` (`int`, required)
- `AsOf` (`string`, required)

### SeatEvent

- `ID` (`string`, required)
- `CustomerID` (`string`, required)
- `FeatureCode` (`string`, required)
- `PreviousBalance` (`int`, required)
- `NewBalance` (`int`, required)
- `Ts` (`string`, required)
- `CreatedAt` (`string`, required)
- `Object` (`string`, required)
- `Livemode` (`bool`, required)

### SeatsSetAllResult

- `Object` (`string`, required)
- `Data` (`[]SeatEvent`, required)
- `HasMore` (`bool`, required)
- `NextCursor` (`string`, optional)

### SentInvoice

- `Sent` (`bool`, required)
- `SentAt` (`string`, required)
- `Object` (`string`, required)
- `Livemode` (`bool`, required)

### SetPlanRegionalPricingParamsFeaturesItem

- `FeatureID` (`string`, required)
- `OverageUnitPrice` (`int`, required)

### SetPlanRegionalPricingParamsPricesItem

- `PriceID` (`string`, required)
- `Price` (`int`, required)
- `IncludedBalance` (`int`, optional)

### SubscriberNeedingAChange

- `SubscriptionID` (`string`, required)
- `MoveID` (`string`, required) — The move it could not complete.
- `FromVersionID` (`string`, required) — The version it stays on.
- `ToVersionID` (`string`, required) — The version it could not move to.
- `Code` (`string`, required) — Why it could not move.
- `Message` (`string`, required)
- `Object` (`string`, required)
- `Livemode` (`bool`, required)

### Subscription

- `ID` (`string`, required)
- `CustomerID` (`string`, required)
- `Plan` (`SubscriptionPlan`, required)
- `Name` (`string`, required)
- `Description` (`string | null`, required)
- `Status` (`SubscriptionStatus`, required)
- `BillingInterval` (`BillingInterval | null`, required)
- `TrialEndsAt` (`string | null`, required)
- `CurrentPeriod` (`SubscriptionCurrentPeriod | null`, required)
- `Cancellation` (`SubscriptionCancellation | null`, required)
- `CancelAtPeriodEnd` (`bool`, required)
- `ScheduledPlanChange` (`SubscriptionScheduledPlanChange | null`, required)
- `StartDate` (`string`, required)
- `EndDate` (`string | null`, required)
- `BillingDayOfMonth` (`int | null`, required)
- `NextBillingDate` (`string | null`, required)
- `CheckoutURL` (`string | null`, required)
- `HasPaymentMethod` (`bool`, required) — Whether Commet can charge this subscription on its own at its next charge (the trial end or the renewal). A trial that started without a payment method reports false until the customer adds one for this subscription; a payment method saved for another subscription does not count. A subscription on a free plan reports false.
- `CreatedAt` (`string`, required)
- `UpdatedAt` (`string`, required)
- `PlanVersion` (`SubscriptionPlanVersion | null`, required) — The plan version whose terms the subscription has now. It keeps paying and receiving everything from this version until it is moved.
- `OfferApplications` (`[]SubscriptionOfferApplication`, required)
- `Pause` (`SubscriptionPause | null`, required)
- `ScheduledVersionChange` (`SubscriptionScheduledVersionChange | null`, required) — A move to another version of the plan pending for the next renewal.
- `PlanGrant` (`SubscriptionPlanGrant`, optional)
- `ConsumptionModel` (`ConsumptionModel | null`, required)
- `Features` (`[]SubscriptionFeaturesItem`, required)
- `Credits` (`SubscriptionCredits | null`, required)
- `Balance` (`SubscriptionBalance | null`, required)
- `PriceID` (`string | null`, required)
- `Object` (`string`, required)
- `Livemode` (`bool`, required)

### SubscriptionAddon

- `AddonID` (`string`, required)
- `Status` (`string`, required)
- `ProratedCharge` (`int`, required)
- `Object` (`string`, required)
- `Livemode` (`bool`, required)

### SubscriptionAutoTopUp

- `Enabled` (`bool`, required) — Whether top-ups run automatically: credits bought on a credits plan, balance added on a balance plan. Always false when the plan does not sell credits or has balance top-ups turned off.
- `Threshold` (`int | null`, required) — The available amount at or below which a top-up runs, or null when never configured: credits (plan plus purchased) on a credits plan, minor units of the subscription currency (plan plus purchased balance) on a balance plan.
- `Quantity` (`int | null`, required) — Credits each top-up buys on a credits plan, or null when never configured or on a balance plan.
- `Amount` (`int | null`, required) — Minor units of the subscription currency each top-up adds on a balance plan, charged plus tax, or null when never configured or on a credits plan.
- `MaxPerPeriod` (`int | null`, required) — Maximum top-ups per renewal period, or null when never configured.
- `Status` (`string`, required) — active tops up when the available amount reaches the threshold; paused stops after three consecutive failed top-ups, or after one whose charge could not be refunded, until the payment method changes, a credit purchase or balance top-up succeeds or it is saved enabled again; off tops up nothing.
- `TopUpsThisPeriod` (`int`, required) — Top-ups made in the current renewal period.
- `PeriodResetsAt` (`string | null`, required) — When the top-up count of the period resets, or null before the subscription starts its first period.
- `LastFailure` (`SubscriptionAutoTopUpLastFailure | null`, required) — The last top-up that failed, or null when none has failed.
- `Object` (`string`, required)
- `Livemode` (`bool`, required)

### SubscriptionAutoTopUpLastFailure

- `Reason` (`string`, required)
- `At` (`string`, required)

### SubscriptionBalance

- `Remaining` (`float64`, required) — Balance available in minor units of the subscription currency. Kept for compatibility; equal to `total`.
- `Included` (`float64`, required)
- `Currency` (`string`, required)
- `Plan` (`float64`, required) — Plan balance left this period, in minor units. Spent first and replaced or rolled over at each reset; negative while overage is owed.
- `Purchased` (`float64`, required) — Purchased balance in minor units: top-ups and manual adjustments. Spent after the plan balance and kept until the subscription ends.
- `Total` (`float64`, required) — Balance available in minor units: `plan` plus `purchased`.

### SubscriptionCancellation

- `ScheduledAt` (`string`, required)
- `Reason` (`string | null`, required)
- `EffectiveAt` (`string`, required)

### SubscriptionCredits

- `Remaining` (`float64`, required) — Plan credits left. Kept for compatibility; equal to `plan`.
- `Included` (`float64`, required)
- `Purchased` (`float64`, required) — Purchased credits: purchases and manual adjustments. Spent after plan credits and kept until the subscription ends.
- `Plan` (`float64`, required) — Plan credits left this period. Spent first and replaced or rolled over at each reset; negative while overage is owed.
- `Total` (`float64`, required) — Credits available: `plan` plus `purchased`.

### SubscriptionCurrentPeriod

- `Start` (`string`, required)
- `End` (`string`, required)
- `DaysRemaining` (`float64`, required)

### SubscriptionFeaturesItem

Variants:

- `SubscriptionFeaturesItemVariant1`
- `SubscriptionFeaturesItemVariant2`
- `SubscriptionFeaturesItemVariant3`
- `SubscriptionFeaturesItemVariant4`

Discriminator: `Type`

- `"boolean"` → `SubscriptionFeaturesItemVariant1`
- `"usage"` → `SubscriptionFeaturesItemVariant2`
- `"seats"` → `SubscriptionFeaturesItemVariant3`
- `"quota"` → `SubscriptionFeaturesItemVariant4`

### SubscriptionFeaturesItemVariant1

- `Code` (`string`, required)
- `Name` (`string`, required)
- `Type` (`string`, required)
- `Enabled` (`bool`, required)
- `BaseAccess` (`SubscriptionFeaturesItemVariant1BaseAccess | null`, optional)

### SubscriptionFeaturesItemVariant1BaseAccess

- `Enabled` (`bool`, required)

### SubscriptionFeaturesItemVariant2

- `Code` (`string`, required)
- `Name` (`string`, required)
- `Type` (`string`, required)
- `Usage` (`SubscriptionFeaturesItemVariant2Usage`, optional)
- `BaseAccess` (`SubscriptionFeaturesItemVariant2BaseAccess | null`, optional)

### SubscriptionFeaturesItemVariant2BaseAccess

- `Included` (`float64`, required)
- `Unlimited` (`bool`, required)

### SubscriptionFeaturesItemVariant2Usage

- `Current` (`float64`, required)
- `Included` (`float64`, required)
- `OverageQuantity` (`float64`, required)
- `OverageUnitPrice` (`float64`, optional)
- `Unlimited` (`bool`, optional)

### SubscriptionFeaturesItemVariant3

- `Code` (`string`, required)
- `Name` (`string`, required)
- `Type` (`string`, required)
- `Usage` (`SubscriptionFeaturesItemVariant3Usage`, required)
- `BaseAccess` (`SubscriptionFeaturesItemVariant3BaseAccess | null`, optional)

### SubscriptionFeaturesItemVariant3BaseAccess

- `Included` (`float64`, required)
- `Unlimited` (`bool`, required)

### SubscriptionFeaturesItemVariant3Usage

- `Current` (`float64`, required)
- `Included` (`float64`, required)
- `OverageQuantity` (`float64`, required)
- `OverageUnitPrice` (`float64`, optional)
- `Unlimited` (`bool`, optional)

### SubscriptionFeaturesItemVariant4

- `Code` (`string`, required)
- `Name` (`string`, required)
- `Type` (`string`, required)
- `Usage` (`SubscriptionFeaturesItemVariant4Usage`, optional)
- `BaseAccess` (`SubscriptionFeaturesItemVariant4BaseAccess | null`, optional)

### SubscriptionFeaturesItemVariant4BaseAccess

- `Included` (`float64`, required)
- `Unlimited` (`bool`, required)

### SubscriptionFeaturesItemVariant4Usage

- `Current` (`float64`, required)
- `Included` (`float64`, required)
- `OverageQuantity` (`float64`, required)
- `OverageUnitPrice` (`float64`, optional)
- `Unlimited` (`bool`, optional)

### SubscriptionOfferApplication

- `ID` (`string`, required)
- `Name` (`string`, required)
- `AppliesTo` (`SubscriptionOfferApplicationAppliesTo`, required)
- `OfferID` (`string | null`, required)
- `Source` (`string`, required)
- `Status` (`string`, required)
- `Currency` (`string | null`, required)
- `Subtotal` (`int | null`, required)
- `DiscountAmount` (`int | null`, required)
- `Total` (`int | null`, required)
- `Phases` (`[]SubscriptionOfferApplicationPhase`, required)
- `QuotedAt` (`string`, required)
- `ExpiresAt` (`string | null`, required)
- `AppliedAt` (`string | null`, required)

### SubscriptionOfferApplicationAppliesTo

Variants:

- `SubscriptionOfferApplicationAppliesToVariant1`
- `SubscriptionOfferApplicationAppliesToVariant2`
- `SubscriptionOfferApplicationAppliesToVariant3`

Discriminator: `Type`

- `"plan_price"` → `SubscriptionOfferApplicationAppliesToVariant1`
- `"addon"` → `SubscriptionOfferApplicationAppliesToVariant2`
- `"credit_pack"` → `SubscriptionOfferApplicationAppliesToVariant3`

### SubscriptionOfferApplicationAppliesToVariant1

- `Type` (`string`, required)
- `ID` (`string`, required)

### SubscriptionOfferApplicationAppliesToVariant2

- `Type` (`string`, required)
- `ID` (`string`, required)

### SubscriptionOfferApplicationAppliesToVariant3

- `Type` (`string`, required)
- `ID` (`string`, required)

### SubscriptionOfferApplicationPhase

Variants:

- `SubscriptionOfferApplicationPhaseVariant1`
- `SubscriptionOfferApplicationPhaseVariant2`
- `SubscriptionOfferApplicationPhaseVariant3`
- `SubscriptionOfferApplicationPhaseVariant4`

Discriminator: `Type`

- `"free_trial"` → `SubscriptionOfferApplicationPhaseVariant1`
- `"percentage"` → `SubscriptionOfferApplicationPhaseVariant2`
- `"amount_off"` → `SubscriptionOfferApplicationPhaseVariant3`
- `"fixed_price"` → `SubscriptionOfferApplicationPhaseVariant4`

### SubscriptionOfferApplicationPhaseVariant1

- `Type` (`string`, required)
- `DurationDays` (`int`, required)
- `DurationInterval` (`string | null`, required)
- `StartsAt` (`string | null`, required)
- `EndsAt` (`string | null`, required)

### SubscriptionOfferApplicationPhaseVariant2

- `Type` (`string`, required)
- `DurationCycles` (`int | null`, required)
- `DurationInterval` (`string | null`, required)
- `Percentage` (`int`, required)
- `StartsAt` (`string | null`, required)
- `EndsAt` (`string | null`, required)

### SubscriptionOfferApplicationPhaseVariant3

- `Type` (`string`, required)
- `DurationCycles` (`int | null`, required)
- `DurationInterval` (`string | null`, required)
- `Amount` (`int`, required)
- `StartsAt` (`string | null`, required)
- `EndsAt` (`string | null`, required)

### SubscriptionOfferApplicationPhaseVariant4

- `Type` (`string`, required)
- `DurationCycles` (`int | null`, required)
- `DurationInterval` (`string | null`, required)
- `Price` (`int`, required)
- `StartsAt` (`string | null`, required)
- `EndsAt` (`string | null`, required)

### SubscriptionPause

Variants:

- `SubscriptionPauseVariant1`
- `SubscriptionPauseVariant2`

Discriminator: `Status`

- `"scheduled"` → `SubscriptionPauseVariant1`
- `"active"` → `SubscriptionPauseVariant2`

### SubscriptionPauseVariant1

- `Status` (`string`, required)
- `Mode` (`string`, required)
- `RequestedAt` (`string`, required)
- `EffectiveAt` (`string`, required)
- `ResumeAt` (`string | null`, required)

### SubscriptionPauseVariant2

- `Status` (`string`, required)
- `Mode` (`string`, required)
- `RequestedAt` (`string`, required)
- `EffectiveAt` (`string`, required)
- `ResumeAt` (`string | null`, required)

### SubscriptionPlan

- `ID` (`string`, required)
- `Name` (`string`, required)
- `BasePrice` (`float64`, required)

### SubscriptionPlanGrant

- `ID` (`string`, required) — The active Plan Grant ID.
- `Plan` (`SubscriptionPlanGrantPlan`, required) — The higher plan whose access is temporarily applied.
- `ExpiresAt` (`string | null`, required) — When the temporary access ends, or null when it lasts until revoked.

### SubscriptionPlanGrantPlan

- `ID` (`string`, required)
- `Name` (`string`, required)

### SubscriptionPlanVersion

- `ID` (`string`, required) — Public plan version ID.
- `Label` (`string`, required) — Main versions are 1, 2, 3; an experiment's versions carry its own code, like price-test 2.

### SubscriptionResume

- `SubscriptionID` (`string`, required)
- `InvoiceID` (`string | null`, required)
- `Status` (`string`, required)
- `Object` (`string`, required)
- `Livemode` (`bool`, required)

### SubscriptionScheduledPlanChange

- `ChangeType` (`string`, required)
- `NewPlanID` (`string | null`, required)
- `NewPlanName` (`string | null`, required)
- `NewBillingInterval` (`string | null`, required)
- `ScheduledFor` (`string`, required)

### SubscriptionScheduledVersionChange

- `PlanVersionID` (`string`, required)
- `Label` (`string`, required)
- `EffectiveAt` (`string`, required)

### SubscriptionsListResult

- `Object` (`string`, required)
- `Data` (`[]SubscriptionSummary`, required)
- `HasMore` (`bool`, required)
- `NextCursor` (`string`, optional)

### SubscriptionSummary

- `ID` (`string`, required)
- `CustomerID` (`string`, required)
- `Plan` (`SubscriptionSummaryPlan`, required)
- `Name` (`string`, required)
- `Description` (`string | null`, required)
- `Status` (`SubscriptionStatus`, required)
- `BillingInterval` (`BillingInterval | null`, required)
- `TrialEndsAt` (`string | null`, required)
- `CurrentPeriod` (`SubscriptionSummaryCurrentPeriod | null`, required)
- `Cancellation` (`SubscriptionSummaryCancellation | null`, required)
- `CancelAtPeriodEnd` (`bool`, required)
- `ScheduledPlanChange` (`SubscriptionSummaryScheduledPlanChange | null`, required)
- `StartDate` (`string`, required)
- `EndDate` (`string | null`, required)
- `BillingDayOfMonth` (`int | null`, required)
- `NextBillingDate` (`string | null`, required)
- `CheckoutURL` (`string | null`, required)
- `HasPaymentMethod` (`bool`, required) — Whether Commet can charge this subscription on its own at its next charge (the trial end or the renewal). A trial that started without a payment method reports false until the customer adds one for this subscription; a payment method saved for another subscription does not count. A subscription on a free plan reports false.
- `CreatedAt` (`string`, required)
- `UpdatedAt` (`string`, required)
- `PlanVersion` (`SubscriptionSummaryPlanVersion | null`, required) — The plan version whose terms the subscription has now. It keeps paying and receiving everything from this version until it is moved.
- `OfferApplications` (`[]SubscriptionOfferApplication`, required)
- `Pause` (`SubscriptionSummaryPause | null`, required)
- `PriceID` (`string | null`, required)
- `Object` (`string`, required)
- `Livemode` (`bool`, required)

### SubscriptionSummaryCancellation

- `ScheduledAt` (`string`, required)
- `Reason` (`string | null`, required)
- `EffectiveAt` (`string`, required)

### SubscriptionSummaryCurrentPeriod

- `Start` (`string`, required)
- `End` (`string`, required)
- `DaysRemaining` (`float64`, required)

### SubscriptionSummaryPause

Variants:

- `SubscriptionSummaryPauseVariant1`
- `SubscriptionSummaryPauseVariant2`

Discriminator: `Status`

- `"scheduled"` → `SubscriptionSummaryPauseVariant1`
- `"active"` → `SubscriptionSummaryPauseVariant2`

### SubscriptionSummaryPauseVariant1

- `Status` (`string`, required)
- `Mode` (`string`, required)
- `RequestedAt` (`string`, required)
- `EffectiveAt` (`string`, required)
- `ResumeAt` (`string | null`, required)

### SubscriptionSummaryPauseVariant2

- `Status` (`string`, required)
- `Mode` (`string`, required)
- `RequestedAt` (`string`, required)
- `EffectiveAt` (`string`, required)
- `ResumeAt` (`string | null`, required)

### SubscriptionSummaryPlan

- `ID` (`string`, required)
- `Name` (`string`, required)

### SubscriptionSummaryPlanVersion

- `ID` (`string`, required) — Public plan version ID.
- `Label` (`string`, required) — Main versions are 1, 2, 3; an experiment's versions carry its own code, like price-test 2.

### SubscriptionSummaryScheduledPlanChange

- `ChangeType` (`string`, required)
- `NewPlanID` (`string | null`, required)
- `NewPlanName` (`string | null`, required)
- `NewBillingInterval` (`string | null`, required)
- `ScheduledFor` (`string`, required)

### TestClock

- `SimulatedTime` (`string | null`, required)
- `IsActive` (`bool`, required)
- `Now` (`string`, required)
- `LatestRun` (`TestClockLatestRun | null`, required)
- `Object` (`string`, required)
- `Livemode` (`bool`, required)

### TestClockLatestRun

- `ID` (`string`, required)
- `Status` (`string`, required)
- `StartedAtTime` (`string`, required)
- `TargetTime` (`string`, required)
- `EstimatedDeadlineCount` (`int`, required)
- `CompletedDeadlineCount` (`int`, required)
- `FailedDeadlineCount` (`int`, required)
- `Error` (`string | null`, required)
- `Items` (`[]TestClockLatestRunItemsItem`, required)

### TestClockLatestRunItemsItem

- `Kind` (`string`, required)
- `Status` (`string`, required)
- `DueAt` (`string`, required)
- `SubscriptionID` (`string`, required)
- `CustomerName` (`string | null`, required)
- `InvoiceNumber` (`string | null`, required)
- `InvoiceID` (`string | null`, required)
- `Outcome` (`string | null`, required)
- `Detail` (`string | null`, required)
- `Error` (`string | null`, required)

### TestClockRun

- `ID` (`string`, required)
- `Status` (`string`, required)
- `StartedAtTime` (`string`, required)
- `TargetTime` (`string`, required)
- `EstimatedDeadlineCount` (`int`, required)
- `CompletedDeadlineCount` (`int`, required)
- `FailedDeadlineCount` (`int`, required)
- `Error` (`string | null`, required)
- `Items` (`[]TestClockRunItemsItem`, required)
- `Object` (`string`, required)
- `Livemode` (`bool`, required)

### TestClockRunItemsItem

- `Kind` (`string`, required)
- `Status` (`string`, required)
- `DueAt` (`string`, required)
- `SubscriptionID` (`string`, required)
- `CustomerName` (`string | null`, required)
- `InvoiceNumber` (`string | null`, required)
- `InvoiceID` (`string | null`, required)
- `Outcome` (`string | null`, required)
- `Detail` (`string | null`, required)
- `Error` (`string | null`, required)

### TrackUsageParamsPropertiesItem

- `Property` (`string`, required)
- `Value` (`string`, required)

### Transaction

- `PaymentContext` (`TransactionPaymentContext | null`, required) — Charge context captured for new payments. Null for historical payments with no captured context.
- `ID` (`string`, required)
- `InvoiceID` (`string | null`, required)
- `GrossAmount` (`int | null`, required) — Gross amount in USD cents. Null when the provider has not reported an honest USD figure; see presentmentAmount.
- `Subtotal` (`int | null`, required) — Subtotal in USD cents (gross minus tax). Null when grossAmount is null.
- `TaxAmount` (`int | null`, required)
- `PresentmentAmount` (`int | null`, required) — Amount in the charge currency's smallest unit, as presented to the customer. Set for non-USD charges; null when the charge was made in USD.
- `AmountRefunded` (`int`, required) — Cumulative confirmed refunds in the payment currency’s smallest unit.
- `Currency` (`string`, required)
- `Provider` (`PaymentProvider`, required) — The payment provider the charge was routed to: stripe, commet, or dlocal.
- `PaymentMethod` (`PaymentMethod | null`, required) — The method used for this charge: card, oxxo, or mercado_pago. Null when unknown.
- `SubPaymentMethod` (`SubPaymentMethod | null`, required) — The source of funds for this charge, when reported by the provider. Null when unavailable or unknown.
- `Status` (`TransactionStatus`, required)
- `CustomerEmail` (`string | null`, required)
- `CustomerName` (`string | null`, required)
- `PaidAt` (`string | null`, required)
- `CreatedAt` (`string`, required)
- `UpdatedAt` (`string`, required)
- `AvailableAt` (`string | null`, required)
- `Object` (`string`, required)
- `Livemode` (`bool`, required)

### TransactionListItem

- `PaymentContext` (`TransactionListItemPaymentContext | null`, required) — Charge context captured for new payments. Null for historical payments with no captured context.
- `ID` (`string`, required)
- `InvoiceID` (`string | null`, required)
- `GrossAmount` (`int | null`, required) — Gross amount in USD cents. Null when the provider has not reported an honest USD figure; see presentmentAmount.
- `Subtotal` (`int | null`, required) — Subtotal in USD cents (gross minus tax). Null when grossAmount is null.
- `TaxAmount` (`int | null`, required)
- `PresentmentAmount` (`int | null`, required) — Amount in the charge currency's smallest unit, as presented to the customer. Set for non-USD charges; null when the charge was made in USD.
- `AmountRefunded` (`int`, required) — Cumulative confirmed refunds in the payment currency’s smallest unit.
- `Currency` (`string`, required)
- `Provider` (`PaymentProvider`, required) — The payment provider the charge was routed to: stripe, commet, or dlocal.
- `PaymentMethod` (`PaymentMethod | null`, required) — The method used for this charge: card, oxxo, or mercado_pago. Null when unknown.
- `SubPaymentMethod` (`SubPaymentMethod | null`, required) — The source of funds for this charge, when reported by the provider. Null when unavailable or unknown.
- `Status` (`TransactionStatus`, required)
- `CustomerEmail` (`string | null`, required)
- `CustomerName` (`string | null`, required)
- `PaidAt` (`string | null`, required)
- `CreatedAt` (`string`, required)
- `UpdatedAt` (`string`, required)
- `Object` (`string`, required)
- `Livemode` (`bool`, required)

### TransactionListItemPaymentContext

- `Reason` (`string`, required) — The original reason for the charge. Recovery never replaces this reason.
- `PaymentLinkID` (`string | null`, required) — The public payment link ID, independently of the reason, or null when no payment link originated the charge.
- `Recovery` (`TransactionListItemPaymentContextRecovery | null`, required)

### TransactionListItemPaymentContextRecovery

Variants:

- `TransactionListItemPaymentContextRecoveryVariant1`
- `TransactionListItemPaymentContextRecoveryVariant2`

Discriminator: `Type`

- `"payment_recovery"` → `TransactionListItemPaymentContextRecoveryVariant1`
- `"dunning_retry"` → `TransactionListItemPaymentContextRecoveryVariant2`

### TransactionListItemPaymentContextRecoveryVariant1

- `Type` (`string`, required)

### TransactionListItemPaymentContextRecoveryVariant2

- `Type` (`string`, required)
- `Attempt` (`int`, required) — Current retry, starting at 1. The original decline is not a retry.
- `MaxAttempts` (`int`, required) — Total retries applicable to this charge's dunning schedule.

### TransactionPaymentContext

- `Reason` (`string`, required) — The original reason for the charge. Recovery never replaces this reason.
- `PaymentLinkID` (`string | null`, required) — The public payment link ID, independently of the reason, or null when no payment link originated the charge.
- `Recovery` (`TransactionPaymentContextRecovery | null`, required)

### TransactionPaymentContextRecovery

Variants:

- `TransactionPaymentContextRecoveryVariant1`
- `TransactionPaymentContextRecoveryVariant2`

Discriminator: `Type`

- `"payment_recovery"` → `TransactionPaymentContextRecoveryVariant1`
- `"dunning_retry"` → `TransactionPaymentContextRecoveryVariant2`

### TransactionPaymentContextRecoveryVariant1

- `Type` (`string`, required)

### TransactionPaymentContextRecoveryVariant2

- `Type` (`string`, required)
- `Attempt` (`int`, required) — Current retry, starting at 1. The original decline is not a retry.
- `MaxAttempts` (`int`, required) — Total retries applicable to this charge's dunning schedule.

### TransactionRetry

- `OriginalTransactionID` (`string`, required)
- `InvoiceID` (`string`, required)
- `Status` (`string`, required)
- `Object` (`string`, required)
- `Livemode` (`bool`, required)

### TransactionsListResult

- `Object` (`string`, required)
- `Data` (`[]TransactionListItem`, required)
- `HasMore` (`bool`, required)
- `NextCursor` (`string`, optional)

### UpdateCustomerParamsAddress

- `Line1` (`string`, required)
- `Line2` (`string`, optional)
- `City` (`string`, required)
- `State` (`string`, optional)
- `PostalCode` (`string`, required)
- `Country` (`string`, required)
- `Region` (`string`, optional)

### UpdateOfferParamsPhasesItem

Variants:

- `UpdateOfferParamsPhasesItemVariant1`
- `UpdateOfferParamsPhasesItemVariant2`
- `UpdateOfferParamsPhasesItemVariant3`

Discriminator: `Type`

- `"percentage"` → `UpdateOfferParamsPhasesItemVariant1`
- `"amount_off"` → `UpdateOfferParamsPhasesItemVariant2`
- `"fixed_price"` → `UpdateOfferParamsPhasesItemVariant3`

### UpdateOfferParamsPhasesItemVariant1

- `Type` (`string`, required)
- `DurationCycles` (`int | null`, required)
- `DurationInterval` (`string | null`, optional) — Unit the phase duration is counted in. Only a fixed-price phase may set it, because its amount is declared rather than derived from the plan. Defaults to the plan's own billing interval.
- `Percentage` (`int`, required) — Discount in basis points. 5000 means 50%.

### UpdateOfferParamsPhasesItemVariant2

- `Type` (`string`, required)
- `DurationCycles` (`int | null`, required)
- `DurationInterval` (`string | null`, optional) — Unit the phase duration is counted in. Only a fixed-price phase may set it, because its amount is declared rather than derived from the plan. Defaults to the plan's own billing interval.
- `Amounts` (`[]UpdateOfferParamsPhasesItemVariant2AmountsItem`, required)

### UpdateOfferParamsPhasesItemVariant2AmountsItem

- `Currency` (`string`, required)
- `Amount` (`int`, required) — Amount in the currency's minor unit (for example, cents for USD).

### UpdateOfferParamsPhasesItemVariant3

- `Type` (`string`, required)
- `DurationCycles` (`int | null`, required)
- `DurationInterval` (`string | null`, optional) — Unit the phase duration is counted in. Only a fixed-price phase may set it, because its amount is declared rather than derived from the plan. Defaults to the plan's own billing interval.
- `Prices` (`[]UpdateOfferParamsPhasesItemVariant3PricesItem`, required)

### UpdateOfferParamsPhasesItemVariant3PricesItem

- `Currency` (`string`, required)
- `Amount` (`int`, required) — Amount in the currency's minor unit (for example, cents for USD).

### UpdatePlanDraftParamsCountriesItem

- `Country` (`string`, required) — ISO country code.
- `Currency` (`string`, required) — The currency the country pays in, USD included.
- `ExchangeRate` (`float64 | null`, optional) — Local units per USD for prices that follow the USD base. USD defaults to 1.
- `Prices` (`[]UpdatePlanDraftParamsCountriesItemPricesItem`, required)
- `TopUpMinimum` (`int | null`, optional) — Smallest balance top-up in this country, in the currency's minor unit. Null follows the USD minimum at the exchange rate; left out keeps it.
- `MaxPlanBalance` (`int | null`, optional) — Most unused plan balance that rolls over in this country. Null follows the USD cap at the exchange rate; left out keeps it.
- `PlanBalanceLimit` (`int | null`, optional) — Most plan balance spent per limit interval in this country. Null follows the USD limit at the exchange rate; left out keeps it.
- `AddonPrices` (`[]UpdatePlanDraftParamsCountriesItemAddonPricesItem`, optional) — What the country pays for each add-on the draft sells. Left out keeps the country's add-on prices; sent, it replaces them, and add-ons not listed follow the USD price.
- `CreditPackPrices` (`[]UpdatePlanDraftParamsCountriesItemCreditPackPricesItem`, optional) — What the country pays for each credit pack the draft sells. Left out keeps the country's pack prices; sent, it replaces them, and packs not listed follow the USD price.

### UpdatePlanDraftParamsCountriesItemAddonPricesItem

Variants:

- `UpdatePlanDraftParamsCountriesItemAddonPricesItemVariant1`
- `UpdatePlanDraftParamsCountriesItemAddonPricesItemVariant2`

### UpdatePlanDraftParamsCountriesItemAddonPricesItemVariant1

- `AddonID` (`string`, required)
- `Price` (`int`, required) — The add-on's price set by hand, in the currency's minor unit.

### UpdatePlanDraftParamsCountriesItemAddonPricesItemVariant2

- `AddonID` (`string`, required)
- `FollowBase` (`any`, required) — Follow the add-on's USD price at the exchange rate.

### UpdatePlanDraftParamsCountriesItemCreditPackPricesItem

Variants:

- `UpdatePlanDraftParamsCountriesItemCreditPackPricesItemVariant1`
- `UpdatePlanDraftParamsCountriesItemCreditPackPricesItemVariant2`

### UpdatePlanDraftParamsCountriesItemCreditPackPricesItemVariant1

- `CreditPackID` (`string`, required)
- `Price` (`int`, required) — The pack's price set by hand, in the currency's minor unit.

### UpdatePlanDraftParamsCountriesItemCreditPackPricesItemVariant2

- `CreditPackID` (`string`, required)
- `FollowBase` (`any`, required) — Follow the pack's USD price at the exchange rate.

### UpdatePlanDraftParamsCountriesItemPricesItem

Variants:

- `UpdatePlanDraftParamsCountriesItemPricesItemVariant1`
- `UpdatePlanDraftParamsCountriesItemPricesItemVariant2`

### UpdatePlanDraftParamsCountriesItemPricesItemVariant1

- `BillingInterval` (`string`, required)
- `Price` (`int`, required) — The price set by hand, in the currency's minor unit.
- `IncludedBalance` (`int`, optional) — The balance included each period in this country, in the currency's minor unit. Required when the plan's price includes a balance.

### UpdatePlanDraftParamsCountriesItemPricesItemVariant2

- `BillingInterval` (`string`, required)
- `FollowBase` (`any`, required) — Follow the USD base at the exchange rate.

### UpdatePlanDraftParamsPricesItem

- `BillingInterval` (`string`, required)
- `Price` (`int`, required) — USD price in cents.
- `IncludedCredits` (`int | null`, optional)
- `IncludedBalance` (`int | null`, optional)
- `TrialDays` (`int | null`, optional) — Days of trial on this price. Null removes the trial.
- `TrialRequiresPaymentMethod` (`bool`, optional)
- `IsDefault` (`bool`, optional) — Show this interval first at checkout and in pricing tables.

### UpdatePlanDraftParamsRules

- `FreeIncludedCredits` (`int | null`, optional) — Credits a free plan includes each month.
- `FreeIncludedBalance` (`int | null`, optional) — Balance a free plan includes each month, in cents.
- `UnusedPlanCredits` (`string`, optional)
- `MaxPlanCredits` (`int | null`, optional)
- `PlanCreditLimit` (`int | null`, optional)
- `PlanCreditLimitInterval` (`string | null`, optional)
- `UnusedPlanBalance` (`string`, optional)
- `MaxPlanBalance` (`int | null`, optional)
- `PlanBalanceLimit` (`int | null`, optional)
- `PlanBalanceLimitInterval` (`string | null`, optional)
- `BlockOnExhaustion` (`bool`, optional)
- `PurchasesEnabled` (`bool`, optional)
- `CreditPurchaseUnitPrice` (`int | null`, optional)
- `CreditPurchaseMinimum` (`int | null`, optional)
- `TopUpMinimum` (`int | null`, optional)

### UpdatePlanFeatureParamsOverage

- `Enabled` (`bool`, optional)
- `UnitPrice` (`int`, optional)

### UpdatePlanParamsPlanBalanceLimit

- `Amount` (`int`, required) — The most plan balance a subscription spends in one window, in USD cents. Other currencies convert it with the plan's exchange rate unless the dashboard sets a custom amount for that currency.
- `Interval` (`string`, required) — The window length: `day` is 24 hours and `week` is 7 days, starting at the time of day of the subscription's billing anchor.

### UpdatePlanParamsPlanCreditLimit

- `Credits` (`int`, required) — The most plan credits a subscription spends in one window.
- `Interval` (`string`, required) — The window length: `day` is 24 hours and `week` is 7 days, starting at the time of day of the subscription's billing anchor.

### UpdatePlanParamsUnusedPlanBalance

Variants:

- `UpdatePlanParamsUnusedPlanBalanceVariant1`
- `UpdatePlanParamsUnusedPlanBalanceVariant2`

Discriminator: `Type`

- `"expire"` → `UpdatePlanParamsUnusedPlanBalanceVariant1`
- `"roll_over"` → `UpdatePlanParamsUnusedPlanBalanceVariant2`

### UpdatePlanParamsUnusedPlanBalanceVariant1

- `Type` (`string`, required)

### UpdatePlanParamsUnusedPlanBalanceVariant2

- `Type` (`string`, required)
- `MaxPlanBalance` (`int | null`, required) — The most plan balance a subscription holds after a reset, in USD cents, including the balance that reset grants; balance above it expires. At least the included balance of every price. Other currencies convert it with the plan's exchange rate unless the dashboard sets a custom amount for that currency. null means no limit other than the largest storable amount, 2,147,483,647 cents.

### UpdatePlanParamsUnusedPlanCredits

Variants:

- `UpdatePlanParamsUnusedPlanCreditsVariant1`
- `UpdatePlanParamsUnusedPlanCreditsVariant2`

Discriminator: `Type`

- `"expire"` → `UpdatePlanParamsUnusedPlanCreditsVariant1`
- `"roll_over"` → `UpdatePlanParamsUnusedPlanCreditsVariant2`

### UpdatePlanParamsUnusedPlanCreditsVariant1

- `Type` (`string`, required)

### UpdatePlanParamsUnusedPlanCreditsVariant2

- `Type` (`string`, required)
- `MaxPlanCredits` (`int | null`, required) — The most plan credits a subscription holds after a reset, including the credits that reset grants; credits above it expire. At least the included credits of every price. null means no limit other than the largest storable balance, 2,147,483,647 plan credits.

### UpdatePlanPriceParamsMarketPricesItem

- `MarketGroupID` (`string`, required)
- `Currency` (`string`, required)
- `Price` (`int`, required)

### UpsertRegionalPricesParamsOverridesItem

- `Currency` (`string`, required)
- `Price` (`int`, required)
- `IncludedBalance` (`int`, optional)

### UsageAdjustment

- `ID` (`string`, required)
- `Value` (`int`, required)
- `PreviousValue` (`int`, required)
- `Adjustment` (`int`, required)
- `CustomerID` (`string`, required)
- `Reason` (`string | null`, required)
- `Ts` (`string`, required)
- `CreatedAt` (`string`, required)
- `FeatureCode` (`string`, required)
- `Object` (`string`, required)
- `Livemode` (`bool`, required)

### UsageCheck

Variants:

- `UsageCheckVariant1`
- `UsageCheckVariant2`
- `UsageCheckVariant3`

Discriminator: `ConsumptionModel`

- `"metered"` → `UsageCheckVariant1`
- `"credits"` → `UsageCheckVariant2`
- `"balance"` → `UsageCheckVariant3`

### UsageCheckVariant1

- `Allowed` (`bool`, required)
- `SubscriptionStatus` (`string`, required)
- `FeatureCode` (`string`, required)
- `Quantity` (`int`, required)
- `Reason` (`string`, optional)
- `Message` (`string`, optional)
- `ConsumptionModel` (`string`, required)
- `Current` (`float64`, required)
- `Remaining` (`float64`, required)
- `Unlimited` (`bool`, required)
- `Included` (`float64`, required)
- `OverageEnabled` (`bool`, required)
- `OverageUnitPrice` (`float64 | null`, required)
- `Object` (`string`, required)
- `Livemode` (`bool`, required)

### UsageCheckVariant2

- `Allowed` (`bool`, required)
- `SubscriptionStatus` (`string`, required)
- `FeatureCode` (`string`, required)
- `Quantity` (`int`, required)
- `Reason` (`string`, optional)
- `Message` (`string`, optional)
- `ConsumptionModel` (`string`, required)
- `CreditsPerUnit` (`int`, required)
- `EstimatedCredits` (`int`, required)
- `PlanCredits` (`int`, required)
- `PurchasedCredits` (`int`, required)
- `TotalCredits` (`int`, required)
- `PlanCreditLimit` (`UsageCheckVariant2PlanCreditLimit | null`, required)
- `BlockOnExhaustion` (`bool`, required) — Whether usage stops when plan and purchased credits run out. False only when the plan allows overage, which is billed at its price per credit; trials and free plans always stop.
- `Object` (`string`, required)
- `Livemode` (`bool`, required)

### UsageCheckVariant2PlanCreditLimit

- `Credits` (`int`, required) — The most plan credits the subscription spends per window.
- `Interval` (`string`, required) — The window length: 24 hours for `day`, 7 days for `week`.
- `Spent` (`int`, required) — Plan credits spent in the current window.
- `Remaining` (`int`, required) — Plan credits the subscription can still spend in the current window. Purchased credits are not limited.
- `ResetsAt` (`string`, required) — When the next window starts and spending resets.

### UsageCheckVariant3

- `Allowed` (`bool`, required)
- `SubscriptionStatus` (`string`, required)
- `FeatureCode` (`string`, required)
- `Quantity` (`int`, required)
- `Reason` (`string`, optional)
- `Message` (`string`, optional)
- `ConsumptionModel` (`string`, required)
- `UnitPrice` (`float64`, required)
- `EstimatedAmount` (`float64`, required)
- `CurrentBalance` (`float64`, required) — The available balance, plan plus purchased, in rate scale (10000 = 1.00 of the subscription currency). Equal to totalBalance.
- `PlanBalance` (`float64`, required) — The plan part of the balance in rate scale. Consumed first, replaced or rolled over at each reset, and negative while overage is owed.
- `PurchasedBalance` (`float64`, required) — The purchased part of the balance in rate scale: top-ups and adjustments. Consumed after the plan part and kept until the subscription ends.
- `TotalBalance` (`float64`, required) — planBalance plus purchasedBalance, in rate scale.
- `PlanBalanceLimit` (`UsageCheckVariant3PlanBalanceLimit | null`, required)
- `BlockOnExhaustion` (`bool`, required)
- `Currency` (`string`, required)
- `Object` (`string`, required)
- `Livemode` (`bool`, required)

### UsageCheckVariant3PlanBalanceLimit

- `Amount` (`int`, required) — The most plan balance the subscription spends per window, in rate scale (10000 = 1.00 of the subscription currency).
- `Interval` (`string`, required) — The window length: 24 hours for `day`, 7 days for `week`.
- `Spent` (`int`, required) — Plan balance spent in the current window, in rate scale.
- `Remaining` (`int`, required) — Plan balance the subscription can still spend in the current window, in rate scale. Purchased balance is not limited.
- `ResetsAt` (`string`, required) — When the next window starts and spending resets.

### UsageEvent

- `ID` (`string`, required)
- `FeatureCode` (`string`, required)
- `Value` (`float64`, required)
- `CustomerID` (`string`, required)
- `EventID` (`string | null`, required)
- `Ts` (`string`, required)
- `CreatedAt` (`string`, required)
- `Properties` (`[]UsageEventPropertiesItem`, required)
- `Consumption` (`UsageEventConsumption`, optional)
- `Object` (`string`, required)
- `Livemode` (`bool`, required)

### UsageEventConsumption

- `Model` (`string`, required)
- `Deducted` (`float64`, required)
- `Remaining` (`float64`, required)
- `Blocked` (`bool`, required)
- `PlanCreditLimit` (`UsageEventConsumptionPlanCreditLimit | null`, optional) — Credits consumption only, omitted for balance. The plan credit limit after this event; null when the plan has no limit.
- `PlanBalanceLimit` (`UsageEventConsumptionPlanBalanceLimit | null`, optional) — Balance consumption on a plan with a plan balance limit: the limit after this event. Omitted for credits and for plans without a limit.

### UsageEventConsumptionPlanBalanceLimit

- `Amount` (`int`, required) — The most plan balance the subscription spends per window, in rate scale (10000 = 1.00 of the subscription currency).
- `Interval` (`string`, required) — The window length: 24 hours for `day`, 7 days for `week`.
- `Spent` (`int`, required) — Plan balance spent in the current window, in rate scale.
- `Remaining` (`int`, required) — Plan balance the subscription can still spend in the current window, in rate scale. Purchased balance is not limited.
- `ResetsAt` (`string`, required) — When the next window starts and spending resets.

### UsageEventConsumptionPlanCreditLimit

- `Credits` (`int`, required) — The most plan credits the subscription spends per window.
- `Interval` (`string`, required) — The window length: 24 hours for `day`, 7 days for `week`.
- `Spent` (`int`, required) — Plan credits spent in the current window.
- `Remaining` (`int`, required) — Plan credits the subscription can still spend in the current window. Purchased credits are not limited.
- `ResetsAt` (`string`, required) — When the next window starts and spending resets.

### UsageEventPropertiesItem

- `Property` (`string`, required)
- `Value` (`string`, required)

### UsageQuota

- `FeatureCode` (`string`, required)
- `Current` (`float64`, required)
- `Included` (`float64`, required)
- `Remaining` (`float64 | null`, required)
- `BilledQuantity` (`float64`, required)
- `Unlimited` (`bool`, required)
- `OverageEnabled` (`bool`, required)
- `AsOf` (`string | null`, required)
- `Object` (`string`, required)
- `Livemode` (`bool`, required)

### UsageQuotaEvent

- `ID` (`string`, required)
- `CustomerID` (`string`, required)
- `FeatureCode` (`string`, required)
- `PreviousBalance` (`int`, required)
- `NewBalance` (`int`, required)
- `Ts` (`string`, required)
- `CreatedAt` (`string`, required)
- `Object` (`string`, required)
- `Livemode` (`bool`, required)

### Webhook

- `ID` (`string`, required)
- `URL` (`string`, required)
- `Events` (`[]string`, required)
- `Description` (`string | null`, required)
- `IsActive` (`bool`, required)
- `APIVersion` (`string | null`, required)
- `CreatedAt` (`string`, required)
- `Object` (`string`, required)
- `Livemode` (`bool`, required)

### WebhookAddonRef

- `ID` (`string`, required)
- `Name` (`string`, required)

### WebhookBalance

- `CurrentBalance` (`float64`, required) — The available balance, plan plus purchased, in rate scale. Equal to totalBalance.
- `PlanBalance` (`float64`, required) — The plan part in rate scale: spent first, replaced or rolled over at each reset, negative while overage is owed.
- `PurchasedBalance` (`float64`, required) — The purchased part in rate scale: top-ups and manual adjustments, spent after the plan part and kept until the subscription ends.
- `TotalBalance` (`float64`, required) — planBalance plus purchasedBalance, in rate scale.

### WebhookBankRef

- `BankName` (`string | null`, required)
- `Last4` (`string`, required)

### WebhookCardInfo

- `Brand` (`string`, required)
- `Last4` (`string`, required)
- `ExpMonth` (`float64`, required)
- `ExpYear` (`float64`, required)

### WebhookCreditsBalance

- `PlanCredits` (`float64`, required)
- `PurchasedCredits` (`float64`, required)
- `TotalCredits` (`float64`, required)

### WebhookPlanGrantTimelineEvent

- `ID` (`string`, required) — The public ID of this plan grant event.
- `Type` (`string`, required) — The durable lifecycle transition recorded by this event.
- `Reason` (`string`, required) — The reason recorded for this transition.
- `Source` (`string`, required) — Where this transition originated.
- `PreviousExpiresAt` (`string | null`, required) — The prior expiration deadline for an update, otherwise null.
- `ExpiresAt` (`string | null`, required) — The expiration deadline after this transition, if any.
- `Duration` (`string | null`, required) — The duration selected by a create or update event.
- `DurationCycles` (`int | null`, required) — The selected cycle count when duration is cycles.
- `RequestedExpiresAt` (`string | null`, required) — The requested deadline when duration is until_date.
- `CreatedAt` (`string`, required) — When this transition occurred.

### WebhookPlanRef

- `ID` (`string`, required)
- `Name` (`string`, required)

### WebhookSeatSummary

- `Code` (`string`, required)
- `Current` (`float64 | null`, required)
- `Included` (`float64 | null`, required)
- `Remaining` (`float64 | null`, required)
- `Unlimited` (`bool | null`, required)

### WebhooksListResult

- `Object` (`string`, required)
- `Data` (`[]Webhook`, required)
- `HasMore` (`bool`, required)
- `NextCursor` (`string`, optional)

### WebhookTest

- `Success` (`bool`, required)
- `DeliveryID` (`string`, required)
- `DeliveredAt` (`string`, required)
- `Object` (`string`, required)
- `Livemode` (`bool`, required)
