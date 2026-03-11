package gocardless

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/google/go-querystring/query"
)

// PaymentAccountService manages payment_accounts
type PaymentAccountServiceImpl struct {
	config Config
}

type PaymentAccountLinks struct {
	Creditor string `url:"creditor,omitempty" json:"creditor,omitempty"`
}

// PaymentAccount model
type PaymentAccount struct {
	AccountBalance      int                  `url:"account_balance,omitempty" json:"account_balance,omitempty"`
	AccountHolderName   string               `url:"account_holder_name,omitempty" json:"account_holder_name,omitempty"`
	AccountNumberEnding string               `url:"account_number_ending,omitempty" json:"account_number_ending,omitempty"`
	BankName            string               `url:"bank_name,omitempty" json:"bank_name,omitempty"`
	Currency            string               `url:"currency,omitempty" json:"currency,omitempty"`
	Id                  string               `url:"id,omitempty" json:"id,omitempty"`
	Links               *PaymentAccountLinks `url:"links,omitempty" json:"links,omitempty"`
}

type PaymentAccountService interface {
	Get(ctx context.Context, identity string, opts ...RequestOption) (*PaymentAccount, error)
	List(ctx context.Context, p PaymentAccountListParams, opts ...RequestOption) (*PaymentAccountListResult, error)
	All(ctx context.Context,
		p PaymentAccountListParams, opts ...RequestOption) *PaymentAccountListPagingIterator
}

// Get
// Retrieves the details of an existing payment account.
func (s *PaymentAccountServiceImpl) Get(ctx context.Context, identity string, opts ...RequestOption) (*PaymentAccount, error) {
	uri, err := url.Parse(fmt.Sprintf(s.config.Endpoint()+"/payment_accounts/%v",
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
	req.Header.Set("GoCardless-Client-Version", ClientLibVersion)
	req.Header.Set("User-Agent", userAgent)

	for key, value := range o.headers {
		req.Header.Set(key, value)
	}

	client := s.config.Client()
	if client == nil {
		client = http.DefaultClient
	}

	var result struct {
		Err            *APIError       `json:"error"`
		PaymentAccount *PaymentAccount `json:"payment_accounts"`
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

	if result.PaymentAccount == nil {
		return nil, errors.New("missing result")
	}

	return result.PaymentAccount, nil
}

// PaymentAccountListParams parameters
type PaymentAccountListParams struct {
	After  string `url:"after,omitempty" json:"after,omitempty"`
	Before string `url:"before,omitempty" json:"before,omitempty"`
	Limit  int    `url:"limit,omitempty" json:"limit,omitempty"`
}

type PaymentAccountListResultMetaCursors struct {
	After  string `url:"after,omitempty" json:"after,omitempty"`
	Before string `url:"before,omitempty" json:"before,omitempty"`
}

type PaymentAccountListResultMeta struct {
	Cursors *PaymentAccountListResultMetaCursors `url:"cursors,omitempty" json:"cursors,omitempty"`
	Limit   int                                  `url:"limit,omitempty" json:"limit,omitempty"`
}

type PaymentAccountListResult struct {
	PaymentAccounts []PaymentAccount             `json:"payment_accounts"`
	Meta            PaymentAccountListResultMeta `url:"meta,omitempty" json:"meta,omitempty"`
}

// List
// Returns a [cursor-paginated](#api-usage-cursor-pagination) list of your
// payment accounts.
func (s *PaymentAccountServiceImpl) List(ctx context.Context, p PaymentAccountListParams, opts ...RequestOption) (*PaymentAccountListResult, error) {
	uri, err := url.Parse(fmt.Sprintf(s.config.Endpoint() + "/payment_accounts"))
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
	req.Header.Set("GoCardless-Client-Version", ClientLibVersion)
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
		*PaymentAccountListResult
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

	if result.PaymentAccountListResult == nil {
		return nil, errors.New("missing result")
	}

	return result.PaymentAccountListResult, nil
}

type PaymentAccountListPagingIterator struct {
	cursor         string
	response       *PaymentAccountListResult
	params         PaymentAccountListParams
	service        *PaymentAccountServiceImpl
	requestOptions []RequestOption
}

func (c *PaymentAccountListPagingIterator) Next() bool {
	if c.cursor == "" && c.response != nil {
		return false
	}

	return true
}

func (c *PaymentAccountListPagingIterator) Value(ctx context.Context) (*PaymentAccountListResult, error) {
	if !c.Next() {
		return c.response, nil
	}

	s := c.service
	p := c.params
	p.After = c.cursor

	uri, err := url.Parse(fmt.Sprintf(s.config.Endpoint() + "/payment_accounts"))

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
	req.Header.Set("GoCardless-Client-Version", ClientLibVersion)
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
		*PaymentAccountListResult
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

	if result.PaymentAccountListResult == nil {
		return nil, errors.New("missing result")
	}

	c.response = result.PaymentAccountListResult
	c.cursor = c.response.Meta.Cursors.After
	return c.response, nil
}

func (s *PaymentAccountServiceImpl) All(ctx context.Context,
	p PaymentAccountListParams,
	opts ...RequestOption) *PaymentAccountListPagingIterator {
	return &PaymentAccountListPagingIterator{
		params:         p,
		service:        s,
		requestOptions: opts,
	}
}
