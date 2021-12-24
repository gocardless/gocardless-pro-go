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

// RefundService manages refunds
type RefundService struct {
	endpoint string
	token    string
	client   *http.Client
}

// Refund model
type Refund struct {
	Amount    int    `url:"amount,omitempty" json:"amount,omitempty"`
	CreatedAt string `url:"created_at,omitempty" json:"created_at,omitempty"`
	Currency  string `url:"currency,omitempty" json:"currency,omitempty"`
	Fx        struct {
		EstimatedExchangeRate string `url:"estimated_exchange_rate,omitempty" json:"estimated_exchange_rate,omitempty"`
		ExchangeRate          string `url:"exchange_rate,omitempty" json:"exchange_rate,omitempty"`
		FxAmount              int    `url:"fx_amount,omitempty" json:"fx_amount,omitempty"`
		FxCurrency            string `url:"fx_currency,omitempty" json:"fx_currency,omitempty"`
	} `url:"fx,omitempty" json:"fx,omitempty"`
	Id    string `url:"id,omitempty" json:"id,omitempty"`
	Links struct {
		Mandate string `url:"mandate,omitempty" json:"mandate,omitempty"`
		Payment string `url:"payment,omitempty" json:"payment,omitempty"`
	} `url:"links,omitempty" json:"links,omitempty"`
	Metadata  map[string]interface{} `url:"metadata,omitempty" json:"metadata,omitempty"`
	Reference string                 `url:"reference,omitempty" json:"reference,omitempty"`
	Status    string                 `url:"status,omitempty" json:"status,omitempty"`
}

// RefundCreateParams parameters
type RefundCreateParams struct {
	Amount int `url:"amount,omitempty" json:"amount,omitempty"`
	Links  struct {
		Mandate string `url:"mandate,omitempty" json:"mandate,omitempty"`
		Payment string `url:"payment,omitempty" json:"payment,omitempty"`
	} `url:"links,omitempty" json:"links,omitempty"`
	Metadata                map[string]interface{} `url:"metadata,omitempty" json:"metadata,omitempty"`
	Reference               string                 `url:"reference,omitempty" json:"reference,omitempty"`
	TotalAmountConfirmation int                    `url:"total_amount_confirmation,omitempty" json:"total_amount_confirmation,omitempty"`
}

// Create
// Creates a new refund object.
//
// This fails with:<a name="total_amount_confirmation_invalid"></a><a
// name="number_of_refunds_exceeded"></a><a
// name="available_refund_amount_insufficient"></a>
//
// - `total_amount_confirmation_invalid` if the confirmation amount doesn't
// match the total amount refunded for the payment. This safeguard is there to
// prevent two processes from creating refunds without awareness of each other.
//
// - `number_of_refunds_exceeded` if five or more refunds have already been
// created against the payment.
//
// - `available_refund_amount_insufficient` if the creditor does not have
// sufficient balance for refunds available to cover the cost of the requested
// refund.
//
func (s *RefundService) Create(ctx context.Context, p RefundCreateParams, opts ...RequestOption) (*Refund, error) {
	uri, err := url.Parse(fmt.Sprintf(s.endpoint + "/refunds"))
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
		"refunds": p,
	})
	if err != nil {
		return nil, err
	}
	body = &buf

	req, err := http.NewRequest("POST", uri.String(), body)
	if err != nil {
		return nil, err
	}
	req.WithContext(ctx)
	req.Header.Set("Authorization", "Bearer "+s.token)

	req.Header.Set("GoCardless-Version", "2015-07-06")

	req.Header.Set("GoCardless-Client-Library", "<no value>")
	req.Header.Set("GoCardless-Client-Version", "1.0.0")
	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Idempotency-Key", o.idempotencyKey)

	for key, value := range o.headers {
		req.Header.Set(key, value)
	}

	client := s.client
	if client == nil {
		client = http.DefaultClient
	}

	var result struct {
		Err    *APIError `json:"error"`
		Refund *Refund   `json:"refunds"`
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

	if result.Refund == nil {
		return nil, errors.New("missing result")
	}

	return result.Refund, nil
}

// RefundListParams parameters
type RefundListParams struct {
	After     string `url:"after,omitempty" json:"after,omitempty"`
	Before    string `url:"before,omitempty" json:"before,omitempty"`
	CreatedAt struct {
		Gt  string `url:"gt,omitempty" json:"gt,omitempty"`
		Gte string `url:"gte,omitempty" json:"gte,omitempty"`
		Lt  string `url:"lt,omitempty" json:"lt,omitempty"`
		Lte string `url:"lte,omitempty" json:"lte,omitempty"`
	} `url:"created_at,omitempty" json:"created_at,omitempty"`
	Limit      int    `url:"limit,omitempty" json:"limit,omitempty"`
	Mandate    string `url:"mandate,omitempty" json:"mandate,omitempty"`
	Payment    string `url:"payment,omitempty" json:"payment,omitempty"`
	RefundType string `url:"refund_type,omitempty" json:"refund_type,omitempty"`
}

