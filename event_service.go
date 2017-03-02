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


type EventService struct {
  endpoint string
  token string
  client *http.Client
}



// EventListParams parameters
type EventListParams struct {
      Action string `json:"action,omitempty"`
        After string `json:"after,omitempty"`
        Before string `json:"before,omitempty"`
        CreatedAt struct {
      Gt string `json:"gt,omitempty"`
        Gte string `json:"gte,omitempty"`
        Lt string `json:"lt,omitempty"`
        Lte string `json:"lte,omitempty"`
        
    } `json:"created_at,omitempty"`
        Include string `json:"include,omitempty"`
        Limit string `json:"limit,omitempty"`
        Mandate string `json:"mandate,omitempty"`
        ParentEvent string `json:"parent_event,omitempty"`
        Payment string `json:"payment,omitempty"`
        Payout string `json:"payout,omitempty"`
        Refund string `json:"refund,omitempty"`
        ResourceType string `json:"resource_type,omitempty"`
        Subscription string `json:"subscription,omitempty"`
        
    }
// EventListResult parameters
type EventListResult struct {
      Events []struct {
      Action string `json:"action,omitempty"`
        CreatedAt string `json:"created_at,omitempty"`
        Details struct {
      Cause string `json:"cause,omitempty"`
        Description string `json:"description,omitempty"`
        Origin string `json:"origin,omitempty"`
        ReasonCode string `json:"reason_code,omitempty"`
        Scheme string `json:"scheme,omitempty"`
        
    } `json:"details,omitempty"`
        Id string `json:"id,omitempty"`
        Links struct {
      Mandate string `json:"mandate,omitempty"`
        NewCustomerBankAccount string `json:"new_customer_bank_account,omitempty"`
        NewMandate string `json:"new_mandate,omitempty"`
        Organisation string `json:"organisation,omitempty"`
        ParentEvent string `json:"parent_event,omitempty"`
        Payment string `json:"payment,omitempty"`
        Payout string `json:"payout,omitempty"`
        PreviousCustomerBankAccount string `json:"previous_customer_bank_account,omitempty"`
        Refund string `json:"refund,omitempty"`
        Subscription string `json:"subscription,omitempty"`
        
    } `json:"links,omitempty"`
        Metadata map[string]interface{} `json:"metadata,omitempty"`
        ResourceType string `json:"resource_type,omitempty"`
        
    } `json:"events,omitempty"`
        Meta struct {
      Cursors struct {
      After string `json:"after,omitempty"`
        Before string `json:"before,omitempty"`
        
    } `json:"cursors,omitempty"`
        Limit int `json:"limit,omitempty"`
        
    } `json:"meta,omitempty"`
        
    }

// List
// Returns a [cursor-paginated](#api-usage-cursor-pagination) list of your
// events.
func (s *EventService) List(
  ctx context.Context,
  p EventListParams) (*EventListResult, error) {
  uri, err := url.Parse(fmt.Sprintf(
      s.endpoint + "/events",))
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
    *EventListResult
    Err *APIError `json:"error"`
  }

  err = json.NewDecoder(res.Body).Decode(&result)
  if err != nil {
    return nil, err
  }

  if result.Err != nil {
    return nil, result.Err
  }

  return result.EventListResult, nil
}


// EventGetResult parameters
type EventGetResult struct {
      Events struct {
      Action string `json:"action,omitempty"`
        CreatedAt string `json:"created_at,omitempty"`
        Details struct {
      Cause string `json:"cause,omitempty"`
        Description string `json:"description,omitempty"`
        Origin string `json:"origin,omitempty"`
        ReasonCode string `json:"reason_code,omitempty"`
        Scheme string `json:"scheme,omitempty"`
        
    } `json:"details,omitempty"`
        Id string `json:"id,omitempty"`
        Links struct {
      Mandate string `json:"mandate,omitempty"`
        NewCustomerBankAccount string `json:"new_customer_bank_account,omitempty"`
        NewMandate string `json:"new_mandate,omitempty"`
        Organisation string `json:"organisation,omitempty"`
        ParentEvent string `json:"parent_event,omitempty"`
        Payment string `json:"payment,omitempty"`
        Payout string `json:"payout,omitempty"`
        PreviousCustomerBankAccount string `json:"previous_customer_bank_account,omitempty"`
        Refund string `json:"refund,omitempty"`
        Subscription string `json:"subscription,omitempty"`
        
    } `json:"links,omitempty"`
        Metadata map[string]interface{} `json:"metadata,omitempty"`
        ResourceType string `json:"resource_type,omitempty"`
        
    } `json:"events,omitempty"`
        
    }

// Get
// Retrieves the details of a single event.
func (s *EventService) Get(
  ctx context.Context,
  identity string) (*EventGetResult, error) {
  uri, err := url.Parse(fmt.Sprintf(
      s.endpoint + "/events/%v",
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
    *EventGetResult
    Err *APIError `json:"error"`
  }

  err = json.NewDecoder(res.Body).Decode(&result)
  if err != nil {
    return nil, err
  }

  if result.Err != nil {
    return nil, result.Err
  }

  return result.EventGetResult, nil
}

