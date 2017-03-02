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


type PaymentService struct {
  endpoint string
  token string
  client *http.Client
}



// PaymentCreateParams parameters
type PaymentCreateParams struct {
      Amount string `json:"amount,omitempty"`
        AppFee string `json:"app_fee,omitempty"`
        ChargeDate string `json:"charge_date,omitempty"`
        Currency string `json:"currency,omitempty"`
        Description string `json:"description,omitempty"`
        Links struct {
      Mandate string `json:"mandate,omitempty"`
        
    } `json:"links,omitempty"`
        Metadata map[string]interface{} `json:"metadata,omitempty"`
        Reference string `json:"reference,omitempty"`
        
    }
// PaymentCreateResult parameters
type PaymentCreateResult struct {
      Payments struct {
      Amount string `json:"amount,omitempty"`
        AmountRefunded string `json:"amount_refunded,omitempty"`
        ChargeDate string `json:"charge_date,omitempty"`
        CreatedAt string `json:"created_at,omitempty"`
        Currency string `json:"currency,omitempty"`
        Description string `json:"description,omitempty"`
        Id string `json:"id,omitempty"`
        Links struct {
      Creditor string `json:"creditor,omitempty"`
        Mandate string `json:"mandate,omitempty"`
        Payout string `json:"payout,omitempty"`
        Subscription string `json:"subscription,omitempty"`
        
    } `json:"links,omitempty"`
        Metadata map[string]interface{} `json:"metadata,omitempty"`
        Reference string `json:"reference,omitempty"`
        Status string `json:"status,omitempty"`
        
    } `json:"payments,omitempty"`
        
    }

// Create
// <a name="mandate_is_inactive"></a>Creates a new payment object.
// 
// This
// fails with a `mandate_is_inactive` error if the linked
// [mandate](#core-endpoints-mandates) is cancelled or has failed. Payments can
// be created against mandates with status of: `pending_customer_approval`,
// `pending_submission`, `submitted`, and `active`.
func (s *PaymentService) Create(
  ctx context.Context,
  p PaymentCreateParams) (*PaymentCreateResult, error) {
  uri, err := url.Parse(fmt.Sprintf(
      s.endpoint + "/payments",))
  if err != nil {
    return nil, err
  }

  var body io.Reader

  var buf bytes.Buffer
  err = json.NewEncoder(&buf).Encode(map[string]interface{}{
    "payments": p,
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
    *PaymentCreateResult
    Err *APIError `json:"error"`
  }

  err = json.NewDecoder(res.Body).Decode(&result)
  if err != nil {
    return nil, err
  }

  if result.Err != nil {
    return nil, result.Err
  }

  return result.PaymentCreateResult, nil
}


// PaymentListParams parameters
type PaymentListParams struct {
      After string `json:"after,omitempty"`
        Before string `json:"before,omitempty"`
        CreatedAt struct {
      Gt string `json:"gt,omitempty"`
        Gte string `json:"gte,omitempty"`
        Lt string `json:"lt,omitempty"`
        Lte string `json:"lte,omitempty"`
        
    } `json:"created_at,omitempty"`
        Creditor string `json:"creditor,omitempty"`
        Currency string `json:"currency,omitempty"`
        Customer string `json:"customer,omitempty"`
        Limit string `json:"limit,omitempty"`
        Mandate string `json:"mandate,omitempty"`
        Status string `json:"status,omitempty"`
        Subscription string `json:"subscription,omitempty"`
        
    }
// PaymentListResult parameters
type PaymentListResult struct {
      Meta struct {
      Cursors struct {
      After string `json:"after,omitempty"`
        Before string `json:"before,omitempty"`
        
    } `json:"cursors,omitempty"`
        Limit int `json:"limit,omitempty"`
        
    } `json:"meta,omitempty"`
        Payments []struct {
      Amount string `json:"amount,omitempty"`
        AmountRefunded string `json:"amount_refunded,omitempty"`
        ChargeDate string `json:"charge_date,omitempty"`
        CreatedAt string `json:"created_at,omitempty"`
        Currency string `json:"currency,omitempty"`
        Description string `json:"description,omitempty"`
        Id string `json:"id,omitempty"`
        Links struct {
      Creditor string `json:"creditor,omitempty"`
        Mandate string `json:"mandate,omitempty"`
        Payout string `json:"payout,omitempty"`
        Subscription string `json:"subscription,omitempty"`
        
    } `json:"links,omitempty"`
        Metadata map[string]interface{} `json:"metadata,omitempty"`
        Reference string `json:"reference,omitempty"`
        Status string `json:"status,omitempty"`
        
    } `json:"payments,omitempty"`
        
    }

// List
// Returns a [cursor-paginated](#api-usage-cursor-pagination) list of your
// payments.
func (s *PaymentService) List(
  ctx context.Context,
  p PaymentListParams) (*PaymentListResult, error) {
  uri, err := url.Parse(fmt.Sprintf(
      s.endpoint + "/payments",))
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
    *PaymentListResult
    Err *APIError `json:"error"`
  }

  err = json.NewDecoder(res.Body).Decode(&result)
  if err != nil {
    return nil, err
  }

  if result.Err != nil {
    return nil, result.Err
  }

  return result.PaymentListResult, nil
}


