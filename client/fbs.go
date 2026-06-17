package client

import (
	"context"
	"net/http"

	"github.com/benice2me11/wb-api-client/transport"
)

// FBSOrdersQuery defines optional filters for list orders call.
type FBSOrdersQuery struct {
	Limit    *int32
	Next     *int64
	DateFrom *int32
	DateTo   *int32
}

// FBSStickerOptions defines optional sticker output query parameters.
type FBSStickerOptions struct {
	Type   *string
	Width  *int32
	Height *int32
}

// FBSSuppliesQuery defines required pagination for supplies list call.
type FBSSuppliesQuery struct {
	Limit int32
	Next  int64
}

// FBSService provides ergonomic access to FBS Orders operations.
type FBSService interface {
	Orders(ctx context.Context, query *FBSOrdersQuery) (*ApiV3OrdersGet200Response, *http.Response, error)
	NewOrders(ctx context.Context) (*ApiV3OrdersNewGet200Response, *http.Response, error)
	OrdersByStatus(ctx context.Context, request ApiV3OrdersStatusPostRequest) (*ApiV3OrdersStatusPost200Response, *http.Response, error)
	OrdersStatusHistory(ctx context.Context, request ApiV3OrdersStatusHistoryPostRequest) (*ApiV3OrdersStatusHistoryPost200Response, *http.Response, error)
	Stickers(ctx context.Context, request ApiV3OrdersStickersPostRequest, options *FBSStickerOptions) (*ApiV3OrdersStickersPost200Response, *http.Response, error)
	CrossBorderStickers(ctx context.Context, request ApiV3OrdersStickersCrossBorderPostRequest) (*ApiV3OrdersStickersCrossBorderPost200Response, *http.Response, error)
	ClientInfo(ctx context.Context, request OrdersRequestAPI) (*CrossborderTurkeyClientInfoResp, *http.Response, error)
	ReshipmentOrders(ctx context.Context) (*ApiV3SuppliesOrdersReshipmentGet200Response, *http.Response, error)
	Supplies(ctx context.Context, query FBSSuppliesQuery) (*ApiV3SuppliesGet200Response, *http.Response, error)
	CreateSupply(ctx context.Context, request ApiV3SuppliesPostRequest) (*ApiV3SuppliesPost201Response, *http.Response, error)
	Supply(ctx context.Context, supplyID string) (*Supply, *http.Response, error)
	DeleteSupply(ctx context.Context, supplyID string) (*http.Response, error)
	DeliverSupply(ctx context.Context, supplyID string) (*http.Response, error)
	SupplyBarcode(ctx context.Context, supplyID string, typ *string) (*ApiV3SuppliesSupplyIdBarcodeGet200Response, *http.Response, error)
	SupplyOrderIDs(ctx context.Context, supplyID string) (*V3SupplyOrderIDsAPI, *http.Response, error)
	AddOrdersToSupply(ctx context.Context, supplyID string, request OrdersRequestAPI) (*http.Response, error)
	Passes(ctx context.Context) ([]Pass, *http.Response, error)
	PassOffices(ctx context.Context) ([]PassOffice, *http.Response, error)
	CreatePass(ctx context.Context, request ApiV3PassesPostRequest) (*ApiV3PassesPost201Response, *http.Response, error)
	UpdatePass(ctx context.Context, passID int64, request ApiV3PassesPostRequest) (*http.Response, error)
	DeletePass(ctx context.Context, passID int64) (*http.Response, error)
	Metadata(ctx context.Context, request V3GetMetaMultiRequest) (*V3OrdersMetaAPI, *http.Response, error)
	DeleteMetadata(ctx context.Context, orderID int64, key string) (*http.Response, error)
	SetSGTINMetadata(ctx context.Context, orderID int64, request ApiV3OrdersOrderIdMetaSgtinPutRequest) (*http.Response, error)
	SetUINMetadata(ctx context.Context, orderID int64, request ApiV3OrdersOrderIdMetaUinPutRequest) (*http.Response, error)
	SetIMEIMetadata(ctx context.Context, orderID int64, request ApiV3OrdersOrderIdMetaImeiPutRequest) (*http.Response, error)
	SetGTINMetadata(ctx context.Context, orderID int64, request ApiV3OrdersOrderIdMetaGtinPutRequest) (*http.Response, error)
	SetExpirationMetadata(ctx context.Context, orderID int64, request ApiV3OrdersOrderIdMetaExpirationPutRequest) (*http.Response, error)
	SetCustomsDeclarationMetadata(ctx context.Context, orderID int64, request ApiMarketplaceV3OrdersOrderIdMetaCustomsDeclarationPutRequest) (*http.Response, error)
	CancelOrder(ctx context.Context, orderID int64) (*http.Response, error)
}

type fbsService struct {
	api *FBSAPIClient
}

