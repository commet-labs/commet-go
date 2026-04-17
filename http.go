package commet

import (
	"bytes"
	"context"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"net/url"
	"strings"
	"time"
	"unicode"
)

const version = "1.10.0"

var baseURLs = map[Environment]string{
	Production: "https://commet.co",
	Sandbox:    "https://sandbox.commet.co",
}

var retryableStatusCodes = map[int]bool{
	408: true,
	429: true,
	500: true,
	502: true,
	503: true,
	504: true,
}

type httpClient struct {
	client     *http.Client
	baseURL    string
	apiKey     string
	maxRetries int
}

func newHTTPClient(apiKey string, environment Environment, timeout time.Duration, retries int) *httpClient {
	return &httpClient{
		client:     &http.Client{Timeout: timeout},
		baseURL:    baseURLs[environment] + "/api",
		apiKey:     apiKey,
		maxRetries: retries,
	}
}

func (h *httpClient) close() {
	h.client.CloseIdleConnections()
}

func (h *httpClient) get(ctx context.Context, endpoint string, params map[string]string) (*rawApiResponse, error) {
	cleanParams := make(map[string]string)
	for k, v := range params {
		if v != "" {
			cleanParams[toCamel(k)] = v
		}
	}
	return h.request(ctx, http.MethodGet, endpoint, nil, cleanParams, "")
}

func (h *httpClient) post(ctx context.Context, endpoint string, body map[string]any, idempotencyKey string) (*rawApiResponse, error) {
	return h.request(ctx, http.MethodPost, endpoint, body, nil, idempotencyKey)
}

func (h *httpClient) put(ctx context.Context, endpoint string, body map[string]any, idempotencyKey string) (*rawApiResponse, error) {
	return h.request(ctx, http.MethodPut, endpoint, body, nil, idempotencyKey)
}

func (h *httpClient) delete(ctx context.Context, endpoint string, body map[string]any, idempotencyKey string) (*rawApiResponse, error) {
	return h.request(ctx, http.MethodDelete, endpoint, body, nil, idempotencyKey)
}

func (h *httpClient) request(ctx context.Context, method string, endpoint string, body map[string]any, params map[string]string, idempotencyKey string) (*rawApiResponse, error) {
	headers := map[string]string{}
	if method == http.MethodPost {
		if idempotencyKey != "" {
			headers["Idempotency-Key"] = idempotencyKey
		} else {
			headers["Idempotency-Key"] = "sdk_" + generateUUID()
		}
	}

	var jsonBody []byte
	if body != nil {
		converted := convertKeys(body, toCamel)
		var err error
		jsonBody, err = json.Marshal(converted)
		if err != nil {
			return nil, fmt.Errorf("commet: failed to marshal request body: %w", err)
		}
	}

	return h.execute(ctx, method, endpoint, jsonBody, params, headers, 1)
}

func (h *httpClient) execute(ctx context.Context, method string, endpoint string, jsonBody []byte, params map[string]string, headers map[string]string, attempt int) (*rawApiResponse, error) {
	fullURL := h.baseURL + endpoint
	if len(params) > 0 {
		query := url.Values{}
		for k, v := range params {
			query.Set(k, v)
		}
		fullURL += "?" + query.Encode()
	}

	var bodyReader io.Reader
	if jsonBody != nil {
		bodyReader = bytes.NewReader(jsonBody)
	}

	req, err := http.NewRequestWithContext(ctx, method, fullURL, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("commet: failed to create request: %w", err)
	}

	req.Header.Set("x-api-key", h.apiKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "commet-go/"+version)
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := h.client.Do(req)
	if err != nil {
		if attempt <= h.maxRetries {
			h.wait(attempt)
			return h.execute(ctx, method, endpoint, jsonBody, params, headers, attempt+1)
		}
		return nil, fmt.Errorf("commet: request failed: %w", err)
	}
	defer resp.Body.Close()

	if retryableStatusCodes[resp.StatusCode] && attempt <= h.maxRetries {
		h.wait(attempt)
		return h.execute(ctx, method, endpoint, jsonBody, params, headers, attempt+1)
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("commet: failed to read response body: %w", err)
	}

	if len(respBody) == 0 {
		return nil, &CommetError{
			Message:    fmt.Sprintf("Empty response with status %d", resp.StatusCode),
			StatusCode: resp.StatusCode,
			Code:       "EMPTY_RESPONSE",
		}
	}

	var rawData map[string]any
	if err := json.Unmarshal(respBody, &rawData); err != nil {
		return nil, &CommetError{
			Message:    fmt.Sprintf("Invalid JSON response: %d", resp.StatusCode),
			StatusCode: resp.StatusCode,
			Code:       "INVALID_JSON",
		}
	}

	if resp.StatusCode >= 400 {
		return nil, h.handleError(resp.StatusCode, rawData)
	}

	converted := convertKeys(rawData, toSnake).(map[string]any)

	apiResp := &rawApiResponse{Success: true}
	if v, ok := converted["success"].(bool); ok {
		apiResp.Success = v
	}
	if v, ok := converted["data"]; ok {
		dataBytes, err := json.Marshal(v)
		if err != nil {
			return nil, fmt.Errorf("commet: failed to re-marshal data: %w", err)
		}
		apiResp.Data = dataBytes
	}
	if v, ok := converted["code"].(string); ok {
		apiResp.Code = v
	}
	if v, ok := converted["message"].(string); ok {
		apiResp.Message = v
	}
	if v, ok := converted["has_more"].(bool); ok {
		apiResp.HasMore = v
	}
	if v, ok := converted["next_cursor"].(string); ok {
		apiResp.NextCursor = v
	}

	return apiResp, nil
}

