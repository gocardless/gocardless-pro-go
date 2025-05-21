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

// PayerAuthorisationService manages payer_authorisations
type PayerAuthorisationServiceImpl struct {
	config Config
}

type PayerAuthorisationBankAccount struct {
	AccountHolderName   string                 `url:"account_holder_name,omitempty" json:"account_holder_name,omitempty"`
	AccountNumber       string                 `url:"account_number,omitempty" json:"account_number,omitempty"`
	AccountNumberEnding string                 `url:"account_number_ending,omitempty" json:"account_number_ending,omitempty"`
	AccountNumberSuffix string                 `url:"account_number_suffix,omitempty" json:"account_number_suffix,omitempty"`
	AccountType         string                 `url:"account_type,omitempty" json:"account_type,omitempty"`
	BankCode            string                 `url:"bank_code,omitempty" json:"bank_code,omitempty"`
	BranchCode          string                 `url:"branch_code,omitempty" json:"branch_code,omitempty"`
	CountryCode         string                 `url:"country_code,omitempty" json:"country_code,omitempty"`
	Currency            string                 `url:"currency,omitempty" json:"currency,omitempty"`
	Iban                string                 `url:"iban,omitempty" json:"iban,omitempty"`
	Metadata            map[string]interface{} `url:"metadata,omitempty" json:"metadata,omitempty"`
}

type PayerAuthorisationCustomer struct {
	AddressLine1          string                 `url:"address_line1,omitempty" json:"address_line1,omitempty"`
	AddressLine2          string                 `url:"address_line2,omitempty" json:"address_line2,omitempty"`
	AddressLine3          string                 `url:"address_line3,omitempty" json:"address_line3,omitempty"`
	City                  string                 `url:"city,omitempty" json:"city,omitempty"`
	CompanyName           string                 `url:"company_name,omitempty" json:"company_name,omitempty"`
	CountryCode           string                 `url:"country_code,omitempty" json:"country_code,omitempty"`
	DanishIdentityNumber  string                 `url:"danish_identity_number,omitempty" json:"danish_identity_number,omitempty"`
	Email                 string                 `url:"email,omitempty" json:"email,omitempty"`
	FamilyName            string                 `url:"family_name,omitempty" json:"family_name,omitempty"`
	GivenName             string                 `url:"given_name,omitempty" json:"given_name,omitempty"`
	Locale                string                 `url:"locale,omitempty" json:"locale,omitempty"`
	Metadata              map[string]interface{} `url:"metadata,omitempty" json:"metadata,omitempty"`
	PostalCode            string                 `url:"postal_code,omitempty" json:"postal_code,omitempty"`
	Region                string                 `url:"region,omitempty" json:"region,omitempty"`
	SwedishIdentityNumber string                 `url:"swedish_identity_number,omitempty" json:"swedish_identity_number,omitempty"`
}

type PayerAuthorisationIncompleteFields struct {
	Field          string `url:"field,omitempty" json:"field,omitempty"`
	Message        string `url:"message,omitempty" json:"message,omitempty"`
	RequestPointer string `url:"request_pointer,omitempty" json:"request_pointer,omitempty"`
}

type PayerAuthorisationLinks struct {
	BankAccount string `url:"bank_account,omitempty" json:"bank_account,omitempty"`
	Customer    string `url:"customer,omitempty" json:"customer,omitempty"`
	Mandate     string `url:"mandate,omitempty" json:"mandate,omitempty"`
}

type PayerAuthorisationMandate struct {
	Metadata       map[string]interface{} `url:"metadata,omitempty" json:"metadata,omitempty"`
	PayerIpAddress string                 `url:"payer_ip_address,omitempty" json:"payer_ip_address,omitempty"`
	Reference      string                 `url:"reference,omitempty" json:"reference,omitempty"`
	Scheme         string                 `url:"scheme,omitempty" json:"scheme,omitempty"`
}

// PayerAuthorisation model
type PayerAuthorisation struct {
	BankAccount      *PayerAuthorisationBankAccount       `url:"bank_account,omitempty" json:"bank_account,omitempty"`
	CreatedAt        string                               `url:"created_at,omitempty" json:"created_at,omitempty"`
	Customer         *PayerAuthorisationCustomer          `url:"customer,omitempty" json:"customer,omitempty"`
	Id               string                               `url:"id,omitempty" json:"id,omitempty"`
	IncompleteFields []PayerAuthorisationIncompleteFields `url:"incomplete_fields,omitempty" json:"incomplete_fields,omitempty"`
	Links            *PayerAuthorisationLinks             `url:"links,omitempty" json:"links,omitempty"`
	Mandate          *PayerAuthorisationMandate           `url:"mandate,omitempty" json:"mandate,omitempty"`
	Status           string                               `url:"status,omitempty" json:"status,omitempty"`
}

