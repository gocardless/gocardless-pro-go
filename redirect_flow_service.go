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

// RedirectFlowService manages redirect_flows
type RedirectFlowServiceImpl struct {
	config Config
}

type RedirectFlowLinks struct {
	BillingRequest      string `url:"billing_request,omitempty" json:"billing_request,omitempty"`
	Creditor            string `url:"creditor,omitempty" json:"creditor,omitempty"`
	Customer            string `url:"customer,omitempty" json:"customer,omitempty"`
	CustomerBankAccount string `url:"customer_bank_account,omitempty" json:"customer_bank_account,omitempty"`
	Mandate             string `url:"mandate,omitempty" json:"mandate,omitempty"`
}

// RedirectFlow model
type RedirectFlow struct {
	ConfirmationUrl    string                 `url:"confirmation_url,omitempty" json:"confirmation_url,omitempty"`
	CreatedAt          string                 `url:"created_at,omitempty" json:"created_at,omitempty"`
	Description        string                 `url:"description,omitempty" json:"description,omitempty"`
	Id                 string                 `url:"id,omitempty" json:"id,omitempty"`
	Links              *RedirectFlowLinks     `url:"links,omitempty" json:"links,omitempty"`
	MandateReference   string                 `url:"mandate_reference,omitempty" json:"mandate_reference,omitempty"`
	Metadata           map[string]interface{} `url:"metadata,omitempty" json:"metadata,omitempty"`
	RedirectUrl        string                 `url:"redirect_url,omitempty" json:"redirect_url,omitempty"`
	Scheme             string                 `url:"scheme,omitempty" json:"scheme,omitempty"`
	SessionToken       string                 `url:"session_token,omitempty" json:"session_token,omitempty"`
	SuccessRedirectUrl string                 `url:"success_redirect_url,omitempty" json:"success_redirect_url,omitempty"`
}

type RedirectFlowService interface {
	Create(ctx context.Context, p RedirectFlowCreateParams, opts ...RequestOption) (*RedirectFlow, error)
	Get(ctx context.Context, identity string, opts ...RequestOption) (*RedirectFlow, error)
	Complete(ctx context.Context, identity string, p RedirectFlowCompleteParams, opts ...RequestOption) (*RedirectFlow, error)
}

type RedirectFlowCreateParamsLinks struct {
	Creditor string `url:"creditor,omitempty" json:"creditor,omitempty"`
}

type RedirectFlowCreateParamsPrefilledBankAccount struct {
	AccountType string `url:"account_type,omitempty" json:"account_type,omitempty"`
}

type RedirectFlowCreateParamsPrefilledCustomer struct {
	AddressLine1          string `url:"address_line1,omitempty" json:"address_line1,omitempty"`
	AddressLine2          string `url:"address_line2,omitempty" json:"address_line2,omitempty"`
	AddressLine3          string `url:"address_line3,omitempty" json:"address_line3,omitempty"`
	City                  string `url:"city,omitempty" json:"city,omitempty"`
	CompanyName           string `url:"company_name,omitempty" json:"company_name,omitempty"`
	CountryCode           string `url:"country_code,omitempty" json:"country_code,omitempty"`
	DanishIdentityNumber  string `url:"danish_identity_number,omitempty" json:"danish_identity_number,omitempty"`
	Email                 string `url:"email,omitempty" json:"email,omitempty"`
	FamilyName            string `url:"family_name,omitempty" json:"family_name,omitempty"`
	GivenName             string `url:"given_name,omitempty" json:"given_name,omitempty"`
	Language              string `url:"language,omitempty" json:"language,omitempty"`
	PhoneNumber           string `url:"phone_number,omitempty" json:"phone_number,omitempty"`
	PostalCode            string `url:"postal_code,omitempty" json:"postal_code,omitempty"`
	Region                string `url:"region,omitempty" json:"region,omitempty"`
	SwedishIdentityNumber string `url:"swedish_identity_number,omitempty" json:"swedish_identity_number,omitempty"`
}

