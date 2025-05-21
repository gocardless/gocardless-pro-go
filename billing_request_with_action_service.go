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

// BillingRequestWithActionService manages billing_request_with_actions
type BillingRequestWithActionServiceImpl struct {
	config Config
}

type BillingRequestWithActionBankAuthorisationsLinks struct {
	BillingRequest string `url:"billing_request,omitempty" json:"billing_request,omitempty"`
	Institution    string `url:"institution,omitempty" json:"institution,omitempty"`
}

type BillingRequestWithActionBankAuthorisations struct {
	AuthorisationType string                                           `url:"authorisation_type,omitempty" json:"authorisation_type,omitempty"`
	AuthorisedAt      string                                           `url:"authorised_at,omitempty" json:"authorised_at,omitempty"`
	CreatedAt         string                                           `url:"created_at,omitempty" json:"created_at,omitempty"`
	ExpiresAt         string                                           `url:"expires_at,omitempty" json:"expires_at,omitempty"`
	Id                string                                           `url:"id,omitempty" json:"id,omitempty"`
	LastVisitedAt     string                                           `url:"last_visited_at,omitempty" json:"last_visited_at,omitempty"`
	Links             *BillingRequestWithActionBankAuthorisationsLinks `url:"links,omitempty" json:"links,omitempty"`
	QrCodeUrl         string                                           `url:"qr_code_url,omitempty" json:"qr_code_url,omitempty"`
	RedirectUri       string                                           `url:"redirect_uri,omitempty" json:"redirect_uri,omitempty"`
	Url               string                                           `url:"url,omitempty" json:"url,omitempty"`
}

type BillingRequestWithActionBillingRequestsActionsAvailableCurrencies struct {
	Currency string `url:"currency,omitempty" json:"currency,omitempty"`
}

type BillingRequestWithActionBillingRequestsActionsBankAuthorisation struct {
	Adapter           string `url:"adapter,omitempty" json:"adapter,omitempty"`
	AuthorisationType string `url:"authorisation_type,omitempty" json:"authorisation_type,omitempty"`
}

type BillingRequestWithActionBillingRequestsActionsCollectCustomerDetailsIncompleteFields struct {
	Customer              []string `url:"customer,omitempty" json:"customer,omitempty"`
	CustomerBillingDetail []string `url:"customer_billing_detail,omitempty" json:"customer_billing_detail,omitempty"`
}

type BillingRequestWithActionBillingRequestsActionsCollectCustomerDetails struct {
	DefaultCountryCode string                                                                                `url:"default_country_code,omitempty" json:"default_country_code,omitempty"`
	IncompleteFields   *BillingRequestWithActionBillingRequestsActionsCollectCustomerDetailsIncompleteFields `url:"incomplete_fields,omitempty" json:"incomplete_fields,omitempty"`
}

type BillingRequestWithActionBillingRequestsActions struct {
	AvailableCurrencies    *[]string                                                             `url:"available_currencies,omitempty" json:"available_currencies,omitempty"`
	BankAuthorisation      *BillingRequestWithActionBillingRequestsActionsBankAuthorisation      `url:"bank_authorisation,omitempty" json:"bank_authorisation,omitempty"`
	CollectCustomerDetails *BillingRequestWithActionBillingRequestsActionsCollectCustomerDetails `url:"collect_customer_details,omitempty" json:"collect_customer_details,omitempty"`
	CompletesActions       []string                                                              `url:"completes_actions,omitempty" json:"completes_actions,omitempty"`
	InstitutionGuessStatus string                                                                `url:"institution_guess_status,omitempty" json:"institution_guess_status,omitempty"`
	Required               bool                                                                  `url:"required,omitempty" json:"required,omitempty"`
	RequiresActions        []string                                                              `url:"requires_actions,omitempty" json:"requires_actions,omitempty"`
	Status                 string                                                                `url:"status,omitempty" json:"status,omitempty"`
	Type                   string                                                                `url:"type,omitempty" json:"type,omitempty"`
}

