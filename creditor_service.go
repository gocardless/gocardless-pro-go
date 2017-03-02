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


type CreditorService struct {
  endpoint string
  token string
  client *http.Client
}



// CreditorCreateParams parameters
type CreditorCreateParams struct {
      AddressLine1 string `json:"address_line1,omitempty"`
        AddressLine2 string `json:"address_line2,omitempty"`
        AddressLine3 string `json:"address_line3,omitempty"`
        City string `json:"city,omitempty"`
        CountryCode string `json:"country_code,omitempty"`
        Links map[string]interface{} `json:"links,omitempty"`
        Name string `json:"name,omitempty"`
        PostalCode string `json:"postal_code,omitempty"`
        Region string `json:"region,omitempty"`
        
    }
// CreditorCreateResult parameters
type CreditorCreateResult struct {
      Creditors struct {
      AddressLine1 string `json:"address_line1,omitempty"`
        AddressLine2 string `json:"address_line2,omitempty"`
        AddressLine3 string `json:"address_line3,omitempty"`
        City string `json:"city,omitempty"`
        CountryCode string `json:"country_code,omitempty"`
        CreatedAt string `json:"created_at,omitempty"`
        Id string `json:"id,omitempty"`
        Links struct {
      DefaultEurPayoutAccount string `json:"default_eur_payout_account,omitempty"`
        DefaultGbpPayoutAccount string `json:"default_gbp_payout_account,omitempty"`
        DefaultSekPayoutAccount string `json:"default_sek_payout_account,omitempty"`
        
    } `json:"links,omitempty"`
        LogoUrl string `json:"logo_url,omitempty"`
        Name string `json:"name,omitempty"`
        PostalCode string `json:"postal_code,omitempty"`
        Region string `json:"region,omitempty"`
        SchemeIdentifiers []struct {
      AddressLine1 string `json:"address_line1,omitempty"`
        AddressLine2 string `json:"address_line2,omitempty"`
        AddressLine3 string `json:"address_line3,omitempty"`
        CanSpecifyMandateReference bool `json:"can_specify_mandate_reference,omitempty"`
        City string `json:"city,omitempty"`
        CountryCode string `json:"country_code,omitempty"`
        Currency string `json:"currency,omitempty"`
        Email string `json:"email,omitempty"`
        MinimumAdvanceNotice int `json:"minimum_advance_notice,omitempty"`
        Name string `json:"name,omitempty"`
        PhoneNumber string `json:"phone_number,omitempty"`
        PostalCode string `json:"postal_code,omitempty"`
        Reference string `json:"reference,omitempty"`
        Region string `json:"region,omitempty"`
        Scheme string `json:"scheme,omitempty"`
        
    } `json:"scheme_identifiers,omitempty"`
        
    } `json:"creditors,omitempty"`
        
    }

// Create
// Creates a new creditor.
func (s *CreditorService) Create(
  ctx context.Context,
  p CreditorCreateParams) (*CreditorCreateResult, error) {
  uri, err := url.Parse(fmt.Sprintf(
      s.endpoint + "/creditors",))
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
    *CreditorCreateResult
    Err *APIError `json:"error"`
  }

  err = json.NewDecoder(res.Body).Decode(&result)
  if err != nil {
    return nil, err
  }

  if result.Err != nil {
    return nil, result.Err
  }

  return result.CreditorCreateResult, nil
}


// CreditorListParams parameters
type CreditorListParams struct {
      After string `json:"after,omitempty"`
        Before string `json:"before,omitempty"`
        CreatedAt struct {
      Gt string `json:"gt,omitempty"`
        Gte string `json:"gte,omitempty"`
        Lt string `json:"lt,omitempty"`
        Lte string `json:"lte,omitempty"`
        
    } `json:"created_at,omitempty"`
        Limit string `json:"limit,omitempty"`
        
    }
