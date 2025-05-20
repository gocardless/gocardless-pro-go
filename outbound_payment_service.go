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

// OutboundPaymentService manages outbound_payments
type OutboundPaymentServiceImpl struct {
	config Config
}

type OutboundPaymentLinks struct {
	Creditor             string `url:"creditor,omitempty" json:"creditor,omitempty"`
	Customer             string `url:"customer,omitempty" json:"customer,omitempty"`
	RecipientBankAccount string `url:"recipient_bank_account,omitempty" json:"recipient_bank_account,omitempty"`
}

type OutboundPaymentVerificationsRecipientBankAccountHolderVerification struct {
	ActualAccountName string `url:"actual_account_name,omitempty" json:"actual_account_name,omitempty"`
	Result            string `url:"result,omitempty" json:"result,omitempty"`
	Type              string `url:"type,omitempty" json:"type,omitempty"`
}

type OutboundPaymentVerifications struct {
	RecipientBankAccountHolderVerification *OutboundPaymentVerificationsRecipientBankAccountHolderVerification `url:"recipient_bank_account_holder_verification,omitempty" json:"recipient_bank_account_holder_verification,omitempty"`
}

// OutboundPayment model
type OutboundPayment struct {
	Amount        int                           `url:"amount,omitempty" json:"amount,omitempty"`
	CreatedAt     string                        `url:"created_at,omitempty" json:"created_at,omitempty"`
	Currency      string                        `url:"currency,omitempty" json:"currency,omitempty"`
	Description   string                        `url:"description,omitempty" json:"description,omitempty"`
	ExecutionDate string                        `url:"execution_date,omitempty" json:"execution_date,omitempty"`
	Id            string                        `url:"id,omitempty" json:"id,omitempty"`
	IsWithdrawal  bool                          `url:"is_withdrawal,omitempty" json:"is_withdrawal,omitempty"`
	Links         *OutboundPaymentLinks         `url:"links,omitempty" json:"links,omitempty"`
	Metadata      map[string]interface{}        `url:"metadata,omitempty" json:"metadata,omitempty"`
	Reference     string                        `url:"reference,omitempty" json:"reference,omitempty"`
	Scheme        string                        `url:"scheme,omitempty" json:"scheme,omitempty"`
	Status        string                        `url:"status,omitempty" json:"status,omitempty"`
	Verifications *OutboundPaymentVerifications `url:"verifications,omitempty" json:"verifications,omitempty"`
}

type OutboundPaymentService interface {
	Create(ctx context.Context, p OutboundPaymentCreateParams, opts ...RequestOption) (*OutboundPayment, error)
	Withdraw(ctx context.Context, p OutboundPaymentWithdrawParams, opts ...RequestOption) (*OutboundPayment, error)
	Cancel(ctx context.Context, identity string, p OutboundPaymentCancelParams, opts ...RequestOption) (*OutboundPayment, error)
	Approve(ctx context.Context, identity string, p OutboundPaymentApproveParams, opts ...RequestOption) (*OutboundPayment, error)
	Get(ctx context.Context, identity string, opts ...RequestOption) (*OutboundPayment, error)
	List(ctx context.Context, p OutboundPaymentListParams, opts ...RequestOption) (*OutboundPaymentListResult, error)
	All(ctx context.Context, p OutboundPaymentListParams, opts ...RequestOption) *OutboundPaymentListPagingIterator
	Update(ctx context.Context, identity string, p OutboundPaymentUpdateParams, opts ...RequestOption) (*OutboundPayment, error)
}

type OutboundPaymentCreateParamsLinks struct {
	Creditor             string `url:"creditor,omitempty" json:"creditor,omitempty"`
	RecipientBankAccount string `url:"recipient_bank_account,omitempty" json:"recipient_bank_account,omitempty"`
}

// OutboundPaymentCreateParams parameters
type OutboundPaymentCreateParams struct {
	Amount        int                              `url:"amount,omitempty" json:"amount,omitempty"`
	Description   string                           `url:"description,omitempty" json:"description,omitempty"`
	ExecutionDate string                           `url:"execution_date,omitempty" json:"execution_date,omitempty"`
	Links         OutboundPaymentCreateParamsLinks `url:"links,omitempty" json:"links,omitempty"`
	Metadata      map[string]interface{}           `url:"metadata,omitempty" json:"metadata,omitempty"`
	Scheme        string                           `url:"scheme,omitempty" json:"scheme,omitempty"`
}

// Create
func (s *OutboundPaymentServiceImpl) Create(ctx context.Context, p OutboundPaymentCreateParams, opts ...RequestOption) (*OutboundPayment, error) {
	uri, err := url.Parse(fmt.Sprintf(s.config.Endpoint() + "/outbound_payments"))
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
		"outbound_payments": p,
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
	req.Header.Set("GoCardless-Client-Version", "4.8.0")
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
		Err             *APIError        `json:"error"`
		OutboundPayment *OutboundPayment `json:"outbound_payments"`
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

	if result.OutboundPayment == nil {
		return nil, errors.New("missing result")
	}

	return result.OutboundPayment, nil
}

