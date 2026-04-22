package client

import (
	"context"
	"net/http"

	"github.com/benice2me11/wb-api-client/transport"
)

// DBWOrdersQuery contains required filters for completed DBW orders list.
type DBWOrdersQuery struct {
	Limit    int32
	Next     int64
	DateFrom int32
	DateTo   int32
}

// DBWService provides ergonomic access to DBW Orders operations.
type DBWService interface {
	Orders(ctx context.Context, query DBWOrdersQuery) (*ApiV3DbwOrdersGet200Response, *http.Response, error)
	NewOrders(ctx context.Context) (*ApiV3DbwOrdersNewGet200Response, *http.Response, error)
	OrdersByStatus(ctx context.Context, request ApiV3DbwOrdersStatusPostRequest) (*ApiV3DbwOrdersStatusPost200Response, *http.Response, error)
	ConfirmOrder(ctx context.Context, orderID int64) (*http.Response, error)
	AssembleOrder(ctx context.Context, orderID int64) (*http.Response, error)
	CancelOrder(ctx context.Context, orderID int64) (*http.Response, error)
}

type dbwService struct {
	api *DBWAPIClient
}

func (s *dbwService) Orders(ctx context.Context, query DBWOrdersQuery) (*ApiV3DbwOrdersGet200Response, *http.Response, error) {
	resp, httpResp, err := s.api.DBWAssemblyOrdersAPI.
		ApiV3DbwOrdersGet(ctx).
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

func (s *dbwService) NewOrders(ctx context.Context) (*ApiV3DbwOrdersNewGet200Response, *http.Response, error) {
	resp, httpResp, err := s.api.DBWAssemblyOrdersAPI.ApiV3DbwOrdersNewGet(ctx).Execute()
	if err != nil {
		return nil, httpResp, transport.WrapResponseError(httpResp, err)
	}
	return resp, httpResp, nil
}

func (s *dbwService) OrdersByStatus(ctx context.Context, request ApiV3DbwOrdersStatusPostRequest) (*ApiV3DbwOrdersStatusPost200Response, *http.Response, error) {
	resp, httpResp, err := s.api.DBWAssemblyOrdersAPI.
		ApiV3DbwOrdersStatusPost(ctx).
		ApiV3DbwOrdersStatusPostRequest(request).
		Execute()
	if err != nil {
		return nil, httpResp, transport.WrapResponseError(httpResp, err)
	}
	return resp, httpResp, nil
}

func (s *dbwService) ConfirmOrder(ctx context.Context, orderID int64) (*http.Response, error) {
	httpResp, err := s.api.DBWAssemblyOrdersAPI.
		ApiV3DbwOrdersOrderIdConfirmPatch(ctx, orderID).
		Execute()
	if err != nil {
		return httpResp, transport.WrapResponseError(httpResp, err)
	}
	return httpResp, nil
}

func (s *dbwService) AssembleOrder(ctx context.Context, orderID int64) (*http.Response, error) {
	httpResp, err := s.api.DBWAssemblyOrdersAPI.
		ApiV3DbwOrdersOrderIdAssemblePatch(ctx, orderID).
		Execute()
	if err != nil {
		return httpResp, transport.WrapResponseError(httpResp, err)
	}
	return httpResp, nil
}

func (s *dbwService) CancelOrder(ctx context.Context, orderID int64) (*http.Response, error) {
	httpResp, err := s.api.DBWAssemblyOrdersAPI.
		ApiV3DbwOrdersOrderIdCancelPatch(ctx, orderID).
		Execute()
	if err != nil {
		return httpResp, transport.WrapResponseError(httpResp, err)
	}
	return httpResp, nil
}
