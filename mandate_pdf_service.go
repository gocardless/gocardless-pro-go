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


type MandatePdfService struct {
  endpoint string
  token string
  client *http.Client
}



// MandatePdfCreateParams parameters
type MandatePdfCreateParams struct {
      AccountHolderName string `url:",omitempty" json:"account_holder_name,omitempty"`
        AccountNumber string `url:",omitempty" json:"account_number,omitempty"`
        BankCode string `url:",omitempty" json:"bank_code,omitempty"`
        Bic string `url:",omitempty" json:"bic,omitempty"`
        BranchCode string `url:",omitempty" json:"branch_code,omitempty"`
        CountryCode string `url:",omitempty" json:"country_code,omitempty"`
        Iban string `url:",omitempty" json:"iban,omitempty"`
        Links struct {
      Mandate string `url:",omitempty" json:"mandate,omitempty"`
        
    } `url:",omitempty" json:"links,omitempty"`
        MandateReference string `url:",omitempty" json:"mandate_reference,omitempty"`
        Scheme string `url:",omitempty" json:"scheme,omitempty"`
        SignatureDate string `url:",omitempty" json:"signature_date,omitempty"`
        SwedishIdentityNumber string `url:",omitempty" json:"swedish_identity_number,omitempty"`
        
    }
// MandatePdfCreateResult parameters
type MandatePdfCreateResult struct {
      MandatePdfs struct {
      ExpiresAt string `url:",omitempty" json:"expires_at,omitempty"`
        Url string `url:",omitempty" json:"url,omitempty"`
        
    } `url:",omitempty" json:"mandate_pdfs,omitempty"`
        
    }

// Create
// Generates a PDF mandate and returns its temporary URL.
// 
// Customer and
// bank account details can be left blank (for a blank mandate), provided
// manually, or inferred from the ID of an existing
// [mandate](#core-endpoints-mandates).
// 
// To generate a PDF mandate in a
// foreign language, set your `Accept-Language` header to the relevant [ISO
// 639-1](http://en.wikipedia.org/wiki/List_of_ISO_639-1_codes#Partial_ISO_639_table)
// language code. Supported languages are Dutch, English, French, German,
// Italian, Portuguese, Spanish and Swedish.
func (s *MandatePdfService) Create(
  ctx context.Context,
  p MandatePdfCreateParams) (*MandatePdfCreateResult, error) {
  uri, err := url.Parse(fmt.Sprintf(
      s.endpoint + "/mandate_pdfs",))
  if err != nil {
    return nil, err
  }

  var body io.Reader

  var buf bytes.Buffer
  err = json.NewEncoder(&buf).Encode(map[string]interface{}{
    "mandate_pdfs": p,
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
    *MandatePdfCreateResult
    Err *APIError `json:"error"`
  }

  err = json.NewDecoder(res.Body).Decode(&result)
  if err != nil {
    return nil, err
  }

  if result.Err != nil {
    return nil, result.Err
  }

  return result.MandatePdfCreateResult, nil
}

