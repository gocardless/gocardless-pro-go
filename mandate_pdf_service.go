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

// MandatePdfService manages mandate_pdfs
type MandatePdfServiceImpl struct {
	config Config
}

// MandatePdf model
type MandatePdf struct {
	ExpiresAt string `url:"expires_at,omitempty" json:"expires_at,omitempty"`
	Url       string `url:"url,omitempty" json:"url,omitempty"`
}

type MandatePdfService interface {
	Create(ctx context.Context, p MandatePdfCreateParams, opts ...RequestOption) (*MandatePdf, error)
}

type MandatePdfCreateParamsLinks struct {
	Mandate string `url:"mandate,omitempty" json:"mandate,omitempty"`
}

// MandatePdfCreateParams parameters
type MandatePdfCreateParams struct {
	AccountHolderName     string                       `url:"account_holder_name,omitempty" json:"account_holder_name,omitempty"`
	AccountNumber         string                       `url:"account_number,omitempty" json:"account_number,omitempty"`
	AccountType           string                       `url:"account_type,omitempty" json:"account_type,omitempty"`
	AddressLine1          string                       `url:"address_line1,omitempty" json:"address_line1,omitempty"`
	AddressLine2          string                       `url:"address_line2,omitempty" json:"address_line2,omitempty"`
	AddressLine3          string                       `url:"address_line3,omitempty" json:"address_line3,omitempty"`
	BankCode              string                       `url:"bank_code,omitempty" json:"bank_code,omitempty"`
	Bic                   string                       `url:"bic,omitempty" json:"bic,omitempty"`
	BranchCode            string                       `url:"branch_code,omitempty" json:"branch_code,omitempty"`
	City                  string                       `url:"city,omitempty" json:"city,omitempty"`
	CountryCode           string                       `url:"country_code,omitempty" json:"country_code,omitempty"`
	DanishIdentityNumber  string                       `url:"danish_identity_number,omitempty" json:"danish_identity_number,omitempty"`
	Iban                  string                       `url:"iban,omitempty" json:"iban,omitempty"`
	Links                 *MandatePdfCreateParamsLinks `url:"links,omitempty" json:"links,omitempty"`
	MandateReference      string                       `url:"mandate_reference,omitempty" json:"mandate_reference,omitempty"`
	PayerIpAddress        string                       `url:"payer_ip_address,omitempty" json:"payer_ip_address,omitempty"`
	PhoneNumber           string                       `url:"phone_number,omitempty" json:"phone_number,omitempty"`
	PostalCode            string                       `url:"postal_code,omitempty" json:"postal_code,omitempty"`
	Region                string                       `url:"region,omitempty" json:"region,omitempty"`
	Scheme                string                       `url:"scheme,omitempty" json:"scheme,omitempty"`
	SignatureDate         string                       `url:"signature_date,omitempty" json:"signature_date,omitempty"`
	SubscriptionAmount    int                          `url:"subscription_amount,omitempty" json:"subscription_amount,omitempty"`
	SubscriptionFrequency string                       `url:"subscription_frequency,omitempty" json:"subscription_frequency,omitempty"`
	SwedishIdentityNumber string                       `url:"swedish_identity_number,omitempty" json:"swedish_identity_number,omitempty"`
}

// Create
// Generates a PDF mandate and returns its temporary URL.
//
// Customer and bank account details can be left blank (for a blank mandate),
// provided manually, or inferred from the ID of an existing
// [mandate](#core-endpoints-mandates).
//
// By default, we'll generate PDF mandates in English.
//
// To generate a PDF mandate in another language, set the `Accept-Language`
// header when creating the PDF mandate to the relevant [ISO
// 639-1](http://en.wikipedia.org/wiki/List_of_ISO_639-1_codes) language code
// supported for the scheme.
//
// | Scheme           | Supported languages
//
//	|
//
// | :--------------- |
// :-------------------------------------------------------------------------------------------------------------------------------------------
// |
// | ACH              | English (`en`)
//
//	|
//
// | Autogiro         | English (`en`), Swedish (`sv`)
//
//	|
//
// | Bacs             | English (`en`)
//
//	|
//
// | BECS             | English (`en`)
//
//	|
//
// | BECS NZ          | English (`en`)
//
//	|
//
// | Betalingsservice | Danish (`da`), English (`en`)
//
//	|
//
// | PAD              | English (`en`)
//
//	|
//
// | SEPA Core        | Danish (`da`), Dutch (`nl`), English (`en`), French
// (`fr`), German (`de`), Italian (`it`), Portuguese (`pt`), Spanish (`es`),
// Swedish (`sv`) |
func (s *MandatePdfServiceImpl) Create(ctx context.Context, p MandatePdfCreateParams, opts ...RequestOption) (*MandatePdf, error) {
	uri, err := url.Parse(fmt.Sprintf(s.config.Endpoint() + "/mandate_pdfs"))
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
		"mandate_pdfs": p,
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
	req.Header.Set("GoCardless-Client-Version", "3.6.0")
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
		Err        *APIError   `json:"error"`
		MandatePdf *MandatePdf `json:"mandate_pdfs"`
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

	if result.MandatePdf == nil {
		return nil, errors.New("missing result")
	}

	return result.MandatePdf, nil
}