// CreditorListResult parameters
type CreditorListResult struct {
      Creditors []struct {
      AddressLine1 string `json:"address_line1,omitempty"`
        AddressLine2 string `json:"address_line2,omitempty"`
        AddressLine3 string `json:"address_line3,omitempty"`
        City string `json:"city,omitempty"`
        CountryCode string `json:"country_code,omitempty"`
        CreatedAt string `json:"created_at,omitempty"`
        Id string `json:"id,omitempty"`
        Links struct {
      DefaultEurPayoutAccount string `json:"default_eur_payout_account,omitempty"`
        DefaultGbpPayoutAccount string `json:"default_gbp_payout_account,omitempty"`
        DefaultSekPayoutAccount string `json:"default_sek_payout_account,omitempty"`
        
    } `json:"links,omitempty"`
        LogoUrl string `json:"logo_url,omitempty"`
        Name string `json:"name,omitempty"`
        PostalCode string `json:"postal_code,omitempty"`
        Region string `json:"region,omitempty"`
        SchemeIdentifiers []struct {
      AddressLine1 string `json:"address_line1,omitempty"`
        AddressLine2 string `json:"address_line2,omitempty"`
        AddressLine3 string `json:"address_line3,omitempty"`
        CanSpecifyMandateReference bool `json:"can_specify_mandate_reference,omitempty"`
        City string `json:"city,omitempty"`
        CountryCode string `json:"country_code,omitempty"`
        Currency string `json:"currency,omitempty"`
        Email string `json:"email,omitempty"`
        MinimumAdvanceNotice int `json:"minimum_advance_notice,omitempty"`
        Name string `json:"name,omitempty"`
        PhoneNumber string `json:"phone_number,omitempty"`
        PostalCode string `json:"postal_code,omitempty"`
        Reference string `json:"reference,omitempty"`
        Region string `json:"region,omitempty"`
        Scheme string `json:"scheme,omitempty"`
        
    } `json:"scheme_identifiers,omitempty"`
        
    } `json:"creditors,omitempty"`
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
// creditors.
func (s *CreditorService) List(
  ctx context.Context,
  p CreditorListParams) (*CreditorListResult, error) {
  uri, err := url.Parse(fmt.Sprintf(
      s.endpoint + "/creditors",))
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
    *CreditorListResult
    Err *APIError `json:"error"`
  }

  err = json.NewDecoder(res.Body).Decode(&result)
  if err != nil {
    return nil, err
  }

  if result.Err != nil {
    return nil, result.Err
  }

  return result.CreditorListResult, nil
}


// CreditorGetParams parameters
type CreditorGetParams map[string]interface{}
// CreditorGetResult parameters
type CreditorGetResult struct {
      Creditors struct {
      AddressLine1 string `json:"address_line1,omitempty"`
        AddressLine2 string `json:"address_line2,omitempty"`
        AddressLine3 string `json:"address_line3,omitempty"`
        City string `json:"city,omitempty"`
        CountryCode string `json:"country_code,omitempty"`
        CreatedAt string `json:"created_at,omitempty"`
        Id string `json:"id,omitempty"`
        Links struct {
      DefaultEurPayoutAccount string `json:"default_eur_payout_account,omitempty"`
        DefaultGbpPayoutAccount string `json:"default_gbp_payout_account,omitempty"`
        DefaultSekPayoutAccount string `json:"default_sek_payout_account,omitempty"`
        
    } `json:"links,omitempty"`
        LogoUrl string `json:"logo_url,omitempty"`
        Name string `json:"name,omitempty"`
        PostalCode string `json:"postal_code,omitempty"`
        Region string `json:"region,omitempty"`
        SchemeIdentifiers []struct {
      AddressLine1 string `json:"address_line1,omitempty"`
        AddressLine2 string `json:"address_line2,omitempty"`
        AddressLine3 string `json:"address_line3,omitempty"`
        CanSpecifyMandateReference bool `json:"can_specify_mandate_reference,omitempty"`
        City string `json:"city,omitempty"`
        CountryCode string `json:"country_code,omitempty"`
        Currency string `json:"currency,omitempty"`
        Email string `json:"email,omitempty"`
        MinimumAdvanceNotice int `json:"minimum_advance_notice,omitempty"`
        Name string `json:"name,omitempty"`
        PhoneNumber string `json:"phone_number,omitempty"`
        PostalCode string `json:"postal_code,omitempty"`
        Reference string `json:"reference,omitempty"`
        Region string `json:"region,omitempty"`
        Scheme string `json:"scheme,omitempty"`
        
    } `json:"scheme_identifiers,omitempty"`
        
    } `json:"creditors,omitempty"`
        
    }

