# Schemas

Generated from Commet API version `2026-07-31`.

## Enums

### BillingInterval

- `"weekly"`
- `"monthly"`
- `"quarterly"`
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

### PaymentProvider

- `"stripe"`
- `"commet"`
- `"dlocal"`

### SubscriptionStatus

- `"draft"`
- `"pending_payment"`
- `"trialing"`
- `"active"`
- `"past_due"`
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

### ClaimLink

- `URL` (`string`, required)
- `ExpiresAt` (`string`, required)
- `Object` (`string`, required)
- `Livemode` (`bool`, required)

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
- `CreatedAt` (`string`, required)
- `UpdatedAt` (`string`, required)
- `OfferApplications` (`[]SubscriptionOfferApplication`, required)
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

### CreatedSubscriptionPlan

- `ID` (`string`, required)
- `Name` (`string`, required)

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
- `CreateOfferParamsPhasesItemVariant4`

Discriminator: `Type`

- `"free_trial"` → `CreateOfferParamsPhasesItemVariant1`
- `"percentage"` → `CreateOfferParamsPhasesItemVariant2`
- `"amount_off"` → `CreateOfferParamsPhasesItemVariant3`
- `"fixed_price"` → `CreateOfferParamsPhasesItemVariant4`

### CreateOfferParamsPhasesItemVariant1

- `Type` (`string`, required)
- `DurationDays` (`int`, required)

### CreateOfferParamsPhasesItemVariant2

- `Type` (`string`, required)
- `DurationCycles` (`int | null`, required)
- `DurationInterval` (`string | null`, optional) — Unit the phase duration is counted in. Only a fixed-price phase may set it, because its amount is declared rather than derived from the plan. Defaults to the plan's own billing interval.
- `Percentage` (`int`, required) — Discount in basis points. 5000 means 50%.

### CreateOfferParamsPhasesItemVariant3

- `Type` (`string`, required)
- `DurationCycles` (`int | null`, required)
- `DurationInterval` (`string | null`, optional) — Unit the phase duration is counted in. Only a fixed-price phase may set it, because its amount is declared rather than derived from the plan. Defaults to the plan's own billing interval.
- `Amounts` (`[]CreateOfferParamsPhasesItemVariant3AmountsItem`, required)

### CreateOfferParamsPhasesItemVariant3AmountsItem

- `Currency` (`string`, required)
- `Amount` (`int`, required) — Amount in the currency's minor unit (for example, cents for USD).

### CreateOfferParamsPhasesItemVariant4

- `Type` (`string`, required)
- `DurationCycles` (`int | null`, required)
- `DurationInterval` (`string | null`, optional) — Unit the phase duration is counted in. Only a fixed-price phase may set it, because its amount is declared rather than derived from the plan. Defaults to the plan's own billing interval.
- `Prices` (`[]CreateOfferParamsPhasesItemVariant4PricesItem`, required)

### CreateOfferParamsPhasesItemVariant4PricesItem

- `Currency` (`string`, required)
- `Amount` (`int`, required) — Amount in the currency's minor unit (for example, cents for USD).

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
- `CreatedAt` (`string`, required)
- `UpdatedAt` (`string`, required)
- `Object` (`string`, required)
- `Livemode` (`bool`, required)

### OfferPhasesItem

Variants:

- `OfferPhasesItemVariant1`
- `OfferPhasesItemVariant2`
- `OfferPhasesItemVariant3`
- `OfferPhasesItemVariant4`

Discriminator: `Type`

- `"free_trial"` → `OfferPhasesItemVariant1`
- `"percentage"` → `OfferPhasesItemVariant2`
- `"amount_off"` → `OfferPhasesItemVariant3`
- `"fixed_price"` → `OfferPhasesItemVariant4`

### OfferPhasesItemVariant1

- `Type` (`string`, required)
- `DurationDays` (`int`, required)

### OfferPhasesItemVariant2

- `Type` (`string`, required)
- `DurationCycles` (`int | null`, required)
- `DurationInterval` (`string | null`, required) — Unit the phase duration is counted in. Only a fixed-price phase may set it, because its amount is declared rather than derived from the plan. Defaults to the plan's own billing interval.
- `Percentage` (`int`, required) — Discount in basis points. 5000 means 50%.

### OfferPhasesItemVariant3

