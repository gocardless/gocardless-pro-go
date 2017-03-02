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


type RedirectFlowService struct {
  endpoint string
  token string
  client *http.Client
}



// RedirectFlowCreateParams parameters
type RedirectFlowCreateParams struct {
      Description string `json:"description,omitempty"`
        Links struct {
      Creditor string `json:"creditor,omitempty"`
        
    } `json:"links,omitempty"`
        PrefilledCustomer struct {
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
        PostalCode string `json:"postal_code,omitempty"`
        Region string `json:"region,omitempty"`
        SwedishIdentityNumber string `json:"swedish_identity_number,omitempty"`
        
    } `json:"prefilled_customer,omitempty"`
        Scheme string `json:"scheme,omitempty"`
        SessionToken string `json:"session_token,omitempty"`
        SuccessRedirectUrl string `json:"success_redirect_url,omitempty"`
        
    }
// RedirectFlowCreateResult parameters
type RedirectFlowCreateResult struct {
      RedirectFlows struct {
      CreatedAt string `json:"created_at,omitempty"`
        Description string `json:"description,omitempty"`
        Id string `json:"id,omitempty"`
        Links struct {
      Creditor string `json:"creditor,omitempty"`
        Customer string `json:"customer,omitempty"`
        CustomerBankAccount string `json:"customer_bank_account,omitempty"`
        Mandate string `json:"mandate,omitempty"`
        
    } `json:"links,omitempty"`
        RedirectUrl string `json:"redirect_url,omitempty"`
        Scheme string `json:"scheme,omitempty"`
        SessionToken string `json:"session_token,omitempty"`
        SuccessRedirectUrl string `json:"success_redirect_url,omitempty"`
        
    } `json:"redirect_flows,omitempty"`
        
    }

// Create
// Creates a redirect flow object which can then be used to redirect your
// customer to the GoCardless hosted payment pages.
func (s *RedirectFlowService) Create(
  ctx context.Context,
  p RedirectFlowCreateParams) (*RedirectFlowCreateResult, error) {
  uri, err := url.Parse(fmt.Sprintf(
      s.endpoint + "/redirect_flows",))
  if err != nil {
    return nil, err
  }

  var body io.Reader

  var buf bytes.Buffer
  err = json.NewEncoder(&buf).Encode(map[string]interface{}{
    "redirect_flows": p,
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
    *RedirectFlowCreateResult
    Err *APIError `json:"error"`
  }

  err = json.NewDecoder(res.Body).Decode(&result)
  if err != nil {
    return nil, err
  }

  if result.Err != nil {
    return nil, result.Err
  }

  return result.RedirectFlowCreateResult, nil
}


// RedirectFlowGetResult parameters
type RedirectFlowGetResult struct {
      RedirectFlows struct {
      CreatedAt string `json:"created_at,omitempty"`
        Description string `json:"description,omitempty"`
        Id string `json:"id,omitempty"`
        Links struct {
      Creditor string `json:"creditor,omitempty"`
        Customer string `json:"customer,omitempty"`
        CustomerBankAccount string `json:"customer_bank_account,omitempty"`
        Mandate string `json:"mandate,omitempty"`
        
    } `json:"links,omitempty"`
        RedirectUrl string `json:"redirect_url,omitempty"`
        Scheme string `json:"scheme,omitempty"`
        SessionToken string `json:"session_token,omitempty"`
        SuccessRedirectUrl string `json:"success_redirect_url,omitempty"`
        
    } `json:"redirect_flows,omitempty"`
        
    }

// Get
// Returns all details about a single redirect flow
func (s *RedirectFlowService) Get(
  ctx context.Context,
  identity string) (*RedirectFlowGetResult, error) {
  uri, err := url.Parse(fmt.Sprintf(
      s.endpoint + "/redirect_flows/%v",
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
    *RedirectFlowGetResult
    Err *APIError `json:"error"`
  }

  err = json.NewDecoder(res.Body).Decode(&result)
  if err != nil {
    return nil, err
  }

  if result.Err != nil {
    return nil, result.Err
  }

  return result.RedirectFlowGetResult, nil
}


// RedirectFlowCompleteParams parameters
type RedirectFlowCompleteParams struct {
      SessionToken string `json:"session_token,omitempty"`
        
    }
// RedirectFlowCompleteResult parameters
type RedirectFlowCompleteResult struct {
      RedirectFlows struct {
      CreatedAt string `json:"created_at,omitempty"`
        Description string `json:"description,omitempty"`
        Id string `json:"id,omitempty"`
        Links struct {
      Creditor string `json:"creditor,omitempty"`
        Customer string `json:"customer,omitempty"`
        CustomerBankAccount string `json:"customer_bank_account,omitempty"`
        Mandate string `json:"mandate,omitempty"`
        
    } `json:"links,omitempty"`
        RedirectUrl string `json:"redirect_url,omitempty"`
        Scheme string `json:"scheme,omitempty"`
        SessionToken string `json:"session_token,omitempty"`
        SuccessRedirectUrl string `json:"success_redirect_url,omitempty"`
        
    } `json:"redirect_flows,omitempty"`
        
    }

// Complete
// This creates a [customer](#core-endpoints-customers), [customer bank
// account](#core-endpoints-customer-bank-accounts), and
// [mandate](#core-endpoints-mandates) using the details supplied by your
// customer and returns the ID of the created mandate.
// 
// This will return a
// `redirect_flow_incomplete` error if your customer has not yet been redirected
// back to your site, and a `redirect_flow_already_completed` error if your
// integration has already completed this flow. It will return a `bad_request`
// error if the `session_token` differs to the one supplied when the redirect
// flow was created.
func (s *RedirectFlowService) Complete(
  ctx context.Context,
  identity string,
  p RedirectFlowCompleteParams) (*RedirectFlowCompleteResult, error) {
  uri, err := url.Parse(fmt.Sprintf(
      s.endpoint + "/redirect_flows/%v/actions/complete",
      identity,))
  if err != nil {
    return nil, err
  }

  var body io.Reader

  var buf bytes.Buffer
  err = json.NewEncoder(&buf).Encode(map[string]interface{}{
    "redirect_flows": p,
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
    *RedirectFlowCompleteResult
    Err *APIError `json:"error"`
  }

  err = json.NewDecoder(res.Body).Decode(&result)
  if err != nil {
    return nil, err
  }

  if result.Err != nil {
    return nil, result.Err
  }

  return result.RedirectFlowCompleteResult, nil
}

