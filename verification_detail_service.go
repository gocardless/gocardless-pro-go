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

// VerificationDetailService manages verification_details
type VerificationDetailServiceImpl struct {
	config Config
}

type VerificationDetailDirectors struct {
	City        string `url:"city,omitempty" json:"city,omitempty"`
	CountryCode string `url:"country_code,omitempty" json:"country_code,omitempty"`
	DateOfBirth string `url:"date_of_birth,omitempty" json:"date_of_birth,omitempty"`
	FamilyName  string `url:"family_name,omitempty" json:"family_name,omitempty"`
	GivenName   string `url:"given_name,omitempty" json:"given_name,omitempty"`
	PostalCode  string `url:"postal_code,omitempty" json:"postal_code,omitempty"`
	Street      string `url:"street,omitempty" json:"street,omitempty"`
}

type VerificationDetailLinks struct {
	Creditor string `url:"creditor,omitempty" json:"creditor,omitempty"`
}

// VerificationDetail model
type VerificationDetail struct {
	AddressLine1  string                        `url:"address_line1,omitempty" json:"address_line1,omitempty"`
	AddressLine2  string                        `url:"address_line2,omitempty" json:"address_line2,omitempty"`
	AddressLine3  string                        `url:"address_line3,omitempty" json:"address_line3,omitempty"`
	City          string                        `url:"city,omitempty" json:"city,omitempty"`
	CompanyNumber string                        `url:"company_number,omitempty" json:"company_number,omitempty"`
	Description   string                        `url:"description,omitempty" json:"description,omitempty"`
	Directors     []VerificationDetailDirectors `url:"directors,omitempty" json:"directors,omitempty"`
	Links         *VerificationDetailLinks      `url:"links,omitempty" json:"links,omitempty"`
	PostalCode    string                        `url:"postal_code,omitempty" json:"postal_code,omitempty"`
}

type VerificationDetailService interface {
	Create(ctx context.Context, identity string, p VerificationDetailCreateParams, opts ...RequestOption) (*VerificationDetail, error)
}

type VerificationDetailCreateParamsDirectors struct {
	City        string `url:"city,omitempty" json:"city,omitempty"`
	CountryCode string `url:"country_code,omitempty" json:"country_code,omitempty"`
	DateOfBirth string `url:"date_of_birth,omitempty" json:"date_of_birth,omitempty"`
	FamilyName  string `url:"family_name,omitempty" json:"family_name,omitempty"`
	GivenName   string `url:"given_name,omitempty" json:"given_name,omitempty"`
	PostalCode  string `url:"postal_code,omitempty" json:"postal_code,omitempty"`
	Street      string `url:"street,omitempty" json:"street,omitempty"`
}

type VerificationDetailCreateParamsLinks struct {
	Creditor string `url:"creditor,omitempty" json:"creditor,omitempty"`
}

// VerificationDetailCreateParams parameters
type VerificationDetailCreateParams struct {
	AddressLine1  string                                    `url:"address_line1,omitempty" json:"address_line1,omitempty"`
	AddressLine2  string                                    `url:"address_line2,omitempty" json:"address_line2,omitempty"`
	AddressLine3  string                                    `url:"address_line3,omitempty" json:"address_line3,omitempty"`
	City          string                                    `url:"city,omitempty" json:"city,omitempty"`
	CompanyNumber string                                    `url:"company_number,omitempty" json:"company_number,omitempty"`
	Description   string                                    `url:"description,omitempty" json:"description,omitempty"`
	Directors     []VerificationDetailCreateParamsDirectors `url:"directors,omitempty" json:"directors,omitempty"`
	Links         VerificationDetailCreateParamsLinks       `url:"links,omitempty" json:"links,omitempty"`
	PostalCode    string                                    `url:"postal_code,omitempty" json:"postal_code,omitempty"`
}

// Create
// Verification details represent any information needed by GoCardless to verify
// a creditor.
// Currently, only UK-based companies are supported.
// In other words, to submit verification details for a creditor, their
// creditor_type must be company and their country_code must be GB.
func (s *VerificationDetailServiceImpl) Create(ctx context.Context, identity string, p VerificationDetailCreateParams, opts ...RequestOption) (*VerificationDetail, error) {
	uri, err := url.Parse(fmt.Sprintf(s.config.Endpoint()+"/verification_details/%v",
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
		"verification_details": p,
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
	req.Header.Set("GoCardless-Client-Version", "2.9.0")
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
		Err                *APIError           `json:"error"`
		VerificationDetail *VerificationDetail `json:"verification_details"`
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

	if result.VerificationDetail == nil {
		return nil, errors.New("missing result")
	}

	return result.VerificationDetail, nil
}
