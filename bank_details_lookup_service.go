package gocardless

import (
  "bytes"
  "context"
  "encoding/json"
  "fmt"
  "io"
  "net/http"
  "net/url"

  "github.com/google/go-querystring/query"
)

var _ = query.Values
var _ = bytes.NewBuffer


type BankDetailsLookupService struct {
  endpoint string
  token string
  client *http.Client
}



// BankDetailsLookupCreateParams parameters
type BankDetailsLookupCreateParams struct {
      AccountNumber string `url:",omitempty",json:"account_number,omitempty"`
        BankCode string `url:",omitempty",json:"bank_code,omitempty"`
        BranchCode string `url:",omitempty",json:"branch_code,omitempty"`
        CountryCode string `url:",omitempty",json:"country_code,omitempty"`
        Iban string `url:",omitempty",json:"iban,omitempty"`
        
    }
// BankDetailsLookupCreateResult parameters
type BankDetailsLookupCreateResult struct {
      BankDetailsLookups struct {
      AvailableDebitSchemes []string `url:",omitempty",json:"available_debit_schemes,omitempty"`
        BankName string `url:",omitempty",json:"bank_name,omitempty"`
        Bic string `url:",omitempty",json:"bic,omitempty"`
        
    } `url:",omitempty",json:"bank_details_lookups,omitempty"`
        
    }

// Create
// Performs a bank details lookup.
// 
// As part of the lookup a modulus check
// and reachability check are performed.
// 
// Bank account details may be
// supplied using [local details](#appendix-local-bank-details) or an IBAN.
//
// 
// _Note:_ Usage of this endpoint is monitored. If your organisation relies
// on GoCardless for
// modulus or reachability checking but not for payment
// collection, please get in touch.
func (s *BankDetailsLookupService) Create(
  ctx context.Context,
  p BankDetailsLookupCreateParams) (*BankDetailsLookupCreateResult, error) {
  uri, err := url.Parse(fmt.Sprintf(
      s.endpoint + "/bank_details_lookups",))
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

  client := s.client
  if client == nil {
    client = http.DefaultClient
  }

  res, err := client.Do(req)
  if err != nil {
    return nil, err
  }
  defer res.Body.Close()

  var result struct {
    *BankDetailsLookupCreateResult
    Err *APIError `json:"error"`
  }

  err = json.NewDecoder(res.Body).Decode(&result)
  if err != nil {
    return nil, err
  }

  if result.Err != nil {
    return nil, result.Err
  }

  return result.BankDetailsLookupCreateResult, nil
}

