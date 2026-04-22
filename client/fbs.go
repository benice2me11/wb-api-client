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

// FBSService provides ergonomic access to FBS Orders operations.
type FBSService interface {
	Orders(ctx context.Context, query *FBSOrdersQuery) (*ApiV3OrdersGet200Response, *http.Response, error)
	NewOrders(ctx context.Context) (*ApiV3OrdersNewGet200Response, *http.Response, error)
	OrdersByStatus(ctx context.Context, request ApiV3OrdersStatusPostRequest) (*ApiV3OrdersStatusPost200Response, *http.Response, error)
	OrdersStatusHistory(ctx context.Context, request ApiV3OrdersStatusHistoryPostRequest) (*ApiV3OrdersStatusHistoryPost200Response, *http.Response, error)
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

func (s *fbsService) CancelOrder(ctx context.Context, orderID int64) (*http.Response, error) {
	httpResp, err := s.api.FBSAssemblyOrdersAPI.
		ApiV3OrdersOrderIdCancelPatch(ctx, orderID).
		Execute()
	if err != nil {
		return httpResp, transport.WrapResponseError(httpResp, err)
	}
	return httpResp, nil
}
