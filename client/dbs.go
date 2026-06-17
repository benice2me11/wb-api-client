package client

import (
	"context"
	"fmt"
	"math"
	"net/http"

	"github.com/benice2me11/wb-api-client/transport"
)

// DBSOrdersQuery contains required filters for completed DBS orders list.
type DBSOrdersQuery struct {
	Limit    int32
	Next     int64
	DateFrom int32
	DateTo   int32
}

// DBSStickerOptions defines optional sticker output query parameters.
type DBSStickerOptions struct {
	Type   *string
	Width  *int32
	Height *int32
}

// DBSService provides ergonomic access to DBS Orders operations.
type DBSService interface {
	Orders(ctx context.Context, query DBSOrdersQuery) (*ApiV3DbsOrdersGet200Response, *http.Response, error)
	NewOrders(ctx context.Context) (*ApiV3DbsOrdersNewGet200Response, *http.Response, error)
	GroupsInfo(ctx context.Context, request ApiOrderGroupsRequest) ([]ApiOrderGroupInner, *http.Response, error)
	BuyerInfo(ctx context.Context, request ApiOrdersRequestV2) (*DbsOnlyClientInfoResp, *http.Response, error)
	B2BInfo(ctx context.Context, request ApiOrdersRequestV2) (*ApiB2bClientInfoResponses, *http.Response, error)
	DeliveryDate(ctx context.Context, request DeliveryDatesRequest) (*DeliveryDatesInfoResp, *http.Response, error)
	Stickers(ctx context.Context, request ApiMarketplaceV3DbsOrdersStickersPostRequest, options *DBSStickerOptions) (*ApiMarketplaceV3DbsOrdersStickersPost200Response, *http.Response, error)
	OrdersStatusInfo(ctx context.Context, request ApiOrdersRequestV2) (*ApiOrderStatusesV2, *http.Response, error)
	ConfirmOrders(ctx context.Context, request ApiOrdersRequestV2) (*ApiStatusSetResponses, *http.Response, error)
	DeliverOrders(ctx context.Context, request ApiOrdersRequestV2) (*ApiStatusSetResponses, *http.Response, error)
	ReceiveOrders(ctx context.Context, request ApiOrdersCodeRequest) (*ApiMarketplaceV3DbsOrdersStatusReceivePost200Response, *http.Response, error)
	RejectOrders(ctx context.Context, request ApiOrdersCodeRequest) (*ApiStatusSetResponses, *http.Response, error)
	CancelOrders(ctx context.Context, request ApiOrdersRequestV2) (*ApiStatusSetResponses, *http.Response, error)
	MetadataInfo(ctx context.Context, request ApiOrdersRequestV2) (*ApiOrdersMetaResponse, *http.Response, error)
	DeleteMetadata(ctx context.Context, request ApiOrdersMetaDeleteRequest) (*ApiStatusSetResponses, *http.Response, error)
	SetSGTINMetadata(ctx context.Context, request ApiOrdersSGTINsSetRequest) (*ApiStatusSetResponses, *http.Response, error)
	SetUINMetadata(ctx context.Context, request ApiOrdersUINSetRequest) (*ApiStatusSetResponses, *http.Response, error)
	SetIMEIMetadata(ctx context.Context, request ApiOrdersIMEISetRequest) (*ApiStatusSetResponses, *http.Response, error)
	SetGTINMetadata(ctx context.Context, request ApiOrdersGTINSetRequest) (*ApiStatusSetResponses, *http.Response, error)
	SetCustomsDeclarationMetadata(ctx context.Context, request ApiMarketplaceV3DbsOrdersMetaCustomsDeclarationPostRequest) (*http.Response, error)
	// Deprecated: use ConfirmOrders.
	ConfirmOrder(ctx context.Context, orderID int64) (*http.Response, error)
	// Deprecated: use DeliverOrders.
	DeliverOrder(ctx context.Context, orderID int64) (*http.Response, error)
	// Deprecated: use ReceiveOrders.
	ReceiveOrder(ctx context.Context, orderID int64, code DBSCode) (*http.Response, error)
	// Deprecated: use RejectOrders.
	RejectOrder(ctx context.Context, orderID int64, code DBSCode) (*http.Response, error)
	// Deprecated: use CancelOrders.
	CancelOrder(ctx context.Context, orderID int64) (*http.Response, error)
}

type dbsService struct {
	api *DBSAPIClient
}

