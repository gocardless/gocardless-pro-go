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

// SubscriptionService manages subscriptions
type SubscriptionServiceImpl struct {
	config Config
}

type SubscriptionLinks struct {
	Mandate string `url:"mandate,omitempty" json:"mandate,omitempty"`
}

type SubscriptionUpcomingPayments struct {
	Amount     int    `url:"amount,omitempty" json:"amount,omitempty"`
	ChargeDate string `url:"charge_date,omitempty" json:"charge_date,omitempty"`
}

// Subscription model
type Subscription struct {
	Amount                        int                            `url:"amount,omitempty" json:"amount,omitempty"`
	AppFee                        int                            `url:"app_fee,omitempty" json:"app_fee,omitempty"`
	Count                         int                            `url:"count,omitempty" json:"count,omitempty"`
	CreatedAt                     string                         `url:"created_at,omitempty" json:"created_at,omitempty"`
	Currency                      string                         `url:"currency,omitempty" json:"currency,omitempty"`
	DayOfMonth                    int                            `url:"day_of_month,omitempty" json:"day_of_month,omitempty"`
	EarliestChargeDateAfterResume string                         `url:"earliest_charge_date_after_resume,omitempty" json:"earliest_charge_date_after_resume,omitempty"`
	EndDate                       string                         `url:"end_date,omitempty" json:"end_date,omitempty"`
	Id                            string                         `url:"id,omitempty" json:"id,omitempty"`
	Interval                      int                            `url:"interval,omitempty" json:"interval,omitempty"`
	IntervalUnit                  string                         `url:"interval_unit,omitempty" json:"interval_unit,omitempty"`
	Links                         *SubscriptionLinks             `url:"links,omitempty" json:"links,omitempty"`
	Metadata                      map[string]interface{}         `url:"metadata,omitempty" json:"metadata,omitempty"`
	Month                         string                         `url:"month,omitempty" json:"month,omitempty"`
	Name                          string                         `url:"name,omitempty" json:"name,omitempty"`
	PaymentReference              string                         `url:"payment_reference,omitempty" json:"payment_reference,omitempty"`
	RetryIfPossible               bool                           `url:"retry_if_possible,omitempty" json:"retry_if_possible,omitempty"`
	StartDate                     string                         `url:"start_date,omitempty" json:"start_date,omitempty"`
	Status                        string                         `url:"status,omitempty" json:"status,omitempty"`
	UpcomingPayments              []SubscriptionUpcomingPayments `url:"upcoming_payments,omitempty" json:"upcoming_payments,omitempty"`
}

type SubscriptionService interface {
	Create(ctx context.Context, p SubscriptionCreateParams, opts ...RequestOption) (*Subscription, error)
	List(ctx context.Context, p SubscriptionListParams, opts ...RequestOption) (*SubscriptionListResult, error)
	All(ctx context.Context, p SubscriptionListParams, opts ...RequestOption) *SubscriptionListPagingIterator
	Get(ctx context.Context, identity string, opts ...RequestOption) (*Subscription, error)
	Update(ctx context.Context, identity string, p SubscriptionUpdateParams, opts ...RequestOption) (*Subscription, error)
	Pause(ctx context.Context, identity string, p SubscriptionPauseParams, opts ...RequestOption) (*Subscription, error)
	Resume(ctx context.Context, identity string, p SubscriptionResumeParams, opts ...RequestOption) (*Subscription, error)
	Cancel(ctx context.Context, identity string, p SubscriptionCancelParams, opts ...RequestOption) (*Subscription, error)
}

type SubscriptionCreateParamsLinks struct {
	Mandate string `url:"mandate,omitempty" json:"mandate,omitempty"`
}

