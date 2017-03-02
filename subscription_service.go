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


type SubscriptionService struct {
  endpoint string
  token string
  client *http.Client
}



// SubscriptionCreateParams parameters
type SubscriptionCreateParams struct {
      Amount string `json:"amount,omitempty"`
        Count string `json:"count,omitempty"`
        Currency string `json:"currency,omitempty"`
        DayOfMonth string `json:"day_of_month,omitempty"`
        EndDate string `json:"end_date,omitempty"`
        Interval string `json:"interval,omitempty"`
        IntervalUnit string `json:"interval_unit,omitempty"`
        Links struct {
      Mandate string `json:"mandate,omitempty"`
        
    } `json:"links,omitempty"`
        Metadata map[string]interface{} `json:"metadata,omitempty"`
        Month string `json:"month,omitempty"`
        Name string `json:"name,omitempty"`
        PaymentReference string `json:"payment_reference,omitempty"`
        StartDate string `json:"start_date,omitempty"`
        
    }
// SubscriptionCreateResult parameters
type SubscriptionCreateResult struct {
      Subscriptions struct {
      Amount string `json:"amount,omitempty"`
        CreatedAt string `json:"created_at,omitempty"`
        Currency string `json:"currency,omitempty"`
        DayOfMonth string `json:"day_of_month,omitempty"`
        EndDate string `json:"end_date,omitempty"`
        Id string `json:"id,omitempty"`
        Interval string `json:"interval,omitempty"`
        IntervalUnit string `json:"interval_unit,omitempty"`
        Links struct {
      Mandate string `json:"mandate,omitempty"`
        
    } `json:"links,omitempty"`
        Metadata map[string]interface{} `json:"metadata,omitempty"`
        Month string `json:"month,omitempty"`
        Name string `json:"name,omitempty"`
        PaymentReference string `json:"payment_reference,omitempty"`
        StartDate string `json:"start_date,omitempty"`
        Status string `json:"status,omitempty"`
        UpcomingPayments []struct {
      Amount string `json:"amount,omitempty"`
        ChargeDate string `json:"charge_date,omitempty"`
        
    } `json:"upcoming_payments,omitempty"`
        
    } `json:"subscriptions,omitempty"`
        
    }

// Create
// Creates a new subscription object
func (s *SubscriptionService) Create(
  ctx context.Context,
  p SubscriptionCreateParams) (*SubscriptionCreateResult, error) {
  uri, err := url.Parse(fmt.Sprintf(
      s.endpoint + "/subscriptions",))
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
    *SubscriptionCreateResult
    Err *APIError `json:"error"`
  }

  err = json.NewDecoder(res.Body).Decode(&result)
  if err != nil {
    return nil, err
  }

  if result.Err != nil {
    return nil, result.Err
  }

  return result.SubscriptionCreateResult, nil
}


// SubscriptionListParams parameters
type SubscriptionListParams struct {
      After string `json:"after,omitempty"`
        Before string `json:"before,omitempty"`
        CreatedAt struct {
      Gt string `json:"gt,omitempty"`
        Gte string `json:"gte,omitempty"`
        Lt string `json:"lt,omitempty"`
        Lte string `json:"lte,omitempty"`
        
    } `json:"created_at,omitempty"`
        Customer string `json:"customer,omitempty"`
        Limit string `json:"limit,omitempty"`
        Mandate string `json:"mandate,omitempty"`
        
    }