func (s *fbsService) Orders(ctx context.Context, query *FBSOrdersQuery) (*ApiV3OrdersGet200Response, *http.Response, error) {
	req := s.api.FBSAssemblyOrdersAPI.ApiV3OrdersGet(ctx)
	if query != nil {
		if query.Limit != nil {
			req = req.Limit(*query.Limit)
		}
		if query.Next != nil {
			req = req.Next(*query.Next)
		}
		if query.DateFrom != nil {
			req = req.DateFrom(*query.DateFrom)
		}
		if query.DateTo != nil {
			req = req.DateTo(*query.DateTo)
		}
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, transport.WrapResponseError(httpResp, err)
	}
	return resp, httpResp, nil
}

func (s *fbsService) NewOrders(ctx context.Context) (*ApiV3OrdersNewGet200Response, *http.Response, error) {
	req := s.api.FBSAssemblyOrdersAPI.ApiV3OrdersNewGet(ctx)

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, transport.WrapResponseError(httpResp, err)
	}
	return resp, httpResp, nil
}

func (s *fbsService) OrdersByStatus(ctx context.Context, request ApiV3OrdersStatusPostRequest) (*ApiV3OrdersStatusPost200Response, *http.Response, error) {
	resp, httpResp, err := s.api.FBSAssemblyOrdersAPI.
		ApiV3OrdersStatusPost(ctx).
		ApiV3OrdersStatusPostRequest(request).
		Execute()
	if err != nil {
		return nil, httpResp, transport.WrapResponseError(httpResp, err)
	}
	return resp, httpResp, nil
}

func (s *fbsService) OrdersStatusHistory(ctx context.Context, request ApiV3OrdersStatusHistoryPostRequest) (*ApiV3OrdersStatusHistoryPost200Response, *http.Response, error) {
	resp, httpResp, err := s.api.FBSAssemblyOrdersAPI.
		ApiV3OrdersStatusHistoryPost(ctx).
		ApiV3OrdersStatusHistoryPostRequest(request).
		Execute()
	if err != nil {
		return nil, httpResp, transport.WrapResponseError(httpResp, err)
	}
	return resp, httpResp, nil
}

func (s *fbsService) Stickers(ctx context.Context, request ApiV3OrdersStickersPostRequest, options *FBSStickerOptions) (*ApiV3OrdersStickersPost200Response, *http.Response, error) {
	req := s.api.FBSAssemblyOrdersAPI.
		ApiV3OrdersStickersPost(ctx).
		ApiV3OrdersStickersPostRequest(request)
	if options != nil {
		if options.Type != nil {
			req = req.Type_(*options.Type)
		}
		if options.Width != nil {
			req = req.Width(*options.Width)
		}
		if options.Height != nil {
			req = req.Height(*options.Height)
		}
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, transport.WrapResponseError(httpResp, err)
	}
	return resp, httpResp, nil
}

func (s *fbsService) CrossBorderStickers(ctx context.Context, request ApiV3OrdersStickersCrossBorderPostRequest) (*ApiV3OrdersStickersCrossBorderPost200Response, *http.Response, error) {
	resp, httpResp, err := s.api.FBSAssemblyOrdersAPI.
		ApiV3OrdersStickersCrossBorderPost(ctx).
		ApiV3OrdersStickersCrossBorderPostRequest(request).
		Execute()
	if err != nil {
		return nil, httpResp, transport.WrapResponseError(httpResp, err)
	}
	return resp, httpResp, nil
}

func (s *fbsService) ClientInfo(ctx context.Context, request OrdersRequestAPI) (*CrossborderTurkeyClientInfoResp, *http.Response, error) {
	resp, httpResp, err := s.api.FBSAssemblyOrdersAPI.
		ApiV3OrdersClientPost(ctx).
		OrdersRequestAPI(request).
		Execute()
	if err != nil {
		return nil, httpResp, transport.WrapResponseError(httpResp, err)
	}
	return resp, httpResp, nil
}

func (s *fbsService) ReshipmentOrders(ctx context.Context) (*ApiV3SuppliesOrdersReshipmentGet200Response, *http.Response, error) {
	resp, httpResp, err := s.api.FBSAssemblyOrdersAPI.
		ApiV3SuppliesOrdersReshipmentGet(ctx).
		Execute()
	if err != nil {
		return nil, httpResp, transport.WrapResponseError(httpResp, err)
	}
	return resp, httpResp, nil
}

func (s *fbsService) CancelOrder(ctx context.Context, orderID int64) (*http.Response, error) {
	httpResp, err := s.api.FBSAssemblyOrdersAPI.
		ApiV3OrdersOrderIdCancelPatch(ctx, orderID).
		Execute()
	if err != nil {
		return httpResp, transport.WrapResponseError(httpResp, err)
	}
	return httpResp, nil
}