// SubscriptionCreateParams parameters
type SubscriptionCreateParams struct {
	Amount           int                           `url:"amount,omitempty" json:"amount,omitempty"`
	AppFee           int                           `url:"app_fee,omitempty" json:"app_fee,omitempty"`
	Count            int                           `url:"count,omitempty" json:"count,omitempty"`
	Currency         string                        `url:"currency,omitempty" json:"currency,omitempty"`
	DayOfMonth       int                           `url:"day_of_month,omitempty" json:"day_of_month,omitempty"`
	EndDate          string                        `url:"end_date,omitempty" json:"end_date,omitempty"`
	Interval         int                           `url:"interval,omitempty" json:"interval,omitempty"`
	IntervalUnit     string                        `url:"interval_unit,omitempty" json:"interval_unit,omitempty"`
	Links            SubscriptionCreateParamsLinks `url:"links,omitempty" json:"links,omitempty"`
	Metadata         map[string]interface{}        `url:"metadata,omitempty" json:"metadata,omitempty"`
	Month            string                        `url:"month,omitempty" json:"month,omitempty"`
	Name             string                        `url:"name,omitempty" json:"name,omitempty"`
	PaymentReference string                        `url:"payment_reference,omitempty" json:"payment_reference,omitempty"`
	RetryIfPossible  bool                          `url:"retry_if_possible,omitempty" json:"retry_if_possible,omitempty"`
	StartDate        string                        `url:"start_date,omitempty" json:"start_date,omitempty"`
}

// Create
// Creates a new subscription object
func (s *SubscriptionServiceImpl) Create(ctx context.Context, p SubscriptionCreateParams, opts ...RequestOption) (*Subscription, error) {
	uri, err := url.Parse(fmt.Sprintf(s.config.Endpoint() + "/subscriptions"))
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
		"subscriptions": p,
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
	req.Header.Set("GoCardless-Client-Version", "3.7.0")
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
		Subscription *Subscription `json:"subscriptions"`
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

	if result.Subscription == nil {
		return nil, errors.New("missing result")
	}

	return result.Subscription, nil
}

type SubscriptionListParamsCreatedAt struct {
	Gt  string `url:"gt,omitempty" json:"gt,omitempty"`
	Gte string `url:"gte,omitempty" json:"gte,omitempty"`
	Lt  string `url:"lt,omitempty" json:"lt,omitempty"`
	Lte string `url:"lte,omitempty" json:"lte,omitempty"`
}

// SubscriptionListParams parameters
type SubscriptionListParams struct {
	After     string                           `url:"after,omitempty" json:"after,omitempty"`
	Before    string                           `url:"before,omitempty" json:"before,omitempty"`
	CreatedAt *SubscriptionListParamsCreatedAt `url:"created_at,omitempty" json:"created_at,omitempty"`
	Customer  string                           `url:"customer,omitempty" json:"customer,omitempty"`
	Limit     int                              `url:"limit,omitempty" json:"limit,omitempty"`
	Mandate   string                           `url:"mandate,omitempty" json:"mandate,omitempty"`
	Status    []string                         `url:"status,omitempty" json:"status,omitempty"`
}

type SubscriptionListResultMetaCursors struct {
	After  string `url:"after,omitempty" json:"after,omitempty"`
	Before string `url:"before,omitempty" json:"before,omitempty"`
}

type SubscriptionListResultMeta struct {
	Cursors *SubscriptionListResultMetaCursors `url:"cursors,omitempty" json:"cursors,omitempty"`
	Limit   int                                `url:"limit,omitempty" json:"limit,omitempty"`
}

type SubscriptionListResult struct {
	Subscriptions []Subscription             `json:"subscriptions"`
	Meta          SubscriptionListResultMeta `url:"meta,omitempty" json:"meta,omitempty"`
}

// List
// Returns a [cursor-paginated](#api-usage-cursor-pagination) list of your
// subscriptions. Please note if the subscriptions are related to customers who
// have been removed, they will not be shown in the response.
func (s *SubscriptionServiceImpl) List(ctx context.Context, p SubscriptionListParams, opts ...RequestOption) (*SubscriptionListResult, error) {
	uri, err := url.Parse(fmt.Sprintf(s.config.Endpoint() + "/subscriptions"))
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
	req.Header.Set("GoCardless-Client-Version", "3.7.0")
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
		*SubscriptionListResult
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

	if result.SubscriptionListResult == nil {
		return nil, errors.New("missing result")
	}

	return result.SubscriptionListResult, nil
}

type SubscriptionListPagingIterator struct {
	cursor         string
	response       *SubscriptionListResult
	params         SubscriptionListParams
	service        *SubscriptionServiceImpl
	requestOptions []RequestOption
}