func parseResponse[T any](raw *rawApiResponse, err error) (*ApiResponse[T], error) {
	if err != nil {
		return nil, err
	}

	result := &ApiResponse[T]{
		Success:    raw.Success,
		Code:       raw.Code,
		Message:    raw.Message,
		HasMore:    raw.HasMore,
		NextCursor: raw.NextCursor,
	}

	if len(raw.Data) > 0 {
		if err := json.Unmarshal(raw.Data, &result.Data); err != nil {
			return nil, fmt.Errorf("commet: failed to unmarshal response data: %w", err)
		}
	}

	return result, nil
}

func (h *httpClient) handleError(statusCode int, data map[string]any) error {
	code, _ := data["code"].(string)
	message, _ := data["message"].(string)

	if code == "validation_error" {
		if details, ok := data["details"].([]any); ok {
			validationErrors := make(map[string][]string)
			for _, d := range details {
				detail, ok := d.(map[string]any)
				if !ok {
					continue
				}
				field, _ := detail["field"].(string)
				if field == "" {
					field = "unknown"
				}
				msg, _ := detail["message"].(string)
				validationErrors[field] = append(validationErrors[field], msg)
			}
			return &ValidationError{
				CommetError: CommetError{
					Message:    orDefault(message, "Validation failed"),
					StatusCode: statusCode,
					Code:       code,
				},
				ValidationErrors: validationErrors,
			}
		}
	}

	return &CommetError{
		Message:    orDefault(message, fmt.Sprintf("Request failed with status %d", statusCode)),
		StatusCode: statusCode,
		Code:       code,
		Details:    data["details"],
	}
}

func (h *httpClient) wait(attempt int) {
	delay := math.Min(1.0*math.Pow(2, float64(attempt-1)), 8.0)
	time.Sleep(time.Duration(delay * float64(time.Second)))
}

// -- Case conversion --

func toSnake(s string) string {
	var result strings.Builder
	for i, r := range s {
		if unicode.IsUpper(r) {
			if i > 0 {
				prev := rune(s[i-1])
				if unicode.IsLower(prev) || unicode.IsDigit(prev) {
					result.WriteRune('_')
				} else if unicode.IsUpper(prev) && i+1 < len(s) && unicode.IsLower(rune(s[i+1])) {
					result.WriteRune('_')
				}
			}
			result.WriteRune(unicode.ToLower(r))
		} else {
			result.WriteRune(r)
		}
	}
	return result.String()
}

func toCamel(s string) string {
	parts := strings.Split(s, "_")
	if len(parts) == 0 {
		return s
	}
	result := parts[0]
	for _, part := range parts[1:] {
		if len(part) > 0 {
			result += strings.ToUpper(part[:1]) + part[1:]
		}
	}
	return result
}

func convertKeys(obj any, fn func(string) string) any {
	switch v := obj.(type) {
	case map[string]any:
		converted := make(map[string]any, len(v))
		for key, val := range v {
			converted[fn(key)] = convertKeys(val, fn)
		}
		return converted
	case []any:
		converted := make([]any, len(v))
		for i, val := range v {
			converted[i] = convertKeys(val, fn)
		}
		return converted
	default:
		return obj
	}
}

// -- Helpers --

func generateUUID() string {
	b := make([]byte, 16)
	_, _ = rand.Read(b)
	return fmt.Sprintf("%x%x%x%x%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}

func buildBody(fields map[string]any) map[string]any {
	body := make(map[string]any)
	for k, v := range fields {
		if v == nil {
			continue
		}
		if s, ok := v.(string); ok && s == "" {
			continue
		}
		body[k] = v
	}
	return body
}

func orDefault(value, fallback string) string {
	if value == "" {
		return fallback
	}
	return value
}
