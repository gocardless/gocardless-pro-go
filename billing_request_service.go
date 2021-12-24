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

// BillingRequestService manages billing_requests
type BillingRequestService struct {
	endpoint string
	token    string
	client   *http.Client
}

// BillingRequest model
type BillingRequest struct {
	Actions []struct {
		BankAuthorisation struct {
			Adapter             string `url:"adapter,omitempty" json:"adapter,omitempty"`
			AuthorisationType   string `url:"authorisation_type,omitempty" json:"authorisation_type,omitempty"`
			RequiresInstitution bool   `url:"requires_institution,omitempty" json:"requires_institution,omitempty"`
		} `url:"bank_authorisation,omitempty" json:"bank_authorisation,omitempty"`
		CollectCustomerDetails struct {
			DefaultCountryCode string `url:"default_country_code,omitempty" json:"default_country_code,omitempty"`
		} `url:"collect_customer_details,omitempty" json:"collect_customer_details,omitempty"`
		CompletesActions []string `url:"completes_actions,omitempty" json:"completes_actions,omitempty"`
		Required         bool     `url:"required,omitempty" json:"required,omitempty"`
		RequiresActions  []string `url:"requires_actions,omitempty" json:"requires_actions,omitempty"`
		Status           string   `url:"status,omitempty" json:"status,omitempty"`
		Type             string   `url:"type,omitempty" json:"type,omitempty"`
	} `url:"actions,omitempty" json:"actions,omitempty"`
	CreatedAt string `url:"created_at,omitempty" json:"created_at,omitempty"`
	Id        string `url:"id,omitempty" json:"id,omitempty"`
	Links     struct {
		BankAuthorisation     string `url:"bank_authorisation,omitempty" json:"bank_authorisation,omitempty"`
		Creditor              string `url:"creditor,omitempty" json:"creditor,omitempty"`
		Customer              string `url:"customer,omitempty" json:"customer,omitempty"`
		CustomerBankAccount   string `url:"customer_bank_account,omitempty" json:"customer_bank_account,omitempty"`
		CustomerBillingDetail string `url:"customer_billing_detail,omitempty" json:"customer_billing_detail,omitempty"`
		MandateRequest        string `url:"mandate_request,omitempty" json:"mandate_request,omitempty"`
		MandateRequestMandate string `url:"mandate_request_mandate,omitempty" json:"mandate_request_mandate,omitempty"`
		PaymentRequest        string `url:"payment_request,omitempty" json:"payment_request,omitempty"`
		PaymentRequestPayment string `url:"payment_request_payment,omitempty" json:"payment_request_payment,omitempty"`
	} `url:"links,omitempty" json:"links,omitempty"`
	MandateRequest struct {
		Currency string `url:"currency,omitempty" json:"currency,omitempty"`
		Links    struct {
			Mandate string `url:"mandate,omitempty" json:"mandate,omitempty"`
		} `url:"links,omitempty" json:"links,omitempty"`
		Scheme string `url:"scheme,omitempty" json:"scheme,omitempty"`
		Verify string `url:"verify,omitempty" json:"verify,omitempty"`
	} `url:"mandate_request,omitempty" json:"mandate_request,omitempty"`
	Metadata       map[string]interface{} `url:"metadata,omitempty" json:"metadata,omitempty"`
	PaymentRequest struct {
		Amount      int    `url:"amount,omitempty" json:"amount,omitempty"`
		AppFee      int    `url:"app_fee,omitempty" json:"app_fee,omitempty"`
		Currency    string `url:"currency,omitempty" json:"currency,omitempty"`
		Description string `url:"description,omitempty" json:"description,omitempty"`
		Links       struct {
			Payment string `url:"payment,omitempty" json:"payment,omitempty"`
		} `url:"links,omitempty" json:"links,omitempty"`
		Scheme string `url:"scheme,omitempty" json:"scheme,omitempty"`
	} `url:"payment_request,omitempty" json:"payment_request,omitempty"`
	Resources struct {
		Customer struct {
			CompanyName string                 `url:"company_name,omitempty" json:"company_name,omitempty"`
			CreatedAt   string                 `url:"created_at,omitempty" json:"created_at,omitempty"`
			Email       string                 `url:"email,omitempty" json:"email,omitempty"`
			FamilyName  string                 `url:"family_name,omitempty" json:"family_name,omitempty"`
			GivenName   string                 `url:"given_name,omitempty" json:"given_name,omitempty"`
			Id          string                 `url:"id,omitempty" json:"id,omitempty"`
			Language    string                 `url:"language,omitempty" json:"language,omitempty"`
			Metadata    map[string]interface{} `url:"metadata,omitempty" json:"metadata,omitempty"`
			PhoneNumber string                 `url:"phone_number,omitempty" json:"phone_number,omitempty"`
		} `url:"customer,omitempty" json:"customer,omitempty"`
		CustomerBankAccount struct {
			AccountHolderName   string `url:"account_holder_name,omitempty" json:"account_holder_name,omitempty"`
			AccountNumberEnding string `url:"account_number_ending,omitempty" json:"account_number_ending,omitempty"`
			AccountType         string `url:"account_type,omitempty" json:"account_type,omitempty"`
			BankName            string `url:"bank_name,omitempty" json:"bank_name,omitempty"`
			CountryCode         string `url:"country_code,omitempty" json:"country_code,omitempty"`
			CreatedAt           string `url:"created_at,omitempty" json:"created_at,omitempty"`
			Currency            string `url:"currency,omitempty" json:"currency,omitempty"`
			Enabled             bool   `url:"enabled,omitempty" json:"enabled,omitempty"`
			Id                  string `url:"id,omitempty" json:"id,omitempty"`
			Links               struct {
				Customer string `url:"customer,omitempty" json:"customer,omitempty"`
			} `url:"links,omitempty" json:"links,omitempty"`
			Metadata map[string]interface{} `url:"metadata,omitempty" json:"metadata,omitempty"`
		} `url:"customer_bank_account,omitempty" json:"customer_bank_account,omitempty"`
		CustomerBillingDetail struct {
			AddressLine1          string   `url:"address_line1,omitempty" json:"address_line1,omitempty"`
			AddressLine2          string   `url:"address_line2,omitempty" json:"address_line2,omitempty"`
			AddressLine3          string   `url:"address_line3,omitempty" json:"address_line3,omitempty"`
			City                  string   `url:"city,omitempty" json:"city,omitempty"`
			CountryCode           string   `url:"country_code,omitempty" json:"country_code,omitempty"`
			CreatedAt             string   `url:"created_at,omitempty" json:"created_at,omitempty"`
			DanishIdentityNumber  string   `url:"danish_identity_number,omitempty" json:"danish_identity_number,omitempty"`
			Id                    string   `url:"id,omitempty" json:"id,omitempty"`
			IpAddress             string   `url:"ip_address,omitempty" json:"ip_address,omitempty"`
			PostalCode            string   `url:"postal_code,omitempty" json:"postal_code,omitempty"`
			Region                string   `url:"region,omitempty" json:"region,omitempty"`
			Schemes               []string `url:"schemes,omitempty" json:"schemes,omitempty"`
			SwedishIdentityNumber string   `url:"swedish_identity_number,omitempty" json:"swedish_identity_number,omitempty"`
		} `url:"customer_billing_detail,omitempty" json:"customer_billing_detail,omitempty"`
	} `url:"resources,omitempty" json:"resources,omitempty"`
	Status string `url:"status,omitempty" json:"status,omitempty"`
}

