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

// MandateImportService manages mandate_imports
type MandateImportServiceImpl struct {
	config Config
}

// MandateImport model
type MandateImport struct {
	CreatedAt string `url:"created_at,omitempty" json:"created_at,omitempty"`
	Id        string `url:"id,omitempty" json:"id,omitempty"`
	Scheme    string `url:"scheme,omitempty" json:"scheme,omitempty"`
	Status    string `url:"status,omitempty" json:"status,omitempty"`
}

type MandateImportService interface {
	Create(ctx context.Context, p MandateImportCreateParams, opts ...RequestOption) (*MandateImport, error)
	Get(ctx context.Context, identity string, p MandateImportGetParams, opts ...RequestOption) (*MandateImport, error)
	Submit(ctx context.Context, identity string, p MandateImportSubmitParams, opts ...RequestOption) (*MandateImport, error)
	Cancel(ctx context.Context, identity string, p MandateImportCancelParams, opts ...RequestOption) (*MandateImport, error)
}

// MandateImportCreateParams parameters
type MandateImportCreateParams struct {
	Scheme string `url:"scheme,omitempty" json:"scheme,omitempty"`
}

// Create
// Mandate imports are first created, before mandates are added one-at-a-time,
// so
// this endpoint merely signals the start of the import process. Once you've
// finished
// adding entries to an import, you should
// [submit](#mandate-imports-submit-a-mandate-import) it.
func (s *MandateImportServiceImpl) Create(ctx context.Context, p MandateImportCreateParams, opts ...RequestOption) (*MandateImport, error) {
	uri, err := url.Parse(fmt.Sprintf(s.config.Endpoint() + "/mandate_imports"))
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
		"mandate_imports": p,
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
	req.Header.Set("GoCardless-Client-Version", "3.1.0")
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
		Err           *APIError      `json:"error"`
		MandateImport *MandateImport `json:"mandate_imports"`
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

	if result.MandateImport == nil {
		return nil, errors.New("missing result")
	}

	return result.MandateImport, nil
}

// MandateImportGetParams parameters
type MandateImportGetParams struct {
}

// Get
// Returns a single mandate import.
func (s *MandateImportServiceImpl) Get(ctx context.Context, identity string, p MandateImportGetParams, opts ...RequestOption) (*MandateImport, error) {
	uri, err := url.Parse(fmt.Sprintf(s.config.Endpoint()+"/mandate_imports/%v",
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
	req.Header.Set("GoCardless-Client-Version", "3.1.0")
	req.Header.Set("User-Agent", userAgent)

	for key, value := range o.headers {
		req.Header.Set(key, value)
	}

	client := s.config.Client()
	if client == nil {
		client = http.DefaultClient
	}

	var result struct {
		Err           *APIError      `json:"error"`
		MandateImport *MandateImport `json:"mandate_imports"`
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

	if result.MandateImport == nil {
		return nil, errors.New("missing result")
	}

	return result.MandateImport, nil
}

// MandateImportSubmitParams parameters
type MandateImportSubmitParams struct {
}

// Submit
// Submits the mandate import, which allows it to be processed by a member of
// the
// GoCardless team. Once the import has been submitted, it can no longer have
// entries
// added to it.
//
// In our sandbox environment, to aid development, we automatically process
// mandate
// imports approximately 10 seconds after they are submitted. This will allow
// you to
// test both the "submitted" response and wait for the webhook to confirm the
// processing has begun.
func (s *MandateImportServiceImpl) Submit(ctx context.Context, identity string, p MandateImportSubmitParams, opts ...RequestOption) (*MandateImport, error) {
	uri, err := url.Parse(fmt.Sprintf(s.config.Endpoint()+"/mandate_imports/%v/actions/submit",
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

	req, err := http.NewRequest("POST", uri.String(), body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	req.Header.Set("Authorization", "Bearer "+s.config.Token())
	req.Header.Set("GoCardless-Version", "2015-07-06")
	req.Header.Set("GoCardless-Client-Library", "gocardless-pro-go")
	req.Header.Set("GoCardless-Client-Version", "3.1.0")
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
		Err           *APIError      `json:"error"`
		MandateImport *MandateImport `json:"mandate_imports"`
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

	if result.MandateImport == nil {
		return nil, errors.New("missing result")
	}

	return result.MandateImport, nil
}

// MandateImportCancelParams parameters
type MandateImportCancelParams struct {
}

// Cancel
// Cancels the mandate import, which aborts the import process and stops the
// mandates
// being set up in GoCardless. Once the import has been cancelled, it can no
// longer have
// entries added to it. Mandate imports which have already been submitted or
// processed
// cannot be cancelled.
func (s *MandateImportServiceImpl) Cancel(ctx context.Context, identity string, p MandateImportCancelParams, opts ...RequestOption) (*MandateImport, error) {
	uri, err := url.Parse(fmt.Sprintf(s.config.Endpoint()+"/mandate_imports/%v/actions/cancel",
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

	req, err := http.NewRequest("POST", uri.String(), body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	req.Header.Set("Authorization", "Bearer "+s.config.Token())
	req.Header.Set("GoCardless-Version", "2015-07-06")
	req.Header.Set("GoCardless-Client-Library", "gocardless-pro-go")
	req.Header.Set("GoCardless-Client-Version", "3.1.0")
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
		Err           *APIError      `json:"error"`
		MandateImport *MandateImport `json:"mandate_imports"`
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

	if result.MandateImport == nil {
		return nil, errors.New("missing result")
	}

	return result.MandateImport, nil
}
