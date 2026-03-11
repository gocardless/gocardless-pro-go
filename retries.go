package gocardless

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type temporary interface {
	Temporary() bool
}

type waiter interface {
	Wait()
}

func try(attempts int, fn func() error) error {
	var err error
	if attempts < 1 {
		attempts = 1
	}
	for i := attempts; i > 0; i-- {
		err = fn()
		if err == nil {
			return nil
		}
		if t, ok := err.(temporary); !ok || !t.Temporary() {
			return err
		}
		if w, ok := err.(waiter); ok {
			w.Wait()
		}
	}
	return err
}

// FlexError encapsulates an error response that may be a string or an object
type FlexError struct {
	Msg *string
	Err *APIError
}

func (e *FlexError) UnmarshalJSON(data []byte) error {
	if len(data) == 0 || string(data) == "null" {
		return nil
	}
	// Read the raw JSON value and see if it's a string
	if data[0] == '"' {
		return json.Unmarshal(data, &e.Msg)
	}
	e.Err = &APIError{}
	return json.Unmarshal(data, e.Err)
}

func responseErr(r *http.Response) error {
	switch {
	case 200 <= r.StatusCode && r.StatusCode < 400:
		return nil
	default:
		var cause error
		var result struct {
			Error FlexError `json:"error"`
		}

		err := json.NewDecoder(r.Body).Decode(&result)
		if err != nil {
			return fmt.Errorf("decoding error response: %w", err)
		}

		if result.Error.Err != nil {
			cause = result.Error.Err
		} else if result.Error.Msg != nil {
			// Wrap the error message into an API erro
			cause = &APIError{
				Message: *result.Error.Msg,
				Code:    r.StatusCode,
			}
		}

		return &responseError{
			res: r,
			err: cause,
		}
	}
}

type responseError struct {
	res *http.Response
	err error
}

func (r *responseError) Cause() error {
	return r.err
}

func (r *responseError) Unwrap() error {
	return r.err
}

func (r *responseError) Error() string {
	return r.res.Status
}

func (r *responseError) Temporary() bool {
	switch {
	case 500 <= r.res.StatusCode:
		return true
	case r.res.StatusCode == http.StatusTooManyRequests:
		return true
	default:
		return false
	}
}

func (r *responseError) Wait() {
	rem, err := strconv.Atoi(r.res.Header.Get("RateLimit-Remaining"))
	if err != nil || rem > 0 {
		return
	}
	reset := r.res.Header.Get("RateLimit-Reset")
	t, err := time.Parse(time.RFC1123, reset)
	if err != nil {
		t, err = time.Parse(time.RFC1123Z, reset)
	}
	if err != nil {
		return
	}
	time.Sleep(time.Until(t))
}

// NewIdempotencyKey generates a random and unique idempotency key
func NewIdempotencyKey() string {
	buf := make([]byte, 10)
	_, err := rand.Read(buf)
	if err != nil {
		panic("failed do generate random key")
	}
	t := time.Now().UTC().UnixNano()
	return fmt.Sprintf("%d-%s", t, base64.RawURLEncoding.EncodeToString(buf))
}