func (c *SubscriptionListPagingIterator) Next() bool {
	if c.cursor == "" && c.response != nil {
		return false
	}

	return true
}

func (c *SubscriptionListPagingIterator) Value(ctx context.Context) (*SubscriptionListResult, error) {
	if !c.Next() {
		return c.response, nil
	}

	s := c.service
	p := c.params
	p.After = c.cursor

	uri, err := url.Parse(fmt.Sprintf(s.config.Endpoint() + "/subscriptions"))

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
	req.Header.Set("GoCardless-Client-Version", "3.7.0")
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
		*SubscriptionListResult
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

	if result.SubscriptionListResult == nil {
		return nil, errors.New("missing result")
	}

	c.response = result.SubscriptionListResult
	c.cursor = c.response.Meta.Cursors.After
	return c.response, nil
}

func (s *SubscriptionServiceImpl) All(ctx context.Context,
	p SubscriptionListParams,
	opts ...RequestOption) *SubscriptionListPagingIterator {
	return &SubscriptionListPagingIterator{
		params:         p,
		service:        s,
		requestOptions: opts,
	}
}

// Get
// Retrieves the details of a single subscription.
func (s *SubscriptionServiceImpl) Get(ctx context.Context, identity string, opts ...RequestOption) (*Subscription, error) {
	uri, err := url.Parse(fmt.Sprintf(s.config.Endpoint()+"/subscriptions/%v",
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
	req.Header.Set("GoCardless-Client-Version", "3.7.0")
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
		Subscription *Subscription `json:"subscriptions"`
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

	if result.Subscription == nil {
		return nil, errors.New("missing result")
	}

	return result.Subscription, nil
}

// SubscriptionUpdateParams parameters
type SubscriptionUpdateParams struct {
	Amount           int                    `url:"amount,omitempty" json:"amount,omitempty"`
	AppFee           int                    `url:"app_fee,omitempty" json:"app_fee,omitempty"`
	Metadata         map[string]interface{} `url:"metadata,omitempty" json:"metadata,omitempty"`
	Name             string                 `url:"name,omitempty" json:"name,omitempty"`
	PaymentReference string                 `url:"payment_reference,omitempty" json:"payment_reference,omitempty"`
	RetryIfPossible  bool                   `url:"retry_if_possible,omitempty" json:"retry_if_possible,omitempty"`
}

// Update
// Updates a subscription object.
//
// This fails with:
//
// - `validation_failed` if invalid data is provided when attempting to update a
// subscription.
//
// - `subscription_not_active` if the subscription is no longer active.
//
// - `subscription_already_ended` if the subscription has taken all payments.
//
// - `mandate_payments_require_approval` if the amount is being changed and the
// mandate requires approval.
//
// - `number_of_subscription_amendments_exceeded` error if the subscription
// amount has already been changed 10 times.
//
// - `forbidden` if the amount is being changed, and the subscription was
// created by an app and you are not authenticated as that app, or if the
// subscription was not created by an app and you are authenticated as an app
//
// - `resource_created_by_another_app` if the app fee is being changed, and the
// subscription was created by an app other than the app you are authenticated
// as
func (s *SubscriptionServiceImpl) Update(ctx context.Context, identity string, p SubscriptionUpdateParams, opts ...RequestOption) (*Subscription, error) {
	uri, err := url.Parse(fmt.Sprintf(s.config.Endpoint()+"/subscriptions/%v",
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
		"subscriptions": p,
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
	req.Header.Set("GoCardless-Client-Version", "3.7.0")
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
		Subscription *Subscription `json:"subscriptions"`
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

	if result.Subscription == nil {
		return nil, errors.New("missing result")
	}

	return result.Subscription, nil
}

// SubscriptionPauseParams parameters
type SubscriptionPauseParams struct {
	Metadata    map[string]interface{} `url:"metadata,omitempty" json:"metadata,omitempty"`
	PauseCycles int                    `url:"pause_cycles,omitempty" json:"pause_cycles,omitempty"`
}

// Pause
// Pause a subscription object.
// No payments will be created until it is resumed.
//
// This can only be used when a subscription is collecting a fixed number of
// payments (created using `count`),
// when they continue forever (created without `count` or `end_date`) or
// the subscription is already paused for a number of cycles.
//
// When `pause_cycles` is omitted the subscription is paused until the [resume
// endpoint](#subscriptions-resume-a-subscription) is called.
// If the subscription is collecting a fixed number of payments, `end_date` will
// be set to `null`.
// When paused indefinitely, `upcoming_payments` will be empty.
//
// When `pause_cycles` is provided the subscription will be paused for the
// number of cycles requested.
// If the subscription is collecting a fixed number of payments, `end_date` will
// be set to a new value.
// When paused for a number of cycles, `upcoming_payments` will still contain
// the upcoming charge dates.
//
// This fails with:
//
// - `forbidden` if the subscription was created by an app and you are not
// authenticated as that app, or if the subscription was not created by an app
// and you are authenticated as an app
//
// - `validation_failed` if invalid data is provided when attempting to pause a
// subscription.
//
// - `subscription_paused_cannot_update_cycles` if the subscription is already
// paused for a number of cycles and the request provides a value for
// `pause_cycle`.
//
// - `subscription_cannot_be_paused` if the subscription cannot be paused.
//
// - `subscription_already_ended` if the subscription has taken all payments.
//
// - `pause_cycles_must_be_greater_than_or_equal_to` if the provided value for
// `pause_cycles` cannot be satisfied.
func (s *SubscriptionServiceImpl) Pause(ctx context.Context, identity string, p SubscriptionPauseParams, opts ...RequestOption) (*Subscription, error) {
	uri, err := url.Parse(fmt.Sprintf(s.config.Endpoint()+"/subscriptions/%v/actions/pause",
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
	req.Header.Set("GoCardless-Client-Version", "3.7.0")
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
		Subscription *Subscription `json:"subscriptions"`
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

	if result.Subscription == nil {
		return nil, errors.New("missing result")
	}

	return result.Subscription, nil
}

// SubscriptionResumeParams parameters
type SubscriptionResumeParams struct {
	Metadata map[string]interface{} `url:"metadata,omitempty" json:"metadata,omitempty"`
}

// Resume
// Resume a subscription object.
// Payments will start to be created again based on the subscriptions recurrence
// rules.
// The `charge_date` on the next payment will be the same as the subscriptions
// `earliest_charge_date_after_resume`
//
// This fails with:
//
// - `forbidden` if the subscription was created by an app and you are not
// authenticated as that app, or if the subscription was not created by an app
// and you are authenticated as an app
//
// - `validation_failed` if invalid data is provided when attempting to resume a
// subscription.
//
// - `subscription_not_paused` if the subscription is not paused.
func (s *SubscriptionServiceImpl) Resume(ctx context.Context, identity string, p SubscriptionResumeParams, opts ...RequestOption) (*Subscription, error) {
	uri, err := url.Parse(fmt.Sprintf(s.config.Endpoint()+"/subscriptions/%v/actions/resume",
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
	req.Header.Set("GoCardless-Client-Version", "3.7.0")
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
		Subscription *Subscription `json:"subscriptions"`
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

	if result.Subscription == nil {
		return nil, errors.New("missing result")
	}

	return result.Subscription, nil
}

// SubscriptionCancelParams parameters
type SubscriptionCancelParams struct {
	Metadata map[string]interface{} `url:"metadata,omitempty" json:"metadata,omitempty"`
}

// Cancel
// Immediately cancels a subscription; no more payments will be created under
// it. Any metadata supplied to this endpoint will be stored on the payment
// cancellation event it causes.
//
// This will fail with a cancellation_failed error if the subscription is
// already cancelled or finished.
func (s *SubscriptionServiceImpl) Cancel(ctx context.Context, identity string, p SubscriptionCancelParams, opts ...RequestOption) (*Subscription, error) {
	uri, err := url.Parse(fmt.Sprintf(s.config.Endpoint()+"/subscriptions/%v/actions/cancel",
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
	req.Header.Set("GoCardless-Client-Version", "3.7.0")
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
		Subscription *Subscription `json:"subscriptions"`
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

	if result.Subscription == nil {
		return nil, errors.New("missing result")
	}

	return result.Subscription, nil
}
