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

// InstalmentScheduleService manages instalment_schedules
type InstalmentScheduleServiceImpl struct {
	config Config
}

type InstalmentScheduleLinks struct {
	Customer string   `url:"customer,omitempty" json:"customer,omitempty"`
	Mandate  string   `url:"mandate,omitempty" json:"mandate,omitempty"`
	Payments []string `url:"payments,omitempty" json:"payments,omitempty"`
}

// InstalmentSchedule model
type InstalmentSchedule struct {
	CreatedAt     string                   `url:"created_at,omitempty" json:"created_at,omitempty"`
	Currency      string                   `url:"currency,omitempty" json:"currency,omitempty"`
	Id            string                   `url:"id,omitempty" json:"id,omitempty"`
	Links         *InstalmentScheduleLinks `url:"links,omitempty" json:"links,omitempty"`
	Metadata      map[string]interface{}   `url:"metadata,omitempty" json:"metadata,omitempty"`
	Name          string                   `url:"name,omitempty" json:"name,omitempty"`
	PaymentErrors map[string]interface{}   `url:"payment_errors,omitempty" json:"payment_errors,omitempty"`
	Status        string                   `url:"status,omitempty" json:"status,omitempty"`
	TotalAmount   int                      `url:"total_amount,omitempty" json:"total_amount,omitempty"`
}

type InstalmentScheduleService interface {
	CreateWithDates(ctx context.Context, p InstalmentScheduleCreateWithDatesParams, opts ...RequestOption) (*InstalmentSchedule, error)
	CreateWithSchedule(ctx context.Context, p InstalmentScheduleCreateWithScheduleParams, opts ...RequestOption) (*InstalmentSchedule, error)
	List(ctx context.Context, p InstalmentScheduleListParams, opts ...RequestOption) (*InstalmentScheduleListResult, error)
	All(ctx context.Context, p InstalmentScheduleListParams, opts ...RequestOption) *InstalmentScheduleListPagingIterator
	Get(ctx context.Context, identity string, opts ...RequestOption) (*InstalmentSchedule, error)
	Update(ctx context.Context, identity string, p InstalmentScheduleUpdateParams, opts ...RequestOption) (*InstalmentSchedule, error)
	Cancel(ctx context.Context, identity string, p InstalmentScheduleCancelParams, opts ...RequestOption) (*InstalmentSchedule, error)
}

type InstalmentScheduleCreateWithDatesParamsInstalments struct {
	Amount      int    `url:"amount,omitempty" json:"amount,omitempty"`
	ChargeDate  string `url:"charge_date,omitempty" json:"charge_date,omitempty"`
	Description string `url:"description,omitempty" json:"description,omitempty"`
}

type InstalmentScheduleCreateWithDatesParamsLinks struct {
	Mandate string `url:"mandate,omitempty" json:"mandate,omitempty"`
}

// InstalmentScheduleCreateWithDatesParams parameters
type InstalmentScheduleCreateWithDatesParams struct {
	AppFee           int                                                  `url:"app_fee,omitempty" json:"app_fee,omitempty"`
	Currency         string                                               `url:"currency,omitempty" json:"currency,omitempty"`
	Instalments      []InstalmentScheduleCreateWithDatesParamsInstalments `url:"instalments,omitempty" json:"instalments,omitempty"`
	Links            InstalmentScheduleCreateWithDatesParamsLinks         `url:"links,omitempty" json:"links,omitempty"`
	Metadata         map[string]interface{}                               `url:"metadata,omitempty" json:"metadata,omitempty"`
	Name             string                                               `url:"name,omitempty" json:"name,omitempty"`
	PaymentReference string                                               `url:"payment_reference,omitempty" json:"payment_reference,omitempty"`
	RetryIfPossible  bool                                                 `url:"retry_if_possible,omitempty" json:"retry_if_possible,omitempty"`
	TotalAmount      int                                                  `url:"total_amount,omitempty" json:"total_amount,omitempty"`
}

// CreateWithDates
// Creates a new instalment schedule object, along with the associated payments.
// This
// API is recommended if you know the specific dates you wish to charge.
// Otherwise,
// please check out the [scheduling
// version](#instalment-schedules-create-with-schedule).
//
// The `instalments` property is an array of payment properties (`amount` and
// `charge_date`).
//
// It can take quite a while to create the associated payments, so the API will
// return
// the status as `pending` initially. When processing has completed, a
// subsequent GET
// request for the instalment schedule will either have the status `success` and
// link
// to the created payments, or the status `error` and detailed information about
// the
// failures.
func (s *InstalmentScheduleServiceImpl) CreateWithDates(ctx context.Context, p InstalmentScheduleCreateWithDatesParams, opts ...RequestOption) (*InstalmentSchedule, error) {
	uri, err := url.Parse(fmt.Sprintf(s.config.Endpoint() + "/instalment_schedules"))
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
		Err                *APIError           `json:"error"`
		InstalmentSchedule *InstalmentSchedule `json:"instalment_schedules"`
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

	if result.InstalmentSchedule == nil {
		return nil, errors.New("missing result")
	}

	return result.InstalmentSchedule, nil
}