type PayerAuthorisationService interface {
	Get(ctx context.Context, identity string, opts ...RequestOption) (*PayerAuthorisation, error)
	Create(ctx context.Context, p PayerAuthorisationCreateParams, opts ...RequestOption) (*PayerAuthorisation, error)
	Update(ctx context.Context, identity string, p PayerAuthorisationUpdateParams, opts ...RequestOption) (*PayerAuthorisation, error)
	Submit(ctx context.Context, identity string, opts ...RequestOption) (*PayerAuthorisation, error)
	Confirm(ctx context.Context, identity string, opts ...RequestOption) (*PayerAuthorisation, error)
}

// Get
// Retrieves the details of a single existing Payer Authorisation. It can be
// used for polling the status of a Payer Authorisation.
func (s *PayerAuthorisationServiceImpl) Get(ctx context.Context, identity string, opts ...RequestOption) (*PayerAuthorisation, error) {
	uri, err := url.Parse(fmt.Sprintf(s.config.Endpoint()+"/payer_authorisations/%v",
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
	req.Header.Set("GoCardless-Client-Version", "5.0.0")
	req.Header.Set("User-Agent", userAgent)

	for key, value := range o.headers {
		req.Header.Set(key, value)
	}

	client := s.config.Client()
	if client == nil {
		client = http.DefaultClient
	}

	var result struct {
		Err                *APIError           `json:"error"`
		PayerAuthorisation *PayerAuthorisation `json:"payer_authorisations"`
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

	if result.PayerAuthorisation == nil {
		return nil, errors.New("missing result")
	}

	return result.PayerAuthorisation, nil
}

type PayerAuthorisationCreateParamsBankAccount struct {
	AccountHolderName   string                 `url:"account_holder_name,omitempty" json:"account_holder_name,omitempty"`
	AccountNumber       string                 `url:"account_number,omitempty" json:"account_number,omitempty"`
	AccountNumberEnding string                 `url:"account_number_ending,omitempty" json:"account_number_ending,omitempty"`
	AccountNumberSuffix string                 `url:"account_number_suffix,omitempty" json:"account_number_suffix,omitempty"`
	AccountType         string                 `url:"account_type,omitempty" json:"account_type,omitempty"`
	BankCode            string                 `url:"bank_code,omitempty" json:"bank_code,omitempty"`
	BranchCode          string                 `url:"branch_code,omitempty" json:"branch_code,omitempty"`
	CountryCode         string                 `url:"country_code,omitempty" json:"country_code,omitempty"`
	Currency            string                 `url:"currency,omitempty" json:"currency,omitempty"`
	Iban                string                 `url:"iban,omitempty" json:"iban,omitempty"`
	Metadata            map[string]interface{} `url:"metadata,omitempty" json:"metadata,omitempty"`
}

type PayerAuthorisationCreateParamsCustomer struct {
	AddressLine1          string                 `url:"address_line1,omitempty" json:"address_line1,omitempty"`
	AddressLine2          string                 `url:"address_line2,omitempty" json:"address_line2,omitempty"`
	AddressLine3          string                 `url:"address_line3,omitempty" json:"address_line3,omitempty"`
	City                  string                 `url:"city,omitempty" json:"city,omitempty"`
	CompanyName           string                 `url:"company_name,omitempty" json:"company_name,omitempty"`
	CountryCode           string                 `url:"country_code,omitempty" json:"country_code,omitempty"`
	DanishIdentityNumber  string                 `url:"danish_identity_number,omitempty" json:"danish_identity_number,omitempty"`
	Email                 string                 `url:"email,omitempty" json:"email,omitempty"`
	FamilyName            string                 `url:"family_name,omitempty" json:"family_name,omitempty"`
	GivenName             string                 `url:"given_name,omitempty" json:"given_name,omitempty"`
	Locale                string                 `url:"locale,omitempty" json:"locale,omitempty"`
	Metadata              map[string]interface{} `url:"metadata,omitempty" json:"metadata,omitempty"`
	PostalCode            string                 `url:"postal_code,omitempty" json:"postal_code,omitempty"`
	Region                string                 `url:"region,omitempty" json:"region,omitempty"`
	SwedishIdentityNumber string                 `url:"swedish_identity_number,omitempty" json:"swedish_identity_number,omitempty"`
}

type PayerAuthorisationCreateParamsMandate struct {
	Metadata       map[string]interface{} `url:"metadata,omitempty" json:"metadata,omitempty"`
	PayerIpAddress string                 `url:"payer_ip_address,omitempty" json:"payer_ip_address,omitempty"`
	Reference      string                 `url:"reference,omitempty" json:"reference,omitempty"`
	Scheme         string                 `url:"scheme,omitempty" json:"scheme,omitempty"`
}

// PayerAuthorisationCreateParams parameters
type PayerAuthorisationCreateParams struct {
	BankAccount PayerAuthorisationCreateParamsBankAccount `url:"bank_account,omitempty" json:"bank_account,omitempty"`
	Customer    PayerAuthorisationCreateParamsCustomer    `url:"customer,omitempty" json:"customer,omitempty"`
	Mandate     PayerAuthorisationCreateParamsMandate     `url:"mandate,omitempty" json:"mandate,omitempty"`
}

// Create
// Creates a Payer Authorisation. The resource is saved to the database even if
// incomplete. An empty array of incomplete_fields means that the resource is
// valid. The ID of the resource is used for the other actions. This endpoint
// has been designed this way so you do not need to save any payer data on your
// servers or the browser while still being able to implement a progressive
// solution, such as a multi-step form.
func (s *PayerAuthorisationServiceImpl) Create(ctx context.Context, p PayerAuthorisationCreateParams, opts ...RequestOption) (*PayerAuthorisation, error) {
	uri, err := url.Parse(fmt.Sprintf(s.config.Endpoint() + "/payer_authorisations"))
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
		"payer_authorisations": p,
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
	req.Header.Set("GoCardless-Client-Version", "5.0.0")
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
		PayerAuthorisation *PayerAuthorisation `json:"payer_authorisations"`
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

	if result.PayerAuthorisation == nil {
		return nil, errors.New("missing result")
	}

	return result.PayerAuthorisation, nil
}

type PayerAuthorisationUpdateParamsBankAccount struct {
	AccountHolderName   string                 `url:"account_holder_name,omitempty" json:"account_holder_name,omitempty"`
	AccountNumber       string                 `url:"account_number,omitempty" json:"account_number,omitempty"`
	AccountNumberEnding string                 `url:"account_number_ending,omitempty" json:"account_number_ending,omitempty"`
	AccountNumberSuffix string                 `url:"account_number_suffix,omitempty" json:"account_number_suffix,omitempty"`
	AccountType         string                 `url:"account_type,omitempty" json:"account_type,omitempty"`
	BankCode            string                 `url:"bank_code,omitempty" json:"bank_code,omitempty"`
	BranchCode          string                 `url:"branch_code,omitempty" json:"branch_code,omitempty"`
	CountryCode         string                 `url:"country_code,omitempty" json:"country_code,omitempty"`
	Currency            string                 `url:"currency,omitempty" json:"currency,omitempty"`
	Iban                string                 `url:"iban,omitempty" json:"iban,omitempty"`
	Metadata            map[string]interface{} `url:"metadata,omitempty" json:"metadata,omitempty"`
}

type PayerAuthorisationUpdateParamsCustomer struct {
	AddressLine1          string                 `url:"address_line1,omitempty" json:"address_line1,omitempty"`
	AddressLine2          string                 `url:"address_line2,omitempty" json:"address_line2,omitempty"`
	AddressLine3          string                 `url:"address_line3,omitempty" json:"address_line3,omitempty"`
	City                  string                 `url:"city,omitempty" json:"city,omitempty"`
	CompanyName           string                 `url:"company_name,omitempty" json:"company_name,omitempty"`
	CountryCode           string                 `url:"country_code,omitempty" json:"country_code,omitempty"`
	DanishIdentityNumber  string                 `url:"danish_identity_number,omitempty" json:"danish_identity_number,omitempty"`
	Email                 string                 `url:"email,omitempty" json:"email,omitempty"`
	FamilyName            string                 `url:"family_name,omitempty" json:"family_name,omitempty"`
	GivenName             string                 `url:"given_name,omitempty" json:"given_name,omitempty"`
	Locale                string                 `url:"locale,omitempty" json:"locale,omitempty"`
	Metadata              map[string]interface{} `url:"metadata,omitempty" json:"metadata,omitempty"`
	PostalCode            string                 `url:"postal_code,omitempty" json:"postal_code,omitempty"`
	Region                string                 `url:"region,omitempty" json:"region,omitempty"`
	SwedishIdentityNumber string                 `url:"swedish_identity_number,omitempty" json:"swedish_identity_number,omitempty"`
}

type PayerAuthorisationUpdateParamsMandate struct {
	Metadata       map[string]interface{} `url:"metadata,omitempty" json:"metadata,omitempty"`
	PayerIpAddress string                 `url:"payer_ip_address,omitempty" json:"payer_ip_address,omitempty"`
	Reference      string                 `url:"reference,omitempty" json:"reference,omitempty"`
	Scheme         string                 `url:"scheme,omitempty" json:"scheme,omitempty"`
}

// PayerAuthorisationUpdateParams parameters
type PayerAuthorisationUpdateParams struct {
	BankAccount PayerAuthorisationUpdateParamsBankAccount `url:"bank_account,omitempty" json:"bank_account,omitempty"`
	Customer    PayerAuthorisationUpdateParamsCustomer    `url:"customer,omitempty" json:"customer,omitempty"`
	Mandate     PayerAuthorisationUpdateParamsMandate     `url:"mandate,omitempty" json:"mandate,omitempty"`
}

// Update
// Updates a Payer Authorisation. Updates the Payer Authorisation with the
// request data. Can be invoked as many times as needed. Only fields present in
// the request will be modified. An empty array of incomplete_fields means that
// the resource is valid. This endpoint has been designed this way so you do not
// need to save any payer data on your servers or the browser while still being
// able to implement a progressive solution, such a multi-step form. <p
// class="notice"> Note that in order to update the `metadata` attribute values
// it must be sent completely as it overrides the previously existing values.
// </p>
func (s *PayerAuthorisationServiceImpl) Update(ctx context.Context, identity string, p PayerAuthorisationUpdateParams, opts ...RequestOption) (*PayerAuthorisation, error) {
	uri, err := url.Parse(fmt.Sprintf(s.config.Endpoint()+"/payer_authorisations/%v",
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
		"payer_authorisations": p,
	})
	if err != nil {
		return nil, err
	}
	body = &buf

	req, err := http.NewRequest("PUT", uri.String(), body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	req.Header.Set("Authorization", "Bearer "+s.config.Token())
	req.Header.Set("GoCardless-Version", "2015-07-06")
	req.Header.Set("GoCardless-Client-Library", "gocardless-pro-go")
	req.Header.Set("GoCardless-Client-Version", "5.0.0")
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
		PayerAuthorisation *PayerAuthorisation `json:"payer_authorisations"`
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

	if result.PayerAuthorisation == nil {
		return nil, errors.New("missing result")
	}

	return result.PayerAuthorisation, nil
}

// Submit
// Submits all the data previously pushed to this PayerAuthorisation for
// verification. This time, a 200 HTTP status is returned if the resource is
// valid and a 422 error response in case of validation errors. After it is
// successfully submitted, the Payer Authorisation can no longer be edited.
func (s *PayerAuthorisationServiceImpl) Submit(ctx context.Context, identity string, opts ...RequestOption) (*PayerAuthorisation, error) {
	uri, err := url.Parse(fmt.Sprintf(s.config.Endpoint()+"/payer_authorisations/%v/actions/submit",
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

	req, err := http.NewRequest("POST", uri.String(), body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	req.Header.Set("Authorization", "Bearer "+s.config.Token())
	req.Header.Set("GoCardless-Version", "2015-07-06")
	req.Header.Set("GoCardless-Client-Library", "gocardless-pro-go")
	req.Header.Set("GoCardless-Client-Version", "5.0.0")
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
		PayerAuthorisation *PayerAuthorisation `json:"payer_authorisations"`
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

	if result.PayerAuthorisation == nil {
		return nil, errors.New("missing result")
	}

	return result.PayerAuthorisation, nil
}

// Confirm
// Confirms the Payer Authorisation, indicating that the resources are ready to
// be created.
// A Payer Authorisation cannot be confirmed if it hasn't been submitted yet.
//
// <p class="notice">
//
//	The main use of the confirm endpoint is to enable integrators to
//
// acknowledge the end of the setup process.
//
//	They might want to make the payers go through some other steps after they
//
// go through our flow or make them go through the necessary verification
// mechanism (upcoming feature).
// </p>
func (s *PayerAuthorisationServiceImpl) Confirm(ctx context.Context, identity string, opts ...RequestOption) (*PayerAuthorisation, error) {
	uri, err := url.Parse(fmt.Sprintf(s.config.Endpoint()+"/payer_authorisations/%v/actions/confirm",
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

	req, err := http.NewRequest("POST", uri.String(), body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	req.Header.Set("Authorization", "Bearer "+s.config.Token())
	req.Header.Set("GoCardless-Version", "2015-07-06")
	req.Header.Set("GoCardless-Client-Library", "gocardless-pro-go")
	req.Header.Set("GoCardless-Client-Version", "5.0.0")
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
		PayerAuthorisation *PayerAuthorisation `json:"payer_authorisations"`
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

	if result.PayerAuthorisation == nil {
		return nil, errors.New("missing result")
	}

	return result.PayerAuthorisation, nil
}
