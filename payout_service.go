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


type PayoutService struct {
  endpoint string
  token string
  client *http.Client
}



// PayoutListParams parameters
type PayoutListParams struct {
      After string `json:"after,omitempty"`
        Before string `json:"before,omitempty"`
        CreatedAt struct {
      Gt string `json:"gt,omitempty"`
        Gte string `json:"gte,omitempty"`
        Lt string `json:"lt,omitempty"`
        Lte string `json:"lte,omitempty"`
        
    } `json:"created_at,omitempty"`
        Creditor string `json:"creditor,omitempty"`
        CreditorBankAccount string `json:"creditor_bank_account,omitempty"`
        Currency string `json:"currency,omitempty"`
        Limit string `json:"limit,omitempty"`
        Status string `json:"status,omitempty"`
        
    }
// PayoutListResult parameters
type PayoutListResult struct {
      Meta struct {
      Cursors struct {
      After string `json:"after,omitempty"`
        Before string `json:"before,omitempty"`
        
    } `json:"cursors,omitempty"`
        Limit int `json:"limit,omitempty"`
        
    } `json:"meta,omitempty"`
        Payouts []struct {
      Amount string `json:"amount,omitempty"`
        ArrivalDate string `json:"arrival_date,omitempty"`
        CreatedAt string `json:"created_at,omitempty"`
        Currency string `json:"currency,omitempty"`
        DeductedFees string `json:"deducted_fees,omitempty"`
        Id string `json:"id,omitempty"`
        Links struct {
      Creditor string `json:"creditor,omitempty"`
        CreditorBankAccount string `json:"creditor_bank_account,omitempty"`
        
    } `json:"links,omitempty"`
        Reference string `json:"reference,omitempty"`
        Status string `json:"status,omitempty"`
        
    } `json:"payouts,omitempty"`
        
    }

// List
// Returns a [cursor-paginated](#api-usage-cursor-pagination) list of your
// payouts.
func (s *PayoutService) List(
  ctx context.Context,
  p PayoutListParams) (*PayoutListResult, error) {
  uri, err := url.Parse(fmt.Sprintf(
      s.endpoint + "/payouts",))
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
    *PayoutListResult
    Err *APIError `json:"error"`
  }

  err = json.NewDecoder(res.Body).Decode(&result)
  if err != nil {
    return nil, err
  }

  if result.Err != nil {
    return nil, result.Err
  }

  return result.PayoutListResult, nil
}


// PayoutGetResult parameters
type PayoutGetResult struct {
      Payouts struct {
      Amount string `json:"amount,omitempty"`
        ArrivalDate string `json:"arrival_date,omitempty"`
        CreatedAt string `json:"created_at,omitempty"`
        Currency string `json:"currency,omitempty"`
        DeductedFees string `json:"deducted_fees,omitempty"`
        Id string `json:"id,omitempty"`
        Links struct {
      Creditor string `json:"creditor,omitempty"`
        CreditorBankAccount string `json:"creditor_bank_account,omitempty"`
        
    } `json:"links,omitempty"`
        Reference string `json:"reference,omitempty"`
        Status string `json:"status,omitempty"`
        
    } `json:"payouts,omitempty"`
        
    }

// Get
// Retrieves the details of a single payout. For an example of how to reconcile
// the transactions in a payout, see [this
// guide](#events-reconciling-payouts-with-events).
func (s *PayoutService) Get(
  ctx context.Context,
  identity string) (*PayoutGetResult, error) {
  uri, err := url.Parse(fmt.Sprintf(
      s.endpoint + "/payouts/%v",
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
    *PayoutGetResult
    Err *APIError `json:"error"`
  }

  err = json.NewDecoder(res.Body).Decode(&result)
  if err != nil {
    return nil, err
  }

  if result.Err != nil {
    return nil, result.Err
  }

  return result.PayoutGetResult, nil
}

