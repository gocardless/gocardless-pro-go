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


// CustomerBankAccountService manages customer_bank_accounts
type CustomerBankAccountService struct {
  endpoint string
  token string
  client *http.Client
}


// CustomerBankAccount model
type CustomerBankAccount struct {
      AccountHolderName string `url:",omitempty" json:"account_holder_name,omitempty"`
      AccountNumberEnding string `url:",omitempty" json:"account_number_ending,omitempty"`
      BankName string `url:",omitempty" json:"bank_name,omitempty"`
      CountryCode string `url:",omitempty" json:"country_code,omitempty"`
      CreatedAt string `url:",omitempty" json:"created_at,omitempty"`
      Currency string `url:",omitempty" json:"currency,omitempty"`
      Enabled bool `url:",omitempty" json:"enabled,omitempty"`
      Id string `url:",omitempty" json:"id,omitempty"`
      Links struct {
      Customer string `url:",omitempty" json:"customer,omitempty"`
      } `url:",omitempty" json:"links,omitempty"`
      Metadata map[string]interface{} `url:",omitempty" json:"metadata,omitempty"`
      }




// CustomerBankAccountCreateParams parameters
type CustomerBankAccountCreateParams struct {
      AccountHolderName string `url:",omitempty" json:"account_holder_name,omitempty"`
      AccountNumber string `url:",omitempty" json:"account_number,omitempty"`
      BankCode string `url:",omitempty" json:"bank_code,omitempty"`
      BranchCode string `url:",omitempty" json:"branch_code,omitempty"`
      CountryCode string `url:",omitempty" json:"country_code,omitempty"`
      Currency string `url:",omitempty" json:"currency,omitempty"`
      Iban string `url:",omitempty" json:"iban,omitempty"`
      Links struct {
      Customer string `url:",omitempty" json:"customer,omitempty"`
      CustomerBankAccountToken string `url:",omitempty" json:"customer_bank_account_token,omitempty"`
      } `url:",omitempty" json:"links,omitempty"`
      Metadata map[string]interface{} `url:",omitempty" json:"metadata,omitempty"`
      }

