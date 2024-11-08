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

// MandateImportEntryService manages mandate_import_entries
type MandateImportEntryServiceImpl struct {
	config Config
}

type MandateImportEntryLinks struct {
	Customer            string `url:"customer,omitempty" json:"customer,omitempty"`
	CustomerBankAccount string `url:"customer_bank_account,omitempty" json:"customer_bank_account,omitempty"`
	Mandate             string `url:"mandate,omitempty" json:"mandate,omitempty"`
	MandateImport       string `url:"mandate_import,omitempty" json:"mandate_import,omitempty"`
}

// MandateImportEntry model
type MandateImportEntry struct {
	CreatedAt        string                   `url:"created_at,omitempty" json:"created_at,omitempty"`
	Links            *MandateImportEntryLinks `url:"links,omitempty" json:"links,omitempty"`
	RecordIdentifier string                   `url:"record_identifier,omitempty" json:"record_identifier,omitempty"`
}

type MandateImportEntryService interface {
	Create(ctx context.Context, p MandateImportEntryCreateParams, opts ...RequestOption) (*MandateImportEntry, error)
	List(ctx context.Context, p MandateImportEntryListParams, opts ...RequestOption) (*MandateImportEntryListResult, error)
	All(ctx context.Context, p MandateImportEntryListParams, opts ...RequestOption) *MandateImportEntryListPagingIterator
}

type MandateImportEntryCreateParamsAmendment struct {
	OriginalCreditorId       string `url:"original_creditor_id,omitempty" json:"original_creditor_id,omitempty"`
	OriginalCreditorName     string `url:"original_creditor_name,omitempty" json:"original_creditor_name,omitempty"`
	OriginalMandateReference string `url:"original_mandate_reference,omitempty" json:"original_mandate_reference,omitempty"`
}

type MandateImportEntryCreateParamsBankAccount struct {
	AccountHolderName string                 `url:"account_holder_name,omitempty" json:"account_holder_name,omitempty"`
	AccountNumber     string                 `url:"account_number,omitempty" json:"account_number,omitempty"`
	AccountType       string                 `url:"account_type,omitempty" json:"account_type,omitempty"`
	BankCode          string                 `url:"bank_code,omitempty" json:"bank_code,omitempty"`
	BranchCode        string                 `url:"branch_code,omitempty" json:"branch_code,omitempty"`
	CountryCode       string                 `url:"country_code,omitempty" json:"country_code,omitempty"`
	Iban              string                 `url:"iban,omitempty" json:"iban,omitempty"`
	Metadata          map[string]interface{} `url:"metadata,omitempty" json:"metadata,omitempty"`
}

type MandateImportEntryCreateParamsCustomer struct {
	AddressLine1          string                 `url:"address_line1,omitempty" json:"address_line1,omitempty"`
	AddressLine2          string                 `url:"address_line2,omitempty" json:"address_line2,omitempty"`
	AddressLine3          string                 `url:"address_line3,omitempty" json:"address_line3,omitempty"`
	City                  string                 `url:"city,omitempty" json:"city,omitempty"`
	CompanyName           string                 `url:"company_name,omitempty" json:"company_name,omitempty"`
	CountryCode           string                 `url:"country_code,omitempty" json:"country_code,omitempty"`
	DanishIdentityNumber  string                 `url:"danish_identity_number,omitempty" json:"danish_identity_number,omitempty"`
	Email                 string                 `url:"email,omitempty" json:"email,omitempty"`
	FamilyName            string                 `url:"family_name,omitempty" json:"family_name,omitempty"`
	GivenName             string                 `url:"given_name,omitempty" json:"given_name,omitempty"`
	Language              string                 `url:"language,omitempty" json:"language,omitempty"`
	Metadata              map[string]interface{} `url:"metadata,omitempty" json:"metadata,omitempty"`
	PhoneNumber           string                 `url:"phone_number,omitempty" json:"phone_number,omitempty"`
	PostalCode            string                 `url:"postal_code,omitempty" json:"postal_code,omitempty"`
	Region                string                 `url:"region,omitempty" json:"region,omitempty"`
	SwedishIdentityNumber string                 `url:"swedish_identity_number,omitempty" json:"swedish_identity_number,omitempty"`
}

type MandateImportEntryCreateParamsLinks struct {
	MandateImport string `url:"mandate_import,omitempty" json:"mandate_import,omitempty"`
}

type MandateImportEntryCreateParamsMandate struct {
	Metadata  map[string]interface{} `url:"metadata,omitempty" json:"metadata,omitempty"`
	Reference string                 `url:"reference,omitempty" json:"reference,omitempty"`
}

