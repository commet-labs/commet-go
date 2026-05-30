package commet

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"testing"
)

type capturingRoundTripper struct {
	lastBody []byte
}

func (rt *capturingRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	if req.Body != nil {
		rt.lastBody, _ = io.ReadAll(req.Body)
	}
	return &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(bytes.NewReader([]byte(`{"data":{"id":"cus_x"}}`))),
		Header:     make(http.Header),
	}, nil
}

func newCapturingClient(t *testing.T) (*Client, *capturingRoundTripper) {
	t.Helper()
	rt := &capturingRoundTripper{}
	client, err := New("ck_test_123", WithHTTPClient(&http.Client{Transport: rt}), WithTelemetry(false))
	if err != nil {
		t.Fatalf("New: %v", err)
	}
	return client, rt
}

func TestCreateSendsID(t *testing.T) {
	client, rt := newCapturingClient(t)

	_, _ = client.Customers.Create(context.Background(), &CreateCustomerParams{
		Email: "a@b.com",
		ID:    "ext_123",
	})

	var body map[string]any
	if err := json.Unmarshal(rt.lastBody, &body); err != nil {
		t.Fatalf("unmarshal body: %v (raw: %s)", err, rt.lastBody)
	}
	if body["id"] != "ext_123" {
		t.Errorf("expected body id=ext_123, got %v (raw: %s)", body["id"], rt.lastBody)
	}
}

func TestCreateOmitsEmptyID(t *testing.T) {
	client, rt := newCapturingClient(t)

	_, _ = client.Customers.Create(context.Background(), &CreateCustomerParams{Email: "a@b.com"})

	var body map[string]any
	if err := json.Unmarshal(rt.lastBody, &body); err != nil {
		t.Fatalf("unmarshal body: %v", err)
	}
	if _, ok := body["id"]; ok {
		t.Errorf("expected no id in body when ID empty, got %s", rt.lastBody)
	}
}

func TestCreateBatchSendsID(t *testing.T) {
	client, rt := newCapturingClient(t)

	_, _ = client.Customers.CreateBatch(context.Background(), []CreateCustomerParams{
		{Email: "a@b.com", ID: "ext_a"},
		{Email: "b@b.com"},
	}, "")

	var body struct {
		Customers []map[string]any `json:"customers"`
	}
	if err := json.Unmarshal(rt.lastBody, &body); err != nil {
		t.Fatalf("unmarshal body: %v (raw: %s)", err, rt.lastBody)
	}
	if len(body.Customers) != 2 {
		t.Fatalf("expected 2 customers, got %d", len(body.Customers))
	}
	if body.Customers[0]["id"] != "ext_a" {
		t.Errorf("expected customers[0].id=ext_a, got %v", body.Customers[0]["id"])
	}
	if _, ok := body.Customers[1]["id"]; ok {
		t.Errorf("expected no id in customers[1], got %v", body.Customers[1])
	}
}
