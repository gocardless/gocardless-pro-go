package gocardless

import (
	"errors"
	"net/http"
	"net/url"
)

// Option used to initialise the client
type Option func(*options) error

type options struct {
	endpoint string
	client   *http.Client
}

// WithEndpoint configures the endpoint hosting the API
func WithEndpoint(endpoint string) Option {
	return func(opts *options) error {
		u, err := url.Parse(endpoint)
		if err != nil {
			return err
		}
		opts.endpoint = u.String()
		return nil
	}
}

// WithClient configures the net/http client
func WithClient(c *http.Client) Option {
	return func(opts *options) error {
		opts.client = c
		return nil
	}
}

// RequestOption is used to configure a given request
type RequestOption func(*requestOptions) error

type requestOptions struct {
	idempotencyKey string
	retries        int
	headers        map[string]string
}

// WithIdempotencyKey sets an idempotency key so multiple calls to a
// non-idempotent endpoint with the same key are actually idempotent
func WithIdempotencyKey(key string) RequestOption {
	return func(opts *requestOptions) error {
		if opts.idempotencyKey != "" {
			return errors.New("idempotency key is already set")
		}
		opts.idempotencyKey = key
		return nil
	}
}

// WithRetries sets the amount of total retries to make for the request
func WithRetries(n int) RequestOption {
	return func(opts *requestOptions) error {
		opts.retries = n
		return nil
	}
}

// WithoutRetries disables retries for this request
func WithoutRetries() RequestOption {
	return WithRetries(0)
}

// WithHeaders sets headers to be sent for this request
func WithHeaders(headers map[string]string) RequestOption {
	return func(opts *requestOptions) error {
		opts.headers = headers
		return nil
	}
}