type BillingRequestWithActionBillingRequestsInstalmentScheduleRequestInstalmentsWithDates struct {
	Amount      int    `url:"amount,omitempty" json:"amount,omitempty"`
	ChargeDate  string `url:"charge_date,omitempty" json:"charge_date,omitempty"`
	Description string `url:"description,omitempty" json:"description,omitempty"`
}

type BillingRequestWithActionBillingRequestsInstalmentScheduleRequestInstalmentsWithSchedule struct {
	Amounts      []int  `url:"amounts,omitempty" json:"amounts,omitempty"`
	Interval     int    `url:"interval,omitempty" json:"interval,omitempty"`
	IntervalUnit string `url:"interval_unit,omitempty" json:"interval_unit,omitempty"`
	StartDate    string `url:"start_date,omitempty" json:"start_date,omitempty"`
}

type BillingRequestWithActionBillingRequestsInstalmentScheduleRequestLinks struct {
	InstalmentSchedule string `url:"instalment_schedule,omitempty" json:"instalment_schedule,omitempty"`
}

type BillingRequestWithActionBillingRequestsInstalmentScheduleRequest struct {
	AppFee                  int                                                                                      `url:"app_fee,omitempty" json:"app_fee,omitempty"`
	Currency                string                                                                                   `url:"currency,omitempty" json:"currency,omitempty"`
	InstalmentsWithDates    []BillingRequestWithActionBillingRequestsInstalmentScheduleRequestInstalmentsWithDates   `url:"instalments_with_dates,omitempty" json:"instalments_with_dates,omitempty"`
	InstalmentsWithSchedule *BillingRequestWithActionBillingRequestsInstalmentScheduleRequestInstalmentsWithSchedule `url:"instalments_with_schedule,omitempty" json:"instalments_with_schedule,omitempty"`
	Links                   *BillingRequestWithActionBillingRequestsInstalmentScheduleRequestLinks                   `url:"links,omitempty" json:"links,omitempty"`
	Metadata                map[string]interface{}                                                                   `url:"metadata,omitempty" json:"metadata,omitempty"`
	Name                    string                                                                                   `url:"name,omitempty" json:"name,omitempty"`
	PaymentReference        string                                                                                   `url:"payment_reference,omitempty" json:"payment_reference,omitempty"`
	RetryIfPossible         bool                                                                                     `url:"retry_if_possible,omitempty" json:"retry_if_possible,omitempty"`
	TotalAmount             int                                                                                      `url:"total_amount,omitempty" json:"total_amount,omitempty"`
}

type BillingRequestWithActionBillingRequestsLinks struct {
	BankAuthorisation                           string `url:"bank_authorisation,omitempty" json:"bank_authorisation,omitempty"`
	Creditor                                    string `url:"creditor,omitempty" json:"creditor,omitempty"`
	Customer                                    string `url:"customer,omitempty" json:"customer,omitempty"`
	CustomerBankAccount                         string `url:"customer_bank_account,omitempty" json:"customer_bank_account,omitempty"`
	CustomerBillingDetail                       string `url:"customer_billing_detail,omitempty" json:"customer_billing_detail,omitempty"`
	InstalmentScheduleRequest                   string `url:"instalment_schedule_request,omitempty" json:"instalment_schedule_request,omitempty"`
	InstalmentScheduleRequestInstalmentSchedule string `url:"instalment_schedule_request_instalment_schedule,omitempty" json:"instalment_schedule_request_instalment_schedule,omitempty"`
	MandateRequest                              string `url:"mandate_request,omitempty" json:"mandate_request,omitempty"`
	MandateRequestMandate                       string `url:"mandate_request_mandate,omitempty" json:"mandate_request_mandate,omitempty"`
	Organisation                                string `url:"organisation,omitempty" json:"organisation,omitempty"`
	PaymentProvider                             string `url:"payment_provider,omitempty" json:"payment_provider,omitempty"`
	PaymentRequest                              string `url:"payment_request,omitempty" json:"payment_request,omitempty"`
	PaymentRequestPayment                       string `url:"payment_request_payment,omitempty" json:"payment_request_payment,omitempty"`
	SubscriptionRequest                         string `url:"subscription_request,omitempty" json:"subscription_request,omitempty"`
	SubscriptionRequestSubscription             string `url:"subscription_request_subscription,omitempty" json:"subscription_request_subscription,omitempty"`
}

