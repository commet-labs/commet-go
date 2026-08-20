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
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
	"unicode"
)

const version = "9.2.0"

const baseURL = "https://commet.co"

const APIVersion = "2026-07-31"

var retryableStatusCodes = map[int]bool{
	408: true,
	429: true,
	500: true,
	502: true,
	503: true,
	504: true,
}

type httpClient struct {
	client             *http.Client
	baseURL            string
	apiKey             string
	apiVersion         string
	maxRetries         int
	telemetryEnabled   bool
	debug              bool
	defaultTimeout     time.Duration
	lastRequestMetrics *requestMetrics
}

func newHTTPClient(apiKey string, apiVersion string, timeout time.Duration, retries int, telemetry bool, debug bool, customHTTPClient *http.Client) *httpClient {
	client := customHTTPClient
	if client == nil {
		client = &http.Client{}
	}
	return &httpClient{
		client:           client,
		baseURL:          baseURL + "/api/v1",
		apiKey:           apiKey,
		apiVersion:       apiVersion,
		maxRetries:       retries,
		telemetryEnabled: telemetry,
		debug:            debug,
		defaultTimeout:   timeout,
	}
}

func (h *httpClient) resolveApiVersion(opts *RequestOptions) string {
	if opts != nil && opts.ApiVersion != "" {
		return opts.ApiVersion
	}
	if h.apiVersion != "" {
		return h.apiVersion
	}
	return APIVersion
}

func (h *httpClient) resolveTimeout(opts *RequestOptions) time.Duration {
	if opts != nil && opts.Timeout > 0 {
		return opts.Timeout
	}
	return h.defaultTimeout
}

var bodyMethods = map[string]bool{
	http.MethodPost:  true,
	http.MethodPut:   true,
	http.MethodPatch: true,
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
	return h.request(ctx, http.MethodGet, endpoint, nil, cleanParams, "", nil)
}

func (h *httpClient) post(ctx context.Context, endpoint string, body map[string]any, idempotencyKey string) (*rawApiResponse, error) {
	return h.request(ctx, http.MethodPost, endpoint, body, nil, idempotencyKey, nil)
}

func (h *httpClient) put(ctx context.Context, endpoint string, body map[string]any, idempotencyKey string) (*rawApiResponse, error) {
	return h.request(ctx, http.MethodPut, endpoint, body, nil, idempotencyKey, nil)
}

func (h *httpClient) patch(ctx context.Context, endpoint string, body map[string]any, idempotencyKey string) (*rawApiResponse, error) {
	return h.request(ctx, http.MethodPatch, endpoint, body, nil, idempotencyKey, nil)
}

func (h *httpClient) delete(ctx context.Context, endpoint string, body map[string]any, idempotencyKey string) (*rawApiResponse, error) {
	return h.request(ctx, http.MethodDelete, endpoint, body, nil, idempotencyKey, nil)
}

func (h *httpClient) request(ctx context.Context, method string, endpoint string, body map[string]any, params map[string]string, idempotencyKey string, opts *RequestOptions) (*rawApiResponse, error) {
	resolvedKey := idempotencyKey
	if opts != nil && opts.IdempotencyKey != "" {
		resolvedKey = opts.IdempotencyKey
	}

	headers := map[string]string{}
	if bodyMethods[method] && h.maxRetries > 0 && resolvedKey == "" {
		headers["Idempotency-Key"] = "commet-go-retry-" + generateUUID()
	} else if resolvedKey != "" {
		headers["Idempotency-Key"] = resolvedKey
	}

	var jsonBody []byte
	if body != nil {
		generic, err := genericize(body)
		if err != nil {
			return nil, fmt.Errorf("commet: failed to normalize request body: %w", err)
		}
		converted := convertKeys(generic, toCamel)
		jsonBody, err = json.Marshal(converted)
		if err != nil {
			return nil, fmt.Errorf("commet: failed to marshal request body: %w", err)
		}
	}

	return h.execute(ctx, method, endpoint, jsonBody, params, headers, opts, 1)
}

