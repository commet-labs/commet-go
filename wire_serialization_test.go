package commet

import (
	"context"
	"testing"
)

// TestCreateFeatureSendsTypedEnumAsWireString verifies a top-level typed enum
// field (FeatureType) serializes to its wire string value in the request body and
// that the snake_case json tag "unit_name" becomes "unitName" on the wire.
func TestCreateFeatureSendsTypedEnumAsWireString(t *testing.T) {
	client, captured := newWireServer(t, 200, `{"id":"feat_1","name":"API Calls","code":"api_calls","type":"usage","description":null,"unitName":"call","createdAt":"2026-01-01","updatedAt":"2026-01-01","object":"feature","livemode":true}`)

	resp, err := client.Features.Create(context.Background(), &CreateFeatureParams{
		Name:     "API Calls",
		Code:     "api_calls",
		Type:     "usage",
		UnitName: strPtr("call"),
	})
	if err != nil {
		t.Fatalf("Create: %v", err)
	}

	body := decodeBody(t, captured.Body)
	if body["type"] != "usage" {
		t.Errorf("body type = %v, want usage (typed enum did not serialize as string) (raw: %s)", body["type"], captured.Body)
	}
	if body["unitName"] != "call" {
		t.Errorf("body unitName = %v, want call (raw: %s)", body["unitName"], captured.Body)
	}
	if _, exists := body["unit_name"]; exists {
		t.Errorf("snake_case unit_name leaked: %s", captured.Body)
	}

	// Response: typed enum decodes, wire null -> nil pointer.
	if resp.Type != FeatureTypeUsage {
		t.Errorf("resp Type = %q, want usage enum", resp.Type)
	}
	if resp.Description != nil {
		t.Errorf("resp Description = %v, want nil for wire null", *resp.Description)
	}
	if resp.UnitName == nil || *resp.UnitName != "call" {
		t.Errorf("resp UnitName = %v, want call", resp.UnitName)
	}
}

// TestListTransactionsSendsEnumStatusAsQueryParam verifies a typed enum filter
// (*TransactionStatus) is emitted as its wire string in the query string, and the
// numeric Limit is stringified. List responses decode the typed Status enum and a
// nullable *string (InvoiceID null -> nil).
func TestListTransactionsSendsEnumStatusAsQueryParam(t *testing.T) {
	client, captured := newWireServer(t, 200, `{"data":[{"id":"txn_1","invoiceId":null,"grossAmount":1000,"subtotal":900,"taxAmount":100,"currency":"usd","status":"succeeded","customerEmail":"a@b.com","customerName":null,"paidAt":"2026-01-01","createdAt":"2026-01-01","updatedAt":"2026-01-01","object":"transaction","livemode":true}],"hasMore":false}`)

	status := TransactionStatusSucceeded
	resp, err := client.Transactions.List(context.Background(), &ListTransactionsParams{
		Status: &status,
		Limit:  intPtr(25),
	})
	if err != nil {
		t.Fatalf("List: %v", err)
	}

	// get() camelizes query keys; "status" stays "status", value is the enum string.
	if captured.Query["status"] != "succeeded" {
		t.Errorf("query status = %q, want succeeded (raw query: %v)", captured.Query["status"], captured.Query)
	}
	if captured.Query["limit"] != "25" {
		t.Errorf("query limit = %q, want 25", captured.Query["limit"])
	}

	if len(resp.Data) != 1 {
		t.Fatalf("len(Data) = %d, want 1", len(resp.Data))
	}
	txn := resp.Data[0]
	if txn.Status != TransactionStatusSucceeded {
		t.Errorf("Status = %q, want succeeded enum", txn.Status)
	}
	if txn.InvoiceID != nil {
		t.Errorf("InvoiceID = %v, want nil for wire null", *txn.InvoiceID)
	}
	if txn.CustomerName != nil {
		t.Errorf("CustomerName = %v, want nil for wire null", *txn.CustomerName)
	}
	if txn.GrossAmount == nil || *txn.GrossAmount != 1000 || txn.TaxAmount == nil || *txn.TaxAmount != 100 {
		t.Errorf("amounts = (%v,%v), want (1000,100)", txn.GrossAmount, txn.TaxAmount)
	}
}

// TestGetInvoiceUnmarshalsTypedEnumLineItemsAndNullables checks an Invoice response
// with the InvoiceType enum, an explicit-null *int (CreditApplied via omitempty),
// nested line items array, and a present optional all decode correctly.
func TestGetInvoiceUnmarshalsTypedEnumLineItemsAndNullables(t *testing.T) {
	client, _ := newWireServer(t, 200, `{
		"id":"inv_1",
		"customerId":"cus_1",
		"subscriptionId":null,
		"invoiceNumber":"INV-001",
		"status":"paid",
		"invoiceType":"plan_change",
		"currency":"usd",
		"subtotal":1000,
		"discountAmount":0,
		"taxAmount":100,
		"total":1100,
		"periodStart":"2026-01-01",
		"periodEnd":"2026-02-01",
		"issueDate":"2026-01-01",
		"dueDate":"2026-01-15",
		"memo":null,
		"metadata":{},
		"createdAt":"2026-01-01",
		"updatedAt":"2026-01-01",
		"object":"invoice",
		"livemode":true,
		"lineItems":[{"lineType":"plan","featureName":null,"description":"Pro plan"}]
	}`)

	resp, err := client.Invoices.Get(context.Background(), "inv_1")
	if err != nil {
		t.Fatalf("Get: %v", err)
	}
	inv := resp

	if inv.InvoiceType != InvoiceTypePlanChange {
		t.Errorf("InvoiceType = %q, want plan_change enum", inv.InvoiceType)
	}
	if inv.SubscriptionID != nil {
		t.Errorf("SubscriptionID = %v, want nil for wire null", *inv.SubscriptionID)
	}
	if inv.Total != 1100 {
		t.Errorf("Total = %d, want 1100", inv.Total)
	}
	if len(inv.LineItems) != 1 {
		t.Fatalf("len(LineItems) = %d, want 1", len(inv.LineItems))
	}
	if inv.LineItems[0].Description != "Pro plan" {
		t.Errorf("LineItems[0].Description = %q, want Pro plan", inv.LineItems[0].Description)
	}
}