type BillingRequestWithActionBillingRequestsMandateRequestConstraintsPeriodicLimits struct {
	Alignment      string `url:"alignment,omitempty" json:"alignment,omitempty"`
	MaxPayments    int    `url:"max_payments,omitempty" json:"max_payments,omitempty"`
	MaxTotalAmount int    `url:"max_total_amount,omitempty" json:"max_total_amount,omitempty"`
	Period         string `url:"period,omitempty" json:"period,omitempty"`
}

type BillingRequestWithActionBillingRequestsMandateRequestConstraints struct {
	EndDate             string                                                                           `url:"end_date,omitempty" json:"end_date,omitempty"`
	MaxAmountPerPayment int                                                                              `url:"max_amount_per_payment,omitempty" json:"max_amount_per_payment,omitempty"`
	PaymentMethod       string                                                                           `url:"payment_method,omitempty" json:"payment_method,omitempty"`
	PeriodicLimits      []BillingRequestWithActionBillingRequestsMandateRequestConstraintsPeriodicLimits `url:"periodic_limits,omitempty" json:"periodic_limits,omitempty"`
	StartDate           string                                                                           `url:"start_date,omitempty" json:"start_date,omitempty"`
}

type BillingRequestWithActionBillingRequestsMandateRequestLinks struct {
	Mandate string `url:"mandate,omitempty" json:"mandate,omitempty"`
}

type BillingRequestWithActionBillingRequestsMandateRequest struct {
	AuthorisationSource         string                                                            `url:"authorisation_source,omitempty" json:"authorisation_source,omitempty"`
	ConsentType                 string                                                            `url:"consent_type,omitempty" json:"consent_type,omitempty"`
	Constraints                 *BillingRequestWithActionBillingRequestsMandateRequestConstraints `url:"constraints,omitempty" json:"constraints,omitempty"`
	Currency                    string                                                            `url:"currency,omitempty" json:"currency,omitempty"`
	Description                 string                                                            `url:"description,omitempty" json:"description,omitempty"`
	Links                       *BillingRequestWithActionBillingRequestsMandateRequestLinks       `url:"links,omitempty" json:"links,omitempty"`
	Metadata                    map[string]interface{}                                            `url:"metadata,omitempty" json:"metadata,omitempty"`
	PayerRequestedDualSignature bool                                                              `url:"payer_requested_dual_signature,omitempty" json:"payer_requested_dual_signature,omitempty"`
	Scheme                      string                                                            `url:"scheme,omitempty" json:"scheme,omitempty"`
	Sweeping                    bool                                                              `url:"sweeping,omitempty" json:"sweeping,omitempty"`
	Verify                      string                                                            `url:"verify,omitempty" json:"verify,omitempty"`
}

type BillingRequestWithActionBillingRequestsPaymentRequestLinks struct {
	Payment string `url:"payment,omitempty" json:"payment,omitempty"`
}

type BillingRequestWithActionBillingRequestsPaymentRequest struct {
	Amount          int                                                         `url:"amount,omitempty" json:"amount,omitempty"`
	AppFee          int                                                         `url:"app_fee,omitempty" json:"app_fee,omitempty"`
	Currency        string                                                      `url:"currency,omitempty" json:"currency,omitempty"`
	Description     string                                                      `url:"description,omitempty" json:"description,omitempty"`
	FundsSettlement string                                                      `url:"funds_settlement,omitempty" json:"funds_settlement,omitempty"`
	Links           *BillingRequestWithActionBillingRequestsPaymentRequestLinks `url:"links,omitempty" json:"links,omitempty"`
	Metadata        map[string]interface{}                                      `url:"metadata,omitempty" json:"metadata,omitempty"`
	Reference       string                                                      `url:"reference,omitempty" json:"reference,omitempty"`
	Scheme          string                                                      `url:"scheme,omitempty" json:"scheme,omitempty"`
}

