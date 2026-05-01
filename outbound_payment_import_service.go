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

// OutboundPaymentImportService manages outbound_payment_imports
type OutboundPaymentImportServiceImpl struct {
	config Config
}

type OutboundPaymentImportEntryCounts struct {
	FailedToProcess           int `url:"failed_to_process,omitempty" json:"failed_to_process,omitempty"`
	Invalid                   int `url:"invalid,omitempty" json:"invalid,omitempty"`
	Processed                 int `url:"processed,omitempty" json:"processed,omitempty"`
	Total                     int `url:"total,omitempty" json:"total,omitempty"`
	Valid                     int `url:"valid,omitempty" json:"valid,omitempty"`
	Verified                  int `url:"verified,omitempty" json:"verified,omitempty"`
	VerifiedWithFullMatch     int `url:"verified_with_full_match,omitempty" json:"verified_with_full_match,omitempty"`
	VerifiedWithNoMatch       int `url:"verified_with_no_match,omitempty" json:"verified_with_no_match,omitempty"`
	VerifiedWithPartialMatch  int `url:"verified_with_partial_match,omitempty" json:"verified_with_partial_match,omitempty"`
	VerifiedWithUnableToMatch int `url:"verified_with_unable_to_match,omitempty" json:"verified_with_unable_to_match,omitempty"`
}

type OutboundPaymentImportLinks struct {
	Creditor string `url:"creditor,omitempty" json:"creditor,omitempty"`
}

// OutboundPaymentImport model
type OutboundPaymentImport struct {
	AmountSum        int                               `url:"amount_sum,omitempty" json:"amount_sum,omitempty"`
	AuthorisationUrl string                            `url:"authorisation_url,omitempty" json:"authorisation_url,omitempty"`
	CreatedAt        string                            `url:"created_at,omitempty" json:"created_at,omitempty"`
	Currency         string                            `url:"currency,omitempty" json:"currency,omitempty"`
	EntryCounts      *OutboundPaymentImportEntryCounts `url:"entry_counts,omitempty" json:"entry_counts,omitempty"`
	Id               string                            `url:"id,omitempty" json:"id,omitempty"`
	Links            *OutboundPaymentImportLinks       `url:"links,omitempty" json:"links,omitempty"`
	Status           string                            `url:"status,omitempty" json:"status,omitempty"`
}

type OutboundPaymentImportService interface {
	Create(ctx context.Context, p OutboundPaymentImportCreateParams, opts ...RequestOption) (*OutboundPaymentImport, error)
	Get(ctx context.Context, identity string, p OutboundPaymentImportGetParams, opts ...RequestOption) (*OutboundPaymentImport, error)
	List(ctx context.Context, p OutboundPaymentImportListParams, opts ...RequestOption) (*OutboundPaymentImportListResult, error)
	All(ctx context.Context,
		p OutboundPaymentImportListParams, opts ...RequestOption) *OutboundPaymentImportListPagingIterator
}

type OutboundPaymentImportCreateParamsEntryItems struct {
	Amount                 int               `url:"amount,omitempty" json:"amount,omitempty"`
	Metadata               map[string]string `url:"metadata,omitempty" json:"metadata,omitempty"`
	RecipientBankAccountId string            `url:"recipient_bank_account_id,omitempty" json:"recipient_bank_account_id,omitempty"`
	Reference              string            `url:"reference,omitempty" json:"reference,omitempty"`
	Scheme                 string            `url:"scheme,omitempty" json:"scheme,omitempty"`
}

type OutboundPaymentImportCreateParamsLinks struct {
	Creditor string `url:"creditor,omitempty" json:"creditor,omitempty"`
}

// OutboundPaymentImportCreateParams parameters
type OutboundPaymentImportCreateParams struct {
	EntryItems []OutboundPaymentImportCreateParamsEntryItems `url:"entry_items,omitempty" json:"entry_items,omitempty"`
	Links      *OutboundPaymentImportCreateParamsLinks       `url:"links,omitempty" json:"links,omitempty"`
}

// Create
func (s *OutboundPaymentImportServiceImpl) Create(ctx context.Context, p OutboundPaymentImportCreateParams, opts ...RequestOption) (*OutboundPaymentImport, error) {
	uri, err := url.Parse(fmt.Sprintf(s.config.Endpoint() + "/outbound_payment_imports"))
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
		"outbound_payment_imports": p,
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
	req.Header.Set("GoCardless-Client-Version", ClientLibVersion)
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
		Err                   *APIError              `json:"error"`
		OutboundPaymentImport *OutboundPaymentImport `json:"outbound_payment_imports"`
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

	if result.OutboundPaymentImport == nil {
		return nil, errors.New("missing result")
	}

	return result.OutboundPaymentImport, nil
}

