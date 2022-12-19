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

// CreditorBankAccountService manages creditor_bank_accounts
type CreditorBankAccountServiceImpl struct {
	config Config
}

type CreditorBankAccountLinks struct {
	Creditor string `url:"creditor,omitempty" json:"creditor,omitempty"`
}

// CreditorBankAccount model
type CreditorBankAccount struct {
	AccountHolderName   string                    `url:"account_holder_name,omitempty" json:"account_holder_name,omitempty"`
	AccountNumberEnding string                    `url:"account_number_ending,omitempty" json:"account_number_ending,omitempty"`
	AccountType         string                    `url:"account_type,omitempty" json:"account_type,omitempty"`
	BankName            string                    `url:"bank_name,omitempty" json:"bank_name,omitempty"`
	CountryCode         string                    `url:"country_code,omitempty" json:"country_code,omitempty"`
	CreatedAt           string                    `url:"created_at,omitempty" json:"created_at,omitempty"`
	Currency            string                    `url:"currency,omitempty" json:"currency,omitempty"`
	Enabled             bool                      `url:"enabled,omitempty" json:"enabled,omitempty"`
	Id                  string                    `url:"id,omitempty" json:"id,omitempty"`
	Links               *CreditorBankAccountLinks `url:"links,omitempty" json:"links,omitempty"`
	Metadata            map[string]interface{}    `url:"metadata,omitempty" json:"metadata,omitempty"`
}

type CreditorBankAccountService interface {
	Create(ctx context.Context, p CreditorBankAccountCreateParams, opts ...RequestOption) (*CreditorBankAccount, error)
	List(ctx context.Context, p CreditorBankAccountListParams, opts ...RequestOption) (*CreditorBankAccountListResult, error)
	All(ctx context.Context, p CreditorBankAccountListParams, opts ...RequestOption) *CreditorBankAccountListPagingIterator
	Get(ctx context.Context, identity string, opts ...RequestOption) (*CreditorBankAccount, error)
	Disable(ctx context.Context, identity string, opts ...RequestOption) (*CreditorBankAccount, error)
}

type CreditorBankAccountCreateParamsLinks struct {
	Creditor string `url:"creditor,omitempty" json:"creditor,omitempty"`
}

// CreditorBankAccountCreateParams parameters
type CreditorBankAccountCreateParams struct {
	AccountHolderName         string                               `url:"account_holder_name,omitempty" json:"account_holder_name,omitempty"`
	AccountNumber             string                               `url:"account_number,omitempty" json:"account_number,omitempty"`
	AccountType               string                               `url:"account_type,omitempty" json:"account_type,omitempty"`
	BankCode                  string                               `url:"bank_code,omitempty" json:"bank_code,omitempty"`
	BranchCode                string                               `url:"branch_code,omitempty" json:"branch_code,omitempty"`
	CountryCode               string                               `url:"country_code,omitempty" json:"country_code,omitempty"`
	Currency                  string                               `url:"currency,omitempty" json:"currency,omitempty"`
	Iban                      string                               `url:"iban,omitempty" json:"iban,omitempty"`
	Links                     CreditorBankAccountCreateParamsLinks `url:"links,omitempty" json:"links,omitempty"`
	Metadata                  map[string]interface{}               `url:"metadata,omitempty" json:"metadata,omitempty"`
	SetAsDefaultPayoutAccount bool                                 `url:"set_as_default_payout_account,omitempty" json:"set_as_default_payout_account,omitempty"`
}

// Create
// Creates a new creditor bank account object.
func (s *CreditorBankAccountServiceImpl) Create(ctx context.Context, p CreditorBankAccountCreateParams, opts ...RequestOption) (*CreditorBankAccount, error) {
	uri, err := url.Parse(fmt.Sprintf(s.config.Endpoint() + "/creditor_bank_accounts"))
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
		"creditor_bank_accounts": p,
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
	req.Header.Set("GoCardless-Client-Version", "2.8.0")
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
		Err                 *APIError            `json:"error"`
		CreditorBankAccount *CreditorBankAccount `json:"creditor_bank_accounts"`
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

	if result.CreditorBankAccount == nil {
		return nil, errors.New("missing result")
	}

	return result.CreditorBankAccount, nil
}

type CreditorBankAccountListParamsCreatedAt struct {
	Gt  string `url:"gt,omitempty" json:"gt,omitempty"`
	Gte string `url:"gte,omitempty" json:"gte,omitempty"`
	Lt  string `url:"lt,omitempty" json:"lt,omitempty"`
	Lte string `url:"lte,omitempty" json:"lte,omitempty"`
}

// CreditorBankAccountListParams parameters
type CreditorBankAccountListParams struct {
	After     string                                  `url:"after,omitempty" json:"after,omitempty"`
	Before    string                                  `url:"before,omitempty" json:"before,omitempty"`
	CreatedAt *CreditorBankAccountListParamsCreatedAt `url:"created_at,omitempty" json:"created_at,omitempty"`
	Creditor  string                                  `url:"creditor,omitempty" json:"creditor,omitempty"`
	Enabled   bool                                    `url:"enabled,omitempty" json:"enabled,omitempty"`
	Limit     int                                     `url:"limit,omitempty" json:"limit,omitempty"`
}

