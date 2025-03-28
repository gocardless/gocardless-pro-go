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

// NegativeBalanceLimitService manages negative_balance_limits
type NegativeBalanceLimitServiceImpl struct {
	config Config
}

type NegativeBalanceLimitLinks struct {
	CreatorUser string `url:"creator_user,omitempty" json:"creator_user,omitempty"`
	Creditor    string `url:"creditor,omitempty" json:"creditor,omitempty"`
}

// NegativeBalanceLimit model
type NegativeBalanceLimit struct {
	BalanceLimit int                        `url:"balance_limit,omitempty" json:"balance_limit,omitempty"`
	CreatedAt    string                     `url:"created_at,omitempty" json:"created_at,omitempty"`
	Currency     string                     `url:"currency,omitempty" json:"currency,omitempty"`
	Id           string                     `url:"id,omitempty" json:"id,omitempty"`
	Links        *NegativeBalanceLimitLinks `url:"links,omitempty" json:"links,omitempty"`
}

type NegativeBalanceLimitService interface {
	List(ctx context.Context, p NegativeBalanceLimitListParams, opts ...RequestOption) (*NegativeBalanceLimitListResult, error)
	All(ctx context.Context, p NegativeBalanceLimitListParams, opts ...RequestOption) *NegativeBalanceLimitListPagingIterator
	Create(ctx context.Context, p NegativeBalanceLimitCreateParams, opts ...RequestOption) (*NegativeBalanceLimit, error)
}

// NegativeBalanceLimitListParams parameters
type NegativeBalanceLimitListParams struct {
	After    string `url:"after,omitempty" json:"after,omitempty"`
	Before   string `url:"before,omitempty" json:"before,omitempty"`
	Creditor string `url:"creditor,omitempty" json:"creditor,omitempty"`
	Currency string `url:"currency,omitempty" json:"currency,omitempty"`
	Limit    int    `url:"limit,omitempty" json:"limit,omitempty"`
}

type NegativeBalanceLimitListResultMetaCursors struct {
	After  string `url:"after,omitempty" json:"after,omitempty"`
	Before string `url:"before,omitempty" json:"before,omitempty"`
}

type NegativeBalanceLimitListResultMeta struct {
	Cursors *NegativeBalanceLimitListResultMetaCursors `url:"cursors,omitempty" json:"cursors,omitempty"`
	Limit   int                                        `url:"limit,omitempty" json:"limit,omitempty"`
}

type NegativeBalanceLimitListResult struct {
	NegativeBalanceLimits []NegativeBalanceLimit             `json:"negative_balance_limits"`
	Meta                  NegativeBalanceLimitListResultMeta `url:"meta,omitempty" json:"meta,omitempty"`
}

// List
// Returns a [cursor-paginated](#api-usage-cursor-pagination) list of negative
// balance limits.
func (s *NegativeBalanceLimitServiceImpl) List(ctx context.Context, p NegativeBalanceLimitListParams, opts ...RequestOption) (*NegativeBalanceLimitListResult, error) {
	uri, err := url.Parse(fmt.Sprintf(s.config.Endpoint() + "/negative_balance_limits"))
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
	req.Header.Set("GoCardless-Client-Version", "4.4.0")
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
		*NegativeBalanceLimitListResult
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

	if result.NegativeBalanceLimitListResult == nil {
		return nil, errors.New("missing result")
	}

	return result.NegativeBalanceLimitListResult, nil
}

type NegativeBalanceLimitListPagingIterator struct {
	cursor         string
	response       *NegativeBalanceLimitListResult
	params         NegativeBalanceLimitListParams
	service        *NegativeBalanceLimitServiceImpl
	requestOptions []RequestOption
}

func (c *NegativeBalanceLimitListPagingIterator) Next() bool {
	if c.cursor == "" && c.response != nil {
		return false
	}

	return true
}

func (c *NegativeBalanceLimitListPagingIterator) Value(ctx context.Context) (*NegativeBalanceLimitListResult, error) {
	if !c.Next() {
		return c.response, nil
	}

	s := c.service
	p := c.params
	p.After = c.cursor

	uri, err := url.Parse(fmt.Sprintf(s.config.Endpoint() + "/negative_balance_limits"))

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
	req.Header.Set("GoCardless-Client-Version", "4.4.0")
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
		*NegativeBalanceLimitListResult
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

	if result.NegativeBalanceLimitListResult == nil {
		return nil, errors.New("missing result")
	}

	c.response = result.NegativeBalanceLimitListResult
	c.cursor = c.response.Meta.Cursors.After
	return c.response, nil
}

func (s *NegativeBalanceLimitServiceImpl) All(ctx context.Context,
	p NegativeBalanceLimitListParams,
	opts ...RequestOption) *NegativeBalanceLimitListPagingIterator {
	return &NegativeBalanceLimitListPagingIterator{
		params:         p,
		service:        s,
		requestOptions: opts,
	}
}

type NegativeBalanceLimitCreateParamsLinks struct {
	Creditor string `url:"creditor,omitempty" json:"creditor,omitempty"`
}

// NegativeBalanceLimitCreateParams parameters
type NegativeBalanceLimitCreateParams struct {
	BalanceLimit int                                    `url:"balance_limit,omitempty" json:"balance_limit,omitempty"`
	Currency     string                                 `url:"currency,omitempty" json:"currency,omitempty"`
	Links        *NegativeBalanceLimitCreateParamsLinks `url:"links,omitempty" json:"links,omitempty"`
}

// Create
// Creates a new negative balance limit, which replaces the existing limit (if
// present) for that currency and creditor combination.
func (s *NegativeBalanceLimitServiceImpl) Create(ctx context.Context, p NegativeBalanceLimitCreateParams, opts ...RequestOption) (*NegativeBalanceLimit, error) {
	uri, err := url.Parse(fmt.Sprintf(s.config.Endpoint() + "/negative_balance_limits"))
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
		"negative_balance_limits": p,
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
	req.Header.Set("GoCardless-Client-Version", "4.4.0")
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
		Err                  *APIError             `json:"error"`
		NegativeBalanceLimit *NegativeBalanceLimit `json:"negative_balance_limits"`
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

	if result.NegativeBalanceLimit == nil {
		return nil, errors.New("missing result")
	}

	return result.NegativeBalanceLimit, nil
}
