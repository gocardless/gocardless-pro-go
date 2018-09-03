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


// MandateImportEntryService manages mandate_import_entries
type MandateImportEntryService struct {
  endpoint string
  token string
  client *http.Client
}


// MandateImportEntry model
type MandateImportEntry struct {
      CreatedAt string `url:",omitempty" json:"created_at,omitempty"`
      Links struct {
      Customer string `url:",omitempty" json:"customer,omitempty"`
      CustomerBankAccount string `url:",omitempty" json:"customer_bank_account,omitempty"`
      Mandate string `url:",omitempty" json:"mandate,omitempty"`
      MandateImport string `url:",omitempty" json:"mandate_import,omitempty"`
      } `url:",omitempty" json:"links,omitempty"`
      RecordIdentifier string `url:",omitempty" json:"record_identifier,omitempty"`
      }




// MandateImportEntryCreateParams parameters
type MandateImportEntryCreateParams struct {
      Amendment struct {
      OriginalCreditorId string `url:",omitempty" json:"original_creditor_id,omitempty"`
      OriginalCreditorName string `url:",omitempty" json:"original_creditor_name,omitempty"`
      OriginalMandateReference string `url:",omitempty" json:"original_mandate_reference,omitempty"`
      } `url:",omitempty" json:"amendment,omitempty"`
      BankAccount struct {
      AccountHolderName string `url:",omitempty" json:"account_holder_name,omitempty"`
      AccountNumber string `url:",omitempty" json:"account_number,omitempty"`
      BankCode string `url:",omitempty" json:"bank_code,omitempty"`
      BranchCode string `url:",omitempty" json:"branch_code,omitempty"`
      CountryCode string `url:",omitempty" json:"country_code,omitempty"`
      Iban string `url:",omitempty" json:"iban,omitempty"`
      } `url:",omitempty" json:"bank_account,omitempty"`
      Customer struct {
      AddressLine1 string `url:",omitempty" json:"address_line1,omitempty"`
      AddressLine2 string `url:",omitempty" json:"address_line2,omitempty"`
      AddressLine3 string `url:",omitempty" json:"address_line3,omitempty"`
      City string `url:",omitempty" json:"city,omitempty"`
      CompanyName string `url:",omitempty" json:"company_name,omitempty"`
      CountryCode string `url:",omitempty" json:"country_code,omitempty"`
      DanishIdentityNumber string `url:",omitempty" json:"danish_identity_number,omitempty"`
      Email string `url:",omitempty" json:"email,omitempty"`
      FamilyName string `url:",omitempty" json:"family_name,omitempty"`
      GivenName string `url:",omitempty" json:"given_name,omitempty"`
      Language string `url:",omitempty" json:"language,omitempty"`
      PostalCode string `url:",omitempty" json:"postal_code,omitempty"`
      Region string `url:",omitempty" json:"region,omitempty"`
      SwedishIdentityNumber string `url:",omitempty" json:"swedish_identity_number,omitempty"`
      } `url:",omitempty" json:"customer,omitempty"`
      Links struct {
      MandateImport string `url:",omitempty" json:"mandate_import,omitempty"`
      } `url:",omitempty" json:"links,omitempty"`
      RecordIdentifier string `url:",omitempty" json:"record_identifier,omitempty"`
      }

// Create
// For an existing [mandate import](#core-endpoints-mandate-imports), this
// endpoint can
// be used to add individual mandates to be imported into GoCardless.
// 
// You can add no more than 30,000 rows to a single mandate import.
// If you attempt to go over this limit, the API will return a
// `record_limit_exceeded` error.
func (s *MandateImportEntryService) Create(ctx context.Context, p MandateImportEntryCreateParams) (*MandateImportEntry,error) {
  uri, err := url.Parse(fmt.Sprintf(s.endpoint + "/mandate_import_entries",))
  if err != nil {
    return nil, err
  }

  var body io.Reader

  var buf bytes.Buffer
  err = json.NewEncoder(&buf).Encode(map[string]interface{}{
    "mandate_import_entries": p,
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
MandateImportEntry *MandateImportEntry `json:"mandate_import_entries"`
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

if result.MandateImportEntry == nil {
    return nil, errors.New("missing result")
  }

  return result.MandateImportEntry, nil
}


// MandateImportEntryListParams parameters
type MandateImportEntryListParams struct {
      After string `url:",omitempty" json:"after,omitempty"`
      Before string `url:",omitempty" json:"before,omitempty"`
      Limit int `url:",omitempty" json:"limit,omitempty"`
      MandateImport string `url:",omitempty" json:"mandate_import,omitempty"`
      }

// MandateImportEntryListResult response including pagination metadata
type MandateImportEntryListResult struct {
  MandateImportEntries []MandateImportEntry `json:"mandate_import_entries"`
  Meta struct {
      Cursors struct {
      After string `url:",omitempty" json:"after,omitempty"`
      Before string `url:",omitempty" json:"before,omitempty"`
      } `url:",omitempty" json:"cursors,omitempty"`
      Limit int `url:",omitempty" json:"limit,omitempty"`
      } `json:"meta"`
}


// List
// For an existing mandate import, this endpoint lists all of the entries
// attached.
// 
// After a mandate import has been submitted, you can use this endpoint to
// associate records
// in your system (using the `record_identifier` that you provided when creating
// the
// mandate import).
// 
func (s *MandateImportEntryService) List(ctx context.Context, p MandateImportEntryListParams) (*MandateImportEntryListResult,error) {
  uri, err := url.Parse(fmt.Sprintf(s.endpoint + "/mandate_import_entries",))
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
*MandateImportEntryListResult
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

if result.MandateImportEntryListResult == nil {
    return nil, errors.New("missing result")
  }

  return result.MandateImportEntryListResult, nil
}

