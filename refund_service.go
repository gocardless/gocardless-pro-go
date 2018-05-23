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
	Amount    int    `url:",omitempty" json:"amount,omitempty"`
	CreatedAt string `url:",omitempty" json:"created_at,omitempty"`
	Currency  string `url:",omitempty" json:"currency,omitempty"`
	Id        string `url:",omitempty" json:"id,omitempty"`
	Links     struct {
		Mandate string `url:",omitempty" json:"mandate,omitempty"`
		Payment string `url:",omitempty" json:"payment,omitempty"`
	} `url:",omitempty" json:"links,omitempty"`
	Metadata  map[string]interface{} `url:",omitempty" json:"metadata,omitempty"`
	Reference string                 `url:",omitempty" json:"reference,omitempty"`
}

// RefundCreateParams parameters
type RefundCreateParams struct {
	Amount int `url:",omitempty" json:"amount,omitempty"`
	Links  struct {
		Payment string `url:",omitempty" json:"payment,omitempty"`
	} `url:",omitempty" json:"links,omitempty"`
	Metadata                map[string]interface{} `url:",omitempty" json:"metadata,omitempty"`
	Reference               string                 `url:",omitempty" json:"reference,omitempty"`
	TotalAmountConfirmation int                    `url:",omitempty" json:"total_amount_confirmation,omitempty"`
}

// Create
// Creates a new refund object.
//
// This fails with:<a name="refund_payment_invalid_state"></a><a
// name="total_amount_confirmation_invalid"></a><a
// name="number_of_refunds_exceeded"></a>
//
// - `refund_payment_invalid_state` error if the linked
// [payment](#core-endpoints-payments) isn't either `confirmed` or `paid_out`.
//
// - `total_amount_confirmation_invalid` if the confirmation amount doesn't
// match the total amount refunded for the payment. This safeguard is there to
// prevent two processes from creating refunds without awareness of each other.
//
// - `number_of_refunds_exceeded` if five or more refunds have already been
// created against the payment.
//
func (s *RefundService) Create(ctx context.Context, p RefundCreateParams) (*Refund, error) {
	uri, err := url.Parse(fmt.Sprintf(s.endpoint + "/refunds"))
	if err != nil {
		return nil, err
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
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Idempotency-Key", NewIdempotencyKey())

	client := s.client
	if client == nil {
		client = http.DefaultClient
	}

	var result struct {
		Err    *APIError `json:"error"`
		Refund *Refund   `json:"refunds"`
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

	if result.Refund == nil {
		return nil, errors.New("missing result")
	}

	return result.Refund, nil
}

// RefundListParams parameters
type RefundListParams struct {
	After     string `url:",omitempty" json:"after,omitempty"`
	Before    string `url:",omitempty" json:"before,omitempty"`
	CreatedAt struct {
		Gt  string `url:",omitempty" json:"gt,omitempty"`
		Gte string `url:",omitempty" json:"gte,omitempty"`
		Lt  string `url:",omitempty" json:"lt,omitempty"`
		Lte string `url:",omitempty" json:"lte,omitempty"`
	} `url:",omitempty" json:"created_at,omitempty"`
	Limit   int    `url:",omitempty" json:"limit,omitempty"`
	Payment string `url:",omitempty" json:"payment,omitempty"`
}

// RefundListResult response including pagination metadata
type RefundListResult struct {
	Refunds []Refund `json:"refunds"`
	Meta    struct {
		Cursors struct {
			After  string `url:",omitempty" json:"after,omitempty"`
			Before string `url:",omitempty" json:"before,omitempty"`
		} `url:",omitempty" json:"cursors,omitempty"`
		Limit int `url:",omitempty" json:"limit,omitempty"`
	} `json:"meta"`
}

// List
// Returns a [cursor-paginated](#api-usage-cursor-pagination) list of your
// refunds.
func (s *RefundService) List(ctx context.Context, p RefundListParams) (*RefundListResult, error) {
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

	return result.RefundListResult, nil
}

// Get
// Retrieves all details for a single refund
func (s *RefundService) Get(ctx context.Context, identity string) (*Refund, error) {
	uri, err := url.Parse(fmt.Sprintf(s.endpoint+"/refunds/%v",
		identity))
	if err != nil {
		return nil, err
	}

	var body io.Reader

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
		Err    *APIError `json:"error"`
		Refund *Refund   `json:"refunds"`
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

	if result.Refund == nil {
		return nil, errors.New("missing result")
	}

	return result.Refund, nil
}

// RefundUpdateParams parameters
type RefundUpdateParams struct {
	Metadata map[string]interface{} `url:",omitempty" json:"metadata,omitempty"`
}

// Update
// Updates a refund object.
func (s *RefundService) Update(ctx context.Context, identity string, p RefundUpdateParams) (*Refund, error) {
	uri, err := url.Parse(fmt.Sprintf(s.endpoint+"/refunds/%v",
		identity))
	if err != nil {
		return nil, err
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
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Idempotency-Key", NewIdempotencyKey())

	client := s.client
	if client == nil {
		client = http.DefaultClient
	}

	var result struct {
		Err    *APIError `json:"error"`
		Refund *Refund   `json:"refunds"`
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

	if result.Refund == nil {
		return nil, errors.New("missing result")
	}

	return result.Refund, nil
}
