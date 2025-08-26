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

// CreditorBankAccountValidateService manages creditor_bank_account_validates
type CreditorBankAccountValidateServiceImpl struct {
	config Config
}

type CreditorBankAccountValidateInvalidReasons struct {
	Field   string `url:"field,omitempty" json:"field,omitempty"`
	Message string `url:"message,omitempty" json:"message,omitempty"`
}

// CreditorBankAccountValidate model
type CreditorBankAccountValidate struct {
	BankName       string                                      `url:"bank_name,omitempty" json:"bank_name,omitempty"`
	IconUrl        string                                      `url:"icon_url,omitempty" json:"icon_url,omitempty"`
	InvalidReasons []CreditorBankAccountValidateInvalidReasons `url:"invalid_reasons,omitempty" json:"invalid_reasons,omitempty"`
	IsValid        bool                                        `url:"is_valid,omitempty" json:"is_valid,omitempty"`
}

type CreditorBankAccountValidateService interface {
	Validate(ctx context.Context, p CreditorBankAccountValidateValidateParams, opts ...RequestOption) (*CreditorBankAccountValidate, error)
}

type CreditorBankAccountValidateValidateParamsLocalDetails struct {
	BankNumber string `url:"bank_number,omitempty" json:"bank_number,omitempty"`
	SortCode   string `url:"sort_code,omitempty" json:"sort_code,omitempty"`
}

// CreditorBankAccountValidateValidateParams parameters
type CreditorBankAccountValidateValidateParams struct {
	Iban         string                                                 `url:"iban,omitempty" json:"iban,omitempty"`
	LocalDetails *CreditorBankAccountValidateValidateParamsLocalDetails `url:"local_details,omitempty" json:"local_details,omitempty"`
}

// Validate
// Validate bank details without creating a creditor bank account
func (s *CreditorBankAccountValidateServiceImpl) Validate(ctx context.Context, p CreditorBankAccountValidateValidateParams, opts ...RequestOption) (*CreditorBankAccountValidate, error) {
	uri, err := url.Parse(fmt.Sprintf(s.config.Endpoint() + "/creditor_bank_accounts/validate"))
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
	req.Header.Set("GoCardless-Client-Version", "5.2.0")
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
		Err                         *APIError                    `json:"error"`
		CreditorBankAccountValidate *CreditorBankAccountValidate `json:"creditor_bank_accounts"`
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

	if result.CreditorBankAccountValidate == nil {
		return nil, errors.New("missing result")
	}

	return result.CreditorBankAccountValidate, nil
}
