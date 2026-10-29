<!-- This file is generated, please add to it using `knope document-change` in the client-library-templates repo -->
# Changelog

## 6.23.0 (2026-10-29)

### Features

#### Add account_information_request to billing_request

Add `account_information_request` field on billing requests, in both create and
response, and the `links.account_information_request` ID.

## 6.22.1 (2026-10-29)

### Fixes

- Update documentation links for new docs site

## 6.22.0 (2026-10-28)

### Features

#### Use Billing Requests instead of Redirect Flows and Payer Authorisations

Redirect Flows and Payer Authorisations are legacy APIs and cannot be used by new integrators. If you use `RedirectFlowService` and `PayerAuthorisationService`, migrate to [Billing Requests](https://developer.gocardless.com/api-reference/#billing-requests), which cover the same flows (and more) through a single API. The [integration planner](https://docs.gocardless.com/docs/collect-payments/integration-planner) is a good place to start.

The legacy methods are now marked as deprecated with `// Deprecated:` doc comments. Behaviour is unchanged. `go doc`, gopls and staticcheck will flag these methods.

## 6.21.0 (2026-10-27)

### Features

#### Add `include` parameter to payments

Includes linked resources in the response. Use `customer` to embed customer details, or `mandate` to embed mandate details.

### Fixes

#### Correct `day_of_month` documentation for subscriptions and plans

For subscriptions and plans, values `1`-`31` or `-1` are accepted, not just `1`-`28`.

`1`-`28` applies to payout intervals, creditors and mandate caps.

## 6.20.0 (2026-10-26)

### Features

#### Time out requests after 60 seconds by default

Previously, unless you passed your own client with `WithClient`, requests were made with `http.DefaultClient`, which has no timeout, so a stalled connection could hang forever. The client now uses an `http.Client` with a 60 second timeout (`gocardless.DefaultTimeout`) covering the whole request, including reading the response body.

Clients passed with `WithClient` are used as they are, so set `Timeout` on them to choose your own, or `0` for no timeout:

```go
config, err := gocardless.NewConfig(token, gocardless.WithClient(&http.Client{Timeout: 10 * time.Second}))
```

## 6.19.0 (2026-10-12)

### Features

#### Add `include` parameter to mandates

Includes linked resources in the response. Use `customer` to embed customer details, or `bank_account` to embed bank account details.

## 6.18.2 (2026-10-07)

### Fixes

- Remove unused definitions from schema

## 6.18.1 (2026-10-06)

### Fixes

#### Fix Institutions include_disabled request param to a string enum

This field is only used in request parameters on GET requests. Parameters on
GET requests are querystrings and can only ever be strings.

Therefore any attempt was rejected with "One of your parameters was incorrectly
typed" - "\"true\" is not a boolean.".

This string enum "true" or "false" is our general pattern for parameters like this.

## 6.18.0 (2026-10-06)

### Features

#### Add previously undocumented fields to billing_request

Add `payment_request.retry_if_possible` - On failure, automatically retry payments using intelligent retries.
Add `actions.available_country_codes` - list of currencies the current mandate supports

## 6.17.0 (2026-10-06)

### Features

- Updates the public institutions API to allow merchants to provide the include_disabled query param and see institution status per feature

## 6.16.1 (2026-10-06)

### Fixes

#### Document `required` fields on several API resource responses

Several resource response schemas were missing `required` field lists,
even though the fields are always present in the actual response.
Added accurate `required` arrays (and fixed a few hardcoded docs
examples that predated them) for: `payment`, `bank_details_lookup`,
`billing_request_template`, `currency_exchange_rate`, `mandate_import`,
`mandate_import_entry`, `negative_balance_limit`, `payer_authorisation`,
`payout`, `payout_item`, `refund`, `scenario_simulator`,
`scheme_identifier`, `tax_rate`, `verification_detail` and `logo`.

This is a documentation-only change; API behaviour is unchanged.

## 6.16.0 (2026-10-06)

### Features

#### Reject URL parameters that could change which endpoint is addressed

A URL parameter is a single path segment — a resource identity — but the escaping
applied to one varied by language, and in Go, Node, PHP and .NET there was none at
all. A value carrying path syntax could move a request to an endpoint the caller
never asked for: `find("../mandates")` reached the mandates collection, and
`find("?limit=500")` injected a query parameter.

Escaping alone cannot fix this, because `.` and `..` are dot segments that a path
resolver strips whether or not they are encoded, and an empty value addresses the
collection rather than one resource. Values that could change which endpoint is
addressed are therefore rejected rather than escaped: `/`, `?`, `#`, control
characters, `.`, `..` and the empty string now raise an error instead of producing a
request that quietly 404s. Everything else is escaped as before.

No valid GoCardless resource identity contains any of these characters, so correct
code is unaffected. Ruby and Java previously encoded `/` as `%2F` and sent the
request; they now raise.

## 6.15.0 (2026-10-05)

### Features

- Fixed postal_code and country_code fields on Scheme Identifiers to correctly allow null values.

## 6.14.1 (2026-10-02)

### Fixes

- Verify signature before parsing webhook

## 6.14.0 (2026-10-02)

### Features

- Fixed response_body_truncated, response_headers_content_truncated, and response_headers_count_truncated fields on Webhooks to correctly allow null values.

## 6.13.0 (2026-10-02)

### Features

- Fixed the email field in merchant_contact_details on Billing Request Flows to correctly allow null values.

### Fixes

- Fix nullability of various fields, mainly meta objects in list responses

## 6.12.1 (2026-10-02)

### Fixes

- Fixed the url field on Bank Authorisations to correctly allow null values.

## 6.12.0 (2026-10-01)

### Features

- Made account_number_ending nullable across all resources that reference this shared definition

## 6.11.0 (2026-10-01)

### Features

- Added `language` and `phone_number` properties to prefilled_customer in Billing Request Flows.

## 6.10.0 (2026-10-01)

### Features

#### Fixes to mandate property definitions

* Add missing `consent_parameters` properties: `id`, `scheme`, `currency`, `fixed_amount_per_payment`
* Make `max_amount_per_period`, `max_payments_per_period` and `end_date` nullable
* Add missing `period_alignment` to `period_alignment` enum

### Fixes

#### Fix `GET /institutions` documentation incorrectly listing `status` and `autocompletes_collect_bank_account` fields

These fields are only returned by `GET /billing_requests/{identity}/institutions`, not by the top-level `GET /institutions` endpoint. This is a documentation-only fix; API behaviour is unchanged.

## 6.9.6 (2026-10-01)

### Fixes

#### Correct supported type list for /customer_notifications/{id}/actions/handle

Only `payment_created`, `mandate_created` and `subscription_created` are supported for now, but the enum implied it was more event types.

The enum is unchanged to avoid breaking consumers dependent on the ordering.

## 6.9.5 (2026-09-30)

### Fixes

- Add missing bank_name property to payer_authorisation bank_account schema

## 6.9.4 (2026-09-30)

### Fixes

- Fix nullable metadata reference in billing_request list response schema

## 6.9.3 (2026-09-30)

### Fixes

#### Refine types for billing_request amount fields

The API and client libraries accept string or integer for amount-type fields.
We only emit those same fields as integer.
However, the schema incorrectly specified string or integer for the response as well as the request.

## 6.9.2 (2026-09-30)

### Fixes

- Prune unused definitions from the schema

## 6.9.1 (2026-09-29)

### Fixes

- Add missing institution properties to billing_request schema

## 6.9.0 (2026-09-29)

### Features

#### Add `payer_name_verification_result` to bank details lookups

Bank details lookups now return a `payer_name_verification_result` field when an `account_holder_name` is supplied and a payer name verification check is performed. It can be `full`, `close`, `cannot_perform_verification` or `null`

## 6.8.0 (2026-09-25)

### Features

#### Add `interval` param to `GET /reporting/metrics` for aggregating results by day, week, or month

You can now pass `interval` (`daily`, `weekly`, or `monthly`) when fetching metrics to have values aggregated over that period, instead of only receiving a single value for the full `start_date`/`end_date` range.

## 6.7.2 (2026-09-22)

### Fixes

- Fix example values for a small number of fields to comply with the schema

## 6.7.1 (2026-09-22)

### Fixes

- Fix schema definition/component names to avoid losing types in openapi schema

## 6.7.0 (2026-09-17)

### Features

- Add "reference" to Create Bank Account Holder Verification

## 6.6.3 (2026-09-16)

### Fixes

- Clean up docs and use a shared definition of event `include` and `resource_type` enums

## 6.6.2 (2026-09-14)

### Fixes

#### Fix nullable field declarations and missing properties across multiple resources

Adds `null` to type declarations for fields that legitimately return nil across redirect_flows, webhooks, scheme_identifiers, customer_bank_accounts, outbound_payments, and billing_request_with_actions. Also adds the missing `period_alignment` property to mandate consent_parameters.

## 6.6.1 (2026-09-14)

### Fixes

#### Add missing enum values to schema definitions

Adds `sepa_credit_transfer` and `sepa_instant_credit_transfer` to the complete scheme enum, adds hosted payment flow sources to the event source/type enum, and makes `creditor_type` nullable for legacy creditors.

## 6.6.0 (2026-09-11)

### Features

#### Use specific sub-endpoints URLs for create /instalment_schedules: with_schedule and with_dates

The two variants for creating an instalment_schedule were surfaced as separate functions. However, they both went to the same URL and endpoint on the backend.

This created some bugs in generating our openapi schema and therefore our API reference documentation.

Therefore, we've added specific URLs for each endpoint aliased to the original one: `POST /instalment_schedules/with_dates` or `POST /instalment_schedules/with_schedule`.

The existing POST /instalment_schedules endpoint is unchanged and will remain available for the foreseeable future.

Client libraries will now use the specific endpoint matching the method - if you are stubbing the HTTP call you may need to update those stubs.

## 6.5.5 (2026-09-09)

### Fixes

- Add remember_me to ui_components bootstrap endpoint

## 6.5.4 (2026-09-08)

### Fixes

#### Remove incorrect Pro/Enterprise restriction from mandate and customer bank account endpoints

The "Create a mandate", "Reinstate a mandate", and "Create a customer bank account" endpoints incorrectly stated they were restricted to GoCardless Pro and Enterprise accounts. Custom payment pages are available to any merchant — they are not package-restricted.

## 6.5.3 (2026-09-07)

### Fixes

- Update code samples to match change to integer types for amounts etc

## 6.5.2 (2026-09-04)

### Fixes

#### Define common titles for common types

The intention is to make it possible to define common types in generated code.
Instead of ~37 different currency enum types which are all equivalent, we could have one.

## 6.5.1 (2026-09-02)

### Fixes

- Fix typo in Mandate next_possible_standard_ach_charge_date description

## 6.5.0 (2026-09-01)

### Features

#### Add `app_connected_organisations` export type

Exports can now be created with `resource_type: app_connected_organisations`, allowing connected merchant details to be exported.

## 6.4.6 (2026-08-27)

### Fixes

- Fix typo in subscription status description

## 6.4.5 (2026-08-13)

### Fixes

- Fix typo in mandate_import_entry status description

## 6.4.4 (2026-08-13)

### Fixes

- Fix typo in mandate_import_entry status description

## 6.4.3 (2026-08-13)

### Fixes

- Fix typo in mandate_import_entry status description

## 6.4.2 (2026-08-12)

### Fixes

- Upgrade inflect dependency from 0.0.1 to 0.0.2

## 6.4.1

Start of changelog tracking with Knope. See [GitHub releases](https://github.com/gocardless/gocardless-pro-go/releases) for the history of earlier versions.
