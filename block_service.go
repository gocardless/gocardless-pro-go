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

// BlockService manages blocks
type BlockServiceImpl struct {
	config Config
}

// Block model
type Block struct {
	Active            bool   `url:"active,omitempty" json:"active,omitempty"`
	BlockType         string `url:"block_type,omitempty" json:"block_type,omitempty"`
	CreatedAt         string `url:"created_at,omitempty" json:"created_at,omitempty"`
	Id                string `url:"id,omitempty" json:"id,omitempty"`
	ReasonDescription string `url:"reason_description,omitempty" json:"reason_description,omitempty"`
	ReasonType        string `url:"reason_type,omitempty" json:"reason_type,omitempty"`
	ResourceReference string `url:"resource_reference,omitempty" json:"resource_reference,omitempty"`
	UpdatedAt         string `url:"updated_at,omitempty" json:"updated_at,omitempty"`
}

type BlockService interface {
	Create(ctx context.Context, p BlockCreateParams, opts ...RequestOption) (*Block, error)
	Get(ctx context.Context, identity string, opts ...RequestOption) (*Block, error)
	List(ctx context.Context, p BlockListParams, opts ...RequestOption) (*BlockListResult, error)
	All(ctx context.Context, p BlockListParams, opts ...RequestOption) *BlockListPagingIterator
	Disable(ctx context.Context, identity string, opts ...RequestOption) (*Block, error)
	Enable(ctx context.Context, identity string, opts ...RequestOption) (*Block, error)
	BlockByRef(ctx context.Context, p BlockBlockByRefParams, opts ...RequestOption) (
		*BlockBlockByRefResult, error)
}

// BlockCreateParams parameters
type BlockCreateParams struct {
	Active            bool   `url:"active,omitempty" json:"active,omitempty"`
	BlockType         string `url:"block_type,omitempty" json:"block_type,omitempty"`
	ReasonDescription string `url:"reason_description,omitempty" json:"reason_description,omitempty"`
	ReasonType        string `url:"reason_type,omitempty" json:"reason_type,omitempty"`
	ResourceReference string `url:"resource_reference,omitempty" json:"resource_reference,omitempty"`
}

// Create
// Creates a new Block of a given type. By default it will be active.
func (s *BlockServiceImpl) Create(ctx context.Context, p BlockCreateParams, opts ...RequestOption) (*Block, error) {
	uri, err := url.Parse(fmt.Sprintf(s.config.Endpoint() + "/blocks"))
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
		"blocks": p,
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
	req.Header.Set("GoCardless-Client-Version", "2.6.0")
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
		Err   *APIError `json:"error"`
		Block *Block    `json:"blocks"`
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

	if result.Block == nil {
		return nil, errors.New("missing result")
	}

	return result.Block, nil
}

