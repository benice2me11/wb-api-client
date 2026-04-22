package transport

import (
	"bytes"
	"context"
	"io"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// RetryPolicy controls retry behavior for HTTP requests.
type RetryPolicy struct {
	// MaxAttempts is total number of attempts including the first call.
	MaxAttempts int
	// BaseDelay is the initial backoff delay.
	BaseDelay time.Duration
	// MaxDelay caps exponential backoff.
	MaxDelay time.Duration
	// Jitter is a fraction in [0,1]. 0.2 means +/-20% randomization.
	Jitter float64
}

// DefaultRetryPolicy returns conservative defaults for WB APIs.
func DefaultRetryPolicy() RetryPolicy {
	return RetryPolicy{
		MaxAttempts: 3,
		BaseDelay:   500 * time.Millisecond,
		MaxDelay:    10 * time.Second,
		Jitter:      0.2,
	}
}

func (p RetryPolicy) normalize() RetryPolicy {
	defaults := DefaultRetryPolicy()
	if p.MaxAttempts <= 0 {
		p.MaxAttempts = defaults.MaxAttempts
	}
	if p.BaseDelay <= 0 {
		p.BaseDelay = defaults.BaseDelay
	}
	if p.MaxDelay <= 0 {
		p.MaxDelay = defaults.MaxDelay
	}
	if p.Jitter < 0 {
		p.Jitter = 0
	}
	if p.Jitter > 1 {
		p.Jitter = 1
	}
	if p.MaxDelay < p.BaseDelay {
		p.MaxDelay = p.BaseDelay
	}
	return p
}

// RetryTransport retries failed responses for 429 and 5xx statuses.
type RetryTransport struct {
	next   http.RoundTripper
	policy RetryPolicy
	sleep  func(context.Context, time.Duration) error
	rand   *rand.Rand
}

// NewRetryTransport creates a retrying round tripper.
func NewRetryTransport(next http.RoundTripper, policy RetryPolicy) *RetryTransport {
	if next == nil {
		next = http.DefaultTransport
	}
	return &RetryTransport{
		next:   next,
		policy: policy.normalize(),
		sleep:  sleepWithContext,
		rand:   rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (t *RetryTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if req == nil {
		return nil, nil
	}

	if err := ensureReplayableBody(req); err != nil {
		return nil, err
	}

	for attempt := 1; attempt <= t.policy.MaxAttempts; attempt++ {
		attemptReq, err := cloneRequest(req)
		if err != nil {
			return nil, err
		}

		resp, err := t.next.RoundTrip(attemptReq)
		if err != nil {
			return nil, err
		}

		if !shouldRetryStatus(resp.StatusCode) || attempt == t.policy.MaxAttempts {
			return resp, nil
		}

		delay := t.retryDelay(resp, attempt)
		drainAndClose(resp.Body)
		if err := t.sleep(req.Context(), delay); err != nil {
			return nil, err
		}
	}

	return nil, nil
}

func shouldRetryStatus(status int) bool {
	return status == http.StatusTooManyRequests || (status >= http.StatusInternalServerError && status <= 599)
}

func (t *RetryTransport) retryDelay(resp *http.Response, attempt int) time.Duration {
	if resp != nil {
		if d, ok := parseRateLimitRetry(resp.Header.Get("X-Ratelimit-Retry")); ok {
			return clampDelay(d, t.policy.MaxDelay)
		}
		if d, ok := parseRetryAfter(resp.Header.Get("Retry-After"), time.Now()); ok {
			return clampDelay(d, t.policy.MaxDelay)
		}
	}

	exp := float64(t.policy.BaseDelay) * math.Pow(2, float64(attempt-1))
	d := time.Duration(exp)
	if d > t.policy.MaxDelay {
		d = t.policy.MaxDelay
	}
	if t.policy.Jitter > 0 {
		min := 1 - t.policy.Jitter
		max := 1 + t.policy.Jitter
		factor := min + t.rand.Float64()*(max-min)
		d = time.Duration(float64(d) * factor)
	}
	if d < 0 {
		return 0
	}
	return d
}

func clampDelay(d, max time.Duration) time.Duration {
	if d < 0 {
		return 0
	}
	if max > 0 && d > max {
		return max
	}
	return d
}

func parseRateLimitRetry(value string) (time.Duration, bool) {
	v := strings.TrimSpace(value)
	if v == "" {
		return 0, false
	}

	if seconds, err := strconv.ParseFloat(v, 64); err == nil {
		if seconds < 0 {
			return 0, false
		}
		return time.Duration(seconds * float64(time.Second)), true
	}

	d, err := time.ParseDuration(v)
	if err != nil || d < 0 {
		return 0, false
	}
	return d, true
}

func parseRetryAfter(value string, now time.Time) (time.Duration, bool) {
	v := strings.TrimSpace(value)
	if v == "" {
		return 0, false
	}

	if seconds, err := strconv.ParseInt(v, 10, 64); err == nil {
		if seconds < 0 {
			return 0, false
		}
		return time.Duration(seconds) * time.Second, true
	}

	tm, err := http.ParseTime(v)
	if err != nil {
		return 0, false
	}
	d := tm.Sub(now)
	if d < 0 {
		return 0, true
	}
	return d, true
}

func ensureReplayableBody(req *http.Request) error {
	if req.Body == nil || req.GetBody != nil {
		return nil
	}

	buf, err := io.ReadAll(req.Body)
	if err != nil {
		return err
	}
	_ = req.Body.Close()

	rebuild := func() io.ReadCloser {
		return io.NopCloser(bytes.NewReader(buf))
	}
	req.Body = rebuild()
	req.GetBody = func() (io.ReadCloser, error) {
		return rebuild(), nil
	}
	return nil
}

func cloneRequest(req *http.Request) (*http.Request, error) {
	clone := req.Clone(req.Context())
	if req.GetBody != nil {
		body, err := req.GetBody()
		if err != nil {
			return nil, err
		}
		clone.Body = body
	}
	return clone, nil
}

func drainAndClose(body io.ReadCloser) {
	if body == nil {
		return
	}
	_, _ = io.Copy(io.Discard, body)
	_ = body.Close()
}

func sleepWithContext(ctx context.Context, d time.Duration) error {
	if d <= 0 {
		return nil
	}
	if ctx == nil {
		time.Sleep(d)
		return nil
	}

	timer := time.NewTimer(d)
	defer timer.Stop()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-timer.C:
		return nil
	}
}
