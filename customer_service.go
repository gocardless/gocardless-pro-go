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


type CustomerService struct {
  endpoint string
  token string
  client *http.Client
}



// CustomerCreateParams parameters
type CustomerCreateParams struct {
      AddressLine1 string `json:"address_line1,omitempty"`
        AddressLine2 string `json:"address_line2,omitempty"`
        AddressLine3 string `json:"address_line3,omitempty"`
        City string `json:"city,omitempty"`
        CompanyName string `json:"company_name,omitempty"`
        CountryCode string `json:"country_code,omitempty"`
        Email string `json:"email,omitempty"`
        FamilyName string `json:"family_name,omitempty"`
        GivenName string `json:"given_name,omitempty"`
        Language string `json:"language,omitempty"`
        Metadata map[string]interface{} `json:"metadata,omitempty"`
        PostalCode string `json:"postal_code,omitempty"`
        Region string `json:"region,omitempty"`
        SwedishIdentityNumber string `json:"swedish_identity_number,omitempty"`
        
    }
// CustomerCreateResult parameters
type CustomerCreateResult struct {
      Customers struct {
      AddressLine1 string `json:"address_line1,omitempty"`
        AddressLine2 string `json:"address_line2,omitempty"`
        AddressLine3 string `json:"address_line3,omitempty"`
        City string `json:"city,omitempty"`
        CompanyName string `json:"company_name,omitempty"`
        CountryCode string `json:"country_code,omitempty"`
        CreatedAt string `json:"created_at,omitempty"`
        Email string `json:"email,omitempty"`
        FamilyName string `json:"family_name,omitempty"`
        GivenName string `json:"given_name,omitempty"`
        Id string `json:"id,omitempty"`
        Language string `json:"language,omitempty"`
        Metadata map[string]interface{} `json:"metadata,omitempty"`
        PostalCode string `json:"postal_code,omitempty"`
        Region string `json:"region,omitempty"`
        SwedishIdentityNumber string `json:"swedish_identity_number,omitempty"`
        
    } `json:"customers,omitempty"`
        
    }

