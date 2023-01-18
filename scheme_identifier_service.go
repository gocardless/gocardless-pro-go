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

// SchemeIdentifierService manages scheme_identifiers
type SchemeIdentifierServiceImpl struct {
	config Config
}

// SchemeIdentifier model
type SchemeIdentifier struct {
	AddressLine1               string `url:"address_line1,omitempty" json:"address_line1,omitempty"`
	AddressLine2               string `url:"address_line2,omitempty" json:"address_line2,omitempty"`
	AddressLine3               string `url:"address_line3,omitempty" json:"address_line3,omitempty"`
	CanSpecifyMandateReference bool   `url:"can_specify_mandate_reference,omitempty" json:"can_specify_mandate_reference,omitempty"`
	City                       string `url:"city,omitempty" json:"city,omitempty"`
	CountryCode                string `url:"country_code,omitempty" json:"country_code,omitempty"`
	CreatedAt                  string `url:"created_at,omitempty" json:"created_at,omitempty"`
	Currency                   string `url:"currency,omitempty" json:"currency,omitempty"`
	Email                      string `url:"email,omitempty" json:"email,omitempty"`
	Id                         string `url:"id,omitempty" json:"id,omitempty"`
	MinimumAdvanceNotice       int    `url:"minimum_advance_notice,omitempty" json:"minimum_advance_notice,omitempty"`
	Name                       string `url:"name,omitempty" json:"name,omitempty"`
	PhoneNumber                string `url:"phone_number,omitempty" json:"phone_number,omitempty"`
	PostalCode                 string `url:"postal_code,omitempty" json:"postal_code,omitempty"`
	Reference                  string `url:"reference,omitempty" json:"reference,omitempty"`
	Region                     string `url:"region,omitempty" json:"region,omitempty"`
	Scheme                     string `url:"scheme,omitempty" json:"scheme,omitempty"`
	Status                     string `url:"status,omitempty" json:"status,omitempty"`
}

type SchemeIdentifierService interface {
	List(ctx context.Context, p SchemeIdentifierListParams, opts ...RequestOption) (*SchemeIdentifierListResult, error)
	All(ctx context.Context, p SchemeIdentifierListParams, opts ...RequestOption) *SchemeIdentifierListPagingIterator
	Get(ctx context.Context, identity string, opts ...RequestOption) (*SchemeIdentifier, error)
}

// SchemeIdentifierListParams parameters
type SchemeIdentifierListParams struct {
	After  string `url:"after,omitempty" json:"after,omitempty"`
	Before string `url:"before,omitempty" json:"before,omitempty"`
	Limit  int    `url:"limit,omitempty" json:"limit,omitempty"`
}

type SchemeIdentifierListResultMetaCursors struct {
	After  string `url:"after,omitempty" json:"after,omitempty"`
	Before string `url:"before,omitempty" json:"before,omitempty"`
}

type SchemeIdentifierListResultMeta struct {
	Cursors *SchemeIdentifierListResultMetaCursors `url:"cursors,omitempty" json:"cursors,omitempty"`
	Limit   int                                    `url:"limit,omitempty" json:"limit,omitempty"`
}

type SchemeIdentifierListResult struct {
	SchemeIdentifiers []SchemeIdentifier             `json:"scheme_identifiers"`
	Meta              SchemeIdentifierListResultMeta `url:"meta,omitempty" json:"meta,omitempty"`
}

// List
// Returns a [cursor-paginated](#api-usage-cursor-pagination) list of your
// scheme identifiers.
func (s *SchemeIdentifierServiceImpl) List(ctx context.Context, p SchemeIdentifierListParams, opts ...RequestOption) (*SchemeIdentifierListResult, error) {
	uri, err := url.Parse(fmt.Sprintf(s.config.Endpoint() + "/scheme_identifiers"))
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
	req.Header.Set("GoCardless-Client-Version", "2.9.0")
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
		*SchemeIdentifierListResult
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

	if result.SchemeIdentifierListResult == nil {
		return nil, errors.New("missing result")
	}

	return result.SchemeIdentifierListResult, nil
}

type SchemeIdentifierListPagingIterator struct {
	cursor         string
	response       *SchemeIdentifierListResult
	params         SchemeIdentifierListParams
	service        *SchemeIdentifierServiceImpl
	requestOptions []RequestOption
}

func (c *SchemeIdentifierListPagingIterator) Next() bool {
	if c.cursor == "" && c.response != nil {
		return false
	}

	return true
}

func (c *SchemeIdentifierListPagingIterator) Value(ctx context.Context) (*SchemeIdentifierListResult, error) {
	if !c.Next() {
		return c.response, nil
	}

	s := c.service
	p := c.params
	p.After = c.cursor

	uri, err := url.Parse(fmt.Sprintf(s.config.Endpoint() + "/scheme_identifiers"))

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
	req.Header.Set("GoCardless-Client-Version", "2.9.0")
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
		*SchemeIdentifierListResult
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

	if result.SchemeIdentifierListResult == nil {
		return nil, errors.New("missing result")
	}

	c.response = result.SchemeIdentifierListResult
	c.cursor = c.response.Meta.Cursors.After
	return c.response, nil
}

func (s *SchemeIdentifierServiceImpl) All(ctx context.Context,
	p SchemeIdentifierListParams,
	opts ...RequestOption) *SchemeIdentifierListPagingIterator {
	return &SchemeIdentifierListPagingIterator{
		params:         p,
		service:        s,
		requestOptions: opts,
	}
}

// Get
// Retrieves the details of an existing scheme identifier.
func (s *SchemeIdentifierServiceImpl) Get(ctx context.Context, identity string, opts ...RequestOption) (*SchemeIdentifier, error) {
	uri, err := url.Parse(fmt.Sprintf(s.config.Endpoint()+"/scheme_identifiers/%v",
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
	req.Header.Set("GoCardless-Client-Version", "2.9.0")
	req.Header.Set("User-Agent", userAgent)

	for key, value := range o.headers {
		req.Header.Set(key, value)
	}

	client := s.config.Client()
	if client == nil {
		client = http.DefaultClient
	}

	var result struct {
		Err              *APIError         `json:"error"`
		SchemeIdentifier *SchemeIdentifier `json:"scheme_identifiers"`
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

	if result.SchemeIdentifier == nil {
		return nil, errors.New("missing result")
	}

	return result.SchemeIdentifier, nil
}