type InstalmentScheduleCreateWithScheduleParamsInstalments struct {
	Amounts      []int  `url:"amounts,omitempty" json:"amounts,omitempty"`
	Interval     int    `url:"interval,omitempty" json:"interval,omitempty"`
	IntervalUnit string `url:"interval_unit,omitempty" json:"interval_unit,omitempty"`
	StartDate    string `url:"start_date,omitempty" json:"start_date,omitempty"`
}

type InstalmentScheduleCreateWithScheduleParamsLinks struct {
	Mandate string `url:"mandate,omitempty" json:"mandate,omitempty"`
}

// InstalmentScheduleCreateWithScheduleParams parameters
type InstalmentScheduleCreateWithScheduleParams struct {
	AppFee           int                                                   `url:"app_fee,omitempty" json:"app_fee,omitempty"`
	Currency         string                                                `url:"currency,omitempty" json:"currency,omitempty"`
	Instalments      InstalmentScheduleCreateWithScheduleParamsInstalments `url:"instalments,omitempty" json:"instalments,omitempty"`
	Links            InstalmentScheduleCreateWithScheduleParamsLinks       `url:"links,omitempty" json:"links,omitempty"`
	Metadata         map[string]interface{}                                `url:"metadata,omitempty" json:"metadata,omitempty"`
	Name             string                                                `url:"name,omitempty" json:"name,omitempty"`
	PaymentReference string                                                `url:"payment_reference,omitempty" json:"payment_reference,omitempty"`
	RetryIfPossible  bool                                                  `url:"retry_if_possible,omitempty" json:"retry_if_possible,omitempty"`
	TotalAmount      int                                                   `url:"total_amount,omitempty" json:"total_amount,omitempty"`
}

// CreateWithSchedule
// Creates a new instalment schedule object, along with the associated payments.
// This
// API is recommended if you wish to use the GoCardless scheduling logic. For
// finer
// control over the individual dates, please check out the [alternative
// version](#instalment-schedules-create-with-dates).
//
// It can take quite a while to create the associated payments, so the API will
// return
// the status as `pending` initially. When processing has completed, a
// subsequent
// GET request for the instalment schedule will either have the status `success`
// and link to
// the created payments, or the status `error` and detailed information about
// the
// failures.
func (s *InstalmentScheduleServiceImpl) CreateWithSchedule(ctx context.Context, p InstalmentScheduleCreateWithScheduleParams, opts ...RequestOption) (*InstalmentSchedule, error) {
	uri, err := url.Parse(fmt.Sprintf(s.config.Endpoint() + "/instalment_schedules"))
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
		Err                *APIError           `json:"error"`
		InstalmentSchedule *InstalmentSchedule `json:"instalment_schedules"`
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

	if result.InstalmentSchedule == nil {
		return nil, errors.New("missing result")
	}

	return result.InstalmentSchedule, nil
}

type InstalmentScheduleListParamsCreatedAt struct {
	Gt  string `url:"gt,omitempty" json:"gt,omitempty"`
	Gte string `url:"gte,omitempty" json:"gte,omitempty"`
	Lt  string `url:"lt,omitempty" json:"lt,omitempty"`
	Lte string `url:"lte,omitempty" json:"lte,omitempty"`
}

// InstalmentScheduleListParams parameters
type InstalmentScheduleListParams struct {
	After     string                                 `url:"after,omitempty" json:"after,omitempty"`
	Before    string                                 `url:"before,omitempty" json:"before,omitempty"`
	CreatedAt *InstalmentScheduleListParamsCreatedAt `url:"created_at,omitempty" json:"created_at,omitempty"`
	Customer  string                                 `url:"customer,omitempty" json:"customer,omitempty"`
	Limit     int                                    `url:"limit,omitempty" json:"limit,omitempty"`
	Mandate   string                                 `url:"mandate,omitempty" json:"mandate,omitempty"`
	Status    []string                               `url:"status,omitempty" json:"status,omitempty"`
}

type InstalmentScheduleListResultMetaCursors struct {
	After  string `url:"after,omitempty" json:"after,omitempty"`
	Before string `url:"before,omitempty" json:"before,omitempty"`
}

type InstalmentScheduleListResultMeta struct {
	Cursors *InstalmentScheduleListResultMetaCursors `url:"cursors,omitempty" json:"cursors,omitempty"`
	Limit   int                                      `url:"limit,omitempty" json:"limit,omitempty"`
}

