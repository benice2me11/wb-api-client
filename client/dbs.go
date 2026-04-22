package client

import (
	"context"
	"net/http"

	wbdbs "github.com/benice2me11/wb-api-client/internal/generated/dbs"
	"github.com/benice2me11/wb-api-client/transport"
)

// DBSOrdersQuery contains required filters for completed DBS orders list.
type DBSOrdersQuery struct {
	Limit    int32
	Next     int64
	DateFrom int32
	DateTo   int32
}

// DBSService provides ergonomic access to DBS Orders operations.
type DBSService interface {
	Raw() *wbdbs.APIClient
	Orders(ctx context.Context, query DBSOrdersQuery) (*wbdbs.ApiV3DbsOrdersGet200Response, *http.Response, error)
	NewOrders(ctx context.Context) (*wbdbs.ApiV3DbsOrdersNewGet200Response, *http.Response, error)
	OrdersStatusInfo(ctx context.Context, request wbdbs.ApiOrdersRequestV2) (*wbdbs.ApiOrderStatusesV2, *http.Response, error)
	ConfirmOrder(ctx context.Context, orderID int64) (*http.Response, error)
	DeliverOrder(ctx context.Context, orderID int64) (*http.Response, error)
	ReceiveOrder(ctx context.Context, orderID int64, code wbdbs.Code) (*http.Response, error)
	RejectOrder(ctx context.Context, orderID int64, code wbdbs.Code) (*http.Response, error)
	CancelOrder(ctx context.Context, orderID int64) (*http.Response, error)
}

type dbsService struct {
	api *wbdbs.APIClient
}

func (s *dbsService) Raw() *wbdbs.APIClient {
	return s.api
}

func (s *dbsService) Orders(ctx context.Context, query DBSOrdersQuery) (*wbdbs.ApiV3DbsOrdersGet200Response, *http.Response, error) {
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

func (s *dbsService) NewOrders(ctx context.Context) (*wbdbs.ApiV3DbsOrdersNewGet200Response, *http.Response, error) {
	resp, httpResp, err := s.api.DBSAssemblyOrdersAPI.ApiV3DbsOrdersNewGet(ctx).Execute()
	if err != nil {
		return nil, httpResp, transport.WrapResponseError(httpResp, err)
	}
	return resp, httpResp, nil
}

func (s *dbsService) OrdersStatusInfo(ctx context.Context, request wbdbs.ApiOrdersRequestV2) (*wbdbs.ApiOrderStatusesV2, *http.Response, error) {
	resp, httpResp, err := s.api.DBSAssemblyOrdersAPI.
		ApiMarketplaceV3DbsOrdersStatusInfoPost(ctx).
		ApiOrdersRequestV2(request).
		Execute()
	if err != nil {
		return nil, httpResp, transport.WrapResponseError(httpResp, err)
	}
	return resp, httpResp, nil
}

func (s *dbsService) ConfirmOrder(ctx context.Context, orderID int64) (*http.Response, error) {
	httpResp, err := s.api.DBSAssemblyOrdersAPI.
		ApiV3DbsOrdersOrderIdConfirmPatch(ctx, orderID).
		Execute()
	if err != nil {
		return httpResp, transport.WrapResponseError(httpResp, err)
	}
	return httpResp, nil
}

func (s *dbsService) DeliverOrder(ctx context.Context, orderID int64) (*http.Response, error) {
	httpResp, err := s.api.DBSAssemblyOrdersAPI.
		ApiV3DbsOrdersOrderIdDeliverPatch(ctx, orderID).
		Execute()
	if err != nil {
		return httpResp, transport.WrapResponseError(httpResp, err)
	}
	return httpResp, nil
}

func (s *dbsService) ReceiveOrder(ctx context.Context, orderID int64, code wbdbs.Code) (*http.Response, error) {
	httpResp, err := s.api.DBSAssemblyOrdersAPI.
		ApiV3DbsOrdersOrderIdReceivePatch(ctx, orderID).
		Code(code).
		Execute()
	if err != nil {
		return httpResp, transport.WrapResponseError(httpResp, err)
	}
	return httpResp, nil
}

func (s *dbsService) RejectOrder(ctx context.Context, orderID int64, code wbdbs.Code) (*http.Response, error) {
	httpResp, err := s.api.DBSAssemblyOrdersAPI.
		ApiV3DbsOrdersOrderIdRejectPatch(ctx, orderID).
		Code(code).
		Execute()
	if err != nil {
		return httpResp, transport.WrapResponseError(httpResp, err)
	}
	return httpResp, nil
}

func (s *dbsService) CancelOrder(ctx context.Context, orderID int64) (*http.Response, error) {
	httpResp, err := s.api.DBSAssemblyOrdersAPI.
		ApiV3DbsOrdersOrderIdCancelPatch(ctx, orderID).
		Execute()
	if err != nil {
		return httpResp, transport.WrapResponseError(httpResp, err)
	}
	return httpResp, nil
}
