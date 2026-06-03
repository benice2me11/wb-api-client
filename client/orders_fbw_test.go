package client_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/benice2me11/wb-api-client/client"
	"github.com/benice2me11/wb-api-client/transport"
)

func TestOrdersFBWFacadeRoutesAndStatus(t *testing.T) {
	t.Parallel()

	calls := map[string]int{}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if auth := r.Header.Get("Authorization"); auth != "Bearer orders-fbw-token" {
			t.Fatalf("auth header mismatch: %q", auth)
		}

		key := r.Method + " " + r.URL.Path
		calls[key]++
		w.Header().Set("Content-Type", "application/json")

		switch key {
		case "POST /api/v1/acceptance/options":
			assertQueryValue(t, r, "warehouseID", "321")
			assertRequestJSON(t, r, `[{"barcode":"barcode-1","quantity":2}]`)
			w.WriteHeader(http.StatusOK)
		case "GET /api/v1/warehouses",
			"GET /api/v1/transit-tariffs",
			"POST /api/v1/supplies",
			"GET /api/v1/supplies/42/goods",
			"GET /api/v1/supplies/42/package":
			if key == "POST /api/v1/supplies" {
				assertQueryValue(t, r, "limit", "50")
				assertQueryValue(t, r, "offset", "5")
				assertRequestJSON(t, r, `{"dates":[{"from":"2026-06-01","till":"2026-06-02","type":"createDate"}],"statusIDs":[1]}`)
			}
			if key == "GET /api/v1/supplies/42/goods" {
				assertQueryValue(t, r, "limit", "50")
				assertQueryValue(t, r, "offset", "5")
				assertQueryValue(t, r, "isPreorderID", "true")
			}
			w.WriteHeader(http.StatusOK)
		case "GET /api/v1/supplies/42":
			assertQueryValue(t, r, "isPreorderID", "true")
			w.WriteHeader(http.StatusOK)
		default:
			t.Fatalf("unexpected request: %s", key)
		}
	}))
	defer server.Close()

	c := client.NewClient(
		client.WithToken("orders-fbw-token"),
		client.WithOrdersFBWBaseURL(server.URL),
	)

	ctx := context.Background()
	assertStatus := assertFacadeStatus(t)
	warehouseID := int32(321)
	limit := int32(50)
	offset := int32(5)
	isPreorder := true

	good := client.ModelsGood{}
	good.SetBarcode("barcode-1")
	good.SetQuantity(2)

	dateFilter := client.ModelsDateFilterRequest{Type: "createDate"}
	dateFilter.SetFrom("2026-06-01")
	dateFilter.SetTill("2026-06-02")
	filters := client.ModelsSuppliesFiltersRequest{
		Dates:     []client.ModelsDateFilterRequest{dateFilter},
		StatusIDs: []client.ModelsHandySupplyStatus{client.ModelsHandySupplyStatus(1)},
	}

	assertStatus(c.OrdersFBW().AcceptanceOptions(ctx, []client.ModelsGood{good}, &client.OrdersFBWAcceptanceOptionsQuery{WarehouseID: &warehouseID}))
	assertStatus(c.OrdersFBW().Warehouses(ctx))
	assertStatus(c.OrdersFBW().TransitTariffs(ctx))
	assertStatus(c.OrdersFBW().Supplies(ctx, client.OrdersFBWSuppliesQuery{Filters: filters, Limit: &limit, Offset: &offset}))
	assertStatus(c.OrdersFBW().Supply(ctx, 42, &client.OrdersFBWSupplyQuery{IsPreorderID: &isPreorder}))
	assertStatus(c.OrdersFBW().SupplyGoods(ctx, 42, &client.OrdersFBWSupplyGoodsQuery{Limit: &limit, Offset: &offset, IsPreorderID: &isPreorder}))
	assertStatus(c.OrdersFBW().SupplyPackage(ctx, 42))

	expected := []string{
		"POST /api/v1/acceptance/options",
		"GET /api/v1/warehouses",
		"GET /api/v1/transit-tariffs",
		"POST /api/v1/supplies",
		"GET /api/v1/supplies/42",
		"GET /api/v1/supplies/42/goods",
		"GET /api/v1/supplies/42/package",
	}
	for _, key := range expected {
		if calls[key] != 1 {
			t.Fatalf("expected one call for %s, got %d", key, calls[key])
		}
	}
}

func TestOrdersFBWFacadeWrapsErrors(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		status    int
		requestID string
		want429   bool
		call      func(context.Context, client.OrdersFBWService) (*http.Response, error)
	}{
		{
			name:      "bad request",
			status:    http.StatusBadRequest,
			requestID: "orders-fbw-400",
			call: func(ctx context.Context, orders client.OrdersFBWService) (*http.Response, error) {
				_, httpResp, err := orders.Warehouses(ctx)
				return httpResp, err
			},
		},
		{
			name:      "rate limited",
			status:    http.StatusTooManyRequests,
			requestID: "orders-fbw-429",
			want429:   true,
			call: func(ctx context.Context, orders client.OrdersFBWService) (*http.Response, error) {
				_, httpResp, err := orders.AcceptanceOptions(ctx, []client.ModelsGood{{}}, nil)
				return httpResp, err
			},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if auth := r.Header.Get("Authorization"); auth != "Bearer orders-fbw-token" {
					t.Fatalf("auth header mismatch: %q", auth)
				}
				w.Header().Set("Content-Type", "application/problem+json")
				w.Header().Set("X-Request-Id", tc.requestID)
				if tc.want429 {
					w.Header().Set("X-Ratelimit-Retry", "1")
					w.Header().Set("X-Ratelimit-Limit", "10")
					w.Header().Set("X-Ratelimit-Remaining", "0")
				}
				w.WriteHeader(tc.status)
				_, _ = fmt.Fprint(w, `{"title":"api error"}`)
			}))
			defer server.Close()

			c := client.NewClient(
				client.WithToken("orders-fbw-token"),
				client.WithOrdersFBWBaseURL(server.URL),
				client.WithRetryPolicy(transport.RetryPolicy{MaxAttempts: 1, BaseDelay: 10, MaxDelay: 10, Jitter: 0}),
			)

			httpResp, err := tc.call(context.Background(), c.OrdersFBW())
			assertProductsFacadeWrappedError(t, err, tc.requestID, tc.want429)
			if httpResp == nil || httpResp.StatusCode != tc.status {
				t.Fatalf("unexpected error response: %+v", httpResp)
			}
		})
	}
}
