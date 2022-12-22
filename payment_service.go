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

// PaymentService manages payments
type PaymentServiceImpl struct {
	config Config
}

type PaymentFx struct {
	EstimatedExchangeRate string `url:"estimated_exchange_rate,omitempty" json:"estimated_exchange_rate,omitempty"`
	ExchangeRate          string `url:"exchange_rate,omitempty" json:"exchange_rate,omitempty"`
	FxAmount              int    `url:"fx_amount,omitempty" json:"fx_amount,omitempty"`
	FxCurrency            string `url:"fx_currency,omitempty" json:"fx_currency,omitempty"`
}

type PaymentLinks struct {
	Creditor           string `url:"creditor,omitempty" json:"creditor,omitempty"`
	InstalmentSchedule string `url:"instalment_schedule,omitempty" json:"instalment_schedule,omitempty"`
	Mandate            string `url:"mandate,omitempty" json:"mandate,omitempty"`
	Payout             string `url:"payout,omitempty" json:"payout,omitempty"`
	Subscription       string `url:"subscription,omitempty" json:"subscription,omitempty"`
}

// Payment model
type Payment struct {
	Amount          int                    `url:"amount,omitempty" json:"amount,omitempty"`
	AmountRefunded  int                    `url:"amount_refunded,omitempty" json:"amount_refunded,omitempty"`
	ChargeDate      string                 `url:"charge_date,omitempty" json:"charge_date,omitempty"`
	CreatedAt       string                 `url:"created_at,omitempty" json:"created_at,omitempty"`
	Currency        string                 `url:"currency,omitempty" json:"currency,omitempty"`
	Description     string                 `url:"description,omitempty" json:"description,omitempty"`
	Fx              *PaymentFx             `url:"fx,omitempty" json:"fx,omitempty"`
	Id              string                 `url:"id,omitempty" json:"id,omitempty"`
	Links           *PaymentLinks          `url:"links,omitempty" json:"links,omitempty"`
	Metadata        map[string]interface{} `url:"metadata,omitempty" json:"metadata,omitempty"`
	Reference       string                 `url:"reference,omitempty" json:"reference,omitempty"`
	RetryIfPossible bool                   `url:"retry_if_possible,omitempty" json:"retry_if_possible,omitempty"`
	Status          string                 `url:"status,omitempty" json:"status,omitempty"`
}

type PaymentService interface {
	Create(ctx context.Context, p PaymentCreateParams, opts ...RequestOption) (*Payment, error)
	List(ctx context.Context, p PaymentListParams, opts ...RequestOption) (*PaymentListResult, error)
	All(ctx context.Context, p PaymentListParams, opts ...RequestOption) *PaymentListPagingIterator
	Get(ctx context.Context, identity string, opts ...RequestOption) (*Payment, error)
	Update(ctx context.Context, identity string, p PaymentUpdateParams, opts ...RequestOption) (*Payment, error)
	Cancel(ctx context.Context, identity string, p PaymentCancelParams, opts ...RequestOption) (*Payment, error)
	Retry(ctx context.Context, identity string, p PaymentRetryParams, opts ...RequestOption) (*Payment, error)
}

type PaymentCreateParamsLinks struct {
	Mandate string `url:"mandate,omitempty" json:"mandate,omitempty"`
}

// PaymentCreateParams parameters
type PaymentCreateParams struct {
	Amount          int                      `url:"amount,omitempty" json:"amount,omitempty"`
	AppFee          int                      `url:"app_fee,omitempty" json:"app_fee,omitempty"`
	ChargeDate      string                   `url:"charge_date,omitempty" json:"charge_date,omitempty"`
	Currency        string                   `url:"currency,omitempty" json:"currency,omitempty"`
	Description     string                   `url:"description,omitempty" json:"description,omitempty"`
	Links           PaymentCreateParamsLinks `url:"links,omitempty" json:"links,omitempty"`
	Metadata        map[string]interface{}   `url:"metadata,omitempty" json:"metadata,omitempty"`
	Reference       string                   `url:"reference,omitempty" json:"reference,omitempty"`
	RetryIfPossible bool                     `url:"retry_if_possible,omitempty" json:"retry_if_possible,omitempty"`
}

