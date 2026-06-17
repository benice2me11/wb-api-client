package client_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/benice2me11/wb-api-client/client"
)

func TestDBWOperations(t *testing.T) {
	t.Parallel()

	calls := map[string]int{}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if auth := r.Header.Get("Authorization"); auth != "Bearer dbw-token" {
			t.Fatalf("auth header mismatch: %q", auth)
		}

		key := r.Method + " " + r.URL.Path
		calls[key]++

		switch key {
		case "POST /api/marketplace/v3/dbw/orders/client":
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprint(w, `{}`)
		case "POST /api/v3/dbw/orders/courier":
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprint(w, `{}`)
		case "POST /api/v3/dbw/orders/delivery-date":
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprint(w, `{}`)
		case "POST /api/v3/dbw/orders/stickers":
			if got := r.URL.Query().Get("type"); got != "pdf" {
				t.Fatalf("unexpected stickers type query param: %s", got)
			}
			if got := r.URL.Query().Get("width"); got != "580" {
				t.Fatalf("unexpected stickers width query param: %s", got)
			}
			if got := r.URL.Query().Get("height"); got != "400" {
				t.Fatalf("unexpected stickers height query param: %s", got)
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprint(w, `{}`)
		case "GET /api/v3/dbw/orders/123/meta":
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprint(w, `{}`)
		case "DELETE /api/v3/dbw/orders/123/meta":
			if got := r.URL.Query().Get("key"); got != "sgtin" {
				t.Fatalf("unexpected metadata key query param: %s", got)
			}
			w.WriteHeader(http.StatusNoContent)
		case "PUT /api/v3/dbw/orders/123/meta/sgtin":
			w.WriteHeader(http.StatusNoContent)
		case "PUT /api/v3/dbw/orders/123/meta/uin":
			w.WriteHeader(http.StatusNoContent)
		case "PUT /api/v3/dbw/orders/123/meta/imei":
			w.WriteHeader(http.StatusNoContent)
		case "PUT /api/v3/dbw/orders/123/meta/gtin":
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
	ordersReq := client.DBWOrdersRequestAPI{Orders: []int32{123}}

	if resp, httpResp, err := c.DBW().BuyerInfo(ctx, ordersReq); err != nil {
		t.Fatalf("BuyerInfo unexpected error: %v", err)
	} else if httpResp == nil || httpResp.StatusCode != http.StatusOK || resp == nil {
		t.Fatalf("BuyerInfo unexpected response: %+v %+v", httpResp, resp)
	}

	if resp, httpResp, err := c.DBW().CourierInfo(ctx, ordersReq); err != nil {
		t.Fatalf("CourierInfo unexpected error: %v", err)
	} else if httpResp == nil || httpResp.StatusCode != http.StatusOK || resp == nil {
		t.Fatalf("CourierInfo unexpected response: %+v %+v", httpResp, resp)
	}

	if resp, httpResp, err := c.DBW().DeliveryDate(ctx, client.DBWDeliveryDatesRequest{Orders: []int32{123}}); err != nil {
		t.Fatalf("DeliveryDate unexpected error: %v", err)
	} else if httpResp == nil || httpResp.StatusCode != http.StatusOK || resp == nil {
		t.Fatalf("DeliveryDate unexpected response: %+v %+v", httpResp, resp)
	}

	stickerType := "pdf"
	stickerWidth := int32(580)
	stickerHeight := int32(400)
	stickerOptions := &client.DBWStickerOptions{
		Type:   &stickerType,
		Width:  &stickerWidth,
		Height: &stickerHeight,
	}
	stickerReq := client.ApiV3DbwOrdersStickersPostRequest{Orders: []int64{123}}
	if resp, httpResp, err := c.DBW().Stickers(ctx, stickerReq, stickerOptions); err != nil {
		t.Fatalf("Stickers unexpected error: %v", err)
	} else if httpResp == nil || httpResp.StatusCode != http.StatusOK || resp == nil {
		t.Fatalf("Stickers unexpected response: %+v %+v", httpResp, resp)
	}

	if resp, httpResp, err := c.DBW().Metadata(ctx, 123); err != nil {
		t.Fatalf("Metadata unexpected error: %v", err)
	} else if httpResp == nil || httpResp.StatusCode != http.StatusOK || resp == nil {
		t.Fatalf("Metadata unexpected response: %+v %+v", httpResp, resp)
	}

	if httpResp, err := c.DBW().DeleteMetadata(ctx, 123, "sgtin"); err != nil {
		t.Fatalf("DeleteMetadata unexpected error: %v", err)
	} else if httpResp == nil || httpResp.StatusCode != http.StatusNoContent {
		t.Fatalf("DeleteMetadata unexpected status: %+v", httpResp)
	}

	sgtinReq := client.ApiV3DbwOrdersOrderIdMetaSgtinPutRequest{}
	sgtinReq.SetSgtins([]string{"010460406000000021ABC"})
	if httpResp, err := c.DBW().SetSGTINMetadata(ctx, 123, sgtinReq); err != nil {
		t.Fatalf("SetSGTINMetadata unexpected error: %v", err)
	} else if httpResp == nil || httpResp.StatusCode != http.StatusNoContent {
		t.Fatalf("SetSGTINMetadata unexpected status: %+v", httpResp)
	}

	uinReq := client.ApiV3DbwOrdersOrderIdMetaUinPutRequest{}
	uinReq.SetUin("uin-1")
	if httpResp, err := c.DBW().SetUINMetadata(ctx, 123, uinReq); err != nil {
		t.Fatalf("SetUINMetadata unexpected error: %v", err)
	} else if httpResp == nil || httpResp.StatusCode != http.StatusNoContent {
		t.Fatalf("SetUINMetadata unexpected status: %+v", httpResp)
	}

	imeiReq := client.ApiV3DbwOrdersOrderIdMetaImeiPutRequest{}
	imeiReq.SetImei("123456789012345")
	if httpResp, err := c.DBW().SetIMEIMetadata(ctx, 123, imeiReq); err != nil {
		t.Fatalf("SetIMEIMetadata unexpected error: %v", err)
	} else if httpResp == nil || httpResp.StatusCode != http.StatusNoContent {
		t.Fatalf("SetIMEIMetadata unexpected status: %+v", httpResp)
	}

	gtinReq := client.ApiV3DbwOrdersOrderIdMetaGtinPutRequest{}
	gtinReq.SetGtin("04604060000000")
	if httpResp, err := c.DBW().SetGTINMetadata(ctx, 123, gtinReq); err != nil {
		t.Fatalf("SetGTINMetadata unexpected error: %v", err)
	} else if httpResp == nil || httpResp.StatusCode != http.StatusNoContent {
		t.Fatalf("SetGTINMetadata unexpected status: %+v", httpResp)
	}

	expected := []string{
		"POST /api/marketplace/v3/dbw/orders/client",
		"POST /api/v3/dbw/orders/courier",
		"POST /api/v3/dbw/orders/delivery-date",
		"POST /api/v3/dbw/orders/stickers",
		"GET /api/v3/dbw/orders/123/meta",
		"DELETE /api/v3/dbw/orders/123/meta",
		"PUT /api/v3/dbw/orders/123/meta/sgtin",
		"PUT /api/v3/dbw/orders/123/meta/uin",
		"PUT /api/v3/dbw/orders/123/meta/imei",
		"PUT /api/v3/dbw/orders/123/meta/gtin",
	}
	for _, key := range expected {
		if calls[key] != 1 {
			t.Fatalf("expected one call for %s, got %d", key, calls[key])
		}
	}
}