type OutboundPaymentWithdrawParamsLinks struct {
	Creditor string `url:"creditor,omitempty" json:"creditor,omitempty"`
}

// OutboundPaymentWithdrawParams parameters
type OutboundPaymentWithdrawParams struct {
	Amount        int                                 `url:"amount,omitempty" json:"amount,omitempty"`
	Description   string                              `url:"description,omitempty" json:"description,omitempty"`
	ExecutionDate string                              `url:"execution_date,omitempty" json:"execution_date,omitempty"`
	Links         *OutboundPaymentWithdrawParamsLinks `url:"links,omitempty" json:"links,omitempty"`
	Metadata      map[string]interface{}              `url:"metadata,omitempty" json:"metadata,omitempty"`
	Scheme        string                              `url:"scheme,omitempty" json:"scheme,omitempty"`
}

// Withdraw
// Creates an outbound payment to your verified business bank account as the
// recipient.
func (s *OutboundPaymentServiceImpl) Withdraw(ctx context.Context, p OutboundPaymentWithdrawParams, opts ...RequestOption) (*OutboundPayment, error) {
	uri, err := url.Parse(fmt.Sprintf(s.config.Endpoint() + "/outbound_payments/withdrawal"))
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
	req.Header.Set("GoCardless-Client-Version", "4.8.0")
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
		Err             *APIError        `json:"error"`
		OutboundPayment *OutboundPayment `json:"outbound_payments"`
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

	if result.OutboundPayment == nil {
		return nil, errors.New("missing result")
	}

	return result.OutboundPayment, nil
}

// OutboundPaymentCancelParams parameters
type OutboundPaymentCancelParams struct {
	Metadata map[string]interface{} `url:"metadata,omitempty" json:"metadata,omitempty"`
}

