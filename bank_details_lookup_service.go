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
type BankDetailsLookupService struct {
  endpoint string
  token string
  client *http.Client
}


// BankDetailsLookup model
type BankDetailsLookup struct {
      AvailableDebitSchemes []string `url:",omitempty" json:"available_debit_schemes,omitempty"`
      BankName string `url:",omitempty" json:"bank_name,omitempty"`
      Bic string `url:",omitempty" json:"bic,omitempty"`
      }




// BankDetailsLookupCreateParams parameters
type BankDetailsLookupCreateParams struct {
      AccountNumber string `url:",omitempty" json:"account_number,omitempty"`
      BankCode string `url:",omitempty" json:"bank_code,omitempty"`
      BranchCode string `url:",omitempty" json:"branch_code,omitempty"`
      CountryCode string `url:",omitempty" json:"country_code,omitempty"`
      Iban string `url:",omitempty" json:"iban,omitempty"`
      }

// Create
// Performs a bank details lookup. As part of the lookup, a modulus check and
// reachability check are performed.
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
// _Note:_ Usage of this endpoint is monitored. If your organisation relies on
// GoCardless for
// modulus or reachability checking but not for payment collection, please get
// in touch.
func (s *BankDetailsLookupService) Create(ctx context.Context, p BankDetailsLookupCreateParams) (*BankDetailsLookup,error) {
  uri, err := url.Parse(fmt.Sprintf(s.endpoint + "/bank_details_lookups",))
  if err != nil {
    return nil, err
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
  req.WithContext(ctx)
  req.Header.Set("Authorization", "Bearer "+s.token)
  req.Header.Set("GoCardless-Version", "2015-07-06")
  req.Header.Set("Content-Type", "application/json")
  req.Header.Set("Idempotency-Key", NewIdempotencyKey())

  client := s.client
  if client == nil {
    client = http.DefaultClient
  }

  var result struct {
    Err *APIError `json:"error"`
BankDetailsLookup *BankDetailsLookup `json:"bank_details_lookups"`
  }

  err = try(3, func() error {
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

