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
	BankAuthorisations      *BankAuthorisationService
	BankDetailsLookups      *BankDetailsLookupService
	BillingRequests         *BillingRequestService
	BillingRequestFlows     *BillingRequestFlowService
	BillingRequestTemplates *BillingRequestTemplateService
	Blocks                  *BlockService
	Creditors               *CreditorService
	CreditorBankAccounts    *CreditorBankAccountService
	CurrencyExchangeRates   *CurrencyExchangeRateService
	Customers               *CustomerService
	CustomerBankAccounts    *CustomerBankAccountService
	CustomerNotifications   *CustomerNotificationService
	Events                  *EventService
	InstalmentSchedules     *InstalmentScheduleService
	Institutions            *InstitutionService
	Mandates                *MandateService
	MandateImports          *MandateImportService
	MandateImportEntries    *MandateImportEntryService
	MandatePdfs             *MandatePdfService
	PayerAuthorisations     *PayerAuthorisationService
	Payments                *PaymentService
	Payouts                 *PayoutService
	PayoutItems             *PayoutItemService
	RedirectFlows           *RedirectFlowService
	Refunds                 *RefundService
	ScenarioSimulators      *ScenarioSimulatorService
	Subscriptions           *SubscriptionService
	TaxRates                *TaxRateService
	Webhooks                *WebhookService
}

func init() {
	initUserAgent()
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

	s.BankAuthorisations = &BankAuthorisationService{
		token:    token,
		endpoint: o.endpoint,
		client:   o.client,
	}
	s.BankDetailsLookups = &BankDetailsLookupService{
		token:    token,
		endpoint: o.endpoint,
		client:   o.client,
	}
	s.BillingRequests = &BillingRequestService{
		token:    token,
		endpoint: o.endpoint,
		client:   o.client,
	}
	s.BillingRequestFlows = &BillingRequestFlowService{
		token:    token,
		endpoint: o.endpoint,
		client:   o.client,
	}
	s.BillingRequestTemplates = &BillingRequestTemplateService{
		token:    token,
		endpoint: o.endpoint,
		client:   o.client,
	}
	s.Blocks = &BlockService{
		token:    token,
		endpoint: o.endpoint,
		client:   o.client,
	}
	s.Creditors = &CreditorService{
		token:    token,
		endpoint: o.endpoint,
		client:   o.client,
	}
	s.CreditorBankAccounts = &CreditorBankAccountService{
		token:    token,
		endpoint: o.endpoint,
		client:   o.client,
	}
	s.CurrencyExchangeRates = &CurrencyExchangeRateService{
		token:    token,
		endpoint: o.endpoint,
		client:   o.client,
	}
	s.Customers = &CustomerService{
		token:    token,
		endpoint: o.endpoint,
		client:   o.client,
	}
	s.CustomerBankAccounts = &CustomerBankAccountService{
		token:    token,
		endpoint: o.endpoint,
		client:   o.client,
	}
	s.CustomerNotifications = &CustomerNotificationService{
		token:    token,
		endpoint: o.endpoint,
		client:   o.client,
	}
	s.Events = &EventService{
		token:    token,
		endpoint: o.endpoint,
		client:   o.client,
	}
	s.InstalmentSchedules = &InstalmentScheduleService{
		token:    token,
		endpoint: o.endpoint,
		client:   o.client,
	}
	s.Institutions = &InstitutionService{
		token:    token,
		endpoint: o.endpoint,
		client:   o.client,
	}
	s.Mandates = &MandateService{
		token:    token,
		endpoint: o.endpoint,
		client:   o.client,
	}
	s.MandateImports = &MandateImportService{
		token:    token,
		endpoint: o.endpoint,
		client:   o.client,
	}
	s.MandateImportEntries = &MandateImportEntryService{
		token:    token,
		endpoint: o.endpoint,
		client:   o.client,
	}
	s.MandatePdfs = &MandatePdfService{
		token:    token,
		endpoint: o.endpoint,
		client:   o.client,
	}
	s.PayerAuthorisations = &PayerAuthorisationService{
		token:    token,
		endpoint: o.endpoint,
		client:   o.client,
	}
	s.Payments = &PaymentService{
		token:    token,
		endpoint: o.endpoint,
		client:   o.client,
	}
	s.Payouts = &PayoutService{
		token:    token,
		endpoint: o.endpoint,
		client:   o.client,
	}
	s.PayoutItems = &PayoutItemService{
		token:    token,
		endpoint: o.endpoint,
		client:   o.client,
	}
	s.RedirectFlows = &RedirectFlowService{
		token:    token,
		endpoint: o.endpoint,
		client:   o.client,
	}
	s.Refunds = &RefundService{
		token:    token,
		endpoint: o.endpoint,
		client:   o.client,
	}
	s.ScenarioSimulators = &ScenarioSimulatorService{
		token:    token,
		endpoint: o.endpoint,
		client:   o.client,
	}
	s.Subscriptions = &SubscriptionService{
		token:    token,
		endpoint: o.endpoint,
		client:   o.client,
	}
	s.TaxRates = &TaxRateService{
		token:    token,
		endpoint: o.endpoint,
		client:   o.client,
	}
	s.Webhooks = &WebhookService{
		token:    token,
		endpoint: o.endpoint,
		client:   o.client,
	}
	return s, nil
}

type APIError struct {
	Message          string            `json:"message"`
	DocumentationUrl string            `json:"documentation_url"`
	Type             string            `json:"type"`
	RequestID        string            `json:"request_id"`
	Errors           []ValidationError `json:"errors"`
	Code             int               `json:"code"`
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
	Message        string `json:"message"`
	Field          string `json:"field"`
	RequestPointer string `json:"request_pointer"`
}
