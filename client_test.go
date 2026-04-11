package commet

import "testing"

func TestNewClient(t *testing.T) {
	tests := []struct {
		name    string
		apiKey  string
		wantErr string
	}{
		{
			name:   "valid API key defaults to sandbox",
			apiKey: "ck_test_abc123",
		},
		{
			name:    "empty API key",
			apiKey:  "",
			wantErr: "commet: API key is required",
		},
		{
			name:    "invalid API key format",
			apiKey:  "sk_invalid_key",
			wantErr: "commet: invalid API key format, expected format: ck_xxx...",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, err := New(tt.apiKey)

			if tt.wantErr != "" {
				if err == nil {
					t.Fatalf("expected error %q, got nil", tt.wantErr)
				}
				if err.Error() != tt.wantErr {
					t.Errorf("error = %q, want %q", err.Error(), tt.wantErr)
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if !client.IsSandbox() {
				t.Error("expected default environment to be sandbox")
			}

			client.Close()
		})
	}
}
