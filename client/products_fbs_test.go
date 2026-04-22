package client_test

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/benice2me11/wb-api-client/client"
	wbfbs "github.com/benice2me11/wb-api-client/internal/generated/fbs"
	wbproducts "github.com/benice2me11/wb-api-client/internal/generated/products"
	"github.com/benice2me11/wb-api-client/internal/testkit"
	"github.com/benice2me11/wb-api-client/transport"
)

func TestProductsGetCardsList(t *testing.T) {
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
			body:       `{}`,
			headers:    map[string]string{"Content-Type": "application/json"},
		},
		{
			name:       "bad request",
			statusCode: http.StatusBadRequest,
			body:       `{"title":"bad request"}`,
			headers:    map[string]string{"Content-Type": "application/problem+json", "X-Request-Id": "products-400"},
			wantErr:    true,
		},
		{
			name:       "unauthorized",
			statusCode: http.StatusUnauthorized,
			body:       `{"title":"unauthorized"}`,
			headers:    map[string]string{"Content-Type": "application/problem+json", "X-Request-Id": "products-401"},
			wantErr:    true,
		},
		{
			name:       "rate limited",
			statusCode: http.StatusTooManyRequests,
			body:       `{"title":"too many requests"}`,
			headers: map[string]string{
				"Content-Type":          "application/problem+json",
				"X-Request-Id":          "products-429",
				"X-Ratelimit-Retry":     "1",
				"X-Ratelimit-Limit":     "10",
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

			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.URL.Path != "/content/v2/get/cards/list" {
					t.Fatalf("unexpected path: %s", r.URL.Path)
				}
				if r.Method != http.MethodPost {
					t.Fatalf("unexpected method: %s", r.Method)
				}
				if auth := r.Header.Get("Authorization"); auth != "Bearer products-token" {
					t.Fatalf("auth header mismatch: %q", auth)
				}
				_, _ = io.ReadAll(r.Body)
				for k, v := range tc.headers {
					w.Header().Set(k, v)
				}
				w.WriteHeader(tc.statusCode)
				_, _ = fmt.Fprint(w, tc.body)
			}))
			defer server.Close()

			c := client.NewClient(
				client.WithToken("products-token"),
				client.WithProductsBaseURL(server.URL),
				client.WithRetryPolicy(transport.RetryPolicy{MaxAttempts: 1, BaseDelay: 10, MaxDelay: 10, Jitter: 0}),
			)

			resp, httpResp, err := c.Products().GetCardsList(context.Background(), wbproducts.ContentV2GetCardsListPostRequest{})
			if tc.wantErr {
				if err == nil {
					t.Fatalf("expected error")
				}
				if httpResp == nil {
					t.Fatalf("expected non-nil http response")
				}
				if tc.want429 {
					var rateErr *transport.RateLimitError
					if !errors.As(err, &rateErr) {
						t.Fatalf("expected RateLimitError, got %T", err)
					}
					if rateErr.Meta.RequestID != "products-429" {
						t.Fatalf("unexpected request id: %s", rateErr.Meta.RequestID)
					}
				} else {
					var apiErr *transport.APIError
					if !errors.As(err, &apiErr) {
						t.Fatalf("expected APIError, got %T", err)
					}
					if apiErr.RequestID == "" {
						t.Fatalf("expected non-empty request id")
					}
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if httpResp == nil || httpResp.StatusCode != http.StatusOK {
				t.Fatalf("unexpected status response: %+v", httpResp)
			}
			if resp == nil {
				t.Fatalf("expected non-nil response")
			}
			testkit.AssertJSONEqual(t, `{}`, resp)
		})
	}
}

