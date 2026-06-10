package client_test

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/benice2me11/wb-api-client/client"
	"github.com/benice2me11/wb-api-client/internal/testkit"
	"github.com/benice2me11/wb-api-client/transport"
)

func TestProductsDeleteInventory(t *testing.T) {
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
			statusCode: http.StatusNoContent,
			headers:    map[string]string{},
		},
		{
			name:       "bad request",
			statusCode: http.StatusBadRequest,
			body:       `{"title":"bad request"}`,
			headers:    map[string]string{"Content-Type": "application/problem+json", "X-Request-Id": "stocks-delete-400"},
			wantErr:    true,
		},
		{
			name:       "unauthorized",
			statusCode: http.StatusUnauthorized,
			body:       `{"title":"unauthorized"}`,
			headers:    map[string]string{"Content-Type": "application/problem+json", "X-Request-Id": "stocks-delete-401"},
			wantErr:    true,
		},
		{
			name:       "rate limited",
			statusCode: http.StatusTooManyRequests,
			body:       `{"title":"too many requests"}`,
			headers: map[string]string{
				"Content-Type":          "application/problem+json",
				"X-Request-Id":          "stocks-delete-429",
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
				if r.URL.Path != "/api/v3/stocks/777" {
					t.Fatalf("unexpected path: %s", r.URL.Path)
				}
				if r.Method != http.MethodDelete {
					t.Fatalf("unexpected method: %s", r.Method)
				}
				if auth := r.Header.Get("Authorization"); auth != "Bearer products-token" {
					t.Fatalf("auth header mismatch: %q", auth)
				}
				assertRequestJSON(t, r, `{"chrtIds":[11,22]}`)

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

			httpResp, err := c.Products().DeleteInventory(
				context.Background(),
				777,
				client.ApiV3StocksWarehouseIdDeleteRequest{ChrtIds: []int32{11, 22}},
			)
			if tc.wantErr {
				assertProductsFacadeWrappedError(t, err, tc.headers["X-Request-Id"], tc.want429)
				if httpResp == nil || httpResp.StatusCode != tc.statusCode {
					t.Fatalf("unexpected error response: %+v", httpResp)
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if httpResp == nil || httpResp.StatusCode != http.StatusNoContent {
				t.Fatalf("unexpected status response: %+v", httpResp)
			}
		})
	}
}

func TestProductsDeleteCards(t *testing.T) {
	t.Parallel()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/content/v2/cards/delete/trash" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		if r.Method != http.MethodPost {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if got := r.Header.Get("Authorization"); got != "Bearer products-token" {
			t.Fatalf("auth header mismatch: %q", got)
		}
		assertRequestJSON(t, r, `{"nmIDs":[11,22]}`)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, `{}`)
	}))
	defer server.Close()

	c := client.NewClient(
		client.WithToken("products-token"),
		client.WithProductsBaseURL(server.URL),
	)

	resp, httpResp, err := c.Products().DeleteCards(context.Background(), client.ContentV2CardsDeleteTrashPostRequest{
		NmIDs: []int32{11, 22},
	})
	if err != nil {
		t.Fatalf("DeleteCards unexpected error: %v", err)
	}
	if httpResp == nil || httpResp.StatusCode != http.StatusOK {
		t.Fatalf("DeleteCards unexpected status: %+v", httpResp)
	}
	if resp == nil {
		t.Fatalf("DeleteCards unexpected response: %+v", resp)
	}
}

func TestProductsResetInventory(t *testing.T) {
	t.Parallel()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/v3/stocks/777" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		if r.Method != http.MethodPut {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if got := r.Header.Get("Authorization"); got != "Bearer products-token" {
			t.Fatalf("auth header mismatch: %q", got)
		}
		assertRequestJSON(t, r, `{"stocks":[{"amount":0,"chrtId":11},{"amount":0,"chrtId":22}]}`)
		w.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	c := client.NewClient(
		client.WithToken("products-token"),
		client.WithProductsBaseURL(server.URL),
	)

	httpResp, err := c.Products().ResetInventory(context.Background(), 777, []int32{11, 22})
	if err != nil {
		t.Fatalf("ResetInventory unexpected error: %v", err)
	}
	if httpResp == nil || httpResp.StatusCode != http.StatusNoContent {
		t.Fatalf("ResetInventory unexpected status: %+v", httpResp)
	}
}

