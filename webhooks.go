package gocardless

import (
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
	HandleEvent(Event) error
}

// EventWithMetaHandler is an optional interface that can be implemented to receive
// the webhook ID along with each event.
type EventWithMetaHandler interface {
	HandleEventWithMeta(event Event, webhookID string) error
}

// WebhookParseResult contains the parsed events and metadata from a webhook.
type WebhookParseResult struct {
	Events    []Event
	WebhookID string
}

// EventHandlerFunc can be used to convert a function into an EventHandler
type EventHandlerFunc func(Event) error

// HandleEvent will call the EventHandlerFunc function
func (h EventHandlerFunc) HandleEvent(e Event) error {
	return h(e)
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
	if len(sig) == 0 {
		http.Error(w, "invalid signature", 498)
		return
	}
	hash := hmac.New(sha256.New, []byte(h.secret))
	body := io.TeeReader(r.Body, hash)

	var webhook struct {
		Events []Event `json:"events"`
		Meta   struct {
			WebhookID string `json:"webhook_id"`
		} `json:"meta"`
	}
	err = json.NewDecoder(body).Decode(&webhook)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if !hmac.Equal(sig, hash.Sum(nil)) {
		http.Error(w, "invalid signature", 498)
		return
	}

	// Check if handler implements EventWithMetaHandler
	if metaHandler, ok := h.EventHandler.(EventWithMetaHandler); ok {
		for _, event := range webhook.Events {
			err := metaHandler.HandleEventWithMeta(event, webhook.Meta.WebhookID)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
	} else {
		for _, event := range webhook.Events {
			err := h.HandleEvent(event)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
	}

	w.WriteHeader(http.StatusNoContent)
}

// ParseWebhook validates the signature and parses a webhook body, returning
// the events and webhook ID. This is useful when you need direct access to
// the parsed webhook data outside of an HTTP handler context.
func ParseWebhook(body []byte, secret string, signatureHeader string) (*WebhookParseResult, error) {
	sig, err := hex.DecodeString(signatureHeader)
	if err != nil || len(sig) == 0 {
		return nil, errors.New("invalid signature")
	}

	hash := hmac.New(sha256.New, []byte(secret))
	hash.Write(body)

	if !hmac.Equal(sig, hash.Sum(nil)) {
		return nil, errors.New("invalid signature")
	}

	var webhook struct {
		Events []Event `json:"events"`
		Meta   struct {
			WebhookID string `json:"webhook_id"`
		} `json:"meta"`
	}
	err = json.Unmarshal(body, &webhook)
	if err != nil {
		return nil, err
	}

	return &WebhookParseResult{
		Events:    webhook.Events,
		WebhookID: webhook.Meta.WebhookID,
	}, nil
}