type BillingRequestWithActionBillingRequestsResourcesCustomer struct {
	CompanyName string                 `url:"company_name,omitempty" json:"company_name,omitempty"`
	CreatedAt   string                 `url:"created_at,omitempty" json:"created_at,omitempty"`
	Email       string                 `url:"email,omitempty" json:"email,omitempty"`
	FamilyName  string                 `url:"family_name,omitempty" json:"family_name,omitempty"`
	GivenName   string                 `url:"given_name,omitempty" json:"given_name,omitempty"`
	Id          string                 `url:"id,omitempty" json:"id,omitempty"`
	Language    string                 `url:"language,omitempty" json:"language,omitempty"`
	Metadata    map[string]interface{} `url:"metadata,omitempty" json:"metadata,omitempty"`
	PhoneNumber string                 `url:"phone_number,omitempty" json:"phone_number,omitempty"`
}

type BillingRequestWithActionBillingRequestsResourcesCustomerBankAccountLinks struct {
	Customer string `url:"customer,omitempty" json:"customer,omitempty"`
}

type BillingRequestWithActionBillingRequestsResourcesCustomerBankAccount struct {
	AccountHolderName   string                                                                    `url:"account_holder_name,omitempty" json:"account_holder_name,omitempty"`
	AccountNumberEnding string                                                                    `url:"account_number_ending,omitempty" json:"account_number_ending,omitempty"`
	AccountType         string                                                                    `url:"account_type,omitempty" json:"account_type,omitempty"`
	BankAccountToken    string                                                                    `url:"bank_account_token,omitempty" json:"bank_account_token,omitempty"`
	BankName            string                                                                    `url:"bank_name,omitempty" json:"bank_name,omitempty"`
	CountryCode         string                                                                    `url:"country_code,omitempty" json:"country_code,omitempty"`
	CreatedAt           string                                                                    `url:"created_at,omitempty" json:"created_at,omitempty"`
	Currency            string                                                                    `url:"currency,omitempty" json:"currency,omitempty"`
	Enabled             bool                                                                      `url:"enabled,omitempty" json:"enabled,omitempty"`
	Id                  string                                                                    `url:"id,omitempty" json:"id,omitempty"`
	Links               *BillingRequestWithActionBillingRequestsResourcesCustomerBankAccountLinks `url:"links,omitempty" json:"links,omitempty"`
	Metadata            map[string]interface{}                                                    `url:"metadata,omitempty" json:"metadata,omitempty"`
}

type BillingRequestWithActionBillingRequestsResourcesCustomerBillingDetail struct {
	AddressLine1          string   `url:"address_line1,omitempty" json:"address_line1,omitempty"`
	AddressLine2          string   `url:"address_line2,omitempty" json:"address_line2,omitempty"`
	AddressLine3          string   `url:"address_line3,omitempty" json:"address_line3,omitempty"`
	City                  string   `url:"city,omitempty" json:"city,omitempty"`
	CountryCode           string   `url:"country_code,omitempty" json:"country_code,omitempty"`
	CreatedAt             string   `url:"created_at,omitempty" json:"created_at,omitempty"`
	DanishIdentityNumber  string   `url:"danish_identity_number,omitempty" json:"danish_identity_number,omitempty"`
	Id                    string   `url:"id,omitempty" json:"id,omitempty"`
	IpAddress             string   `url:"ip_address,omitempty" json:"ip_address,omitempty"`
	PostalCode            string   `url:"postal_code,omitempty" json:"postal_code,omitempty"`
	Region                string   `url:"region,omitempty" json:"region,omitempty"`
	Schemes               []string `url:"schemes,omitempty" json:"schemes,omitempty"`
	SwedishIdentityNumber string   `url:"swedish_identity_number,omitempty" json:"swedish_identity_number,omitempty"`
}