// Create
// <a name="mandate_is_inactive"></a>Creates a new payment object.
//
// This fails with a `mandate_is_inactive` error if the linked
// [mandate](#core-endpoints-mandates) is cancelled or has failed. Payments can
// be created against mandates with status of: `pending_customer_approval`,
// `pending_submission`, `submitted`, and `active`.
func (s *PaymentServiceImpl) Create(ctx context.Context, p PaymentCreateParams, opts ...RequestOption) (*Payment, error) {
	uri, err := url.Parse(fmt.Sprintf(s.config.Endpoint() + "/payments"))
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
		"payments": p,
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
	req.Header.Set("GoCardless-Client-Version", "2.9.0")
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
		Err     *APIError `json:"error"`
		Payment *Payment  `json:"payments"`
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

	if result.Payment == nil {
		return nil, errors.New("missing result")
	}

	return result.Payment, nil
}

type PaymentListParamsChargeDate struct {
	Gt  string `url:"gt,omitempty" json:"gt,omitempty"`
	Gte string `url:"gte,omitempty" json:"gte,omitempty"`
	Lt  string `url:"lt,omitempty" json:"lt,omitempty"`
	Lte string `url:"lte,omitempty" json:"lte,omitempty"`
}

type PaymentListParamsCreatedAt struct {
	Gt  string `url:"gt,omitempty" json:"gt,omitempty"`
	Gte string `url:"gte,omitempty" json:"gte,omitempty"`
	Lt  string `url:"lt,omitempty" json:"lt,omitempty"`
	Lte string `url:"lte,omitempty" json:"lte,omitempty"`
}

// PaymentListParams parameters
type PaymentListParams struct {
	After         string                       `url:"after,omitempty" json:"after,omitempty"`
	Before        string                       `url:"before,omitempty" json:"before,omitempty"`
	ChargeDate    *PaymentListParamsChargeDate `url:"charge_date,omitempty" json:"charge_date,omitempty"`
	CreatedAt     *PaymentListParamsCreatedAt  `url:"created_at,omitempty" json:"created_at,omitempty"`
	Creditor      string                       `url:"creditor,omitempty" json:"creditor,omitempty"`
	Currency      string                       `url:"currency,omitempty" json:"currency,omitempty"`
	Customer      string                       `url:"customer,omitempty" json:"customer,omitempty"`
	Limit         int                          `url:"limit,omitempty" json:"limit,omitempty"`
	Mandate       string                       `url:"mandate,omitempty" json:"mandate,omitempty"`
	SortDirection string                       `url:"sort_direction,omitempty" json:"sort_direction,omitempty"`
	SortField     string                       `url:"sort_field,omitempty" json:"sort_field,omitempty"`
	Status        string                       `url:"status,omitempty" json:"status,omitempty"`
	Subscription  string                       `url:"subscription,omitempty" json:"subscription,omitempty"`
}

type PaymentListResultMetaCursors struct {
	After  string `url:"after,omitempty" json:"after,omitempty"`
	Before string `url:"before,omitempty" json:"before,omitempty"`
}

type PaymentListResultMeta struct {
	Cursors *PaymentListResultMetaCursors `url:"cursors,omitempty" json:"cursors,omitempty"`
	Limit   int                           `url:"limit,omitempty" json:"limit,omitempty"`
}

type PaymentListResult struct {
	Payments []Payment             `json:"payments"`
	Meta     PaymentListResultMeta `url:"meta,omitempty" json:"meta,omitempty"`
}

// List
// Returns a [cursor-paginated](#api-usage-cursor-pagination) list of your
// payments.
func (s *PaymentServiceImpl) List(ctx context.Context, p PaymentListParams, opts ...RequestOption) (*PaymentListResult, error) {
	uri, err := url.Parse(fmt.Sprintf(s.config.Endpoint() + "/payments"))
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
	req.Header.Set("GoCardless-Client-Version", "2.9.0")
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
		*PaymentListResult
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

	if result.PaymentListResult == nil {
		return nil, errors.New("missing result")
	}

	return result.PaymentListResult, nil
}

type PaymentListPagingIterator struct {
	cursor         string
	response       *PaymentListResult
	params         PaymentListParams
	service        *PaymentServiceImpl
	requestOptions []RequestOption
}

func (c *PaymentListPagingIterator) Next() bool {
	if c.cursor == "" && c.response != nil {
		return false
	}

	return true
}

func (c *PaymentListPagingIterator) Value(ctx context.Context) (*PaymentListResult, error) {
	if !c.Next() {
		return c.response, nil
	}

	s := c.service
	p := c.params
	p.After = c.cursor

	uri, err := url.Parse(fmt.Sprintf(s.config.Endpoint() + "/payments"))

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
	req.Header.Set("GoCardless-Client-Version", "2.9.0")
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
		*PaymentListResult
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

	if result.PaymentListResult == nil {
		return nil, errors.New("missing result")
	}

	c.response = result.PaymentListResult
	c.cursor = c.response.Meta.Cursors.After
	return c.response, nil
}

