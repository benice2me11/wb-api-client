package client_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/benice2me11/wb-api-client/client"
	"github.com/benice2me11/wb-api-client/internal/testkit"
	"github.com/benice2me11/wb-api-client/transport"
)

func TestGeneralPing(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		statusCode int
		body       string
		headers    map[string]string
		wantErr    bool
		want429    bool
	}{
		{
			name:       "ok",
			statusCode: http.StatusOK,
			body:       `{"TS":"2026-04-22T12:00:00Z","Status":"ok"}`,
			headers:    map[string]string{"Content-Type": "application/json"},
		},
		{
			name:       "bad request",
			statusCode: http.StatusBadRequest,
			body:       `{"title":"bad request"}`,
			headers:    map[string]string{"Content-Type": "application/problem+json", "X-Request-Id": "req-400"},
			wantErr:    true,
		},
		{
			name:       "unauthorized",
			statusCode: http.StatusUnauthorized,
			body:       `{"title":"unauthorized"}`,
			headers:    map[string]string{"Content-Type": "application/problem+json", "X-Request-Id": "req-401"},
			wantErr:    true,
		},
		{
			name:       "rate limited",
			statusCode: http.StatusTooManyRequests,
			body:       `{"title":"too many requests"}`,
			headers: map[string]string{
				"Content-Type":          "application/problem+json",
				"X-Request-Id":          "req-429",
				"X-Ratelimit-Retry":     "1",
				"X-Ratelimit-Limit":     "3",
				"X-Ratelimit-Remaining": "0",
			},
			wantErr: true,
			want429: true,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			var gotAuth string
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				gotAuth = r.Header.Get("Authorization")
				if r.URL.Path != "/ping" {
					t.Fatalf("unexpected path: %s", r.URL.Path)
				}
				for k, v := range tc.headers {
					w.Header().Set(k, v)
				}
				w.WriteHeader(tc.statusCode)
				_, _ = fmt.Fprint(w, tc.body)
			}))
			defer server.Close()

			c := client.NewClient(
				client.WithToken("test-token"),
				client.WithGeneralBaseURL(server.URL),
				client.WithRetryPolicy(transport.RetryPolicy{MaxAttempts: 1, BaseDelay: 10, MaxDelay: 10, Jitter: 0}),
			)

			resp, httpResp, err := c.General().Ping(context.Background())

			if gotAuth != "Bearer test-token" {
				t.Fatalf("auth header mismatch: %q", gotAuth)
			}

			if tc.wantErr {
				if err == nil {
					t.Fatalf("expected error")
				}
				if tc.want429 {
					if !transport.IsRateLimit(err) {
						t.Fatalf("expected rate limit error, got %T", err)
					}
				} else {
					var apiErr *transport.APIError
					if !errors.As(err, &apiErr) {
						t.Fatalf("expected APIError, got %T", err)
					}
				}
				if httpResp == nil {
					t.Fatalf("expected non-nil http response")
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if httpResp == nil || httpResp.StatusCode != http.StatusOK {
				t.Fatalf("unexpected status response: %+v", httpResp)
			}

			testkit.AssertJSONEqual(t, tc.body, resp)
		})
	}
}
