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

// TransferredMandateService manages transferred_mandates
type TransferredMandateServiceImpl struct {
	config Config
}

type TransferredMandateLinks struct {
	CustomerBankAccount string `url:"customer_bank_account,omitempty" json:"customer_bank_account,omitempty"`
	Mandate             string `url:"mandate,omitempty" json:"mandate,omitempty"`
}

// TransferredMandate model
type TransferredMandate struct {
	EncryptedCustomerBankDetails string                   `url:"encrypted_customer_bank_details,omitempty" json:"encrypted_customer_bank_details,omitempty"`
	EncryptedDecryptionKey       string                   `url:"encrypted_decryption_key,omitempty" json:"encrypted_decryption_key,omitempty"`
	Links                        *TransferredMandateLinks `url:"links,omitempty" json:"links,omitempty"`
	PublicKeyId                  string                   `url:"public_key_id,omitempty" json:"public_key_id,omitempty"`
}

type TransferredMandateService interface {
	TransferredMandates(ctx context.Context, identity string, p TransferredMandateTransferredMandatesParams, opts ...RequestOption) (*TransferredMandate, error)
}

// TransferredMandateTransferredMandatesParams parameters
type TransferredMandateTransferredMandatesParams struct {
}

// TransferredMandates
// Returns new customer bank details for a mandate that's been recently
// transferred
func (s *TransferredMandateServiceImpl) TransferredMandates(ctx context.Context, identity string, p TransferredMandateTransferredMandatesParams, opts ...RequestOption) (*TransferredMandate, error) {
	uri, err := url.Parse(fmt.Sprintf(s.config.Endpoint()+"/transferred_mandates/%v",
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
	req.Header.Set("GoCardless-Client-Version", "4.0.0")
	req.Header.Set("User-Agent", userAgent)

	for key, value := range o.headers {
		req.Header.Set(key, value)
	}

	client := s.config.Client()
	if client == nil {
		client = http.DefaultClient
	}

	var result struct {
		Err                *APIError           `json:"error"`
		TransferredMandate *TransferredMandate `json:"transferred_mandates"`
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

	if result.TransferredMandate == nil {
		return nil, errors.New("missing result")
	}

	return result.TransferredMandate, nil
}
