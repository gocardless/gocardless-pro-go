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

// CreditorService manages creditors
type CreditorServiceImpl struct {
	config Config
}

type CreditorLinks struct {
	DefaultAudPayoutAccount string `url:"default_aud_payout_account,omitempty" json:"default_aud_payout_account,omitempty"`
	DefaultCadPayoutAccount string `url:"default_cad_payout_account,omitempty" json:"default_cad_payout_account,omitempty"`
	DefaultDkkPayoutAccount string `url:"default_dkk_payout_account,omitempty" json:"default_dkk_payout_account,omitempty"`
	DefaultEurPayoutAccount string `url:"default_eur_payout_account,omitempty" json:"default_eur_payout_account,omitempty"`
	DefaultGbpPayoutAccount string `url:"default_gbp_payout_account,omitempty" json:"default_gbp_payout_account,omitempty"`
	DefaultNzdPayoutAccount string `url:"default_nzd_payout_account,omitempty" json:"default_nzd_payout_account,omitempty"`
	DefaultSekPayoutAccount string `url:"default_sek_payout_account,omitempty" json:"default_sek_payout_account,omitempty"`
	DefaultUsdPayoutAccount string `url:"default_usd_payout_account,omitempty" json:"default_usd_payout_account,omitempty"`
}

type CreditorSchemeIdentifiers struct {
	AddressLine1               string `url:"address_line1,omitempty" json:"address_line1,omitempty"`
	AddressLine2               string `url:"address_line2,omitempty" json:"address_line2,omitempty"`
	AddressLine3               string `url:"address_line3,omitempty" json:"address_line3,omitempty"`
	CanSpecifyMandateReference bool   `url:"can_specify_mandate_reference,omitempty" json:"can_specify_mandate_reference,omitempty"`
	City                       string `url:"city,omitempty" json:"city,omitempty"`
	CountryCode                string `url:"country_code,omitempty" json:"country_code,omitempty"`
	Currency                   string `url:"currency,omitempty" json:"currency,omitempty"`
	Email                      string `url:"email,omitempty" json:"email,omitempty"`
	MinimumAdvanceNotice       int    `url:"minimum_advance_notice,omitempty" json:"minimum_advance_notice,omitempty"`
	Name                       string `url:"name,omitempty" json:"name,omitempty"`
	PhoneNumber                string `url:"phone_number,omitempty" json:"phone_number,omitempty"`
	PostalCode                 string `url:"postal_code,omitempty" json:"postal_code,omitempty"`
	Reference                  string `url:"reference,omitempty" json:"reference,omitempty"`
	Region                     string `url:"region,omitempty" json:"region,omitempty"`
	Scheme                     string `url:"scheme,omitempty" json:"scheme,omitempty"`
}

// Creditor model
type Creditor struct {
	AddressLine1                        string                      `url:"address_line1,omitempty" json:"address_line1,omitempty"`
	AddressLine2                        string                      `url:"address_line2,omitempty" json:"address_line2,omitempty"`
	AddressLine3                        string                      `url:"address_line3,omitempty" json:"address_line3,omitempty"`
	CanCreateRefunds                    bool                        `url:"can_create_refunds,omitempty" json:"can_create_refunds,omitempty"`
	City                                string                      `url:"city,omitempty" json:"city,omitempty"`
	CountryCode                         string                      `url:"country_code,omitempty" json:"country_code,omitempty"`
	CreatedAt                           string                      `url:"created_at,omitempty" json:"created_at,omitempty"`
	CustomPaymentPagesEnabled           bool                        `url:"custom_payment_pages_enabled,omitempty" json:"custom_payment_pages_enabled,omitempty"`
	FxPayoutCurrency                    string                      `url:"fx_payout_currency,omitempty" json:"fx_payout_currency,omitempty"`
	Id                                  string                      `url:"id,omitempty" json:"id,omitempty"`
	Links                               *CreditorLinks              `url:"links,omitempty" json:"links,omitempty"`
	LogoUrl                             string                      `url:"logo_url,omitempty" json:"logo_url,omitempty"`
	MandateImportsEnabled               bool                        `url:"mandate_imports_enabled,omitempty" json:"mandate_imports_enabled,omitempty"`
	MerchantResponsibleForNotifications bool                        `url:"merchant_responsible_for_notifications,omitempty" json:"merchant_responsible_for_notifications,omitempty"`
	Name                                string                      `url:"name,omitempty" json:"name,omitempty"`
	PostalCode                          string                      `url:"postal_code,omitempty" json:"postal_code,omitempty"`
	Region                              string                      `url:"region,omitempty" json:"region,omitempty"`
	SchemeIdentifiers                   []CreditorSchemeIdentifiers `url:"scheme_identifiers,omitempty" json:"scheme_identifiers,omitempty"`
	VerificationStatus                  string                      `url:"verification_status,omitempty" json:"verification_status,omitempty"`
}

