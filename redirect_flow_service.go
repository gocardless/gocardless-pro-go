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


// RedirectFlowService manages redirect_flows
type RedirectFlowService struct {
  endpoint string
  token string
  client *http.Client
}


// RedirectFlow model
type RedirectFlow struct {
      ConfirmationUrl string `url:",omitempty" json:"confirmation_url,omitempty"`
      CreatedAt string `url:",omitempty" json:"created_at,omitempty"`
      Description string `url:",omitempty" json:"description,omitempty"`
      Id string `url:",omitempty" json:"id,omitempty"`
      Links struct {
      Creditor string `url:",omitempty" json:"creditor,omitempty"`
      Customer string `url:",omitempty" json:"customer,omitempty"`
      CustomerBankAccount string `url:",omitempty" json:"customer_bank_account,omitempty"`
      Mandate string `url:",omitempty" json:"mandate,omitempty"`
      } `url:",omitempty" json:"links,omitempty"`
      RedirectUrl string `url:",omitempty" json:"redirect_url,omitempty"`
      Scheme string `url:",omitempty" json:"scheme,omitempty"`
      SessionToken string `url:",omitempty" json:"session_token,omitempty"`
      SuccessRedirectUrl string `url:",omitempty" json:"success_redirect_url,omitempty"`
      }




// RedirectFlowCreateParams parameters
type RedirectFlowCreateParams struct {
      Description string `url:",omitempty" json:"description,omitempty"`
      Links struct {
      Creditor string `url:",omitempty" json:"creditor,omitempty"`
      } `url:",omitempty" json:"links,omitempty"`
      PrefilledCustomer struct {
      AddressLine1 string `url:",omitempty" json:"address_line1,omitempty"`
      AddressLine2 string `url:",omitempty" json:"address_line2,omitempty"`
      AddressLine3 string `url:",omitempty" json:"address_line3,omitempty"`
      City string `url:",omitempty" json:"city,omitempty"`
      CompanyName string `url:",omitempty" json:"company_name,omitempty"`
      CountryCode string `url:",omitempty" json:"country_code,omitempty"`
      Email string `url:",omitempty" json:"email,omitempty"`
      FamilyName string `url:",omitempty" json:"family_name,omitempty"`
      GivenName string `url:",omitempty" json:"given_name,omitempty"`
      Language string `url:",omitempty" json:"language,omitempty"`
      PostalCode string `url:",omitempty" json:"postal_code,omitempty"`
      Region string `url:",omitempty" json:"region,omitempty"`
      SwedishIdentityNumber string `url:",omitempty" json:"swedish_identity_number,omitempty"`
      } `url:",omitempty" json:"prefilled_customer,omitempty"`
      Scheme string `url:",omitempty" json:"scheme,omitempty"`
      SessionToken string `url:",omitempty" json:"session_token,omitempty"`
      SuccessRedirectUrl string `url:",omitempty" json:"success_redirect_url,omitempty"`
      }

// Create
// Creates a redirect flow object which can then be used to redirect your
// customer to the GoCardless hosted payment pages.
func (s *RedirectFlowService) Create(ctx context.Context, p RedirectFlowCreateParams) (*RedirectFlow,error) {
  uri, err := url.Parse(fmt.Sprintf(s.endpoint + "/redirect_flows",))
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
  req.Header.Set("Idempotency-Key", NewIdempotencyKey())

  client := s.client
  if client == nil {
    client = http.DefaultClient
  }

  var result struct {
    Err *APIError `json:"error"`
RedirectFlow *RedirectFlow `json:"redirect_flows"`
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

if result.RedirectFlow == nil {
    return nil, errors.New("missing result")
  }

  return result.RedirectFlow, nil
}



// Get
// Returns all details about a single redirect flow
func (s *RedirectFlowService) Get(ctx context.Context,identity string) (*RedirectFlow,error) {
  uri, err := url.Parse(fmt.Sprintf(s.endpoint + "/redirect_flows/%v",
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
RedirectFlow *RedirectFlow `json:"redirect_flows"`
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

if result.RedirectFlow == nil {
    return nil, errors.New("missing result")
  }

  return result.RedirectFlow, nil
}


// RedirectFlowCompleteParams parameters
type RedirectFlowCompleteParams struct {
      SessionToken string `url:",omitempty" json:"session_token,omitempty"`
      }

// Complete
// This creates a [customer](#core-endpoints-customers), [customer bank
// account](#core-endpoints-customer-bank-accounts), and
// [mandate](#core-endpoints-mandates) using the details supplied by your
// customer and returns the ID of the created mandate.
// 
// This will return a `redirect_flow_incomplete` error if your customer has not
// yet been redirected back to your site, and a
// `redirect_flow_already_completed` error if your integration has already
// completed this flow. It will return a `bad_request` error if the
// `session_token` differs to the one supplied when the redirect flow was
// created.
func (s *RedirectFlowService) Complete(ctx context.Context,identity string, p RedirectFlowCompleteParams) (*RedirectFlow,error) {
  uri, err := url.Parse(fmt.Sprintf(s.endpoint + "/redirect_flows/%v/actions/complete",
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
RedirectFlow *RedirectFlow `json:"redirect_flows"`
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

if result.RedirectFlow == nil {
    return nil, errors.New("missing result")
  }

  return result.RedirectFlow, nil
}