func (s *dbsService) Orders(ctx context.Context, query DBSOrdersQuery) (*ApiV3DbsOrdersGet200Response, *http.Response, error) {
	resp, httpResp, err := s.api.DBSAssemblyOrdersAPI.
		ApiV3DbsOrdersGet(ctx).
		Limit(query.Limit).
		Next(query.Next).
		DateFrom(query.DateFrom).
		DateTo(query.DateTo).
		Execute()
	if err != nil {
		return nil, httpResp, transport.WrapResponseError(httpResp, err)
	}
	return resp, httpResp, nil
}

func (s *dbsService) NewOrders(ctx context.Context) (*ApiV3DbsOrdersNewGet200Response, *http.Response, error) {
	resp, httpResp, err := s.api.DBSAssemblyOrdersAPI.ApiV3DbsOrdersNewGet(ctx).Execute()
	if err != nil {
		return nil, httpResp, transport.WrapResponseError(httpResp, err)
	}
	return resp, httpResp, nil
}

func (s *dbsService) GroupsInfo(ctx context.Context, request ApiOrderGroupsRequest) ([]ApiOrderGroupInner, *http.Response, error) {
	resp, httpResp, err := s.api.DBSAssemblyOrdersAPI.
		ApiV3DbsGroupsInfoPost(ctx).
		ApiOrderGroupsRequest(request).
		Execute()
	if err != nil {
		return nil, httpResp, transport.WrapResponseError(httpResp, err)
	}
	return resp, httpResp, nil
}

func (s *dbsService) BuyerInfo(ctx context.Context, request ApiOrdersRequestV2) (*DbsOnlyClientInfoResp, *http.Response, error) {
	resp, httpResp, err := s.api.DBSAssemblyOrdersAPI.
		ApiV3DbsOrdersClientPost(ctx).
		OrdersRequestAPI(dbsOrdersRequestAPI(request)).
		Execute()
	if err != nil {
		return nil, httpResp, transport.WrapResponseError(httpResp, err)
	}
	return resp, httpResp, nil
}

func (s *dbsService) B2BInfo(ctx context.Context, request ApiOrdersRequestV2) (*ApiB2bClientInfoResponses, *http.Response, error) {
	resp, httpResp, err := s.api.DBSAssemblyOrdersAPI.
		ApiMarketplaceV3DbsOrdersB2bInfoPost(ctx).
		ApiOrdersRequestV2(request).
		Execute()
	if err != nil {
		return nil, httpResp, transport.WrapResponseError(httpResp, err)
	}
	return resp, httpResp, nil
}

func (s *dbsService) DeliveryDate(ctx context.Context, request DeliveryDatesRequest) (*DeliveryDatesInfoResp, *http.Response, error) {
	resp, httpResp, err := s.api.DBSAssemblyOrdersAPI.
		ApiV3DbsOrdersDeliveryDatePost(ctx).
		DeliveryDatesRequest(request).
		Execute()
	if err != nil {
		return nil, httpResp, transport.WrapResponseError(httpResp, err)
	}
	return resp, httpResp, nil
}

