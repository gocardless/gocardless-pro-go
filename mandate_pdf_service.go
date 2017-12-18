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


// MandatePdfService manages mandate_pdfs
type MandatePdfService struct {
  endpoint string
  token string
  client *http.Client
}


// MandatePdf model
type MandatePdf struct {
      ExpiresAt string `url:",omitempty" json:"expires_at,omitempty"`
      Url string `url:",omitempty" json:"url,omitempty"`
      }




// MandatePdfCreateParams parameters
type MandatePdfCreateParams struct {
      AccountHolderName string `url:",omitempty" json:"account_holder_name,omitempty"`
      AccountNumber string `url:",omitempty" json:"account_number,omitempty"`
      BankCode string `url:",omitempty" json:"bank_code,omitempty"`
      Bic string `url:",omitempty" json:"bic,omitempty"`
      BranchCode string `url:",omitempty" json:"branch_code,omitempty"`
      CountryCode string `url:",omitempty" json:"country_code,omitempty"`
      DanishIdentityNumber string `url:",omitempty" json:"danish_identity_number,omitempty"`
      Iban string `url:",omitempty" json:"iban,omitempty"`
      Links struct {
      Mandate string `url:",omitempty" json:"mandate,omitempty"`
      } `url:",omitempty" json:"links,omitempty"`
      MandateReference string `url:",omitempty" json:"mandate_reference,omitempty"`
      Scheme string `url:",omitempty" json:"scheme,omitempty"`
      SignatureDate string `url:",omitempty" json:"signature_date,omitempty"`
      SwedishIdentityNumber string `url:",omitempty" json:"swedish_identity_number,omitempty"`
      }

// Create
// Generates a PDF mandate and returns its temporary URL.
// 
// Customer and bank account details can be left blank (for a blank mandate),
// provided manually, or inferred from the ID of an existing
// [mandate](#core-endpoints-mandates).
// 
// To generate a PDF mandate in a foreign language, set your `Accept-Language`
// header to the relevant [ISO
// 639-1](http://en.wikipedia.org/wiki/List_of_ISO_639-1_codes#Partial_ISO_639_table)
// language code. Supported languages are Dutch, English, French, German,
// Italian, Portuguese, Spanish and Swedish.
func (s *MandatePdfService) Create(ctx context.Context, p MandatePdfCreateParams) (*MandatePdf,error) {
  uri, err := url.Parse(fmt.Sprintf(s.endpoint + "/mandate_pdfs",))
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
  req.Header.Set("Idempotency-Key", NewIdempotencyKey())

  client := s.client
  if client == nil {
    client = http.DefaultClient
  }

  var result struct {
    Err *APIError `json:"error"`
MandatePdf *MandatePdf `json:"mandate_pdfs"`
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

if result.MandatePdf == nil {
    return nil, errors.New("missing result")
  }

  return result.MandatePdf, nil
}