// Get
// Retrieves the details of an existing creditor.
func (s *CreditorService) Get(
  ctx context.Context,
  identity string,
  p CreditorGetParams) (*CreditorGetResult, error) {
  uri, err := url.Parse(fmt.Sprintf(
      s.endpoint + "/creditors/%v",
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

  res, err := client.Do(req)
  if err != nil {
    return nil, err
  }
  defer res.Body.Close()

  var result struct {
    *CreditorGetResult
    Err *APIError `json:"error"`
  }

  err = json.NewDecoder(res.Body).Decode(&result)
  if err != nil {
    return nil, err
  }

  if result.Err != nil {
    return nil, result.Err
  }

  return result.CreditorGetResult, nil
}


// CreditorUpdateParams parameters
type CreditorUpdateParams struct {
      AddressLine1 string `json:"address_line1,omitempty"`
        AddressLine2 string `json:"address_line2,omitempty"`
        AddressLine3 string `json:"address_line3,omitempty"`
        City string `json:"city,omitempty"`
        CountryCode string `json:"country_code,omitempty"`
        Links struct {
      DefaultEurPayoutAccount string `json:"default_eur_payout_account,omitempty"`
        DefaultGbpPayoutAccount string `json:"default_gbp_payout_account,omitempty"`
        DefaultSekPayoutAccount string `json:"default_sek_payout_account,omitempty"`
        
    } `json:"links,omitempty"`
        Name string `json:"name,omitempty"`
        PostalCode string `json:"postal_code,omitempty"`
        Region string `json:"region,omitempty"`
        
    }
// CreditorUpdateResult parameters
type CreditorUpdateResult struct {
      Creditors struct {
      AddressLine1 string `json:"address_line1,omitempty"`
        AddressLine2 string `json:"address_line2,omitempty"`
        AddressLine3 string `json:"address_line3,omitempty"`
        City string `json:"city,omitempty"`
        CountryCode string `json:"country_code,omitempty"`
        CreatedAt string `json:"created_at,omitempty"`
        Id string `json:"id,omitempty"`
        Links struct {
      DefaultEurPayoutAccount string `json:"default_eur_payout_account,omitempty"`
        DefaultGbpPayoutAccount string `json:"default_gbp_payout_account,omitempty"`
        DefaultSekPayoutAccount string `json:"default_sek_payout_account,omitempty"`
        
    } `json:"links,omitempty"`
        LogoUrl string `json:"logo_url,omitempty"`
        Name string `json:"name,omitempty"`
        PostalCode string `json:"postal_code,omitempty"`
        Region string `json:"region,omitempty"`
        SchemeIdentifiers []struct {
      AddressLine1 string `json:"address_line1,omitempty"`
        AddressLine2 string `json:"address_line2,omitempty"`
        AddressLine3 string `json:"address_line3,omitempty"`
        CanSpecifyMandateReference bool `json:"can_specify_mandate_reference,omitempty"`
        City string `json:"city,omitempty"`
        CountryCode string `json:"country_code,omitempty"`
        Currency string `json:"currency,omitempty"`
        Email string `json:"email,omitempty"`
        MinimumAdvanceNotice int `json:"minimum_advance_notice,omitempty"`
        Name string `json:"name,omitempty"`
        PhoneNumber string `json:"phone_number,omitempty"`
        PostalCode string `json:"postal_code,omitempty"`
        Reference string `json:"reference,omitempty"`
        Region string `json:"region,omitempty"`
        Scheme string `json:"scheme,omitempty"`
        
    } `json:"scheme_identifiers,omitempty"`
        
    } `json:"creditors,omitempty"`
        
    }

// Update
// Updates a creditor object. Supports all of the fields supported when creating
// a creditor.
func (s *CreditorService) Update(
  ctx context.Context,
  identity string,
  p CreditorUpdateParams) (*CreditorUpdateResult, error) {
  uri, err := url.Parse(fmt.Sprintf(
      s.endpoint + "/creditors/%v",
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
    *CreditorUpdateResult
    Err *APIError `json:"error"`
  }

  err = json.NewDecoder(res.Body).Decode(&result)
  if err != nil {
    return nil, err
  }

  if result.Err != nil {
    return nil, result.Err
  }

  return result.CreditorUpdateResult, nil
}

