package commet

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// capturedRequest records what the SDK actually sent over the wire so tests can
// assert on the path, query, headers and the raw JSON body (post-camelization).
type capturedRequest struct {
	Method string
	Path   string
	Query  map[string]string
	Header http.Header
	Body   []byte
}

// newWireServer spins up an httptest server that records the inbound request and
// replies with the given JSON status + body. It returns a Client pointed at the
// server and a pointer to the captured request (populated after the SDK call).
//
// This mirrors the real layout: the SDK builds requests against http.baseURL,
// which we override to "<server>/api/v1" exactly as subscriptions_test.go does.
func newWireServer(t *testing.T, status int, responseBody string) (*Client, *capturedRequest) {
	t.Helper()

	captured := &capturedRequest{Query: map[string]string{}}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		captured.Method = r.Method
		captured.Path = r.URL.Path
		for k, v := range r.URL.Query() {
			if len(v) > 0 {
				captured.Query[k] = v[0]
			}
		}
		captured.Header = r.Header.Clone()
		if r.Body != nil {
			captured.Body, _ = io.ReadAll(r.Body)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		_, _ = w.Write([]byte(responseBody))
	}))
	t.Cleanup(server.Close)

	client, err := New("ck_test_123", WithTelemetry(false), WithRetries(0))
	if err != nil {
		t.Fatalf("New: %v", err)
	}
	client.http.baseURL = server.URL + "/api/v1"

	return client, captured
}

// decodeBody unmarshals the captured wire body into a generic map for key assertions.
func decodeBody(t *testing.T, raw []byte) map[string]any {
	t.Helper()
	var body map[string]any
	if err := json.Unmarshal(raw, &body); err != nil {
		t.Fatalf("unmarshal request body: %v (raw: %s)", err, raw)
	}
	return body
}

func intPtr(i int) *int    { return &i }
func boolPtr(b bool) *bool { return &b }

var _ = context.Background
