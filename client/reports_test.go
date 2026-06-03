package client_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/benice2me11/wb-api-client/client"
	"github.com/benice2me11/wb-api-client/transport"
)

func TestReportsFacadeRoutesAndStatus(t *testing.T) {
	t.Parallel()

	calls := map[string]int{}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if auth := r.Header.Get("Authorization"); auth != "Bearer reports-token" {
			t.Fatalf("auth header mismatch: %q", auth)
		}

		key := r.Method + " " + r.URL.Path
		calls[key]++
		w.Header().Set("Content-Type", "application/json")

		switch key {
		case "GET /api/v1/warehouse_remains":
			assertQueryValue(t, r, "locale", "ru")
			assertQueryValue(t, r, "groupByBrand", "true")
			assertQueryValue(t, r, "filterPics", "1")
			w.WriteHeader(http.StatusOK)
		case "GET /api/v1/warehouse_remains/tasks/remains-task/status",
			"GET /api/v1/acceptance_report/tasks/acceptance-task/status",
			"GET /api/v1/paid_storage/tasks/storage-task/status":
			w.WriteHeader(http.StatusOK)
		case "GET /api/v1/warehouse_remains/tasks/remains-task/download",
			"GET /api/v1/acceptance_report/tasks/acceptance-task/download",
			"GET /api/v1/paid_storage/tasks/storage-task/download",
			"GET /api/v1/supplier/stocks",
			"GET /api/v1/supplier/orders",
			"GET /api/v1/supplier/sales":
			w.WriteHeader(http.StatusOK)
		case "POST /api/v1/analytics/excise-report":
			assertQueryValue(t, r, "dateFrom", "2026-06-01")
			assertQueryValue(t, r, "dateTo", "2026-06-02")
			assertRequestJSON(t, r, `{"countries":["RU"]}`)
			w.WriteHeader(http.StatusOK)
		case "GET /api/v1/acceptance_report",
			"GET /api/v1/paid_storage",
			"GET /api/v1/analytics/antifraud-details",
			"GET /api/v1/analytics/banned-products/blocked",
			"GET /api/v1/analytics/banned-products/shadowed",
			"GET /api/v1/analytics/brand-share/brands",
			"GET /api/v1/analytics/brand-share",
			"GET /api/v1/analytics/brand-share/parent-subjects",
			"GET /api/v1/analytics/goods-labeling",
			"GET /api/v1/analytics/goods-return",
			"GET /api/v1/analytics/region-sale",
			"GET /api/analytics/v1/deductions",
			"GET /api/analytics/v1/measurement-penalties",
			"GET /api/analytics/v1/warehouse-measurements":
			w.WriteHeader(http.StatusOK)
		default:
			t.Fatalf("unexpected request: %s", key)
		}
	}))
	defer server.Close()

	c := client.NewClient(
		client.WithToken("reports-token"),
		client.WithReportsBaseURL(server.URL),
	)

	ctx := context.Background()
	assertStatus := assertFacadeStatus(t)
	day := time.Date(2026, 6, 1, 0, 0, 0, 0, time.UTC)
	nextDay := day.AddDate(0, 0, 1)
	locale := "ru"
	group := true
	filterPics := int32(1)
	flag := int32(1)
	sort := "nmId"
	order := "asc"
	offset := int32(5)

	assertStatus(c.Reports().CreateWarehouseRemainsReport(ctx, &client.ReportsWarehouseRemainsQuery{
		Locale:       &locale,
		GroupByBrand: &group,
		FilterPics:   &filterPics,
	}))
	assertStatus(c.Reports().WarehouseRemainsReportStatus(ctx, "remains-task"))
	assertStatus(c.Reports().DownloadWarehouseRemainsReport(ctx, "remains-task"))
	assertStatus(c.Reports().SupplierStocks(ctx, client.ReportsSupplierStocksQuery{DateFrom: day}))
	assertStatus(c.Reports().SupplierOrders(ctx, client.ReportsSupplierOrdersQuery{DateFrom: day, Flag: &flag}))
	assertStatus(c.Reports().SupplierSales(ctx, client.ReportsSupplierOrdersQuery{DateFrom: day, Flag: &flag}))
	assertStatus(c.Reports().ExciseReport(ctx, client.ReportsDateRangeQuery{DateFrom: "2026-06-01", DateTo: "2026-06-02"}, client.ExciseReportRequest{Countries: []string{"RU"}}))
	assertStatus(c.Reports().CreateAcceptanceReport(ctx, client.ReportsDateRangeQuery{DateFrom: "2026-06-01", DateTo: "2026-06-02"}))
	assertStatus(c.Reports().AcceptanceReportStatus(ctx, "acceptance-task"))
	assertStatus(c.Reports().DownloadAcceptanceReport(ctx, "acceptance-task"))
	assertStatus(c.Reports().CreatePaidStorageReport(ctx, client.ReportsDateRangeQuery{DateFrom: "2026-06-01", DateTo: "2026-06-02"}))
	assertStatus(c.Reports().PaidStorageReportStatus(ctx, "storage-task"))
	assertStatus(c.Reports().DownloadPaidStorageReport(ctx, "storage-task"))
	assertStatus(c.Reports().AntifraudDetails(ctx, "2026-06-01"))
	assertStatus(c.Reports().BannedProductsBlocked(ctx, &client.ReportsBannedProductsQuery{Sort: &sort, Order: &order}))
	assertStatus(c.Reports().BannedProductsShadowed(ctx, &client.ReportsBannedProductsQuery{Sort: &sort, Order: &order}))
	assertStatus(c.Reports().BrandShareBrands(ctx))
	assertStatus(c.Reports().BrandShare(ctx, client.ReportsBrandShareQuery{ParentID: 123, Brand: "Brand", DateFrom: "2026-06-01", DateTo: "2026-06-02"}))
	assertStatus(c.Reports().BrandShareParentSubjects(ctx, client.ReportsBrandShareParentSubjectsQuery{Brand: "Brand", DateFrom: "2026-06-01", DateTo: "2026-06-02", Locale: &locale}))
	assertStatus(c.Reports().GoodsLabeling(ctx, client.ReportsDateRangeQuery{DateFrom: "2026-06-01", DateTo: "2026-06-02"}))
	assertStatus(c.Reports().GoodsReturn(ctx, client.ReportsDateRangeQuery{DateFrom: "2026-06-01", DateTo: "2026-06-02"}))
	assertStatus(c.Reports().RegionSale(ctx, client.ReportsDateRangeQuery{DateFrom: "2026-06-01", DateTo: "2026-06-02"}))
	assertStatus(c.Reports().Deductions(ctx, client.ReportsDeductionsQuery{DateTo: nextDay, Limit: 10, DateFrom: &day, Sort: &sort, Order: &order, Offset: &offset}))
	assertStatus(c.Reports().MeasurementPenalties(ctx, client.ReportsPagedTimeQuery{DateTo: nextDay, Limit: 10, DateFrom: &day, Offset: &offset}))
	assertStatus(c.Reports().WarehouseMeasurements(ctx, client.ReportsPagedTimeQuery{DateTo: nextDay, Limit: 10, DateFrom: &day, Offset: &offset}))

	expected := []string{
		"GET /api/v1/warehouse_remains",
		"GET /api/v1/warehouse_remains/tasks/remains-task/status",
		"GET /api/v1/warehouse_remains/tasks/remains-task/download",
		"GET /api/v1/supplier/stocks",
		"GET /api/v1/supplier/orders",
		"GET /api/v1/supplier/sales",
		"POST /api/v1/analytics/excise-report",
		"GET /api/v1/acceptance_report",
		"GET /api/v1/acceptance_report/tasks/acceptance-task/status",
		"GET /api/v1/acceptance_report/tasks/acceptance-task/download",
		"GET /api/v1/paid_storage",
		"GET /api/v1/paid_storage/tasks/storage-task/status",
		"GET /api/v1/paid_storage/tasks/storage-task/download",
		"GET /api/v1/analytics/antifraud-details",
		"GET /api/v1/analytics/banned-products/blocked",
		"GET /api/v1/analytics/banned-products/shadowed",
		"GET /api/v1/analytics/brand-share/brands",
		"GET /api/v1/analytics/brand-share",
		"GET /api/v1/analytics/brand-share/parent-subjects",
		"GET /api/v1/analytics/goods-labeling",
		"GET /api/v1/analytics/goods-return",
		"GET /api/v1/analytics/region-sale",
		"GET /api/analytics/v1/deductions",
		"GET /api/analytics/v1/measurement-penalties",
		"GET /api/analytics/v1/warehouse-measurements",
	}
	for _, key := range expected {
		if calls[key] != 1 {
			t.Fatalf("expected one call for %s, got %d", key, calls[key])
		}
	}
}

