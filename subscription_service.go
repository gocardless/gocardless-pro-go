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


// SubscriptionService manages subscriptions
type SubscriptionService struct {
  endpoint string
  token string
  client *http.Client
}


// Subscription model
type Subscription struct {
      Amount int `url:",omitempty" json:"amount,omitempty"`
      AppFee int `url:",omitempty" json:"app_fee,omitempty"`
      CreatedAt string `url:",omitempty" json:"created_at,omitempty"`
      Currency string `url:",omitempty" json:"currency,omitempty"`
      DayOfMonth int `url:",omitempty" json:"day_of_month,omitempty"`
      EndDate string `url:",omitempty" json:"end_date,omitempty"`
      Id string `url:",omitempty" json:"id,omitempty"`
      Interval int `url:",omitempty" json:"interval,omitempty"`
      IntervalUnit string `url:",omitempty" json:"interval_unit,omitempty"`
      Links struct {
      Mandate string `url:",omitempty" json:"mandate,omitempty"`
      } `url:",omitempty" json:"links,omitempty"`
      Metadata map[string]interface{} `url:",omitempty" json:"metadata,omitempty"`
      Month string `url:",omitempty" json:"month,omitempty"`
      Name string `url:",omitempty" json:"name,omitempty"`
      PaymentReference string `url:",omitempty" json:"payment_reference,omitempty"`
      StartDate string `url:",omitempty" json:"start_date,omitempty"`
      Status string `url:",omitempty" json:"status,omitempty"`
      UpcomingPayments []struct {
      Amount int `url:",omitempty" json:"amount,omitempty"`
      ChargeDate string `url:",omitempty" json:"charge_date,omitempty"`
      } `url:",omitempty" json:"upcoming_payments,omitempty"`
      }




// SubscriptionCreateParams parameters
type SubscriptionCreateParams struct {
      Amount int `url:",omitempty" json:"amount,omitempty"`
      AppFee int `url:",omitempty" json:"app_fee,omitempty"`
      Count int `url:",omitempty" json:"count,omitempty"`
      Currency string `url:",omitempty" json:"currency,omitempty"`
      DayOfMonth int `url:",omitempty" json:"day_of_month,omitempty"`
      EndDate string `url:",omitempty" json:"end_date,omitempty"`
      Interval int `url:",omitempty" json:"interval,omitempty"`
      IntervalUnit string `url:",omitempty" json:"interval_unit,omitempty"`
      Links struct {
      Mandate string `url:",omitempty" json:"mandate,omitempty"`
      } `url:",omitempty" json:"links,omitempty"`
      Metadata map[string]interface{} `url:",omitempty" json:"metadata,omitempty"`
      Month string `url:",omitempty" json:"month,omitempty"`
      Name string `url:",omitempty" json:"name,omitempty"`
      PaymentReference string `url:",omitempty" json:"payment_reference,omitempty"`
      StartDate string `url:",omitempty" json:"start_date,omitempty"`
      }

