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

// PayoutService manages payouts
type PayoutServiceImpl struct {
	config Config
}

type PayoutFx struct {
	EstimatedExchangeRate string `url:"estimated_exchange_rate,omitempty" json:"estimated_exchange_rate,omitempty"`
	ExchangeRate          string `url:"exchange_rate,omitempty" json:"exchange_rate,omitempty"`
	FxAmount              int    `url:"fx_amount,omitempty" json:"fx_amount,omitempty"`
	FxCurrency            string `url:"fx_currency,omitempty" json:"fx_currency,omitempty"`
}

type PayoutLinks struct {
	Creditor            string `url:"creditor,omitempty" json:"creditor,omitempty"`
	CreditorBankAccount string `url:"creditor_bank_account,omitempty" json:"creditor_bank_account,omitempty"`
}

// Payout model
type Payout struct {
	Amount       int                    `url:"amount,omitempty" json:"amount,omitempty"`
	ArrivalDate  string                 `url:"arrival_date,omitempty" json:"arrival_date,omitempty"`
	CreatedAt    string                 `url:"created_at,omitempty" json:"created_at,omitempty"`
	Currency     string                 `url:"currency,omitempty" json:"currency,omitempty"`
	DeductedFees int                    `url:"deducted_fees,omitempty" json:"deducted_fees,omitempty"`
	Fx           *PayoutFx              `url:"fx,omitempty" json:"fx,omitempty"`
	Id           string                 `url:"id,omitempty" json:"id,omitempty"`
	Links        *PayoutLinks           `url:"links,omitempty" json:"links,omitempty"`
	Metadata     map[string]interface{} `url:"metadata,omitempty" json:"metadata,omitempty"`
	PayoutType   string                 `url:"payout_type,omitempty" json:"payout_type,omitempty"`
	Reference    string                 `url:"reference,omitempty" json:"reference,omitempty"`
	Status       string                 `url:"status,omitempty" json:"status,omitempty"`
	TaxCurrency  string                 `url:"tax_currency,omitempty" json:"tax_currency,omitempty"`
}

type PayoutService interface {
	List(ctx context.Context, p PayoutListParams, opts ...RequestOption) (*PayoutListResult, error)
	All(ctx context.Context, p PayoutListParams, opts ...RequestOption) *PayoutListPagingIterator
	Get(ctx context.Context, identity string, opts ...RequestOption) (*Payout, error)
	Update(ctx context.Context, identity string, p PayoutUpdateParams, opts ...RequestOption) (*Payout, error)
}

type PayoutListParamsCreatedAt struct {
	Gt  string `url:"gt,omitempty" json:"gt,omitempty"`
	Gte string `url:"gte,omitempty" json:"gte,omitempty"`
	Lt  string `url:"lt,omitempty" json:"lt,omitempty"`
	Lte string `url:"lte,omitempty" json:"lte,omitempty"`
}

// PayoutListParams parameters
type PayoutListParams struct {
	After               string                     `url:"after,omitempty" json:"after,omitempty"`
	Before              string                     `url:"before,omitempty" json:"before,omitempty"`
	CreatedAt           *PayoutListParamsCreatedAt `url:"created_at,omitempty" json:"created_at,omitempty"`
	Creditor            string                     `url:"creditor,omitempty" json:"creditor,omitempty"`
	CreditorBankAccount string                     `url:"creditor_bank_account,omitempty" json:"creditor_bank_account,omitempty"`
	Currency            string                     `url:"currency,omitempty" json:"currency,omitempty"`
	Limit               int                        `url:"limit,omitempty" json:"limit,omitempty"`
	Metadata            map[string]interface{}     `url:"metadata,omitempty" json:"metadata,omitempty"`
	PayoutType          string                     `url:"payout_type,omitempty" json:"payout_type,omitempty"`
	Reference           string                     `url:"reference,omitempty" json:"reference,omitempty"`
	Status              string                     `url:"status,omitempty" json:"status,omitempty"`
}

type PayoutListResultMetaCursors struct {
	After  string `url:"after,omitempty" json:"after,omitempty"`
	Before string `url:"before,omitempty" json:"before,omitempty"`
}