type CreditorService interface {
	Create(ctx context.Context, p CreditorCreateParams, opts ...RequestOption) (*Creditor, error)
	List(ctx context.Context, p CreditorListParams, opts ...RequestOption) (*CreditorListResult, error)
	All(ctx context.Context, p CreditorListParams, opts ...RequestOption) *CreditorListPagingIterator
	Get(ctx context.Context, identity string, p CreditorGetParams, opts ...RequestOption) (*Creditor, error)
	Update(ctx context.Context, identity string, p CreditorUpdateParams, opts ...RequestOption) (*Creditor, error)
}

// CreditorCreateParams parameters
type CreditorCreateParams struct {
	AddressLine1 string                 `url:"address_line1,omitempty" json:"address_line1,omitempty"`
	AddressLine2 string                 `url:"address_line2,omitempty" json:"address_line2,omitempty"`
	AddressLine3 string                 `url:"address_line3,omitempty" json:"address_line3,omitempty"`
	City         string                 `url:"city,omitempty" json:"city,omitempty"`
	CountryCode  string                 `url:"country_code,omitempty" json:"country_code,omitempty"`
	Links        map[string]interface{} `url:"links,omitempty" json:"links,omitempty"`
	Name         string                 `url:"name,omitempty" json:"name,omitempty"`
	PostalCode   string                 `url:"postal_code,omitempty" json:"postal_code,omitempty"`
	Region       string                 `url:"region,omitempty" json:"region,omitempty"`
}

// Create
// Creates a new creditor.
func (s *CreditorServiceImpl) Create(ctx context.Context, p CreditorCreateParams, opts ...RequestOption) (*Creditor, error) {
	uri, err := url.Parse(fmt.Sprintf(s.config.Endpoint() + "/creditors"))
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
		"creditors": p,
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
	req.Header.Set("GoCardless-Client-Version", "2.5.0")
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
		Err      *APIError `json:"error"`
		Creditor *Creditor `json:"creditors"`
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

	if result.Creditor == nil {
		return nil, errors.New("missing result")
	}

	return result.Creditor, nil
}

type CreditorListParamsCreatedAt struct {
	Gt  string `url:"gt,omitempty" json:"gt,omitempty"`
	Gte string `url:"gte,omitempty" json:"gte,omitempty"`
	Lt  string `url:"lt,omitempty" json:"lt,omitempty"`
	Lte string `url:"lte,omitempty" json:"lte,omitempty"`
}

// CreditorListParams parameters
type CreditorListParams struct {
	After     string                       `url:"after,omitempty" json:"after,omitempty"`
	Before    string                       `url:"before,omitempty" json:"before,omitempty"`
	CreatedAt *CreditorListParamsCreatedAt `url:"created_at,omitempty" json:"created_at,omitempty"`
	Limit     int                          `url:"limit,omitempty" json:"limit,omitempty"`
}

type CreditorListResultMetaCursors struct {
	After  string `url:"after,omitempty" json:"after,omitempty"`
	Before string `url:"before,omitempty" json:"before,omitempty"`
}

type CreditorListResultMeta struct {
	Cursors *CreditorListResultMetaCursors `url:"cursors,omitempty" json:"cursors,omitempty"`
	Limit   int                            `url:"limit,omitempty" json:"limit,omitempty"`
}

type CreditorListResult struct {
	Creditors []Creditor             `json:"creditors"`
	Meta      CreditorListResultMeta `url:"meta,omitempty" json:"meta,omitempty"`
}

// List
// Returns a [cursor-paginated](#api-usage-cursor-pagination) list of your
// creditors.
func (s *CreditorServiceImpl) List(ctx context.Context, p CreditorListParams, opts ...RequestOption) (*CreditorListResult, error) {
	uri, err := url.Parse(fmt.Sprintf(s.config.Endpoint() + "/creditors"))
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
	req.Header.Set("GoCardless-Client-Version", "2.5.0")
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
		*CreditorListResult
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

	if result.CreditorListResult == nil {
		return nil, errors.New("missing result")
	}

	return result.CreditorListResult, nil
}

type CreditorListPagingIterator struct {
	cursor         string
	response       *CreditorListResult
	params         CreditorListParams
	service        *CreditorServiceImpl
	requestOptions []RequestOption
}

func (c *CreditorListPagingIterator) Next() bool {
	if c.cursor == "" && c.response != nil {
		return false
	}

	return true
}

