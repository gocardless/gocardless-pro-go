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


type CustomerBankAccountService struct {
  endpoint string
  token string
  client *http.Client
}



// CustomerBankAccountCreateParams parameters
type CustomerBankAccountCreateParams struct {
      AccountHolderName string `json:"account_holder_name,omitempty"`
        AccountNumber string `json:"account_number,omitempty"`
        BankCode string `json:"bank_code,omitempty"`
        BranchCode string `json:"branch_code,omitempty"`
        CountryCode string `json:"country_code,omitempty"`
        Currency string `json:"currency,omitempty"`
        Iban string `json:"iban,omitempty"`
        Links struct {
      Customer string `json:"customer,omitempty"`
        CustomerBankAccountToken string `json:"customer_bank_account_token,omitempty"`
        
    } `json:"links,omitempty"`
        Metadata map[string]interface{} `json:"metadata,omitempty"`
        
    }
// CustomerBankAccountCreateResult parameters
type CustomerBankAccountCreateResult struct {
      CustomerBankAccounts struct {
      AccountHolderName string `json:"account_holder_name,omitempty"`
        AccountNumberEnding string `json:"account_number_ending,omitempty"`
        BankName string `json:"bank_name,omitempty"`
        CountryCode string `json:"country_code,omitempty"`
        CreatedAt string `json:"created_at,omitempty"`
        Currency string `json:"currency,omitempty"`
        Enabled bool `json:"enabled,omitempty"`
        Id string `json:"id,omitempty"`
        Links struct {
      Customer string `json:"customer,omitempty"`
        
    } `json:"links,omitempty"`
        Metadata map[string]interface{} `json:"metadata,omitempty"`
        
    } `json:"customer_bank_accounts,omitempty"`
        
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
func (s *CustomerBankAccountService) Create(
  ctx context.Context,
  p CustomerBankAccountCreateParams) (*CustomerBankAccountCreateResult, error) {
  uri, err := url.Parse(fmt.Sprintf(
      s.endpoint + "/customer_bank_accounts",))
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
    *CustomerBankAccountCreateResult
    Err *APIError `json:"error"`
  }

  err = json.NewDecoder(res.Body).Decode(&result)
  if err != nil {
    return nil, err
  }

  if result.Err != nil {
    return nil, result.Err
  }

  return result.CustomerBankAccountCreateResult, nil
}


// CustomerBankAccountListParams parameters
type CustomerBankAccountListParams struct {
      After string `json:"after,omitempty"`
        Before string `json:"before,omitempty"`
        CreatedAt struct {
      Gt string `json:"gt,omitempty"`
        Gte string `json:"gte,omitempty"`
        Lt string `json:"lt,omitempty"`
        Lte string `json:"lte,omitempty"`
        
    } `json:"created_at,omitempty"`
        Customer string `json:"customer,omitempty"`
        Enabled bool `json:"enabled,omitempty"`
        Limit string `json:"limit,omitempty"`
        
    }
// CustomerBankAccountListResult parameters
type CustomerBankAccountListResult struct {
      CustomerBankAccounts []struct {
      AccountHolderName string `json:"account_holder_name,omitempty"`
        AccountNumberEnding string `json:"account_number_ending,omitempty"`
        BankName string `json:"bank_name,omitempty"`
        CountryCode string `json:"country_code,omitempty"`
        CreatedAt string `json:"created_at,omitempty"`
        Currency string `json:"currency,omitempty"`
        Enabled bool `json:"enabled,omitempty"`
        Id string `json:"id,omitempty"`
        Links struct {
      Customer string `json:"customer,omitempty"`
        
    } `json:"links,omitempty"`
        Metadata map[string]interface{} `json:"metadata,omitempty"`
        
    } `json:"customer_bank_accounts,omitempty"`
        Meta struct {
      Cursors struct {
      After string `json:"after,omitempty"`
        Before string `json:"before,omitempty"`
        
    } `json:"cursors,omitempty"`
        Limit int `json:"limit,omitempty"`
        
    } `json:"meta,omitempty"`
        
    }

// List
// Returns a [cursor-paginated](#api-usage-cursor-pagination) list of your bank
// accounts.
func (s *CustomerBankAccountService) List(
  ctx context.Context,
  p CustomerBankAccountListParams) (*CustomerBankAccountListResult, error) {
  uri, err := url.Parse(fmt.Sprintf(
      s.endpoint + "/customer_bank_accounts",))
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

  res, err := client.Do(req)
  if err != nil {
    return nil, err
  }
  defer res.Body.Close()

  var result struct {
    *CustomerBankAccountListResult
    Err *APIError `json:"error"`
  }

  err = json.NewDecoder(res.Body).Decode(&result)
  if err != nil {
    return nil, err
  }

  if result.Err != nil {
    return nil, result.Err
  }

  return result.CustomerBankAccountListResult, nil
}


