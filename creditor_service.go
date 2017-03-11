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


// CreditorService manages creditors
type CreditorService struct {
  endpoint string
  token string
  client *http.Client
}


// Creditor model
type Creditor struct {
      AddressLine1 string `url:",omitempty" json:"address_line1,omitempty"`
      AddressLine2 string `url:",omitempty" json:"address_line2,omitempty"`
      AddressLine3 string `url:",omitempty" json:"address_line3,omitempty"`
      City string `url:",omitempty" json:"city,omitempty"`
      CountryCode string `url:",omitempty" json:"country_code,omitempty"`
      CreatedAt string `url:",omitempty" json:"created_at,omitempty"`
      Id string `url:",omitempty" json:"id,omitempty"`
      Links struct {
      DefaultEurPayoutAccount string `url:",omitempty" json:"default_eur_payout_account,omitempty"`
      DefaultGbpPayoutAccount string `url:",omitempty" json:"default_gbp_payout_account,omitempty"`
      DefaultSekPayoutAccount string `url:",omitempty" json:"default_sek_payout_account,omitempty"`
      } `url:",omitempty" json:"links,omitempty"`
      LogoUrl string `url:",omitempty" json:"logo_url,omitempty"`
      Name string `url:",omitempty" json:"name,omitempty"`
      PostalCode string `url:",omitempty" json:"postal_code,omitempty"`
      Region string `url:",omitempty" json:"region,omitempty"`
      SchemeIdentifiers []struct {
      AddressLine1 string `url:",omitempty" json:"address_line1,omitempty"`
      AddressLine2 string `url:",omitempty" json:"address_line2,omitempty"`
      AddressLine3 string `url:",omitempty" json:"address_line3,omitempty"`
      CanSpecifyMandateReference bool `url:",omitempty" json:"can_specify_mandate_reference,omitempty"`
      City string `url:",omitempty" json:"city,omitempty"`
      CountryCode string `url:",omitempty" json:"country_code,omitempty"`
      Currency string `url:",omitempty" json:"currency,omitempty"`
      Email string `url:",omitempty" json:"email,omitempty"`
      MinimumAdvanceNotice int `url:",omitempty" json:"minimum_advance_notice,omitempty"`
      Name string `url:",omitempty" json:"name,omitempty"`
      PhoneNumber string `url:",omitempty" json:"phone_number,omitempty"`
      PostalCode string `url:",omitempty" json:"postal_code,omitempty"`
      Reference string `url:",omitempty" json:"reference,omitempty"`
      Region string `url:",omitempty" json:"region,omitempty"`
      Scheme string `url:",omitempty" json:"scheme,omitempty"`
      } `url:",omitempty" json:"scheme_identifiers,omitempty"`
      }




// CreditorCreateParams parameters
type CreditorCreateParams struct {
      AddressLine1 string `url:",omitempty" json:"address_line1,omitempty"`
      AddressLine2 string `url:",omitempty" json:"address_line2,omitempty"`
      AddressLine3 string `url:",omitempty" json:"address_line3,omitempty"`
      City string `url:",omitempty" json:"city,omitempty"`
      CountryCode string `url:",omitempty" json:"country_code,omitempty"`
      Links map[string]interface{} `url:",omitempty" json:"links,omitempty"`
      Name string `url:",omitempty" json:"name,omitempty"`
      PostalCode string `url:",omitempty" json:"postal_code,omitempty"`
      Region string `url:",omitempty" json:"region,omitempty"`
      }

