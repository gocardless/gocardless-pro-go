package gocardless

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/google/go-querystring/query"
)

var _ = query.Values
var _ = bytes.NewBuffer
var _ = json.NewDecoder
var _ = errors.New

// CustomerNotificationService manages customer_notifications
type CustomerNotificationServiceImpl struct {
	config Config
}

type CustomerNotificationLinks struct {
	Customer     string `url:"customer,omitempty" json:"customer,omitempty"`
	Event        string `url:"event,omitempty" json:"event,omitempty"`
	Mandate      string `url:"mandate,omitempty" json:"mandate,omitempty"`
	Payment      string `url:"payment,omitempty" json:"payment,omitempty"`
	Refund       string `url:"refund,omitempty" json:"refund,omitempty"`
	Subscription string `url:"subscription,omitempty" json:"subscription,omitempty"`
}

// CustomerNotification model
type CustomerNotification struct {
	ActionTaken   string                     `url:"action_taken,omitempty" json:"action_taken,omitempty"`
	ActionTakenAt string                     `url:"action_taken_at,omitempty" json:"action_taken_at,omitempty"`
	ActionTakenBy string                     `url:"action_taken_by,omitempty" json:"action_taken_by,omitempty"`
	Id            string                     `url:"id,omitempty" json:"id,omitempty"`
	Links         *CustomerNotificationLinks `url:"links,omitempty" json:"links,omitempty"`
	Type          string                     `url:"type,omitempty" json:"type,omitempty"`
}

type CustomerNotificationService interface {
	Handle(ctx context.Context, identity string, p CustomerNotificationHandleParams, opts ...RequestOption) (*CustomerNotification, error)
}

// CustomerNotificationHandleParams parameters
type CustomerNotificationHandleParams struct {
}

// Handle
// "Handling" a notification means that you have sent the notification yourself
// (and
// don't want GoCardless to send it).
// If the notification has already been actioned, or the deadline to notify has
// passed,
// this endpoint will return an `already_actioned` error and you should not take
// further action. This endpoint takes no additional parameters.
//
func (s *CustomerNotificationServiceImpl) Handle(ctx context.Context, identity string, p CustomerNotificationHandleParams, opts ...RequestOption) (*CustomerNotification, error) {
	uri, err := url.Parse(fmt.Sprintf(s.config.Endpoint()+"/customer_notifications/%v/actions/handle",
		identity))
	if err != nil {
		return nil, err
	}

	o := &requestOptions{
		retries: 3,
	}
	for _, opt := range opts {
		err := opt(o)
		if err != nil {
			return nil, err
		}
	}
	if o.idempotencyKey == "" {
		o.idempotencyKey = NewIdempotencyKey()
	}

	var body io.Reader

	var buf bytes.Buffer
	err = json.NewEncoder(&buf).Encode(map[string]interface{}{
		"data": p,
	})
	if err != nil {
		return nil, err
	}
	body = &buf

	req, err := http.NewRequest("POST", uri.String(), body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	req.Header.Set("Authorization", "Bearer "+s.config.Token())
	req.Header.Set("GoCardless-Version", "2015-07-06")
	req.Header.Set("GoCardless-Client-Library", "gocardless-pro-go")
	req.Header.Set("GoCardless-Client-Version", "2.5.0")
	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Idempotency-Key", o.idempotencyKey)

	for key, value := range o.headers {
		req.Header.Set(key, value)
	}

	client := s.config.Client()
	if client == nil {
		client = http.DefaultClient
	}

	var result struct {
		Err                  *APIError             `json:"error"`
		CustomerNotification *CustomerNotification `json:"customer_notifications"`
	}

	err = try(o.retries, func() error {
		res, err := client.Do(req)
		if err != nil {
			return err
		}
		defer res.Body.Close()

		err = responseErr(res)
		if err != nil {
			return err
		}

		err = json.NewDecoder(res.Body).Decode(&result)
		if err != nil {
			return err
		}

		if result.Err != nil {
			return result.Err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	if result.CustomerNotification == nil {
		return nil, errors.New("missing result")
	}

	return result.CustomerNotification, nil
}