// SubscriptionListResult parameters
type SubscriptionListResult struct {
      Meta struct {
      Cursors struct {
      After string `json:"after,omitempty"`
        Before string `json:"before,omitempty"`
        
    } `json:"cursors,omitempty"`
        Limit int `json:"limit,omitempty"`
        
    } `json:"meta,omitempty"`
        Subscriptions []struct {
      Amount string `json:"amount,omitempty"`
        CreatedAt string `json:"created_at,omitempty"`
        Currency string `json:"currency,omitempty"`
        DayOfMonth string `json:"day_of_month,omitempty"`
        EndDate string `json:"end_date,omitempty"`
        Id string `json:"id,omitempty"`
        Interval string `json:"interval,omitempty"`
        IntervalUnit string `json:"interval_unit,omitempty"`
        Links struct {
      Mandate string `json:"mandate,omitempty"`
        
    } `json:"links,omitempty"`
        Metadata map[string]interface{} `json:"metadata,omitempty"`
        Month string `json:"month,omitempty"`
        Name string `json:"name,omitempty"`
        PaymentReference string `json:"payment_reference,omitempty"`
        StartDate string `json:"start_date,omitempty"`
        Status string `json:"status,omitempty"`
        UpcomingPayments []struct {
      Amount string `json:"amount,omitempty"`
        ChargeDate string `json:"charge_date,omitempty"`
        
    } `json:"upcoming_payments,omitempty"`
        
    } `json:"subscriptions,omitempty"`
        
    }

// List
// Returns a [cursor-paginated](#api-usage-cursor-pagination) list of your
// subscriptions.
func (s *SubscriptionService) List(
  ctx context.Context,
  p SubscriptionListParams) (*SubscriptionListResult, error) {
  uri, err := url.Parse(fmt.Sprintf(
      s.endpoint + "/subscriptions",))
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
    *SubscriptionListResult
    Err *APIError `json:"error"`
  }

  err = json.NewDecoder(res.Body).Decode(&result)
  if err != nil {
    return nil, err
  }

  if result.Err != nil {
    return nil, result.Err
  }

  return result.SubscriptionListResult, nil
}


// SubscriptionGetResult parameters
type SubscriptionGetResult struct {
      Subscriptions struct {
      Amount string `json:"amount,omitempty"`
        CreatedAt string `json:"created_at,omitempty"`
        Currency string `json:"currency,omitempty"`
        DayOfMonth string `json:"day_of_month,omitempty"`
        EndDate string `json:"end_date,omitempty"`
        Id string `json:"id,omitempty"`
        Interval string `json:"interval,omitempty"`
        IntervalUnit string `json:"interval_unit,omitempty"`
        Links struct {
      Mandate string `json:"mandate,omitempty"`
        
    } `json:"links,omitempty"`
        Metadata map[string]interface{} `json:"metadata,omitempty"`
        Month string `json:"month,omitempty"`
        Name string `json:"name,omitempty"`
        PaymentReference string `json:"payment_reference,omitempty"`
        StartDate string `json:"start_date,omitempty"`
        Status string `json:"status,omitempty"`
        UpcomingPayments []struct {
      Amount string `json:"amount,omitempty"`
        ChargeDate string `json:"charge_date,omitempty"`
        
    } `json:"upcoming_payments,omitempty"`
        
    } `json:"subscriptions,omitempty"`
        
    }

