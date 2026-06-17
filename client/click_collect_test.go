package client_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/benice2me11/wb-api-client/client"
)

func TestClickCollectFacade(t *testing.T) {
	t.Parallel()

	calls := map[string]int{}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if auth := r.Header.Get("Authorization"); auth != "Bearer cc-token" {
			t.Fatalf("auth header mismatch: %q", auth)
		}

		key := r.Method + " " + r.URL.Path
		calls[key]++
		switch key {
		case "GET /api/v3/click-collect/orders/new":
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprint(w, `{}`)
		case "GET /api/v3/click-collect/orders":
			if got := r.URL.Query().Get("limit"); got != "100" {
				t.Fatalf("unexpected orders limit query param: %s", got)
			}
			if got := r.URL.Query().Get("next"); got != "0" {
				t.Fatalf("unexpected orders next query param: %s", got)
			}
			if got := r.URL.Query().Get("dateFrom"); got != "1700000000" {
				t.Fatalf("unexpected orders dateFrom query param: %s", got)
			}
			if got := r.URL.Query().Get("dateTo"); got != "1700003600" {
				t.Fatalf("unexpected orders dateTo query param: %s", got)
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprint(w, `{}`)
		case "POST /api/v3/click-collect/orders/client":
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprint(w, `{}`)
		case "POST /api/v3/click-collect/orders/client/identity":
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprint(w, `{"ok":true}`)
		case "POST /api/marketplace/v3/click-collect/orders/status/info":
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprint(w, `{}`)
		case "POST /api/marketplace/v3/click-collect/orders/status/confirm":
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprint(w, `{"requestId":"req","results":[]}`)
		case "POST /api/marketplace/v3/click-collect/orders/status/prepare":
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprint(w, `{"requestId":"req","results":[]}`)
		case "POST /api/marketplace/v3/click-collect/orders/status/receive":
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprint(w, `{"requestId":"req","results":[]}`)
		case "POST /api/marketplace/v3/click-collect/orders/status/reject":
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprint(w, `{"requestId":"req","results":[]}`)
		case "POST /api/marketplace/v3/click-collect/orders/status/cancel":
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprint(w, `{"requestId":"req","results":[]}`)
		case "POST /api/marketplace/v3/click-collect/orders/meta/info":
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprint(w, `{"meta":[]}`)
		case "POST /api/marketplace/v3/click-collect/orders/meta/details":
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprint(w, `{"requestId":"req","orders":[]}`)
		case "POST /api/marketplace/v3/click-collect/orders/meta/delete":
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprint(w, `{"requestId":"req","results":[]}`)
		case "POST /api/marketplace/v3/click-collect/orders/meta/sgtin":
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprint(w, `{"requestId":"req","results":[]}`)
		case "POST /api/marketplace/v3/click-collect/orders/meta/uin":
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprint(w, `{"requestId":"req","results":[]}`)
		case "POST /api/marketplace/v3/click-collect/orders/meta/imei":
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprint(w, `{"requestId":"req","results":[]}`)
		case "POST /api/marketplace/v3/click-collect/orders/meta/gtin":
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprint(w, `{"requestId":"req","results":[]}`)
		default:
			t.Fatalf("unexpected request: %s", key)
		}
	}))
	defer server.Close()

	c := client.NewClient(
		client.WithToken("cc-token"),
		client.WithClickCollectBaseURL(server.URL),
	)

	resp, httpResp, err := c.ClickCollect().NewOrders(context.Background())
	if err != nil {
		t.Fatalf("NewOrders unexpected error: %v", err)
	}
	if httpResp == nil || httpResp.StatusCode != http.StatusOK || resp == nil {
		t.Fatalf("NewOrders unexpected response: %+v %+v", httpResp, resp)
	}

	ordersReq := client.ClickCollectApiOrdersRequest{Orders: []int32{123}}
	ordersV2Req := client.ClickCollectApiOrdersRequestV2{OrdersIds: []int32{123}}
	ordersQuery := client.ClickCollectOrdersQuery{
		Limit:    100,
		Next:     0,
		DateFrom: 1700000000,
		DateTo:   1700003600,
	}

	if resp, httpResp, err := c.ClickCollect().Orders(context.Background(), ordersQuery); err != nil {
		t.Fatalf("Orders unexpected error: %v", err)
	} else if httpResp == nil || httpResp.StatusCode != http.StatusOK || resp == nil {
		t.Fatalf("Orders unexpected response: %+v %+v", httpResp, resp)
	}

	if resp, httpResp, err := c.ClickCollect().BuyerInfo(context.Background(), ordersReq); err != nil {
		t.Fatalf("BuyerInfo unexpected error: %v", err)
	} else if httpResp == nil || httpResp.StatusCode != http.StatusOK || resp == nil {
		t.Fatalf("BuyerInfo unexpected response: %+v %+v", httpResp, resp)
	}

	identityReq := client.ClickCollectApiCheckIdentityRequest{}
	identityReq.SetOrderCode("order-code")
	identityReq.SetPasscode("1234")
	if resp, httpResp, err := c.ClickCollect().CheckBuyerIdentity(context.Background(), identityReq); err != nil {
		t.Fatalf("CheckBuyerIdentity unexpected error: %v", err)
	} else if httpResp == nil || httpResp.StatusCode != http.StatusOK || resp == nil {
		t.Fatalf("CheckBuyerIdentity unexpected response: %+v %+v", httpResp, resp)
	}

	if resp, httpResp, err := c.ClickCollect().OrdersStatusInfo(context.Background(), ordersV2Req); err != nil {
		t.Fatalf("OrdersStatusInfo unexpected error: %v", err)
	} else if httpResp == nil || httpResp.StatusCode != http.StatusOK || resp == nil {
		t.Fatalf("OrdersStatusInfo unexpected response: %+v %+v", httpResp, resp)
	}

	if resp, httpResp, err := c.ClickCollect().ConfirmOrders(context.Background(), ordersV2Req); err != nil {
		t.Fatalf("ConfirmOrders unexpected error: %v", err)
	} else if httpResp == nil || httpResp.StatusCode != http.StatusOK || resp == nil {
		t.Fatalf("ConfirmOrders unexpected response: %+v %+v", httpResp, resp)
	}

	if resp, httpResp, err := c.ClickCollect().PrepareOrders(context.Background(), ordersV2Req); err != nil {
		t.Fatalf("PrepareOrders unexpected error: %v", err)
	} else if httpResp == nil || httpResp.StatusCode != http.StatusOK || resp == nil {
		t.Fatalf("PrepareOrders unexpected response: %+v %+v", httpResp, resp)
	}

	if resp, httpResp, err := c.ClickCollect().ReceiveOrders(context.Background(), ordersV2Req); err != nil {
		t.Fatalf("ReceiveOrders unexpected error: %v", err)
	} else if httpResp == nil || httpResp.StatusCode != http.StatusOK || resp == nil {
		t.Fatalf("ReceiveOrders unexpected response: %+v %+v", httpResp, resp)
	}

	if resp, httpResp, err := c.ClickCollect().RejectOrders(context.Background(), ordersV2Req); err != nil {
		t.Fatalf("RejectOrders unexpected error: %v", err)
	} else if httpResp == nil || httpResp.StatusCode != http.StatusOK || resp == nil {
		t.Fatalf("RejectOrders unexpected response: %+v %+v", httpResp, resp)
	}

	if resp, httpResp, err := c.ClickCollect().CancelOrders(context.Background(), ordersV2Req); err != nil {
		t.Fatalf("CancelOrders unexpected error: %v", err)
	} else if httpResp == nil || httpResp.StatusCode != http.StatusOK || resp == nil {
		t.Fatalf("CancelOrders unexpected response: %+v %+v", httpResp, resp)
	}

	if resp, httpResp, err := c.ClickCollect().MetadataInfo(context.Background(), ordersV2Req); err != nil {
		t.Fatalf("MetadataInfo unexpected error: %v", err)
	} else if httpResp == nil || httpResp.StatusCode != http.StatusOK || resp == nil {
		t.Fatalf("MetadataInfo unexpected response: %+v %+v", httpResp, resp)
	}

	if resp, httpResp, err := c.ClickCollect().MetadataDetails(context.Background(), ordersV2Req); err != nil {
		t.Fatalf("MetadataDetails unexpected error: %v", err)
	} else if httpResp == nil || httpResp.StatusCode != http.StatusOK || resp == nil {
		t.Fatalf("MetadataDetails unexpected response: %+v %+v", httpResp, resp)
	}

	if resp, httpResp, err := c.ClickCollect().DeleteMetadata(context.Background(), client.ClickCollectApiOrdersMetaDeleteRequest{}); err != nil {
		t.Fatalf("DeleteMetadata unexpected error: %v", err)
	} else if httpResp == nil || httpResp.StatusCode != http.StatusOK || resp == nil {
		t.Fatalf("DeleteMetadata unexpected response: %+v %+v", httpResp, resp)
	}

	if resp, httpResp, err := c.ClickCollect().SetSGTINMetadata(context.Background(), client.ClickCollectApiOrdersSGTINsSetRequest{}); err != nil {
		t.Fatalf("SetSGTINMetadata unexpected error: %v", err)
	} else if httpResp == nil || httpResp.StatusCode != http.StatusOK || resp == nil {
		t.Fatalf("SetSGTINMetadata unexpected response: %+v %+v", httpResp, resp)
	}

	if resp, httpResp, err := c.ClickCollect().SetUINMetadata(context.Background(), client.ClickCollectApiOrdersUINSetRequest{}); err != nil {
		t.Fatalf("SetUINMetadata unexpected error: %v", err)
	} else if httpResp == nil || httpResp.StatusCode != http.StatusOK || resp == nil {
		t.Fatalf("SetUINMetadata unexpected response: %+v %+v", httpResp, resp)
	}

	if resp, httpResp, err := c.ClickCollect().SetIMEIMetadata(context.Background(), client.ClickCollectApiOrdersIMEISetRequest{}); err != nil {
		t.Fatalf("SetIMEIMetadata unexpected error: %v", err)
	} else if httpResp == nil || httpResp.StatusCode != http.StatusOK || resp == nil {
		t.Fatalf("SetIMEIMetadata unexpected response: %+v %+v", httpResp, resp)
	}

	if resp, httpResp, err := c.ClickCollect().SetGTINMetadata(context.Background(), client.ClickCollectApiOrdersGTINSetRequest{}); err != nil {
		t.Fatalf("SetGTINMetadata unexpected error: %v", err)
	} else if httpResp == nil || httpResp.StatusCode != http.StatusOK || resp == nil {
		t.Fatalf("SetGTINMetadata unexpected response: %+v %+v", httpResp, resp)
	}

	expected := []string{
		"GET /api/v3/click-collect/orders/new",
		"GET /api/v3/click-collect/orders",
		"POST /api/v3/click-collect/orders/client",
		"POST /api/v3/click-collect/orders/client/identity",
		"POST /api/marketplace/v3/click-collect/orders/status/info",
		"POST /api/marketplace/v3/click-collect/orders/status/confirm",
		"POST /api/marketplace/v3/click-collect/orders/status/prepare",
		"POST /api/marketplace/v3/click-collect/orders/status/receive",
		"POST /api/marketplace/v3/click-collect/orders/status/reject",
		"POST /api/marketplace/v3/click-collect/orders/status/cancel",
		"POST /api/marketplace/v3/click-collect/orders/meta/info",
		"POST /api/marketplace/v3/click-collect/orders/meta/details",
		"POST /api/marketplace/v3/click-collect/orders/meta/delete",
		"POST /api/marketplace/v3/click-collect/orders/meta/sgtin",
		"POST /api/marketplace/v3/click-collect/orders/meta/uin",
		"POST /api/marketplace/v3/click-collect/orders/meta/imei",
		"POST /api/marketplace/v3/click-collect/orders/meta/gtin",
	}
	for _, key := range expected {
		if calls[key] != 1 {
			t.Fatalf("expected one call for %s, got %d", key, calls[key])
		}
	}
}