type CreditorBankAccountListResultMetaCursors struct {
	After  string `url:"after,omitempty" json:"after,omitempty"`
	Before string `url:"before,omitempty" json:"before,omitempty"`
}

type CreditorBankAccountListResultMeta struct {
	Cursors *CreditorBankAccountListResultMetaCursors `url:"cursors,omitempty" json:"cursors,omitempty"`
	Limit   int                                       `url:"limit,omitempty" json:"limit,omitempty"`
}

type CreditorBankAccountListResult struct {
	CreditorBankAccounts []CreditorBankAccount             `json:"creditor_bank_accounts"`
	Meta                 CreditorBankAccountListResultMeta `url:"meta,omitempty" json:"meta,omitempty"`
}

// List
// Returns a [cursor-paginated](#api-usage-cursor-pagination) list of your
// creditor bank accounts.
func (s *CreditorBankAccountServiceImpl) List(ctx context.Context, p CreditorBankAccountListParams, opts ...RequestOption) (*CreditorBankAccountListResult, error) {
	uri, err := url.Parse(fmt.Sprintf(s.config.Endpoint() + "/creditor_bank_accounts"))
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
	req.Header.Set("GoCardless-Client-Version", "2.8.0")
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
		*CreditorBankAccountListResult
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

	if result.CreditorBankAccountListResult == nil {
		return nil, errors.New("missing result")
	}

	return result.CreditorBankAccountListResult, nil
}

type CreditorBankAccountListPagingIterator struct {
	cursor         string
	response       *CreditorBankAccountListResult
	params         CreditorBankAccountListParams
	service        *CreditorBankAccountServiceImpl
	requestOptions []RequestOption
}

func (c *CreditorBankAccountListPagingIterator) Next() bool {
	if c.cursor == "" && c.response != nil {
		return false
	}

	return true
}

func (c *CreditorBankAccountListPagingIterator) Value(ctx context.Context) (*CreditorBankAccountListResult, error) {
	if !c.Next() {
		return c.response, nil
	}

	s := c.service
	p := c.params
	p.After = c.cursor

	uri, err := url.Parse(fmt.Sprintf(s.config.Endpoint() + "/creditor_bank_accounts"))

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
	req.Header.Set("GoCardless-Client-Version", "2.8.0")
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
		*CreditorBankAccountListResult
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

	if result.CreditorBankAccountListResult == nil {
		return nil, errors.New("missing result")
	}

	c.response = result.CreditorBankAccountListResult
	c.cursor = c.response.Meta.Cursors.After
	return c.response, nil
}

func (s *CreditorBankAccountServiceImpl) All(ctx context.Context,
	p CreditorBankAccountListParams,
	opts ...RequestOption) *CreditorBankAccountListPagingIterator {
	return &CreditorBankAccountListPagingIterator{
		params:         p,
		service:        s,
		requestOptions: opts,
	}
}

// Get
// Retrieves the details of an existing creditor bank account.
func (s *CreditorBankAccountServiceImpl) Get(ctx context.Context, identity string, opts ...RequestOption) (*CreditorBankAccount, error) {
	uri, err := url.Parse(fmt.Sprintf(s.config.Endpoint()+"/creditor_bank_accounts/%v",
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
	req.Header.Set("GoCardless-Client-Version", "2.8.0")
	req.Header.Set("User-Agent", userAgent)

	for key, value := range o.headers {
		req.Header.Set(key, value)
	}

	client := s.config.Client()
	if client == nil {
		client = http.DefaultClient
	}

	var result struct {
		Err                 *APIError            `json:"error"`
		CreditorBankAccount *CreditorBankAccount `json:"creditor_bank_accounts"`
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

	if result.CreditorBankAccount == nil {
		return nil, errors.New("missing result")
	}

	return result.CreditorBankAccount, nil
}

// Disable
// Immediately disables the bank account, no money can be paid out to a disabled
// account.
//
// This will return a `disable_failed` error if the bank account has already
// been disabled.
//
// A disabled bank account can be re-enabled by creating a new bank account
// resource with the same details.
func (s *CreditorBankAccountServiceImpl) Disable(ctx context.Context, identity string, opts ...RequestOption) (*CreditorBankAccount, error) {
	uri, err := url.Parse(fmt.Sprintf(s.config.Endpoint()+"/creditor_bank_accounts/%v/actions/disable",
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

	req, err := http.NewRequest("POST", uri.String(), body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	req.Header.Set("Authorization", "Bearer "+s.config.Token())
	req.Header.Set("GoCardless-Version", "2015-07-06")
	req.Header.Set("GoCardless-Client-Library", "gocardless-pro-go")
	req.Header.Set("GoCardless-Client-Version", "2.8.0")
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
		Err                 *APIError            `json:"error"`
		CreditorBankAccount *CreditorBankAccount `json:"creditor_bank_accounts"`
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

	if result.CreditorBankAccount == nil {
		return nil, errors.New("missing result")
	}

	return result.CreditorBankAccount, nil
}