type InstalmentScheduleListResult struct {
	InstalmentSchedules []InstalmentSchedule             `json:"instalment_schedules"`
	Meta                InstalmentScheduleListResultMeta `url:"meta,omitempty" json:"meta,omitempty"`
}

// List
// Returns a [cursor-paginated](#api-usage-cursor-pagination) list of your
// instalment schedules.
func (s *InstalmentScheduleServiceImpl) List(ctx context.Context, p InstalmentScheduleListParams, opts ...RequestOption) (*InstalmentScheduleListResult, error) {
	uri, err := url.Parse(fmt.Sprintf(s.config.Endpoint() + "/instalment_schedules"))
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
		*InstalmentScheduleListResult
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

	if result.InstalmentScheduleListResult == nil {
		return nil, errors.New("missing result")
	}

	return result.InstalmentScheduleListResult, nil
}

type InstalmentScheduleListPagingIterator struct {
	cursor         string
	response       *InstalmentScheduleListResult
	params         InstalmentScheduleListParams
	service        *InstalmentScheduleServiceImpl
	requestOptions []RequestOption
}

func (c *InstalmentScheduleListPagingIterator) Next() bool {
	if c.cursor == "" && c.response != nil {
		return false
	}

	return true
}

func (c *InstalmentScheduleListPagingIterator) Value(ctx context.Context) (*InstalmentScheduleListResult, error) {
	if !c.Next() {
		return c.response, nil
	}

	s := c.service
	p := c.params
	p.After = c.cursor

	uri, err := url.Parse(fmt.Sprintf(s.config.Endpoint() + "/instalment_schedules"))

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
		*InstalmentScheduleListResult
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

	if result.InstalmentScheduleListResult == nil {
		return nil, errors.New("missing result")
	}

	c.response = result.InstalmentScheduleListResult
	c.cursor = c.response.Meta.Cursors.After
	return c.response, nil
}

func (s *InstalmentScheduleServiceImpl) All(ctx context.Context,
	p InstalmentScheduleListParams,
	opts ...RequestOption) *InstalmentScheduleListPagingIterator {
	return &InstalmentScheduleListPagingIterator{
		params:         p,
		service:        s,
		requestOptions: opts,
	}
}

// Get
// Retrieves the details of an existing instalment schedule.
func (s *InstalmentScheduleServiceImpl) Get(ctx context.Context, identity string, opts ...RequestOption) (*InstalmentSchedule, error) {
	uri, err := url.Parse(fmt.Sprintf(s.config.Endpoint()+"/instalment_schedules/%v",
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
		Err                *APIError           `json:"error"`
		InstalmentSchedule *InstalmentSchedule `json:"instalment_schedules"`
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

	if result.InstalmentSchedule == nil {
		return nil, errors.New("missing result")
	}

	return result.InstalmentSchedule, nil
}

// InstalmentScheduleUpdateParams parameters
type InstalmentScheduleUpdateParams struct {
	Metadata map[string]interface{} `url:"metadata,omitempty" json:"metadata,omitempty"`
}

// Update
// Updates an instalment schedule. This accepts only the metadata parameter.
func (s *InstalmentScheduleServiceImpl) Update(ctx context.Context, identity string, p InstalmentScheduleUpdateParams, opts ...RequestOption) (*InstalmentSchedule, error) {
	uri, err := url.Parse(fmt.Sprintf(s.config.Endpoint()+"/instalment_schedules/%v",
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
		"instalment_schedules": p,
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
		Err                *APIError           `json:"error"`
		InstalmentSchedule *InstalmentSchedule `json:"instalment_schedules"`
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

	if result.InstalmentSchedule == nil {
		return nil, errors.New("missing result")
	}

	return result.InstalmentSchedule, nil
}

// InstalmentScheduleCancelParams parameters
type InstalmentScheduleCancelParams struct {
}

// Cancel
// Immediately cancels an instalment schedule; no further payments will be
// collected for it.
//
// This will fail with a `cancellation_failed` error if the instalment schedule
// is already cancelled or has completed.
func (s *InstalmentScheduleServiceImpl) Cancel(ctx context.Context, identity string, p InstalmentScheduleCancelParams, opts ...RequestOption) (*InstalmentSchedule, error) {
	uri, err := url.Parse(fmt.Sprintf(s.config.Endpoint()+"/instalment_schedules/%v/actions/cancel",
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
		Err                *APIError           `json:"error"`
		InstalmentSchedule *InstalmentSchedule `json:"instalment_schedules"`
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

	if result.InstalmentSchedule == nil {
		return nil, errors.New("missing result")
	}

	return result.InstalmentSchedule, nil
}
