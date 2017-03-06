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
      Amount int `url:",omitempty" json:"amount,omitempty"`
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
// SubscriptionCreateResult parameters
type SubscriptionCreateResult struct {
      Subscriptions struct {
      Amount int `url:",omitempty" json:"amount,omitempty"`
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
        
    } `url:",omitempty" json:"subscriptions,omitempty"`
        
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
        
    }
// SubscriptionListResult parameters
type SubscriptionListResult struct {
      Meta struct {
      Cursors struct {
      After string `url:",omitempty" json:"after,omitempty"`
        Before string `url:",omitempty" json:"before,omitempty"`
        
    } `url:",omitempty" json:"cursors,omitempty"`
        Limit int `url:",omitempty" json:"limit,omitempty"`
        
    } `url:",omitempty" json:"meta,omitempty"`
        Subscriptions []struct {
      Amount int `url:",omitempty" json:"amount,omitempty"`
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
        
    } `url:",omitempty" json:"subscriptions,omitempty"`
        
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
      Amount int `url:",omitempty" json:"amount,omitempty"`
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
        
    } `url:",omitempty" json:"subscriptions,omitempty"`
        
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
      Metadata map[string]interface{} `url:",omitempty" json:"metadata,omitempty"`
        Name string `url:",omitempty" json:"name,omitempty"`
        PaymentReference string `url:",omitempty" json:"payment_reference,omitempty"`
        
    }
// SubscriptionUpdateResult parameters
type SubscriptionUpdateResult struct {
      Subscriptions struct {
      Amount int `url:",omitempty" json:"amount,omitempty"`
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
        
    } `url:",omitempty" json:"subscriptions,omitempty"`
        
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
      Metadata map[string]interface{} `url:",omitempty" json:"metadata,omitempty"`
        
    }
// SubscriptionCancelResult parameters
type SubscriptionCancelResult struct {
      Subscriptions struct {
      Amount int `url:",omitempty" json:"amount,omitempty"`
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
        
    } `url:",omitempty" json:"subscriptions,omitempty"`
        
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

