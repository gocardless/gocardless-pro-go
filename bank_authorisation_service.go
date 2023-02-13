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

// BankAuthorisationService manages bank_authorisations
type BankAuthorisationServiceImpl struct {
	config Config
}

type BankAuthorisationLinks struct {
	BillingRequest string `url:"billing_request,omitempty" json:"billing_request,omitempty"`
	Institution    string `url:"institution,omitempty" json:"institution,omitempty"`
}

// BankAuthorisation model
type BankAuthorisation struct {
	AuthorisationType string                  `url:"authorisation_type,omitempty" json:"authorisation_type,omitempty"`
	AuthorisedAt      string                  `url:"authorised_at,omitempty" json:"authorised_at,omitempty"`
	CreatedAt         string                  `url:"created_at,omitempty" json:"created_at,omitempty"`
	ExpiresAt         string                  `url:"expires_at,omitempty" json:"expires_at,omitempty"`
	Id                string                  `url:"id,omitempty" json:"id,omitempty"`
	LastVisitedAt     string                  `url:"last_visited_at,omitempty" json:"last_visited_at,omitempty"`
	Links             *BankAuthorisationLinks `url:"links,omitempty" json:"links,omitempty"`
	RedirectUri       string                  `url:"redirect_uri,omitempty" json:"redirect_uri,omitempty"`
	Url               string                  `url:"url,omitempty" json:"url,omitempty"`
}

type BankAuthorisationService interface {
	Create(ctx context.Context, p BankAuthorisationCreateParams, opts ...RequestOption) (*BankAuthorisation, error)
	Get(ctx context.Context, identity string, opts ...RequestOption) (*BankAuthorisation, error)
}

type BankAuthorisationCreateParamsLinks struct {
	BillingRequest string `url:"billing_request,omitempty" json:"billing_request,omitempty"`
	Institution    string `url:"institution,omitempty" json:"institution,omitempty"`
}

// BankAuthorisationCreateParams parameters
type BankAuthorisationCreateParams struct {
	AuthorisationType string                             `url:"authorisation_type,omitempty" json:"authorisation_type,omitempty"`
	Links             BankAuthorisationCreateParamsLinks `url:"links,omitempty" json:"links,omitempty"`
	RedirectUri       string                             `url:"redirect_uri,omitempty" json:"redirect_uri,omitempty"`
}

// Create
// Create a Bank Authorisation.
func (s *BankAuthorisationServiceImpl) Create(ctx context.Context, p BankAuthorisationCreateParams, opts ...RequestOption) (*BankAuthorisation, error) {
	uri, err := url.Parse(fmt.Sprintf(s.config.Endpoint() + "/bank_authorisations"))
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
		"bank_authorisations": p,
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
	req.Header.Set("GoCardless-Client-Version", "2.10.0")
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
		Err               *APIError          `json:"error"`
		BankAuthorisation *BankAuthorisation `json:"bank_authorisations"`
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

	if result.BankAuthorisation == nil {
		return nil, errors.New("missing result")
	}

	return result.BankAuthorisation, nil
}

// Get
// Get a single bank authorisation.
func (s *BankAuthorisationServiceImpl) Get(ctx context.Context, identity string, opts ...RequestOption) (*BankAuthorisation, error) {
	uri, err := url.Parse(fmt.Sprintf(s.config.Endpoint()+"/bank_authorisations/%v",
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
	req.Header.Set("GoCardless-Client-Version", "2.10.0")
	req.Header.Set("User-Agent", userAgent)

	for key, value := range o.headers {
		req.Header.Set(key, value)
	}

	client := s.config.Client()
	if client == nil {
		client = http.DefaultClient
	}

	var result struct {
		Err               *APIError          `json:"error"`
		BankAuthorisation *BankAuthorisation `json:"bank_authorisations"`
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

	if result.BankAuthorisation == nil {
		return nil, errors.New("missing result")
	}

	return result.BankAuthorisation, nil
}