// Create
// Creates a new customer bank account object.
// 
// There are three different
// ways to supply bank account details:
// 
// - [Local
// details](#appendix-local-bank-details)
// 
// - IBAN
// 
// - [Customer Bank
// Account Tokens](#javascript-flow-create-a-customer-bank-account-token)
// 
//
// For more information on the different fields required in each country, see
// [local bank details](#appendix-local-bank-details).
func (s *CustomerBankAccountService) Create(ctx context.Context, p CustomerBankAccountCreateParams) (*CustomerBankAccount,error) {
  uri, err := url.Parse(fmt.Sprintf(s.endpoint + "/customer_bank_accounts",))
  if err != nil {
    return nil, err
  }

  var body io.Reader

  var buf bytes.Buffer
  err = json.NewEncoder(&buf).Encode(map[string]interface{}{
    "customer_bank_accounts": p,
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
CustomerBankAccount *CustomerBankAccount `json:"customer_bank_accounts"`
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

if result.CustomerBankAccount == nil {
    return nil, errors.New("missing result")
  }

  return result.CustomerBankAccount, nil
}


// CustomerBankAccountListParams parameters
type CustomerBankAccountListParams struct {
      After string `url:",omitempty" json:"after,omitempty"`
      Before string `url:",omitempty" json:"before,omitempty"`
      CreatedAt struct {
      Gt string `url:",omitempty" json:"gt,omitempty"`
      Gte string `url:",omitempty" json:"gte,omitempty"`
      Lt string `url:",omitempty" json:"lt,omitempty"`
      Lte string `url:",omitempty" json:"lte,omitempty"`
      } `url:",omitempty" json:"created_at,omitempty"`
      Customer string `url:",omitempty" json:"customer,omitempty"`
      Enabled bool `url:",omitempty" json:"enabled,omitempty"`
      Limit int `url:",omitempty" json:"limit,omitempty"`
      }

// CustomerBankAccountListResult response including pagination metadata
type CustomerBankAccountListResult struct {
  CustomerBankAccounts []CustomerBankAccount `json:"customer_bank_accounts"`
  Meta struct {
      Cursors struct {
      After string `url:",omitempty" json:"after,omitempty"`
      Before string `url:",omitempty" json:"before,omitempty"`
      } `url:",omitempty" json:"cursors,omitempty"`
      Limit int `url:",omitempty" json:"limit,omitempty"`
      } `json:"meta"`
}


// List
// Returns a [cursor-paginated](#api-usage-cursor-pagination) list of your bank
// accounts.
func (s *CustomerBankAccountService) List(ctx context.Context, p CustomerBankAccountListParams) (*CustomerBankAccountListResult,error) {
  uri, err := url.Parse(fmt.Sprintf(s.endpoint + "/customer_bank_accounts",))
  if err != nil {
    return nil, err
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
  req.WithContext(ctx)
  req.Header.Set("Authorization", "Bearer "+s.token)
  req.Header.Set("GoCardless-Version", "2015-07-06")
  

  client := s.client
  if client == nil {
    client = http.DefaultClient
  }

  var result struct {
    Err *APIError `json:"error"`
*CustomerBankAccountListResult
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

if result.CustomerBankAccountListResult == nil {
    return nil, errors.New("missing result")
  }

  return result.CustomerBankAccountListResult, nil
}



// Get
// Retrieves the details of an existing bank account.
func (s *CustomerBankAccountService) Get(ctx context.Context,identity string) (*CustomerBankAccount,error) {
  uri, err := url.Parse(fmt.Sprintf(s.endpoint + "/customer_bank_accounts/%v",
      identity,))
  if err != nil {
    return nil, err
  }

  var body io.Reader

  

  req, err := http.NewRequest("GET", uri.String(), body)
  if err != nil {
    return nil, err
  }
  req.WithContext(ctx)
  req.Header.Set("Authorization", "Bearer "+s.token)
  req.Header.Set("GoCardless-Version", "2015-07-06")
  

  client := s.client
  if client == nil {
    client = http.DefaultClient
  }

  var result struct {
    Err *APIError `json:"error"`
CustomerBankAccount *CustomerBankAccount `json:"customer_bank_accounts"`
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

if result.CustomerBankAccount == nil {
    return nil, errors.New("missing result")
  }

  return result.CustomerBankAccount, nil
}


// CustomerBankAccountUpdateParams parameters
type CustomerBankAccountUpdateParams struct {
      Metadata map[string]interface{} `url:",omitempty" json:"metadata,omitempty"`
      }

// Update
// Updates a customer bank account object. Only the metadata parameter is
// allowed.
func (s *CustomerBankAccountService) Update(ctx context.Context,identity string, p CustomerBankAccountUpdateParams) (*CustomerBankAccount,error) {
  uri, err := url.Parse(fmt.Sprintf(s.endpoint + "/customer_bank_accounts/%v",
      identity,))
  if err != nil {
    return nil, err
  }

  var body io.Reader

  var buf bytes.Buffer
  err = json.NewEncoder(&buf).Encode(map[string]interface{}{
    "customer_bank_accounts": p,
  })
  if err != nil {
    return nil, err
  }
  body = &buf

  req, err := http.NewRequest("PUT", uri.String(), body)
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
CustomerBankAccount *CustomerBankAccount `json:"customer_bank_accounts"`
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

if result.CustomerBankAccount == nil {
    return nil, errors.New("missing result")
  }

  return result.CustomerBankAccount, nil
}



// Disable
// Immediately cancels all associated mandates and cancellable payments.
// 
//
// This will return a `disable_failed` error if the bank account has already
// been disabled.
// 
// A disabled bank account can be re-enabled by creating a
// new bank account resource with the same details.
func (s *CustomerBankAccountService) Disable(ctx context.Context,identity string) (*CustomerBankAccount,error) {
  uri, err := url.Parse(fmt.Sprintf(s.endpoint + "/customer_bank_accounts/%v/actions/disable",
      identity,))
  if err != nil {
    return nil, err
  }

  var body io.Reader

  

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
CustomerBankAccount *CustomerBankAccount `json:"customer_bank_accounts"`
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

if result.CustomerBankAccount == nil {
    return nil, errors.New("missing result")
  }

  return result.CustomerBankAccount, nil
}