type BillingRequestWithActionBillingRequestsResources struct {
	Customer              *BillingRequestWithActionBillingRequestsResourcesCustomer              `url:"customer,omitempty" json:"customer,omitempty"`
	CustomerBankAccount   *BillingRequestWithActionBillingRequestsResourcesCustomerBankAccount   `url:"customer_bank_account,omitempty" json:"customer_bank_account,omitempty"`
	CustomerBillingDetail *BillingRequestWithActionBillingRequestsResourcesCustomerBillingDetail `url:"customer_billing_detail,omitempty" json:"customer_billing_detail,omitempty"`
}

type BillingRequestWithActionBillingRequestsSubscriptionRequestLinks struct {
	Subscription string `url:"subscription,omitempty" json:"subscription,omitempty"`
}

type BillingRequestWithActionBillingRequestsSubscriptionRequest struct {
	Amount           int                                                              `url:"amount,omitempty" json:"amount,omitempty"`
	AppFee           int                                                              `url:"app_fee,omitempty" json:"app_fee,omitempty"`
	Count            int                                                              `url:"count,omitempty" json:"count,omitempty"`
	Currency         string                                                           `url:"currency,omitempty" json:"currency,omitempty"`
	DayOfMonth       int                                                              `url:"day_of_month,omitempty" json:"day_of_month,omitempty"`
	Interval         int                                                              `url:"interval,omitempty" json:"interval,omitempty"`
	IntervalUnit     string                                                           `url:"interval_unit,omitempty" json:"interval_unit,omitempty"`
	Links            *BillingRequestWithActionBillingRequestsSubscriptionRequestLinks `url:"links,omitempty" json:"links,omitempty"`
	Metadata         map[string]interface{}                                           `url:"metadata,omitempty" json:"metadata,omitempty"`
	Month            string                                                           `url:"month,omitempty" json:"month,omitempty"`
	Name             string                                                           `url:"name,omitempty" json:"name,omitempty"`
	PaymentReference string                                                           `url:"payment_reference,omitempty" json:"payment_reference,omitempty"`
	RetryIfPossible  bool                                                             `url:"retry_if_possible,omitempty" json:"retry_if_possible,omitempty"`
	StartDate        string                                                           `url:"start_date,omitempty" json:"start_date,omitempty"`
}

type BillingRequestWithActionBillingRequests struct {
	Actions                   []BillingRequestWithActionBillingRequestsActions                  `url:"actions,omitempty" json:"actions,omitempty"`
	CreatedAt                 string                                                            `url:"created_at,omitempty" json:"created_at,omitempty"`
	FallbackEnabled           bool                                                              `url:"fallback_enabled,omitempty" json:"fallback_enabled,omitempty"`
	FallbackOccurred          bool                                                              `url:"fallback_occurred,omitempty" json:"fallback_occurred,omitempty"`
	Id                        string                                                            `url:"id,omitempty" json:"id,omitempty"`
	InstalmentScheduleRequest *BillingRequestWithActionBillingRequestsInstalmentScheduleRequest `url:"instalment_schedule_request,omitempty" json:"instalment_schedule_request,omitempty"`
	Links                     *BillingRequestWithActionBillingRequestsLinks                     `url:"links,omitempty" json:"links,omitempty"`
	MandateRequest            *BillingRequestWithActionBillingRequestsMandateRequest            `url:"mandate_request,omitempty" json:"mandate_request,omitempty"`
	Metadata                  map[string]interface{}                                            `url:"metadata,omitempty" json:"metadata,omitempty"`
	PaymentRequest            *BillingRequestWithActionBillingRequestsPaymentRequest            `url:"payment_request,omitempty" json:"payment_request,omitempty"`
	PurposeCode               string                                                            `url:"purpose_code,omitempty" json:"purpose_code,omitempty"`
	Resources                 *BillingRequestWithActionBillingRequestsResources                 `url:"resources,omitempty" json:"resources,omitempty"`
	Status                    string                                                            `url:"status,omitempty" json:"status,omitempty"`
	SubscriptionRequest       *BillingRequestWithActionBillingRequestsSubscriptionRequest       `url:"subscription_request,omitempty" json:"subscription_request,omitempty"`
}

