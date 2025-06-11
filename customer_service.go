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

// CustomerService manages customers
type CustomerServiceImpl struct {
	config Config
}

// Customer model
type Customer struct {
	AddressLine1          string                 `url:"address_line1,omitempty" json:"address_line1,omitempty"`
	AddressLine2          string                 `url:"address_line2,omitempty" json:"address_line2,omitempty"`
	AddressLine3          string                 `url:"address_line3,omitempty" json:"address_line3,omitempty"`
	City                  string                 `url:"city,omitempty" json:"city,omitempty"`
	CompanyName           string                 `url:"company_name,omitempty" json:"company_name,omitempty"`
	CountryCode           string                 `url:"country_code,omitempty" json:"country_code,omitempty"`
	CreatedAt             string                 `url:"created_at,omitempty" json:"created_at,omitempty"`
	DanishIdentityNumber  string                 `url:"danish_identity_number,omitempty" json:"danish_identity_number,omitempty"`
	Email                 string                 `url:"email,omitempty" json:"email,omitempty"`
	FamilyName            string                 `url:"family_name,omitempty" json:"family_name,omitempty"`
	GivenName             string                 `url:"given_name,omitempty" json:"given_name,omitempty"`
	Id                    string                 `url:"id,omitempty" json:"id,omitempty"`
	Language              string                 `url:"language,omitempty" json:"language,omitempty"`
	Metadata              map[string]interface{} `url:"metadata,omitempty" json:"metadata,omitempty"`
	PhoneNumber           string                 `url:"phone_number,omitempty" json:"phone_number,omitempty"`
	PostalCode            string                 `url:"postal_code,omitempty" json:"postal_code,omitempty"`
	Region                string                 `url:"region,omitempty" json:"region,omitempty"`
	SwedishIdentityNumber string                 `url:"swedish_identity_number,omitempty" json:"swedish_identity_number,omitempty"`
}

type CustomerService interface {
	Create(ctx context.Context, p CustomerCreateParams, opts ...RequestOption) (*Customer, error)
	List(ctx context.Context, p CustomerListParams, opts ...RequestOption) (*CustomerListResult, error)
	All(ctx context.Context, p CustomerListParams, opts ...RequestOption) *CustomerListPagingIterator
	Get(ctx context.Context, identity string, opts ...RequestOption) (*Customer, error)
	Update(ctx context.Context, identity string, p CustomerUpdateParams, opts ...RequestOption) (*Customer, error)
	Remove(ctx context.Context, identity string, p CustomerRemoveParams, opts ...RequestOption) (*Customer, error)
}

// CustomerCreateParams parameters
type CustomerCreateParams struct {
	AddressLine1          string                 `url:"address_line1,omitempty" json:"address_line1,omitempty"`
	AddressLine2          string                 `url:"address_line2,omitempty" json:"address_line2,omitempty"`
	AddressLine3          string                 `url:"address_line3,omitempty" json:"address_line3,omitempty"`
	City                  string                 `url:"city,omitempty" json:"city,omitempty"`
	CompanyName           string                 `url:"company_name,omitempty" json:"company_name,omitempty"`
	CountryCode           string                 `url:"country_code,omitempty" json:"country_code,omitempty"`
	DanishIdentityNumber  string                 `url:"danish_identity_number,omitempty" json:"danish_identity_number,omitempty"`
	Email                 string                 `url:"email,omitempty" json:"email,omitempty"`
	FamilyName            string                 `url:"family_name,omitempty" json:"family_name,omitempty"`
	GivenName             string                 `url:"given_name,omitempty" json:"given_name,omitempty"`
	Language              string                 `url:"language,omitempty" json:"language,omitempty"`
	Metadata              map[string]interface{} `url:"metadata,omitempty" json:"metadata,omitempty"`
	PhoneNumber           string                 `url:"phone_number,omitempty" json:"phone_number,omitempty"`
	PostalCode            string                 `url:"postal_code,omitempty" json:"postal_code,omitempty"`
	Region                string                 `url:"region,omitempty" json:"region,omitempty"`
	SwedishIdentityNumber string                 `url:"swedish_identity_number,omitempty" json:"swedish_identity_number,omitempty"`
}

