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

// OutboundPaymentImportEntryService manages outbound_payment_import_entries
type OutboundPaymentImportEntryServiceImpl struct {
	config Config
}

type OutboundPaymentImportEntryLinks struct {
	OutboundPayment       string `url:"outbound_payment,omitempty" json:"outbound_payment,omitempty"`
	OutboundPaymentImport string `url:"outbound_payment_import,omitempty" json:"outbound_payment_import,omitempty"`
	RecipientBankAccount  string `url:"recipient_bank_account,omitempty" json:"recipient_bank_account,omitempty"`
}

type OutboundPaymentImportEntryValidationErrorsOutboundPayment struct {
	Amount               []string `url:"amount,omitempty" json:"amount,omitempty"`
	RecipientBankAccount []string `url:"recipient_bank_account,omitempty" json:"recipient_bank_account,omitempty"`
	Reference            []string `url:"reference,omitempty" json:"reference,omitempty"`
	Scheme               string   `url:"scheme,omitempty" json:"scheme,omitempty"`
}

type OutboundPaymentImportEntryValidationErrors struct {
	OutboundPayment *OutboundPaymentImportEntryValidationErrorsOutboundPayment `url:"outbound_payment,omitempty" json:"outbound_payment,omitempty"`
}

// OutboundPaymentImportEntry model
type OutboundPaymentImportEntry struct {
	Amount             int                                         `url:"amount,omitempty" json:"amount,omitempty"`
	CreatedAt          string                                      `url:"created_at,omitempty" json:"created_at,omitempty"`
	Id                 string                                      `url:"id,omitempty" json:"id,omitempty"`
	Links              *OutboundPaymentImportEntryLinks            `url:"links,omitempty" json:"links,omitempty"`
	Metadata           map[string]string                           `url:"metadata,omitempty" json:"metadata,omitempty"`
	ProcessedAt        string                                      `url:"processed_at,omitempty" json:"processed_at,omitempty"`
	Reference          string                                      `url:"reference,omitempty" json:"reference,omitempty"`
	Scheme             string                                      `url:"scheme,omitempty" json:"scheme,omitempty"`
	ValidationErrors   *OutboundPaymentImportEntryValidationErrors `url:"validation_errors,omitempty" json:"validation_errors,omitempty"`
	VerificationResult string                                      `url:"verification_result,omitempty" json:"verification_result,omitempty"`
}

type OutboundPaymentImportEntryService interface {
	List(ctx context.Context, p OutboundPaymentImportEntryListParams, opts ...RequestOption) (*OutboundPaymentImportEntryListResult, error)
	All(ctx context.Context,
		p OutboundPaymentImportEntryListParams, opts ...RequestOption) *OutboundPaymentImportEntryListPagingIterator
}

// OutboundPaymentImportEntryListParams parameters
type OutboundPaymentImportEntryListParams struct {
	After                 string `url:"after,omitempty" json:"after,omitempty"`
	Before                string `url:"before,omitempty" json:"before,omitempty"`
	Limit                 int    `url:"limit,omitempty" json:"limit,omitempty"`
	OutboundPaymentImport string `url:"outbound_payment_import,omitempty" json:"outbound_payment_import,omitempty"`
}

type OutboundPaymentImportEntryListResultMetaCursors struct {
	After  string `url:"after,omitempty" json:"after,omitempty"`
	Before string `url:"before,omitempty" json:"before,omitempty"`
}

type OutboundPaymentImportEntryListResultMeta struct {
	Cursors OutboundPaymentImportEntryListResultMetaCursors `url:"cursors,omitempty" json:"cursors,omitempty"`
	Limit   int                                             `url:"limit,omitempty" json:"limit,omitempty"`
}

type OutboundPaymentImportEntryListResult struct {
	OutboundPaymentImportEntries []OutboundPaymentImportEntry             `json:"outbound_payment_import_entries"`
	Meta                         OutboundPaymentImportEntryListResultMeta `url:"meta,omitempty" json:"meta,omitempty"`
}

// List
// Returns a [cursor-paginated](#api-usage-cursor-pagination) list of the
// entries for a given outbound payment import.
func (s *OutboundPaymentImportEntryServiceImpl) List(ctx context.Context, p OutboundPaymentImportEntryListParams, opts ...RequestOption) (*OutboundPaymentImportEntryListResult, error) {
	uri, err := url.Parse(fmt.Sprintf(s.config.Endpoint() + "/outbound_payment_import_entries"))
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
		*OutboundPaymentImportEntryListResult
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

	if result.OutboundPaymentImportEntryListResult == nil {
		return nil, errors.New("missing result")
	}

	return result.OutboundPaymentImportEntryListResult, nil
}

type OutboundPaymentImportEntryListPagingIterator struct {
	cursor         string
	response       *OutboundPaymentImportEntryListResult
	params         OutboundPaymentImportEntryListParams
	service        *OutboundPaymentImportEntryServiceImpl
	requestOptions []RequestOption
}

func (c *OutboundPaymentImportEntryListPagingIterator) Next() bool {
	if c.cursor == "" && c.response != nil {
		return false
	}

	return true
}

func (c *OutboundPaymentImportEntryListPagingIterator) Value(ctx context.Context) (*OutboundPaymentImportEntryListResult, error) {
	if !c.Next() {
		return c.response, nil
	}

	s := c.service
	p := c.params
	p.After = c.cursor

	uri, err := url.Parse(fmt.Sprintf(s.config.Endpoint() + "/outbound_payment_import_entries"))

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
		*OutboundPaymentImportEntryListResult
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

	if result.OutboundPaymentImportEntryListResult == nil {
		return nil, errors.New("missing result")
	}

	c.response = result.OutboundPaymentImportEntryListResult
	c.cursor = c.response.Meta.Cursors.After
	return c.response, nil
}

func (s *OutboundPaymentImportEntryServiceImpl) All(ctx context.Context,
	p OutboundPaymentImportEntryListParams,
	opts ...RequestOption) *OutboundPaymentImportEntryListPagingIterator {
	return &OutboundPaymentImportEntryListPagingIterator{
		params:         p,
		service:        s,
		requestOptions: opts,
	}
}
