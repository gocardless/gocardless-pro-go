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
	opts := WithEndpoint(url)
	token := "dummy_token"
	service, err := New(token, opts)
	if err != nil {
		t.Fatal(err)
	}
	return service, nil
}
