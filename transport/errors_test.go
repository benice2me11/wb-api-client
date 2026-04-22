package transport

import (
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"
)

func TestWrapResponseError_RateLimit(t *testing.T) {
	resp := &http.Response{
		StatusCode: http.StatusTooManyRequests,
		Header: http.Header{
			"X-Ratelimit-Retry":     []string{"2"},
			"X-Ratelimit-Limit":     []string{"300"},
			"X-Ratelimit-Remaining": []string{"0"},
			"X-Request-Id":          []string{"req-123"},
		},
		Body: io.NopCloser(strings.NewReader(`{"detail":"too many"}`)),
	}

	err := WrapResponseError(resp, errors.New("raw error"))
	if !IsRateLimit(err) {
		t.Fatalf("expected rate limit error, got %T", err)
	}

	var rateErr *RateLimitError
	if !errors.As(err, &rateErr) {
		t.Fatalf("expected RateLimitError via errors.As")
	}

	if rateErr.StatusCode != http.StatusTooManyRequests {
		t.Fatalf("unexpected status code: %d", rateErr.StatusCode)
	}
	if rateErr.Meta.RetryAfter != 2*time.Second {
		t.Fatalf("unexpected retry-after: %s", rateErr.Meta.RetryAfter)
	}
	if rateErr.Meta.RequestID != "req-123" {
		t.Fatalf("unexpected request id: %s", rateErr.Meta.RequestID)
	}
	if rateErr.Meta.Limit == nil || *rateErr.Meta.Limit != 300 {
		t.Fatalf("unexpected limit: %v", rateErr.Meta.Limit)
	}
}

func TestWrapResponseError_APIError(t *testing.T) {
	resp := &http.Response{
		StatusCode: http.StatusUnauthorized,
		Header: http.Header{
			"X-Request-Id": []string{"req-401"},
		},
		Body: io.NopCloser(strings.NewReader(`{"detail":"unauthorized"}`)),
	}

	err := WrapResponseError(resp, errors.New("raw error"))
	if IsRateLimit(err) {
		t.Fatalf("unexpected rate limit error")
	}

	var apiErr *APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("expected APIError, got %T", err)
	}
	if apiErr.StatusCode != http.StatusUnauthorized {
		t.Fatalf("unexpected status code: %d", apiErr.StatusCode)
	}
	if apiErr.RequestID != "req-401" {
		t.Fatalf("unexpected request id: %s", apiErr.RequestID)
	}
	if string(apiErr.Body) != `{"detail":"unauthorized"}` {
		t.Fatalf("unexpected body: %s", string(apiErr.Body))
	}
}
