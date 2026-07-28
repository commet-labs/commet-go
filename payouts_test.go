package commet

import (
	"context"
	"testing"
)

// TestAddBankAccountCamelizesSnakeFieldsAndOmitsNilOptionals exercises both
// http.go bug classes on a single request: every snake_case json tag on
// AddPayoutBankAccountParams must reach the wire as camelCase (toCamel runs over
// the whole body), and the unset optional *string / *bool pointers must NOT leak
// as JSON null — buildBody + isNilValue should drop them entirely.
func TestAddBankAccountCamelizesSnakeFieldsAndOmitsNilOptionals(t *testing.T) {
	client, captured := newWireServer(t, 200, `{"data":{"id":"pba_1"}}`)

	_, err := client.Payouts.AddBankAccount(context.Background(), &AddPayoutBankAccountParams{
		AccountNumber:     "000123456789",
		AccountHolderName: "Acme Inc",
		// RoutingNumber, AccountType, SetDefault left nil on purpose.
	})
	if err != nil {
		t.Fatalf("AddBankAccount: %v", err)
	}

	body := decodeBody(t, captured.Body)

	if body["accountNumber"] != "000123456789" {
		t.Errorf("accountNumber = %v, want 000123456789 (raw: %s)", body["accountNumber"], captured.Body)
	}
	if body["accountHolderName"] != "Acme Inc" {
		t.Errorf("accountHolderName = %v, want Acme Inc (raw: %s)", body["accountHolderName"], captured.Body)
	}

	// No snake_case keys may survive the camelization pass.
	for _, snake := range []string{"account_number", "account_holder_name", "routing_number", "account_type", "set_default"} {
		if _, exists := body[snake]; exists {
			t.Errorf("snake_case key %q leaked on the wire (camelization did not run): %s", snake, captured.Body)
		}
	}

	// Unset optional pointers must be absent, not present-as-null. A null leak is
	// one of the two real http.go bugs this guards against.
	for _, optional := range []string{"routingNumber", "accountType", "setDefault"} {
		if v, exists := body[optional]; exists {
			t.Errorf("unset optional %q must be omitted, got present with value %v (raw: %s)", optional, v, captured.Body)
		}
	}
}

// TestAddBankAccountSendsOptionalsWhenSet confirms the set optionals DO reach the
// wire (camelCased), so the omission test above is not passing simply because the
// fields never serialize.
func TestAddBankAccountSendsOptionalsWhenSet(t *testing.T) {
	client, captured := newWireServer(t, 200, `{"data":{"id":"pba_1"}}`)

	_, err := client.Payouts.AddBankAccount(context.Background(), &AddPayoutBankAccountParams{
		AccountNumber:     "000123456789",
		AccountHolderName: "Acme Inc",
		RoutingNumber:     strPtr("110000000"),
		AccountType:       strPtr("checking"),
		SetDefault:        boolPtr(true),
	})
	if err != nil {
		t.Fatalf("AddBankAccount: %v", err)
	}

	body := decodeBody(t, captured.Body)
	if body["routingNumber"] != "110000000" {
		t.Errorf("routingNumber = %v, want 110000000 (raw: %s)", body["routingNumber"], captured.Body)
	}
	if body["accountType"] != "checking" {
		t.Errorf("accountType = %v, want checking (raw: %s)", body["accountType"], captured.Body)
	}
	if body["setDefault"] != true {
		t.Errorf("setDefault = %v, want true (raw: %s)", body["setDefault"], captured.Body)
	}
}