func (s *dbsService) Stickers(ctx context.Context, request ApiMarketplaceV3DbsOrdersStickersPostRequest, options *DBSStickerOptions) (*ApiMarketplaceV3DbsOrdersStickersPost200Response, *http.Response, error) {
	req := s.api.DBSAssemblyOrdersAPI.
		ApiMarketplaceV3DbsOrdersStickersPost(ctx).
		ApiMarketplaceV3DbsOrdersStickersPostRequest(request)
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

func (s *dbsService) OrdersStatusInfo(ctx context.Context, request ApiOrdersRequestV2) (*ApiOrderStatusesV2, *http.Response, error) {
	resp, httpResp, err := s.api.DBSAssemblyOrdersAPI.
		ApiMarketplaceV3DbsOrdersStatusInfoPost(ctx).
		ApiOrdersRequestV2(request).
		Execute()
	if err != nil {
		return nil, httpResp, transport.WrapResponseError(httpResp, err)
	}
	return resp, httpResp, nil
}

func (s *dbsService) ConfirmOrders(ctx context.Context, request ApiOrdersRequestV2) (*ApiStatusSetResponses, *http.Response, error) {
	resp, httpResp, err := s.api.DBSAssemblyOrdersAPI.
		ApiMarketplaceV3DbsOrdersStatusConfirmPost(ctx).
		ApiOrdersRequestV2(request).
		Execute()
	if err != nil {
		return nil, httpResp, transport.WrapResponseError(httpResp, err)
	}
	return resp, httpResp, nil
}

func (s *dbsService) DeliverOrders(ctx context.Context, request ApiOrdersRequestV2) (*ApiStatusSetResponses, *http.Response, error) {
	resp, httpResp, err := s.api.DBSAssemblyOrdersAPI.
		ApiMarketplaceV3DbsOrdersStatusDeliverPost(ctx).
		ApiOrdersRequestV2(request).
		Execute()
	if err != nil {
		return nil, httpResp, transport.WrapResponseError(httpResp, err)
	}
	return resp, httpResp, nil
}

func (s *dbsService) ReceiveOrders(ctx context.Context, request ApiOrdersCodeRequest) (*ApiMarketplaceV3DbsOrdersStatusReceivePost200Response, *http.Response, error) {
	resp, httpResp, err := s.api.DBSAssemblyOrdersAPI.
		ApiMarketplaceV3DbsOrdersStatusReceivePost(ctx).
		ApiOrdersCodeRequest(request).
		Execute()
	if err != nil {
		return nil, httpResp, transport.WrapResponseError(httpResp, err)
	}
	return resp, httpResp, nil
}

func (s *dbsService) RejectOrders(ctx context.Context, request ApiOrdersCodeRequest) (*ApiStatusSetResponses, *http.Response, error) {
	resp, httpResp, err := s.api.DBSAssemblyOrdersAPI.
		ApiMarketplaceV3DbsOrdersStatusRejectPost(ctx).
		ApiOrdersCodeRequest(request).
		Execute()
	if err != nil {
		return nil, httpResp, transport.WrapResponseError(httpResp, err)
	}
	return resp, httpResp, nil
}

func (s *dbsService) CancelOrders(ctx context.Context, request ApiOrdersRequestV2) (*ApiStatusSetResponses, *http.Response, error) {
	resp, httpResp, err := s.api.DBSAssemblyOrdersAPI.
		ApiMarketplaceV3DbsOrdersStatusCancelPost(ctx).
		ApiOrdersRequestV2(request).
		Execute()
	if err != nil {
		return nil, httpResp, transport.WrapResponseError(httpResp, err)
	}
	return resp, httpResp, nil
}

func (s *dbsService) MetadataInfo(ctx context.Context, request ApiOrdersRequestV2) (*ApiOrdersMetaResponse, *http.Response, error) {
	resp, httpResp, err := s.api.DBSMetadataAPI.
		ApiMarketplaceV3DbsOrdersMetaInfoPost(ctx).
		ApiOrdersRequestV2(request).
		Execute()
	if err != nil {
		return nil, httpResp, transport.WrapResponseError(httpResp, err)
	}
	return resp, httpResp, nil
}

func (s *dbsService) DeleteMetadata(ctx context.Context, request ApiOrdersMetaDeleteRequest) (*ApiStatusSetResponses, *http.Response, error) {
	resp, httpResp, err := s.api.DBSMetadataAPI.
		ApiMarketplaceV3DbsOrdersMetaDeletePost(ctx).
		ApiOrdersMetaDeleteRequest(request).
		Execute()
	if err != nil {
		return nil, httpResp, transport.WrapResponseError(httpResp, err)
	}
	return resp, httpResp, nil
}

func (s *dbsService) SetSGTINMetadata(ctx context.Context, request ApiOrdersSGTINsSetRequest) (*ApiStatusSetResponses, *http.Response, error) {
	resp, httpResp, err := s.api.DBSMetadataAPI.
		ApiMarketplaceV3DbsOrdersMetaSgtinPost(ctx).
		ApiOrdersSGTINsSetRequest(request).
		Execute()
	if err != nil {
		return nil, httpResp, transport.WrapResponseError(httpResp, err)
	}
	return resp, httpResp, nil
}

func (s *dbsService) SetUINMetadata(ctx context.Context, request ApiOrdersUINSetRequest) (*ApiStatusSetResponses, *http.Response, error) {
	resp, httpResp, err := s.api.DBSMetadataAPI.
		ApiMarketplaceV3DbsOrdersMetaUinPost(ctx).
		ApiOrdersUINSetRequest(request).
		Execute()
	if err != nil {
		return nil, httpResp, transport.WrapResponseError(httpResp, err)
	}
	return resp, httpResp, nil
}

func (s *dbsService) SetIMEIMetadata(ctx context.Context, request ApiOrdersIMEISetRequest) (*ApiStatusSetResponses, *http.Response, error) {
	resp, httpResp, err := s.api.DBSMetadataAPI.
		ApiMarketplaceV3DbsOrdersMetaImeiPost(ctx).
		ApiOrdersIMEISetRequest(request).
		Execute()
	if err != nil {
		return nil, httpResp, transport.WrapResponseError(httpResp, err)
	}
	return resp, httpResp, nil
}

func (s *dbsService) SetGTINMetadata(ctx context.Context, request ApiOrdersGTINSetRequest) (*ApiStatusSetResponses, *http.Response, error) {
	resp, httpResp, err := s.api.DBSMetadataAPI.
		ApiMarketplaceV3DbsOrdersMetaGtinPost(ctx).
		ApiOrdersGTINSetRequest(request).
		Execute()
	if err != nil {
		return nil, httpResp, transport.WrapResponseError(httpResp, err)
	}
	return resp, httpResp, nil
}

func (s *dbsService) SetCustomsDeclarationMetadata(ctx context.Context, request ApiMarketplaceV3DbsOrdersMetaCustomsDeclarationPostRequest) (*http.Response, error) {
	httpResp, err := s.api.DBSMetadataAPI.
		ApiMarketplaceV3DbsOrdersMetaCustomsDeclarationPost(ctx).
		ApiMarketplaceV3DbsOrdersMetaCustomsDeclarationPostRequest(request).
		Execute()
	if err != nil {
		return httpResp, transport.WrapResponseError(httpResp, err)
	}
	return httpResp, nil
}

func (s *dbsService) ConfirmOrder(ctx context.Context, orderID int64) (*http.Response, error) {
	request, err := dbsSingleOrderRequest(orderID)
	if err != nil {
		return nil, err
	}
	_, httpResp, err := s.ConfirmOrders(ctx, request)
	return httpResp, err
}

func (s *dbsService) DeliverOrder(ctx context.Context, orderID int64) (*http.Response, error) {
	request, err := dbsSingleOrderRequest(orderID)
	if err != nil {
		return nil, err
	}
	_, httpResp, err := s.DeliverOrders(ctx, request)
	return httpResp, err
}

func (s *dbsService) ReceiveOrder(ctx context.Context, orderID int64, code DBSCode) (*http.Response, error) {
	request, err := dbsSingleOrderCodeRequest(orderID, code)
	if err != nil {
		return nil, err
	}
	_, httpResp, err := s.ReceiveOrders(ctx, request)
	return httpResp, err
}

func (s *dbsService) RejectOrder(ctx context.Context, orderID int64, code DBSCode) (*http.Response, error) {
	request, err := dbsSingleOrderCodeRequest(orderID, code)
	if err != nil {
		return nil, err
	}
	_, httpResp, err := s.RejectOrders(ctx, request)
	return httpResp, err
}

func (s *dbsService) CancelOrder(ctx context.Context, orderID int64) (*http.Response, error) {
	request, err := dbsSingleOrderRequest(orderID)
	if err != nil {
		return nil, err
	}
	_, httpResp, err := s.CancelOrders(ctx, request)
	return httpResp, err
}

func dbsSingleOrderRequest(orderID int64) (ApiOrdersRequestV2, error) {
	if orderID < 1 || orderID > math.MaxInt32 {
		return ApiOrdersRequestV2{}, fmt.Errorf("dbs order ID %d is outside supported batch range", orderID)
	}
	return ApiOrdersRequestV2{OrdersIds: []int32{int32(orderID)}}, nil
}

func dbsSingleOrderCodeRequest(orderID int64, code DBSCode) (ApiOrdersCodeRequest, error) {
	if orderID < 1 || orderID > math.MaxInt32 {
		return ApiOrdersCodeRequest{}, fmt.Errorf("dbs order ID %d is outside supported batch range", orderID)
	}

	order := ApiOrderCodeRequest{}
	order.SetOrderId(int32(orderID))
	order.SetCode(code.GetCode())

	return ApiOrdersCodeRequest{Orders: []ApiOrderCodeRequest{order}}, nil
}

func dbsOrdersRequestAPI(request ApiOrdersRequestV2) DBSOrdersRequestAPI {
	return DBSOrdersRequestAPI{Orders: request.GetOrdersIds()}
}
