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

// EventService manages events
type EventServiceImpl struct {
	config Config
}

type EventCustomerNotifications struct {
	Deadline  string `url:"deadline,omitempty" json:"deadline,omitempty"`
	Id        string `url:"id,omitempty" json:"id,omitempty"`
	Mandatory bool   `url:"mandatory,omitempty" json:"mandatory,omitempty"`
	Type      string `url:"type,omitempty" json:"type,omitempty"`
}

type EventDetails struct {
	BankAccountId    string `url:"bank_account_id,omitempty" json:"bank_account_id,omitempty"`
	Cause            string `url:"cause,omitempty" json:"cause,omitempty"`
	Currency         string `url:"currency,omitempty" json:"currency,omitempty"`
	Description      string `url:"description,omitempty" json:"description,omitempty"`
	NotRetriedReason string `url:"not_retried_reason,omitempty" json:"not_retried_reason,omitempty"`
	Origin           string `url:"origin,omitempty" json:"origin,omitempty"`
	Property         string `url:"property,omitempty" json:"property,omitempty"`
	ReasonCode       string `url:"reason_code,omitempty" json:"reason_code,omitempty"`
	Scheme           string `url:"scheme,omitempty" json:"scheme,omitempty"`
	WillAttemptRetry bool   `url:"will_attempt_retry,omitempty" json:"will_attempt_retry,omitempty"`
}

type EventLinks struct {
	BankAuthorisation           string `url:"bank_authorisation,omitempty" json:"bank_authorisation,omitempty"`
	BillingRequest              string `url:"billing_request,omitempty" json:"billing_request,omitempty"`
	BillingRequestFlow          string `url:"billing_request_flow,omitempty" json:"billing_request_flow,omitempty"`
	Creditor                    string `url:"creditor,omitempty" json:"creditor,omitempty"`
	Customer                    string `url:"customer,omitempty" json:"customer,omitempty"`
	CustomerBankAccount         string `url:"customer_bank_account,omitempty" json:"customer_bank_account,omitempty"`
	InstalmentSchedule          string `url:"instalment_schedule,omitempty" json:"instalment_schedule,omitempty"`
	Mandate                     string `url:"mandate,omitempty" json:"mandate,omitempty"`
	MandateRequestMandate       string `url:"mandate_request_mandate,omitempty" json:"mandate_request_mandate,omitempty"`
	NewCustomerBankAccount      string `url:"new_customer_bank_account,omitempty" json:"new_customer_bank_account,omitempty"`
	NewMandate                  string `url:"new_mandate,omitempty" json:"new_mandate,omitempty"`
	Organisation                string `url:"organisation,omitempty" json:"organisation,omitempty"`
	ParentEvent                 string `url:"parent_event,omitempty" json:"parent_event,omitempty"`
	PayerAuthorisation          string `url:"payer_authorisation,omitempty" json:"payer_authorisation,omitempty"`
	Payment                     string `url:"payment,omitempty" json:"payment,omitempty"`
	PaymentRequestPayment       string `url:"payment_request_payment,omitempty" json:"payment_request_payment,omitempty"`
	Payout                      string `url:"payout,omitempty" json:"payout,omitempty"`
	PreviousCustomerBankAccount string `url:"previous_customer_bank_account,omitempty" json:"previous_customer_bank_account,omitempty"`
	Refund                      string `url:"refund,omitempty" json:"refund,omitempty"`
	Subscription                string `url:"subscription,omitempty" json:"subscription,omitempty"`
}

// Event model
type Event struct {
	Action                string                       `url:"action,omitempty" json:"action,omitempty"`
	CreatedAt             string                       `url:"created_at,omitempty" json:"created_at,omitempty"`
	CustomerNotifications []EventCustomerNotifications `url:"customer_notifications,omitempty" json:"customer_notifications,omitempty"`
	Details               *EventDetails                `url:"details,omitempty" json:"details,omitempty"`
	Id                    string                       `url:"id,omitempty" json:"id,omitempty"`
	Links                 *EventLinks                  `url:"links,omitempty" json:"links,omitempty"`
	Metadata              map[string]interface{}       `url:"metadata,omitempty" json:"metadata,omitempty"`
	ResourceType          string                       `url:"resource_type,omitempty" json:"resource_type,omitempty"`
}

type EventService interface {
	List(ctx context.Context, p EventListParams, opts ...RequestOption) (*EventListResult, error)
	All(ctx context.Context, p EventListParams, opts ...RequestOption) *EventListPagingIterator
	Get(ctx context.Context, identity string, opts ...RequestOption) (*Event, error)
}

type EventListParamsCreatedAt struct {
	Gt  string `url:"gt,omitempty" json:"gt,omitempty"`
	Gte string `url:"gte,omitempty" json:"gte,omitempty"`
	Lt  string `url:"lt,omitempty" json:"lt,omitempty"`
	Lte string `url:"lte,omitempty" json:"lte,omitempty"`
}