func TestFBSOrders(t *testing.T) {
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
			body:       `{}`,
			headers:    map[string]string{"Content-Type": "application/json"},
		},
		{
			name:       "bad request",
			statusCode: http.StatusBadRequest,
			body:       `{"code":"bad request"}`,
			headers:    map[string]string{"Content-Type": "application/json", "X-Request-Id": "fbs-400"},
			wantErr:    true,
		},
		{
			name:       "unauthorized",
			statusCode: http.StatusUnauthorized,
			body:       `{"code":"unauthorized"}`,
			headers:    map[string]string{"Content-Type": "application/json", "X-Request-Id": "fbs-401"},
			wantErr:    true,
		},
		{
			name:       "rate limited",
			statusCode: http.StatusTooManyRequests,
			body:       `{"code":"too many requests"}`,
			headers: map[string]string{
				"Content-Type":          "application/json",
				"X-Request-Id":          "fbs-429",
				"X-Ratelimit-Retry":     "1",
				"X-Ratelimit-Limit":     "300",
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

			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.URL.Path != "/api/v3/orders" {
					t.Fatalf("unexpected path: %s", r.URL.Path)
				}
				if r.Method != http.MethodGet {
					t.Fatalf("unexpected method: %s", r.Method)
				}
				if got := r.URL.Query().Get("limit"); got != "2" {
					t.Fatalf("unexpected limit query param: %s", got)
				}
				if got := r.URL.Query().Get("next"); got != "0" {
					t.Fatalf("unexpected next query param: %s", got)
				}
				if auth := r.Header.Get("Authorization"); auth != "Bearer fbs-token" {
					t.Fatalf("auth header mismatch: %q", auth)
				}
				for k, v := range tc.headers {
					w.Header().Set(k, v)
				}
				w.WriteHeader(tc.statusCode)
				_, _ = fmt.Fprint(w, tc.body)
			}))
			defer server.Close()

			c := client.NewClient(
				client.WithToken("fbs-token"),
				client.WithFBSBaseURL(server.URL),
				client.WithRetryPolicy(transport.RetryPolicy{MaxAttempts: 1, BaseDelay: 10, MaxDelay: 10, Jitter: 0}),
			)

			limit := int32(2)
			next := int64(0)
			resp, httpResp, err := c.FBS().Orders(context.Background(), &client.FBSOrdersQuery{Limit: &limit, Next: &next})
			if tc.wantErr {
				if err == nil {
					t.Fatalf("expected error")
				}
				if httpResp == nil {
					t.Fatalf("expected non-nil http response")
				}
				if tc.want429 {
					var rateErr *transport.RateLimitError
					if !errors.As(err, &rateErr) {
						t.Fatalf("expected RateLimitError, got %T", err)
					}
					if rateErr.Meta.RequestID != "fbs-429" {
						t.Fatalf("unexpected request id: %s", rateErr.Meta.RequestID)
					}
				} else {
					var apiErr *transport.APIError
					if !errors.As(err, &apiErr) {
						t.Fatalf("expected APIError, got %T", err)
					}
					if apiErr.RequestID == "" {
						t.Fatalf("expected non-empty request id")
					}
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if httpResp == nil || httpResp.StatusCode != http.StatusOK {
				t.Fatalf("unexpected status response: %+v", httpResp)
			}
			if resp == nil {
				t.Fatalf("expected non-nil response")
			}
			testkit.AssertJSONEqual(t, `{}`, resp)
		})
	}
}

func TestProductsOperationsEndpoints(t *testing.T) {
	t.Parallel()

	calls := map[string]int{}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if auth := r.Header.Get("Authorization"); auth != "Bearer products-token" {
			t.Fatalf("auth header mismatch: %q", auth)
		}

		key := r.Method + " " + r.URL.Path
		calls[key]++

		switch key {
		case "POST /content/v2/cards/upload":
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprint(w, `{}`)
		case "POST /content/v2/cards/upload/add":
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprint(w, `{}`)
		case "POST /content/v2/cards/update":
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprint(w, `{}`)
		case "POST /api/v2/upload/task":
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprint(w, `{}`)
		case "POST /api/v2/upload/task/size":
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprint(w, `{}`)
		case "GET /api/v2/history/goods/task":
			if got := r.URL.Query().Get("limit"); got != "1" {
				t.Fatalf("unexpected limit: %s", got)
			}
			if got := r.URL.Query().Get("uploadID"); got != "2" {
				t.Fatalf("unexpected uploadID: %s", got)
			}
			if got := r.URL.Query().Get("offset"); got != "0" {
				t.Fatalf("unexpected offset: %s", got)
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprint(w, `{}`)
		case "POST /api/v3/stocks/10":
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprint(w, `{}`)
		case "PUT /api/v3/stocks/10":
			w.WriteHeader(http.StatusNoContent)
		default:
			t.Fatalf("unexpected request: %s", key)
		}
	}))
	defer server.Close()

	c := client.NewClient(
		client.WithToken("products-token"),
		client.WithProductsBaseURL(server.URL),
	)

	ctx := context.Background()

	if _, httpResp, err := c.Products().CreateCards(ctx, []wbproducts.ContentV2CardsUploadPostRequestInner{}); err != nil {
		t.Fatalf("CreateCards unexpected error: %v", err)
	} else if httpResp == nil || httpResp.StatusCode != http.StatusOK {
		t.Fatalf("CreateCards unexpected status: %+v", httpResp)
	}

	if _, httpResp, err := c.Products().AddCards(ctx, wbproducts.ContentV2CardsUploadAddPostRequest{}); err != nil {
		t.Fatalf("AddCards unexpected error: %v", err)
	} else if httpResp == nil || httpResp.StatusCode != http.StatusOK {
		t.Fatalf("AddCards unexpected status: %+v", httpResp)
	}

	if _, httpResp, err := c.Products().UpdateCards(ctx, []wbproducts.ContentV2CardsUpdatePostRequestInner{}); err != nil {
		t.Fatalf("UpdateCards unexpected error: %v", err)
	} else if httpResp == nil || httpResp.StatusCode != http.StatusOK {
		t.Fatalf("UpdateCards unexpected status: %+v", httpResp)
	}

	if _, httpResp, err := c.Products().SetPrices(ctx, wbproducts.ApiV2UploadTaskPostRequest{Data: []wbproducts.Good{}}); err != nil {
		t.Fatalf("SetPrices unexpected error: %v", err)
	} else if httpResp == nil || httpResp.StatusCode != http.StatusOK {
		t.Fatalf("SetPrices unexpected status: %+v", httpResp)
	}

	if _, httpResp, err := c.Products().SetSizePrices(ctx, wbproducts.ApiV2UploadTaskSizePostRequest{Data: []wbproducts.SizeGoodReq{}}); err != nil {
		t.Fatalf("SetSizePrices unexpected error: %v", err)
	} else if httpResp == nil || httpResp.StatusCode != http.StatusOK {
		t.Fatalf("SetSizePrices unexpected status: %+v", httpResp)
	}

	offset := int32(0)
	if _, httpResp, err := c.Products().PriceTaskDetails(ctx, client.PriceTaskDetailsQuery{Limit: 1, UploadID: 2, Offset: &offset}); err != nil {
		t.Fatalf("PriceTaskDetails unexpected error: %v", err)
	} else if httpResp == nil || httpResp.StatusCode != http.StatusOK {
		t.Fatalf("PriceTaskDetails unexpected status: %+v", httpResp)
	}

	if _, httpResp, err := c.Products().GetInventory(ctx, 10, wbproducts.ApiV3StocksWarehouseIdPostRequest{ChrtIds: []int32{1}}); err != nil {
		t.Fatalf("GetInventory unexpected error: %v", err)
	} else if httpResp == nil || httpResp.StatusCode != http.StatusOK {
		t.Fatalf("GetInventory unexpected status: %+v", httpResp)
	}

	if httpResp, err := c.Products().UpdateInventory(ctx, 10, wbproducts.ApiV3StocksWarehouseIdPutRequest{Stocks: []wbproducts.ApiV3StocksWarehouseIdPutRequestStocksInner{}}); err != nil {
		t.Fatalf("UpdateInventory unexpected error: %v", err)
	} else if httpResp == nil || httpResp.StatusCode != http.StatusNoContent {
		t.Fatalf("UpdateInventory unexpected status: %+v", httpResp)
	}

	expected := []string{
		"POST /content/v2/cards/upload",
		"POST /content/v2/cards/upload/add",
		"POST /content/v2/cards/update",
		"POST /api/v2/upload/task",
		"POST /api/v2/upload/task/size",
		"GET /api/v2/history/goods/task",
		"POST /api/v3/stocks/10",
		"PUT /api/v3/stocks/10",
	}
	for _, key := range expected {
		if calls[key] != 1 {
			t.Fatalf("expected one call for %s, got %d", key, calls[key])
		}
	}
}