// Create
// Creates a new customer object.
func (s *CustomerServiceImpl) Create(ctx context.Context, p CustomerCreateParams, opts ...RequestOption) (*Customer, error) {
	uri, err := url.Parse(fmt.Sprintf(s.config.Endpoint() + "/customers"))
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
		"customers": p,
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
	req.Header.Set("GoCardless-Client-Version", "5.1.0")
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
		Err      *APIError `json:"error"`
		Customer *Customer `json:"customers"`
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

	if result.Customer == nil {
		return nil, errors.New("missing result")
	}

	return result.Customer, nil
}

type CustomerListParamsCreatedAt struct {
	Gt  string `url:"gt,omitempty" json:"gt,omitempty"`
	Gte string `url:"gte,omitempty" json:"gte,omitempty"`
	Lt  string `url:"lt,omitempty" json:"lt,omitempty"`
	Lte string `url:"lte,omitempty" json:"lte,omitempty"`
}

// CustomerListParams parameters
type CustomerListParams struct {
	After         string                       `url:"after,omitempty" json:"after,omitempty"`
	Before        string                       `url:"before,omitempty" json:"before,omitempty"`
	CreatedAt     *CustomerListParamsCreatedAt `url:"created_at,omitempty" json:"created_at,omitempty"`
	Currency      string                       `url:"currency,omitempty" json:"currency,omitempty"`
	Limit         int                          `url:"limit,omitempty" json:"limit,omitempty"`
	SortDirection string                       `url:"sort_direction,omitempty" json:"sort_direction,omitempty"`
	SortField     string                       `url:"sort_field,omitempty" json:"sort_field,omitempty"`
}

type CustomerListResultMetaCursors struct {
	After  string `url:"after,omitempty" json:"after,omitempty"`
	Before string `url:"before,omitempty" json:"before,omitempty"`
}

type CustomerListResultMeta struct {
	Cursors *CustomerListResultMetaCursors `url:"cursors,omitempty" json:"cursors,omitempty"`
	Limit   int                            `url:"limit,omitempty" json:"limit,omitempty"`
}

type CustomerListResult struct {
	Customers []Customer             `json:"customers"`
	Meta      CustomerListResultMeta `url:"meta,omitempty" json:"meta,omitempty"`
}

// List
// Returns a [cursor-paginated](#api-usage-cursor-pagination) list of your
// customers.
func (s *CustomerServiceImpl) List(ctx context.Context, p CustomerListParams, opts ...RequestOption) (*CustomerListResult, error) {
	uri, err := url.Parse(fmt.Sprintf(s.config.Endpoint() + "/customers"))
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
	req.Header.Set("GoCardless-Client-Version", "5.1.0")
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
		*CustomerListResult
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

	if result.CustomerListResult == nil {
		return nil, errors.New("missing result")
	}

	return result.CustomerListResult, nil
}

type CustomerListPagingIterator struct {
	cursor         string
	response       *CustomerListResult
	params         CustomerListParams
	service        *CustomerServiceImpl
	requestOptions []RequestOption
}

func (c *CustomerListPagingIterator) Next() bool {
	if c.cursor == "" && c.response != nil {
		return false
	}

	return true
}

func (c *CustomerListPagingIterator) Value(ctx context.Context) (*CustomerListResult, error) {
	if !c.Next() {
		return c.response, nil
	}

	s := c.service
	p := c.params
	p.After = c.cursor

	uri, err := url.Parse(fmt.Sprintf(s.config.Endpoint() + "/customers"))

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
	req.Header.Set("GoCardless-Client-Version", "5.1.0")
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
		*CustomerListResult
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

	if result.CustomerListResult == nil {
		return nil, errors.New("missing result")
	}

	c.response = result.CustomerListResult
	c.cursor = c.response.Meta.Cursors.After
	return c.response, nil
}

func (s *CustomerServiceImpl) All(ctx context.Context,
	p CustomerListParams,
	opts ...RequestOption) *CustomerListPagingIterator {
	return &CustomerListPagingIterator{
		params:         p,
		service:        s,
		requestOptions: opts,
	}
}

