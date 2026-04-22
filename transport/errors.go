package transport

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

// RateLimitMeta includes parsed rate-limit hints.
type RateLimitMeta struct {
	RetryAfter time.Duration
	Limit      *int
	Remaining  *int
	RequestID  string
}

// APIError is a normalized HTTP/API error.
type APIError struct {
	StatusCode int
	Body       []byte
	RequestID  string
	Err        error
}

func (e *APIError) Error() string {
	if e == nil {
		return "<nil>"
	}
	if e.RequestID != "" {
		return fmt.Sprintf("api error status=%d request_id=%s", e.StatusCode, e.RequestID)
	}
	return fmt.Sprintf("api error status=%d", e.StatusCode)
}

func (e *APIError) Unwrap() error {
	if e == nil {
		return nil
	}
	return e.Err
}

// RateLimitError is APIError for HTTP 429 with parsed metadata.
type RateLimitError struct {
	APIError
	Meta RateLimitMeta
}

func (e *RateLimitError) Error() string {
	if e == nil {
		return "<nil>"
	}
	if e.Meta.RetryAfter > 0 {
		return fmt.Sprintf("rate limit exceeded (retry after %s)", e.Meta.RetryAfter)
	}
	return "rate limit exceeded"
}

// IsRateLimit reports whether error is a RateLimitError.
func IsRateLimit(err error) bool {
	var rateErr *RateLimitError
	return errors.As(err, &rateErr)
}

// WrapResponseError converts generated/client errors into normalized APIError/RateLimitError.
func WrapResponseError(resp *http.Response, err error) error {
	if err == nil {
		return nil
	}
	if resp == nil {
		return err
	}

	body := preserveAndReadBody(resp)
	requestID := firstHeader(resp.Header, "X-Request-ID", "X-Request-Id", "X-Request-Id")

	apiErr := &APIError{
		StatusCode: resp.StatusCode,
		Body:       body,
		RequestID:  requestID,
		Err:        err,
	}

	if resp.StatusCode != http.StatusTooManyRequests {
		return apiErr
	}

	rateMeta := RateLimitMeta{RequestID: requestID}
	if d, ok := parseRateLimitRetry(resp.Header.Get("X-Ratelimit-Retry")); ok {
		rateMeta.RetryAfter = d
	}
	rateMeta.Limit = parseHeaderInt(resp.Header.Get("X-Ratelimit-Limit"))
	rateMeta.Remaining = parseHeaderInt(resp.Header.Get("X-Ratelimit-Remaining"))

	return &RateLimitError{
		APIError: *apiErr,
		Meta:     rateMeta,
	}
}

func preserveAndReadBody(resp *http.Response) []byte {
	if resp == nil || resp.Body == nil {
		return nil
	}
	b, _ := io.ReadAll(resp.Body)
	resp.Body = io.NopCloser(bytes.NewBuffer(b))
	return b
}

func firstHeader(h http.Header, keys ...string) string {
	for _, k := range keys {
		if v := h.Get(k); v != "" {
			return v
		}
	}
	return ""
}

func parseHeaderInt(v string) *int {
	if v == "" {
		return nil
	}
	n, err := strconv.Atoi(v)
	if err != nil {
		return nil
	}
	return &n
}
