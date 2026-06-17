package client_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/benice2me11/wb-api-client/client"
)

func TestDBWLifecycle(t *testing.T) {
	t.Parallel()

	calls := map[string]int{}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if auth := r.Header.Get("Authorization"); auth != "Bearer dbw-token" {
			t.Fatalf("auth header mismatch: %q", auth)
		}

		key := r.Method + " " + r.URL.Path
		calls[key]++

		switch key {
		case "GET /api/v3/dbw/orders":
			if got := r.URL.Query().Get("limit"); got != "2" {
				t.Fatalf("unexpected limit query param: %s", got)
			}
			if got := r.URL.Query().Get("next"); got != "0" {
				t.Fatalf("unexpected next query param: %s", got)
			}
			if got := r.URL.Query().Get("dateFrom"); got != "1" {
				t.Fatalf("unexpected dateFrom query param: %s", got)
			}
			if got := r.URL.Query().Get("dateTo"); got != "2" {
				t.Fatalf("unexpected dateTo query param: %s", got)
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprint(w, `{}`)
		case "GET /api/v3/dbw/orders/new":
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprint(w, `{}`)
		case "POST /api/v3/dbw/orders/status":
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprint(w, `{}`)
		case "PATCH /api/v3/dbw/orders/123/confirm":
			w.WriteHeader(http.StatusNoContent)
		case "PATCH /api/v3/dbw/orders/123/assemble":
			w.WriteHeader(http.StatusNoContent)
		case "PATCH /api/v3/dbw/orders/123/cancel":
			w.WriteHeader(http.StatusNoContent)
		default:
			t.Fatalf("unexpected request: %s", key)
		}
	}))
	defer server.Close()

	c := client.NewClient(
		client.WithToken("dbw-token"),
		client.WithDBWBaseURL(server.URL),
	)

	ctx := context.Background()
	if resp, httpResp, err := c.DBW().Orders(ctx, client.DBWOrdersQuery{Limit: 2, Next: 0, DateFrom: 1, DateTo: 2}); err != nil {
		t.Fatalf("Orders unexpected error: %v", err)
	} else if httpResp == nil || httpResp.StatusCode != http.StatusOK || resp == nil {
		t.Fatalf("Orders unexpected response: %+v %+v", httpResp, resp)
	}

	if resp, httpResp, err := c.DBW().NewOrders(ctx); err != nil {
		t.Fatalf("NewOrders unexpected error: %v", err)
	} else if httpResp == nil || httpResp.StatusCode != http.StatusOK || resp == nil {
		t.Fatalf("NewOrders unexpected response: %+v %+v", httpResp, resp)
	}

	if resp, httpResp, err := c.DBW().OrdersByStatus(ctx, client.ApiV3DbwOrdersStatusPostRequest{Orders: []int64{123}}); err != nil {
		t.Fatalf("OrdersByStatus unexpected error: %v", err)
	} else if httpResp == nil || httpResp.StatusCode != http.StatusOK || resp == nil {
		t.Fatalf("OrdersByStatus unexpected response: %+v %+v", httpResp, resp)
	}

	if httpResp, err := c.DBW().ConfirmOrder(ctx, 123); err != nil {
		t.Fatalf("ConfirmOrder unexpected error: %v", err)
	} else if httpResp == nil || httpResp.StatusCode != http.StatusNoContent {
		t.Fatalf("ConfirmOrder unexpected status: %+v", httpResp)
	}

	if httpResp, err := c.DBW().AssembleOrder(ctx, 123); err != nil {
		t.Fatalf("AssembleOrder unexpected error: %v", err)
	} else if httpResp == nil || httpResp.StatusCode != http.StatusNoContent {
		t.Fatalf("AssembleOrder unexpected status: %+v", httpResp)
	}

	if httpResp, err := c.DBW().CancelOrder(ctx, 123); err != nil {
		t.Fatalf("CancelOrder unexpected error: %v", err)
	} else if httpResp == nil || httpResp.StatusCode != http.StatusNoContent {
		t.Fatalf("CancelOrder unexpected status: %+v", httpResp)
	}

	expected := []string{
		"GET /api/v3/dbw/orders",
		"GET /api/v3/dbw/orders/new",
		"POST /api/v3/dbw/orders/status",
		"PATCH /api/v3/dbw/orders/123/confirm",
		"PATCH /api/v3/dbw/orders/123/assemble",
		"PATCH /api/v3/dbw/orders/123/cancel",
	}
	for _, key := range expected {
		if calls[key] != 1 {
			t.Fatalf("expected one call for %s, got %d", key, calls[key])
		}
	}
}