// Get
// Retrieves the details of an existing block.
func (s *BlockServiceImpl) Get(ctx context.Context, identity string, opts ...RequestOption) (*Block, error) {
	uri, err := url.Parse(fmt.Sprintf(s.config.Endpoint()+"/blocks/%v",
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
	req.Header.Set("GoCardless-Client-Version", "2.6.0")
	req.Header.Set("User-Agent", userAgent)

	for key, value := range o.headers {
		req.Header.Set(key, value)
	}

	client := s.config.Client()
	if client == nil {
		client = http.DefaultClient
	}

	var result struct {
		Err   *APIError `json:"error"`
		Block *Block    `json:"blocks"`
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

	if result.Block == nil {
		return nil, errors.New("missing result")
	}

	return result.Block, nil
}

// BlockListParams parameters
type BlockListParams struct {
	After      string `url:"after,omitempty" json:"after,omitempty"`
	Before     string `url:"before,omitempty" json:"before,omitempty"`
	Block      string `url:"block,omitempty" json:"block,omitempty"`
	BlockType  string `url:"block_type,omitempty" json:"block_type,omitempty"`
	CreatedAt  string `url:"created_at,omitempty" json:"created_at,omitempty"`
	Limit      int    `url:"limit,omitempty" json:"limit,omitempty"`
	ReasonType string `url:"reason_type,omitempty" json:"reason_type,omitempty"`
	UpdatedAt  string `url:"updated_at,omitempty" json:"updated_at,omitempty"`
}

type BlockListResultMetaCursors struct {
	After  string `url:"after,omitempty" json:"after,omitempty"`
	Before string `url:"before,omitempty" json:"before,omitempty"`
}

type BlockListResultMeta struct {
	Cursors *BlockListResultMetaCursors `url:"cursors,omitempty" json:"cursors,omitempty"`
	Limit   int                         `url:"limit,omitempty" json:"limit,omitempty"`
}

type BlockListResult struct {
	Blocks []Block             `json:"blocks"`
	Meta   BlockListResultMeta `url:"meta,omitempty" json:"meta,omitempty"`
}

// List
// Returns a [cursor-paginated](#api-usage-cursor-pagination) list of your
// blocks.
func (s *BlockServiceImpl) List(ctx context.Context, p BlockListParams, opts ...RequestOption) (*BlockListResult, error) {
	uri, err := url.Parse(fmt.Sprintf(s.config.Endpoint() + "/blocks"))
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
	req.Header.Set("GoCardless-Client-Version", "2.6.0")
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
		*BlockListResult
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

	if result.BlockListResult == nil {
		return nil, errors.New("missing result")
	}

	return result.BlockListResult, nil
}

type BlockListPagingIterator struct {
	cursor         string
	response       *BlockListResult
	params         BlockListParams
	service        *BlockServiceImpl
	requestOptions []RequestOption
}

func (c *BlockListPagingIterator) Next() bool {
	if c.cursor == "" && c.response != nil {
		return false
	}

	return true
}

func (c *BlockListPagingIterator) Value(ctx context.Context) (*BlockListResult, error) {
	if !c.Next() {
		return c.response, nil
	}

	s := c.service
	p := c.params
	p.After = c.cursor

	uri, err := url.Parse(fmt.Sprintf(s.config.Endpoint() + "/blocks"))

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
	req.Header.Set("GoCardless-Client-Version", "2.6.0")
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
		*BlockListResult
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

	if result.BlockListResult == nil {
		return nil, errors.New("missing result")
	}

	c.response = result.BlockListResult
	c.cursor = c.response.Meta.Cursors.After
	return c.response, nil
}

func (s *BlockServiceImpl) All(ctx context.Context,
	p BlockListParams,
	opts ...RequestOption) *BlockListPagingIterator {
	return &BlockListPagingIterator{
		params:         p,
		service:        s,
		requestOptions: opts,
	}
}

// Disable
// Disables a block so that it no longer will prevent mandate creation.
func (s *BlockServiceImpl) Disable(ctx context.Context, identity string, opts ...RequestOption) (*Block, error) {
	uri, err := url.Parse(fmt.Sprintf(s.config.Endpoint()+"/blocks/%v/actions/disable",
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
	req.Header.Set("GoCardless-Client-Version", "2.6.0")
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
		Err   *APIError `json:"error"`
		Block *Block    `json:"blocks"`
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

	if result.Block == nil {
		return nil, errors.New("missing result")
	}

	return result.Block, nil
}

// Enable
// Enables a previously disabled block so that it will prevent mandate creation
func (s *BlockServiceImpl) Enable(ctx context.Context, identity string, opts ...RequestOption) (*Block, error) {
	uri, err := url.Parse(fmt.Sprintf(s.config.Endpoint()+"/blocks/%v/actions/enable",
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
	req.Header.Set("GoCardless-Client-Version", "2.6.0")
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
		Err   *APIError `json:"error"`
		Block *Block    `json:"blocks"`
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

	if result.Block == nil {
		return nil, errors.New("missing result")
	}

	return result.Block, nil
}

// BlockBlockByRefParams parameters
type BlockBlockByRefParams struct {
	Active            bool   `url:"active,omitempty" json:"active,omitempty"`
	ReasonDescription string `url:"reason_description,omitempty" json:"reason_description,omitempty"`
	ReasonType        string `url:"reason_type,omitempty" json:"reason_type,omitempty"`
	ReferenceType     string `url:"reference_type,omitempty" json:"reference_type,omitempty"`
	ReferenceValue    string `url:"reference_value,omitempty" json:"reference_value,omitempty"`
}

type BlockBlockByRefResultMetaCursors struct {
	After  string `url:"after,omitempty" json:"after,omitempty"`
	Before string `url:"before,omitempty" json:"before,omitempty"`
}

type BlockBlockByRefResultMeta struct {
	Cursors *BlockBlockByRefResultMetaCursors `url:"cursors,omitempty" json:"cursors,omitempty"`
	Limit   int                               `url:"limit,omitempty" json:"limit,omitempty"`
}

type BlockBlockByRefResult struct {
	Blocks []Block                   `json:"blocks"`
	Meta   BlockBlockByRefResultMeta `url:"meta,omitempty" json:"meta,omitempty"`
}

// BlockByRef
// Creates new blocks for a given reference. By default blocks will be active.
// Returns 201 if at least one block was created. Returns 200 if there were no
// new
// blocks created.
func (s *BlockServiceImpl) BlockByRef(ctx context.Context, p BlockBlockByRefParams, opts ...RequestOption) (
	*BlockBlockByRefResult, error) {
	uri, err := url.Parse(fmt.Sprintf(s.config.Endpoint() + "/block_by_ref"))
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

	req, err := http.NewRequest("POST", uri.String(), body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	req.Header.Set("Authorization", "Bearer "+s.config.Token())
	req.Header.Set("GoCardless-Version", "2015-07-06")
	req.Header.Set("GoCardless-Client-Library", "gocardless-pro-go")
	req.Header.Set("GoCardless-Client-Version", "2.6.0")
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
		Err *APIError `json:"error"`

		*BlockBlockByRefResult
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

	if result.BlockBlockByRefResult == nil {
		return nil, errors.New("missing result")
	}

	return result.BlockBlockByRefResult, nil
}