// TestCompleteVerificationCamelizesDeeplyNestedTypedStructs is the strongest guard
// for the nested-camelization bug: CompletePayoutVerificationParams carries typed
// nested structs (Bank, Individual with its own Address). Their json tags are
// snake_case, so without genericize() the inner keys would ship as snake_case.
// We assert the inner-of-inner Address keys are camelCase on the wire.
func TestCompleteVerificationCamelizesDeeplyNestedTypedStructs(t *testing.T) {
	client, captured := newWireServer(t, 200, `{"data":{"provider_account_id":"acct_1","status":"pending","transfers_enabled":false}}`)

	_, err := client.Payouts.CompleteVerification(context.Background(), &CompletePayoutVerificationParams{
		Email:        "ops@acme.com",
		BusinessType: "individual",
		BusinessURL:  "https://acme.com",
		DocumentURL:  "https://files/doc.pdf",
		Bank: CompletePayoutVerificationParamsBank{
			AccountNumber:     "000123456789",
			AccountHolderName: "Jane Roe",
			RoutingNumber:     strPtr("110000000"),
		},
		Individual: &CompletePayoutVerificationParamsIndividual{
			FirstName:   "Jane",
			LastName:    "Roe",
			Phone:       "+15555550123",
			DateOfBirth: "1990-01-01",
			Address: CompletePayoutVerificationParamsIndividualAddress{
				Line1:      "1 Main St",
				City:       "Springfield",
				PostalCode: "11111",
				Country:    "US",
			},
		},
	})
	if err != nil {
		t.Fatalf("CompleteVerification: %v", err)
	}

	body := decodeBody(t, captured.Body)

	// Top-level camelCase.
	if body["businessType"] != "individual" {
		t.Errorf("businessType = %v, want individual (raw: %s)", body["businessType"], captured.Body)
	}
	if body["businessUrl"] != "https://acme.com" {
		t.Errorf("businessUrl = %v, want https://acme.com (raw: %s)", body["businessUrl"], captured.Body)
	}

	// Nested typed struct: bank.accountHolderName must be camelCase, not snake.
	bank, ok := body["bank"].(map[string]any)
	if !ok {
		t.Fatalf("bank = %v, want nested object (raw: %s)", body["bank"], captured.Body)
	}
	if bank["accountHolderName"] != "Jane Roe" {
		t.Errorf("bank.accountHolderName = %v, want Jane Roe (raw: %s)", bank["accountHolderName"], captured.Body)
	}
	if _, exists := bank["account_holder_name"]; exists {
		t.Errorf("nested snake_case bank.account_holder_name leaked (genericize/camelization broken): %s", captured.Body)
	}

	// Doubly-nested: individual.address fields. This is the level a top-level-only
	// re-key would miss entirely.
	individual, ok := body["individual"].(map[string]any)
	if !ok {
		t.Fatalf("individual = %v, want nested object (raw: %s)", body["individual"], captured.Body)
	}
	if individual["firstName"] != "Jane" {
		t.Errorf("individual.firstName = %v, want Jane (raw: %s)", individual["firstName"], captured.Body)
	}
	if individual["dateOfBirth"] != "1990-01-01" {
		t.Errorf("individual.dateOfBirth = %v, want 1990-01-01 (raw: %s)", individual["dateOfBirth"], captured.Body)
	}
	address, ok := individual["address"].(map[string]any)
	if !ok {
		t.Fatalf("individual.address = %v, want nested object (raw: %s)", individual["address"], captured.Body)
	}
	if address["postalCode"] != "11111" {
		t.Errorf("individual.address.postalCode = %v, want 11111 (raw: %s)", address["postalCode"], captured.Body)
	}
	if _, exists := address["postal_code"]; exists {
		t.Errorf("doubly-nested snake_case address.postal_code leaked (deep camelization broken): %s", captured.Body)
	}

	// Optional nested struct left nil must be fully omitted, not null.
	if v, exists := body["company"]; exists {
		t.Errorf("unset optional Company must be omitted, got %v (raw: %s)", v, captured.Body)
	}
	// Individual.ssn_last4 was nil and must not leak as null inside the nested object.
	if v, exists := individual["ssnLast4"]; exists && v == nil {
		t.Errorf("nested optional ssnLast4 leaked as null inside individual (raw: %s)", captured.Body)
	}
}

// TestRequestPayoutUnmarshalsNullableNetAmountAndPointer verifies a camelCase
// response unmarshals into the typed Payout struct, including a wire null mapping
// to a nil *string Description, and integer amount fields decoding correctly.
func TestRequestPayoutUnmarshalsNullableResponse(t *testing.T) {
	client, captured := newWireServer(t, 200, `{
		"id":"po_1",
		"status":"pending",
		"amount":10000,
		"fee":0,
		"netAmount":10000,
		"currency":"usd",
		"description":null,
		"providerTransferId":"tr_abc",
		"createdAt":"2026-06-08T00:00:00Z",
		"object":"payout",
		"livemode":true
	}`)

	resp, err := client.Payouts.Request(context.Background(), &RequestPayoutParams{Amount: 10000})
	if err != nil {
		t.Fatalf("Request: %v", err)
	}

	// The request body should carry amount as a number (no snake_case leakage).
	body := decodeBody(t, captured.Body)
	if body["amount"] != float64(10000) {
		t.Errorf("request amount = %v, want 10000 (raw: %s)", body["amount"], captured.Body)
	}

	p := resp
	if p.ID != "po_1" {
		t.Errorf("ID = %q, want po_1", p.ID)
	}
	if p.NetAmount != 10000 {
		t.Errorf("NetAmount = %d, want 10000", p.NetAmount)
	}
	// http.go converts the camelCase wire keys to snake before re-marshal, so the
	// typed struct's snake json tags ("provider_transfer_id") must populate.
	if p.ProviderTransferID != "tr_abc" {
		t.Errorf("ProviderTransferID = %q, want tr_abc", p.ProviderTransferID)
	}
	// Wire null -> nil pointer.
	if p.Description != nil {
		t.Errorf("Description = %v, want nil for wire null", *p.Description)
	}
	if !p.Livemode {
		t.Error("Livemode = false, want true")
	}
}