// OutboundPaymentImportGetParams parameters
type OutboundPaymentImportGetParams struct {
}

// Get
// Returns a single outbound payment import.
func (s *OutboundPaymentImportServiceImpl) Get(ctx context.Context, identity string, p OutboundPaymentImportGetParams, opts ...RequestOption) (*OutboundPaymentImport, error) {
	uri, err := url.Parse(fmt.Sprintf(s.config.Endpoint()+"/outbound_payment_imports/%v",
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
		Err                   *APIError              `json:"error"`
		OutboundPaymentImport *OutboundPaymentImport `json:"outbound_payment_imports"`
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

	if result.OutboundPaymentImport == nil {
		return nil, errors.New("missing result")
	}

	return result.OutboundPaymentImport, nil
}

type OutboundPaymentImportListParamsCreatedAt struct {
	Gt  string `url:"gt,omitempty" json:"gt,omitempty"`
	Gte string `url:"gte,omitempty" json:"gte,omitempty"`
	Lt  string `url:"lt,omitempty" json:"lt,omitempty"`
	Lte string `url:"lte,omitempty" json:"lte,omitempty"`
}

// OutboundPaymentImportListParams parameters
type OutboundPaymentImportListParams struct {
	After     string                                    `url:"after,omitempty" json:"after,omitempty"`
	Before    string                                    `url:"before,omitempty" json:"before,omitempty"`
	CreatedAt *OutboundPaymentImportListParamsCreatedAt `url:"created_at,omitempty" json:"created_at,omitempty"`
	Limit     int                                       `url:"limit,omitempty" json:"limit,omitempty"`
	Status    string                                    `url:"status,omitempty" json:"status,omitempty"`
}

type OutboundPaymentImportListResultMetaCursors struct {
	After  string `url:"after,omitempty" json:"after,omitempty"`
	Before string `url:"before,omitempty" json:"before,omitempty"`
}

type OutboundPaymentImportListResultMeta struct {
	Cursors *OutboundPaymentImportListResultMetaCursors `url:"cursors,omitempty" json:"cursors,omitempty"`
	Limit   int                                         `url:"limit,omitempty" json:"limit,omitempty"`
}

type OutboundPaymentImportListResult struct {
	OutboundPaymentImports []OutboundPaymentImport             `json:"outbound_payment_imports"`
	Meta                   OutboundPaymentImportListResultMeta `url:"meta,omitempty" json:"meta,omitempty"`
}

// List
// Returns a [cursor-paginated](#api-usage-cursor-pagination) list of your
// outbound payment imports.
func (s *OutboundPaymentImportServiceImpl) List(ctx context.Context, p OutboundPaymentImportListParams, opts ...RequestOption) (*OutboundPaymentImportListResult, error) {
	uri, err := url.Parse(fmt.Sprintf(s.config.Endpoint() + "/outbound_payment_imports"))
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
		*OutboundPaymentImportListResult
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

	if result.OutboundPaymentImportListResult == nil {
		return nil, errors.New("missing result")
	}

	return result.OutboundPaymentImportListResult, nil
}

type OutboundPaymentImportListPagingIterator struct {
	cursor         string
	response       *OutboundPaymentImportListResult
	params         OutboundPaymentImportListParams
	service        *OutboundPaymentImportServiceImpl
	requestOptions []RequestOption
}

func (c *OutboundPaymentImportListPagingIterator) Next() bool {
	if c.cursor == "" && c.response != nil {
		return false
	}

	return true
}

func (c *OutboundPaymentImportListPagingIterator) Value(ctx context.Context) (*OutboundPaymentImportListResult, error) {
	if !c.Next() {
		return c.response, nil
	}

	s := c.service
	p := c.params
	p.After = c.cursor

	uri, err := url.Parse(fmt.Sprintf(s.config.Endpoint() + "/outbound_payment_imports"))

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
		*OutboundPaymentImportListResult
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

	if result.OutboundPaymentImportListResult == nil {
		return nil, errors.New("missing result")
	}

	c.response = result.OutboundPaymentImportListResult
	c.cursor = c.response.Meta.Cursors.After
	return c.response, nil
}

func (s *OutboundPaymentImportServiceImpl) All(ctx context.Context,
	p OutboundPaymentImportListParams,
	opts ...RequestOption) *OutboundPaymentImportListPagingIterator {
	return &OutboundPaymentImportListPagingIterator{
		params:         p,
		service:        s,
		requestOptions: opts,
	}
}
