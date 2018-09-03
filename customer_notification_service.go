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


// CustomerNotificationService manages customer_notifications
type CustomerNotificationService struct {
  endpoint string
  token string
  client *http.Client
}


// CustomerNotification model
type CustomerNotification struct {
      ActionTaken string `url:",omitempty" json:"action_taken,omitempty"`
      ActionTakenAt string `url:",omitempty" json:"action_taken_at,omitempty"`
      ActionTakenBy string `url:",omitempty" json:"action_taken_by,omitempty"`
      Id string `url:",omitempty" json:"id,omitempty"`
      Links struct {
      Customer string `url:",omitempty" json:"customer,omitempty"`
      Event string `url:",omitempty" json:"event,omitempty"`
      Mandate string `url:",omitempty" json:"mandate,omitempty"`
      Payment string `url:",omitempty" json:"payment,omitempty"`
      Refund string `url:",omitempty" json:"refund,omitempty"`
      Subscription string `url:",omitempty" json:"subscription,omitempty"`
      } `url:",omitempty" json:"links,omitempty"`
      Type string `url:",omitempty" json:"type,omitempty"`
      }




// CustomerNotificationHandleParams parameters
type CustomerNotificationHandleParams map[string]interface{}

// Handle
// "Handling" a notification means that you have sent the notification yourself
// (and
// don't want GoCardless to send it).
// If the notification has already been actioned, or the deadline to notify has
// passed,
// this endpoint will return an `already_actioned` error and you should not take
// further action.
// 
func (s *CustomerNotificationService) Handle(ctx context.Context,identity string, p CustomerNotificationHandleParams) (*CustomerNotification,error) {
  uri, err := url.Parse(fmt.Sprintf(s.endpoint + "/customer_notifications/%v/actions/handle",
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
CustomerNotification *CustomerNotification `json:"customer_notifications"`
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

if result.CustomerNotification == nil {
    return nil, errors.New("missing result")
  }

  return result.CustomerNotification, nil
}

