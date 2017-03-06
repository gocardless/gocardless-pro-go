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


type MandateService struct {
  endpoint string
  token string
  client *http.Client
}



// MandateCreateParams parameters
type MandateCreateParams struct {
      Links struct {
      Creditor string `url:",omitempty" json:"creditor,omitempty"`
        CustomerBankAccount string `url:",omitempty" json:"customer_bank_account,omitempty"`
        
    } `url:",omitempty" json:"links,omitempty"`
        Metadata map[string]interface{} `url:",omitempty" json:"metadata,omitempty"`
        Reference string `url:",omitempty" json:"reference,omitempty"`
        Scheme string `url:",omitempty" json:"scheme,omitempty"`
        
    }
// MandateCreateResult parameters
type MandateCreateResult struct {
      Mandates struct {
      CreatedAt string `url:",omitempty" json:"created_at,omitempty"`
        Id string `url:",omitempty" json:"id,omitempty"`
        Links struct {
      Creditor string `url:",omitempty" json:"creditor,omitempty"`
        Customer string `url:",omitempty" json:"customer,omitempty"`
        CustomerBankAccount string `url:",omitempty" json:"customer_bank_account,omitempty"`
        NewMandate string `url:",omitempty" json:"new_mandate,omitempty"`
        
    } `url:",omitempty" json:"links,omitempty"`
        Metadata map[string]interface{} `url:",omitempty" json:"metadata,omitempty"`
        NextPossibleChargeDate string `url:",omitempty" json:"next_possible_charge_date,omitempty"`
        PaymentsRequireApproval bool `url:",omitempty" json:"payments_require_approval,omitempty"`
        Reference string `url:",omitempty" json:"reference,omitempty"`
        Scheme string `url:",omitempty" json:"scheme,omitempty"`
        Status string `url:",omitempty" json:"status,omitempty"`
        
    } `url:",omitempty" json:"mandates,omitempty"`
        
    }

