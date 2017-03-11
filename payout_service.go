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


// PayoutService manages payouts
type PayoutService struct {
  endpoint string
  token string
  client *http.Client
}


// Payout model
type Payout struct {
      Amount int `url:",omitempty" json:"amount,omitempty"`
      ArrivalDate string `url:",omitempty" json:"arrival_date,omitempty"`
      CreatedAt string `url:",omitempty" json:"created_at,omitempty"`
      Currency string `url:",omitempty" json:"currency,omitempty"`
      DeductedFees int `url:",omitempty" json:"deducted_fees,omitempty"`
      Id string `url:",omitempty" json:"id,omitempty"`
      Links struct {
      Creditor string `url:",omitempty" json:"creditor,omitempty"`
      CreditorBankAccount string `url:",omitempty" json:"creditor_bank_account,omitempty"`
      } `url:",omitempty" json:"links,omitempty"`
      Reference string `url:",omitempty" json:"reference,omitempty"`
      Status string `url:",omitempty" json:"status,omitempty"`
      }




// PayoutListParams parameters
type PayoutListParams struct {
      After string `url:",omitempty" json:"after,omitempty"`
      Before string `url:",omitempty" json:"before,omitempty"`
      CreatedAt struct {
      Gt string `url:",omitempty" json:"gt,omitempty"`
      Gte string `url:",omitempty" json:"gte,omitempty"`
      Lt string `url:",omitempty" json:"lt,omitempty"`
      Lte string `url:",omitempty" json:"lte,omitempty"`
      } `url:",omitempty" json:"created_at,omitempty"`
      Creditor string `url:",omitempty" json:"creditor,omitempty"`
      CreditorBankAccount string `url:",omitempty" json:"creditor_bank_account,omitempty"`
      Currency string `url:",omitempty" json:"currency,omitempty"`
      Limit int `url:",omitempty" json:"limit,omitempty"`
      Status string `url:",omitempty" json:"status,omitempty"`
      }// PayoutListResult response including pagination metadata
type PayoutListResult struct {
      Meta struct {
      Cursors struct {
      After string `url:",omitempty" json:"after,omitempty"`
      Before string `url:",omitempty" json:"before,omitempty"`
      } `url:",omitempty" json:"cursors,omitempty"`
      Limit int `url:",omitempty" json:"limit,omitempty"`
      } `url:",omitempty" json:"meta,omitempty"`
      Payouts []struct {
      Amount int `url:",omitempty" json:"amount,omitempty"`
      ArrivalDate string `url:",omitempty" json:"arrival_date,omitempty"`
      CreatedAt string `url:",omitempty" json:"created_at,omitempty"`
      Currency string `url:",omitempty" json:"currency,omitempty"`
      DeductedFees int `url:",omitempty" json:"deducted_fees,omitempty"`
      Id string `url:",omitempty" json:"id,omitempty"`
      Links struct {
      Creditor string `url:",omitempty" json:"creditor,omitempty"`
      CreditorBankAccount string `url:",omitempty" json:"creditor_bank_account,omitempty"`
      } `url:",omitempty" json:"links,omitempty"`
      Reference string `url:",omitempty" json:"reference,omitempty"`
      Status string `url:",omitempty" json:"status,omitempty"`
      } `url:",omitempty" json:"payouts,omitempty"`
      }


// List
// Returns a [cursor-paginated](#api-usage-cursor-pagination) list of your
// payouts.
func (s *PayoutService) List(ctx context.Context, p PayoutListParams) (*PayoutListResult,error) {
  uri, err := url.Parse(fmt.Sprintf(s.endpoint + "/payouts",))
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
*PayoutListResult
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

if result.PayoutListResult == nil {
    return nil, errors.New("missing result")
  }

  return result.PayoutListResult, nil
}



// Get
// Retrieves the details of a single payout. For an example of how to reconcile
// the transactions in a payout, see [this
// guide](#events-reconciling-payouts-with-events).
func (s *PayoutService) Get(ctx context.Context,identity string) (*Payout,error) {
  uri, err := url.Parse(fmt.Sprintf(s.endpoint + "/payouts/%v",
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
Payout *Payout `json:"payouts"`
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

if result.Payout == nil {
    return nil, errors.New("missing result")
  }

  return result.Payout, nil
}

