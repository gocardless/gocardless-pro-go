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

// BillingRequestTemplateService manages billing_request_templates
type BillingRequestTemplateService struct {
	endpoint string
	token    string
	client   *http.Client
}

// BillingRequestTemplate model
type BillingRequestTemplate struct {
	AuthorisationUrl          string                 `url:"authorisation_url,omitempty" json:"authorisation_url,omitempty"`
	CreatedAt                 string                 `url:"created_at,omitempty" json:"created_at,omitempty"`
	Id                        string                 `url:"id,omitempty" json:"id,omitempty"`
	MandateRequestCurrency    string                 `url:"mandate_request_currency,omitempty" json:"mandate_request_currency,omitempty"`
	MandateRequestMetadata    map[string]interface{} `url:"mandate_request_metadata,omitempty" json:"mandate_request_metadata,omitempty"`
	MandateRequestScheme      string                 `url:"mandate_request_scheme,omitempty" json:"mandate_request_scheme,omitempty"`
	MandateRequestVerify      string                 `url:"mandate_request_verify,omitempty" json:"mandate_request_verify,omitempty"`
	Metadata                  map[string]interface{} `url:"metadata,omitempty" json:"metadata,omitempty"`
	Name                      string                 `url:"name,omitempty" json:"name,omitempty"`
	PaymentRequestAmount      int                    `url:"payment_request_amount,omitempty" json:"payment_request_amount,omitempty"`
	PaymentRequestCurrency    string                 `url:"payment_request_currency,omitempty" json:"payment_request_currency,omitempty"`
	PaymentRequestDescription string                 `url:"payment_request_description,omitempty" json:"payment_request_description,omitempty"`
	PaymentRequestMetadata    map[string]interface{} `url:"payment_request_metadata,omitempty" json:"payment_request_metadata,omitempty"`
	PaymentRequestScheme      string                 `url:"payment_request_scheme,omitempty" json:"payment_request_scheme,omitempty"`
	RedirectUri               string                 `url:"redirect_uri,omitempty" json:"redirect_uri,omitempty"`
	UpdatedAt                 string                 `url:"updated_at,omitempty" json:"updated_at,omitempty"`
}

// BillingRequestTemplateListParams parameters
type BillingRequestTemplateListParams struct {
	After  string `url:"after,omitempty" json:"after,omitempty"`
	Before string `url:"before,omitempty" json:"before,omitempty"`
	Limit  int    `url:"limit,omitempty" json:"limit,omitempty"`
}

// BillingRequestTemplateListResult response including pagination metadata
type BillingRequestTemplateListResult struct {
	BillingRequestTemplates []BillingRequestTemplate `json:"billing_request_templates"`
	Meta                    struct {
		Cursors struct {
			After  string `url:"after,omitempty" json:"after,omitempty"`
			Before string `url:"before,omitempty" json:"before,omitempty"`
		} `url:"cursors,omitempty" json:"cursors,omitempty"`
		Limit int `url:"limit,omitempty" json:"limit,omitempty"`
	} `json:"meta"`
}

// List
// Returns a [cursor-paginated](#api-usage-cursor-pagination) list of your
// Billing Request Templates.
func (s *BillingRequestTemplateService) List(ctx context.Context, p BillingRequestTemplateListParams, opts ...RequestOption) (*BillingRequestTemplateListResult, error) {
	uri, err := url.Parse(fmt.Sprintf(s.endpoint + "/billing_request_templates"))
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
		*BillingRequestTemplateListResult
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

	if result.BillingRequestTemplateListResult == nil {
		return nil, errors.New("missing result")
	}

	return result.BillingRequestTemplateListResult, nil
}

type BillingRequestTemplateListPagingIterator struct {
	cursor         string
	response       *BillingRequestTemplateListResult
	params         BillingRequestTemplateListParams
	service        *BillingRequestTemplateService
	requestOptions []RequestOption
}

func (c *BillingRequestTemplateListPagingIterator) Next() bool {
	if c.cursor == "" && c.response != nil {
		return false
	}

	return true
}

