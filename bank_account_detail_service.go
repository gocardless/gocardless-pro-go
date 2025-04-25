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

// BankAccountDetailService manages bank_account_details
type BankAccountDetailServiceImpl struct {
	config Config
}

// BankAccountDetail model
type BankAccountDetail struct {
	Ciphertext   string `url:"ciphertext,omitempty" json:"ciphertext,omitempty"`
	EncryptedKey string `url:"encrypted_key,omitempty" json:"encrypted_key,omitempty"`
	Iv           string `url:"iv,omitempty" json:"iv,omitempty"`
	Protected    string `url:"protected,omitempty" json:"protected,omitempty"`
	Tag          string `url:"tag,omitempty" json:"tag,omitempty"`
}

type BankAccountDetailService interface {
	Get(ctx context.Context, identity string, p BankAccountDetailGetParams, opts ...RequestOption) (*BankAccountDetail, error)
}

// BankAccountDetailGetParams parameters
type BankAccountDetailGetParams struct {
}

// Get
// Returns bank account details in the flattened JSON Web Encryption format
// described in RFC 7516
func (s *BankAccountDetailServiceImpl) Get(ctx context.Context, identity string, p BankAccountDetailGetParams, opts ...RequestOption) (*BankAccountDetail, error) {
	uri, err := url.Parse(fmt.Sprintf(s.config.Endpoint()+"/bank_account_details/%v",
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
	req.Header.Set("GoCardless-Client-Version", "4.7.0")
	req.Header.Set("User-Agent", userAgent)

	for key, value := range o.headers {
		req.Header.Set(key, value)
	}

	client := s.config.Client()
	if client == nil {
		client = http.DefaultClient
	}

	var result struct {
		Err               *APIError          `json:"error"`
		BankAccountDetail *BankAccountDetail `json:"bank_account_details"`
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

	if result.BankAccountDetail == nil {
		return nil, errors.New("missing result")
	}

	return result.BankAccountDetail, nil
}
