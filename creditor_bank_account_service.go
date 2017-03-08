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
var _ = json.NewDecoder


type CreditorBankAccountService struct {
  endpoint string
  token string
  client *http.Client
}



// CreditorBankAccountCreateParams parameters
type CreditorBankAccountCreateParams struct {
      AccountHolderName string `url:",omitempty" json:"account_holder_name,omitempty"`
        AccountNumber string `url:",omitempty" json:"account_number,omitempty"`
        BankCode string `url:",omitempty" json:"bank_code,omitempty"`
        BranchCode string `url:",omitempty" json:"branch_code,omitempty"`
        CountryCode string `url:",omitempty" json:"country_code,omitempty"`
        Currency string `url:",omitempty" json:"currency,omitempty"`
        Iban string `url:",omitempty" json:"iban,omitempty"`
        Links struct {
      Creditor string `url:",omitempty" json:"creditor,omitempty"`
        
    } `url:",omitempty" json:"links,omitempty"`
        Metadata map[string]interface{} `url:",omitempty" json:"metadata,omitempty"`
        SetAsDefaultPayoutAccount bool `url:",omitempty" json:"set_as_default_payout_account,omitempty"`
        
    }
// CreditorBankAccountCreateResult parameters
type CreditorBankAccountCreateResult struct {
      CreditorBankAccounts struct {
      AccountHolderName string `url:",omitempty" json:"account_holder_name,omitempty"`
        AccountNumberEnding string `url:",omitempty" json:"account_number_ending,omitempty"`
        BankName string `url:",omitempty" json:"bank_name,omitempty"`
        CountryCode string `url:",omitempty" json:"country_code,omitempty"`
        CreatedAt string `url:",omitempty" json:"created_at,omitempty"`
        Currency string `url:",omitempty" json:"currency,omitempty"`
        Enabled bool `url:",omitempty" json:"enabled,omitempty"`
        Id string `url:",omitempty" json:"id,omitempty"`
        Links struct {
      Creditor string `url:",omitempty" json:"creditor,omitempty"`
        
    } `url:",omitempty" json:"links,omitempty"`
        Metadata map[string]interface{} `url:",omitempty" json:"metadata,omitempty"`
        
    } `url:",omitempty" json:"creditor_bank_accounts,omitempty"`
        
    }

// Create
// Creates a new creditor bank account object.
func (s *CreditorBankAccountService) Create(ctx context.Context, p CreditorBankAccountCreateParams) (*CreditorBankAccountCreateResult, error) {
  uri, err := url.Parse(fmt.Sprintf(
      s.endpoint + "/creditor_bank_accounts",))
  if err != nil {
    return nil, err
  }

  var body io.Reader

  var buf bytes.Buffer
  err = json.NewEncoder(&buf).Encode(map[string]interface{}{
    "creditor_bank_accounts": p,
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
    *CreditorBankAccountCreateResult
  }

  try(3, func() error {
      res, err := client.Do(req)
      if err != nil {
        return err
      }
      defer res.Body.Close()

      err = responseErr(res)
      if err != nil {
        return err
      }

      return nil
  })
  if err != nil {
    return nil, err
  }

  return result.CreditorBankAccountCreateResult, nil
}


// CreditorBankAccountListParams parameters
type CreditorBankAccountListParams struct {
      After string `url:",omitempty" json:"after,omitempty"`
        Before string `url:",omitempty" json:"before,omitempty"`
        CreatedAt struct {
      Gt string `url:",omitempty" json:"gt,omitempty"`
        Gte string `url:",omitempty" json:"gte,omitempty"`
        Lt string `url:",omitempty" json:"lt,omitempty"`
        Lte string `url:",omitempty" json:"lte,omitempty"`
        
    } `url:",omitempty" json:"created_at,omitempty"`
        Creditor string `url:",omitempty" json:"creditor,omitempty"`
        Enabled bool `url:",omitempty" json:"enabled,omitempty"`
        Limit int `url:",omitempty" json:"limit,omitempty"`
        
    }
// CreditorBankAccountListResult parameters
type CreditorBankAccountListResult struct {
      CreditorBankAccounts []struct {
      AccountHolderName string `url:",omitempty" json:"account_holder_name,omitempty"`
        AccountNumberEnding string `url:",omitempty" json:"account_number_ending,omitempty"`
        BankName string `url:",omitempty" json:"bank_name,omitempty"`
        CountryCode string `url:",omitempty" json:"country_code,omitempty"`
        CreatedAt string `url:",omitempty" json:"created_at,omitempty"`
        Currency string `url:",omitempty" json:"currency,omitempty"`
        Enabled bool `url:",omitempty" json:"enabled,omitempty"`
        Id string `url:",omitempty" json:"id,omitempty"`
        Links struct {
      Creditor string `url:",omitempty" json:"creditor,omitempty"`
        
    } `url:",omitempty" json:"links,omitempty"`
        Metadata map[string]interface{} `url:",omitempty" json:"metadata,omitempty"`
        
    } `url:",omitempty" json:"creditor_bank_accounts,omitempty"`
        Meta struct {
      Cursors struct {
      After string `url:",omitempty" json:"after,omitempty"`
        Before string `url:",omitempty" json:"before,omitempty"`
        
    } `url:",omitempty" json:"cursors,omitempty"`
        Limit int `url:",omitempty" json:"limit,omitempty"`
        
    } `url:",omitempty" json:"meta,omitempty"`
        
    }

// List
// Returns a [cursor-paginated](#api-usage-cursor-pagination) list of your
// creditor bank accounts.
func (s *CreditorBankAccountService) List(ctx context.Context, p CreditorBankAccountListParams) (*CreditorBankAccountListResult, error) {
  uri, err := url.Parse(fmt.Sprintf(
      s.endpoint + "/creditor_bank_accounts",))
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
    *CreditorBankAccountListResult
  }

  try(3, func() error {
      res, err := client.Do(req)
      if err != nil {
        return err
      }
      defer res.Body.Close()

      err = responseErr(res)
      if err != nil {
        return err
      }

      return nil
  })
  if err != nil {
    return nil, err
  }

  return result.CreditorBankAccountListResult, nil
}


