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


type RefundService struct {
  endpoint string
  token string
  client *http.Client
}



// RefundCreateParams parameters
type RefundCreateParams struct {
      Amount string `json:"amount,omitempty"`
        Links struct {
      Payment string `json:"payment,omitempty"`
        
    } `json:"links,omitempty"`
        Metadata map[string]interface{} `json:"metadata,omitempty"`
        Reference string `json:"reference,omitempty"`
        TotalAmountConfirmation string `json:"total_amount_confirmation,omitempty"`
        
    }
// RefundCreateResult parameters
type RefundCreateResult struct {
      Refunds struct {
      Amount string `json:"amount,omitempty"`
        CreatedAt string `json:"created_at,omitempty"`
        Currency string `json:"currency,omitempty"`
        Id string `json:"id,omitempty"`
        Links struct {
      Payment string `json:"payment,omitempty"`
        
    } `json:"links,omitempty"`
        Metadata map[string]interface{} `json:"metadata,omitempty"`
        Reference string `json:"reference,omitempty"`
        
    } `json:"refunds,omitempty"`
        
    }

// Create
// Creates a new refund object.
// 
// This fails with:<a
// name="refund_payment_invalid_state"></a><a
// name="total_amount_confirmation_invalid"></a><a
// name="number_of_refunds_exceeded"></a>
// 
// -
// `refund_payment_invalid_state` error if the linked
// [payment](#core-endpoints-payments) isn't either `confirmed` or
// `paid_out`.
// 
// - `total_amount_confirmation_invalid` if the confirmation
// amount doesn't match the total amount refunded for the payment. This
// safeguard is there to prevent two processes from creating refunds without
// awareness of each other.
// 
// - `number_of_refunds_exceeded` if five or
// more refunds have already been created against the payment.
// 
func (s *RefundService) Create(
  ctx context.Context,
  p RefundCreateParams) (*RefundCreateResult, error) {
  uri, err := url.Parse(fmt.Sprintf(
      s.endpoint + "/refunds",))
  if err != nil {
    return nil, err
  }

  var body io.Reader

  var buf bytes.Buffer
  err = json.NewEncoder(&buf).Encode(map[string]interface{}{
    "refunds": p,
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
    *RefundCreateResult
    Err *APIError `json:"error"`
  }

  err = json.NewDecoder(res.Body).Decode(&result)
  if err != nil {
    return nil, err
  }

  if result.Err != nil {
    return nil, result.Err
  }

  return result.RefundCreateResult, nil
}


// RefundListParams parameters
type RefundListParams struct {
      After string `json:"after,omitempty"`
        Before string `json:"before,omitempty"`
        CreatedAt struct {
      Gt string `json:"gt,omitempty"`
        Gte string `json:"gte,omitempty"`
        Lt string `json:"lt,omitempty"`
        Lte string `json:"lte,omitempty"`
        
    } `json:"created_at,omitempty"`
        Limit string `json:"limit,omitempty"`
        Payment string `json:"payment,omitempty"`
        
    }
// RefundListResult parameters
type RefundListResult struct {
      Meta struct {
      Cursors struct {
      After string `json:"after,omitempty"`
        Before string `json:"before,omitempty"`
        
    } `json:"cursors,omitempty"`
        Limit int `json:"limit,omitempty"`
        
    } `json:"meta,omitempty"`
        Refunds []struct {
      Amount string `json:"amount,omitempty"`
        CreatedAt string `json:"created_at,omitempty"`
        Currency string `json:"currency,omitempty"`
        Id string `json:"id,omitempty"`
        Links struct {
      Payment string `json:"payment,omitempty"`
        
    } `json:"links,omitempty"`
        Metadata map[string]interface{} `json:"metadata,omitempty"`
        Reference string `json:"reference,omitempty"`
        
    } `json:"refunds,omitempty"`
        
    }

// List
// Returns a [cursor-paginated](#api-usage-cursor-pagination) list of your
// refunds.
func (s *RefundService) List(
  ctx context.Context,
  p RefundListParams) (*RefundListResult, error) {
  uri, err := url.Parse(fmt.Sprintf(
      s.endpoint + "/refunds",))
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
    *RefundListResult
    Err *APIError `json:"error"`
  }

  err = json.NewDecoder(res.Body).Decode(&result)
  if err != nil {
    return nil, err
  }

  if result.Err != nil {
    return nil, result.Err
  }

  return result.RefundListResult, nil
}


// RefundGetResult parameters
type RefundGetResult struct {
      Refunds struct {
      Amount string `json:"amount,omitempty"`
        CreatedAt string `json:"created_at,omitempty"`
        Currency string `json:"currency,omitempty"`
        Id string `json:"id,omitempty"`
        Links struct {
      Payment string `json:"payment,omitempty"`
        
    } `json:"links,omitempty"`
        Metadata map[string]interface{} `json:"metadata,omitempty"`
        Reference string `json:"reference,omitempty"`
        
    } `json:"refunds,omitempty"`
        
    }

// Get
// Retrieves all details for a single refund
func (s *RefundService) Get(
  ctx context.Context,
  identity string) (*RefundGetResult, error) {
  uri, err := url.Parse(fmt.Sprintf(
      s.endpoint + "/refunds/%v",
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
    *RefundGetResult
    Err *APIError `json:"error"`
  }

  err = json.NewDecoder(res.Body).Decode(&result)
  if err != nil {
    return nil, err
  }

  if result.Err != nil {
    return nil, result.Err
  }

  return result.RefundGetResult, nil
}


// RefundUpdateParams parameters
type RefundUpdateParams struct {
      Metadata map[string]interface{} `json:"metadata,omitempty"`
        
    }
// RefundUpdateResult parameters
type RefundUpdateResult struct {
      Refunds struct {
      Amount string `json:"amount,omitempty"`
        CreatedAt string `json:"created_at,omitempty"`
        Currency string `json:"currency,omitempty"`
        Id string `json:"id,omitempty"`
        Links struct {
      Payment string `json:"payment,omitempty"`
        
    } `json:"links,omitempty"`
        Metadata map[string]interface{} `json:"metadata,omitempty"`
        Reference string `json:"reference,omitempty"`
        
    } `json:"refunds,omitempty"`
        
    }

// Update
// Updates a refund object.
func (s *RefundService) Update(
  ctx context.Context,
  identity string,
  p RefundUpdateParams) (*RefundUpdateResult, error) {
  uri, err := url.Parse(fmt.Sprintf(
      s.endpoint + "/refunds/%v",
      identity,))
  if err != nil {
    return nil, err
  }

  var body io.Reader

  var buf bytes.Buffer
  err = json.NewEncoder(&buf).Encode(map[string]interface{}{
    "refunds": p,
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
    *RefundUpdateResult
    Err *APIError `json:"error"`
  }

  err = json.NewDecoder(res.Body).Decode(&result)
  if err != nil {
    return nil, err
  }

  if result.Err != nil {
    return nil, result.Err
  }

  return result.RefundUpdateResult, nil
}

