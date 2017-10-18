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


// PaymentService manages payments
type PaymentService struct {
  endpoint string
  token string
  client *http.Client
}


// Payment model
type Payment struct {
      Amount int `url:",omitempty" json:"amount,omitempty"`
      AmountRefunded int `url:",omitempty" json:"amount_refunded,omitempty"`
      ChargeDate string `url:",omitempty" json:"charge_date,omitempty"`
      CreatedAt string `url:",omitempty" json:"created_at,omitempty"`
      Currency string `url:",omitempty" json:"currency,omitempty"`
      Description string `url:",omitempty" json:"description,omitempty"`
      Id string `url:",omitempty" json:"id,omitempty"`
      Links struct {
      Creditor string `url:",omitempty" json:"creditor,omitempty"`
      Mandate string `url:",omitempty" json:"mandate,omitempty"`
      Payout string `url:",omitempty" json:"payout,omitempty"`
      Subscription string `url:",omitempty" json:"subscription,omitempty"`
      } `url:",omitempty" json:"links,omitempty"`
      Metadata map[string]interface{} `url:",omitempty" json:"metadata,omitempty"`
      Reference string `url:",omitempty" json:"reference,omitempty"`
      Status string `url:",omitempty" json:"status,omitempty"`
      }




// PaymentCreateParams parameters
type PaymentCreateParams struct {
      Amount int `url:",omitempty" json:"amount,omitempty"`
      AppFee int `url:",omitempty" json:"app_fee,omitempty"`
      ChargeDate string `url:",omitempty" json:"charge_date,omitempty"`
      Currency string `url:",omitempty" json:"currency,omitempty"`
      Description string `url:",omitempty" json:"description,omitempty"`
      Links struct {
      Mandate string `url:",omitempty" json:"mandate,omitempty"`
      } `url:",omitempty" json:"links,omitempty"`
      Metadata map[string]interface{} `url:",omitempty" json:"metadata,omitempty"`
      Reference string `url:",omitempty" json:"reference,omitempty"`
      }