func TestProductsWarehouseLifecycle(t *testing.T) {
	t.Parallel()

	calls := map[string]int{}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if auth := r.Header.Get("Authorization"); auth != "Bearer products-token" {
			t.Fatalf("auth header mismatch: %q", auth)
		}

		key := r.Method + " " + r.URL.Path
		calls[key]++

		switch key {
		case "GET /api/v3/offices":
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprint(w, `[{"id":12,"name":"Office 12","city":"Kazan"}]`)
		case "GET /api/v3/warehouses":
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprint(w, `[{"id":123,"name":"Warehouse 123","officeId":12}]`)
		case "POST /api/v3/warehouses":
			assertRequestJSON(t, r, `{"name":"New Warehouse","officeId":12}`)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			_, _ = fmt.Fprint(w, `{"id":123}`)
		case "PUT /api/v3/warehouses/123":
			assertRequestJSON(t, r, `{"name":"Updated Warehouse","officeId":12}`)
			w.WriteHeader(http.StatusNoContent)
		case "DELETE /api/v3/warehouses/123":
			w.WriteHeader(http.StatusNoContent)
		case "GET /api/v3/dbw/warehouses/123/contacts":
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprint(w, `{"contacts":[{"comment":"ops","phone":"+79998887766"}]}`)
		case "PUT /api/v3/dbw/warehouses/123/contacts":
			assertRequestJSON(t, r, `{"contacts":[{"comment":"ops","phone":"+79998887766"}]}`)
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
	if resp, httpResp, err := c.Products().Offices(ctx); err != nil {
		t.Fatalf("Offices unexpected error: %v", err)
	} else if httpResp == nil || httpResp.StatusCode != http.StatusOK || resp == nil {
		t.Fatalf("Offices unexpected response: %+v %+v", httpResp, resp)
	} else {
		testkit.AssertJSONEqual(t, `[{"id":12,"name":"Office 12","city":"Kazan"}]`, resp)
	}

	if resp, httpResp, err := c.Products().Warehouses(ctx); err != nil {
		t.Fatalf("Warehouses unexpected error: %v", err)
	} else if httpResp == nil || httpResp.StatusCode != http.StatusOK || resp == nil {
		t.Fatalf("Warehouses unexpected response: %+v %+v", httpResp, resp)
	} else {
		testkit.AssertJSONEqual(t, `[{"id":123,"name":"Warehouse 123","officeId":12}]`, resp)
	}

	createResp, createHTTPResp, err := c.Products().CreateWarehouse(ctx, client.ApiV3WarehousesPostRequest{Name: "New Warehouse", OfficeId: 12})
	if err != nil {
		t.Fatalf("CreateWarehouse unexpected error: %v", err)
	}
	if createHTTPResp == nil || createHTTPResp.StatusCode != http.StatusCreated || createResp == nil {
		t.Fatalf("CreateWarehouse unexpected response: %+v %+v", createHTTPResp, createResp)
	}
	testkit.AssertJSONEqual(t, `{"id":123}`, createResp)

	if httpResp, err := c.Products().UpdateWarehouse(ctx, 123, client.ApiV3WarehousesWarehouseIdPutRequest{Name: "Updated Warehouse", OfficeId: 12}); err != nil {
		t.Fatalf("UpdateWarehouse unexpected error: %v", err)
	} else if httpResp == nil || httpResp.StatusCode != http.StatusNoContent {
		t.Fatalf("UpdateWarehouse unexpected status: %+v", httpResp)
	}

	if httpResp, err := c.Products().DeleteWarehouse(ctx, 123); err != nil {
		t.Fatalf("DeleteWarehouse unexpected error: %v", err)
	} else if httpResp == nil || httpResp.StatusCode != http.StatusNoContent {
		t.Fatalf("DeleteWarehouse unexpected status: %+v", httpResp)
	}

	if resp, httpResp, err := c.Products().DBWWarehouseContacts(ctx, 123); err != nil {
		t.Fatalf("DBWWarehouseContacts unexpected error: %v", err)
	} else if httpResp == nil || httpResp.StatusCode != http.StatusOK || resp == nil {
		t.Fatalf("DBWWarehouseContacts unexpected response: %+v %+v", httpResp, resp)
	} else {
		testkit.AssertJSONEqual(t, `{"contacts":[{"comment":"ops","phone":"+79998887766"}]}`, resp)
	}

	contact := client.StoreContactRequestBodyContactsInner{}
	contact.SetComment("ops")
	contact.SetPhone("+79998887766")
	if httpResp, err := c.Products().UpdateDBWWarehouseContacts(ctx, 123, client.StoreContactRequestBody{
		Contacts: []client.StoreContactRequestBodyContactsInner{contact},
	}); err != nil {
		t.Fatalf("UpdateDBWWarehouseContacts unexpected error: %v", err)
	} else if httpResp == nil || httpResp.StatusCode != http.StatusNoContent {
		t.Fatalf("UpdateDBWWarehouseContacts unexpected status: %+v", httpResp)
	}

	expected := []string{
		"GET /api/v3/offices",
		"GET /api/v3/warehouses",
		"POST /api/v3/warehouses",
		"PUT /api/v3/warehouses/123",
		"DELETE /api/v3/warehouses/123",
		"GET /api/v3/dbw/warehouses/123/contacts",
		"PUT /api/v3/dbw/warehouses/123/contacts",
	}
	for _, key := range expected {
		if calls[key] != 1 {
			t.Fatalf("expected one call for %s, got %d", key, calls[key])
		}
	}
}

func TestProductsWarehouseFacadesWrapErrors(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		status    int
		requestID string
		want429   bool
		call      func(context.Context, client.ProductsService) (*http.Response, error)
	}{
		{
			name:      "offices bad request",
			status:    http.StatusBadRequest,
			requestID: "offices-400",
			call: func(ctx context.Context, products client.ProductsService) (*http.Response, error) {
				_, httpResp, err := products.Offices(ctx)
				return httpResp, err
			},
		},
		{
			name:      "create warehouse unauthorized",
			status:    http.StatusUnauthorized,
			requestID: "warehouse-create-401",
			call: func(ctx context.Context, products client.ProductsService) (*http.Response, error) {
				_, httpResp, err := products.CreateWarehouse(ctx, client.ApiV3WarehousesPostRequest{Name: "Warehouse", OfficeId: 12})
				return httpResp, err
			},
		},
		{
			name:      "update warehouse rate limited",
			status:    http.StatusTooManyRequests,
			requestID: "warehouse-update-429",
			want429:   true,
			call: func(ctx context.Context, products client.ProductsService) (*http.Response, error) {
				return products.UpdateWarehouse(ctx, 123, client.ApiV3WarehousesWarehouseIdPutRequest{Name: "Warehouse", OfficeId: 12})
			},
		},
		{
			name:      "dbw contacts bad request",
			status:    http.StatusBadRequest,
			requestID: "dbw-contacts-400",
			call: func(ctx context.Context, products client.ProductsService) (*http.Response, error) {
				_, httpResp, err := products.DBWWarehouseContacts(ctx, 123)
				return httpResp, err
			},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if auth := r.Header.Get("Authorization"); auth != "Bearer products-token" {
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
				client.WithToken("products-token"),
				client.WithProductsBaseURL(server.URL),
				client.WithRetryPolicy(transport.RetryPolicy{MaxAttempts: 1, BaseDelay: 10, MaxDelay: 10, Jitter: 0}),
			)

			httpResp, err := tc.call(context.Background(), c.Products())
			assertProductsFacadeWrappedError(t, err, tc.requestID, tc.want429)
			if httpResp == nil || httpResp.StatusCode != tc.status {
				t.Fatalf("unexpected error response: %+v", httpResp)
			}
		})
	}
}

func assertRequestJSON(t *testing.T, r *http.Request, want string) {
	t.Helper()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		t.Fatalf("read request body: %v", err)
	}

	var got interface{}
	if err := json.Unmarshal(body, &got); err != nil {
		t.Fatalf("decode request body: %v", err)
	}
	testkit.AssertJSONEqual(t, want, got)
}

func assertProductsFacadeWrappedError(t *testing.T, err error, requestID string, wantRateLimit bool) {
	t.Helper()

	if err == nil {
		t.Fatalf("expected error")
	}
	if wantRateLimit {
		var rateErr *transport.RateLimitError
		if !errors.As(err, &rateErr) {
			t.Fatalf("expected RateLimitError, got %T", err)
		}
		if rateErr.Meta.RequestID != requestID {
			t.Fatalf("unexpected request id: %s", rateErr.Meta.RequestID)
		}
		return
	}

	var apiErr *transport.APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("expected APIError, got %T", err)
	}
	if apiErr.RequestID != requestID {
		t.Fatalf("unexpected request id: %s", apiErr.RequestID)
	}
}
