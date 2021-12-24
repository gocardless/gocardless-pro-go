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
type PayoutItemService struct {
	endpoint string
	token    string
	client   *http.Client
}

// PayoutItem model
type PayoutItem struct {
	Amount string `url:"amount,omitempty" json:"amount,omitempty"`
	Links  struct {
		Mandate string `url:"mandate,omitempty" json:"mandate,omitempty"`
		Payment string `url:"payment,omitempty" json:"payment,omitempty"`
		Refund  string `url:"refund,omitempty" json:"refund,omitempty"`
	} `url:"links,omitempty" json:"links,omitempty"`
	Taxes []struct {
		Amount              string `url:"amount,omitempty" json:"amount,omitempty"`
		Currency            string `url:"currency,omitempty" json:"currency,omitempty"`
		DestinationAmount   string `url:"destination_amount,omitempty" json:"destination_amount,omitempty"`
		DestinationCurrency string `url:"destination_currency,omitempty" json:"destination_currency,omitempty"`
		ExchangeRate        string `url:"exchange_rate,omitempty" json:"exchange_rate,omitempty"`
		TaxRateId           string `url:"tax_rate_id,omitempty" json:"tax_rate_id,omitempty"`
	} `url:"taxes,omitempty" json:"taxes,omitempty"`
	Type string `url:"type,omitempty" json:"type,omitempty"`
}

// PayoutItemListParams parameters
type PayoutItemListParams struct {
	After                 string `url:"after,omitempty" json:"after,omitempty"`
	Before                string `url:"before,omitempty" json:"before,omitempty"`
	Include2020TaxCutover string `url:"include_2020_tax_cutover,omitempty" json:"include_2020_tax_cutover,omitempty"`
	Limit                 int    `url:"limit,omitempty" json:"limit,omitempty"`
	Payout                string `url:"payout,omitempty" json:"payout,omitempty"`
}

// PayoutItemListResult response including pagination metadata
type PayoutItemListResult struct {
	PayoutItems []PayoutItem `json:"payout_items"`
	Meta        struct {
		Cursors struct {
			After  string `url:"after,omitempty" json:"after,omitempty"`
			Before string `url:"before,omitempty" json:"before,omitempty"`
		} `url:"cursors,omitempty" json:"cursors,omitempty"`
		Limit int `url:"limit,omitempty" json:"limit,omitempty"`
	} `json:"meta"`
}

// List
// Returns a [cursor-paginated](#api-usage-cursor-pagination) list of items in
// the payout.
//
func (s *PayoutItemService) List(ctx context.Context, p PayoutItemListParams, opts ...RequestOption) (*PayoutItemListResult, error) {
	uri, err := url.Parse(fmt.Sprintf(s.endpoint + "/payout_items"))
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
	service        *PayoutItemService
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

	uri, err := url.Parse(fmt.Sprintf(s.endpoint + "/payout_items"))

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

func (s *PayoutItemService) All(ctx context.Context,
	p PayoutItemListParams,
	opts ...RequestOption) *PayoutItemListPagingIterator {
	return &PayoutItemListPagingIterator{
		params:         p,
		service:        s,
		requestOptions: opts,
	}
}
