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


// PayoutItemService manages payout_items
type PayoutItemService struct {
  endpoint string
  token string
  client *http.Client
}


// PayoutItem model
type PayoutItem struct {
      Amount string `url:",omitempty" json:"amount,omitempty"`
      Links struct {
      Mandate string `url:",omitempty" json:"mandate,omitempty"`
      Payment string `url:",omitempty" json:"payment,omitempty"`
      } `url:",omitempty" json:"links,omitempty"`
      Type string `url:",omitempty" json:"type,omitempty"`
      }




// PayoutItemListParams parameters
type PayoutItemListParams struct {
      After string `url:",omitempty" json:"after,omitempty"`
      Before string `url:",omitempty" json:"before,omitempty"`
      Limit int `url:",omitempty" json:"limit,omitempty"`
      Payout string `url:",omitempty" json:"payout,omitempty"`
      }

// PayoutItemListResult response including pagination metadata
type PayoutItemListResult struct {
  PayoutItems []PayoutItem `json:"payout_items"`
  Meta struct {
      Cursors struct {
      After string `url:",omitempty" json:"after,omitempty"`
      Before string `url:",omitempty" json:"before,omitempty"`
      } `url:",omitempty" json:"cursors,omitempty"`
      Limit int `url:",omitempty" json:"limit,omitempty"`
      } `json:"meta"`
}


// List
// Returns a [cursor-paginated](#api-usage-cursor-pagination) list of items in
// the payout.
// 
func (s *PayoutItemService) List(ctx context.Context, p PayoutItemListParams) (*PayoutItemListResult,error) {
  uri, err := url.Parse(fmt.Sprintf(s.endpoint + "/payout_items",))
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
*PayoutItemListResult
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

if result.PayoutItemListResult == nil {
    return nil, errors.New("missing result")
  }

  return result.PayoutItemListResult, nil
}