// EventListParams parameters
type EventListParams struct {
	Action             string                    `url:"action,omitempty" json:"action,omitempty"`
	After              string                    `url:"after,omitempty" json:"after,omitempty"`
	Before             string                    `url:"before,omitempty" json:"before,omitempty"`
	BillingRequest     string                    `url:"billing_request,omitempty" json:"billing_request,omitempty"`
	CreatedAt          *EventListParamsCreatedAt `url:"created_at,omitempty" json:"created_at,omitempty"`
	Include            string                    `url:"include,omitempty" json:"include,omitempty"`
	InstalmentSchedule string                    `url:"instalment_schedule,omitempty" json:"instalment_schedule,omitempty"`
	Limit              int                       `url:"limit,omitempty" json:"limit,omitempty"`
	Mandate            string                    `url:"mandate,omitempty" json:"mandate,omitempty"`
	ParentEvent        string                    `url:"parent_event,omitempty" json:"parent_event,omitempty"`
	PayerAuthorisation string                    `url:"payer_authorisation,omitempty" json:"payer_authorisation,omitempty"`
	Payment            string                    `url:"payment,omitempty" json:"payment,omitempty"`
	Payout             string                    `url:"payout,omitempty" json:"payout,omitempty"`
	Refund             string                    `url:"refund,omitempty" json:"refund,omitempty"`
	ResourceType       string                    `url:"resource_type,omitempty" json:"resource_type,omitempty"`
	Subscription       string                    `url:"subscription,omitempty" json:"subscription,omitempty"`
}

type EventListResultMetaCursors struct {
	After  string `url:"after,omitempty" json:"after,omitempty"`
	Before string `url:"before,omitempty" json:"before,omitempty"`
}

type EventListResultMeta struct {
	Cursors *EventListResultMetaCursors `url:"cursors,omitempty" json:"cursors,omitempty"`
	Limit   int                         `url:"limit,omitempty" json:"limit,omitempty"`
}

type EventListResult struct {
	Events []Event             `json:"events"`
	Meta   EventListResultMeta `url:"meta,omitempty" json:"meta,omitempty"`
}

// List
// Returns a [cursor-paginated](#api-usage-cursor-pagination) list of your
// events.
func (s *EventServiceImpl) List(ctx context.Context, p EventListParams, opts ...RequestOption) (*EventListResult, error) {
	uri, err := url.Parse(fmt.Sprintf(s.config.Endpoint() + "/events"))
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
	req = req.WithContext(ctx)
	req.Header.Set("Authorization", "Bearer "+s.config.Token())
	req.Header.Set("GoCardless-Version", "2015-07-06")
	req.Header.Set("GoCardless-Client-Library", "gocardless-pro-go")
	req.Header.Set("GoCardless-Client-Version", "2.11.0")
	req.Header.Set("User-Agent", userAgent)

	for key, value := range o.headers {
		req.Header.Set(key, value)
	}

	client := s.config.Client()
	if client == nil {
		client = http.DefaultClient
	}

	var result struct {
		Err *APIError `json:"error"`
		*EventListResult
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

	if result.EventListResult == nil {
		return nil, errors.New("missing result")
	}

	return result.EventListResult, nil
}

type EventListPagingIterator struct {
	cursor         string
	response       *EventListResult
	params         EventListParams
	service        *EventServiceImpl
	requestOptions []RequestOption
}

func (c *EventListPagingIterator) Next() bool {
	if c.cursor == "" && c.response != nil {
		return false
	}

	return true
}

func (c *EventListPagingIterator) Value(ctx context.Context) (*EventListResult, error) {
	if !c.Next() {
		return c.response, nil
	}

	s := c.service
	p := c.params
	p.After = c.cursor

	uri, err := url.Parse(fmt.Sprintf(s.config.Endpoint() + "/events"))

	if err != nil {
		return nil, err
	}

	o := &requestOptions{
		retries: 3,
	}
	for _, opt := range c.requestOptions {
		err := opt(o)
		if err != nil {
			return nil, err
		}
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

	req = req.WithContext(ctx)
	req.Header.Set("Authorization", "Bearer "+s.config.Token())
	req.Header.Set("GoCardless-Version", "2015-07-06")
	req.Header.Set("GoCardless-Client-Library", "gocardless-pro-go")
	req.Header.Set("GoCardless-Client-Version", "2.11.0")
	req.Header.Set("User-Agent", userAgent)

	for key, value := range o.headers {
		req.Header.Set(key, value)
	}
	client := s.config.Client()
	if client == nil {
		client = http.DefaultClient
	}

	var result struct {
		Err *APIError `json:"error"`
		*EventListResult
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

	if result.EventListResult == nil {
		return nil, errors.New("missing result")
	}

	c.response = result.EventListResult
	c.cursor = c.response.Meta.Cursors.After
	return c.response, nil
}

func (s *EventServiceImpl) All(ctx context.Context,
	p EventListParams,
	opts ...RequestOption) *EventListPagingIterator {
	return &EventListPagingIterator{
		params:         p,
		service:        s,
		requestOptions: opts,
	}
}

// Get
// Retrieves the details of a single event.
func (s *EventServiceImpl) Get(ctx context.Context, identity string, opts ...RequestOption) (*Event, error) {
	uri, err := url.Parse(fmt.Sprintf(s.config.Endpoint()+"/events/%v",
		identity))
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

	var body io.Reader

	req, err := http.NewRequest("GET", uri.String(), body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	req.Header.Set("Authorization", "Bearer "+s.config.Token())
	req.Header.Set("GoCardless-Version", "2015-07-06")
	req.Header.Set("GoCardless-Client-Library", "gocardless-pro-go")
	req.Header.Set("GoCardless-Client-Version", "2.11.0")
	req.Header.Set("User-Agent", userAgent)

	for key, value := range o.headers {
		req.Header.Set(key, value)
	}

	client := s.config.Client()
	if client == nil {
		client = http.DefaultClient
	}

	var result struct {
		Err   *APIError `json:"error"`
		Event *Event    `json:"events"`
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

	if result.Event == nil {
		return nil, errors.New("missing result")
	}

	return result.Event, nil
}
