package client_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/benice2me11/wb-api-client/client"
)

func TestFBSOperations(t *testing.T) {
	t.Parallel()

	calls := map[string]int{}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if auth := r.Header.Get("Authorization"); auth != "Bearer fbs-token" {
			t.Fatalf("auth header mismatch: %q", auth)
		}

		key := r.Method + " " + r.URL.Path
		calls[key]++

		switch key {
		case "POST /api/v3/orders/stickers":
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
		case "POST /api/v3/orders/stickers/cross-border":
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprint(w, `{}`)
		case "POST /api/v3/orders/client":
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprint(w, `{}`)
		case "GET /api/v3/supplies/orders/reshipment":
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprint(w, `{}`)
		case "GET /api/v3/supplies":
			if got := r.URL.Query().Get("limit"); got != "100" {
				t.Fatalf("unexpected supplies limit query param: %s", got)
			}
			if got := r.URL.Query().Get("next"); got != "0" {
				t.Fatalf("unexpected supplies next query param: %s", got)
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprint(w, `{}`)
		case "POST /api/v3/supplies":
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			_, _ = fmt.Fprint(w, `{}`)
		case "GET /api/v3/supplies/WB-GI-1":
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprint(w, `{}`)
		case "DELETE /api/v3/supplies/WB-GI-1":
			w.WriteHeader(http.StatusNoContent)
		case "PATCH /api/v3/supplies/WB-GI-1/deliver":
			w.WriteHeader(http.StatusNoContent)
		case "GET /api/v3/supplies/WB-GI-1/barcode":
			if got := r.URL.Query().Get("type"); got != "svg" {
				t.Fatalf("unexpected supply barcode type query param: %s", got)
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprint(w, `{}`)
		case "GET /api/marketplace/v3/supplies/WB-GI-1/order-ids":
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprint(w, `{}`)
		case "PATCH /api/marketplace/v3/supplies/WB-GI-1/orders":
			w.WriteHeader(http.StatusNoContent)
		case "GET /api/v3/passes":
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprint(w, `[]`)
		case "GET /api/v3/passes/offices":
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprint(w, `[]`)
		case "POST /api/v3/passes":
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			_, _ = fmt.Fprint(w, `{}`)
		case "PUT /api/v3/passes/123":
			w.WriteHeader(http.StatusNoContent)
		case "DELETE /api/v3/passes/123":
			w.WriteHeader(http.StatusNoContent)
		case "POST /api/marketplace/v3/orders/meta":
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprint(w, `{}`)
		case "DELETE /api/v3/orders/123/meta":
			if got := r.URL.Query().Get("key"); got != "sgtin" {
				t.Fatalf("unexpected metadata key query param: %s", got)
			}
			w.WriteHeader(http.StatusNoContent)
		case "PUT /api/v3/orders/123/meta/sgtin":
			w.WriteHeader(http.StatusNoContent)
		case "PUT /api/v3/orders/123/meta/uin":
			w.WriteHeader(http.StatusNoContent)
		case "PUT /api/v3/orders/123/meta/imei":
			w.WriteHeader(http.StatusNoContent)
		case "PUT /api/v3/orders/123/meta/gtin":
			w.WriteHeader(http.StatusNoContent)
		case "PUT /api/v3/orders/123/meta/expiration":
			w.WriteHeader(http.StatusNoContent)
		case "PUT /api/marketplace/v3/orders/123/meta/customs-declaration":
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
	orderReq := client.OrdersRequestAPI{Orders: []int32{123}}

	stickerType := "pdf"
	stickerWidth := int32(580)
	stickerHeight := int32(400)
	stickerOptions := &client.FBSStickerOptions{
		Type:   &stickerType,
		Width:  &stickerWidth,
		Height: &stickerHeight,
	}
	if resp, httpResp, err := c.FBS().Stickers(ctx, client.ApiV3OrdersStickersPostRequest{Orders: []int64{123}}, stickerOptions); err != nil {
		t.Fatalf("Stickers unexpected error: %v", err)
	} else if httpResp == nil || httpResp.StatusCode != http.StatusOK || resp == nil {
		t.Fatalf("Stickers unexpected response: %+v %+v", httpResp, resp)
	}

	if resp, httpResp, err := c.FBS().CrossBorderStickers(ctx, client.ApiV3OrdersStickersCrossBorderPostRequest{Orders: []int64{123}}); err != nil {
		t.Fatalf("CrossBorderStickers unexpected error: %v", err)
	} else if httpResp == nil || httpResp.StatusCode != http.StatusOK || resp == nil {
		t.Fatalf("CrossBorderStickers unexpected response: %+v %+v", httpResp, resp)
	}

	if resp, httpResp, err := c.FBS().ClientInfo(ctx, orderReq); err != nil {
		t.Fatalf("ClientInfo unexpected error: %v", err)
	} else if httpResp == nil || httpResp.StatusCode != http.StatusOK || resp == nil {
		t.Fatalf("ClientInfo unexpected response: %+v %+v", httpResp, resp)
	}

	if resp, httpResp, err := c.FBS().ReshipmentOrders(ctx); err != nil {
		t.Fatalf("ReshipmentOrders unexpected error: %v", err)
	} else if httpResp == nil || httpResp.StatusCode != http.StatusOK || resp == nil {
		t.Fatalf("ReshipmentOrders unexpected response: %+v %+v", httpResp, resp)
	}

	if resp, httpResp, err := c.FBS().Supplies(ctx, client.FBSSuppliesQuery{Limit: 100, Next: 0}); err != nil {
		t.Fatalf("Supplies unexpected error: %v", err)
	} else if httpResp == nil || httpResp.StatusCode != http.StatusOK || resp == nil {
		t.Fatalf("Supplies unexpected response: %+v %+v", httpResp, resp)
	}

	supplyReq := client.ApiV3SuppliesPostRequest{}
	supplyReq.SetName("daily supply")
	if resp, httpResp, err := c.FBS().CreateSupply(ctx, supplyReq); err != nil {
		t.Fatalf("CreateSupply unexpected error: %v", err)
	} else if httpResp == nil || httpResp.StatusCode != http.StatusCreated || resp == nil {
		t.Fatalf("CreateSupply unexpected response: %+v %+v", httpResp, resp)
	}

	if resp, httpResp, err := c.FBS().Supply(ctx, "WB-GI-1"); err != nil {
		t.Fatalf("Supply unexpected error: %v", err)
	} else if httpResp == nil || httpResp.StatusCode != http.StatusOK || resp == nil {
		t.Fatalf("Supply unexpected response: %+v %+v", httpResp, resp)
	}

	if httpResp, err := c.FBS().DeleteSupply(ctx, "WB-GI-1"); err != nil {
		t.Fatalf("DeleteSupply unexpected error: %v", err)
	} else if httpResp == nil || httpResp.StatusCode != http.StatusNoContent {
		t.Fatalf("DeleteSupply unexpected status: %+v", httpResp)
	}

	if httpResp, err := c.FBS().DeliverSupply(ctx, "WB-GI-1"); err != nil {
		t.Fatalf("DeliverSupply unexpected error: %v", err)
	} else if httpResp == nil || httpResp.StatusCode != http.StatusNoContent {
		t.Fatalf("DeliverSupply unexpected status: %+v", httpResp)
	}

	barcodeType := "svg"
	if resp, httpResp, err := c.FBS().SupplyBarcode(ctx, "WB-GI-1", &barcodeType); err != nil {
		t.Fatalf("SupplyBarcode unexpected error: %v", err)
	} else if httpResp == nil || httpResp.StatusCode != http.StatusOK || resp == nil {
		t.Fatalf("SupplyBarcode unexpected response: %+v %+v", httpResp, resp)
	}

	if resp, httpResp, err := c.FBS().SupplyOrderIDs(ctx, "WB-GI-1"); err != nil {
		t.Fatalf("SupplyOrderIDs unexpected error: %v", err)
	} else if httpResp == nil || httpResp.StatusCode != http.StatusOK || resp == nil {
		t.Fatalf("SupplyOrderIDs unexpected response: %+v %+v", httpResp, resp)
	}

	if httpResp, err := c.FBS().AddOrdersToSupply(ctx, "WB-GI-1", orderReq); err != nil {
		t.Fatalf("AddOrdersToSupply unexpected error: %v", err)
	} else if httpResp == nil || httpResp.StatusCode != http.StatusNoContent {
		t.Fatalf("AddOrdersToSupply unexpected status: %+v", httpResp)
	}

	if resp, httpResp, err := c.FBS().Passes(ctx); err != nil {
		t.Fatalf("Passes unexpected error: %v", err)
	} else if httpResp == nil || httpResp.StatusCode != http.StatusOK || resp == nil {
		t.Fatalf("Passes unexpected response: %+v %+v", httpResp, resp)
	}

	if resp, httpResp, err := c.FBS().PassOffices(ctx); err != nil {
		t.Fatalf("PassOffices unexpected error: %v", err)
	} else if httpResp == nil || httpResp.StatusCode != http.StatusOK || resp == nil {
		t.Fatalf("PassOffices unexpected response: %+v %+v", httpResp, resp)
	}

	passReq := client.ApiV3PassesPostRequest{
		FirstName: "Ivan",
		LastName:  "Ivanov",
		CarModel:  "Lada",
		CarNumber: "A001AA",
		OfficeId:  1,
	}
	if resp, httpResp, err := c.FBS().CreatePass(ctx, passReq); err != nil {
		t.Fatalf("CreatePass unexpected error: %v", err)
	} else if httpResp == nil || httpResp.StatusCode != http.StatusCreated || resp == nil {
		t.Fatalf("CreatePass unexpected response: %+v %+v", httpResp, resp)
	}

	if httpResp, err := c.FBS().UpdatePass(ctx, 123, passReq); err != nil {
		t.Fatalf("UpdatePass unexpected error: %v", err)
	} else if httpResp == nil || httpResp.StatusCode != http.StatusNoContent {
		t.Fatalf("UpdatePass unexpected status: %+v", httpResp)
	}

	if httpResp, err := c.FBS().DeletePass(ctx, 123); err != nil {
		t.Fatalf("DeletePass unexpected error: %v", err)
	} else if httpResp == nil || httpResp.StatusCode != http.StatusNoContent {
		t.Fatalf("DeletePass unexpected status: %+v", httpResp)
	}

	if resp, httpResp, err := c.FBS().Metadata(ctx, client.V3GetMetaMultiRequest{Orders: []int32{123}}); err != nil {
		t.Fatalf("Metadata unexpected error: %v", err)
	} else if httpResp == nil || httpResp.StatusCode != http.StatusOK || resp == nil {
		t.Fatalf("Metadata unexpected response: %+v %+v", httpResp, resp)
	}

	if httpResp, err := c.FBS().DeleteMetadata(ctx, 123, "sgtin"); err != nil {
		t.Fatalf("DeleteMetadata unexpected error: %v", err)
	} else if httpResp == nil || httpResp.StatusCode != http.StatusNoContent {
		t.Fatalf("DeleteMetadata unexpected status: %+v", httpResp)
	}

	sgtinReq := client.ApiV3OrdersOrderIdMetaSgtinPutRequest{}
	sgtinReq.SetSgtins([]string{"010460406000000021ABC"})
	if httpResp, err := c.FBS().SetSGTINMetadata(ctx, 123, sgtinReq); err != nil {
		t.Fatalf("SetSGTINMetadata unexpected error: %v", err)
	} else if httpResp == nil || httpResp.StatusCode != http.StatusNoContent {
		t.Fatalf("SetSGTINMetadata unexpected status: %+v", httpResp)
	}

	uinReq := client.ApiV3OrdersOrderIdMetaUinPutRequest{}
	uinReq.SetUin("uin-1")
	if httpResp, err := c.FBS().SetUINMetadata(ctx, 123, uinReq); err != nil {
		t.Fatalf("SetUINMetadata unexpected error: %v", err)
	} else if httpResp == nil || httpResp.StatusCode != http.StatusNoContent {
		t.Fatalf("SetUINMetadata unexpected status: %+v", httpResp)
	}

	imeiReq := client.ApiV3OrdersOrderIdMetaImeiPutRequest{}
	imeiReq.SetImei("123456789012345")
	if httpResp, err := c.FBS().SetIMEIMetadata(ctx, 123, imeiReq); err != nil {
		t.Fatalf("SetIMEIMetadata unexpected error: %v", err)
	} else if httpResp == nil || httpResp.StatusCode != http.StatusNoContent {
		t.Fatalf("SetIMEIMetadata unexpected status: %+v", httpResp)
	}

	gtinReq := client.ApiV3OrdersOrderIdMetaGtinPutRequest{}
	gtinReq.SetGtin("04604060000000")
	if httpResp, err := c.FBS().SetGTINMetadata(ctx, 123, gtinReq); err != nil {
		t.Fatalf("SetGTINMetadata unexpected error: %v", err)
	} else if httpResp == nil || httpResp.StatusCode != http.StatusNoContent {
		t.Fatalf("SetGTINMetadata unexpected status: %+v", httpResp)
	}

	expirationReq := client.ApiV3OrdersOrderIdMetaExpirationPutRequest{}
	expirationReq.SetExpiration("2026-12-31")
	if httpResp, err := c.FBS().SetExpirationMetadata(ctx, 123, expirationReq); err != nil {
		t.Fatalf("SetExpirationMetadata unexpected error: %v", err)
	} else if httpResp == nil || httpResp.StatusCode != http.StatusNoContent {
		t.Fatalf("SetExpirationMetadata unexpected status: %+v", httpResp)
	}

	customsReq := client.ApiMarketplaceV3OrdersOrderIdMetaCustomsDeclarationPutRequest{}
	customsReq.SetCustomsDeclaration("declaration-1")
	if httpResp, err := c.FBS().SetCustomsDeclarationMetadata(ctx, 123, customsReq); err != nil {
		t.Fatalf("SetCustomsDeclarationMetadata unexpected error: %v", err)
	} else if httpResp == nil || httpResp.StatusCode != http.StatusNoContent {
		t.Fatalf("SetCustomsDeclarationMetadata unexpected status: %+v", httpResp)
	}

	expected := []string{
		"POST /api/v3/orders/stickers",
		"POST /api/v3/orders/stickers/cross-border",
		"POST /api/v3/orders/client",
		"GET /api/v3/supplies/orders/reshipment",
		"GET /api/v3/supplies",
		"POST /api/v3/supplies",
		"GET /api/v3/supplies/WB-GI-1",
		"DELETE /api/v3/supplies/WB-GI-1",
		"PATCH /api/v3/supplies/WB-GI-1/deliver",
		"GET /api/v3/supplies/WB-GI-1/barcode",
		"GET /api/marketplace/v3/supplies/WB-GI-1/order-ids",
		"PATCH /api/marketplace/v3/supplies/WB-GI-1/orders",
		"GET /api/v3/passes",
		"GET /api/v3/passes/offices",
		"POST /api/v3/passes",
		"PUT /api/v3/passes/123",
		"DELETE /api/v3/passes/123",
		"POST /api/marketplace/v3/orders/meta",
		"DELETE /api/v3/orders/123/meta",
		"PUT /api/v3/orders/123/meta/sgtin",
		"PUT /api/v3/orders/123/meta/uin",
		"PUT /api/v3/orders/123/meta/imei",
		"PUT /api/v3/orders/123/meta/gtin",
		"PUT /api/v3/orders/123/meta/expiration",
		"PUT /api/marketplace/v3/orders/123/meta/customs-declaration",
	}
	for _, key := range expected {
		if calls[key] != 1 {
			t.Fatalf("expected one call for %s, got %d", key, calls[key])
		}
	}
}