- `Type` (`string`, required)
- `DurationCycles` (`int | null`, required)
- `DurationInterval` (`string | null`, required) — Unit the phase duration is counted in. Only a fixed-price phase may set it, because its amount is declared rather than derived from the plan. Defaults to the plan's own billing interval.
- `Amounts` (`[]OfferPhasesItemVariant3AmountsItem`, required)

### OfferPhasesItemVariant3AmountsItem

- `Currency` (`string`, required)
- `Amount` (`int`, required) — Amount in the currency's minor unit (for example, cents for USD).

### OfferPhasesItemVariant4

- `Type` (`string`, required)
- `DurationCycles` (`int | null`, required)
- `DurationInterval` (`string | null`, required) — Unit the phase duration is counted in. Only a fixed-price phase may set it, because its amount is declared rather than derived from the plan. Defaults to the plan's own billing interval.
- `Prices` (`[]OfferPhasesItemVariant4PricesItem`, required)

### OfferPhasesItemVariant4PricesItem

- `Currency` (`string`, required)
- `Amount` (`int`, required) — Amount in the currency's minor unit (for example, cents for USD).

### OffersListResult

- `Object` (`string`, required)
- `Data` (`[]Offer`, required)
- `HasMore` (`bool`, required)
- `NextCursor` (`string`, optional)

### Payment

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
- `BlockOnExhaustion` (`bool | null`, required)
- `SortOrder` (`int`, required)
- `PlanGroupID` (`string | null`, required)
- `Metadata` (`map[string]any | null`, required)
- `CreatedAt` (`string`, required)
- `UpdatedAt` (`string`, required)
- `Features` (`[]PlanFeaturesItem`, required)
- `Prices` (`[]PlanPricesItem`, required)
- `ExchangeRates` (`[]PlanExchangeRatesItem`, required)
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

### PlanExchangeRatesItem

- `Currency` (`string`, required)
- `ExchangeRate` (`float64`, required)

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

### PlanPrice

- `ID` (`string`, required) — Public plan price ID.
- `PlanID` (`string`, required)
- `BillingInterval` (`BillingInterval`, required)
- `Price` (`int`, required) — Price in the currency's minor unit (for example, cents for USD).
- `IsDefault` (`bool`, required)
- `TrialDays` (`int`, required)
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

### PlansListResult

- `Object` (`string`, required)
- `Data` (`[]Plan`, required)
- `HasMore` (`bool`, required)
- `NextCursor` (`string`, optional)

### PortalAccess

- `PortalURL` (`string`, required)
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
- `OfferApplication` (`PreviewChangeOfferApplication`, optional)
- `Object` (`string`, required)
- `Livemode` (`bool`, required)

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
- `CreatedAt` (`string`, required)
- `UpdatedAt` (`string`, required)
- `Object` (`string`, required)
- `Livemode` (`bool`, required)

### PromoCodesListResult

- `Object` (`string`, required)
- `Data` (`[]PromoCode`, required)
- `HasMore` (`bool`, required)
- `NextCursor` (`string`, optional)

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
- `CreatedAt` (`string`, required)
- `UpdatedAt` (`string`, required)
- `OfferApplications` (`[]SubscriptionOfferApplication`, required)
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

### SubscriptionBalance

- `Remaining` (`float64`, required)
- `Included` (`float64`, required)
- `Currency` (`string`, required)

### SubscriptionCancellation

- `ScheduledAt` (`string`, required)
- `Reason` (`string | null`, required)
- `EffectiveAt` (`string`, required)

### SubscriptionCredits

- `Remaining` (`float64`, required)
- `Included` (`float64`, required)
- `Purchased` (`float64`, required)

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

### SubscriptionScheduledPlanChange

- `ChangeType` (`string`, required)
- `NewPlanID` (`string | null`, required)
- `NewPlanName` (`string | null`, required)
- `NewBillingInterval` (`string | null`, required)
- `ScheduledFor` (`string`, required)

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
- `CreatedAt` (`string`, required)
- `UpdatedAt` (`string`, required)
- `OfferApplications` (`[]SubscriptionOfferApplication`, required)
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

### SubscriptionSummaryPlan

- `ID` (`string`, required)
- `Name` (`string`, required)

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

