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

// BalanceService manages balances
type BalanceServiceImpl struct {
	config Config
}

type BalanceLinks struct {
	Creditor string `url:"creditor,omitempty" json:"creditor,omitempty"`
}

// Balance model
type Balance struct {
	Amount        int           `url:"amount,omitempty" json:"amount,omitempty"`
	BalanceType   string        `url:"balance_type,omitempty" json:"balance_type,omitempty"`
	Currency      string        `url:"currency,omitempty" json:"currency,omitempty"`
	LastUpdatedAt string        `url:"last_updated_at,omitempty" json:"last_updated_at,omitempty"`
	Links         *BalanceLinks `url:"links,omitempty" json:"links,omitempty"`
}

type BalanceService interface {
	List(ctx context.Context, p BalanceListParams, opts ...RequestOption) (*BalanceListResult, error)
	All(ctx context.Context, p BalanceListParams, opts ...RequestOption) *BalanceListPagingIterator
}

// BalanceListParams parameters
type BalanceListParams struct {
	After    string `url:"after,omitempty" json:"after,omitempty"`
	Before   string `url:"before,omitempty" json:"before,omitempty"`
	Creditor string `url:"creditor,omitempty" json:"creditor,omitempty"`
	Limit    int    `url:"limit,omitempty" json:"limit,omitempty"`
}

type BalanceListResultMetaCursors struct {
	After  string `url:"after,omitempty" json:"after,omitempty"`
	Before string `url:"before,omitempty" json:"before,omitempty"`
}

type BalanceListResultMeta struct {
	Cursors *BalanceListResultMetaCursors `url:"cursors,omitempty" json:"cursors,omitempty"`
	Limit   int                           `url:"limit,omitempty" json:"limit,omitempty"`
}

type BalanceListResult struct {
	Balances []Balance             `json:"balances"`
	Meta     BalanceListResultMeta `url:"meta,omitempty" json:"meta,omitempty"`
}

// List
// Returns a [cursor-paginated](#api-usage-cursor-pagination) list of balances
// for a given creditor. This endpoint is rate limited to 60 requests per
// minute.
func (s *BalanceServiceImpl) List(ctx context.Context, p BalanceListParams, opts ...RequestOption) (*BalanceListResult, error) {
	uri, err := url.Parse(fmt.Sprintf(s.config.Endpoint() + "/balances"))
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
	req.Header.Set("GoCardless-Client-Version", "4.3.0")
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
		*BalanceListResult
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

	if result.BalanceListResult == nil {
		return nil, errors.New("missing result")
	}

	return result.BalanceListResult, nil
}

type BalanceListPagingIterator struct {
	cursor         string
	response       *BalanceListResult
	params         BalanceListParams
	service        *BalanceServiceImpl
	requestOptions []RequestOption
}

func (c *BalanceListPagingIterator) Next() bool {
	if c.cursor == "" && c.response != nil {
		return false
	}

	return true
}

func (c *BalanceListPagingIterator) Value(ctx context.Context) (*BalanceListResult, error) {
	if !c.Next() {
		return c.response, nil
	}

	s := c.service
	p := c.params
	p.After = c.cursor

	uri, err := url.Parse(fmt.Sprintf(s.config.Endpoint() + "/balances"))

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
	req.Header.Set("GoCardless-Client-Version", "4.3.0")
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
		*BalanceListResult
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

	if result.BalanceListResult == nil {
		return nil, errors.New("missing result")
	}

	c.response = result.BalanceListResult
	c.cursor = c.response.Meta.Cursors.After
	return c.response, nil
}

func (s *BalanceServiceImpl) All(ctx context.Context,
	p BalanceListParams,
	opts ...RequestOption) *BalanceListPagingIterator {
	return &BalanceListPagingIterator{
		params:         p,
		service:        s,
		requestOptions: opts,
	}
}