func TestDBSLifecycle(t *testing.T) {
	t.Parallel()

	calls := map[string]int{}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if auth := r.Header.Get("Authorization"); auth != "Bearer dbs-token" {
			t.Fatalf("auth header mismatch: %q", auth)
		}

		key := r.Method + " " + r.URL.Path
		calls[key]++

		switch key {
		case "GET /api/v3/dbs/orders":
			if got := r.URL.Query().Get("limit"); got != "2" {
				t.Fatalf("unexpected limit query param: %s", got)
			}
			if got := r.URL.Query().Get("next"); got != "0" {
				t.Fatalf("unexpected next query param: %s", got)
			}
			if got := r.URL.Query().Get("dateFrom"); got != "1" {
				t.Fatalf("unexpected dateFrom query param: %s", got)
			}
			if got := r.URL.Query().Get("dateTo"); got != "2" {
				t.Fatalf("unexpected dateTo query param: %s", got)
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprint(w, `{}`)
		case "GET /api/v3/dbs/orders/new":
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprint(w, `{}`)
		case "POST /api/marketplace/v3/dbs/orders/status/info":
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprint(w, `{}`)
		case "POST /api/marketplace/v3/dbs/orders/status/confirm":
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprint(w, `{"results":[{"orderId":321,"isError":false}]}`)
		case "POST /api/marketplace/v3/dbs/orders/status/deliver":
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprint(w, `{"results":[{"orderId":321,"isError":false}]}`)
		case "POST /api/marketplace/v3/dbs/orders/status/receive":
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprint(w, `{"results":[{"orderId":321,"isError":false}]}`)
		case "POST /api/marketplace/v3/dbs/orders/status/reject":
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprint(w, `{"results":[{"orderId":321,"isError":false}]}`)
		case "POST /api/marketplace/v3/dbs/orders/status/cancel":
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprint(w, `{"results":[{"orderId":321,"isError":false}]}`)
		default:
			t.Fatalf("unexpected request: %s", key)
		}
	}))
	defer server.Close()

	c := client.NewClient(
		client.WithToken("dbs-token"),
		client.WithDBSBaseURL(server.URL),
	)

	ctx := context.Background()
	if resp, httpResp, err := c.DBS().Orders(ctx, client.DBSOrdersQuery{Limit: 2, Next: 0, DateFrom: 1, DateTo: 2}); err != nil {
		t.Fatalf("Orders unexpected error: %v", err)
	} else if httpResp == nil || httpResp.StatusCode != http.StatusOK || resp == nil {
		t.Fatalf("Orders unexpected response: %+v %+v", httpResp, resp)
	}

	if resp, httpResp, err := c.DBS().NewOrders(ctx); err != nil {
		t.Fatalf("NewOrders unexpected error: %v", err)
	} else if httpResp == nil || httpResp.StatusCode != http.StatusOK || resp == nil {
		t.Fatalf("NewOrders unexpected response: %+v %+v", httpResp, resp)
	}

	ordersReq := client.ApiOrdersRequestV2{OrdersIds: []int32{321}}
	if resp, httpResp, err := c.DBS().OrdersStatusInfo(ctx, ordersReq); err != nil {
		t.Fatalf("OrdersStatusInfo unexpected error: %v", err)
	} else if httpResp == nil || httpResp.StatusCode != http.StatusOK || resp == nil {
		t.Fatalf("OrdersStatusInfo unexpected response: %+v %+v", httpResp, resp)
	}

	codeReq := client.ApiOrdersCodeRequest{Orders: []client.ApiOrderCodeRequest{{}}}
	codeReq.Orders[0].SetOrderId(321)
	codeReq.Orders[0].SetCode("1234")

	if resp, httpResp, err := c.DBS().ConfirmOrders(ctx, ordersReq); err != nil {
		t.Fatalf("ConfirmOrders unexpected error: %v", err)
	} else if httpResp == nil || httpResp.StatusCode != http.StatusOK || resp == nil {
		t.Fatalf("ConfirmOrders unexpected response: %+v %+v", httpResp, resp)
	}

	if resp, httpResp, err := c.DBS().DeliverOrders(ctx, ordersReq); err != nil {
		t.Fatalf("DeliverOrders unexpected error: %v", err)
	} else if httpResp == nil || httpResp.StatusCode != http.StatusOK || resp == nil {
		t.Fatalf("DeliverOrders unexpected response: %+v %+v", httpResp, resp)
	}

	if resp, httpResp, err := c.DBS().ReceiveOrders(ctx, codeReq); err != nil {
		t.Fatalf("ReceiveOrders unexpected error: %v", err)
	} else if httpResp == nil || httpResp.StatusCode != http.StatusOK || resp == nil {
		t.Fatalf("ReceiveOrders unexpected response: %+v %+v", httpResp, resp)
	}

	if resp, httpResp, err := c.DBS().RejectOrders(ctx, codeReq); err != nil {
		t.Fatalf("RejectOrders unexpected error: %v", err)
	} else if httpResp == nil || httpResp.StatusCode != http.StatusOK || resp == nil {
		t.Fatalf("RejectOrders unexpected response: %+v %+v", httpResp, resp)
	}

	if resp, httpResp, err := c.DBS().CancelOrders(ctx, ordersReq); err != nil {
		t.Fatalf("CancelOrders unexpected error: %v", err)
	} else if httpResp == nil || httpResp.StatusCode != http.StatusOK || resp == nil {
		t.Fatalf("CancelOrders unexpected response: %+v %+v", httpResp, resp)
	}

	if httpResp, err := c.DBS().ConfirmOrder(ctx, 321); err != nil {
		t.Fatalf("ConfirmOrder unexpected error: %v", err)
	} else if httpResp == nil || httpResp.StatusCode != http.StatusOK {
		t.Fatalf("ConfirmOrder unexpected status: %+v", httpResp)
	}

	if httpResp, err := c.DBS().DeliverOrder(ctx, 321); err != nil {
		t.Fatalf("DeliverOrder unexpected error: %v", err)
	} else if httpResp == nil || httpResp.StatusCode != http.StatusOK {
		t.Fatalf("DeliverOrder unexpected status: %+v", httpResp)
	}

	code := client.NewDBSCode("1234")

	if httpResp, err := c.DBS().ReceiveOrder(ctx, 321, code); err != nil {
		t.Fatalf("ReceiveOrder unexpected error: %v", err)
	} else if httpResp == nil || httpResp.StatusCode != http.StatusOK {
		t.Fatalf("ReceiveOrder unexpected status: %+v", httpResp)
	}

	if httpResp, err := c.DBS().RejectOrder(ctx, 321, code); err != nil {
		t.Fatalf("RejectOrder unexpected error: %v", err)
	} else if httpResp == nil || httpResp.StatusCode != http.StatusOK {
		t.Fatalf("RejectOrder unexpected status: %+v", httpResp)
	}

	if httpResp, err := c.DBS().CancelOrder(ctx, 321); err != nil {
		t.Fatalf("CancelOrder unexpected error: %v", err)
	} else if httpResp == nil || httpResp.StatusCode != http.StatusOK {
		t.Fatalf("CancelOrder unexpected status: %+v", httpResp)
	}

	expected := map[string]int{
		"GET /api/v3/dbs/orders":                             1,
		"GET /api/v3/dbs/orders/new":                         1,
		"POST /api/marketplace/v3/dbs/orders/status/info":    1,
		"POST /api/marketplace/v3/dbs/orders/status/confirm": 2,
		"POST /api/marketplace/v3/dbs/orders/status/deliver": 2,
		"POST /api/marketplace/v3/dbs/orders/status/receive": 2,
		"POST /api/marketplace/v3/dbs/orders/status/reject":  2,
		"POST /api/marketplace/v3/dbs/orders/status/cancel":  2,
	}
	for key, want := range expected {
		if calls[key] != want {
			t.Fatalf("expected %d call(s) for %s, got %d", want, key, calls[key])
		}
	}
}
