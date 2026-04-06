package commet

// ApiResponse represents a standard API response from Commet.
type ApiResponse struct {
	Success    bool   `json:"success"`
	Data       any    `json:"data,omitempty"`
	Code       string `json:"code,omitempty"`
	Message    string `json:"message,omitempty"`
	HasMore    bool   `json:"has_more,omitempty"`
	NextCursor string `json:"next_cursor,omitempty"`
}
