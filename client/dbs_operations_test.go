package client_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/benice2me11/wb-api-client/client"
)

func TestDBSOperations(t *testing.T) {
	t.Parallel()

	calls := map[string]int{}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if auth := r.Header.Get("Authorization"); auth != "Bearer dbs-token" {
			t.Fatalf("auth header mismatch: %q", auth)
		}

		key := r.Method + " " + r.URL.Path
		calls[key]++

		switch key {
		case "POST /api/v3/dbs/groups/info":
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprint(w, `[]`)
		case "POST /api/v3/dbs/orders/client":
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprint(w, `{}`)
		case "POST /api/marketplace/v3/dbs/orders/b2b/info":
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprint(w, `{"requestId":"req","results":[]}`)
		case "POST /api/v3/dbs/orders/delivery-date":
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprint(w, `{}`)
		case "POST /api/marketplace/v3/dbs/orders/stickers":
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
		case "POST /api/marketplace/v3/dbs/orders/meta/info":
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprint(w, `{}`)
		case "POST /api/marketplace/v3/dbs/orders/meta/delete":
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprint(w, `{}`)
		case "POST /api/marketplace/v3/dbs/orders/meta/sgtin":
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprint(w, `{}`)
		case "POST /api/marketplace/v3/dbs/orders/meta/uin":
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprint(w, `{}`)
		case "POST /api/marketplace/v3/dbs/orders/meta/imei":
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprint(w, `{}`)
		case "POST /api/marketplace/v3/dbs/orders/meta/gtin":
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprint(w, `{}`)
		case "POST /api/marketplace/v3/dbs/orders/meta/customs-declaration":
			w.WriteHeader(http.StatusNoContent)
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
	ordersReq := client.ApiOrdersRequestV2{OrdersIds: []int32{321}}

	if resp, httpResp, err := c.DBS().GroupsInfo(ctx, client.ApiOrderGroupsRequest{Groups: []string{"group-1"}}); err != nil {
		t.Fatalf("GroupsInfo unexpected error: %v", err)
	} else if httpResp == nil || httpResp.StatusCode != http.StatusOK || resp == nil {
		t.Fatalf("GroupsInfo unexpected response: %+v %+v", httpResp, resp)
	}

	if resp, httpResp, err := c.DBS().BuyerInfo(ctx, ordersReq); err != nil {
		t.Fatalf("BuyerInfo unexpected error: %v", err)
	} else if httpResp == nil || httpResp.StatusCode != http.StatusOK || resp == nil {
		t.Fatalf("BuyerInfo unexpected response: %+v %+v", httpResp, resp)
	}

	if resp, httpResp, err := c.DBS().B2BInfo(ctx, ordersReq); err != nil {
		t.Fatalf("B2BInfo unexpected error: %v", err)
	} else if httpResp == nil || httpResp.StatusCode != http.StatusOK || resp == nil {
		t.Fatalf("B2BInfo unexpected response: %+v %+v", httpResp, resp)
	}

	if resp, httpResp, err := c.DBS().DeliveryDate(ctx, client.DeliveryDatesRequest{Orders: []int32{321}}); err != nil {
		t.Fatalf("DeliveryDate unexpected error: %v", err)
	} else if httpResp == nil || httpResp.StatusCode != http.StatusOK || resp == nil {
		t.Fatalf("DeliveryDate unexpected response: %+v %+v", httpResp, resp)
	}

	stickerType := "pdf"
	stickerWidth := int32(580)
	stickerHeight := int32(400)
	stickerOptions := &client.DBSStickerOptions{
		Type:   &stickerType,
		Width:  &stickerWidth,
		Height: &stickerHeight,
	}
	stickerReq := client.ApiMarketplaceV3DbsOrdersStickersPostRequest{Orders: []int64{321}}
	if resp, httpResp, err := c.DBS().Stickers(ctx, stickerReq, stickerOptions); err != nil {
		t.Fatalf("Stickers unexpected error: %v", err)
	} else if httpResp == nil || httpResp.StatusCode != http.StatusOK || resp == nil {
		t.Fatalf("Stickers unexpected response: %+v %+v", httpResp, resp)
	}

	if resp, httpResp, err := c.DBS().MetadataInfo(ctx, ordersReq); err != nil {
		t.Fatalf("MetadataInfo unexpected error: %v", err)
	} else if httpResp == nil || httpResp.StatusCode != http.StatusOK || resp == nil {
		t.Fatalf("MetadataInfo unexpected response: %+v %+v", httpResp, resp)
	}

	deleteMetaReq := client.ApiOrdersMetaDeleteRequest{Key: "sgtin", OrderIds: []int32{321}}
	if resp, httpResp, err := c.DBS().DeleteMetadata(ctx, deleteMetaReq); err != nil {
		t.Fatalf("DeleteMetadata unexpected error: %v", err)
	} else if httpResp == nil || httpResp.StatusCode != http.StatusOK || resp == nil {
		t.Fatalf("DeleteMetadata unexpected response: %+v %+v", httpResp, resp)
	}

	sgtin := client.ApiSGTINs{}
	sgtin.SetOrderId(321)
	sgtin.SetSgtins([]string{"010460406000000021ABC"})
	if resp, httpResp, err := c.DBS().SetSGTINMetadata(ctx, client.ApiOrdersSGTINsSetRequest{Orders: []client.ApiSGTINs{sgtin}}); err != nil {
		t.Fatalf("SetSGTINMetadata unexpected error: %v", err)
	} else if httpResp == nil || httpResp.StatusCode != http.StatusOK || resp == nil {
		t.Fatalf("SetSGTINMetadata unexpected response: %+v %+v", httpResp, resp)
	}

	uin := client.ApiUIN{}
	uin.SetOrderId(321)
	uin.SetUin("uin-1")
	if resp, httpResp, err := c.DBS().SetUINMetadata(ctx, client.ApiOrdersUINSetRequest{Orders: []client.ApiUIN{uin}}); err != nil {
		t.Fatalf("SetUINMetadata unexpected error: %v", err)
	} else if httpResp == nil || httpResp.StatusCode != http.StatusOK || resp == nil {
		t.Fatalf("SetUINMetadata unexpected response: %+v %+v", httpResp, resp)
	}

	imei := client.ApiIMEI{}
	imei.SetOrderId(321)
	imei.SetImei("123456789012345")
	if resp, httpResp, err := c.DBS().SetIMEIMetadata(ctx, client.ApiOrdersIMEISetRequest{Orders: []client.ApiIMEI{imei}}); err != nil {
		t.Fatalf("SetIMEIMetadata unexpected error: %v", err)
	} else if httpResp == nil || httpResp.StatusCode != http.StatusOK || resp == nil {
		t.Fatalf("SetIMEIMetadata unexpected response: %+v %+v", httpResp, resp)
	}

	gtin := client.ApiGTIN{}
	gtin.SetOrderId(321)
	gtin.SetGtin("04604060000000")
	if resp, httpResp, err := c.DBS().SetGTINMetadata(ctx, client.ApiOrdersGTINSetRequest{Orders: []client.ApiGTIN{gtin}}); err != nil {
		t.Fatalf("SetGTINMetadata unexpected error: %v", err)
	} else if httpResp == nil || httpResp.StatusCode != http.StatusOK || resp == nil {
		t.Fatalf("SetGTINMetadata unexpected response: %+v %+v", httpResp, resp)
	}

	customsReq := client.ApiMarketplaceV3DbsOrdersMetaCustomsDeclarationPostRequest{
		Orders: []client.ApiMarketplaceV3DbsOrdersMetaCustomsDeclarationPostRequestOrdersInner{
			{OrderId: 321, CustomsDeclaration: "declaration-1"},
		},
	}
	if httpResp, err := c.DBS().SetCustomsDeclarationMetadata(ctx, customsReq); err != nil {
		t.Fatalf("SetCustomsDeclarationMetadata unexpected error: %v", err)
	} else if httpResp == nil || httpResp.StatusCode != http.StatusNoContent {
		t.Fatalf("SetCustomsDeclarationMetadata unexpected status: %+v", httpResp)
	}

	expected := []string{
		"POST /api/v3/dbs/groups/info",
		"POST /api/v3/dbs/orders/client",
		"POST /api/marketplace/v3/dbs/orders/b2b/info",
		"POST /api/v3/dbs/orders/delivery-date",
		"POST /api/marketplace/v3/dbs/orders/stickers",
		"POST /api/marketplace/v3/dbs/orders/meta/info",
		"POST /api/marketplace/v3/dbs/orders/meta/delete",
		"POST /api/marketplace/v3/dbs/orders/meta/sgtin",
		"POST /api/marketplace/v3/dbs/orders/meta/uin",
		"POST /api/marketplace/v3/dbs/orders/meta/imei",
		"POST /api/marketplace/v3/dbs/orders/meta/gtin",
		"POST /api/marketplace/v3/dbs/orders/meta/customs-declaration",
	}
	for _, key := range expected {
		if calls[key] != 1 {
			t.Fatalf("expected one call for %s, got %d", key, calls[key])
		}
	}
}
