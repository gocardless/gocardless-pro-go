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

// ExportService manages exports
type ExportServiceImpl struct {
	config Config
}

// Export model
type Export struct {
	CreatedAt   string `url:"created_at,omitempty" json:"created_at,omitempty"`
	Currency    string `url:"currency,omitempty" json:"currency,omitempty"`
	DownloadUrl string `url:"download_url,omitempty" json:"download_url,omitempty"`
	ExportType  string `url:"export_type,omitempty" json:"export_type,omitempty"`
	Id          string `url:"id,omitempty" json:"id,omitempty"`
}

type ExportService interface {
	Get(ctx context.Context, identity string, opts ...RequestOption) (*Export, error)
	List(ctx context.Context, p ExportListParams, opts ...RequestOption) (*ExportListResult, error)
	All(ctx context.Context, p ExportListParams, opts ...RequestOption) *ExportListPagingIterator
}

// Get
// Returns a single export.
func (s *ExportServiceImpl) Get(ctx context.Context, identity string, opts ...RequestOption) (*Export, error) {
	uri, err := url.Parse(fmt.Sprintf(s.config.Endpoint()+"/exports/%v",
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
	req.Header.Set("GoCardless-Client-Version", "4.6.0")
	req.Header.Set("User-Agent", userAgent)

	for key, value := range o.headers {
		req.Header.Set(key, value)
	}

	client := s.config.Client()
	if client == nil {
		client = http.DefaultClient
	}

	var result struct {
		Err    *APIError `json:"error"`
		Export *Export   `json:"exports"`
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

	if result.Export == nil {
		return nil, errors.New("missing result")
	}

	return result.Export, nil
}

// ExportListParams parameters
type ExportListParams struct {
	After  string `url:"after,omitempty" json:"after,omitempty"`
	Before string `url:"before,omitempty" json:"before,omitempty"`
	Limit  int    `url:"limit,omitempty" json:"limit,omitempty"`
}

type ExportListResultMetaCursors struct {
	After  string `url:"after,omitempty" json:"after,omitempty"`
	Before string `url:"before,omitempty" json:"before,omitempty"`
}

type ExportListResultMeta struct {
	Cursors *ExportListResultMetaCursors `url:"cursors,omitempty" json:"cursors,omitempty"`
	Limit   int                          `url:"limit,omitempty" json:"limit,omitempty"`
}

type ExportListResult struct {
	Exports []Export             `json:"exports"`
	Meta    ExportListResultMeta `url:"meta,omitempty" json:"meta,omitempty"`
}

// List
// Returns a list of exports which are available for download.
func (s *ExportServiceImpl) List(ctx context.Context, p ExportListParams, opts ...RequestOption) (*ExportListResult, error) {
	uri, err := url.Parse(fmt.Sprintf(s.config.Endpoint() + "/exports"))
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
	req.Header.Set("GoCardless-Client-Version", "4.6.0")
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
		*ExportListResult
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

	if result.ExportListResult == nil {
		return nil, errors.New("missing result")
	}

	return result.ExportListResult, nil
}

type ExportListPagingIterator struct {
	cursor         string
	response       *ExportListResult
	params         ExportListParams
	service        *ExportServiceImpl
	requestOptions []RequestOption
}

func (c *ExportListPagingIterator) Next() bool {
	if c.cursor == "" && c.response != nil {
		return false
	}

	return true
}

func (c *ExportListPagingIterator) Value(ctx context.Context) (*ExportListResult, error) {
	if !c.Next() {
		return c.response, nil
	}

	s := c.service
	p := c.params
	p.After = c.cursor

	uri, err := url.Parse(fmt.Sprintf(s.config.Endpoint() + "/exports"))

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
	req.Header.Set("GoCardless-Client-Version", "4.6.0")
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
		*ExportListResult
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

	if result.ExportListResult == nil {
		return nil, errors.New("missing result")
	}

	c.response = result.ExportListResult
	c.cursor = c.response.Meta.Cursors.After
	return c.response, nil
}

func (s *ExportServiceImpl) All(ctx context.Context,
	p ExportListParams,
	opts ...RequestOption) *ExportListPagingIterator {
	return &ExportListPagingIterator{
		params:         p,
		service:        s,
		requestOptions: opts,
	}
}
