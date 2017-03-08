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

func responseErr(r *http.Response) error {
	switch {
	case 200 <= r.StatusCode && r.StatusCode < 400:
		return nil
	default:
		var cause error
		var result struct {
			Err *APIError `json:"error"`
		}

		json.NewDecoder(r.Body).Decode(&result)
		if result.Err != nil {
			cause = result.Err
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
