package gocardless

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"
)

func TestFlexError_UnmarshalJSON_String(t *testing.T) {
	data := `"something went wrong"`
	var fe FlexError
	if err := json.Unmarshal([]byte(data), &fe); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if fe.Msg == nil {
		t.Fatal("expected Msg to be set")
	}
	if *fe.Msg != "something went wrong" {
		t.Fatalf("expected Msg to be %q, got %q", "something went wrong", *fe.Msg)
	}
	if fe.Err != nil {
		t.Fatal("expected Err to be nil")
	}
}

func TestFlexError_UnmarshalJSON_Object(t *testing.T) {
	data := `{"message":"bad request","type":"validation_error","code":400}`
	var fe FlexError
	if err := json.Unmarshal([]byte(data), &fe); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if fe.Err == nil {
		t.Fatal("expected Err to be set")
	}
	if fe.Err.Message != "bad request" {
		t.Fatalf("expected Message %q, got %q", "bad request", fe.Err.Message)
	}
	if fe.Err.Code != 400 {
		t.Fatalf("expected Code 400, got %d", fe.Err.Code)
	}
	if fe.Msg != nil {
		t.Fatal("expected Msg to be nil")
	}
}

func TestFlexError_UnmarshalJSON_Null(t *testing.T) {
	data := `null`
	var fe FlexError
	if err := json.Unmarshal([]byte(data), &fe); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if fe.Msg != nil || fe.Err != nil {
		t.Fatal("expected both Msg and Err to be nil for null input")
	}
}

func TestResponseErr_Success(t *testing.T) {
	resp := &http.Response{StatusCode: 200}
	if err := responseErr(resp); err != nil {
		t.Fatalf("expected nil error for 200, got %v", err)
	}
}

func TestResponseErr_StringError(t *testing.T) {
	body := `"internal server error"`
	resp := &http.Response{
		StatusCode: 500,
		Status:     "500 Internal Server Error",
		Body:       io.NopCloser(strings.NewReader(body)),
	}

	err := responseErr(resp)
	if err == nil {
		t.Fatal("expected error for 500 response")
	}

	var re *responseError
	if !errors.As(err, &re) {
		t.Fatalf("expected *responseError, got %T", err)
	}

	var apiErr *APIError
	if !errors.As(re.err, &apiErr) {
		t.Fatalf("expected string to be wrapped as *APIError, got %T", re.err)
	}
	if apiErr.Message != "internal server error" {
		t.Fatalf("expected message %q, got %q", "internal server error", apiErr.Message)
	}
	if apiErr.Code != 500 {
		t.Fatalf("expected code 500, got %d", apiErr.Code)
	}
}

func TestResponseErr_ObjectError(t *testing.T) {
	body := `{"message":"not found","type":"invalid_api_usage","code":404}`
	resp := &http.Response{
		StatusCode: 404,
		Status:     "404 Not Found",
		Body:       io.NopCloser(strings.NewReader(body)),
	}

	err := responseErr(resp)
	if err == nil {
		t.Fatal("expected error for 404 response")
	}

	var re *responseError
	if !errors.As(err, &re) {
		t.Fatalf("expected *responseError, got %T", err)
	}

	var apiErr *APIError
	if !errors.As(re.err, &apiErr) {
		t.Fatalf("expected *APIError cause, got %T", re.err)
	}
	if apiErr.Message != "not found" {
		t.Fatalf("expected message %q, got %q", "not found", apiErr.Message)
	}
}

func TestResponseError_Temporary(t *testing.T) {
	tests := []struct {
		code int
		want bool
	}{
		{400, false},
		{404, false},
		{429, true},
		{500, true},
		{502, true},
	}

	for _, tt := range tests {
		re := &responseError{
			res: &http.Response{StatusCode: tt.code},
		}
		if got := re.Temporary(); got != tt.want {
			t.Errorf("Temporary() for status %d: got %v, want %v", tt.code, got, tt.want)
		}
	}
}