// RedirectFlowCreateParams parameters
type RedirectFlowCreateParams struct {
	Description          string                                        `url:"description,omitempty" json:"description,omitempty"`
	Links                *RedirectFlowCreateParamsLinks                `url:"links,omitempty" json:"links,omitempty"`
	Metadata             map[string]interface{}                        `url:"metadata,omitempty" json:"metadata,omitempty"`
	PrefilledBankAccount *RedirectFlowCreateParamsPrefilledBankAccount `url:"prefilled_bank_account,omitempty" json:"prefilled_bank_account,omitempty"`
	PrefilledCustomer    *RedirectFlowCreateParamsPrefilledCustomer    `url:"prefilled_customer,omitempty" json:"prefilled_customer,omitempty"`
	Scheme               string                                        `url:"scheme,omitempty" json:"scheme,omitempty"`
	SessionToken         string                                        `url:"session_token,omitempty" json:"session_token,omitempty"`
	SuccessRedirectUrl   string                                        `url:"success_redirect_url,omitempty" json:"success_redirect_url,omitempty"`
}

// Create
// Creates a redirect flow object which can then be used to redirect your
// customer to the GoCardless hosted payment pages.
func (s *RedirectFlowServiceImpl) Create(ctx context.Context, p RedirectFlowCreateParams, opts ...RequestOption) (*RedirectFlow, error) {
	uri, err := url.Parse(fmt.Sprintf(s.config.Endpoint() + "/redirect_flows"))
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
		"redirect_flows": p,
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
	req.Header.Set("GoCardless-Client-Version", "3.5.0")
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
		Err          *APIError     `json:"error"`
		RedirectFlow *RedirectFlow `json:"redirect_flows"`
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

	if result.RedirectFlow == nil {
		return nil, errors.New("missing result")
	}

	return result.RedirectFlow, nil
}

// Get
// Returns all details about a single redirect flow
func (s *RedirectFlowServiceImpl) Get(ctx context.Context, identity string, opts ...RequestOption) (*RedirectFlow, error) {
	uri, err := url.Parse(fmt.Sprintf(s.config.Endpoint()+"/redirect_flows/%v",
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

	var body io.Reader

	req, err := http.NewRequest("GET", uri.String(), body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	req.Header.Set("Authorization", "Bearer "+s.config.Token())
	req.Header.Set("GoCardless-Version", "2015-07-06")
	req.Header.Set("GoCardless-Client-Library", "gocardless-pro-go")
	req.Header.Set("GoCardless-Client-Version", "3.5.0")
	req.Header.Set("User-Agent", userAgent)

	for key, value := range o.headers {
		req.Header.Set(key, value)
	}

	client := s.config.Client()
	if client == nil {
		client = http.DefaultClient
	}

	var result struct {
		Err          *APIError     `json:"error"`
		RedirectFlow *RedirectFlow `json:"redirect_flows"`
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

	if result.RedirectFlow == nil {
		return nil, errors.New("missing result")
	}

	return result.RedirectFlow, nil
}

// RedirectFlowCompleteParams parameters
type RedirectFlowCompleteParams struct {
	SessionToken string `url:"session_token,omitempty" json:"session_token,omitempty"`
}

// Complete
// This creates a [customer](#core-endpoints-customers), [customer bank
// account](#core-endpoints-customer-bank-accounts), and
// [mandate](#core-endpoints-mandates) using the details supplied by your
// customer and returns the ID of the created mandate.
//
// This will return a `redirect_flow_incomplete` error if your customer has not
// yet been redirected back to your site, and a
// `redirect_flow_already_completed` error if your integration has already
// completed this flow. It will return a `bad_request` error if the
// `session_token` differs to the one supplied when the redirect flow was
// created.
func (s *RedirectFlowServiceImpl) Complete(ctx context.Context, identity string, p RedirectFlowCompleteParams, opts ...RequestOption) (*RedirectFlow, error) {
	uri, err := url.Parse(fmt.Sprintf(s.config.Endpoint()+"/redirect_flows/%v/actions/complete",
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
	req.Header.Set("GoCardless-Client-Version", "3.5.0")
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
		Err          *APIError     `json:"error"`
		RedirectFlow *RedirectFlow `json:"redirect_flows"`
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

	if result.RedirectFlow == nil {
		return nil, errors.New("missing result")
	}

	return result.RedirectFlow, nil
}