// Create
// Creates a new subscription object
func (s *SubscriptionService) Create(ctx context.Context, p SubscriptionCreateParams) (*Subscription,error) {
  uri, err := url.Parse(fmt.Sprintf(s.endpoint + "/subscriptions",))
  if err != nil {
    return nil, err
  }

  var body io.Reader

  var buf bytes.Buffer
  err = json.NewEncoder(&buf).Encode(map[string]interface{}{
    "subscriptions": p,
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
Subscription *Subscription `json:"subscriptions"`
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

if result.Subscription == nil {
    return nil, errors.New("missing result")
  }

  return result.Subscription, nil
}


// SubscriptionListParams parameters
type SubscriptionListParams struct {
      After string `url:",omitempty" json:"after,omitempty"`
      Before string `url:",omitempty" json:"before,omitempty"`
      CreatedAt struct {
      Gt string `url:",omitempty" json:"gt,omitempty"`
      Gte string `url:",omitempty" json:"gte,omitempty"`
      Lt string `url:",omitempty" json:"lt,omitempty"`
      Lte string `url:",omitempty" json:"lte,omitempty"`
      } `url:",omitempty" json:"created_at,omitempty"`
      Customer string `url:",omitempty" json:"customer,omitempty"`
      Limit int `url:",omitempty" json:"limit,omitempty"`
      Mandate string `url:",omitempty" json:"mandate,omitempty"`
      Status []string `url:",omitempty" json:"status,omitempty"`
      }

// SubscriptionListResult response including pagination metadata
type SubscriptionListResult struct {
  Subscriptions []Subscription `json:"subscriptions"`
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
// subscriptions.
func (s *SubscriptionService) List(ctx context.Context, p SubscriptionListParams) (*SubscriptionListResult,error) {
  uri, err := url.Parse(fmt.Sprintf(s.endpoint + "/subscriptions",))
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
*SubscriptionListResult
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

if result.SubscriptionListResult == nil {
    return nil, errors.New("missing result")
  }

  return result.SubscriptionListResult, nil
}



// Get
// Retrieves the details of a single subscription.
func (s *SubscriptionService) Get(ctx context.Context,identity string) (*Subscription,error) {
  uri, err := url.Parse(fmt.Sprintf(s.endpoint + "/subscriptions/%v",
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
Subscription *Subscription `json:"subscriptions"`
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

if result.Subscription == nil {
    return nil, errors.New("missing result")
  }

  return result.Subscription, nil
}


// SubscriptionUpdateParams parameters
type SubscriptionUpdateParams struct {
      Amount int `url:",omitempty" json:"amount,omitempty"`
      AppFee int `url:",omitempty" json:"app_fee,omitempty"`
      Metadata map[string]interface{} `url:",omitempty" json:"metadata,omitempty"`
      Name string `url:",omitempty" json:"name,omitempty"`
      PaymentReference string `url:",omitempty" json:"payment_reference,omitempty"`
      }

// Update
// Updates a subscription object.
// 
// This fails with:
// 
// - `validation_failed` if invalid data is provided when attempting to update a
// subscription.
// 
// - `subscription_not_active` if the subscription is no longer active.
// 
// - `subscription_already_ended` if the subscription has taken all payments.
// 
// - `mandate_payments_require_approval` if the amount is being changed and the
// mandate requires approval.
// 
// - `number_of_subscription_amendments_exceeded` error if the subscription
// amount has already been changed 10 times.
// 
// - `forbidden` if the amount is being changed, and the subscription was
// created by an app and you are not authenticated as that app, or if the
// subscription was not created by an app and you are authenticated as an app
// 
// - `resource_created_by_another_app` if the app fee is being changed, and the
// subscription was created by an app other than the app you are authenticated
// as
// 
func (s *SubscriptionService) Update(ctx context.Context,identity string, p SubscriptionUpdateParams) (*Subscription,error) {
  uri, err := url.Parse(fmt.Sprintf(s.endpoint + "/subscriptions/%v",
      identity,))
  if err != nil {
    return nil, err
  }

  var body io.Reader

  var buf bytes.Buffer
  err = json.NewEncoder(&buf).Encode(map[string]interface{}{
    "subscriptions": p,
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
Subscription *Subscription `json:"subscriptions"`
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

if result.Subscription == nil {
    return nil, errors.New("missing result")
  }

  return result.Subscription, nil
}


// SubscriptionCancelParams parameters
type SubscriptionCancelParams struct {
      Metadata map[string]interface{} `url:",omitempty" json:"metadata,omitempty"`
      }

// Cancel
// Immediately cancels a subscription; no more payments will be created under
// it. Any metadata supplied to this endpoint will be stored on the payment
// cancellation event it causes.
// 
// This will fail with a cancellation_failed error if the subscription is
// already cancelled or finished.
func (s *SubscriptionService) Cancel(ctx context.Context,identity string, p SubscriptionCancelParams) (*Subscription,error) {
  uri, err := url.Parse(fmt.Sprintf(s.endpoint + "/subscriptions/%v/actions/cancel",
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
Subscription *Subscription `json:"subscriptions"`
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

if result.Subscription == nil {
    return nil, errors.New("missing result")
  }

  return result.Subscription, nil
}