// Create
// Creates a new mandate object.
func (s *MandateService) Create(
  ctx context.Context,
  p MandateCreateParams) (*MandateCreateResult, error) {
  uri, err := url.Parse(fmt.Sprintf(
      s.endpoint + "/mandates",))
  if err != nil {
    return nil, err
  }

  var body io.Reader

  var buf bytes.Buffer
  err = json.NewEncoder(&buf).Encode(map[string]interface{}{
    "mandates": p,
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
    *MandateCreateResult
    Err *APIError `json:"error"`
  }

  err = json.NewDecoder(res.Body).Decode(&result)
  if err != nil {
    return nil, err
  }

  if result.Err != nil {
    return nil, result.Err
  }

  return result.MandateCreateResult, nil
}


// MandateListParams parameters
type MandateListParams struct {
      After string `url:",omitempty" json:"after,omitempty"`
        Before string `url:",omitempty" json:"before,omitempty"`
        CreatedAt struct {
      Gt string `url:",omitempty" json:"gt,omitempty"`
        Gte string `url:",omitempty" json:"gte,omitempty"`
        Lt string `url:",omitempty" json:"lt,omitempty"`
        Lte string `url:",omitempty" json:"lte,omitempty"`
        
    } `url:",omitempty" json:"created_at,omitempty"`
        Creditor string `url:",omitempty" json:"creditor,omitempty"`
        Customer string `url:",omitempty" json:"customer,omitempty"`
        CustomerBankAccount string `url:",omitempty" json:"customer_bank_account,omitempty"`
        Limit int `url:",omitempty" json:"limit,omitempty"`
        Reference string `url:",omitempty" json:"reference,omitempty"`
        Status []string `url:",omitempty" json:"status,omitempty"`
        
    }
// MandateListResult parameters
type MandateListResult struct {
      Mandates []struct {
      CreatedAt string `url:",omitempty" json:"created_at,omitempty"`
        Id string `url:",omitempty" json:"id,omitempty"`
        Links struct {
      Creditor string `url:",omitempty" json:"creditor,omitempty"`
        Customer string `url:",omitempty" json:"customer,omitempty"`
        CustomerBankAccount string `url:",omitempty" json:"customer_bank_account,omitempty"`
        NewMandate string `url:",omitempty" json:"new_mandate,omitempty"`
        
    } `url:",omitempty" json:"links,omitempty"`
        Metadata map[string]interface{} `url:",omitempty" json:"metadata,omitempty"`
        NextPossibleChargeDate string `url:",omitempty" json:"next_possible_charge_date,omitempty"`
        PaymentsRequireApproval bool `url:",omitempty" json:"payments_require_approval,omitempty"`
        Reference string `url:",omitempty" json:"reference,omitempty"`
        Scheme string `url:",omitempty" json:"scheme,omitempty"`
        Status string `url:",omitempty" json:"status,omitempty"`
        
    } `url:",omitempty" json:"mandates,omitempty"`
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
// mandates.
func (s *MandateService) List(
  ctx context.Context,
  p MandateListParams) (*MandateListResult, error) {
  uri, err := url.Parse(fmt.Sprintf(
      s.endpoint + "/mandates",))
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
    *MandateListResult
    Err *APIError `json:"error"`
  }

  err = json.NewDecoder(res.Body).Decode(&result)
  if err != nil {
    return nil, err
  }

  if result.Err != nil {
    return nil, result.Err
  }

  return result.MandateListResult, nil
}


// MandateGetResult parameters
type MandateGetResult struct {
      Mandates struct {
      CreatedAt string `url:",omitempty" json:"created_at,omitempty"`
        Id string `url:",omitempty" json:"id,omitempty"`
        Links struct {
      Creditor string `url:",omitempty" json:"creditor,omitempty"`
        Customer string `url:",omitempty" json:"customer,omitempty"`
        CustomerBankAccount string `url:",omitempty" json:"customer_bank_account,omitempty"`
        NewMandate string `url:",omitempty" json:"new_mandate,omitempty"`
        
    } `url:",omitempty" json:"links,omitempty"`
        Metadata map[string]interface{} `url:",omitempty" json:"metadata,omitempty"`
        NextPossibleChargeDate string `url:",omitempty" json:"next_possible_charge_date,omitempty"`
        PaymentsRequireApproval bool `url:",omitempty" json:"payments_require_approval,omitempty"`
        Reference string `url:",omitempty" json:"reference,omitempty"`
        Scheme string `url:",omitempty" json:"scheme,omitempty"`
        Status string `url:",omitempty" json:"status,omitempty"`
        
    } `url:",omitempty" json:"mandates,omitempty"`
        
    }

// Get
// Retrieves the details of an existing mandate.
func (s *MandateService) Get(
  ctx context.Context,
  identity string) (*MandateGetResult, error) {
  uri, err := url.Parse(fmt.Sprintf(
      s.endpoint + "/mandates/%v",
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
    *MandateGetResult
    Err *APIError `json:"error"`
  }

  err = json.NewDecoder(res.Body).Decode(&result)
  if err != nil {
    return nil, err
  }

  if result.Err != nil {
    return nil, result.Err
  }

  return result.MandateGetResult, nil
}


// MandateUpdateParams parameters
type MandateUpdateParams struct {
      Metadata map[string]interface{} `url:",omitempty" json:"metadata,omitempty"`
        
    }
// MandateUpdateResult parameters
type MandateUpdateResult struct {
      Mandates struct {
      CreatedAt string `url:",omitempty" json:"created_at,omitempty"`
        Id string `url:",omitempty" json:"id,omitempty"`
        Links struct {
      Creditor string `url:",omitempty" json:"creditor,omitempty"`
        Customer string `url:",omitempty" json:"customer,omitempty"`
        CustomerBankAccount string `url:",omitempty" json:"customer_bank_account,omitempty"`
        NewMandate string `url:",omitempty" json:"new_mandate,omitempty"`
        
    } `url:",omitempty" json:"links,omitempty"`
        Metadata map[string]interface{} `url:",omitempty" json:"metadata,omitempty"`
        NextPossibleChargeDate string `url:",omitempty" json:"next_possible_charge_date,omitempty"`
        PaymentsRequireApproval bool `url:",omitempty" json:"payments_require_approval,omitempty"`
        Reference string `url:",omitempty" json:"reference,omitempty"`
        Scheme string `url:",omitempty" json:"scheme,omitempty"`
        Status string `url:",omitempty" json:"status,omitempty"`
        
    } `url:",omitempty" json:"mandates,omitempty"`
        
    }

// Update
// Updates a mandate object. This accepts only the metadata parameter.
func (s *MandateService) Update(
  ctx context.Context,
  identity string,
  p MandateUpdateParams) (*MandateUpdateResult, error) {
  uri, err := url.Parse(fmt.Sprintf(
      s.endpoint + "/mandates/%v",
      identity,))
  if err != nil {
    return nil, err
  }

  var body io.Reader

  var buf bytes.Buffer
  err = json.NewEncoder(&buf).Encode(map[string]interface{}{
    "mandates": p,
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
    *MandateUpdateResult
    Err *APIError `json:"error"`
  }

  err = json.NewDecoder(res.Body).Decode(&result)
  if err != nil {
    return nil, err
  }

  if result.Err != nil {
    return nil, result.Err
  }

  return result.MandateUpdateResult, nil
}


// MandateCancelParams parameters
type MandateCancelParams struct {
      Metadata map[string]interface{} `url:",omitempty" json:"metadata,omitempty"`
        
    }
// MandateCancelResult parameters
type MandateCancelResult struct {
      Mandates struct {
      CreatedAt string `url:",omitempty" json:"created_at,omitempty"`
        Id string `url:",omitempty" json:"id,omitempty"`
        Links struct {
      Creditor string `url:",omitempty" json:"creditor,omitempty"`
        Customer string `url:",omitempty" json:"customer,omitempty"`
        CustomerBankAccount string `url:",omitempty" json:"customer_bank_account,omitempty"`
        NewMandate string `url:",omitempty" json:"new_mandate,omitempty"`
        
    } `url:",omitempty" json:"links,omitempty"`
        Metadata map[string]interface{} `url:",omitempty" json:"metadata,omitempty"`
        NextPossibleChargeDate string `url:",omitempty" json:"next_possible_charge_date,omitempty"`
        PaymentsRequireApproval bool `url:",omitempty" json:"payments_require_approval,omitempty"`
        Reference string `url:",omitempty" json:"reference,omitempty"`
        Scheme string `url:",omitempty" json:"scheme,omitempty"`
        Status string `url:",omitempty" json:"status,omitempty"`
        
    } `url:",omitempty" json:"mandates,omitempty"`
        
    }

// Cancel
// Immediately cancels a mandate and all associated cancellable payments. Any
// metadata supplied to this endpoint will be stored on the mandate cancellation
// event it causes.
// 
// This will fail with a `cancellation_failed` error if
// the mandate is already cancelled.
func (s *MandateService) Cancel(
  ctx context.Context,
  identity string,
  p MandateCancelParams) (*MandateCancelResult, error) {
  uri, err := url.Parse(fmt.Sprintf(
      s.endpoint + "/mandates/%v/actions/cancel",
      identity,))
  if err != nil {
    return nil, err
  }

  var body io.Reader

  var buf bytes.Buffer
  err = json.NewEncoder(&buf).Encode(map[string]interface{}{
    "mandates": p,
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
    *MandateCancelResult
    Err *APIError `json:"error"`
  }

  err = json.NewDecoder(res.Body).Decode(&result)
  if err != nil {
    return nil, err
  }

  if result.Err != nil {
    return nil, result.Err
  }

  return result.MandateCancelResult, nil
}


// MandateReinstateParams parameters
type MandateReinstateParams struct {
      Metadata map[string]interface{} `url:",omitempty" json:"metadata,omitempty"`
        
    }
// MandateReinstateResult parameters
type MandateReinstateResult struct {
      Mandates struct {
      CreatedAt string `url:",omitempty" json:"created_at,omitempty"`
        Id string `url:",omitempty" json:"id,omitempty"`
        Links struct {
      Creditor string `url:",omitempty" json:"creditor,omitempty"`
        Customer string `url:",omitempty" json:"customer,omitempty"`
        CustomerBankAccount string `url:",omitempty" json:"customer_bank_account,omitempty"`
        NewMandate string `url:",omitempty" json:"new_mandate,omitempty"`
        
    } `url:",omitempty" json:"links,omitempty"`
        Metadata map[string]interface{} `url:",omitempty" json:"metadata,omitempty"`
        NextPossibleChargeDate string `url:",omitempty" json:"next_possible_charge_date,omitempty"`
        PaymentsRequireApproval bool `url:",omitempty" json:"payments_require_approval,omitempty"`
        Reference string `url:",omitempty" json:"reference,omitempty"`
        Scheme string `url:",omitempty" json:"scheme,omitempty"`
        Status string `url:",omitempty" json:"status,omitempty"`
        
    } `url:",omitempty" json:"mandates,omitempty"`
        
    }

// Reinstate
// <a name="mandate_not_inactive"></a>Reinstates a cancelled or expired mandate
// to the banks. You will receive a `resubmission_requested` webhook, but after
// that reinstating the mandate follows the same process as its initial
// creation, so you will receive a `submitted` webhook, followed by a
// `reinstated` or `failed` webhook up to two working days later. Any metadata
// supplied to this endpoint will be stored on the `resubmission_requested`
// event it causes.
// 
// This will fail with a `mandate_not_inactive` error if
// the mandate is already being submitted, or is active.
// 
// Mandates can be
// resubmitted up to 3 times.
func (s *MandateService) Reinstate(
  ctx context.Context,
  identity string,
  p MandateReinstateParams) (*MandateReinstateResult, error) {
  uri, err := url.Parse(fmt.Sprintf(
      s.endpoint + "/mandates/%v/actions/reinstate",
      identity,))
  if err != nil {
    return nil, err
  }

  var body io.Reader

  var buf bytes.Buffer
  err = json.NewEncoder(&buf).Encode(map[string]interface{}{
    "mandates": p,
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
    *MandateReinstateResult
    Err *APIError `json:"error"`
  }

  err = json.NewDecoder(res.Body).Decode(&result)
  if err != nil {
    return nil, err
  }

  if result.Err != nil {
    return nil, result.Err
  }

  return result.MandateReinstateResult, nil
}

