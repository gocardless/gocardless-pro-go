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
)

// BankAccountHolderVerificationService manages bank_account_holder_verifications
type BankAccountHolderVerificationServiceImpl struct {
	config Config
}

// BankAccountHolderVerification model
type BankAccountHolderVerification struct {
	ActualAccountName string `url:"actual_account_name,omitempty" json:"actual_account_name,omitempty"`
	Id                string `url:"id,omitempty" json:"id,omitempty"`
	Result            string `url:"result,omitempty" json:"result,omitempty"`
	Status            string `url:"status,omitempty" json:"status,omitempty"`
	Type              string `url:"type,omitempty" json:"type,omitempty"`
}

type BankAccountHolderVerificationService interface {
	Create(ctx context.Context, p BankAccountHolderVerificationCreateParams, opts ...RequestOption) (*BankAccountHolderVerification, error)
	Get(ctx context.Context, identity string, opts ...RequestOption) (*BankAccountHolderVerification, error)
}

type BankAccountHolderVerificationCreateParamsLinks struct {
	BankAccount string `url:"bank_account,omitempty" json:"bank_account,omitempty"`
}

// BankAccountHolderVerificationCreateParams parameters
type BankAccountHolderVerificationCreateParams struct {
	Links BankAccountHolderVerificationCreateParamsLinks `url:"links,omitempty" json:"links,omitempty"`
	Type  string                                         `url:"type,omitempty" json:"type,omitempty"`
}

// Create
// Verify the account holder of the bank account. A complete verification can be
// attached when creating an outbound payment. This endpoint allows partner
// merchants to create Confirmation of Payee checks on customer bank accounts
// before sending outbound payments.
func (s *BankAccountHolderVerificationServiceImpl) Create(ctx context.Context, p BankAccountHolderVerificationCreateParams, opts ...RequestOption) (*BankAccountHolderVerification, error) {
	uri, err := url.Parse(fmt.Sprintf(s.config.Endpoint() + "/bank_account_holder_verifications"))
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
		"bank_account_holder_verifications": p,
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
	req.Header.Set("GoCardless-Client-Version", "6.1.0")
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
		Err                           *APIError                      `json:"error"`
		BankAccountHolderVerification *BankAccountHolderVerification `json:"bank_account_holder_verifications"`
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

	if result.BankAccountHolderVerification == nil {
		return nil, errors.New("missing result")
	}

	return result.BankAccountHolderVerification, nil
}

// Get
// Fetches a bank account holder verification by ID.
func (s *BankAccountHolderVerificationServiceImpl) Get(ctx context.Context, identity string, opts ...RequestOption) (*BankAccountHolderVerification, error) {
	uri, err := url.Parse(fmt.Sprintf(s.config.Endpoint()+"/bank_account_holder_verifications/%v",
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
	req.Header.Set("GoCardless-Client-Version", "6.1.0")
	req.Header.Set("User-Agent", userAgent)

	for key, value := range o.headers {
		req.Header.Set(key, value)
	}

	client := s.config.Client()
	if client == nil {
		client = http.DefaultClient
	}

	var result struct {
		Err                           *APIError                      `json:"error"`
		BankAccountHolderVerification *BankAccountHolderVerification `json:"bank_account_holder_verifications"`
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

	if result.BankAccountHolderVerification == nil {
		return nil, errors.New("missing result")
	}

	return result.BankAccountHolderVerification, nil
}