func (c *BillingRequestTemplateListPagingIterator) Value(ctx context.Context) (*BillingRequestTemplateListResult, error) {
	if !c.Next() {
		return c.response, nil
	}

	s := c.service
	p := c.params
	p.After = c.cursor

	uri, err := url.Parse(fmt.Sprintf(s.endpoint + "/billing_request_templates"))

	if err != nil {
		return nil, err
	}

	o := &requestOptions{
		retries: 3,
	}
	for _, opt := range c.requestOptions {
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

	for key, value := range o.headers {
		req.Header.Set(key, value)
	}
	client := s.client
	if client == nil {
		client = http.DefaultClient
	}

	var result struct {
		Err *APIError `json:"error"`
		*BillingRequestTemplateListResult
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

	if result.BillingRequestTemplateListResult == nil {
		return nil, errors.New("missing result")
	}

	c.response = result.BillingRequestTemplateListResult
	c.cursor = c.response.Meta.Cursors.After
	return c.response, nil
}

func (s *BillingRequestTemplateService) All(ctx context.Context,
	p BillingRequestTemplateListParams,
	opts ...RequestOption) *BillingRequestTemplateListPagingIterator {
	return &BillingRequestTemplateListPagingIterator{
		params:         p,
		service:        s,
		requestOptions: opts,
	}
}

// Get
// Fetches a Billing Request Template
func (s *BillingRequestTemplateService) Get(ctx context.Context, identity string, opts ...RequestOption) (*BillingRequestTemplate, error) {
	uri, err := url.Parse(fmt.Sprintf(s.endpoint+"/billing_request_templates/%v",
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
		Err                    *APIError               `json:"error"`
		BillingRequestTemplate *BillingRequestTemplate `json:"billing_request_templates"`
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

	if result.BillingRequestTemplate == nil {
		return nil, errors.New("missing result")
	}

	return result.BillingRequestTemplate, nil
}

// BillingRequestTemplateCreateParams parameters
type BillingRequestTemplateCreateParams struct {
	Links struct {
		Creditor string `url:"creditor,omitempty" json:"creditor,omitempty"`
	} `url:"links,omitempty" json:"links,omitempty"`
	MandateRequestCurrency    string                 `url:"mandate_request_currency,omitempty" json:"mandate_request_currency,omitempty"`
	MandateRequestMetadata    map[string]interface{} `url:"mandate_request_metadata,omitempty" json:"mandate_request_metadata,omitempty"`
	MandateRequestScheme      string                 `url:"mandate_request_scheme,omitempty" json:"mandate_request_scheme,omitempty"`
	MandateRequestVerify      string                 `url:"mandate_request_verify,omitempty" json:"mandate_request_verify,omitempty"`
	Metadata                  map[string]interface{} `url:"metadata,omitempty" json:"metadata,omitempty"`
	Name                      string                 `url:"name,omitempty" json:"name,omitempty"`
	PaymentRequestAmount      int                    `url:"payment_request_amount,omitempty" json:"payment_request_amount,omitempty"`
	PaymentRequestCurrency    string                 `url:"payment_request_currency,omitempty" json:"payment_request_currency,omitempty"`
	PaymentRequestDescription string                 `url:"payment_request_description,omitempty" json:"payment_request_description,omitempty"`
	PaymentRequestMetadata    map[string]interface{} `url:"payment_request_metadata,omitempty" json:"payment_request_metadata,omitempty"`
	PaymentRequestScheme      string                 `url:"payment_request_scheme,omitempty" json:"payment_request_scheme,omitempty"`
	RedirectUri               string                 `url:"redirect_uri,omitempty" json:"redirect_uri,omitempty"`
}

// Create
//
func (s *BillingRequestTemplateService) Create(ctx context.Context, p BillingRequestTemplateCreateParams, opts ...RequestOption) (*BillingRequestTemplate, error) {
	uri, err := url.Parse(fmt.Sprintf(s.endpoint + "/billing_request_templates"))
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
		"billing_request_templates": p,
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
		Err                    *APIError               `json:"error"`
		BillingRequestTemplate *BillingRequestTemplate `json:"billing_request_templates"`
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

	if result.BillingRequestTemplate == nil {
		return nil, errors.New("missing result")
	}

	return result.BillingRequestTemplate, nil
}

// BillingRequestTemplateUpdateParams parameters
type BillingRequestTemplateUpdateParams struct {
	MandateRequestCurrency    string                 `url:"mandate_request_currency,omitempty" json:"mandate_request_currency,omitempty"`
	MandateRequestMetadata    map[string]interface{} `url:"mandate_request_metadata,omitempty" json:"mandate_request_metadata,omitempty"`
	MandateRequestScheme      string                 `url:"mandate_request_scheme,omitempty" json:"mandate_request_scheme,omitempty"`
	MandateRequestVerify      string                 `url:"mandate_request_verify,omitempty" json:"mandate_request_verify,omitempty"`
	Metadata                  map[string]interface{} `url:"metadata,omitempty" json:"metadata,omitempty"`
	Name                      string                 `url:"name,omitempty" json:"name,omitempty"`
	PaymentRequestAmount      int                    `url:"payment_request_amount,omitempty" json:"payment_request_amount,omitempty"`
	PaymentRequestCurrency    string                 `url:"payment_request_currency,omitempty" json:"payment_request_currency,omitempty"`
	PaymentRequestDescription string                 `url:"payment_request_description,omitempty" json:"payment_request_description,omitempty"`
	PaymentRequestMetadata    map[string]interface{} `url:"payment_request_metadata,omitempty" json:"payment_request_metadata,omitempty"`
	PaymentRequestScheme      string                 `url:"payment_request_scheme,omitempty" json:"payment_request_scheme,omitempty"`
	RedirectUri               string                 `url:"redirect_uri,omitempty" json:"redirect_uri,omitempty"`
}

// Update
// Updates a Billing Request Template, which will affect all future Billing
// Requests created by this template.
func (s *BillingRequestTemplateService) Update(ctx context.Context, identity string, p BillingRequestTemplateUpdateParams, opts ...RequestOption) (*BillingRequestTemplate, error) {
	uri, err := url.Parse(fmt.Sprintf(s.endpoint+"/billing_request_templates/%v",
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
		"billing_request_templates": p,
	})
	if err != nil {
		return nil, err
	}
	body = &buf

	req, err := http.NewRequest("PUT", uri.String(), body)
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
		Err                    *APIError               `json:"error"`
		BillingRequestTemplate *BillingRequestTemplate `json:"billing_request_templates"`
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

	if result.BillingRequestTemplate == nil {
		return nil, errors.New("missing result")
	}

	return result.BillingRequestTemplate, nil
}
