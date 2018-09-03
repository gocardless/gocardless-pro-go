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


// EventService manages events
type EventService struct {
  endpoint string
  token string
  client *http.Client
}


// Event model
type Event struct {
      Action string `url:",omitempty" json:"action,omitempty"`
      CreatedAt string `url:",omitempty" json:"created_at,omitempty"`
      CustomerNotifications []struct {
      Deadline string `url:",omitempty" json:"deadline,omitempty"`
      Id string `url:",omitempty" json:"id,omitempty"`
      Mandatory bool `url:",omitempty" json:"mandatory,omitempty"`
      Type string `url:",omitempty" json:"type,omitempty"`
      } `url:",omitempty" json:"customer_notifications,omitempty"`
      Details struct {
      Cause string `url:",omitempty" json:"cause,omitempty"`
      Description string `url:",omitempty" json:"description,omitempty"`
      Origin string `url:",omitempty" json:"origin,omitempty"`
      ReasonCode string `url:",omitempty" json:"reason_code,omitempty"`
      Scheme string `url:",omitempty" json:"scheme,omitempty"`
      } `url:",omitempty" json:"details,omitempty"`
      Id string `url:",omitempty" json:"id,omitempty"`
      Links struct {
      Mandate string `url:",omitempty" json:"mandate,omitempty"`
      NewCustomerBankAccount string `url:",omitempty" json:"new_customer_bank_account,omitempty"`
      NewMandate string `url:",omitempty" json:"new_mandate,omitempty"`
      Organisation string `url:",omitempty" json:"organisation,omitempty"`
      ParentEvent string `url:",omitempty" json:"parent_event,omitempty"`
      Payment string `url:",omitempty" json:"payment,omitempty"`
      Payout string `url:",omitempty" json:"payout,omitempty"`
      PreviousCustomerBankAccount string `url:",omitempty" json:"previous_customer_bank_account,omitempty"`
      Refund string `url:",omitempty" json:"refund,omitempty"`
      Subscription string `url:",omitempty" json:"subscription,omitempty"`
      } `url:",omitempty" json:"links,omitempty"`
      Metadata map[string]interface{} `url:",omitempty" json:"metadata,omitempty"`
      ResourceType string `url:",omitempty" json:"resource_type,omitempty"`
      }




// EventListParams parameters
type EventListParams struct {
      Action string `url:",omitempty" json:"action,omitempty"`
      After string `url:",omitempty" json:"after,omitempty"`
      Before string `url:",omitempty" json:"before,omitempty"`
      CreatedAt struct {
      Gt string `url:",omitempty" json:"gt,omitempty"`
      Gte string `url:",omitempty" json:"gte,omitempty"`
      Lt string `url:",omitempty" json:"lt,omitempty"`
      Lte string `url:",omitempty" json:"lte,omitempty"`
      } `url:",omitempty" json:"created_at,omitempty"`
      Include string `url:",omitempty" json:"include,omitempty"`
      Limit int `url:",omitempty" json:"limit,omitempty"`
      Mandate string `url:",omitempty" json:"mandate,omitempty"`
      ParentEvent string `url:",omitempty" json:"parent_event,omitempty"`
      Payment string `url:",omitempty" json:"payment,omitempty"`
      Payout string `url:",omitempty" json:"payout,omitempty"`
      Refund string `url:",omitempty" json:"refund,omitempty"`
      ResourceType string `url:",omitempty" json:"resource_type,omitempty"`
      Subscription string `url:",omitempty" json:"subscription,omitempty"`
      }

// EventListResult response including pagination metadata
type EventListResult struct {
  Events []Event `json:"events"`
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
// events.
func (s *EventService) List(ctx context.Context, p EventListParams) (*EventListResult,error) {
  uri, err := url.Parse(fmt.Sprintf(s.endpoint + "/events",))
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
*EventListResult
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

if result.EventListResult == nil {
    return nil, errors.New("missing result")
  }

  return result.EventListResult, nil
}



// Get
// Retrieves the details of a single event.
func (s *EventService) Get(ctx context.Context,identity string) (*Event,error) {
  uri, err := url.Parse(fmt.Sprintf(s.endpoint + "/events/%v",
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
Event *Event `json:"events"`
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

if result.Event == nil {
    return nil, errors.New("missing result")
  }

  return result.Event, nil
}

