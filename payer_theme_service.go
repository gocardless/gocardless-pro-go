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

// PayerThemeService manages payer_themes
type PayerThemeServiceImpl struct {
	config Config
}

// PayerTheme model
type PayerTheme struct {
	Id string `url:"id,omitempty" json:"id,omitempty"`
}

type PayerThemeService interface {
	CreateForCreditor(ctx context.Context, p PayerThemeCreateForCreditorParams, opts ...RequestOption) (*PayerTheme, error)
}

type PayerThemeCreateForCreditorParamsLinks struct {
	Creditor string `url:"creditor,omitempty" json:"creditor,omitempty"`
}

// PayerThemeCreateForCreditorParams parameters
type PayerThemeCreateForCreditorParams struct {
	ButtonBackgroundColour string                                  `url:"button_background_colour,omitempty" json:"button_background_colour,omitempty"`
	ContentBoxBorderColour string                                  `url:"content_box_border_colour,omitempty" json:"content_box_border_colour,omitempty"`
	HeaderBackgroundColour string                                  `url:"header_background_colour,omitempty" json:"header_background_colour,omitempty"`
	LinkTextColour         string                                  `url:"link_text_colour,omitempty" json:"link_text_colour,omitempty"`
	Links                  *PayerThemeCreateForCreditorParamsLinks `url:"links,omitempty" json:"links,omitempty"`
}

// CreateForCreditor
// Creates a new payer theme associated with a creditor. If a creditor already
// has payer themes, this will update the existing payer theme linked to the
// creditor.
func (s *PayerThemeServiceImpl) CreateForCreditor(ctx context.Context, p PayerThemeCreateForCreditorParams, opts ...RequestOption) (*PayerTheme, error) {
	uri, err := url.Parse(fmt.Sprintf(s.config.Endpoint() + "/branding/payer_themes"))
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
	req.Header.Set("GoCardless-Client-Version", "4.6.0")
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
		Err        *APIError   `json:"error"`
		PayerTheme *PayerTheme `json:"payer_themes"`
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

	if result.PayerTheme == nil {
		return nil, errors.New("missing result")
	}

	return result.PayerTheme, nil
}