func TestReportsFacadeWrapsErrors(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		status    int
		requestID string
		want429   bool
		call      func(context.Context, client.ReportsService) (*http.Response, error)
	}{
		{
			name:      "bad request",
			status:    http.StatusBadRequest,
			requestID: "reports-400",
			call: func(ctx context.Context, reports client.ReportsService) (*http.Response, error) {
				_, httpResp, err := reports.CreateWarehouseRemainsReport(ctx, nil)
				return httpResp, err
			},
		},
		{
			name:      "rate limited",
			status:    http.StatusTooManyRequests,
			requestID: "reports-429",
			want429:   true,
			call: func(ctx context.Context, reports client.ReportsService) (*http.Response, error) {
				_, httpResp, err := reports.SupplierStocks(ctx, client.ReportsSupplierStocksQuery{DateFrom: time.Date(2026, 6, 1, 0, 0, 0, 0, time.UTC)})
				return httpResp, err
			},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if auth := r.Header.Get("Authorization"); auth != "Bearer reports-token" {
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
				client.WithToken("reports-token"),
				client.WithReportsBaseURL(server.URL),
				client.WithRetryPolicy(transport.RetryPolicy{MaxAttempts: 1, BaseDelay: 10, MaxDelay: 10, Jitter: 0}),
			)

			httpResp, err := tc.call(context.Background(), c.Reports())
			assertProductsFacadeWrappedError(t, err, tc.requestID, tc.want429)
			if httpResp == nil || httpResp.StatusCode != tc.status {
				t.Fatalf("unexpected error response: %+v", httpResp)
			}
		})
	}
}

func assertFacadeStatus(t *testing.T) func(interface{}, *http.Response, error) {
	t.Helper()

	return func(_ interface{}, httpResp *http.Response, err error) {
		t.Helper()

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if httpResp == nil || httpResp.StatusCode < 200 || httpResp.StatusCode > 299 {
			t.Fatalf("unexpected status response: %+v", httpResp)
		}
	}
}

func assertQueryValue(t *testing.T, r *http.Request, key string, want string) {
	t.Helper()

	if got := r.URL.Query().Get(key); got != want {
		t.Fatalf("unexpected %s query param: %q", key, got)
	}
}