// TestSetAllSeatsSendsMapBody verifies the seats bulk endpoint serializes a
// map[string]int body verbatim (keys are user-supplied feature codes, NOT struct
// fields, so they must NOT be camelized) under the camelCase "customerId" key.
func TestSetAllSeatsSendsMapBody(t *testing.T) {
	client, captured := newWireServer(t, 200, `{"data":[{"object":"bulk_seat_update","livemode":true}]}`)

	_, err := client.Seats.SetAll(context.Background(), &BulkSetSeatsParams{
		CustomerID: "cus_1",
		Seats:      map[string]int{"editor_seats": 3, "viewer_seats": 10},
	})
	if err != nil {
		t.Fatalf("SetAll: %v", err)
	}

	body := decodeBody(t, captured.Body)
	if body["customerId"] != "cus_1" {
		t.Errorf("customerId = %v, want cus_1 (raw: %s)", body["customerId"], captured.Body)
	}
	seats, ok := body["seats"].(map[string]any)
	if !ok {
		t.Fatalf("seats = %v, want object (raw: %s)", body["seats"], captured.Body)
	}
	// Map keys are data, not field names: convertKeys does run over them (toCamel),
	// so "editor_seats" would become "editorSeats". Document the actual behavior so
	// callers passing snake-cased feature codes know they get camelized on the wire.
	if seats["editorSeats"] != float64(3) {
		t.Errorf("seats[editorSeats] = %v, want 3 (map key camelization changed) (raw: %s)", seats["editorSeats"], captured.Body)
	}
	if seats["viewerSeats"] != float64(10) {
		t.Errorf("seats[viewerSeats] = %v, want 10 (raw: %s)", seats["viewerSeats"], captured.Body)
	}
}

// TestCreateCreditPackOmitsNilOptionalsAndKeepsZeroPrice guards the nil-pointer
// omission for *bool IsActive while ensuring a legitimately-zero int (Price=0) is
// still sent — buildBody must distinguish "unset pointer" from "zero value".
func TestCreateCreditPackOmitsNilOptionalsAndKeepsZeroPrice(t *testing.T) {
	client, captured := newWireServer(t, 200, `{"id":"cp_1","name":"Starter","credits":100,"price":0,"object":"credit_pack","livemode":true}`)

	_, err := client.CreditPacks.Create(context.Background(), &CreateCreditPackParams{
		Name:    "Starter",
		Credits: 100,
		Price:   0,
		// IsActive nil on purpose.
	})
	if err != nil {
		t.Fatalf("Create: %v", err)
	}

	body := decodeBody(t, captured.Body)
	if body["credits"] != float64(100) {
		t.Errorf("credits = %v, want 100 (raw: %s)", body["credits"], captured.Body)
	}
	if body["price"] != float64(0) {
		t.Errorf("price = %v, want 0 — zero int must still be sent (raw: %s)", body["price"], captured.Body)
	}
	if v, exists := body["isActive"]; exists {
		t.Errorf("unset IsActive must be omitted, got %v (raw: %s)", v, captured.Body)
	}
	if _, exists := body["is_active"]; exists {
		t.Errorf("snake_case is_active leaked: %s", captured.Body)
	}
}

// TestCreateAddonSendsRequiredAndOmitsOptionals exercises a mixed body: required
// snake-tagged fields (base_price, feature_id, consumption_model) camelize, while
// the unset *int optionals are dropped (not null).
func TestCreateAddonSendsRequiredAndOmitsOptionals(t *testing.T) {
	client, captured := newWireServer(t, 200, `{"id":"addon_1","name":"Extra","slug":"extra","basePrice":500,"consumptionModel":"metered","featureCode":"x","featureName":"X","object":"addon","livemode":true}`)

	_, err := client.Addons.Create(context.Background(), &CreateAddonParams{
		Name:             "Extra",
		BasePrice:        500,
		FeatureID:        "feat_1",
		ConsumptionModel: "metered",
		// IncludedUnits, OverageRate, CreditCost nil.
	})
	if err != nil {
		t.Fatalf("Create: %v", err)
	}

	body := decodeBody(t, captured.Body)
	if body["basePrice"] != float64(500) {
		t.Errorf("basePrice = %v, want 500 (raw: %s)", body["basePrice"], captured.Body)
	}
	if body["featureId"] != "feat_1" {
		t.Errorf("featureId = %v, want feat_1 (raw: %s)", body["featureId"], captured.Body)
	}
	if body["consumptionModel"] != "metered" {
		t.Errorf("consumptionModel = %v, want metered (raw: %s)", body["consumptionModel"], captured.Body)
	}
	for _, snake := range []string{"base_price", "feature_id", "consumption_model"} {
		if _, exists := body[snake]; exists {
			t.Errorf("snake_case key %q leaked: %s", snake, captured.Body)
		}
	}
	for _, optional := range []string{"includedUnits", "overageRate", "creditCost"} {
		if v, exists := body[optional]; exists {
			t.Errorf("unset optional %q must be omitted, got %v (raw: %s)", optional, v, captured.Body)
		}
	}
}
