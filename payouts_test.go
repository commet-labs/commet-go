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

func TestCompleteVerificationSendsNoKYCBody(t *testing.T) {
	client, captured := newWireServer(t, 200, `{"success":true,"data":null}`)

	result, err := client.Payouts.CompleteVerification(context.Background())
	if err != nil {
		t.Fatalf("CompleteVerification: %v", err)
	}
	if result != nil {
		t.Fatalf("CompleteVerification result = %v, want nil", result)
	}

	body := decodeBody(t, captured.Body)
	if len(body) != 0 {
		t.Errorf("expected empty JSON body, got %s", captured.Body)
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
