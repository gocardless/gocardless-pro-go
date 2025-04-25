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

// ScenarioSimulatorService manages scenario_simulators
type ScenarioSimulatorServiceImpl struct {
	config Config
}

// ScenarioSimulator model
type ScenarioSimulator struct {
	Id string `url:"id,omitempty" json:"id,omitempty"`
}

type ScenarioSimulatorService interface {
	Run(ctx context.Context, identity string, p ScenarioSimulatorRunParams, opts ...RequestOption) (*ScenarioSimulator, error)
}

type ScenarioSimulatorRunParamsLinks struct {
	Resource string `url:"resource,omitempty" json:"resource,omitempty"`
}

// ScenarioSimulatorRunParams parameters
type ScenarioSimulatorRunParams struct {
	Links *ScenarioSimulatorRunParamsLinks `url:"links,omitempty" json:"links,omitempty"`
}

// Run
// Runs the specific scenario simulator against the specific resource
func (s *ScenarioSimulatorServiceImpl) Run(ctx context.Context, identity string, p ScenarioSimulatorRunParams, opts ...RequestOption) (*ScenarioSimulator, error) {
	uri, err := url.Parse(fmt.Sprintf(s.config.Endpoint()+"/scenario_simulators/%v/actions/run",
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
	req.Header.Set("GoCardless-Client-Version", "4.7.0")
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
		Err               *APIError          `json:"error"`
		ScenarioSimulator *ScenarioSimulator `json:"scenario_simulators"`
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

	if result.ScenarioSimulator == nil {
		return nil, errors.New("missing result")
	}

	return result.ScenarioSimulator, nil
}