- `ID` (`string`, required)
- `InvoiceID` (`string | null`, required)
- `GrossAmount` (`int | null`, required) — Gross amount in USD cents. Null when the provider has not reported an honest USD figure; see presentmentAmount.
- `Subtotal` (`int | null`, required) — Subtotal in USD cents (gross minus tax). Null when grossAmount is null.
- `TaxAmount` (`int | null`, required)
- `PresentmentAmount` (`int | null`, required) — Amount in the charge currency's smallest unit, as presented to the customer. Set for non-USD charges; null when the charge was made in USD.
- `Currency` (`string`, required)
- `Provider` (`PaymentProvider`, required) — The payment provider the charge was routed to: stripe, commet, or dlocal.
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

- `ID` (`string`, required)
- `InvoiceID` (`string | null`, required)
- `GrossAmount` (`int | null`, required) — Gross amount in USD cents. Null when the provider has not reported an honest USD figure; see presentmentAmount.
- `Subtotal` (`int | null`, required) — Subtotal in USD cents (gross minus tax). Null when grossAmount is null.
- `TaxAmount` (`int | null`, required)
- `PresentmentAmount` (`int | null`, required) — Amount in the charge currency's smallest unit, as presented to the customer. Set for non-USD charges; null when the charge was made in USD.
- `Currency` (`string`, required)
- `Provider` (`PaymentProvider`, required) — The payment provider the charge was routed to: stripe, commet, or dlocal.
- `Status` (`TransactionStatus`, required)
- `CustomerEmail` (`string | null`, required)
- `CustomerName` (`string | null`, required)
- `PaidAt` (`string | null`, required)
- `CreatedAt` (`string`, required)
- `UpdatedAt` (`string`, required)
- `Object` (`string`, required)
- `Livemode` (`bool`, required)

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
- `UpdateOfferParamsPhasesItemVariant4`

Discriminator: `Type`

- `"free_trial"` → `UpdateOfferParamsPhasesItemVariant1`
- `"percentage"` → `UpdateOfferParamsPhasesItemVariant2`
- `"amount_off"` → `UpdateOfferParamsPhasesItemVariant3`
- `"fixed_price"` → `UpdateOfferParamsPhasesItemVariant4`

### UpdateOfferParamsPhasesItemVariant1

- `Type` (`string`, required)
- `DurationDays` (`int`, required)

### UpdateOfferParamsPhasesItemVariant2

- `Type` (`string`, required)
- `DurationCycles` (`int | null`, required)
- `DurationInterval` (`string | null`, optional) — Unit the phase duration is counted in. Only a fixed-price phase may set it, because its amount is declared rather than derived from the plan. Defaults to the plan's own billing interval.
- `Percentage` (`int`, required) — Discount in basis points. 5000 means 50%.

### UpdateOfferParamsPhasesItemVariant3

- `Type` (`string`, required)
- `DurationCycles` (`int | null`, required)
- `DurationInterval` (`string | null`, optional) — Unit the phase duration is counted in. Only a fixed-price phase may set it, because its amount is declared rather than derived from the plan. Defaults to the plan's own billing interval.
- `Amounts` (`[]UpdateOfferParamsPhasesItemVariant3AmountsItem`, required)

### UpdateOfferParamsPhasesItemVariant3AmountsItem

- `Currency` (`string`, required)
- `Amount` (`int`, required) — Amount in the currency's minor unit (for example, cents for USD).

### UpdateOfferParamsPhasesItemVariant4

- `Type` (`string`, required)
- `DurationCycles` (`int | null`, required)
- `DurationInterval` (`string | null`, optional) — Unit the phase duration is counted in. Only a fixed-price phase may set it, because its amount is declared rather than derived from the plan. Defaults to the plan's own billing interval.
- `Prices` (`[]UpdateOfferParamsPhasesItemVariant4PricesItem`, required)

### UpdateOfferParamsPhasesItemVariant4PricesItem

- `Currency` (`string`, required)
- `Amount` (`int`, required) — Amount in the currency's minor unit (for example, cents for USD).

### UpdatePlanFeatureParamsOverage

- `Enabled` (`bool`, optional)
- `UnitPrice` (`int`, optional)

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
- `Object` (`string`, required)
- `Livemode` (`bool`, required)

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
- `CurrentBalance` (`float64`, required)
- `BlockOnExhaustion` (`bool`, required)
- `Currency` (`string`, required)
- `Object` (`string`, required)
- `Livemode` (`bool`, required)

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

- `CurrentBalance` (`float64`, required)

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
