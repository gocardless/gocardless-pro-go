package gocardless

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

// EventHandler is the interface that must be implemented to handle events from a webhook.
type EventHandler interface {
	HandleEvent(context.Context, Event) error
}

// EventHandlerFunc can be used to convert a function into an EventHandler
type EventHandlerFunc func(context.Context, Event) error

// HandleEvent will call the EventHandlerFunc function
func (h EventHandlerFunc) HandleEvent(ctx context.Context, e Event) error {
	return h(ctx, e)
}

// WebhookHandler allows you to process incoming events from webhooks.
type WebhookHandler struct {
	EventHandler
	secret string
}

// NewWebhookHandler instantiates a WebhookHandler which can be mounted as a net/http Handler.
func NewWebhookHandler(secret string, h EventHandler) (*WebhookHandler, error) {
	if secret == "" {
		return nil, errors.New("missing secret")
	}
	return &WebhookHandler{
		EventHandler: h,
		secret:       secret,
	}, nil
}

// ServeHTTP processes incoming webhooks and dispatches events to the corresponsing handlers.
func (h *WebhookHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	sig, err := hex.DecodeString(r.Header.Get("Webhook-Signature"))
	if len(sig) == 0 || err != nil {
		http.Error(w, "invalid signature", 498)
		return
	}
	hash := hmac.New(sha256.New, []byte(h.secret))
	body := io.TeeReader(r.Body, hash)

	var events struct {
		Events []Event `json:"events"`
	}
	err = json.NewDecoder(body).Decode(&events)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if !hmac.Equal(sig, hash.Sum(nil)) {
		http.Error(w, "invalid signature", 498)
		return
	}

	for _, event := range events.Events {
		err := h.HandleEvent(r.Context(), event)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusNoContent)
}