// PaymentGetResult parameters
type PaymentGetResult struct {
      Payments struct {
      Amount string `json:"amount,omitempty"`
        AmountRefunded string `json:"amount_refunded,omitempty"`
        ChargeDate string `json:"charge_date,omitempty"`
        CreatedAt string `json:"created_at,omitempty"`
        Currency string `json:"currency,omitempty"`
        Description string `json:"description,omitempty"`
        Id string `json:"id,omitempty"`
        Links struct {
      Creditor string `json:"creditor,omitempty"`
        Mandate string `json:"mandate,omitempty"`
        Payout string `json:"payout,omitempty"`
        Subscription string `json:"subscription,omitempty"`
        
    } `json:"links,omitempty"`
        Metadata map[string]interface{} `json:"metadata,omitempty"`
        Reference string `json:"reference,omitempty"`
        Status string `json:"status,omitempty"`
        
    } `json:"payments,omitempty"`
        
    }

// Get
// Retrieves the details of a single existing payment.
func (s *PaymentService) Get(
  ctx context.Context,
  identity string) (*PaymentGetResult, error) {
  uri, err := url.Parse(fmt.Sprintf(
      s.endpoint + "/payments/%v",
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
    *PaymentGetResult
    Err *APIError `json:"error"`
  }

  err = json.NewDecoder(res.Body).Decode(&result)
  if err != nil {
    return nil, err
  }

  if result.Err != nil {
    return nil, result.Err
  }

  return result.PaymentGetResult, nil
}


// PaymentUpdateParams parameters
type PaymentUpdateParams struct {
      Metadata map[string]interface{} `json:"metadata,omitempty"`
        
    }
// PaymentUpdateResult parameters
type PaymentUpdateResult struct {
      Payments struct {
      Amount string `json:"amount,omitempty"`
        AmountRefunded string `json:"amount_refunded,omitempty"`
        ChargeDate string `json:"charge_date,omitempty"`
        CreatedAt string `json:"created_at,omitempty"`
        Currency string `json:"currency,omitempty"`
        Description string `json:"description,omitempty"`
        Id string `json:"id,omitempty"`
        Links struct {
      Creditor string `json:"creditor,omitempty"`
        Mandate string `json:"mandate,omitempty"`
        Payout string `json:"payout,omitempty"`
        Subscription string `json:"subscription,omitempty"`
        
    } `json:"links,omitempty"`
        Metadata map[string]interface{} `json:"metadata,omitempty"`
        Reference string `json:"reference,omitempty"`
        Status string `json:"status,omitempty"`
        
    } `json:"payments,omitempty"`
        
    }

// Update
// Updates a payment object. This accepts only the metadata parameter.
func (s *PaymentService) Update(
  ctx context.Context,
  identity string,
  p PaymentUpdateParams) (*PaymentUpdateResult, error) {
  uri, err := url.Parse(fmt.Sprintf(
      s.endpoint + "/payments/%v",
      identity,))
  if err != nil {
    return nil, err
  }

  var body io.Reader

  var buf bytes.Buffer
  err = json.NewEncoder(&buf).Encode(map[string]interface{}{
    "payments": p,
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
    *PaymentUpdateResult
    Err *APIError `json:"error"`
  }

  err = json.NewDecoder(res.Body).Decode(&result)
  if err != nil {
    return nil, err
  }

  if result.Err != nil {
    return nil, result.Err
  }

  return result.PaymentUpdateResult, nil
}


// PaymentCancelParams parameters
type PaymentCancelParams struct {
      Metadata map[string]interface{} `json:"metadata,omitempty"`
        
    }
// PaymentCancelResult parameters
type PaymentCancelResult struct {
      Payments struct {
      Amount string `json:"amount,omitempty"`
        AmountRefunded string `json:"amount_refunded,omitempty"`
        ChargeDate string `json:"charge_date,omitempty"`
        CreatedAt string `json:"created_at,omitempty"`
        Currency string `json:"currency,omitempty"`
        Description string `json:"description,omitempty"`
        Id string `json:"id,omitempty"`
        Links struct {
      Creditor string `json:"creditor,omitempty"`
        Mandate string `json:"mandate,omitempty"`
        Payout string `json:"payout,omitempty"`
        Subscription string `json:"subscription,omitempty"`
        
    } `json:"links,omitempty"`
        Metadata map[string]interface{} `json:"metadata,omitempty"`
        Reference string `json:"reference,omitempty"`
        Status string `json:"status,omitempty"`
        
    } `json:"payments,omitempty"`
        
    }

// Cancel
// Cancels the payment if it has not already been submitted to the banks. Any
// metadata supplied to this endpoint will be stored on the payment cancellation
// event it causes.
// 
// This will fail with a `cancellation_failed` error
// unless the payment's status is `pending_submission`.
func (s *PaymentService) Cancel(
  ctx context.Context,
  identity string,
  p PaymentCancelParams) (*PaymentCancelResult, error) {
  uri, err := url.Parse(fmt.Sprintf(
      s.endpoint + "/payments/%v/actions/cancel",
      identity,))
  if err != nil {
    return nil, err
  }

  var body io.Reader

  var buf bytes.Buffer
  err = json.NewEncoder(&buf).Encode(map[string]interface{}{
    "payments": p,
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
    *PaymentCancelResult
    Err *APIError `json:"error"`
  }

  err = json.NewDecoder(res.Body).Decode(&result)
  if err != nil {
    return nil, err
  }

  if result.Err != nil {
    return nil, result.Err
  }

  return result.PaymentCancelResult, nil
}


// PaymentRetryParams parameters
type PaymentRetryParams struct {
      Metadata map[string]interface{} `json:"metadata,omitempty"`
        
    }
// PaymentRetryResult parameters
type PaymentRetryResult struct {
      Payments struct {
      Amount string `json:"amount,omitempty"`
        AmountRefunded string `json:"amount_refunded,omitempty"`
        ChargeDate string `json:"charge_date,omitempty"`
        CreatedAt string `json:"created_at,omitempty"`
        Currency string `json:"currency,omitempty"`
        Description string `json:"description,omitempty"`
        Id string `json:"id,omitempty"`
        Links struct {
      Creditor string `json:"creditor,omitempty"`
        Mandate string `json:"mandate,omitempty"`
        Payout string `json:"payout,omitempty"`
        Subscription string `json:"subscription,omitempty"`
        
    } `json:"links,omitempty"`
        Metadata map[string]interface{} `json:"metadata,omitempty"`
        Reference string `json:"reference,omitempty"`
        Status string `json:"status,omitempty"`
        
    } `json:"payments,omitempty"`
        
    }

// Retry
// <a name="retry_failed"></a>Retries a failed payment if the underlying mandate
// is active. You will receive a `resubmission_requested` webhook, but after
// that retrying the payment follows the same process as its initial creation,
// so you will receive a `submitted` webhook, followed by a `confirmed` or
// `failed` event. Any metadata supplied to this endpoint will be stored against
// the payment submission event it causes.
// 
// This will return a
// `retry_failed` error if the payment has not failed.
// 
// Payments can be
// retried up to 3 times.
func (s *PaymentService) Retry(
  ctx context.Context,
  identity string,
  p PaymentRetryParams) (*PaymentRetryResult, error) {
  uri, err := url.Parse(fmt.Sprintf(
      s.endpoint + "/payments/%v/actions/retry",
      identity,))
  if err != nil {
    return nil, err
  }

  var body io.Reader

  var buf bytes.Buffer
  err = json.NewEncoder(&buf).Encode(map[string]interface{}{
    "payments": p,
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
    *PaymentRetryResult
    Err *APIError `json:"error"`
  }

  err = json.NewDecoder(res.Body).Decode(&result)
  if err != nil {
    return nil, err
  }

  if result.Err != nil {
    return nil, result.Err
  }

  return result.PaymentRetryResult, nil
}

