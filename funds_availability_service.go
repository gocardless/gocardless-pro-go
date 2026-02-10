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

// FundsAvailabilityService manages funds_availabilities
type FundsAvailabilityServiceImpl struct {
	config Config
}

// FundsAvailability model
type FundsAvailability struct {
	Available bool `url:"available,omitempty" json:"available,omitempty"`
}

type FundsAvailabilityService interface {
	Check(ctx context.Context, identity string, p FundsAvailabilityCheckParams, opts ...RequestOption) (*FundsAvailability, error)
}

// FundsAvailabilityCheckParams parameters
type FundsAvailabilityCheckParams struct {
	Amount string `url:"amount,omitempty" json:"amount,omitempty"`
}

// Check
//
//	Checks if the payer's current balance is sufficient to cover the amount
//	the merchant wants to charge within the consent parameters defined on the
//	mandate.
func (s *FundsAvailabilityServiceImpl) Check(ctx context.Context, identity string, p FundsAvailabilityCheckParams, opts ...RequestOption) (*FundsAvailability, error) {
	uri, err := url.Parse(fmt.Sprintf(s.config.Endpoint()+"/funds_availability/%v",
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
	req.Header.Set("GoCardless-Client-Version", "5.3.0")
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
		*FundsAvailability
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

	if result.FundsAvailability == nil {
		return nil, errors.New("missing result")
	}

	return result.FundsAvailability, nil
}