func (h *httpClient) execute(ctx context.Context, method string, endpoint string, jsonBody []byte, params map[string]string, headers map[string]string, opts *RequestOptions, attempt int) (*rawApiResponse, error) {
	fullURL := h.baseURL + endpoint
	if len(params) > 0 {
		query := url.Values{}
		for k, v := range params {
			query.Set(k, v)
		}
		fullURL += "?" + query.Encode()
	}

	if h.debug {
		fmt.Fprintf(os.Stderr, "[Commet SDK] %s %s\n", method, fullURL)
		if jsonBody != nil {
			fmt.Fprintf(os.Stderr, "[Commet SDK] Request body: %s\n", string(jsonBody))
		}
	}

	var bodyReader io.Reader
	if jsonBody != nil {
		bodyReader = bytes.NewReader(jsonBody)
	}

	timeout := h.resolveTimeout(opts)
	reqCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	req, err := http.NewRequestWithContext(reqCtx, method, fullURL, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("commet: failed to create request: %w", err)
	}

	req.Header.Set("x-api-key", h.apiKey)
	req.Header.Set("commet-version", h.resolveApiVersion(opts))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", getUserAgent())
	if h.telemetryEnabled {
		req.Header.Set("commet-client-info", getClientInfoHeader())
		if h.lastRequestMetrics != nil {
			req.Header.Set("commet-client-telemetry", formatTelemetryHeader(*h.lastRequestMetrics))
			h.lastRequestMetrics = nil
		}
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	requestStart := time.Now()
	resp, err := h.client.Do(req)
	if err != nil {
		if attempt <= h.maxRetries {
			delay := h.retryDelay(attempt)
			if h.debug {
				fmt.Fprintf(os.Stderr, "[Commet SDK] Network error, retrying in %v (attempt %d/%d)\n", delay, attempt, h.maxRetries)
			}
			time.Sleep(delay)
			return h.execute(ctx, method, endpoint, jsonBody, params, headers, opts, attempt+1)
		}
		return nil, fmt.Errorf("commet: request failed: %w", err)
	}
	defer resp.Body.Close()

	if h.debug {
		fmt.Fprintf(os.Stderr, "[Commet SDK] Response status: %d\n", resp.StatusCode)
	}

	if retryableStatusCodes[resp.StatusCode] && attempt <= h.maxRetries {
		if delay, retryable := h.statusRetryDelay(resp, attempt); retryable {
			if h.debug {
				fmt.Fprintf(os.Stderr, "[Commet SDK] Retrying in %v (attempt %d/%d)\n", delay, attempt, h.maxRetries)
			}
			time.Sleep(delay)
			return h.execute(ctx, method, endpoint, jsonBody, params, headers, opts, attempt+1)
		}
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

	if resp.StatusCode < 400 && bytes.Equal(bytes.TrimSpace(respBody), []byte("null")) {
		return &rawApiResponse{Success: true, Data: []byte("null")}, nil
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
	success, hasSuccess := converted["success"].(bool)
	data, hasData := converted["data"]
	isEnvelope := hasSuccess && hasData
	if isEnvelope {
		apiResp.Success = success
	}
	if !isEnvelope {
		data = converted
		hasData = true
	}
	if hasData {
		dataBytes, err := json.Marshal(data)
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

	if h.telemetryEnabled {
		requestID := resp.Header.Get("x-request-id")
		if requestID == "" {
			requestID = fmt.Sprintf("req_%d", time.Now().UnixMilli())
		}
		h.lastRequestMetrics = &requestMetrics{
			RequestID:  requestID,
			DurationMs: time.Since(requestStart).Milliseconds(),
		}
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

func parseDirectResponse[T any](raw *rawApiResponse, err error) (*T, error) {
	if err == nil && bytes.Equal(bytes.TrimSpace(raw.Data), []byte("null")) {
		return nil, nil
	}
	response, err := parseResponse[T](raw, err)
	if err != nil {
		return nil, err
	}
	return &response.Data, nil
}

func (h *httpClient) handleError(statusCode int, data map[string]any) error {
	errorObj := data
	if nested, ok := data["error"].(map[string]any); ok {
		errorObj = nested
	}

	errType, _ := errorObj["type"].(string)
	code, _ := errorObj["code"].(string)
	message, _ := errorObj["message"].(string)
	param, _ := errorObj["param"].(string)
	docURL, _ := errorObj["doc_url"].(string)
	details := errorObj["details"]

	if errType == "" {
		errType = "api_error"
	}
	if code == "" {
		code = "unknown"
	}

	if code == "validation_error" {
		if detailsList, ok := details.([]any); ok {
			validationErrors := make(map[string][]string)
			for _, d := range detailsList {
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
					Type:       errType,
					Param:      param,
					DocURL:     docURL,
					Details:    details,
				},
				ValidationErrors: validationErrors,
			}
		}
	}

	return &CommetError{
		Message:    orDefault(message, fmt.Sprintf("Request failed with status %d", statusCode)),
		StatusCode: statusCode,
		Code:       code,
		Type:       errType,
		Param:      param,
		DocURL:     docURL,
		Details:    details,
	}
}

func (h *httpClient) retryDelay(attempt int) time.Duration {
	delay := math.Min(1.0*math.Pow(2, float64(attempt-1)), 8.0)
	return time.Duration(delay * float64(time.Second))
}

const retryAfterCap = 30 * time.Second

// 429 retries wait exactly what the rate limiter reports in Retry-After
// (seconds until the window resets); a 429 without the header did not come
// from the rate limiter, so it is not retried. Exponential backoff only
// applies to statuses that carry no server-provided wait.
func (h *httpClient) statusRetryDelay(resp *http.Response, attempt int) (time.Duration, bool) {
	if resp.StatusCode != http.StatusTooManyRequests {
		return h.retryDelay(attempt), true
	}
	seconds, err := strconv.ParseFloat(resp.Header.Get("Retry-After"), 64)
	if err != nil || seconds <= 0 {
		return 0, false
	}
	return min(time.Duration(seconds*float64(time.Second)), retryAfterCap), true
}

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
		// A typed nil pointer (e.g. (*string)(nil)) boxed into `any` is not == nil,
		// so unset optional fields would otherwise leak as JSON null. Skip them.
		if isNilValue(v) {
			continue
		}
		if s, ok := v.(string); ok && s == "" {
			continue
		}
		body[k] = v
	}
	return body
}

// isNilValue reports whether v wraps a nil pointer. The plain `v == nil` check
// only catches an untyped nil; a typed nil pointer boxed into `any` is not equal
// to nil, so it needs reflection.
func isNilValue(v any) bool {
	rv := reflect.ValueOf(v)
	return rv.Kind() == reflect.Ptr && rv.IsNil()
}

// genericize round-trips a value through JSON so typed nested structs/slices lose
// their snake_case json-tag identity and become plain map[string]any / []any.
// Without this, convertKeys(toCamel) only re-keys the top-level map and nested
// typed values keep their snake_case keys on the wire.
func genericize(v any) (any, error) {
	encoded, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	var generic any
	if err := json.Unmarshal(encoded, &generic); err != nil {
		return nil, err
	}
	return generic, nil
}

func orDefault(value, fallback string) string {
	if value == "" {
		return fallback
	}
	return value
}
