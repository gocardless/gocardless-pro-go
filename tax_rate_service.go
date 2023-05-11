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

// TaxRateService manages tax_rates
type TaxRateServiceImpl struct {
	config Config
}

// TaxRate model
type TaxRate struct {
	EndDate      string `url:"end_date,omitempty" json:"end_date,omitempty"`
	Id           string `url:"id,omitempty" json:"id,omitempty"`
	Jurisdiction string `url:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	Percentage   string `url:"percentage,omitempty" json:"percentage,omitempty"`
	StartDate    string `url:"start_date,omitempty" json:"start_date,omitempty"`
	Type         string `url:"type,omitempty" json:"type,omitempty"`
}

type TaxRateService interface {
	List(ctx context.Context, p TaxRateListParams, opts ...RequestOption) (*TaxRateListResult, error)
	All(ctx context.Context, p TaxRateListParams, opts ...RequestOption) *TaxRateListPagingIterator
	Get(ctx context.Context, identity string, opts ...RequestOption) (*TaxRate, error)
}

// TaxRateListParams parameters
type TaxRateListParams struct {
	After        string `url:"after,omitempty" json:"after,omitempty"`
	Before       string `url:"before,omitempty" json:"before,omitempty"`
	Jurisdiction string `url:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
}

type TaxRateListResultMetaCursors struct {
	After  string `url:"after,omitempty" json:"after,omitempty"`
	Before string `url:"before,omitempty" json:"before,omitempty"`
}

type TaxRateListResultMeta struct {
	Cursors TaxRateListResultMetaCursors `url:"cursors,omitempty" json:"cursors,omitempty"`
	Limit   int                          `url:"limit,omitempty" json:"limit,omitempty"`
}

type TaxRateListResult struct {
	TaxRates []TaxRate             `json:"tax_rates"`
	Meta     TaxRateListResultMeta `url:"meta,omitempty" json:"meta,omitempty"`
}

// List
// Returns a [cursor-paginated](#api-usage-cursor-pagination) list of all tax
// rates.
func (s *TaxRateServiceImpl) List(ctx context.Context, p TaxRateListParams, opts ...RequestOption) (*TaxRateListResult, error) {
	uri, err := url.Parse(fmt.Sprintf(s.config.Endpoint() + "/tax_rates"))
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
	req.Header.Set("GoCardless-Client-Version", "3.2.0")
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
		*TaxRateListResult
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

	if result.TaxRateListResult == nil {
		return nil, errors.New("missing result")
	}

	return result.TaxRateListResult, nil
}

type TaxRateListPagingIterator struct {
	cursor         string
	response       *TaxRateListResult
	params         TaxRateListParams
	service        *TaxRateServiceImpl
	requestOptions []RequestOption
}

func (c *TaxRateListPagingIterator) Next() bool {
	if c.cursor == "" && c.response != nil {
		return false
	}

	return true
}

func (c *TaxRateListPagingIterator) Value(ctx context.Context) (*TaxRateListResult, error) {
	if !c.Next() {
		return c.response, nil
	}

	s := c.service
	p := c.params
	p.After = c.cursor

	uri, err := url.Parse(fmt.Sprintf(s.config.Endpoint() + "/tax_rates"))

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
	req.Header.Set("GoCardless-Client-Version", "3.2.0")
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
		*TaxRateListResult
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

	if result.TaxRateListResult == nil {
		return nil, errors.New("missing result")
	}

	c.response = result.TaxRateListResult
	c.cursor = c.response.Meta.Cursors.After
	return c.response, nil
}

func (s *TaxRateServiceImpl) All(ctx context.Context,
	p TaxRateListParams,
	opts ...RequestOption) *TaxRateListPagingIterator {
	return &TaxRateListPagingIterator{
		params:         p,
		service:        s,
		requestOptions: opts,
	}
}

// Get
// Retrieves the details of a tax rate.
func (s *TaxRateServiceImpl) Get(ctx context.Context, identity string, opts ...RequestOption) (*TaxRate, error) {
	uri, err := url.Parse(fmt.Sprintf(s.config.Endpoint()+"/tax_rates/%v",
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
	req.Header.Set("GoCardless-Client-Version", "3.2.0")
	req.Header.Set("User-Agent", userAgent)

	for key, value := range o.headers {
		req.Header.Set(key, value)
	}

	client := s.config.Client()
	if client == nil {
		client = http.DefaultClient
	}

	var result struct {
		Err     *APIError `json:"error"`
		TaxRate *TaxRate  `json:"tax_rates"`
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

	if result.TaxRate == nil {
		return nil, errors.New("missing result")
	}

	return result.TaxRate, nil
}