// Get
// Retrieves the details of a single subscription.
func (s *SubscriptionService) Get(
  ctx context.Context,
  identity string) (*SubscriptionGetResult, error) {
  uri, err := url.Parse(fmt.Sprintf(
      s.endpoint + "/subscriptions/%v",
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
    *SubscriptionGetResult
    Err *APIError `json:"error"`
  }

  err = json.NewDecoder(res.Body).Decode(&result)
  if err != nil {
    return nil, err
  }

  if result.Err != nil {
    return nil, result.Err
  }

  return result.SubscriptionGetResult, nil
}


// SubscriptionUpdateParams parameters
type SubscriptionUpdateParams struct {
      Metadata map[string]interface{} `json:"metadata,omitempty"`
        Name string `json:"name,omitempty"`
        PaymentReference string `json:"payment_reference,omitempty"`
        
    }
// SubscriptionUpdateResult parameters
type SubscriptionUpdateResult struct {
      Subscriptions struct {
      Amount string `json:"amount,omitempty"`
        CreatedAt string `json:"created_at,omitempty"`
        Currency string `json:"currency,omitempty"`
        DayOfMonth string `json:"day_of_month,omitempty"`
        EndDate string `json:"end_date,omitempty"`
        Id string `json:"id,omitempty"`
        Interval string `json:"interval,omitempty"`
        IntervalUnit string `json:"interval_unit,omitempty"`
        Links struct {
      Mandate string `json:"mandate,omitempty"`
        
    } `json:"links,omitempty"`
        Metadata map[string]interface{} `json:"metadata,omitempty"`
        Month string `json:"month,omitempty"`
        Name string `json:"name,omitempty"`
        PaymentReference string `json:"payment_reference,omitempty"`
        StartDate string `json:"start_date,omitempty"`
        Status string `json:"status,omitempty"`
        UpcomingPayments []struct {
      Amount string `json:"amount,omitempty"`
        ChargeDate string `json:"charge_date,omitempty"`
        
    } `json:"upcoming_payments,omitempty"`
        
    } `json:"subscriptions,omitempty"`
        
    }

// Update
// Updates a subscription object.
func (s *SubscriptionService) Update(
  ctx context.Context,
  identity string,
  p SubscriptionUpdateParams) (*SubscriptionUpdateResult, error) {
  uri, err := url.Parse(fmt.Sprintf(
      s.endpoint + "/subscriptions/%v",
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
    *SubscriptionUpdateResult
    Err *APIError `json:"error"`
  }

  err = json.NewDecoder(res.Body).Decode(&result)
  if err != nil {
    return nil, err
  }

  if result.Err != nil {
    return nil, result.Err
  }

  return result.SubscriptionUpdateResult, nil
}


// SubscriptionCancelParams parameters
type SubscriptionCancelParams struct {
      Metadata map[string]interface{} `json:"metadata,omitempty"`
        
    }
// SubscriptionCancelResult parameters
type SubscriptionCancelResult struct {
      Subscriptions struct {
      Amount string `json:"amount,omitempty"`
        CreatedAt string `json:"created_at,omitempty"`
        Currency string `json:"currency,omitempty"`
        DayOfMonth string `json:"day_of_month,omitempty"`
        EndDate string `json:"end_date,omitempty"`
        Id string `json:"id,omitempty"`
        Interval string `json:"interval,omitempty"`
        IntervalUnit string `json:"interval_unit,omitempty"`
        Links struct {
      Mandate string `json:"mandate,omitempty"`
        
    } `json:"links,omitempty"`
        Metadata map[string]interface{} `json:"metadata,omitempty"`
        Month string `json:"month,omitempty"`
        Name string `json:"name,omitempty"`
        PaymentReference string `json:"payment_reference,omitempty"`
        StartDate string `json:"start_date,omitempty"`
        Status string `json:"status,omitempty"`
        UpcomingPayments []struct {
      Amount string `json:"amount,omitempty"`
        ChargeDate string `json:"charge_date,omitempty"`
        
    } `json:"upcoming_payments,omitempty"`
        
    } `json:"subscriptions,omitempty"`
        
    }

// Cancel
// Immediately cancels a subscription; no more payments will be created under
// it. Any metadata supplied to this endpoint will be stored on the payment
// cancellation event it causes.
// 
// This will fail with a
// cancellation_failed error if the subscription is already cancelled or
// finished.
func (s *SubscriptionService) Cancel(
  ctx context.Context,
  identity string,
  p SubscriptionCancelParams) (*SubscriptionCancelResult, error) {
  uri, err := url.Parse(fmt.Sprintf(
      s.endpoint + "/subscriptions/%v/actions/cancel",
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
    *SubscriptionCancelResult
    Err *APIError `json:"error"`
  }

  err = json.NewDecoder(res.Body).Decode(&result)
  if err != nil {
    return nil, err
  }

  if result.Err != nil {
    return nil, result.Err
  }

  return result.SubscriptionCancelResult, nil
}