// Create
// <a name="mandate_is_inactive"></a>Creates a new payment object.
// 
// This fails with a `mandate_is_inactive` error if the linked
// [mandate](#core-endpoints-mandates) is cancelled or has failed. Payments can
// be created against mandates with status of: `pending_customer_approval`,
// `pending_submission`, `submitted`, and `active`.
func (s *PaymentService) Create(ctx context.Context, p PaymentCreateParams) (*Payment,error) {
  uri, err := url.Parse(fmt.Sprintf(s.endpoint + "/payments",))
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
  req.Header.Set("Idempotency-Key", NewIdempotencyKey())

  client := s.client
  if client == nil {
    client = http.DefaultClient
  }

  var result struct {
    Err *APIError `json:"error"`
Payment *Payment `json:"payments"`
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

if result.Payment == nil {
    return nil, errors.New("missing result")
  }

  return result.Payment, nil
}


// PaymentListParams parameters
type PaymentListParams struct {
      After string `url:",omitempty" json:"after,omitempty"`
      Before string `url:",omitempty" json:"before,omitempty"`
      CreatedAt struct {
      Gt string `url:",omitempty" json:"gt,omitempty"`
      Gte string `url:",omitempty" json:"gte,omitempty"`
      Lt string `url:",omitempty" json:"lt,omitempty"`
      Lte string `url:",omitempty" json:"lte,omitempty"`
      } `url:",omitempty" json:"created_at,omitempty"`
      Creditor string `url:",omitempty" json:"creditor,omitempty"`
      Currency string `url:",omitempty" json:"currency,omitempty"`
      Customer string `url:",omitempty" json:"customer,omitempty"`
      Limit int `url:",omitempty" json:"limit,omitempty"`
      Mandate string `url:",omitempty" json:"mandate,omitempty"`
      Status string `url:",omitempty" json:"status,omitempty"`
      Subscription string `url:",omitempty" json:"subscription,omitempty"`
      }

// PaymentListResult response including pagination metadata
type PaymentListResult struct {
  Payments []Payment `json:"payments"`
  Meta struct {
      Cursors struct {
      After string `url:",omitempty" json:"after,omitempty"`
      Before string `url:",omitempty" json:"before,omitempty"`
      } `url:",omitempty" json:"cursors,omitempty"`
      Limit int `url:",omitempty" json:"limit,omitempty"`
      } `json:"meta"`
}


// List
// Returns a [cursor-paginated](#api-usage-cursor-pagination) list of your
// payments.
func (s *PaymentService) List(ctx context.Context, p PaymentListParams) (*PaymentListResult,error) {
  uri, err := url.Parse(fmt.Sprintf(s.endpoint + "/payments",))
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
*PaymentListResult
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

if result.PaymentListResult == nil {
    return nil, errors.New("missing result")
  }

  return result.PaymentListResult, nil
}



// Get
// Retrieves the details of a single existing payment.
func (s *PaymentService) Get(ctx context.Context,identity string) (*Payment,error) {
  uri, err := url.Parse(fmt.Sprintf(s.endpoint + "/payments/%v",
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
Payment *Payment `json:"payments"`
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

if result.Payment == nil {
    return nil, errors.New("missing result")
  }

  return result.Payment, nil
}


// PaymentUpdateParams parameters
type PaymentUpdateParams struct {
      Metadata map[string]interface{} `url:",omitempty" json:"metadata,omitempty"`
      }

// Update
// Updates a payment object. This accepts only the metadata parameter.
func (s *PaymentService) Update(ctx context.Context,identity string, p PaymentUpdateParams) (*Payment,error) {
  uri, err := url.Parse(fmt.Sprintf(s.endpoint + "/payments/%v",
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
  req.Header.Set("Idempotency-Key", NewIdempotencyKey())

  client := s.client
  if client == nil {
    client = http.DefaultClient
  }

  var result struct {
    Err *APIError `json:"error"`
Payment *Payment `json:"payments"`
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

if result.Payment == nil {
    return nil, errors.New("missing result")
  }

  return result.Payment, nil
}


// PaymentCancelParams parameters
type PaymentCancelParams struct {
      Metadata map[string]interface{} `url:",omitempty" json:"metadata,omitempty"`
      }

// Cancel
// Cancels the payment if it has not already been submitted to the banks. Any
// metadata supplied to this endpoint will be stored on the payment cancellation
// event it causes.
// 
// This will fail with a `cancellation_failed` error unless the payment's status
// is `pending_submission`.
func (s *PaymentService) Cancel(ctx context.Context,identity string, p PaymentCancelParams) (*Payment,error) {
  uri, err := url.Parse(fmt.Sprintf(s.endpoint + "/payments/%v/actions/cancel",
      identity,))
  if err != nil {
    return nil, err
  }

  var body io.Reader

  var buf bytes.Buffer
  err = json.NewEncoder(&buf).Encode(map[string]interface{}{
    "data": p,
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
Payment *Payment `json:"payments"`
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

if result.Payment == nil {
    return nil, errors.New("missing result")
  }

  return result.Payment, nil
}


// PaymentRetryParams parameters
type PaymentRetryParams struct {
      Metadata map[string]interface{} `url:",omitempty" json:"metadata,omitempty"`
      }

// Retry
// <a name="retry_failed"></a>Retries a failed payment if the underlying mandate
// is active. You will receive a `resubmission_requested` webhook, but after
// that retrying the payment follows the same process as its initial creation,
// so you will receive a `submitted` webhook, followed by a `confirmed` or
// `failed` event. Any metadata supplied to this endpoint will be stored against
// the payment submission event it causes.
// 
// This will return a `retry_failed` error if the payment has not failed.
// 
// Payments can be retried up to 3 times.
func (s *PaymentService) Retry(ctx context.Context,identity string, p PaymentRetryParams) (*Payment,error) {
  uri, err := url.Parse(fmt.Sprintf(s.endpoint + "/payments/%v/actions/retry",
      identity,))
  if err != nil {
    return nil, err
  }

  var body io.Reader

  var buf bytes.Buffer
  err = json.NewEncoder(&buf).Encode(map[string]interface{}{
    "data": p,
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
Payment *Payment `json:"payments"`
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

if result.Payment == nil {
    return nil, errors.New("missing result")
  }

  return result.Payment, nil
}