type PayoutListResultMeta struct {
	Cursors *PayoutListResultMetaCursors `url:"cursors,omitempty" json:"cursors,omitempty"`
	Limit   int                          `url:"limit,omitempty" json:"limit,omitempty"`
}

type PayoutListResult struct {
	Payouts []Payout             `json:"payouts"`
	Meta    PayoutListResultMeta `url:"meta,omitempty" json:"meta,omitempty"`
}

// List
// Returns a [cursor-paginated](#api-usage-cursor-pagination) list of your
// payouts.
func (s *PayoutServiceImpl) List(ctx context.Context, p PayoutListParams, opts ...RequestOption) (*PayoutListResult, error) {
	uri, err := url.Parse(fmt.Sprintf(s.config.Endpoint() + "/payouts"))
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
	req.Header.Set("GoCardless-Client-Version", "3.9.0")
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
		*PayoutListResult
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

	if result.PayoutListResult == nil {
		return nil, errors.New("missing result")
	}

	return result.PayoutListResult, nil
}

type PayoutListPagingIterator struct {
	cursor         string
	response       *PayoutListResult
	params         PayoutListParams
	service        *PayoutServiceImpl
	requestOptions []RequestOption
}

func (c *PayoutListPagingIterator) Next() bool {
	if c.cursor == "" && c.response != nil {
		return false
	}

	return true
}

func (c *PayoutListPagingIterator) Value(ctx context.Context) (*PayoutListResult, error) {
	if !c.Next() {
		return c.response, nil
	}

	s := c.service
	p := c.params
	p.After = c.cursor

	uri, err := url.Parse(fmt.Sprintf(s.config.Endpoint() + "/payouts"))

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
	req.Header.Set("GoCardless-Client-Version", "3.9.0")
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
		*PayoutListResult
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

	if result.PayoutListResult == nil {
		return nil, errors.New("missing result")
	}

	c.response = result.PayoutListResult
	c.cursor = c.response.Meta.Cursors.After
	return c.response, nil
}

func (s *PayoutServiceImpl) All(ctx context.Context,
	p PayoutListParams,
	opts ...RequestOption) *PayoutListPagingIterator {
	return &PayoutListPagingIterator{
		params:         p,
		service:        s,
		requestOptions: opts,
	}
}

// Get
// Retrieves the details of a single payout. For an example of how to reconcile
// the transactions in a payout, see [this
// guide](#events-reconciling-payouts-with-events).
func (s *PayoutServiceImpl) Get(ctx context.Context, identity string, opts ...RequestOption) (*Payout, error) {
	uri, err := url.Parse(fmt.Sprintf(s.config.Endpoint()+"/payouts/%v",
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
	req.Header.Set("GoCardless-Client-Version", "3.9.0")
	req.Header.Set("User-Agent", userAgent)

	for key, value := range o.headers {
		req.Header.Set(key, value)
	}

	client := s.config.Client()
	if client == nil {
		client = http.DefaultClient
	}

	var result struct {
		Err    *APIError `json:"error"`
		Payout *Payout   `json:"payouts"`
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

	if result.Payout == nil {
		return nil, errors.New("missing result")
	}

	return result.Payout, nil
}

// PayoutUpdateParams parameters
type PayoutUpdateParams struct {
	Metadata map[string]interface{} `url:"metadata,omitempty" json:"metadata,omitempty"`
}

// Update
// Updates a payout object. This accepts only the metadata parameter.
func (s *PayoutServiceImpl) Update(ctx context.Context, identity string, p PayoutUpdateParams, opts ...RequestOption) (*Payout, error) {
	uri, err := url.Parse(fmt.Sprintf(s.config.Endpoint()+"/payouts/%v",
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
		"payouts": p,
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
	req.Header.Set("GoCardless-Client-Version", "3.9.0")
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
		Err    *APIError `json:"error"`
		Payout *Payout   `json:"payouts"`
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

	if result.Payout == nil {
		return nil, errors.New("missing result")
	}

	return result.Payout, nil
}