// BillingRequestWithAction model
type BillingRequestWithAction struct {
	BankAuthorisations *BillingRequestWithActionBankAuthorisations `url:"bank_authorisations,omitempty" json:"bank_authorisations,omitempty"`
	BillingRequests    *BillingRequestWithActionBillingRequests    `url:"billing_requests,omitempty" json:"billing_requests,omitempty"`
}

type BillingRequestWithActionService interface {
	CreateWithActions(ctx context.Context, p BillingRequestWithActionCreateWithActionsParams, opts ...RequestOption) (*BillingRequestWithAction, error)
}

type BillingRequestWithActionCreateWithActionsParamsActionsCollectBankAccount struct {
	AccountHolderName   string                 `url:"account_holder_name,omitempty" json:"account_holder_name,omitempty"`
	AccountNumber       string                 `url:"account_number,omitempty" json:"account_number,omitempty"`
	AccountNumberSuffix string                 `url:"account_number_suffix,omitempty" json:"account_number_suffix,omitempty"`
	AccountType         string                 `url:"account_type,omitempty" json:"account_type,omitempty"`
	BankCode            string                 `url:"bank_code,omitempty" json:"bank_code,omitempty"`
	BranchCode          string                 `url:"branch_code,omitempty" json:"branch_code,omitempty"`
	CountryCode         string                 `url:"country_code,omitempty" json:"country_code,omitempty"`
	Currency            string                 `url:"currency,omitempty" json:"currency,omitempty"`
	Iban                string                 `url:"iban,omitempty" json:"iban,omitempty"`
	Metadata            map[string]interface{} `url:"metadata,omitempty" json:"metadata,omitempty"`
	PayId               string                 `url:"pay_id,omitempty" json:"pay_id,omitempty"`
}

type BillingRequestWithActionCreateWithActionsParamsActionsCollectCustomerDetailsCustomer struct {
	CompanyName string                 `url:"company_name,omitempty" json:"company_name,omitempty"`
	Email       string                 `url:"email,omitempty" json:"email,omitempty"`
	FamilyName  string                 `url:"family_name,omitempty" json:"family_name,omitempty"`
	GivenName   string                 `url:"given_name,omitempty" json:"given_name,omitempty"`
	Language    string                 `url:"language,omitempty" json:"language,omitempty"`
	Metadata    map[string]interface{} `url:"metadata,omitempty" json:"metadata,omitempty"`
	PhoneNumber string                 `url:"phone_number,omitempty" json:"phone_number,omitempty"`
}

type BillingRequestWithActionCreateWithActionsParamsActionsCollectCustomerDetailsCustomerBillingDetail struct {
	AddressLine1          string `url:"address_line1,omitempty" json:"address_line1,omitempty"`
	AddressLine2          string `url:"address_line2,omitempty" json:"address_line2,omitempty"`
	AddressLine3          string `url:"address_line3,omitempty" json:"address_line3,omitempty"`
	City                  string `url:"city,omitempty" json:"city,omitempty"`
	CountryCode           string `url:"country_code,omitempty" json:"country_code,omitempty"`
	DanishIdentityNumber  string `url:"danish_identity_number,omitempty" json:"danish_identity_number,omitempty"`
	IpAddress             string `url:"ip_address,omitempty" json:"ip_address,omitempty"`
	PostalCode            string `url:"postal_code,omitempty" json:"postal_code,omitempty"`
	Region                string `url:"region,omitempty" json:"region,omitempty"`
	SwedishIdentityNumber string `url:"swedish_identity_number,omitempty" json:"swedish_identity_number,omitempty"`
}

type BillingRequestWithActionCreateWithActionsParamsActionsCollectCustomerDetails struct {
	Customer              *BillingRequestWithActionCreateWithActionsParamsActionsCollectCustomerDetailsCustomer              `url:"customer,omitempty" json:"customer,omitempty"`
	CustomerBillingDetail *BillingRequestWithActionCreateWithActionsParamsActionsCollectCustomerDetailsCustomerBillingDetail `url:"customer_billing_detail,omitempty" json:"customer_billing_detail,omitempty"`
}

