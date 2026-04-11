package commet

import "encoding/json"

type ApiResponse[T any] struct {
	Success    bool   `json:"success"`
	Data       T      `json:"data,omitempty"`
	Code       string `json:"code,omitempty"`
	Message    string `json:"message,omitempty"`
	HasMore    bool   `json:"has_more,omitempty"`
	NextCursor string `json:"next_cursor,omitempty"`
}

type rawApiResponse struct {
	Success    bool            `json:"success"`
	Data       json.RawMessage `json:"data,omitempty"`
	Code       string          `json:"code,omitempty"`
	Message    string          `json:"message,omitempty"`
	HasMore    bool            `json:"has_more,omitempty"`
	NextCursor string          `json:"next_cursor,omitempty"`
}
