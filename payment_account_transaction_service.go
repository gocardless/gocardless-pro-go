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

// PaymentAccountTransactionService manages payment_account_transactions
type PaymentAccountTransactionServiceImpl struct {
	config Config
}

type PaymentAccountTransactionLinks struct {
	OutboundPayment    string `url:"outbound_payment,omitempty" json:"outbound_payment,omitempty"`
	PaymentBankAccount string `url:"payment_bank_account,omitempty" json:"payment_bank_account,omitempty"`
	Payout             string `url:"payout,omitempty" json:"payout,omitempty"`
}

// PaymentAccountTransaction model
type PaymentAccountTransaction struct {
	Amount                  int                             `url:"amount,omitempty" json:"amount,omitempty"`
	BalanceAfterTransaction int                             `url:"balance_after_transaction,omitempty" json:"balance_after_transaction,omitempty"`
	CounterpartyName        string                          `url:"counterparty_name,omitempty" json:"counterparty_name,omitempty"`
	Currency                string                          `url:"currency,omitempty" json:"currency,omitempty"`
	Description             string                          `url:"description,omitempty" json:"description,omitempty"`
	Direction               string                          `url:"direction,omitempty" json:"direction,omitempty"`
	Id                      string                          `url:"id,omitempty" json:"id,omitempty"`
	Links                   *PaymentAccountTransactionLinks `url:"links,omitempty" json:"links,omitempty"`
	Reference               string                          `url:"reference,omitempty" json:"reference,omitempty"`
	ValueDate               string                          `url:"value_date,omitempty" json:"value_date,omitempty"`
}

type PaymentAccountTransactionService interface {
	Get(ctx context.Context, identity string, opts ...RequestOption) (*PaymentAccountTransaction, error)
	List(ctx context.Context, identity string, p PaymentAccountTransactionListParams, opts ...RequestOption) (*PaymentAccountTransactionListResult, error)
	All(ctx context.Context,
		identity string,
		p PaymentAccountTransactionListParams, opts ...RequestOption) *PaymentAccountTransactionListPagingIterator
}

// Get
// Retrieves the details of an existing payment account transaction.
func (s *PaymentAccountTransactionServiceImpl) Get(ctx context.Context, identity string, opts ...RequestOption) (*PaymentAccountTransaction, error) {
	uri, err := url.Parse(fmt.Sprintf(s.config.Endpoint()+"/payment_account_transactions/%v",
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
		Err                       *APIError                  `json:"error"`
		PaymentAccountTransaction *PaymentAccountTransaction `json:"payment_account_transactions"`
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

	if result.PaymentAccountTransaction == nil {
		return nil, errors.New("missing result")
	}

	return result.PaymentAccountTransaction, nil
}

// PaymentAccountTransactionListParams parameters
type PaymentAccountTransactionListParams struct {
	After         string `url:"after,omitempty" json:"after,omitempty"`
	Before        string `url:"before,omitempty" json:"before,omitempty"`
	Direction     string `url:"direction,omitempty" json:"direction,omitempty"`
	Limit         int    `url:"limit,omitempty" json:"limit,omitempty"`
	ValueDateFrom string `url:"value_date_from,omitempty" json:"value_date_from,omitempty"`
	ValueDateTo   string `url:"value_date_to,omitempty" json:"value_date_to,omitempty"`
}

type PaymentAccountTransactionListResultMetaCursors struct {
	After  string `url:"after,omitempty" json:"after,omitempty"`
	Before string `url:"before,omitempty" json:"before,omitempty"`
}

type PaymentAccountTransactionListResultMeta struct {
	Cursors *PaymentAccountTransactionListResultMetaCursors `url:"cursors,omitempty" json:"cursors,omitempty"`
	Limit   int                                             `url:"limit,omitempty" json:"limit,omitempty"`
}

type PaymentAccountTransactionListResult struct {
	PaymentAccountTransactions []PaymentAccountTransaction             `json:"payment_account_transactions"`
	Meta                       PaymentAccountTransactionListResultMeta `url:"meta,omitempty" json:"meta,omitempty"`
}

// List
// List transactions for a given payment account.
func (s *PaymentAccountTransactionServiceImpl) List(ctx context.Context, identity string, p PaymentAccountTransactionListParams, opts ...RequestOption) (*PaymentAccountTransactionListResult, error) {
	uri, err := url.Parse(fmt.Sprintf(s.config.Endpoint()+"/payment_accounts/%v/transactions",
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
		*PaymentAccountTransactionListResult
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

	if result.PaymentAccountTransactionListResult == nil {
		return nil, errors.New("missing result")
	}

	return result.PaymentAccountTransactionListResult, nil
}

type PaymentAccountTransactionListPagingIterator struct {
	cursor         string
	response       *PaymentAccountTransactionListResult
	params         PaymentAccountTransactionListParams
	service        *PaymentAccountTransactionServiceImpl
	requestOptions []RequestOption
	identity       string
}

func (c *PaymentAccountTransactionListPagingIterator) Next() bool {
	if c.cursor == "" && c.response != nil {
		return false
	}

	return true
}

func (c *PaymentAccountTransactionListPagingIterator) Value(ctx context.Context) (*PaymentAccountTransactionListResult, error) {
	if !c.Next() {
		return c.response, nil
	}

	s := c.service
	p := c.params
	p.After = c.cursor

	uri, err := url.Parse(fmt.Sprintf(s.config.Endpoint()+"/payment_accounts/%v/transactions",
		c.identity))

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
		*PaymentAccountTransactionListResult
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

	if result.PaymentAccountTransactionListResult == nil {
		return nil, errors.New("missing result")
	}

	c.response = result.PaymentAccountTransactionListResult
	c.cursor = c.response.Meta.Cursors.After
	return c.response, nil
}

func (s *PaymentAccountTransactionServiceImpl) All(ctx context.Context,
	identity string,
	p PaymentAccountTransactionListParams,
	opts ...RequestOption) *PaymentAccountTransactionListPagingIterator {
	return &PaymentAccountTransactionListPagingIterator{
		params:         p,
		service:        s,
		requestOptions: opts,
		identity:       identity,
	}
}
