package gocardless

import (
  "bytes"
  "errors"
  "fmt"
)

const (

  // Live environment
  LiveEndpoint = "https://api.gocardless.com"

  // Sandbox environment
  SandboxEndpoint = "https://api-sandbox.gocardless.com"

)

type Service struct {

  BankDetailsLookups *BankDetailsLookupService
  Creditors *CreditorService
  CreditorBankAccounts *CreditorBankAccountService
  Customers *CustomerService
  CustomerBankAccounts *CustomerBankAccountService
  Events *EventService
  Mandates *MandateService
  MandatePdfs *MandatePdfService
  Payments *PaymentService
  Payouts *PayoutService
  RedirectFlows *RedirectFlowService
  Refunds *RefundService
  Subscriptions *SubscriptionService
}

func New(token string, opts ...Option) (*Service, error) {
  if token == "" {
    return nil, errors.New("token required")
  }

  o := &options{
    endpoint: LiveEndpoint,
  }
  for _, opt := range opts {
    if err := opt(o); err != nil {
      return nil, err
    }
  }

  s := &Service{}

  s.BankDetailsLookups = &BankDetailsLookupService{
    token: token,
    endpoint: o.endpoint,
    client: o.client,
  }
  s.Creditors = &CreditorService{
    token: token,
    endpoint: o.endpoint,
    client: o.client,
  }
  s.CreditorBankAccounts = &CreditorBankAccountService{
    token: token,
    endpoint: o.endpoint,
    client: o.client,
  }
  s.Customers = &CustomerService{
    token: token,
    endpoint: o.endpoint,
    client: o.client,
  }
  s.CustomerBankAccounts = &CustomerBankAccountService{
    token: token,
    endpoint: o.endpoint,
    client: o.client,
  }
  s.Events = &EventService{
    token: token,
    endpoint: o.endpoint,
    client: o.client,
  }
  s.Mandates = &MandateService{
    token: token,
    endpoint: o.endpoint,
    client: o.client,
  }
  s.MandatePdfs = &MandatePdfService{
    token: token,
    endpoint: o.endpoint,
    client: o.client,
  }
  s.Payments = &PaymentService{
    token: token,
    endpoint: o.endpoint,
    client: o.client,
  }
  s.Payouts = &PayoutService{
    token: token,
    endpoint: o.endpoint,
    client: o.client,
  }
  s.RedirectFlows = &RedirectFlowService{
    token: token,
    endpoint: o.endpoint,
    client: o.client,
  }
  s.Refunds = &RefundService{
    token: token,
    endpoint: o.endpoint,
    client: o.client,
  }
  s.Subscriptions = &SubscriptionService{
    token: token,
    endpoint: o.endpoint,
    client: o.client,
  }
  return s, nil
}

type APIError struct {
  Message string `json:"message"`
  DocumentationUrl string `json:"documentation_url"`
  Type string `json:"type"`
  RequestID string `json:"request_id"`
  Errors []ValidationError `json:"errors"`
  Code int `json:"code"`
}

func (err *APIError) Error() string {
  if len(err.Errors) == 0 {
    return err.Message
  }
  var msg bytes.Buffer
  fmt.Fprintf(&msg, "%s:", err.Message)
  for _, err := range err.Errors {
    fmt.Fprintf(&msg, "\n * %s: %s", err.Field, err.Message)
  }
  return msg.String()
}

type ValidationError struct {
  Message string `json:"message"`
  Field string `json:"field"`
  RequestPointer string `json:"request_pointer"`
}