// BillingRequestListParams parameters
type BillingRequestListParams struct {
	After     string `url:"after,omitempty" json:"after,omitempty"`
	Before    string `url:"before,omitempty" json:"before,omitempty"`
	CreatedAt string `url:"created_at,omitempty" json:"created_at,omitempty"`
	Customer  string `url:"customer,omitempty" json:"customer,omitempty"`
	Limit     int    `url:"limit,omitempty" json:"limit,omitempty"`
	Status    string `url:"status,omitempty" json:"status,omitempty"`
}

// BillingRequestListResult response including pagination metadata
type BillingRequestListResult struct {
	BillingRequests []BillingRequest `json:"billing_requests"`
	Meta            struct {
		Cursors struct {
			After  string `url:"after,omitempty" json:"after,omitempty"`
			Before string `url:"before,omitempty" json:"before,omitempty"`
		} `url:"cursors,omitempty" json:"cursors,omitempty"`
		Limit int `url:"limit,omitempty" json:"limit,omitempty"`
	} `json:"meta"`
}

// List
// Returns a [cursor-paginated](#api-usage-cursor-pagination) list of your
// billing_requests.
func (s *BillingRequestService) List(ctx context.Context, p BillingRequestListParams, opts ...RequestOption) (*BillingRequestListResult, error) {
	uri, err := url.Parse(fmt.Sprintf(s.endpoint + "/billing_requests"))
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

	v, err := query.Values(p)
	if err != nil {
		return nil, err
	}
	uri.RawQuery = v.Encode()

	req, err := http.NewRequest("GET", uri.String(), body)
	if err != nil {
		return nil, err
	}
	req.WithContext(ctx)
	req.Header.Set("Authorization", "Bearer "+s.token)

	req.Header.Set("GoCardless-Version", "2015-07-06")

	req.Header.Set("GoCardless-Client-Library", "<no value>")
	req.Header.Set("GoCardless-Client-Version", "1.0.0")
	req.Header.Set("User-Agent", userAgent)

	for key, value := range o.headers {
		req.Header.Set(key, value)
	}

	client := s.client
	if client == nil {
		client = http.DefaultClient
	}

	var result struct {
		Err *APIError `json:"error"`
		*BillingRequestListResult
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

	if result.BillingRequestListResult == nil {
		return nil, errors.New("missing result")
	}

	return result.BillingRequestListResult, nil
}

type BillingRequestListPagingIterator struct {
	cursor   string
	response *BillingRequestListResult
	params   BillingRequestListParams
	service  *BillingRequestService
}

func (c *BillingRequestListPagingIterator) Next() bool {
	if c.cursor == "" && c.response != nil {
		return false
	}

	return true
}

func (c *BillingRequestListPagingIterator) Value(ctx context.Context) (*BillingRequestListResult, error) {
	if !c.Next() {
		return c.response, nil
	}

	s := c.service
	p := c.params
	p.After = c.cursor

	uri, err := url.Parse(fmt.Sprintf(s.endpoint + "/billing_requests"))

	if err != nil {
		return nil, err
	}

	var body io.Reader

	v, err := query.Values(p)
	if err != nil {
		return nil, err
	}
	uri.RawQuery = v.Encode()

	req, err := http.NewRequest("GET", uri.String(), body)
	if err != nil {
		return nil, err
	}
	req.WithContext(ctx)
	req.Header.Set("Authorization", "Bearer "+s.token)
	req.Header.Set("GoCardless-Version", "2015-07-06")

	client := s.client
	if client == nil {
		client = http.DefaultClient
	}

	var result struct {
		Err *APIError `json:"error"`
		*BillingRequestListResult
	}

	err = try(3, func() error {
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

	if result.BillingRequestListResult == nil {
		return nil, errors.New("missing result")
	}

	c.response = result.BillingRequestListResult
	c.cursor = c.response.Meta.Cursors.After
	return c.response, nil
}

func (s *BillingRequestService) All(ctx context.Context, p BillingRequestListParams) *BillingRequestListPagingIterator {
	return &BillingRequestListPagingIterator{
		params:  p,
		service: s,
	}
}

// BillingRequestCreateParams parameters
type BillingRequestCreateParams struct {
	Links struct {
		Creditor            string `url:"creditor,omitempty" json:"creditor,omitempty"`
		Customer            string `url:"customer,omitempty" json:"customer,omitempty"`
		CustomerBankAccount string `url:"customer_bank_account,omitempty" json:"customer_bank_account,omitempty"`
	} `url:"links,omitempty" json:"links,omitempty"`
	MandateRequest struct {
		Currency string `url:"currency,omitempty" json:"currency,omitempty"`
		Scheme   string `url:"scheme,omitempty" json:"scheme,omitempty"`
	} `url:"mandate_request,omitempty" json:"mandate_request,omitempty"`
	Metadata       map[string]interface{} `url:"metadata,omitempty" json:"metadata,omitempty"`
	PaymentRequest struct {
		Amount      int    `url:"amount,omitempty" json:"amount,omitempty"`
		AppFee      int    `url:"app_fee,omitempty" json:"app_fee,omitempty"`
		Currency    string `url:"currency,omitempty" json:"currency,omitempty"`
		Description string `url:"description,omitempty" json:"description,omitempty"`
	} `url:"payment_request,omitempty" json:"payment_request,omitempty"`
}

// Create
//
func (s *BillingRequestService) Create(ctx context.Context, p BillingRequestCreateParams, opts ...RequestOption) (*BillingRequest, error) {
	uri, err := url.Parse(fmt.Sprintf(s.endpoint + "/billing_requests"))
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
		"billing_requests": p,
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
		Err            *APIError       `json:"error"`
		BillingRequest *BillingRequest `json:"billing_requests"`
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

	if result.BillingRequest == nil {
		return nil, errors.New("missing result")
	}

	return result.BillingRequest, nil
}

// Get
// Fetches a billing request
func (s *BillingRequestService) Get(ctx context.Context, identity string, opts ...RequestOption) (*BillingRequest, error) {
	uri, err := url.Parse(fmt.Sprintf(s.endpoint+"/billing_requests/%v",
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
	req.WithContext(ctx)
	req.Header.Set("Authorization", "Bearer "+s.token)

	req.Header.Set("GoCardless-Version", "2015-07-06")

	req.Header.Set("GoCardless-Client-Library", "<no value>")
	req.Header.Set("GoCardless-Client-Version", "1.0.0")
	req.Header.Set("User-Agent", userAgent)

	for key, value := range o.headers {
		req.Header.Set(key, value)
	}

	client := s.client
	if client == nil {
		client = http.DefaultClient
	}

	var result struct {
		Err            *APIError       `json:"error"`
		BillingRequest *BillingRequest `json:"billing_requests"`
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

	if result.BillingRequest == nil {
		return nil, errors.New("missing result")
	}

	return result.BillingRequest, nil
}

// BillingRequestCollectCustomerDetailsParams parameters
type BillingRequestCollectCustomerDetailsParams struct {
	Customer struct {
		CompanyName string                 `url:"company_name,omitempty" json:"company_name,omitempty"`
		Email       string                 `url:"email,omitempty" json:"email,omitempty"`
		FamilyName  string                 `url:"family_name,omitempty" json:"family_name,omitempty"`
		GivenName   string                 `url:"given_name,omitempty" json:"given_name,omitempty"`
		Language    string                 `url:"language,omitempty" json:"language,omitempty"`
		Metadata    map[string]interface{} `url:"metadata,omitempty" json:"metadata,omitempty"`
		PhoneNumber string                 `url:"phone_number,omitempty" json:"phone_number,omitempty"`
	} `url:"customer,omitempty" json:"customer,omitempty"`
	CustomerBillingDetail struct {
		AddressLine1          string `url:"address_line1,omitempty" json:"address_line1,omitempty"`
		AddressLine2          string `url:"address_line2,omitempty" json:"address_line2,omitempty"`
		AddressLine3          string `url:"address_line3,omitempty" json:"address_line3,omitempty"`
		City                  string `url:"city,omitempty" json:"city,omitempty"`
		CountryCode           string `url:"country_code,omitempty" json:"country_code,omitempty"`
		DanishIdentityNumber  string `url:"danish_identity_number,omitempty" json:"danish_identity_number,omitempty"`
		PostalCode            string `url:"postal_code,omitempty" json:"postal_code,omitempty"`
		Region                string `url:"region,omitempty" json:"region,omitempty"`
		SwedishIdentityNumber string `url:"swedish_identity_number,omitempty" json:"swedish_identity_number,omitempty"`
	} `url:"customer_billing_detail,omitempty" json:"customer_billing_detail,omitempty"`
}

// CollectCustomerDetails
// If the billing request has a pending <code>collect_customer_details</code>
// action, this endpoint can be used to collect the details in order to
// complete it.
//
// The endpoint takes the same payload as Customers, but checks that the
// customer fields are populated correctly for the billing request scheme.
//
// Whatever is provided to this endpoint is used to update the referenced
// customer, and will take effect immediately after the request is
// successful.
func (s *BillingRequestService) CollectCustomerDetails(ctx context.Context, identity string, p BillingRequestCollectCustomerDetailsParams, opts ...RequestOption) (*BillingRequest, error) {
	uri, err := url.Parse(fmt.Sprintf(s.endpoint+"/billing_requests/%v/actions/collect_customer_details",
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
		Err            *APIError       `json:"error"`
		BillingRequest *BillingRequest `json:"billing_requests"`
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

	if result.BillingRequest == nil {
		return nil, errors.New("missing result")
	}

	return result.BillingRequest, nil
}

// BillingRequestCollectBankAccountParams parameters
type BillingRequestCollectBankAccountParams struct {
	AccountHolderName   string                 `url:"account_holder_name,omitempty" json:"account_holder_name,omitempty"`
	AccountNumber       string                 `url:"account_number,omitempty" json:"account_number,omitempty"`
	AccountNumberSuffix string                 `url:"account_number_suffix,omitempty" json:"account_number_suffix,omitempty"`
	AccountType         string                 `url:"account_type,omitempty" json:"account_type,omitempty"`
	BankCode            string                 `url:"bank_code,omitempty" json:"bank_code,omitempty"`
	BranchCode          string                 `url:"branch_code,omitempty" json:"branch_code,omitempty"`
	CountryCode         string                 `url:"country_code,omitempty" json:"country_code,omitempty"`
	Currency            string                 `url:"currency,omitempty" json:"currency,omitempty"`
	Iban                string                 `url:"iban,omitempty" json:"iban,omitempty"`
	Metadata            map[string]interface{} `url:"metadata,omitempty" json:"metadata,omitempty"`
}

// CollectBankAccount
// If the billing request has a pending
// <code>collect_bank_account</code> action, this endpoint can be
// used to collect the details in order to complete it.
//
// The endpoint takes the same payload as Customer Bank Accounts, but check
// the bank account is valid for the billing request scheme before creating
// and attaching it.
func (s *BillingRequestService) CollectBankAccount(ctx context.Context, identity string, p BillingRequestCollectBankAccountParams, opts ...RequestOption) (*BillingRequest, error) {
	uri, err := url.Parse(fmt.Sprintf(s.endpoint+"/billing_requests/%v/actions/collect_bank_account",
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
		Err            *APIError       `json:"error"`
		BillingRequest *BillingRequest `json:"billing_requests"`
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

	if result.BillingRequest == nil {
		return nil, errors.New("missing result")
	}

	return result.BillingRequest, nil
}

// BillingRequestFulfilParams parameters
type BillingRequestFulfilParams struct {
	Metadata map[string]interface{} `url:"metadata,omitempty" json:"metadata,omitempty"`
}

// Fulfil
// If a billing request is ready to be fulfilled, call this endpoint to cause
// it to fulfil, executing the payment.
func (s *BillingRequestService) Fulfil(ctx context.Context, identity string, p BillingRequestFulfilParams, opts ...RequestOption) (*BillingRequest, error) {
	uri, err := url.Parse(fmt.Sprintf(s.endpoint+"/billing_requests/%v/actions/fulfil",
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
		Err            *APIError       `json:"error"`
		BillingRequest *BillingRequest `json:"billing_requests"`
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

	if result.BillingRequest == nil {
		return nil, errors.New("missing result")
	}

	return result.BillingRequest, nil
}

// BillingRequestConfirmPayerDetailsParams parameters
type BillingRequestConfirmPayerDetailsParams struct {
	Metadata map[string]interface{} `url:"metadata,omitempty" json:"metadata,omitempty"`
}

// ConfirmPayerDetails
// This is needed when you have mandate_request. As a scheme compliance rule we
// are required to
// allow the payer to crosscheck the details entered by them and confirm it.
func (s *BillingRequestService) ConfirmPayerDetails(ctx context.Context, identity string, p BillingRequestConfirmPayerDetailsParams, opts ...RequestOption) (*BillingRequest, error) {
	uri, err := url.Parse(fmt.Sprintf(s.endpoint+"/billing_requests/%v/actions/confirm_payer_details",
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
		Err            *APIError       `json:"error"`
		BillingRequest *BillingRequest `json:"billing_requests"`
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

	if result.BillingRequest == nil {
		return nil, errors.New("missing result")
	}

	return result.BillingRequest, nil
}

// BillingRequestCancelParams parameters
type BillingRequestCancelParams struct {
	Metadata map[string]interface{} `url:"metadata,omitempty" json:"metadata,omitempty"`
}

// Cancel
// Immediately cancels a billing request, causing all billing request flows
// to expire.
func (s *BillingRequestService) Cancel(ctx context.Context, identity string, p BillingRequestCancelParams, opts ...RequestOption) (*BillingRequest, error) {
	uri, err := url.Parse(fmt.Sprintf(s.endpoint+"/billing_requests/%v/actions/cancel",
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
		Err            *APIError       `json:"error"`
		BillingRequest *BillingRequest `json:"billing_requests"`
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

	if result.BillingRequest == nil {
		return nil, errors.New("missing result")
	}

	return result.BillingRequest, nil
}

// BillingRequestNotifyParams parameters
type BillingRequestNotifyParams struct {
	NotificationType string `url:"notification_type,omitempty" json:"notification_type,omitempty"`
	RedirectUri      string `url:"redirect_uri,omitempty" json:"redirect_uri,omitempty"`
}

// Notify
// Notifies the customer linked to the billing request, asking them to authorise
// it.
// Currently, the customer can only be notified by email.
func (s *BillingRequestService) Notify(ctx context.Context, identity string, p BillingRequestNotifyParams, opts ...RequestOption) (*BillingRequest, error) {
	uri, err := url.Parse(fmt.Sprintf(s.endpoint+"/billing_requests/%v/actions/notify",
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
		Err            *APIError       `json:"error"`
		BillingRequest *BillingRequest `json:"billing_requests"`
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

	if result.BillingRequest == nil {
		return nil, errors.New("missing result")
	}

	return result.BillingRequest, nil
}