// Create
// Creates a new customer object.
func (s *CustomerService) Create(
  ctx context.Context,
  p CustomerCreateParams) (*CustomerCreateResult, error) {
  uri, err := url.Parse(fmt.Sprintf(
      s.endpoint + "/customers",))
  if err != nil {
    return nil, err
  }

  var body io.Reader

  var buf bytes.Buffer
  err = json.NewEncoder(&buf).Encode(map[string]interface{}{
    "customers": p,
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
    *CustomerCreateResult
    Err *APIError `json:"error"`
  }

  err = json.NewDecoder(res.Body).Decode(&result)
  if err != nil {
    return nil, err
  }

  if result.Err != nil {
    return nil, result.Err
  }

  return result.CustomerCreateResult, nil
}


// CustomerListParams parameters
type CustomerListParams struct {
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
// CustomerListResult parameters
type CustomerListResult struct {
      Customers []struct {
      AddressLine1 string `json:"address_line1,omitempty"`
        AddressLine2 string `json:"address_line2,omitempty"`
        AddressLine3 string `json:"address_line3,omitempty"`
        City string `json:"city,omitempty"`
        CompanyName string `json:"company_name,omitempty"`
        CountryCode string `json:"country_code,omitempty"`
        CreatedAt string `json:"created_at,omitempty"`
        Email string `json:"email,omitempty"`
        FamilyName string `json:"family_name,omitempty"`
        GivenName string `json:"given_name,omitempty"`
        Id string `json:"id,omitempty"`
        Language string `json:"language,omitempty"`
        Metadata map[string]interface{} `json:"metadata,omitempty"`
        PostalCode string `json:"postal_code,omitempty"`
        Region string `json:"region,omitempty"`
        SwedishIdentityNumber string `json:"swedish_identity_number,omitempty"`
        
    } `json:"customers,omitempty"`
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
// customers.
func (s *CustomerService) List(
  ctx context.Context,
  p CustomerListParams) (*CustomerListResult, error) {
  uri, err := url.Parse(fmt.Sprintf(
      s.endpoint + "/customers",))
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
    *CustomerListResult
    Err *APIError `json:"error"`
  }

  err = json.NewDecoder(res.Body).Decode(&result)
  if err != nil {
    return nil, err
  }

  if result.Err != nil {
    return nil, result.Err
  }

  return result.CustomerListResult, nil
}


// CustomerGetResult parameters
type CustomerGetResult struct {
      Customers struct {
      AddressLine1 string `json:"address_line1,omitempty"`
        AddressLine2 string `json:"address_line2,omitempty"`
        AddressLine3 string `json:"address_line3,omitempty"`
        City string `json:"city,omitempty"`
        CompanyName string `json:"company_name,omitempty"`
        CountryCode string `json:"country_code,omitempty"`
        CreatedAt string `json:"created_at,omitempty"`
        Email string `json:"email,omitempty"`
        FamilyName string `json:"family_name,omitempty"`
        GivenName string `json:"given_name,omitempty"`
        Id string `json:"id,omitempty"`
        Language string `json:"language,omitempty"`
        Metadata map[string]interface{} `json:"metadata,omitempty"`
        PostalCode string `json:"postal_code,omitempty"`
        Region string `json:"region,omitempty"`
        SwedishIdentityNumber string `json:"swedish_identity_number,omitempty"`
        
    } `json:"customers,omitempty"`
        
    }

// Get
// Retrieves the details of an existing customer.
func (s *CustomerService) Get(
  ctx context.Context,
  identity string) (*CustomerGetResult, error) {
  uri, err := url.Parse(fmt.Sprintf(
      s.endpoint + "/customers/%v",
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
    *CustomerGetResult
    Err *APIError `json:"error"`
  }

  err = json.NewDecoder(res.Body).Decode(&result)
  if err != nil {
    return nil, err
  }

  if result.Err != nil {
    return nil, result.Err
  }

  return result.CustomerGetResult, nil
}


// CustomerUpdateParams parameters
type CustomerUpdateParams struct {
      AddressLine1 string `json:"address_line1,omitempty"`
        AddressLine2 string `json:"address_line2,omitempty"`
        AddressLine3 string `json:"address_line3,omitempty"`
        City string `json:"city,omitempty"`
        CompanyName string `json:"company_name,omitempty"`
        CountryCode string `json:"country_code,omitempty"`
        Email string `json:"email,omitempty"`
        FamilyName string `json:"family_name,omitempty"`
        GivenName string `json:"given_name,omitempty"`
        Language string `json:"language,omitempty"`
        Metadata map[string]interface{} `json:"metadata,omitempty"`
        PostalCode string `json:"postal_code,omitempty"`
        Region string `json:"region,omitempty"`
        SwedishIdentityNumber string `json:"swedish_identity_number,omitempty"`
        
    }
// CustomerUpdateResult parameters
type CustomerUpdateResult struct {
      Customers struct {
      AddressLine1 string `json:"address_line1,omitempty"`
        AddressLine2 string `json:"address_line2,omitempty"`
        AddressLine3 string `json:"address_line3,omitempty"`
        City string `json:"city,omitempty"`
        CompanyName string `json:"company_name,omitempty"`
        CountryCode string `json:"country_code,omitempty"`
        CreatedAt string `json:"created_at,omitempty"`
        Email string `json:"email,omitempty"`
        FamilyName string `json:"family_name,omitempty"`
        GivenName string `json:"given_name,omitempty"`
        Id string `json:"id,omitempty"`
        Language string `json:"language,omitempty"`
        Metadata map[string]interface{} `json:"metadata,omitempty"`
        PostalCode string `json:"postal_code,omitempty"`
        Region string `json:"region,omitempty"`
        SwedishIdentityNumber string `json:"swedish_identity_number,omitempty"`
        
    } `json:"customers,omitempty"`
        
    }

// Update
// Updates a customer object. Supports all of the fields supported when creating
// a customer.
func (s *CustomerService) Update(
  ctx context.Context,
  identity string,
  p CustomerUpdateParams) (*CustomerUpdateResult, error) {
  uri, err := url.Parse(fmt.Sprintf(
      s.endpoint + "/customers/%v",
      identity,))
  if err != nil {
    return nil, err
  }

  var body io.Reader

  var buf bytes.Buffer
  err = json.NewEncoder(&buf).Encode(map[string]interface{}{
    "customers": p,
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
    *CustomerUpdateResult
    Err *APIError `json:"error"`
  }

  err = json.NewDecoder(res.Body).Decode(&result)
  if err != nil {
    return nil, err
  }

  if result.Err != nil {
    return nil, result.Err
  }

  return result.CustomerUpdateResult, nil
}

