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
type CustomerService struct {
	endpoint string
	token    string
	client   *http.Client
}

// Customer model
type Customer struct {
	AddressLine1          string                 `url:",omitempty" json:"address_line1,omitempty"`
	AddressLine2          string                 `url:",omitempty" json:"address_line2,omitempty"`
	AddressLine3          string                 `url:",omitempty" json:"address_line3,omitempty"`
	City                  string                 `url:",omitempty" json:"city,omitempty"`
	CompanyName           string                 `url:",omitempty" json:"company_name,omitempty"`
	CountryCode           string                 `url:",omitempty" json:"country_code,omitempty"`
	CreatedAt             string                 `url:",omitempty" json:"created_at,omitempty"`
	DanishIdentityNumber  string                 `url:",omitempty" json:"danish_identity_number,omitempty"`
	Email                 string                 `url:",omitempty" json:"email,omitempty"`
	FamilyName            string                 `url:",omitempty" json:"family_name,omitempty"`
	GivenName             string                 `url:",omitempty" json:"given_name,omitempty"`
	Id                    string                 `url:",omitempty" json:"id,omitempty"`
	Language              string                 `url:",omitempty" json:"language,omitempty"`
	Metadata              map[string]interface{} `url:",omitempty" json:"metadata,omitempty"`
	PostalCode            string                 `url:",omitempty" json:"postal_code,omitempty"`
	Region                string                 `url:",omitempty" json:"region,omitempty"`
	SwedishIdentityNumber string                 `url:",omitempty" json:"swedish_identity_number,omitempty"`
}

// CustomerCreateParams parameters
type CustomerCreateParams struct {
	AddressLine1          string                 `url:",omitempty" json:"address_line1,omitempty"`
	AddressLine2          string                 `url:",omitempty" json:"address_line2,omitempty"`
	AddressLine3          string                 `url:",omitempty" json:"address_line3,omitempty"`
	City                  string                 `url:",omitempty" json:"city,omitempty"`
	CompanyName           string                 `url:",omitempty" json:"company_name,omitempty"`
	CountryCode           string                 `url:",omitempty" json:"country_code,omitempty"`
	DanishIdentityNumber  string                 `url:",omitempty" json:"danish_identity_number,omitempty"`
	Email                 string                 `url:",omitempty" json:"email,omitempty"`
	FamilyName            string                 `url:",omitempty" json:"family_name,omitempty"`
	GivenName             string                 `url:",omitempty" json:"given_name,omitempty"`
	Language              string                 `url:",omitempty" json:"language,omitempty"`
	Metadata              map[string]interface{} `url:",omitempty" json:"metadata,omitempty"`
	PostalCode            string                 `url:",omitempty" json:"postal_code,omitempty"`
	Region                string                 `url:",omitempty" json:"region,omitempty"`
	SwedishIdentityNumber string                 `url:",omitempty" json:"swedish_identity_number,omitempty"`
}

// Create
// Creates a new customer object.
func (s *CustomerService) Create(ctx context.Context, p CustomerCreateParams) (*Customer, error) {
	uri, err := url.Parse(fmt.Sprintf(s.endpoint + "/customers"))
	if err != nil {
		return nil, err
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
		Err      *APIError `json:"error"`
		Customer *Customer `json:"customers"`
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

	if result.Customer == nil {
		return nil, errors.New("missing result")
	}

	return result.Customer, nil
}

// CustomerListParams parameters
type CustomerListParams struct {
	After     string `url:",omitempty" json:"after,omitempty"`
	Before    string `url:",omitempty" json:"before,omitempty"`
	CreatedAt struct {
		Gt  string `url:",omitempty" json:"gt,omitempty"`
		Gte string `url:",omitempty" json:"gte,omitempty"`
		Lt  string `url:",omitempty" json:"lt,omitempty"`
		Lte string `url:",omitempty" json:"lte,omitempty"`
	} `url:",omitempty" json:"created_at,omitempty"`
	Limit int `url:",omitempty" json:"limit,omitempty"`
}

// CustomerListResult response including pagination metadata
type CustomerListResult struct {
	Customers []Customer `json:"customers"`
	Meta      struct {
		Cursors struct {
			After  string `url:",omitempty" json:"after,omitempty"`
			Before string `url:",omitempty" json:"before,omitempty"`
		} `url:",omitempty" json:"cursors,omitempty"`
		Limit int `url:",omitempty" json:"limit,omitempty"`
	} `json:"meta"`
}

// List
// Returns a [cursor-paginated](#api-usage-cursor-pagination) list of your
// customers.
func (s *CustomerService) List(ctx context.Context, p CustomerListParams) (*CustomerListResult, error) {
	uri, err := url.Parse(fmt.Sprintf(s.endpoint + "/customers"))
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
		*CustomerListResult
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

	if result.CustomerListResult == nil {
		return nil, errors.New("missing result")
	}

	return result.CustomerListResult, nil
}

// Get
// Retrieves the details of an existing customer.
func (s *CustomerService) Get(ctx context.Context, identity string) (*Customer, error) {
	uri, err := url.Parse(fmt.Sprintf(s.endpoint+"/customers/%v",
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
		Err      *APIError `json:"error"`
		Customer *Customer `json:"customers"`
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

	if result.Customer == nil {
		return nil, errors.New("missing result")
	}

	return result.Customer, nil
}

// CustomerUpdateParams parameters
type CustomerUpdateParams struct {
	AddressLine1          string                 `url:",omitempty" json:"address_line1,omitempty"`
	AddressLine2          string                 `url:",omitempty" json:"address_line2,omitempty"`
	AddressLine3          string                 `url:",omitempty" json:"address_line3,omitempty"`
	City                  string                 `url:",omitempty" json:"city,omitempty"`
	CompanyName           string                 `url:",omitempty" json:"company_name,omitempty"`
	CountryCode           string                 `url:",omitempty" json:"country_code,omitempty"`
	DanishIdentityNumber  string                 `url:",omitempty" json:"danish_identity_number,omitempty"`
	Email                 string                 `url:",omitempty" json:"email,omitempty"`
	FamilyName            string                 `url:",omitempty" json:"family_name,omitempty"`
	GivenName             string                 `url:",omitempty" json:"given_name,omitempty"`
	Language              string                 `url:",omitempty" json:"language,omitempty"`
	Metadata              map[string]interface{} `url:",omitempty" json:"metadata,omitempty"`
	PostalCode            string                 `url:",omitempty" json:"postal_code,omitempty"`
	Region                string                 `url:",omitempty" json:"region,omitempty"`
	SwedishIdentityNumber string                 `url:",omitempty" json:"swedish_identity_number,omitempty"`
}

// Update
// Updates a customer object. Supports all of the fields supported when creating
// a customer.
func (s *CustomerService) Update(ctx context.Context, identity string, p CustomerUpdateParams) (*Customer, error) {
	uri, err := url.Parse(fmt.Sprintf(s.endpoint+"/customers/%v",
		identity))
	if err != nil {
		return nil, err
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
		Err      *APIError `json:"error"`
		Customer *Customer `json:"customers"`
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

	if result.Customer == nil {
		return nil, errors.New("missing result")
	}

	return result.Customer, nil
}