// Create
// Creates a new creditor.
func (s *CreditorService) Create(ctx context.Context, p CreditorCreateParams) (*Creditor,error) {
  uri, err := url.Parse(fmt.Sprintf(s.endpoint + "/creditors",))
  if err != nil {
    return nil, err
  }

  var body io.Reader

  var buf bytes.Buffer
  err = json.NewEncoder(&buf).Encode(map[string]interface{}{
    "creditors": p,
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
Creditor *Creditor `json:"creditors"`
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

if result.Creditor == nil {
    return nil, errors.New("missing result")
  }

  return result.Creditor, nil
}


// CreditorListParams parameters
type CreditorListParams struct {
      After string `url:",omitempty" json:"after,omitempty"`
      Before string `url:",omitempty" json:"before,omitempty"`
      CreatedAt struct {
      Gt string `url:",omitempty" json:"gt,omitempty"`
      Gte string `url:",omitempty" json:"gte,omitempty"`
      Lt string `url:",omitempty" json:"lt,omitempty"`
      Lte string `url:",omitempty" json:"lte,omitempty"`
      } `url:",omitempty" json:"created_at,omitempty"`
      Limit int `url:",omitempty" json:"limit,omitempty"`
      }// CreditorListResult response including pagination metadata
type CreditorListResult struct {
      Creditors []struct {
      AddressLine1 string `url:",omitempty" json:"address_line1,omitempty"`
      AddressLine2 string `url:",omitempty" json:"address_line2,omitempty"`
      AddressLine3 string `url:",omitempty" json:"address_line3,omitempty"`
      City string `url:",omitempty" json:"city,omitempty"`
      CountryCode string `url:",omitempty" json:"country_code,omitempty"`
      CreatedAt string `url:",omitempty" json:"created_at,omitempty"`
      Id string `url:",omitempty" json:"id,omitempty"`
      Links struct {
      DefaultEurPayoutAccount string `url:",omitempty" json:"default_eur_payout_account,omitempty"`
      DefaultGbpPayoutAccount string `url:",omitempty" json:"default_gbp_payout_account,omitempty"`
      DefaultSekPayoutAccount string `url:",omitempty" json:"default_sek_payout_account,omitempty"`
      } `url:",omitempty" json:"links,omitempty"`
      LogoUrl string `url:",omitempty" json:"logo_url,omitempty"`
      Name string `url:",omitempty" json:"name,omitempty"`
      PostalCode string `url:",omitempty" json:"postal_code,omitempty"`
      Region string `url:",omitempty" json:"region,omitempty"`
      SchemeIdentifiers []struct {
      AddressLine1 string `url:",omitempty" json:"address_line1,omitempty"`
      AddressLine2 string `url:",omitempty" json:"address_line2,omitempty"`
      AddressLine3 string `url:",omitempty" json:"address_line3,omitempty"`
      CanSpecifyMandateReference bool `url:",omitempty" json:"can_specify_mandate_reference,omitempty"`
      City string `url:",omitempty" json:"city,omitempty"`
      CountryCode string `url:",omitempty" json:"country_code,omitempty"`
      Currency string `url:",omitempty" json:"currency,omitempty"`
      Email string `url:",omitempty" json:"email,omitempty"`
      MinimumAdvanceNotice int `url:",omitempty" json:"minimum_advance_notice,omitempty"`
      Name string `url:",omitempty" json:"name,omitempty"`
      PhoneNumber string `url:",omitempty" json:"phone_number,omitempty"`
      PostalCode string `url:",omitempty" json:"postal_code,omitempty"`
      Reference string `url:",omitempty" json:"reference,omitempty"`
      Region string `url:",omitempty" json:"region,omitempty"`
      Scheme string `url:",omitempty" json:"scheme,omitempty"`
      } `url:",omitempty" json:"scheme_identifiers,omitempty"`
      } `url:",omitempty" json:"creditors,omitempty"`
      Meta struct {
      Cursors struct {
      After string `url:",omitempty" json:"after,omitempty"`
      Before string `url:",omitempty" json:"before,omitempty"`
      } `url:",omitempty" json:"cursors,omitempty"`
      Limit int `url:",omitempty" json:"limit,omitempty"`
      } `url:",omitempty" json:"meta,omitempty"`
      }


// List
// Returns a [cursor-paginated](#api-usage-cursor-pagination) list of your
// creditors.
func (s *CreditorService) List(ctx context.Context, p CreditorListParams) (*CreditorListResult,error) {
  uri, err := url.Parse(fmt.Sprintf(s.endpoint + "/creditors",))
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
*CreditorListResult
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

if result.CreditorListResult == nil {
    return nil, errors.New("missing result")
  }

  return result.CreditorListResult, nil
}


// CreditorGetParams parameters
type CreditorGetParams map[string]interface{}

// Get
// Retrieves the details of an existing creditor.
func (s *CreditorService) Get(ctx context.Context,identity string, p CreditorGetParams) (*Creditor,error) {
  uri, err := url.Parse(fmt.Sprintf(s.endpoint + "/creditors/%v",
      identity,))
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
Creditor *Creditor `json:"creditors"`
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

if result.Creditor == nil {
    return nil, errors.New("missing result")
  }

  return result.Creditor, nil
}


// CreditorUpdateParams parameters
type CreditorUpdateParams struct {
      AddressLine1 string `url:",omitempty" json:"address_line1,omitempty"`
      AddressLine2 string `url:",omitempty" json:"address_line2,omitempty"`
      AddressLine3 string `url:",omitempty" json:"address_line3,omitempty"`
      City string `url:",omitempty" json:"city,omitempty"`
      CountryCode string `url:",omitempty" json:"country_code,omitempty"`
      Links struct {
      DefaultEurPayoutAccount string `url:",omitempty" json:"default_eur_payout_account,omitempty"`
      DefaultGbpPayoutAccount string `url:",omitempty" json:"default_gbp_payout_account,omitempty"`
      DefaultSekPayoutAccount string `url:",omitempty" json:"default_sek_payout_account,omitempty"`
      } `url:",omitempty" json:"links,omitempty"`
      Name string `url:",omitempty" json:"name,omitempty"`
      PostalCode string `url:",omitempty" json:"postal_code,omitempty"`
      Region string `url:",omitempty" json:"region,omitempty"`
      }

// Update
// Updates a creditor object. Supports all of the fields supported when creating
// a creditor.
func (s *CreditorService) Update(ctx context.Context,identity string, p CreditorUpdateParams) (*Creditor,error) {
  uri, err := url.Parse(fmt.Sprintf(s.endpoint + "/creditors/%v",
      identity,))
  if err != nil {
    return nil, err
  }

  var body io.Reader

  var buf bytes.Buffer
  err = json.NewEncoder(&buf).Encode(map[string]interface{}{
    "creditors": p,
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
Creditor *Creditor `json:"creditors"`
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

if result.Creditor == nil {
    return nil, errors.New("missing result")
  }

  return result.Creditor, nil
}

