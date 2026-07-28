package commet

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
)

type WebhooksResource struct {
	*GeneratedWebhooksResource
}

func (w *WebhooksResource) Verify(payload string, signature string, secret string) bool {
	if payload == "" || signature == "" || secret == "" {
		return false
	}

	expected := sign(payload, secret)
	return hmac.Equal([]byte(signature), []byte(expected))
}

func (w *WebhooksResource) VerifyAndParse(rawBody string, signature string, secret string) (map[string]any, error) {
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