// MandateImportEntryCreateParams parameters
type MandateImportEntryCreateParams struct {
	Amendment        *MandateImportEntryCreateParamsAmendment  `url:"amendment,omitempty" json:"amendment,omitempty"`
	BankAccount      MandateImportEntryCreateParamsBankAccount `url:"bank_account,omitempty" json:"bank_account,omitempty"`
	Customer         MandateImportEntryCreateParamsCustomer    `url:"customer,omitempty" json:"customer,omitempty"`
	Links            MandateImportEntryCreateParamsLinks       `url:"links,omitempty" json:"links,omitempty"`
	Mandate          *MandateImportEntryCreateParamsMandate    `url:"mandate,omitempty" json:"mandate,omitempty"`
	RecordIdentifier string                                    `url:"record_identifier,omitempty" json:"record_identifier,omitempty"`
}

// Create
// For an existing [mandate import](#core-endpoints-mandate-imports), this
// endpoint can
// be used to add individual mandates to be imported into GoCardless.
//
// You can add no more than 30,000 rows to a single mandate import.
// If you attempt to go over this limit, the API will return a
// `record_limit_exceeded` error.
func (s *MandateImportEntryServiceImpl) Create(ctx context.Context, p MandateImportEntryCreateParams, opts ...RequestOption) (*MandateImportEntry, error) {
	uri, err := url.Parse(fmt.Sprintf(s.config.Endpoint() + "/mandate_import_entries"))
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
		"mandate_import_entries": p,
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
	req.Header.Set("GoCardless-Client-Version", "4.1.0")
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
		MandateImportEntry *MandateImportEntry `json:"mandate_import_entries"`
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

	if result.MandateImportEntry == nil {
		return nil, errors.New("missing result")
	}

	return result.MandateImportEntry, nil
}

// MandateImportEntryListParams parameters
type MandateImportEntryListParams struct {
	After         string `url:"after,omitempty" json:"after,omitempty"`
	Before        string `url:"before,omitempty" json:"before,omitempty"`
	Limit         int    `url:"limit,omitempty" json:"limit,omitempty"`
	MandateImport string `url:"mandate_import,omitempty" json:"mandate_import,omitempty"`
}

type MandateImportEntryListResultMetaCursors struct {
	After  string `url:"after,omitempty" json:"after,omitempty"`
	Before string `url:"before,omitempty" json:"before,omitempty"`
}

type MandateImportEntryListResultMeta struct {
	Cursors *MandateImportEntryListResultMetaCursors `url:"cursors,omitempty" json:"cursors,omitempty"`
	Limit   int                                      `url:"limit,omitempty" json:"limit,omitempty"`
}

type MandateImportEntryListResult struct {
	MandateImportEntries []MandateImportEntry             `json:"mandate_import_entries"`
	Meta                 MandateImportEntryListResultMeta `url:"meta,omitempty" json:"meta,omitempty"`
}

// List
// For an existing mandate import, this endpoint lists all of the entries
// attached.
//
// After a mandate import has been submitted, you can use this endpoint to
// associate records
// in your system (using the `record_identifier` that you provided when creating
// the
// mandate import).
func (s *MandateImportEntryServiceImpl) List(ctx context.Context, p MandateImportEntryListParams, opts ...RequestOption) (*MandateImportEntryListResult, error) {
	uri, err := url.Parse(fmt.Sprintf(s.config.Endpoint() + "/mandate_import_entries"))
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
	req.Header.Set("GoCardless-Client-Version", "4.1.0")
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
		*MandateImportEntryListResult
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

	if result.MandateImportEntryListResult == nil {
		return nil, errors.New("missing result")
	}

	return result.MandateImportEntryListResult, nil
}

type MandateImportEntryListPagingIterator struct {
	cursor         string
	response       *MandateImportEntryListResult
	params         MandateImportEntryListParams
	service        *MandateImportEntryServiceImpl
	requestOptions []RequestOption
}

func (c *MandateImportEntryListPagingIterator) Next() bool {
	if c.cursor == "" && c.response != nil {
		return false
	}

	return true
}

func (c *MandateImportEntryListPagingIterator) Value(ctx context.Context) (*MandateImportEntryListResult, error) {
	if !c.Next() {
		return c.response, nil
	}

	s := c.service
	p := c.params
	p.After = c.cursor

	uri, err := url.Parse(fmt.Sprintf(s.config.Endpoint() + "/mandate_import_entries"))

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
	req.Header.Set("GoCardless-Client-Version", "4.1.0")
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
		*MandateImportEntryListResult
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

	if result.MandateImportEntryListResult == nil {
		return nil, errors.New("missing result")
	}

	c.response = result.MandateImportEntryListResult
	c.cursor = c.response.Meta.Cursors.After
	return c.response, nil
}

func (s *MandateImportEntryServiceImpl) All(ctx context.Context,
	p MandateImportEntryListParams,
	opts ...RequestOption) *MandateImportEntryListPagingIterator {
	return &MandateImportEntryListPagingIterator{
		params:         p,
		service:        s,
		requestOptions: opts,
	}
}
