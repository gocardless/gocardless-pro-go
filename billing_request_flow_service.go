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
type BillingRequestFlowServiceImpl struct {
	config Config
}

type BillingRequestFlowLinks struct {
	BillingRequest string `url:"billing_request,omitempty" json:"billing_request,omitempty"`
}

type BillingRequestFlowPrefilledBankAccount struct {
	AccountType string `url:"account_type,omitempty" json:"account_type,omitempty"`
}

type BillingRequestFlowPrefilledCustomer struct {
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
	PostalCode            string `url:"postal_code,omitempty" json:"postal_code,omitempty"`
	Region                string `url:"region,omitempty" json:"region,omitempty"`
	SwedishIdentityNumber string `url:"swedish_identity_number,omitempty" json:"swedish_identity_number,omitempty"`
}

// BillingRequestFlow model
type BillingRequestFlow struct {
	AuthorisationUrl          string                                  `url:"authorisation_url,omitempty" json:"authorisation_url,omitempty"`
	AutoFulfil                bool                                    `url:"auto_fulfil,omitempty" json:"auto_fulfil,omitempty"`
	CreatedAt                 string                                  `url:"created_at,omitempty" json:"created_at,omitempty"`
	ExitUri                   string                                  `url:"exit_uri,omitempty" json:"exit_uri,omitempty"`
	ExpiresAt                 string                                  `url:"expires_at,omitempty" json:"expires_at,omitempty"`
	Id                        string                                  `url:"id,omitempty" json:"id,omitempty"`
	Language                  string                                  `url:"language,omitempty" json:"language,omitempty"`
	Links                     *BillingRequestFlowLinks                `url:"links,omitempty" json:"links,omitempty"`
	LockBankAccount           bool                                    `url:"lock_bank_account,omitempty" json:"lock_bank_account,omitempty"`
	LockCurrency              bool                                    `url:"lock_currency,omitempty" json:"lock_currency,omitempty"`
	LockCustomerDetails       bool                                    `url:"lock_customer_details,omitempty" json:"lock_customer_details,omitempty"`
	PrefilledBankAccount      *BillingRequestFlowPrefilledBankAccount `url:"prefilled_bank_account,omitempty" json:"prefilled_bank_account,omitempty"`
	PrefilledCustomer         *BillingRequestFlowPrefilledCustomer    `url:"prefilled_customer,omitempty" json:"prefilled_customer,omitempty"`
	RedirectUri               string                                  `url:"redirect_uri,omitempty" json:"redirect_uri,omitempty"`
	SessionToken              string                                  `url:"session_token,omitempty" json:"session_token,omitempty"`
	ShowRedirectButtons       bool                                    `url:"show_redirect_buttons,omitempty" json:"show_redirect_buttons,omitempty"`
	ShowSuccessRedirectButton bool                                    `url:"show_success_redirect_button,omitempty" json:"show_success_redirect_button,omitempty"`
}

type BillingRequestFlowService interface {
	Create(ctx context.Context, p BillingRequestFlowCreateParams, opts ...RequestOption) (*BillingRequestFlow, error)
	Initialise(ctx context.Context, identity string, p BillingRequestFlowInitialiseParams, opts ...RequestOption) (*BillingRequestFlow, error)
}

type BillingRequestFlowCreateParamsLinks struct {
	BillingRequest string `url:"billing_request,omitempty" json:"billing_request,omitempty"`
}

type BillingRequestFlowCreateParamsPrefilledBankAccount struct {
	AccountType string `url:"account_type,omitempty" json:"account_type,omitempty"`
}

type BillingRequestFlowCreateParamsPrefilledCustomer struct {
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
	PostalCode            string `url:"postal_code,omitempty" json:"postal_code,omitempty"`
	Region                string `url:"region,omitempty" json:"region,omitempty"`
	SwedishIdentityNumber string `url:"swedish_identity_number,omitempty" json:"swedish_identity_number,omitempty"`
}

// BillingRequestFlowCreateParams parameters
type BillingRequestFlowCreateParams struct {
	AutoFulfil                bool                                                `url:"auto_fulfil,omitempty" json:"auto_fulfil,omitempty"`
	ExitUri                   string                                              `url:"exit_uri,omitempty" json:"exit_uri,omitempty"`
	Language                  string                                              `url:"language,omitempty" json:"language,omitempty"`
	Links                     BillingRequestFlowCreateParamsLinks                 `url:"links,omitempty" json:"links,omitempty"`
	LockBankAccount           bool                                                `url:"lock_bank_account,omitempty" json:"lock_bank_account,omitempty"`
	LockCurrency              bool                                                `url:"lock_currency,omitempty" json:"lock_currency,omitempty"`
	LockCustomerDetails       bool                                                `url:"lock_customer_details,omitempty" json:"lock_customer_details,omitempty"`
	PrefilledBankAccount      *BillingRequestFlowCreateParamsPrefilledBankAccount `url:"prefilled_bank_account,omitempty" json:"prefilled_bank_account,omitempty"`
	PrefilledCustomer         *BillingRequestFlowCreateParamsPrefilledCustomer    `url:"prefilled_customer,omitempty" json:"prefilled_customer,omitempty"`
	RedirectUri               string                                              `url:"redirect_uri,omitempty" json:"redirect_uri,omitempty"`
	ShowRedirectButtons       bool                                                `url:"show_redirect_buttons,omitempty" json:"show_redirect_buttons,omitempty"`
	ShowSuccessRedirectButton bool                                                `url:"show_success_redirect_button,omitempty" json:"show_success_redirect_button,omitempty"`
}

// Create
// Creates a new billing request flow.
func (s *BillingRequestFlowServiceImpl) Create(ctx context.Context, p BillingRequestFlowCreateParams, opts ...RequestOption) (*BillingRequestFlow, error) {
	uri, err := url.Parse(fmt.Sprintf(s.config.Endpoint() + "/billing_request_flows"))
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
	req = req.WithContext(ctx)
	req.Header.Set("Authorization", "Bearer "+s.config.Token())
	req.Header.Set("GoCardless-Version", "2015-07-06")
	req.Header.Set("GoCardless-Client-Library", "gocardless-pro-go")
	req.Header.Set("GoCardless-Client-Version", "3.2.0")
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
type BillingRequestFlowInitialiseParams struct {
}

// Initialise
// Returns the flow having generated a fresh session token which can be used to
// power
// integrations that manipulate the flow.
func (s *BillingRequestFlowServiceImpl) Initialise(ctx context.Context, identity string, p BillingRequestFlowInitialiseParams, opts ...RequestOption) (*BillingRequestFlow, error) {
	uri, err := url.Parse(fmt.Sprintf(s.config.Endpoint()+"/billing_request_flows/%v/actions/initialise",
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
	req.Header.Set("GoCardless-Client-Version", "3.2.0")
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
