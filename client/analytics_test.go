package client_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/benice2me11/wb-api-client/client"
	"github.com/benice2me11/wb-api-client/transport"
)

func TestAnalyticsFacadeRoutesAndStatus(t *testing.T) {
	t.Parallel()

	calls := map[string]int{}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if auth := r.Header.Get("Authorization"); auth != "Bearer analytics-token" {
			t.Fatalf("auth header mismatch: %q", auth)
		}

		key := r.Method + " " + r.URL.Path
		calls[key]++

		switch key {
		case "GET /api/v2/nm-report/downloads":
			got := r.URL.Query()["filter[downloadIds]"]
			if len(got) != 2 || got[0] != "csv-1" || got[1] != "csv-2" {
				t.Fatalf("unexpected filter[downloadIds] query: %#v", got)
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
		case "GET /api/v2/nm-report/downloads/file/csv-1":
			w.Header().Set("Content-Type", "application/zip")
			w.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprint(w, "zip-bytes")
		case "POST /api/v2/nm-report/downloads",
			"POST /api/v2/nm-report/downloads/retry",
			"POST /api/analytics/v3/sales-funnel/products",
			"POST /api/analytics/v3/sales-funnel/grouped/history",
			"POST /api/v2/search-report/report",
			"POST /api/v2/search-report/table/groups",
			"POST /api/v2/search-report/table/details",
			"POST /api/v2/search-report/product/search-texts",
			"POST /api/v2/search-report/product/orders",
			"POST /api/analytics/v1/stocks-report/wb-warehouses",
			"POST /api/v2/stocks-report/products/groups",
			"POST /api/v2/stocks-report/products/products",
			"POST /api/v2/stocks-report/products/sizes",
			"POST /api/v2/stocks-report/offices",
			"POST /api/analytics/v1/item-rating":
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
		case "POST /api/analytics/v3/sales-funnel/products/history":
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
		default:
			t.Fatalf("unexpected request: %s", key)
		}
	}))
	defer server.Close()

	c := client.NewClient(
		client.WithToken("analytics-token"),
		client.WithAnalyticsBaseURL(server.URL),
	)

	ctx := context.Background()
	assertStatus := assertFacadeStatus(t)
	limit := int32(30)
	searchTextsRequest := client.ProductSearchTextsRequest{
		CurrentPeriod: client.Period{Start: "2026-06-01", End: "2026-06-02"},
		NmIds:         []int32{123456},
		TopOrderBy:    "openCard",
		OrderBy:       client.OrderByGrTe{Field: "avgPosition", Mode: "asc"},
		Limit:         client.Int32AsTextLimit(&limit),
	}
	reportReq := client.InventoryHistoryReportReqAsApiV2NmReportDownloadsPostRequest(&client.InventoryHistoryReportReq{
		Id:         "00000000-0000-0000-0000-000000000001",
		ReportType: "STOCK_HISTORY_DAILY_CSV",
		Params:     client.InventoryHistoryReportReqParams{},
	})

	assertStatus(c.Analytics().CSVReports(ctx, []string{"csv-1", "csv-2"}))
	assertStatus(c.Analytics().CreateCSVReport(ctx, reportReq))
	assertStatus(c.Analytics().RetryCSVReport(ctx, client.NmReportRetryReportRequest{}))

	file, httpResp, err := c.Analytics().DownloadCSVReport(ctx, "csv-1")
	if err != nil {
		t.Fatalf("DownloadCSVReport unexpected error: %v", err)
	}
	if httpResp == nil || httpResp.StatusCode != http.StatusOK {
		t.Fatalf("DownloadCSVReport unexpected status: %+v", httpResp)
	}
	if file == nil {
		t.Fatalf("DownloadCSVReport expected file")
	}
	name := file.Name()
	_ = file.Close()
	_ = os.Remove(name)

	assertStatus(c.Analytics().SalesFunnelProducts(ctx, client.ProductsRequest{}))
	assertStatus(c.Analytics().SalesFunnelProductsHistory(ctx, client.ProductHistoryRequest{}))
	assertStatus(c.Analytics().SalesFunnelGroupedHistory(ctx, client.GroupedHistoryRequest{}))
	assertStatus(c.Analytics().SearchReport(ctx, client.MainRequest{}))
	assertStatus(c.Analytics().SearchReportTableGroups(ctx, client.TableGroupRequest{}))
	assertStatus(c.Analytics().SearchReportTableDetails(ctx, client.TableDetailsRequest{}))
	assertStatus(c.Analytics().SearchReportProductSearchTexts(ctx, searchTextsRequest))
	assertStatus(c.Analytics().SearchReportProductOrders(ctx, client.ProductOrdersRequest{}))
	assertStatus(c.Analytics().StocksReportWBWarehouses(ctx, client.InventoryRequest{}))
	assertStatus(c.Analytics().StocksReportProductGroups(ctx, client.TableGroupRequestSt{}))
	assertStatus(c.Analytics().StocksReportProducts(ctx, client.TableProductRequest{}))
	assertStatus(c.Analytics().StocksReportProductSizes(ctx, client.CommonSizeFilters{}))
	assertStatus(c.Analytics().StocksReportOffices(ctx, client.CommonShippingOfficeFilters{}))
	assertStatus(c.Analytics().ItemRating(ctx, client.ItemRatingRequest{}))

	expected := []string{
		"GET /api/v2/nm-report/downloads",
		"POST /api/v2/nm-report/downloads",
		"POST /api/v2/nm-report/downloads/retry",
		"GET /api/v2/nm-report/downloads/file/csv-1",
		"POST /api/analytics/v3/sales-funnel/products",
		"POST /api/analytics/v3/sales-funnel/products/history",
		"POST /api/analytics/v3/sales-funnel/grouped/history",
		"POST /api/v2/search-report/report",
		"POST /api/v2/search-report/table/groups",
		"POST /api/v2/search-report/table/details",
		"POST /api/v2/search-report/product/search-texts",
		"POST /api/v2/search-report/product/orders",
		"POST /api/analytics/v1/stocks-report/wb-warehouses",
		"POST /api/v2/stocks-report/products/groups",
		"POST /api/v2/stocks-report/products/products",
		"POST /api/v2/stocks-report/products/sizes",
		"POST /api/v2/stocks-report/offices",
		"POST /api/analytics/v1/item-rating",
	}
	for _, key := range expected {
		if calls[key] != 1 {
			t.Fatalf("expected one call for %s, got %d", key, calls[key])
		}
	}
}