// CustomerBankAccountGetResult parameters
type CustomerBankAccountGetResult struct {
      CustomerBankAccounts struct {
      AccountHolderName string `json:"account_holder_name,omitempty"`
        AccountNumberEnding string `json:"account_number_ending,omitempty"`
        BankName string `json:"bank_name,omitempty"`
        CountryCode string `json:"country_code,omitempty"`
        CreatedAt string `json:"created_at,omitempty"`
        Currency string `json:"currency,omitempty"`
        Enabled bool `json:"enabled,omitempty"`
        Id string `json:"id,omitempty"`
        Links struct {
      Customer string `json:"customer,omitempty"`
        
    } `json:"links,omitempty"`
        Metadata map[string]interface{} `json:"metadata,omitempty"`
        
    } `json:"customer_bank_accounts,omitempty"`
        
    }

// Get
// Retrieves the details of an existing bank account.
func (s *CustomerBankAccountService) Get(
  ctx context.Context,
  identity string) (*CustomerBankAccountGetResult, error) {
  uri, err := url.Parse(fmt.Sprintf(
      s.endpoint + "/customer_bank_accounts/%v",
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

  res, err := client.Do(req)
  if err != nil {
    return nil, err
  }
  defer res.Body.Close()

  var result struct {
    *CustomerBankAccountGetResult
    Err *APIError `json:"error"`
  }

  err = json.NewDecoder(res.Body).Decode(&result)
  if err != nil {
    return nil, err
  }

  if result.Err != nil {
    return nil, result.Err
  }

  return result.CustomerBankAccountGetResult, nil
}


// CustomerBankAccountUpdateParams parameters
type CustomerBankAccountUpdateParams struct {
      Metadata map[string]interface{} `json:"metadata,omitempty"`
        
    }
// CustomerBankAccountUpdateResult parameters
type CustomerBankAccountUpdateResult struct {
      CustomerBankAccounts struct {
      AccountHolderName string `json:"account_holder_name,omitempty"`
        AccountNumberEnding string `json:"account_number_ending,omitempty"`
        BankName string `json:"bank_name,omitempty"`
        CountryCode string `json:"country_code,omitempty"`
        CreatedAt string `json:"created_at,omitempty"`
        Currency string `json:"currency,omitempty"`
        Enabled bool `json:"enabled,omitempty"`
        Id string `json:"id,omitempty"`
        Links struct {
      Customer string `json:"customer,omitempty"`
        
    } `json:"links,omitempty"`
        Metadata map[string]interface{} `json:"metadata,omitempty"`
        
    } `json:"customer_bank_accounts,omitempty"`
        
    }

// Update
// Updates a customer bank account object. Only the metadata parameter is
// allowed.
func (s *CustomerBankAccountService) Update(
  ctx context.Context,
  identity string,
  p CustomerBankAccountUpdateParams) (*CustomerBankAccountUpdateResult, error) {
  uri, err := url.Parse(fmt.Sprintf(
      s.endpoint + "/customer_bank_accounts/%v",
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
    *CustomerBankAccountUpdateResult
    Err *APIError `json:"error"`
  }

  err = json.NewDecoder(res.Body).Decode(&result)
  if err != nil {
    return nil, err
  }

  if result.Err != nil {
    return nil, result.Err
  }

  return result.CustomerBankAccountUpdateResult, nil
}


// CustomerBankAccountDisableResult parameters
type CustomerBankAccountDisableResult struct {
      CustomerBankAccounts struct {
      AccountHolderName string `json:"account_holder_name,omitempty"`
        AccountNumberEnding string `json:"account_number_ending,omitempty"`
        BankName string `json:"bank_name,omitempty"`
        CountryCode string `json:"country_code,omitempty"`
        CreatedAt string `json:"created_at,omitempty"`
        Currency string `json:"currency,omitempty"`
        Enabled bool `json:"enabled,omitempty"`
        Id string `json:"id,omitempty"`
        Links struct {
      Customer string `json:"customer,omitempty"`
        
    } `json:"links,omitempty"`
        Metadata map[string]interface{} `json:"metadata,omitempty"`
        
    } `json:"customer_bank_accounts,omitempty"`
        
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
func (s *CustomerBankAccountService) Disable(
  ctx context.Context,
  identity string) (*CustomerBankAccountDisableResult, error) {
  uri, err := url.Parse(fmt.Sprintf(
      s.endpoint + "/customer_bank_accounts/%v/actions/disable",
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
    *CustomerBankAccountDisableResult
    Err *APIError `json:"error"`
  }

  err = json.NewDecoder(res.Body).Decode(&result)
  if err != nil {
    return nil, err
  }

  if result.Err != nil {
    return nil, result.Err
  }

  return result.CustomerBankAccountDisableResult, nil
}

