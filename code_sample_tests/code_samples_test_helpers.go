package code_sample_tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
)

// RunCodeSampleServer creates a test server that returns mock responses for code sample tests.
// It's exported so it can be used by tests in separate packages (e.g., gocardless_test).
func RunCodeSampleServer(envelope string, isList bool) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var response map[string]any
		if isList {
			response = map[string]any{
				envelope: []any{map[string]any{}},
				"meta": map[string]any{
					"cursors": map[string]any{},
					"limit":   50,
				},
			}
		} else {
			response = map[string]any{
				envelope: map[string]any{},
			}
		}
		json.NewEncoder(w).Encode(response)
	}))
}