func TestAnalyticsFacadeWrapsErrors(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		status    int
		requestID string
		want429   bool
		call      func(context.Context, client.AnalyticsService) (*http.Response, error)
	}{
		{
			name:      "bad request",
			status:    http.StatusBadRequest,
			requestID: "analytics-400",
			call: func(ctx context.Context, analytics client.AnalyticsService) (*http.Response, error) {
				_, httpResp, err := analytics.StocksReportProductGroups(ctx, client.TableGroupRequestSt{})
				return httpResp, err
			},
		},
		{
			name:      "rate limited",
			status:    http.StatusTooManyRequests,
			requestID: "analytics-429",
			want429:   true,
			call: func(ctx context.Context, analytics client.AnalyticsService) (*http.Response, error) {
				_, httpResp, err := analytics.CSVReports(ctx, nil)
				return httpResp, err
			},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if auth := r.Header.Get("Authorization"); auth != "Bearer analytics-token" {
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
				client.WithToken("analytics-token"),
				client.WithAnalyticsBaseURL(server.URL),
				client.WithRetryPolicy(transport.RetryPolicy{MaxAttempts: 1, BaseDelay: 10, MaxDelay: 10, Jitter: 0}),
			)

			httpResp, err := tc.call(context.Background(), c.Analytics())
			assertProductsFacadeWrappedError(t, err, tc.requestID, tc.want429)
			if httpResp == nil || httpResp.StatusCode != tc.status {
				t.Fatalf("unexpected error response: %+v", httpResp)
			}
		})
	}
}