type BillingRequestWithActionCreateWithActionsParamsActionsConfirmPayerDetails struct {
	Metadata                    map[string]interface{} `url:"metadata,omitempty" json:"metadata,omitempty"`
	PayerRequestedDualSignature bool                   `url:"payer_requested_dual_signature,omitempty" json:"payer_requested_dual_signature,omitempty"`
}

type BillingRequestWithActionCreateWithActionsParamsActionsSelectInstitution struct {
	CountryCode string `url:"country_code,omitempty" json:"country_code,omitempty"`
	Institution string `url:"institution,omitempty" json:"institution,omitempty"`
}

type BillingRequestWithActionCreateWithActionsParamsActions struct {
	BankAuthorisationRedirectUri string                                                                        `url:"bank_authorisation_redirect_uri,omitempty" json:"bank_authorisation_redirect_uri,omitempty"`
	CollectBankAccount           *BillingRequestWithActionCreateWithActionsParamsActionsCollectBankAccount     `url:"collect_bank_account,omitempty" json:"collect_bank_account,omitempty"`
	CollectCustomerDetails       *BillingRequestWithActionCreateWithActionsParamsActionsCollectCustomerDetails `url:"collect_customer_details,omitempty" json:"collect_customer_details,omitempty"`
	ConfirmPayerDetails          *BillingRequestWithActionCreateWithActionsParamsActionsConfirmPayerDetails    `url:"confirm_payer_details,omitempty" json:"confirm_payer_details,omitempty"`
	CreateBankAuthorisation      bool                                                                          `url:"create_bank_authorisation,omitempty" json:"create_bank_authorisation,omitempty"`
	SelectInstitution            *BillingRequestWithActionCreateWithActionsParamsActionsSelectInstitution      `url:"select_institution,omitempty" json:"select_institution,omitempty"`
}

type BillingRequestWithActionCreateWithActionsParamsLinks struct {
	Creditor            string `url:"creditor,omitempty" json:"creditor,omitempty"`
	Customer            string `url:"customer,omitempty" json:"customer,omitempty"`
	CustomerBankAccount string `url:"customer_bank_account,omitempty" json:"customer_bank_account,omitempty"`
}

type BillingRequestWithActionCreateWithActionsParamsMandateRequestConstraintsPeriodicLimits struct {
	Alignment      string `url:"alignment,omitempty" json:"alignment,omitempty"`
	MaxPayments    int    `url:"max_payments,omitempty" json:"max_payments,omitempty"`
	MaxTotalAmount int    `url:"max_total_amount,omitempty" json:"max_total_amount,omitempty"`
	Period         string `url:"period,omitempty" json:"period,omitempty"`
}

type BillingRequestWithActionCreateWithActionsParamsMandateRequestConstraints struct {
	EndDate             string                                                                                   `url:"end_date,omitempty" json:"end_date,omitempty"`
	MaxAmountPerPayment int                                                                                      `url:"max_amount_per_payment,omitempty" json:"max_amount_per_payment,omitempty"`
	PaymentMethod       string                                                                                   `url:"payment_method,omitempty" json:"payment_method,omitempty"`
	PeriodicLimits      []BillingRequestWithActionCreateWithActionsParamsMandateRequestConstraintsPeriodicLimits `url:"periodic_limits,omitempty" json:"periodic_limits,omitempty"`
	StartDate           string                                                                                   `url:"start_date,omitempty" json:"start_date,omitempty"`
}

type BillingRequestWithActionCreateWithActionsParamsMandateRequest struct {
	AuthorisationSource string                                                                    `url:"authorisation_source,omitempty" json:"authorisation_source,omitempty"`
	Constraints         *BillingRequestWithActionCreateWithActionsParamsMandateRequestConstraints `url:"constraints,omitempty" json:"constraints,omitempty"`
	Currency            string                                                                    `url:"currency,omitempty" json:"currency,omitempty"`
	Description         string                                                                    `url:"description,omitempty" json:"description,omitempty"`
	Metadata            map[string]interface{}                                                    `url:"metadata,omitempty" json:"metadata,omitempty"`
	Reference           string                                                                    `url:"reference,omitempty" json:"reference,omitempty"`
	Scheme              string                                                                    `url:"scheme,omitempty" json:"scheme,omitempty"`
	Sweeping            bool                                                                      `url:"sweeping,omitempty" json:"sweeping,omitempty"`
	Verify              string                                                                    `url:"verify,omitempty" json:"verify,omitempty"`
}

