package gocardless

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestWebhookFailsWithInvalidSignature(t *testing.T) {
	wh, err := NewWebhookHandler("testing", EventHandlerFunc(func(ctx context.Context, e Event) error {
		t.Error("unexpected call")
		return nil
	}))

	if err != nil {
		t.Fatal(err)
	}

	f, err := os.Open("testdata/webhook_request.json")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/webhook", f)

	wh.ServeHTTP(w, r)
	if w.Code != 498 {
		t.Fatalf("Expected %d, got %d", 498, w.Code)
	}
}

func TestWebhookFailsWithValidSignature(t *testing.T) {
	var called int

	wh, err := NewWebhookHandler("testing", EventHandlerFunc(func(ctx context.Context, e Event) error {
		called++
		expectedID := "EVTESTNE86TNZS"
		if e.Id != expectedID {
			t.Fatalf("Expected %q, got %q", expectedID, e.Id)
		}
		return nil
	}))

	if err != nil {
		t.Fatal(err)
	}

	f, err := os.Open("testdata/webhook_request.json")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/webhook", f)
	r.Header.Set("Webhook-Signature", "243f3efa57743c24eec7c5e10edc475b547830d256d5b583745afe319dd90936")

	wh.ServeHTTP(w, r)
	if w.Code != http.StatusNoContent {
		t.Fatalf("Expected %d, got %d", http.StatusNoContent, w.Code)
	}

	if called != 1 {
		t.Fatalf("Expected 1 call, got %d", called)
	}
}

func TestWebhookWhenHandlerFails(t *testing.T) {
	var called int

	wh, err := NewWebhookHandler("testing", EventHandlerFunc(func(ctx context.Context, e Event) error {
		called++
		expectedID := "EVTESTNE86TNZS"
		if e.Id != expectedID {
			t.Fatalf("Expected %q, got %q", expectedID, e.Id)
		}
		if called == 0 {
			return nil
		}
		return errors.New("failed")
	}))

	if err != nil {
		t.Fatal(err)
	}

	f, err := os.Open("testdata/webhook_request.json")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/webhook", f)
	r.Header.Set("Webhook-Signature", "243f3efa57743c24eec7c5e10edc475b547830d256d5b583745afe319dd90936")

	wh.ServeHTTP(w, r)
	if w.Code != http.StatusInternalServerError {
		t.Fatalf("Expected %d, got %d", http.StatusInternalServerError, w.Code)
	}

	if called != 1 {
		t.Fatalf("Expected 1 call, got %d", called)
	}
}