func (c *CreditorListPagingIterator) Value(ctx context.Context) (*CreditorListResult, error) {
	if !c.Next() {
		return c.response, nil
	}

	s := c.service
	p := c.params
	p.After = c.cursor

	uri, err := url.Parse(fmt.Sprintf(s.config.Endpoint() + "/creditors"))

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
	req.Header.Set("GoCardless-Client-Version", "2.5.0")
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
		*CreditorListResult
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

	if result.CreditorListResult == nil {
		return nil, errors.New("missing result")
	}

	c.response = result.CreditorListResult
	c.cursor = c.response.Meta.Cursors.After
	return c.response, nil
}

func (s *CreditorServiceImpl) All(ctx context.Context,
	p CreditorListParams,
	opts ...RequestOption) *CreditorListPagingIterator {
	return &CreditorListPagingIterator{
		params:         p,
		service:        s,
		requestOptions: opts,
	}
}

// CreditorGetParams parameters
type CreditorGetParams struct {
}

// Get
// Retrieves the details of an existing creditor.
func (s *CreditorServiceImpl) Get(ctx context.Context, identity string, p CreditorGetParams, opts ...RequestOption) (*Creditor, error) {
	uri, err := url.Parse(fmt.Sprintf(s.config.Endpoint()+"/creditors/%v",
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
	req.Header.Set("GoCardless-Client-Version", "2.5.0")
	req.Header.Set("User-Agent", userAgent)

	for key, value := range o.headers {
		req.Header.Set(key, value)
	}

	client := s.config.Client()
	if client == nil {
		client = http.DefaultClient
	}

	var result struct {
		Err      *APIError `json:"error"`
		Creditor *Creditor `json:"creditors"`
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

	if result.Creditor == nil {
		return nil, errors.New("missing result")
	}

	return result.Creditor, nil
}

type CreditorUpdateParamsLinks struct {
	DefaultAudPayoutAccount string `url:"default_aud_payout_account,omitempty" json:"default_aud_payout_account,omitempty"`
	DefaultCadPayoutAccount string `url:"default_cad_payout_account,omitempty" json:"default_cad_payout_account,omitempty"`
	DefaultDkkPayoutAccount string `url:"default_dkk_payout_account,omitempty" json:"default_dkk_payout_account,omitempty"`
	DefaultEurPayoutAccount string `url:"default_eur_payout_account,omitempty" json:"default_eur_payout_account,omitempty"`
	DefaultGbpPayoutAccount string `url:"default_gbp_payout_account,omitempty" json:"default_gbp_payout_account,omitempty"`
	DefaultNzdPayoutAccount string `url:"default_nzd_payout_account,omitempty" json:"default_nzd_payout_account,omitempty"`
	DefaultSekPayoutAccount string `url:"default_sek_payout_account,omitempty" json:"default_sek_payout_account,omitempty"`
	DefaultUsdPayoutAccount string `url:"default_usd_payout_account,omitempty" json:"default_usd_payout_account,omitempty"`
}

// CreditorUpdateParams parameters
type CreditorUpdateParams struct {
	AddressLine1 string                     `url:"address_line1,omitempty" json:"address_line1,omitempty"`
	AddressLine2 string                     `url:"address_line2,omitempty" json:"address_line2,omitempty"`
	AddressLine3 string                     `url:"address_line3,omitempty" json:"address_line3,omitempty"`
	City         string                     `url:"city,omitempty" json:"city,omitempty"`
	CountryCode  string                     `url:"country_code,omitempty" json:"country_code,omitempty"`
	Links        *CreditorUpdateParamsLinks `url:"links,omitempty" json:"links,omitempty"`
	Name         string                     `url:"name,omitempty" json:"name,omitempty"`
	PostalCode   string                     `url:"postal_code,omitempty" json:"postal_code,omitempty"`
	Region       string                     `url:"region,omitempty" json:"region,omitempty"`
}

// Update
// Updates a creditor object. Supports all of the fields supported when creating
// a creditor.
func (s *CreditorServiceImpl) Update(ctx context.Context, identity string, p CreditorUpdateParams, opts ...RequestOption) (*Creditor, error) {
	uri, err := url.Parse(fmt.Sprintf(s.config.Endpoint()+"/creditors/%v",
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
		"creditors": p,
	})
	if err != nil {
		return nil, err
	}
	body = &buf

	req, err := http.NewRequest("PUT", uri.String(), body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	req.Header.Set("Authorization", "Bearer "+s.config.Token())
	req.Header.Set("GoCardless-Version", "2015-07-06")
	req.Header.Set("GoCardless-Client-Library", "gocardless-pro-go")
	req.Header.Set("GoCardless-Client-Version", "2.5.0")
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
		Err      *APIError `json:"error"`
		Creditor *Creditor `json:"creditors"`
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

	if result.Creditor == nil {
		return nil, errors.New("missing result")
	}

	return result.Creditor, nil
}