type BillingRequestWithActionCreateWithActionsParamsPaymentRequest struct {
	Amount          int                    `url:"amount,omitempty" json:"amount,omitempty"`
	AppFee          int                    `url:"app_fee,omitempty" json:"app_fee,omitempty"`
	Currency        string                 `url:"currency,omitempty" json:"currency,omitempty"`
	Description     string                 `url:"description,omitempty" json:"description,omitempty"`
	FundsSettlement string                 `url:"funds_settlement,omitempty" json:"funds_settlement,omitempty"`
	Metadata        map[string]interface{} `url:"metadata,omitempty" json:"metadata,omitempty"`
	Reference       string                 `url:"reference,omitempty" json:"reference,omitempty"`
	RetryIfPossible bool                   `url:"retry_if_possible,omitempty" json:"retry_if_possible,omitempty"`
	Scheme          string                 `url:"scheme,omitempty" json:"scheme,omitempty"`
}

// BillingRequestWithActionCreateWithActionsParams parameters
type BillingRequestWithActionCreateWithActionsParams struct {
	Actions         *BillingRequestWithActionCreateWithActionsParamsActions        `url:"actions,omitempty" json:"actions,omitempty"`
	FallbackEnabled bool                                                           `url:"fallback_enabled,omitempty" json:"fallback_enabled,omitempty"`
	Links           *BillingRequestWithActionCreateWithActionsParamsLinks          `url:"links,omitempty" json:"links,omitempty"`
	MandateRequest  *BillingRequestWithActionCreateWithActionsParamsMandateRequest `url:"mandate_request,omitempty" json:"mandate_request,omitempty"`
	Metadata        map[string]interface{}                                         `url:"metadata,omitempty" json:"metadata,omitempty"`
	PaymentRequest  *BillingRequestWithActionCreateWithActionsParamsPaymentRequest `url:"payment_request,omitempty" json:"payment_request,omitempty"`
	PurposeCode     string                                                         `url:"purpose_code,omitempty" json:"purpose_code,omitempty"`
}

// CreateWithActions
// Creates a billing request and completes any specified actions in a single
// request.
// This endpoint allows you to create a billing request and immediately complete
// actions
// such as collecting customer details, bank account details, or other required
// actions.
func (s *BillingRequestWithActionServiceImpl) CreateWithActions(ctx context.Context, p BillingRequestWithActionCreateWithActionsParams, opts ...RequestOption) (*BillingRequestWithAction, error) {
	uri, err := url.Parse(fmt.Sprintf(s.config.Endpoint() + "/billing_requests/create_with_actions"))
	if err != nil {
		return nil, err
	}

	o := &requestOptions{
		retries: 3,
	}
	for _, opt := range opts {
		err := opt(o)
		if err != nil {
			return nil, err
		}
	}
	if o.idempotencyKey == "" {
		o.idempotencyKey = NewIdempotencyKey()
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
	req = req.WithContext(ctx)
	req.Header.Set("Authorization", "Bearer "+s.config.Token())
	req.Header.Set("GoCardless-Version", "2015-07-06")
	req.Header.Set("GoCardless-Client-Library", "gocardless-pro-go")
	req.Header.Set("GoCardless-Client-Version", "5.0.0")
	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Idempotency-Key", o.idempotencyKey)

	for key, value := range o.headers {
		req.Header.Set(key, value)
	}

	client := s.config.Client()
	if client == nil {
		client = http.DefaultClient
	}

	var result struct {
		Err                      *APIError                 `json:"error"`
		BillingRequestWithAction *BillingRequestWithAction `json:"billing_request_with_actions"`
	}

	err = try(o.retries, func() error {
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

	if result.BillingRequestWithAction == nil {
		return nil, errors.New("missing result")
	}

	return result.BillingRequestWithAction, nil
}