func (s *PaymentServiceImpl) All(ctx context.Context,
	p PaymentListParams,
	opts ...RequestOption) *PaymentListPagingIterator {
	return &PaymentListPagingIterator{
		params:         p,
		service:        s,
		requestOptions: opts,
	}
}

// Get
// Retrieves the details of a single existing payment.
func (s *PaymentServiceImpl) Get(ctx context.Context, identity string, opts ...RequestOption) (*Payment, error) {
	uri, err := url.Parse(fmt.Sprintf(s.config.Endpoint()+"/payments/%v",
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
	req.Header.Set("GoCardless-Client-Version", "2.9.0")
	req.Header.Set("User-Agent", userAgent)

	for key, value := range o.headers {
		req.Header.Set(key, value)
	}

	client := s.config.Client()
	if client == nil {
		client = http.DefaultClient
	}

	var result struct {
		Err     *APIError `json:"error"`
		Payment *Payment  `json:"payments"`
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

	if result.Payment == nil {
		return nil, errors.New("missing result")
	}

	return result.Payment, nil
}

// PaymentUpdateParams parameters
type PaymentUpdateParams struct {
	Metadata        map[string]interface{} `url:"metadata,omitempty" json:"metadata,omitempty"`
	RetryIfPossible bool                   `url:"retry_if_possible,omitempty" json:"retry_if_possible,omitempty"`
}

// Update
// Updates a payment object. This accepts only the metadata parameter.
func (s *PaymentServiceImpl) Update(ctx context.Context, identity string, p PaymentUpdateParams, opts ...RequestOption) (*Payment, error) {
	uri, err := url.Parse(fmt.Sprintf(s.config.Endpoint()+"/payments/%v",
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
		"payments": p,
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
	req.Header.Set("GoCardless-Client-Version", "2.9.0")
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
		Err     *APIError `json:"error"`
		Payment *Payment  `json:"payments"`
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

	if result.Payment == nil {
		return nil, errors.New("missing result")
	}

	return result.Payment, nil
}

// PaymentCancelParams parameters
type PaymentCancelParams struct {
	Metadata map[string]interface{} `url:"metadata,omitempty" json:"metadata,omitempty"`
}

// Cancel
// Cancels the payment if it has not already been submitted to the banks. Any
// metadata supplied to this endpoint will be stored on the payment cancellation
// event it causes.
//
// This will fail with a `cancellation_failed` error unless the payment's status
// is `pending_submission`.
func (s *PaymentServiceImpl) Cancel(ctx context.Context, identity string, p PaymentCancelParams, opts ...RequestOption) (*Payment, error) {
	uri, err := url.Parse(fmt.Sprintf(s.config.Endpoint()+"/payments/%v/actions/cancel",
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
	req.Header.Set("GoCardless-Client-Version", "2.9.0")
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
		Err     *APIError `json:"error"`
		Payment *Payment  `json:"payments"`
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

	if result.Payment == nil {
		return nil, errors.New("missing result")
	}

	return result.Payment, nil
}

// PaymentRetryParams parameters
type PaymentRetryParams struct {
	ChargeDate string                 `url:"charge_date,omitempty" json:"charge_date,omitempty"`
	Metadata   map[string]interface{} `url:"metadata,omitempty" json:"metadata,omitempty"`
}

// Retry
// <a name="retry_failed"></a>Retries a failed payment if the underlying mandate
// is active. You will receive a `resubmission_requested` webhook, but after
// that retrying the payment follows the same process as its initial creation,
// so you will receive a `submitted` webhook, followed by a `confirmed` or
// `failed` event. Any metadata supplied to this endpoint will be stored against
// the payment submission event it causes.
//
// This will return a `retry_failed` error if the payment has not failed.
//
// Payments can be retried up to 3 times.
func (s *PaymentServiceImpl) Retry(ctx context.Context, identity string, p PaymentRetryParams, opts ...RequestOption) (*Payment, error) {
	uri, err := url.Parse(fmt.Sprintf(s.config.Endpoint()+"/payments/%v/actions/retry",
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
	req.Header.Set("GoCardless-Client-Version", "2.9.0")
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
		Err     *APIError `json:"error"`
		Payment *Payment  `json:"payments"`
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

	if result.Payment == nil {
		return nil, errors.New("missing result")
	}

	return result.Payment, nil
}
