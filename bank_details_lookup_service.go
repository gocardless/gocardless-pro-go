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

// BankDetailsLookupService manages bank_details_lookups
type BankDetailsLookupServiceImpl struct {
	config Config
}

// BankDetailsLookup model
type BankDetailsLookup struct {
	AvailableDebitSchemes []string `url:"available_debit_schemes,omitempty" json:"available_debit_schemes,omitempty"`
	BankName              string   `url:"bank_name,omitempty" json:"bank_name,omitempty"`
	Bic                   string   `url:"bic,omitempty" json:"bic,omitempty"`
}

type BankDetailsLookupService interface {
	Create(ctx context.Context, p BankDetailsLookupCreateParams, opts ...RequestOption) (*BankDetailsLookup, error)
}

// BankDetailsLookupCreateParams parameters
type BankDetailsLookupCreateParams struct {
	AccountHolderName string `url:"account_holder_name,omitempty" json:"account_holder_name,omitempty"`
	AccountNumber     string `url:"account_number,omitempty" json:"account_number,omitempty"`
	BankCode          string `url:"bank_code,omitempty" json:"bank_code,omitempty"`
	BranchCode        string `url:"branch_code,omitempty" json:"branch_code,omitempty"`
	CountryCode       string `url:"country_code,omitempty" json:"country_code,omitempty"`
	Iban              string `url:"iban,omitempty" json:"iban,omitempty"`
}

// Create
// Performs a bank details lookup. As part of the lookup, a modulus check and
// reachability check are performed.
//
// For UK-based bank accounts, where an account holder name is provided (and an
// account number, a sort code or an iban
// are already present), we verify that the account holder name and bank account
// number match the details held by
// the relevant bank.
//
// If your request returns an [error](#api-usage-errors) or the
// `available_debit_schemes`
// attribute is an empty array, you will not be able to collect payments from
// the
// specified bank account. GoCardless may be able to collect payments from an
// account
// even if no `bic` is returned.
//
// Bank account details may be supplied using [local
// details](#appendix-local-bank-details) or an IBAN.
//
// _ACH scheme_ For compliance reasons, an extra validation step is done using
// a third-party provider to make sure the customer's bank account can accept
// Direct Debit. If a bank account is discovered to be closed or invalid, the
// customer is requested to adjust the account number/routing number and
// succeed in this check to continue with the flow.
//
// _Note:_ Usage of this endpoint is monitored. If your organisation relies on
// GoCardless for
// modulus or reachability checking but not for payment collection, please get
// in touch.
func (s *BankDetailsLookupServiceImpl) Create(ctx context.Context, p BankDetailsLookupCreateParams, opts ...RequestOption) (*BankDetailsLookup, error) {
	uri, err := url.Parse(fmt.Sprintf(s.config.Endpoint() + "/bank_details_lookups"))
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
		"bank_details_lookups": p,
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
	req.Header.Set("GoCardless-Client-Version", "3.7.0")
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
		BankDetailsLookup *BankDetailsLookup `json:"bank_details_lookups"`
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

	if result.BankDetailsLookup == nil {
		return nil, errors.New("missing result")
	}

	return result.BankDetailsLookup, nil
}
