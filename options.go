package gocardless

import (
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
