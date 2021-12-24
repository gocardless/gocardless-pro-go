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

// BillingRequestFlowService manages billing_request_flows
type BillingRequestFlowService struct {
	endpoint string
	token    string
	client   *http.Client
}

// BillingRequestFlow model
type BillingRequestFlow struct {
	AuthorisationUrl string `url:"authorisation_url,omitempty" json:"authorisation_url,omitempty"`
	AutoFulfil       bool   `url:"auto_fulfil,omitempty" json:"auto_fulfil,omitempty"`
	CreatedAt        string `url:"created_at,omitempty" json:"created_at,omitempty"`
	ExitUri          string `url:"exit_uri,omitempty" json:"exit_uri,omitempty"`
	ExpiresAt        string `url:"expires_at,omitempty" json:"expires_at,omitempty"`
	Id               string `url:"id,omitempty" json:"id,omitempty"`
	Links            struct {
		BillingRequest string `url:"billing_request,omitempty" json:"billing_request,omitempty"`
	} `url:"links,omitempty" json:"links,omitempty"`
	LockBankAccount     bool   `url:"lock_bank_account,omitempty" json:"lock_bank_account,omitempty"`
	LockCustomerDetails bool   `url:"lock_customer_details,omitempty" json:"lock_customer_details,omitempty"`
	RedirectUri         string `url:"redirect_uri,omitempty" json:"redirect_uri,omitempty"`
	SessionToken        string `url:"session_token,omitempty" json:"session_token,omitempty"`
}

// BillingRequestFlowCreateParams parameters
type BillingRequestFlowCreateParams struct {
	AutoFulfil bool   `url:"auto_fulfil,omitempty" json:"auto_fulfil,omitempty"`
	ExitUri    string `url:"exit_uri,omitempty" json:"exit_uri,omitempty"`
	Links      struct {
		BillingRequest string `url:"billing_request,omitempty" json:"billing_request,omitempty"`
	} `url:"links,omitempty" json:"links,omitempty"`
	LockBankAccount     bool   `url:"lock_bank_account,omitempty" json:"lock_bank_account,omitempty"`
	LockCustomerDetails bool   `url:"lock_customer_details,omitempty" json:"lock_customer_details,omitempty"`
	RedirectUri         string `url:"redirect_uri,omitempty" json:"redirect_uri,omitempty"`
}

// Create
// Creates a new billing request flow.
func (s *BillingRequestFlowService) Create(ctx context.Context, p BillingRequestFlowCreateParams, opts ...RequestOption) (*BillingRequestFlow, error) {
	uri, err := url.Parse(fmt.Sprintf(s.endpoint + "/billing_request_flows"))
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
		"billing_request_flows": p,
	})
	if err != nil {
		return nil, err
	}
	body = &buf

	req, err := http.NewRequest("POST", uri.String(), body)
	if err != nil {
		return nil, err
	}
	req.WithContext(ctx)
	req.Header.Set("Authorization", "Bearer "+s.token)

	req.Header.Set("GoCardless-Version", "2015-07-06")

	req.Header.Set("GoCardless-Client-Library", "<no value>")
	req.Header.Set("GoCardless-Client-Version", "1.0.0")
	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Idempotency-Key", o.idempotencyKey)

	for key, value := range o.headers {
		req.Header.Set(key, value)
	}

	client := s.client
	if client == nil {
		client = http.DefaultClient
	}

	var result struct {
		Err                *APIError           `json:"error"`
		BillingRequestFlow *BillingRequestFlow `json:"billing_request_flows"`
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

	if result.BillingRequestFlow == nil {
		return nil, errors.New("missing result")
	}

	return result.BillingRequestFlow, nil
}

// BillingRequestFlowInitialiseParams parameters
type BillingRequestFlowInitialiseParams map[string]interface{}

// Initialise
// Returns the flow having generated a fresh session token which can be used to
// power
// integrations that manipulate the flow.
func (s *BillingRequestFlowService) Initialise(ctx context.Context, identity string, p BillingRequestFlowInitialiseParams, opts ...RequestOption) (*BillingRequestFlow, error) {
	uri, err := url.Parse(fmt.Sprintf(s.endpoint+"/billing_request_flows/%v/actions/initialise",
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
	req.WithContext(ctx)
	req.Header.Set("Authorization", "Bearer "+s.token)

	req.Header.Set("GoCardless-Version", "2015-07-06")

	req.Header.Set("GoCardless-Client-Library", "<no value>")
	req.Header.Set("GoCardless-Client-Version", "1.0.0")
	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Idempotency-Key", o.idempotencyKey)

	for key, value := range o.headers {
		req.Header.Set(key, value)
	}

	client := s.client
	if client == nil {
		client = http.DefaultClient
	}

	var result struct {
		Err                *APIError           `json:"error"`
		BillingRequestFlow *BillingRequestFlow `json:"billing_request_flows"`
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

	if result.BillingRequestFlow == nil {
		return nil, errors.New("missing result")
	}

	return result.BillingRequestFlow, nil
}
