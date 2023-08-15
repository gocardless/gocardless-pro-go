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

// PayoutItemService manages payout_items
type PayoutItemServiceImpl struct {
	config Config
}

type PayoutItemLinks struct {
	Mandate string `url:"mandate,omitempty" json:"mandate,omitempty"`
	Payment string `url:"payment,omitempty" json:"payment,omitempty"`
	Refund  string `url:"refund,omitempty" json:"refund,omitempty"`
}

type PayoutItemTaxes struct {
	Amount              string `url:"amount,omitempty" json:"amount,omitempty"`
	Currency            string `url:"currency,omitempty" json:"currency,omitempty"`
	DestinationAmount   string `url:"destination_amount,omitempty" json:"destination_amount,omitempty"`
	DestinationCurrency string `url:"destination_currency,omitempty" json:"destination_currency,omitempty"`
	ExchangeRate        string `url:"exchange_rate,omitempty" json:"exchange_rate,omitempty"`
	TaxRateId           string `url:"tax_rate_id,omitempty" json:"tax_rate_id,omitempty"`
}

// PayoutItem model
type PayoutItem struct {
	Amount string            `url:"amount,omitempty" json:"amount,omitempty"`
	Links  *PayoutItemLinks  `url:"links,omitempty" json:"links,omitempty"`
	Taxes  []PayoutItemTaxes `url:"taxes,omitempty" json:"taxes,omitempty"`
	Type   string            `url:"type,omitempty" json:"type,omitempty"`
}

type PayoutItemService interface {
	List(ctx context.Context, p PayoutItemListParams, opts ...RequestOption) (*PayoutItemListResult, error)
	All(ctx context.Context, p PayoutItemListParams, opts ...RequestOption) *PayoutItemListPagingIterator
}

// PayoutItemListParams parameters
type PayoutItemListParams struct {
	After                 string `url:"after,omitempty" json:"after,omitempty"`
	Before                string `url:"before,omitempty" json:"before,omitempty"`
	Include2020TaxCutover string `url:"include_2020_tax_cutover,omitempty" json:"include_2020_tax_cutover,omitempty"`
	Limit                 int    `url:"limit,omitempty" json:"limit,omitempty"`
	Payout                string `url:"payout,omitempty" json:"payout,omitempty"`
}

type PayoutItemListResultMetaCursors struct {
	After  string `url:"after,omitempty" json:"after,omitempty"`
	Before string `url:"before,omitempty" json:"before,omitempty"`
}

type PayoutItemListResultMeta struct {
	Cursors *PayoutItemListResultMetaCursors `url:"cursors,omitempty" json:"cursors,omitempty"`
	Limit   int                              `url:"limit,omitempty" json:"limit,omitempty"`
}

type PayoutItemListResult struct {
	PayoutItems []PayoutItem             `json:"payout_items"`
	Meta        PayoutItemListResultMeta `url:"meta,omitempty" json:"meta,omitempty"`
}

// List
// Returns a [cursor-paginated](#api-usage-cursor-pagination) list of items in
// the payout.
//
// <strong>This endpoint only serves requests for payouts created in the last 6
// months. Requests for older payouts will return an HTTP status <code>410
// Gone</code>.</strong>
func (s *PayoutItemServiceImpl) List(ctx context.Context, p PayoutItemListParams, opts ...RequestOption) (*PayoutItemListResult, error) {
	uri, err := url.Parse(fmt.Sprintf(s.config.Endpoint() + "/payout_items"))
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
	req.Header.Set("GoCardless-Client-Version", "3.5.0")
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
		*PayoutItemListResult
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

	if result.PayoutItemListResult == nil {
		return nil, errors.New("missing result")
	}

	return result.PayoutItemListResult, nil
}

type PayoutItemListPagingIterator struct {
	cursor         string
	response       *PayoutItemListResult
	params         PayoutItemListParams
	service        *PayoutItemServiceImpl
	requestOptions []RequestOption
}

func (c *PayoutItemListPagingIterator) Next() bool {
	if c.cursor == "" && c.response != nil {
		return false
	}

	return true
}

func (c *PayoutItemListPagingIterator) Value(ctx context.Context) (*PayoutItemListResult, error) {
	if !c.Next() {
		return c.response, nil
	}

	s := c.service
	p := c.params
	p.After = c.cursor

	uri, err := url.Parse(fmt.Sprintf(s.config.Endpoint() + "/payout_items"))

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
	req.Header.Set("GoCardless-Client-Version", "3.5.0")
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
		*PayoutItemListResult
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

	if result.PayoutItemListResult == nil {
		return nil, errors.New("missing result")
	}

	c.response = result.PayoutItemListResult
	c.cursor = c.response.Meta.Cursors.After
	return c.response, nil
}

func (s *PayoutItemServiceImpl) All(ctx context.Context,
	p PayoutItemListParams,
	opts ...RequestOption) *PayoutItemListPagingIterator {
	return &PayoutItemListPagingIterator{
		params:         p,
		service:        s,
		requestOptions: opts,
	}
}
