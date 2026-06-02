package commet

import (
	"encoding/json"
	"fmt"
	"os"
	"runtime"
	"sort"
	"sync"
	"testing"
)

type requestMetrics struct {
	RequestID  string `json:"request_id"`
	DurationMs int64  `json:"duration_ms"`
}

type clientInfo struct {
	SDK              string   `json:"sdk"`
	SDKVersion       string   `json:"sdk_version"`
	Lang             string   `json:"lang"`
	LangVersion      string   `json:"lang_version"`
	Platform         string   `json:"platform"`
	Arch             string   `json:"arch"`
	Runtime          string   `json:"runtime"`
	RuntimeVersion   string   `json:"runtime_version"`
	Integrations     []string `json:"integrations"`
	ExecutionContext string   `json:"execution_context,omitempty"`
}

var (
	telemetryMu            sync.Mutex
	cachedClientInfo       string
	cachedUserAgent        string
	registeredIntegrations = map[string]struct{}{}
)

// RegisterIntegration records a wrapper package (name@version) so it is reported
// in the commet-client-info header. Mirrors registerIntegration in the Node SDK.
func RegisterIntegration(name string, version string) {
	telemetryMu.Lock()
	defer telemetryMu.Unlock()
	registeredIntegrations[fmt.Sprintf("%s@%s", name, version)] = struct{}{}
	cachedClientInfo = ""
}

func detectExecutionContext() string {
	if testing.Testing() {
		return "test"
	}
	if os.Getenv("CI") != "" ||
		os.Getenv("GITHUB_ACTIONS") != "" ||
		os.Getenv("GITLAB_CI") != "" ||
		os.Getenv("CIRCLECI") != "" {
		return "ci"
	}
	return ""
}

func collectClientInfo() clientInfo {
	integrations := make([]string, 0, len(registeredIntegrations))
	for integration := range registeredIntegrations {
		integrations = append(integrations, integration)
	}
	sort.Strings(integrations)

	return clientInfo{
		SDK:              "commet-go",
		SDKVersion:       version,
		Lang:             "go",
		LangVersion:      runtime.Version(),
		Platform:         runtime.GOOS,
		Arch:             runtime.GOARCH,
		Runtime:          "go",
		RuntimeVersion:   runtime.Version(),
		Integrations:     integrations,
		ExecutionContext: detectExecutionContext(),
	}
}

func getClientInfoHeader() string {
	telemetryMu.Lock()
	defer telemetryMu.Unlock()
	if cachedClientInfo == "" {
		b, _ := json.Marshal(collectClientInfo())
		cachedClientInfo = string(b)
	}
	return cachedClientInfo
}

func getUserAgent() string {
	telemetryMu.Lock()
	defer telemetryMu.Unlock()
	if cachedUserAgent == "" {
		cachedUserAgent = fmt.Sprintf("commet-go/%s %s %s/%s", version, runtime.Version(), runtime.GOOS, runtime.GOARCH)
	}
	return cachedUserAgent
}

func formatTelemetryHeader(m requestMetrics) string {
	payload := struct {
		LastRequestMetrics requestMetrics `json:"last_request_metrics"`
	}{LastRequestMetrics: m}
	b, _ := json.Marshal(payload)
	return string(b)
}
