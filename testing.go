package gocardless

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func runServer(t *testing.T, fixtureFile string, method string) *httptest.Server {
	f, _ := os.Open(fixtureFile)
	defer f.Close()

	responseBody, err := io.ReadAll(f)
	if err != nil {
		t.Fatal(err)
	}

	var response map[string]interface{}
	json.Unmarshal(responseBody, &response)
	body, ok := response[method].(map[string]interface{})
	if !ok {
		t.Fatal(fmt.Printf("error, response[%s] is nil or value stored is not of type map[string]interface", method))
	}
	bodyStr, ok := body["body"].(map[string]interface{})
	if !ok {
		t.Fatal("error, body[\"body\"] is nil or value stored is not of type map[string]interface")
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(bodyStr)
	}))

	return server
}

func getClient(t *testing.T, url string) (*Service, error) {
	token := "dummy_token"
	config, err := NewConfig(token, WithEndpoint(url))
	if err != nil {
		t.Fatal(err)
	}
	service, err := New(config)
	if err != nil {
		t.Fatal(err)
	}
	return service, nil
}

// GetClient creates a test client for use in code sample tests.
// It's exported so it can be used by tests in separate packages (e.g., gocardless_test).
func GetClient(t *testing.T, url string) (*Service, error) {
	return getClient(t, url)
}