func TestFBSAccountingEndpoints(t *testing.T) {
	t.Parallel()

	calls := map[string]int{}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if auth := r.Header.Get("Authorization"); auth != "Bearer fbs-token" {
			t.Fatalf("auth header mismatch: %q", auth)
		}

		key := r.Method + " " + r.URL.Path
		calls[key]++

		switch key {
		case "POST /api/v3/orders/status":
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprint(w, `{}`)
		case "POST /api/v3/orders/status/history":
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprint(w, `{}`)
		case "PATCH /api/v3/orders/123/cancel":
			w.WriteHeader(http.StatusNoContent)
		default:
			t.Fatalf("unexpected request: %s", key)
		}
	}))
	defer server.Close()

	c := client.NewClient(
		client.WithToken("fbs-token"),
		client.WithFBSBaseURL(server.URL),
	)

	ctx := context.Background()

	if _, httpResp, err := c.FBS().OrdersByStatus(ctx, wbfbs.ApiV3OrdersStatusPostRequest{Orders: []int64{123}}); err != nil {
		t.Fatalf("OrdersByStatus unexpected error: %v", err)
	} else if httpResp == nil || httpResp.StatusCode != http.StatusOK {
		t.Fatalf("OrdersByStatus unexpected status: %+v", httpResp)
	}

	if _, httpResp, err := c.FBS().OrdersStatusHistory(ctx, wbfbs.ApiV3OrdersStatusHistoryPostRequest{Orders: []int32{123}}); err != nil {
		t.Fatalf("OrdersStatusHistory unexpected error: %v", err)
	} else if httpResp == nil || httpResp.StatusCode != http.StatusOK {
		t.Fatalf("OrdersStatusHistory unexpected status: %+v", httpResp)
	}

	if httpResp, err := c.FBS().CancelOrder(ctx, 123); err != nil {
		t.Fatalf("CancelOrder unexpected error: %v", err)
	} else if httpResp == nil || httpResp.StatusCode != http.StatusNoContent {
		t.Fatalf("CancelOrder unexpected status: %+v", httpResp)
	}

	expected := []string{
		"POST /api/v3/orders/status",
		"POST /api/v3/orders/status/history",
		"PATCH /api/v3/orders/123/cancel",
	}
	for _, key := range expected {
		if calls[key] != 1 {
			t.Fatalf("expected one call for %s, got %d", key, calls[key])
		}
	}
}