// RefundListResult response including pagination metadata
type RefundListResult struct {
	Refunds []Refund `json:"refunds"`
	Meta    struct {
		Cursors struct {
			After  string `url:"after,omitempty" json:"after,omitempty"`
			Before string `url:"before,omitempty" json:"before,omitempty"`
		} `url:"cursors,omitempty" json:"cursors,omitempty"`
		Limit int `url:"limit,omitempty" json:"limit,omitempty"`
	} `json:"meta"`
}

// List
// Returns a [cursor-paginated](#api-usage-cursor-pagination) list of your
// refunds.
func (s *RefundService) List(ctx context.Context, p RefundListParams, opts ...RequestOption) (*RefundListResult, error) {
	uri, err := url.Parse(fmt.Sprintf(s.endpoint + "/refunds"))
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
	req.WithContext(ctx)
	req.Header.Set("Authorization", "Bearer "+s.token)

	req.Header.Set("GoCardless-Version", "2015-07-06")

	req.Header.Set("GoCardless-Client-Library", "<no value>")
	req.Header.Set("GoCardless-Client-Version", "1.0.0")
	req.Header.Set("User-Agent", userAgent)

	for key, value := range o.headers {
		req.Header.Set(key, value)
	}

	client := s.client
	if client == nil {
		client = http.DefaultClient
	}

	var result struct {
		Err *APIError `json:"error"`
		*RefundListResult
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

	if result.RefundListResult == nil {
		return nil, errors.New("missing result")
	}

	return result.RefundListResult, nil
}

type RefundListPagingIterator struct {
	cursor   string
	response *RefundListResult
	params   RefundListParams
	service  *RefundService
}

func (c *RefundListPagingIterator) Next() bool {
	if c.cursor == "" && c.response != nil {
		return false
	}

	return true
}

func (c *RefundListPagingIterator) Value(ctx context.Context) (*RefundListResult, error) {
	if !c.Next() {
		return c.response, nil
	}

	s := c.service
	p := c.params
	p.After = c.cursor

	uri, err := url.Parse(fmt.Sprintf(s.endpoint + "/refunds"))

	if err != nil {
		return nil, err
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
	req.WithContext(ctx)
	req.Header.Set("Authorization", "Bearer "+s.token)
	req.Header.Set("GoCardless-Version", "2015-07-06")

	client := s.client
	if client == nil {
		client = http.DefaultClient
	}

	var result struct {
		Err *APIError `json:"error"`
		*RefundListResult
	}

	err = try(3, func() error {
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

	if result.RefundListResult == nil {
		return nil, errors.New("missing result")
	}

	c.response = result.RefundListResult
	c.cursor = c.response.Meta.Cursors.After
	return c.response, nil
}

func (s *RefundService) All(ctx context.Context, p RefundListParams) *RefundListPagingIterator {
	return &RefundListPagingIterator{
		params:  p,
		service: s,
	}
}

// Get
// Retrieves all details for a single refund
func (s *RefundService) Get(ctx context.Context, identity string, opts ...RequestOption) (*Refund, error) {
	uri, err := url.Parse(fmt.Sprintf(s.endpoint+"/refunds/%v",
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
	req.WithContext(ctx)
	req.Header.Set("Authorization", "Bearer "+s.token)

	req.Header.Set("GoCardless-Version", "2015-07-06")

	req.Header.Set("GoCardless-Client-Library", "<no value>")
	req.Header.Set("GoCardless-Client-Version", "1.0.0")
	req.Header.Set("User-Agent", userAgent)

	for key, value := range o.headers {
		req.Header.Set(key, value)
	}

	client := s.client
	if client == nil {
		client = http.DefaultClient
	}

	var result struct {
		Err    *APIError `json:"error"`
		Refund *Refund   `json:"refunds"`
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

	if result.Refund == nil {
		return nil, errors.New("missing result")
	}

	return result.Refund, nil
}

// RefundUpdateParams parameters
type RefundUpdateParams struct {
	Metadata map[string]interface{} `url:"metadata,omitempty" json:"metadata,omitempty"`
}

// Update
// Updates a refund object.
func (s *RefundService) Update(ctx context.Context, identity string, p RefundUpdateParams, opts ...RequestOption) (*Refund, error) {
	uri, err := url.Parse(fmt.Sprintf(s.endpoint+"/refunds/%v",
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
		"refunds": p,
	})
	if err != nil {
		return nil, err
	}
	body = &buf

	req, err := http.NewRequest("PUT", uri.String(), body)
	if err != nil {
		return nil, err
	}
	req.WithContext(ctx)
	req.Header.Set("Authorization", "Bearer "+s.token)

	req.Header.Set("GoCardless-Version", "2015-07-06")

	req.Header.Set("GoCardless-Client-Library", "<no value>")
	req.Header.Set("GoCardless-Client-Version", "1.0.0")
	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Idempotency-Key", o.idempotencyKey)

	for key, value := range o.headers {
		req.Header.Set(key, value)
	}

	client := s.client
	if client == nil {
		client = http.DefaultClient
	}

	var result struct {
		Err    *APIError `json:"error"`
		Refund *Refund   `json:"refunds"`
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

	if result.Refund == nil {
		return nil, errors.New("missing result")
	}

	return result.Refund, nil
}
