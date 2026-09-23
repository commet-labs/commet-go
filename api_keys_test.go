package commet

import (
	"context"
	"testing"
)

func TestCreateApiKeyWirePermissions(t *testing.T) {
	for _, test := range []struct {
		name        string
		permissions map[string][]string
		wantPresent bool
	}{
		{name: "omitted permissions"},
		{name: "empty permissions", permissions: map[string][]string{}, wantPresent: true},
		{name: "underscored resource", permissions: map[string][]string{"plan_group": {"read"}}, wantPresent: true},
	} {
		t.Run(test.name, func(t *testing.T) {
			client, captured := newWireServer(t, 201, `{}`)
			defer client.Close()
			_, err := client.ApiKeys.Create(context.Background(), &CreateApiKeyParams{
				Name: "Example", Permissions: test.permissions,
			})
			if err != nil {
				t.Fatal(err)
			}
			body := decodeBody(t, captured.Body)
			permissions, present := body["permissions"]
			if present != test.wantPresent {
				t.Fatalf("permissions present = %v, want %v: %s", present, test.wantPresent, captured.Body)
			}
			if test.name == "underscored resource" {
				grants := permissions.(map[string]any)
				if _, ok := grants["plan_group"]; !ok {
					t.Fatalf("plan_group missing: %s", captured.Body)
				}
			}
			if test.name == "empty permissions" {
				if len(permissions.(map[string]any)) != 0 {
					t.Fatalf("permissions not empty: %s", captured.Body)
				}
			}
		})
	}
}
