package gocardless

import (
	"errors"
	"net/http"
	"net/url"
)

const (

	// Live environment
	LiveEndpoint = "https://api.gocardless.com"

	// Sandbox environment
	SandboxEndpoint = "https://api-sandbox.gocardless.com"
)

// ConfigOption used to initialise the client
type ConfigOption func(Config) error

type Config interface {
	Token() string
	Endpoint() string
	Client() *http.Client
}

type config struct {
	token    string
	endpoint string
	client   *http.Client
}

func (c *config) Token() string {
	return c.token
}

func (c *config) Endpoint() string {
	return c.endpoint
}

func (c *config) Client() *http.Client {
	return c.client
}

// WithEndpoint configures the endpoint hosting the API
func WithEndpoint(endpoint string) ConfigOption {
	return func(cfg Config) error {
		u, err := url.Parse(endpoint)
		if err != nil {
			return err
		}
		if c, ok := cfg.(*config); ok {
			c.endpoint = u.String()
		} else {
			return errors.New("invalid input, input is not of type config")
		}
		return nil
	}
}

// WithClient configures the net/http client
func WithClient(client *http.Client) ConfigOption {
	return func(cfg Config) error {
		if c, ok := cfg.(*config); ok {
			c.client = client
		} else {
			return errors.New("invalid input, input is not of type config")
		}
		return nil
	}
}

func NewConfig(token string, configOpts ...ConfigOption) (Config, error) {
	if token == "" {
		return nil, errors.New("token required")
	}

	config := &config{
		token:    token,
		endpoint: LiveEndpoint,
	}

	for _, configOpt := range configOpts {
		if err := configOpt(config); err != nil {
			return nil, err
		}
	}

	return config, nil
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
