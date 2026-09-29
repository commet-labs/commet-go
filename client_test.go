package commet

import "testing"

func TestNewClient(t *testing.T) {
	tests := []struct {
		name    string
		apiKey  string
		wantErr string
	}{
		{
			name:   "valid API key",
			apiKey: "ck_test_abc123",
		},
		{
			name:   "restricted API key",
			apiKey: "rk_test_abc123",
		},
		{
			name:   "live restricted API key",
			apiKey: "rk_live_abc123",
		},
		{
			name:   "sandbox restricted API key",
			apiKey: "rk_sandbox_abc123",
		},
		{
			name:   "live API key",
			apiKey: "ck_live_abc123",
		},
		{
			name:   "sandbox API key",
			apiKey: "ck_sandbox_abc123",
		},
		{
			name:    "empty API key",
			apiKey:  "",
			wantErr: "commet: API key is required",
		},
		{
			name:    "invalid API key format",
			apiKey:  "sk_invalid_key",
			wantErr: "commet: invalid API key format, expected prefix ck_ or rk_",
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

			client.Close()
		})
	}
}
