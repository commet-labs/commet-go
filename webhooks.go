package commet

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
)

// Webhooks provides webhook signature verification.
type Webhooks struct{}

// Verify checks that a webhook payload signature is valid.
func (w *Webhooks) Verify(payload string, signature string, secret string) bool {
	if payload == "" || signature == "" || secret == "" {
		return false
	}

	expected := sign(payload, secret)
	return hmac.Equal([]byte(signature), []byte(expected))
}

// VerifyAndParse verifies the signature and parses the payload as JSON.
func (w *Webhooks) VerifyAndParse(rawBody string, signature string, secret string) (map[string]any, error) {
	if !w.Verify(rawBody, signature, secret) {
		return nil, &CommetError{
			Message: "Invalid webhook signature",
			Code:    "INVALID_SIGNATURE",
		}
	}

	var result map[string]any
	if err := json.Unmarshal([]byte(rawBody), &result); err != nil {
		return nil, &CommetError{
			Message: "Failed to parse webhook payload",
			Code:    "INVALID_JSON",
		}
	}

	return result, nil
}

func sign(payload string, secret string) string {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(payload))
	return hex.EncodeToString(mac.Sum(nil))
}