// Get
// Retrieves the details of an existing customer.
func (s *CustomerServiceImpl) Get(ctx context.Context, identity string, opts ...RequestOption) (*Customer, error) {
	uri, err := url.Parse(fmt.Sprintf(s.config.Endpoint()+"/customers/%v",
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
	req.Header.Set("GoCardless-Client-Version", "5.1.0")
	req.Header.Set("User-Agent", userAgent)

	for key, value := range o.headers {
		req.Header.Set(key, value)
	}

	client := s.config.Client()
	if client == nil {
		client = http.DefaultClient
	}

	var result struct {
		Err      *APIError `json:"error"`
		Customer *Customer `json:"customers"`
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

	if result.Customer == nil {
		return nil, errors.New("missing result")
	}

	return result.Customer, nil
}

// CustomerUpdateParams parameters
type CustomerUpdateParams struct {
	AddressLine1          string                 `url:"address_line1,omitempty" json:"address_line1,omitempty"`
	AddressLine2          string                 `url:"address_line2,omitempty" json:"address_line2,omitempty"`
	AddressLine3          string                 `url:"address_line3,omitempty" json:"address_line3,omitempty"`
	City                  string                 `url:"city,omitempty" json:"city,omitempty"`
	CompanyName           string                 `url:"company_name,omitempty" json:"company_name,omitempty"`
	CountryCode           string                 `url:"country_code,omitempty" json:"country_code,omitempty"`
	DanishIdentityNumber  string                 `url:"danish_identity_number,omitempty" json:"danish_identity_number,omitempty"`
	Email                 string                 `url:"email,omitempty" json:"email,omitempty"`
	FamilyName            string                 `url:"family_name,omitempty" json:"family_name,omitempty"`
	GivenName             string                 `url:"given_name,omitempty" json:"given_name,omitempty"`
	Language              string                 `url:"language,omitempty" json:"language,omitempty"`
	Metadata              map[string]interface{} `url:"metadata,omitempty" json:"metadata,omitempty"`
	PhoneNumber           string                 `url:"phone_number,omitempty" json:"phone_number,omitempty"`
	PostalCode            string                 `url:"postal_code,omitempty" json:"postal_code,omitempty"`
	Region                string                 `url:"region,omitempty" json:"region,omitempty"`
	SwedishIdentityNumber string                 `url:"swedish_identity_number,omitempty" json:"swedish_identity_number,omitempty"`
}

// Update
// Updates a customer object. Supports all of the fields supported when creating
// a customer.
func (s *CustomerServiceImpl) Update(ctx context.Context, identity string, p CustomerUpdateParams, opts ...RequestOption) (*Customer, error) {
	uri, err := url.Parse(fmt.Sprintf(s.config.Endpoint()+"/customers/%v",
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

	var buf bytes.Buffer
	err = json.NewEncoder(&buf).Encode(map[string]interface{}{
		"customers": p,
	})
	if err != nil {
		return nil, err
	}
	body = &buf

	req, err := http.NewRequest("PUT", uri.String(), body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	req.Header.Set("Authorization", "Bearer "+s.config.Token())
	req.Header.Set("GoCardless-Version", "2015-07-06")
	req.Header.Set("GoCardless-Client-Library", "gocardless-pro-go")
	req.Header.Set("GoCardless-Client-Version", "5.1.0")
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
		Err      *APIError `json:"error"`
		Customer *Customer `json:"customers"`
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

	if result.Customer == nil {
		return nil, errors.New("missing result")
	}

	return result.Customer, nil
}

// CustomerRemoveParams parameters
type CustomerRemoveParams struct {
}

// Remove
// Removed customers will not appear in search results or lists of customers (in
// our API
// or exports), and it will not be possible to load an individually removed
// customer by
// ID.
//
// <p class="restricted-notice"><strong>The action of removing a customer cannot
// be reversed, so please use with care.</strong></p>
func (s *CustomerServiceImpl) Remove(ctx context.Context, identity string, p CustomerRemoveParams, opts ...RequestOption) (*Customer, error) {
	uri, err := url.Parse(fmt.Sprintf(s.config.Endpoint()+"/customers/%v",
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

	var buf bytes.Buffer
	err = json.NewEncoder(&buf).Encode(map[string]interface{}{
		"data": p,
	})
	if err != nil {
		return nil, err
	}
	body = &buf

	req, err := http.NewRequest("DELETE", uri.String(), body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	req.Header.Set("Authorization", "Bearer "+s.config.Token())
	req.Header.Set("GoCardless-Version", "2015-07-06")
	req.Header.Set("GoCardless-Client-Library", "gocardless-pro-go")
	req.Header.Set("GoCardless-Client-Version", "5.1.0")
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
		Err      *APIError `json:"error"`
		Customer *Customer `json:"customers"`
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

	if result.Customer == nil {
		return nil, errors.New("missing result")
	}

	return result.Customer, nil
}
