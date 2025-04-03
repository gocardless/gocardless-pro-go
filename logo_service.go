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

// LogoService manages logos
type LogoServiceImpl struct {
	config Config
}

// Logo model
type Logo struct {
	Id string `url:"id,omitempty" json:"id,omitempty"`
}

type LogoService interface {
	CreateForCreditor(ctx context.Context, p LogoCreateForCreditorParams, opts ...RequestOption) (*Logo, error)
}

type LogoCreateForCreditorParamsLinks struct {
	Creditor string `url:"creditor,omitempty" json:"creditor,omitempty"`
}

// LogoCreateForCreditorParams parameters
type LogoCreateForCreditorParams struct {
	Image string                            `url:"image,omitempty" json:"image,omitempty"`
	Links *LogoCreateForCreditorParamsLinks `url:"links,omitempty" json:"links,omitempty"`
}

// CreateForCreditor
// Creates a new logo associated with a creditor. If a creditor already has a
// logo, this will update the existing logo linked to the creditor.
//
// We support JPG and PNG formats. Your logo will be scaled to a maximum of
// 300px by 40px. For more guidance on how to upload logos that will look
// great across your customer payment page and notification emails see
// [here](https://developer.gocardless.com/gc-embed/setting-up-branding#tips_for_uploading_your_logo).
func (s *LogoServiceImpl) CreateForCreditor(ctx context.Context, p LogoCreateForCreditorParams, opts ...RequestOption) (*Logo, error) {
	uri, err := url.Parse(fmt.Sprintf(s.config.Endpoint() + "/branding/logos"))
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
	req.Header.Set("GoCardless-Client-Version", "4.5.0")
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
		Err  *APIError `json:"error"`
		Logo *Logo     `json:"logos"`
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

	if result.Logo == nil {
		return nil, errors.New("missing result")
	}

	return result.Logo, nil
}
