package commet

import (
	"encoding/json"
	"fmt"
	"runtime"
	"sync"
)

type requestMetrics struct {
	RequestID  string `json:"request_id"`
	DurationMs int64  `json:"duration_ms"`
}

type clientInfo struct {
	SDK            string `json:"sdk"`
	SDKVersion     string `json:"sdk_version"`
	Lang           string `json:"lang"`
	LangVersion    string `json:"lang_version"`
	Platform       string `json:"platform"`
	Arch           string `json:"arch"`
	Runtime        string `json:"runtime"`
	RuntimeVersion string `json:"runtime_version"`
}

var (
	cachedClientInfo string
	cachedUserAgent  string
	telemetryOnce    sync.Once
)

func initTelemetryCache() {
	telemetryOnce.Do(func() {
		info := clientInfo{
			SDK:            "commet-go",
			SDKVersion:     version,
			Lang:           "go",
			LangVersion:    runtime.Version(),
			Platform:       runtime.GOOS,
			Arch:           runtime.GOARCH,
			Runtime:        "go",
			RuntimeVersion: runtime.Version(),
		}
		b, _ := json.Marshal(info)
		cachedClientInfo = string(b)
		cachedUserAgent = fmt.Sprintf("commet-go/%s %s %s/%s", version, runtime.Version(), runtime.GOOS, runtime.GOARCH)
	})
}

func getClientInfoHeader() string {
	initTelemetryCache()
	return cachedClientInfo
}

func getUserAgent() string {
	initTelemetryCache()
	return cachedUserAgent
}

func formatTelemetryHeader(m requestMetrics) string {
	payload := struct {
		LastRequestMetrics requestMetrics `json:"last_request_metrics"`
	}{LastRequestMetrics: m}
	b, _ := json.Marshal(payload)
	return string(b)
}