// CreditorBankAccountGetResult parameters
type CreditorBankAccountGetResult struct {
      CreditorBankAccounts struct {
      AccountHolderName string `url:",omitempty" json:"account_holder_name,omitempty"`
        AccountNumberEnding string `url:",omitempty" json:"account_number_ending,omitempty"`
        BankName string `url:",omitempty" json:"bank_name,omitempty"`
        CountryCode string `url:",omitempty" json:"country_code,omitempty"`
        CreatedAt string `url:",omitempty" json:"created_at,omitempty"`
        Currency string `url:",omitempty" json:"currency,omitempty"`
        Enabled bool `url:",omitempty" json:"enabled,omitempty"`
        Id string `url:",omitempty" json:"id,omitempty"`
        Links struct {
      Creditor string `url:",omitempty" json:"creditor,omitempty"`
        
    } `url:",omitempty" json:"links,omitempty"`
        Metadata map[string]interface{} `url:",omitempty" json:"metadata,omitempty"`
        
    } `url:",omitempty" json:"creditor_bank_accounts,omitempty"`
        
    }

// Get
// Retrieves the details of an existing creditor bank account.
func (s *CreditorBankAccountService) Get(ctx context.Context,identity string) (*CreditorBankAccountGetResult, error) {
  uri, err := url.Parse(fmt.Sprintf(
      s.endpoint + "/creditor_bank_accounts/%v",
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
    *CreditorBankAccountGetResult
  }

  try(3, func() error {
      res, err := client.Do(req)
      if err != nil {
        return err
      }
      defer res.Body.Close()

      err = responseErr(res)
      if err != nil {
        return err
      }

      return nil
  })
  if err != nil {
    return nil, err
  }

  return result.CreditorBankAccountGetResult, nil
}


// CreditorBankAccountDisableResult parameters
type CreditorBankAccountDisableResult struct {
      CreditorBankAccounts struct {
      AccountHolderName string `url:",omitempty" json:"account_holder_name,omitempty"`
        AccountNumberEnding string `url:",omitempty" json:"account_number_ending,omitempty"`
        BankName string `url:",omitempty" json:"bank_name,omitempty"`
        CountryCode string `url:",omitempty" json:"country_code,omitempty"`
        CreatedAt string `url:",omitempty" json:"created_at,omitempty"`
        Currency string `url:",omitempty" json:"currency,omitempty"`
        Enabled bool `url:",omitempty" json:"enabled,omitempty"`
        Id string `url:",omitempty" json:"id,omitempty"`
        Links struct {
      Creditor string `url:",omitempty" json:"creditor,omitempty"`
        
    } `url:",omitempty" json:"links,omitempty"`
        Metadata map[string]interface{} `url:",omitempty" json:"metadata,omitempty"`
        
    } `url:",omitempty" json:"creditor_bank_accounts,omitempty"`
        
    }

// Disable
// Immediately disables the bank account, no money can be paid out to a disabled
// account.
// 
// This will return a `disable_failed` error if the bank account
// has already been disabled.
// 
// A disabled bank account can be re-enabled
// by creating a new bank account resource with the same details.
func (s *CreditorBankAccountService) Disable(ctx context.Context,identity string) (*CreditorBankAccountDisableResult, error) {
  uri, err := url.Parse(fmt.Sprintf(
      s.endpoint + "/creditor_bank_accounts/%v/actions/disable",
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
    *CreditorBankAccountDisableResult
  }

  try(3, func() error {
      res, err := client.Do(req)
      if err != nil {
        return err
      }
      defer res.Body.Close()

      err = responseErr(res)
      if err != nil {
        return err
      }

      return nil
  })
  if err != nil {
    return nil, err
  }

  return result.CreditorBankAccountDisableResult, nil
}

