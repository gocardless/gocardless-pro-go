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

// MandateService manages mandates
type MandateService struct {
	endpoint string
	token    string
	client   *http.Client
}

type MandateLinks struct {
	Creditor            string `url:"creditor,omitempty" json:"creditor,omitempty"`
	Customer            string `url:"customer,omitempty" json:"customer,omitempty"`
	CustomerBankAccount string `url:"customer_bank_account,omitempty" json:"customer_bank_account,omitempty"`
	NewMandate          string `url:"new_mandate,omitempty" json:"new_mandate,omitempty"`
}

// Mandate model
type Mandate struct {
	CreatedAt               string                 `url:"created_at,omitempty" json:"created_at,omitempty"`
	Id                      string                 `url:"id,omitempty" json:"id,omitempty"`
	Links                   *MandateLinks          `url:"links,omitempty" json:"links,omitempty"`
	Metadata                map[string]interface{} `url:"metadata,omitempty" json:"metadata,omitempty"`
	NextPossibleChargeDate  string                 `url:"next_possible_charge_date,omitempty" json:"next_possible_charge_date,omitempty"`
	PaymentsRequireApproval bool                   `url:"payments_require_approval,omitempty" json:"payments_require_approval,omitempty"`
	Reference               string                 `url:"reference,omitempty" json:"reference,omitempty"`
	Scheme                  string                 `url:"scheme,omitempty" json:"scheme,omitempty"`
	Status                  string                 `url:"status,omitempty" json:"status,omitempty"`
}

type MandateCreateParamsLinks struct {
	Creditor            string `url:"creditor,omitempty" json:"creditor,omitempty"`
	CustomerBankAccount string `url:"customer_bank_account,omitempty" json:"customer_bank_account,omitempty"`
}

// MandateCreateParams parameters
type MandateCreateParams struct {
	Links          MandateCreateParamsLinks `url:"links,omitempty" json:"links,omitempty"`
	Metadata       map[string]interface{}   `url:"metadata,omitempty" json:"metadata,omitempty"`
	PayerIpAddress string                   `url:"payer_ip_address,omitempty" json:"payer_ip_address,omitempty"`
	Reference      string                   `url:"reference,omitempty" json:"reference,omitempty"`
	Scheme         string                   `url:"scheme,omitempty" json:"scheme,omitempty"`
}

// Create
// Creates a new mandate object.
func (s *MandateService) Create(ctx context.Context, p MandateCreateParams, opts ...RequestOption) (*Mandate, error) {
	uri, err := url.Parse(fmt.Sprintf(s.endpoint + "/mandates"))
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
		"mandates": p,
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
	req.Header.Set("GoCardless-Client-Library", "gocardless-pro-go")
	req.Header.Set("GoCardless-Client-Version", "2.0.0")
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
		Err     *APIError `json:"error"`
		Mandate *Mandate  `json:"mandates"`
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

	if result.Mandate == nil {
		return nil, errors.New("missing result")
	}

	return result.Mandate, nil
}

type MandateListParamsCreatedAt struct {
	Gt  string `url:"gt,omitempty" json:"gt,omitempty"`
	Gte string `url:"gte,omitempty" json:"gte,omitempty"`
	Lt  string `url:"lt,omitempty" json:"lt,omitempty"`
	Lte string `url:"lte,omitempty" json:"lte,omitempty"`
}

// MandateListParams parameters
type MandateListParams struct {
	After               string                      `url:"after,omitempty" json:"after,omitempty"`
	Before              string                      `url:"before,omitempty" json:"before,omitempty"`
	CreatedAt           *MandateListParamsCreatedAt `url:"created_at,omitempty" json:"created_at,omitempty"`
	Creditor            string                      `url:"creditor,omitempty" json:"creditor,omitempty"`
	Customer            string                      `url:"customer,omitempty" json:"customer,omitempty"`
	CustomerBankAccount string                      `url:"customer_bank_account,omitempty" json:"customer_bank_account,omitempty"`
	Limit               int                         `url:"limit,omitempty" json:"limit,omitempty"`
	Reference           string                      `url:"reference,omitempty" json:"reference,omitempty"`
	Scheme              []string                    `url:"scheme,omitempty" json:"scheme,omitempty"`
	Status              []string                    `url:"status,omitempty" json:"status,omitempty"`
}

type MandateListResultMetaCursors struct {
	After  string `url:"after,omitempty" json:"after,omitempty"`
	Before string `url:"before,omitempty" json:"before,omitempty"`
}

type MandateListResultMeta struct {
	Cursors *MandateListResultMetaCursors `url:"cursors,omitempty" json:"cursors,omitempty"`
	Limit   int                           `url:"limit,omitempty" json:"limit,omitempty"`
}

type MandateListResult struct {
	Mandates []Mandate             `json:"mandates"`
	Meta     MandateListResultMeta `url:"meta,omitempty" json:"meta,omitempty"`
}

// List
// Returns a [cursor-paginated](#api-usage-cursor-pagination) list of your
// mandates.
func (s *MandateService) List(ctx context.Context, p MandateListParams, opts ...RequestOption) (*MandateListResult, error) {
	uri, err := url.Parse(fmt.Sprintf(s.endpoint + "/mandates"))
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
	req.Header.Set("GoCardless-Client-Library", "gocardless-pro-go")
	req.Header.Set("GoCardless-Client-Version", "2.0.0")
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
		*MandateListResult
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

	if result.MandateListResult == nil {
		return nil, errors.New("missing result")
	}

	return result.MandateListResult, nil
}

