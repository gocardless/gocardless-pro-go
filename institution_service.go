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
type InstitutionService struct {
	endpoint string
	token    string
	client   *http.Client
}

// Institution model
type Institution struct {
	CountryCode string `url:"country_code,omitempty" json:"country_code,omitempty"`
	IconUrl     string `url:"icon_url,omitempty" json:"icon_url,omitempty"`
	Id          string `url:"id,omitempty" json:"id,omitempty"`
	LogoUrl     string `url:"logo_url,omitempty" json:"logo_url,omitempty"`
	Name        string `url:"name,omitempty" json:"name,omitempty"`
}

// InstitutionListParams parameters
type InstitutionListParams struct {
	CountryCode string `url:"country_code,omitempty" json:"country_code,omitempty"`
}

// InstitutionListResult response including pagination metadata
type InstitutionListResult struct {
	Institutions []Institution `json:"institutions"`
	Meta         struct {
		Cursors struct {
			After  string `url:"after,omitempty" json:"after,omitempty"`
			Before string `url:"before,omitempty" json:"before,omitempty"`
		} `url:"cursors,omitempty" json:"cursors,omitempty"`
		Limit int `url:"limit,omitempty" json:"limit,omitempty"`
	} `json:"meta"`
}

// List
// Returns a list of all supported institutions.
func (s *InstitutionService) List(ctx context.Context, p InstitutionListParams, opts ...RequestOption) (*InstitutionListResult, error) {
	uri, err := url.Parse(fmt.Sprintf(s.endpoint + "/institutions"))
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
