package commet

type WebhookPlanRef struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type WebhookAddonRef struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type WebhookCardInfo struct {
	Brand    string `json:"brand"`
	Last4    string `json:"last4"`
	ExpMonth int    `json:"expMonth"`
	ExpYear  int    `json:"expYear"`
}

type WebhookBankRef struct {
	BankName string `json:"bankName"`
	Last4    string `json:"last4"`
}

type WebhookFeatureAccess struct {
	Code             string   `json:"code"`
	Name             string   `json:"name"`
	Type             string   `json:"type"`
	Allowed          bool     `json:"allowed"`
	Enabled          *bool    `json:"enabled"`
	Current          *float64 `json:"current"`
	Included         *float64 `json:"included"`
	Remaining        *float64 `json:"remaining"`
	OverageQuantity  *float64 `json:"overageQuantity"`
	OverageUnitPrice *float64 `json:"overageUnitPrice"`
	Unlimited        *bool    `json:"unlimited"`
	OverageEnabled   *bool    `json:"overageEnabled"`
	BilledQuantity   *float64 `json:"billedQuantity"`
}

type WebhookSeatSummary struct {
	Code      string   `json:"code"`
	Current   *float64 `json:"current"`
	Included  *float64 `json:"included"`
	Remaining *float64 `json:"remaining"`
	Unlimited *bool    `json:"unlimited"`
}

type WebhookCreditsBalance struct {
	PlanCredits      float64 `json:"planCredits"`
	PurchasedCredits float64 `json:"purchasedCredits"`
	TotalCredits     float64 `json:"totalCredits"`
}

type WebhookBalance struct {
	CurrentBalance float64 `json:"currentBalance"`
}