type MandateListPagingIterator struct {
	cursor         string
	response       *MandateListResult
	params         MandateListParams
	service        *MandateService
	requestOptions []RequestOption
}

func (c *MandateListPagingIterator) Next() bool {
	if c.cursor == "" && c.response != nil {
		return false
	}

	return true
}

func (c *MandateListPagingIterator) Value(ctx context.Context) (*MandateListResult, error) {
	if !c.Next() {
		return c.response, nil
	}

	s := c.service
	p := c.params
	p.After = c.cursor

	uri, err := url.Parse(fmt.Sprintf(s.endpoint + "/mandates"))

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
	req.Header.Set("GoCardless-Client-Library", "gocardless-pro-go")
	req.Header.Set("GoCardless-Client-Version", "2.0.0")
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
		*MandateListResult
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

	if result.MandateListResult == nil {
		return nil, errors.New("missing result")
	}

	c.response = result.MandateListResult
	c.cursor = c.response.Meta.Cursors.After
	return c.response, nil
}

func (s *MandateService) All(ctx context.Context,
	p MandateListParams,
	opts ...RequestOption) *MandateListPagingIterator {
	return &MandateListPagingIterator{
		params:         p,
		service:        s,
		requestOptions: opts,
	}
}

// Get
// Retrieves the details of an existing mandate.
func (s *MandateService) Get(ctx context.Context, identity string, opts ...RequestOption) (*Mandate, error) {
	uri, err := url.Parse(fmt.Sprintf(s.endpoint+"/mandates/%v",
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
	req.Header.Set("GoCardless-Client-Library", "gocardless-pro-go")
	req.Header.Set("GoCardless-Client-Version", "2.0.0")
	req.Header.Set("User-Agent", userAgent)

	for key, value := range o.headers {
		req.Header.Set(key, value)
	}

	client := s.client
	if client == nil {
		client = http.DefaultClient
	}

	var result struct {
		Err     *APIError `json:"error"`
		Mandate *Mandate  `json:"mandates"`
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

	if result.Mandate == nil {
		return nil, errors.New("missing result")
	}

	return result.Mandate, nil
}

// MandateUpdateParams parameters
type MandateUpdateParams struct {
	Metadata map[string]interface{} `url:"metadata,omitempty" json:"metadata,omitempty"`
}

// Update
// Updates a mandate object. This accepts only the metadata parameter.
func (s *MandateService) Update(ctx context.Context, identity string, p MandateUpdateParams, opts ...RequestOption) (*Mandate, error) {
	uri, err := url.Parse(fmt.Sprintf(s.endpoint+"/mandates/%v",
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
		"mandates": p,
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
	req.Header.Set("GoCardless-Client-Library", "gocardless-pro-go")
	req.Header.Set("GoCardless-Client-Version", "2.0.0")
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
		Err     *APIError `json:"error"`
		Mandate *Mandate  `json:"mandates"`
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

	if result.Mandate == nil {
		return nil, errors.New("missing result")
	}

	return result.Mandate, nil
}

// MandateCancelParams parameters
type MandateCancelParams struct {
	Metadata map[string]interface{} `url:"metadata,omitempty" json:"metadata,omitempty"`
}

// Cancel
// Immediately cancels a mandate and all associated cancellable payments. Any
// metadata supplied to this endpoint will be stored on the mandate cancellation
// event it causes.
//
// This will fail with a `cancellation_failed` error if the mandate is already
// cancelled.
func (s *MandateService) Cancel(ctx context.Context, identity string, p MandateCancelParams, opts ...RequestOption) (*Mandate, error) {
	uri, err := url.Parse(fmt.Sprintf(s.endpoint+"/mandates/%v/actions/cancel",
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
	req.Header.Set("GoCardless-Client-Library", "gocardless-pro-go")
	req.Header.Set("GoCardless-Client-Version", "2.0.0")
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
		Err     *APIError `json:"error"`
		Mandate *Mandate  `json:"mandates"`
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

	if result.Mandate == nil {
		return nil, errors.New("missing result")
	}

	return result.Mandate, nil
}

// MandateReinstateParams parameters
type MandateReinstateParams struct {
	Metadata map[string]interface{} `url:"metadata,omitempty" json:"metadata,omitempty"`
}

// Reinstate
// <a name="mandate_not_inactive"></a>Reinstates a cancelled or expired mandate
// to the banks. You will receive a `resubmission_requested` webhook, but after
// that reinstating the mandate follows the same process as its initial
// creation, so you will receive a `submitted` webhook, followed by a
// `reinstated` or `failed` webhook up to two working days later. Any metadata
// supplied to this endpoint will be stored on the `resubmission_requested`
// event it causes.
//
// This will fail with a `mandate_not_inactive` error if the mandate is already
// being submitted, or is active.
//
// Mandates can be resubmitted up to 10 times.
func (s *MandateService) Reinstate(ctx context.Context, identity string, p MandateReinstateParams, opts ...RequestOption) (*Mandate, error) {
	uri, err := url.Parse(fmt.Sprintf(s.endpoint+"/mandates/%v/actions/reinstate",
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
	req.Header.Set("GoCardless-Client-Library", "gocardless-pro-go")
	req.Header.Set("GoCardless-Client-Version", "2.0.0")
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
		Err     *APIError `json:"error"`
		Mandate *Mandate  `json:"mandates"`
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

	if result.Mandate == nil {
		return nil, errors.New("missing result")
	}

	return result.Mandate, nil
}
