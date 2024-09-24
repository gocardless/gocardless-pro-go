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

// InstitutionService manages institutions
type InstitutionServiceImpl struct {
	config Config
}

// Institution model
type Institution struct {
	AutocompletesCollectBankAccount bool   `url:"autocompletes_collect_bank_account,omitempty" json:"autocompletes_collect_bank_account,omitempty"`
	CountryCode                     string `url:"country_code,omitempty" json:"country_code,omitempty"`
	IconUrl                         string `url:"icon_url,omitempty" json:"icon_url,omitempty"`
	Id                              string `url:"id,omitempty" json:"id,omitempty"`
	LogoUrl                         string `url:"logo_url,omitempty" json:"logo_url,omitempty"`
	Name                            string `url:"name,omitempty" json:"name,omitempty"`
}

type InstitutionService interface {
	List(ctx context.Context, p InstitutionListParams, opts ...RequestOption) (*InstitutionListResult, error)
	ListForBillingRequest(ctx context.Context, identity string, p InstitutionListForBillingRequestParams, opts ...RequestOption) (
		*InstitutionListForBillingRequestResult, error)
}

// InstitutionListParams parameters
type InstitutionListParams struct {
	CountryCode string `url:"country_code,omitempty" json:"country_code,omitempty"`
}

type InstitutionListResultMetaCursors struct {
	After  string `url:"after,omitempty" json:"after,omitempty"`
	Before string `url:"before,omitempty" json:"before,omitempty"`
}

type InstitutionListResultMeta struct {
	Cursors *InstitutionListResultMetaCursors `url:"cursors,omitempty" json:"cursors,omitempty"`
	Limit   int                               `url:"limit,omitempty" json:"limit,omitempty"`
}

type InstitutionListResult struct {
	Institutions []Institution             `json:"institutions"`
	Meta         InstitutionListResultMeta `url:"meta,omitempty" json:"meta,omitempty"`
}

// List
// Returns a list of supported institutions.
func (s *InstitutionServiceImpl) List(ctx context.Context, p InstitutionListParams, opts ...RequestOption) (*InstitutionListResult, error) {
	uri, err := url.Parse(fmt.Sprintf(s.config.Endpoint() + "/institutions"))
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
	req.Header.Set("GoCardless-Client-Version", "3.7.0")
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
		*InstitutionListResult
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

	if result.InstitutionListResult == nil {
		return nil, errors.New("missing result")
	}

	return result.InstitutionListResult, nil
}

// InstitutionListForBillingRequestParams parameters
type InstitutionListForBillingRequestParams struct {
	CountryCode string   `url:"country_code,omitempty" json:"country_code,omitempty"`
	Ids         []string `url:"ids,omitempty" json:"ids,omitempty"`
	Search      string   `url:"search,omitempty" json:"search,omitempty"`
}

type InstitutionListForBillingRequestResultMetaCursors struct {
	After  string `url:"after,omitempty" json:"after,omitempty"`
	Before string `url:"before,omitempty" json:"before,omitempty"`
}

type InstitutionListForBillingRequestResultMeta struct {
	Cursors *InstitutionListForBillingRequestResultMetaCursors `url:"cursors,omitempty" json:"cursors,omitempty"`
	Limit   int                                                `url:"limit,omitempty" json:"limit,omitempty"`
}

type InstitutionListForBillingRequestResult struct {
	Institutions []Institution                              `json:"institutions"`
	Meta         InstitutionListForBillingRequestResultMeta `url:"meta,omitempty" json:"meta,omitempty"`
}

// ListForBillingRequest
// Returns all institutions valid for a Billing Request.
//
// This endpoint is currently supported only for FasterPayments.
func (s *InstitutionServiceImpl) ListForBillingRequest(ctx context.Context, identity string, p InstitutionListForBillingRequestParams, opts ...RequestOption) (
	*InstitutionListForBillingRequestResult, error) {
	uri, err := url.Parse(fmt.Sprintf(s.config.Endpoint()+"/billing_requests/%v/institutions",
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
	req.Header.Set("GoCardless-Client-Version", "3.7.0")
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

		*InstitutionListForBillingRequestResult
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

	if result.InstitutionListForBillingRequestResult == nil {
		return nil, errors.New("missing result")
	}

	return result.InstitutionListForBillingRequestResult, nil
}