// Cancel
// Cancels an outbound payment. Only outbound payments with either `verifying`,
// `pending_approval`, or `scheduled` status can be cancelled.
// Once an outbound payment is `executing`, the money moving process has begun
// and cannot be reversed.
func (s *OutboundPaymentServiceImpl) Cancel(ctx context.Context, identity string, p OutboundPaymentCancelParams, opts ...RequestOption) (*OutboundPayment, error) {
	uri, err := url.Parse(fmt.Sprintf(s.config.Endpoint()+"/outbound_payments/%v/actions/cancel",
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
	req.Header.Set("GoCardless-Client-Version", "4.8.0")
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
		Err             *APIError        `json:"error"`
		OutboundPayment *OutboundPayment `json:"outbound_payments"`
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

	if result.OutboundPayment == nil {
		return nil, errors.New("missing result")
	}

	return result.OutboundPayment, nil
}

// OutboundPaymentApproveParams parameters
type OutboundPaymentApproveParams struct {
}

// Approve
// Approves an outbound payment. Only outbound payments with the
// “pending_approval” status can be approved.
func (s *OutboundPaymentServiceImpl) Approve(ctx context.Context, identity string, p OutboundPaymentApproveParams, opts ...RequestOption) (*OutboundPayment, error) {
	uri, err := url.Parse(fmt.Sprintf(s.config.Endpoint()+"/outbound_payments/%v/actions/approve",
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
	req.Header.Set("GoCardless-Client-Version", "4.8.0")
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
		Err             *APIError        `json:"error"`
		OutboundPayment *OutboundPayment `json:"outbound_payments"`
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

	if result.OutboundPayment == nil {
		return nil, errors.New("missing result")
	}

	return result.OutboundPayment, nil
}

// Get
// Fetches an outbound_payment by ID
func (s *OutboundPaymentServiceImpl) Get(ctx context.Context, identity string, opts ...RequestOption) (*OutboundPayment, error) {
	uri, err := url.Parse(fmt.Sprintf(s.config.Endpoint()+"/outbound_payments/%v",
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
	req.Header.Set("GoCardless-Client-Version", "4.8.0")
	req.Header.Set("User-Agent", userAgent)

	for key, value := range o.headers {
		req.Header.Set(key, value)
	}

	client := s.config.Client()
	if client == nil {
		client = http.DefaultClient
	}

	var result struct {
		Err             *APIError        `json:"error"`
		OutboundPayment *OutboundPayment `json:"outbound_payments"`
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

	if result.OutboundPayment == nil {
		return nil, errors.New("missing result")
	}

	return result.OutboundPayment, nil
}

// OutboundPaymentListParams parameters
type OutboundPaymentListParams struct {
	After       string `url:"after,omitempty" json:"after,omitempty"`
	Before      string `url:"before,omitempty" json:"before,omitempty"`
	CreatedFrom string `url:"created_from,omitempty" json:"created_from,omitempty"`
	CreatedTo   string `url:"created_to,omitempty" json:"created_to,omitempty"`
	Limit       int    `url:"limit,omitempty" json:"limit,omitempty"`
	Status      string `url:"status,omitempty" json:"status,omitempty"`
}

type OutboundPaymentListResultMetaCursors struct {
	After  string `url:"after,omitempty" json:"after,omitempty"`
	Before string `url:"before,omitempty" json:"before,omitempty"`
}

type OutboundPaymentListResultMeta struct {
	Cursors *OutboundPaymentListResultMetaCursors `url:"cursors,omitempty" json:"cursors,omitempty"`
	Limit   int                                   `url:"limit,omitempty" json:"limit,omitempty"`
}

type OutboundPaymentListResult struct {
	OutboundPayments []OutboundPayment             `json:"outbound_payments"`
	Meta             OutboundPaymentListResultMeta `url:"meta,omitempty" json:"meta,omitempty"`
}

// List
// Returns a [cursor-paginated](#api-usage-cursor-pagination) list of outbound
// payments.
func (s *OutboundPaymentServiceImpl) List(ctx context.Context, p OutboundPaymentListParams, opts ...RequestOption) (*OutboundPaymentListResult, error) {
	uri, err := url.Parse(fmt.Sprintf(s.config.Endpoint() + "/outbound_payments"))
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
	req = req.WithContext(ctx)
	req.Header.Set("Authorization", "Bearer "+s.config.Token())
	req.Header.Set("GoCardless-Version", "2015-07-06")
	req.Header.Set("GoCardless-Client-Library", "gocardless-pro-go")
	req.Header.Set("GoCardless-Client-Version", "4.8.0")
	req.Header.Set("User-Agent", userAgent)

	for key, value := range o.headers {
		req.Header.Set(key, value)
	}

	client := s.config.Client()
	if client == nil {
		client = http.DefaultClient
	}

	var result struct {
		Err *APIError `json:"error"`
		*OutboundPaymentListResult
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

	if result.OutboundPaymentListResult == nil {
		return nil, errors.New("missing result")
	}

	return result.OutboundPaymentListResult, nil
}

type OutboundPaymentListPagingIterator struct {
	cursor         string
	response       *OutboundPaymentListResult
	params         OutboundPaymentListParams
	service        *OutboundPaymentServiceImpl
	requestOptions []RequestOption
}

func (c *OutboundPaymentListPagingIterator) Next() bool {
	if c.cursor == "" && c.response != nil {
		return false
	}

	return true
}

func (c *OutboundPaymentListPagingIterator) Value(ctx context.Context) (*OutboundPaymentListResult, error) {
	if !c.Next() {
		return c.response, nil
	}

	s := c.service
	p := c.params
	p.After = c.cursor

	uri, err := url.Parse(fmt.Sprintf(s.config.Endpoint() + "/outbound_payments"))

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

	req = req.WithContext(ctx)
	req.Header.Set("Authorization", "Bearer "+s.config.Token())
	req.Header.Set("GoCardless-Version", "2015-07-06")
	req.Header.Set("GoCardless-Client-Library", "gocardless-pro-go")
	req.Header.Set("GoCardless-Client-Version", "4.8.0")
	req.Header.Set("User-Agent", userAgent)

	for key, value := range o.headers {
		req.Header.Set(key, value)
	}
	client := s.config.Client()
	if client == nil {
		client = http.DefaultClient
	}

	var result struct {
		Err *APIError `json:"error"`
		*OutboundPaymentListResult
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

	if result.OutboundPaymentListResult == nil {
		return nil, errors.New("missing result")
	}

	c.response = result.OutboundPaymentListResult
	c.cursor = c.response.Meta.Cursors.After
	return c.response, nil
}

func (s *OutboundPaymentServiceImpl) All(ctx context.Context,
	p OutboundPaymentListParams,
	opts ...RequestOption) *OutboundPaymentListPagingIterator {
	return &OutboundPaymentListPagingIterator{
		params:         p,
		service:        s,
		requestOptions: opts,
	}
}

// OutboundPaymentUpdateParams parameters
type OutboundPaymentUpdateParams struct {
	Metadata map[string]interface{} `url:"metadata,omitempty" json:"metadata,omitempty"`
}

// Update
// Updates an outbound payment object. This accepts only the metadata parameter.
func (s *OutboundPaymentServiceImpl) Update(ctx context.Context, identity string, p OutboundPaymentUpdateParams, opts ...RequestOption) (*OutboundPayment, error) {
	uri, err := url.Parse(fmt.Sprintf(s.config.Endpoint()+"/outbound_payments/%v",
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
		"outbound_payments": p,
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
	req.Header.Set("GoCardless-Client-Version", "4.8.0")
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
		Err             *APIError        `json:"error"`
		OutboundPayment *OutboundPayment `json:"outbound_payments"`
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

	if result.OutboundPayment == nil {
		return nil, errors.New("missing result")
	}

	return result.OutboundPayment, nil
}
