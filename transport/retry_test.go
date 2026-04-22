package transport

import (
	"context"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"
)

type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}

func newResponse(status int, body string, headers map[string]string) *http.Response {
	h := make(http.Header)
	for k, v := range headers {
		h.Set(k, v)
	}
	return &http.Response{
		StatusCode: status,
		Header:     h,
		Body:       io.NopCloser(strings.NewReader(body)),
	}
}

func TestRetryTransport_Retry429Then200(t *testing.T) {
	attempts := 0
	delays := make([]time.Duration, 0, 2)

	rt := NewRetryTransport(roundTripFunc(func(req *http.Request) (*http.Response, error) {
		attempts++
		if attempts == 1 {
			return newResponse(http.StatusTooManyRequests, `{"detail":"limited"}`, map[string]string{
				"X-Ratelimit-Retry": "1",
			}), nil
		}
		return newResponse(http.StatusOK, `{"ok":true}`, nil), nil
	}), RetryPolicy{MaxAttempts: 3, BaseDelay: 10 * time.Millisecond, MaxDelay: time.Second, Jitter: 0})

	rt.sleep = func(ctx context.Context, d time.Duration) error {
		delays = append(delays, d)
		return nil
	}

	req, _ := http.NewRequest(http.MethodGet, "https://example.com/ping", nil)
	resp, err := rt.RoundTrip(req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("unexpected status: %d", resp.StatusCode)
	}
	if attempts != 2 {
		t.Fatalf("expected 2 attempts, got %d", attempts)
	}
	if len(delays) != 1 || delays[0] != time.Second {
		t.Fatalf("expected one 1s delay, got %v", delays)
	}
}

func TestRetryTransport_InvalidRateLimitHeaderFallsBackToBackoff(t *testing.T) {
	attempts := 0
	delays := make([]time.Duration, 0, 2)

	rt := NewRetryTransport(roundTripFunc(func(req *http.Request) (*http.Response, error) {
		attempts++
		if attempts == 1 {
			return newResponse(http.StatusTooManyRequests, `{"detail":"limited"}`, map[string]string{
				"X-Ratelimit-Retry": "invalid",
			}), nil
		}
		return newResponse(http.StatusOK, `{"ok":true}`, nil), nil
	}), RetryPolicy{MaxAttempts: 3, BaseDelay: 50 * time.Millisecond, MaxDelay: time.Second, Jitter: 0})

	rt.sleep = func(ctx context.Context, d time.Duration) error {
		delays = append(delays, d)
		return nil
	}

	req, _ := http.NewRequest(http.MethodGet, "https://example.com/ping", nil)
	resp, err := rt.RoundTrip(req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("unexpected status: %d", resp.StatusCode)
	}
	if attempts != 2 {
		t.Fatalf("expected 2 attempts, got %d", attempts)
	}
	if len(delays) != 1 || delays[0] != 50*time.Millisecond {
		t.Fatalf("expected one 50ms delay, got %v", delays)
	}
}

func TestRetryTransport_StopsOnMaxAttempts(t *testing.T) {
	attempts := 0
	rt := NewRetryTransport(roundTripFunc(func(req *http.Request) (*http.Response, error) {
		attempts++
		return newResponse(http.StatusTooManyRequests, `{"detail":"limited"}`, nil), nil
	}), RetryPolicy{MaxAttempts: 3, BaseDelay: time.Millisecond, MaxDelay: time.Millisecond, Jitter: 0})

	rt.sleep = func(ctx context.Context, d time.Duration) error { return nil }

	req, _ := http.NewRequest(http.MethodGet, "https://example.com/ping", nil)
	resp, err := rt.RoundTrip(req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusTooManyRequests {
		t.Fatalf("unexpected status: %d", resp.StatusCode)
	}
	if attempts != 3 {
		t.Fatalf("expected 3 attempts, got %d", attempts)
	}
}

func TestRetryTransport_Retry5xx(t *testing.T) {
	attempts := 0
	rt := NewRetryTransport(roundTripFunc(func(req *http.Request) (*http.Response, error) {
		attempts++
		if attempts == 1 {
			return newResponse(http.StatusInternalServerError, `{"detail":"boom"}`, nil), nil
		}
		return newResponse(http.StatusOK, `{"ok":true}`, nil), nil
	}), RetryPolicy{MaxAttempts: 3, BaseDelay: time.Millisecond, MaxDelay: time.Millisecond, Jitter: 0})

	rt.sleep = func(ctx context.Context, d time.Duration) error { return nil }

	req, _ := http.NewRequest(http.MethodGet, "https://example.com/ping", nil)
	resp, err := rt.RoundTrip(req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("unexpected status: %d", resp.StatusCode)
	}
	if attempts != 2 {
		t.Fatalf("expected 2 attempts, got %d", attempts)
	}
}

func TestRetryTransport_ContextCancel(t *testing.T) {
	rt := NewRetryTransport(roundTripFunc(func(req *http.Request) (*http.Response, error) {
		return newResponse(http.StatusInternalServerError, `{"detail":"boom"}`, nil), nil
	}), RetryPolicy{MaxAttempts: 3, BaseDelay: time.Second, MaxDelay: time.Second, Jitter: 0})

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, "https://example.com/ping", nil)
	_, err := rt.RoundTrip(req)
	if err == nil {
		t.Fatalf("expected context cancellation error")
	}
	if err != context.Canceled {
		t.Fatalf("expected context.Canceled, got %v", err)
	}
}
