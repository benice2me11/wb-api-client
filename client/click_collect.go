package client

import (
	"context"
	"net/http"

	"github.com/benice2me11/wb-api-client/transport"
)

// ClickCollectOrdersQuery defines required filters for completed Click Collect orders.
type ClickCollectOrdersQuery struct {
	Limit    int32
	Next     int32
	DateFrom int32
	DateTo   int32
}

// ClickCollectService provides ergonomic access to In-Store Pickup orders operations.
type ClickCollectService interface {
	NewOrders(ctx context.Context) (*ClickCollectApiNewOrders, *http.Response, error)
	Orders(ctx context.Context, query ClickCollectOrdersQuery) (*ClickCollectApiOrders, *http.Response, error)
	BuyerInfo(ctx context.Context, request ClickCollectApiOrdersRequest) (*ClickCollectApiOrderClientInfoResp, *http.Response, error)
	CheckBuyerIdentity(ctx context.Context, request ClickCollectApiCheckIdentityRequest) (*ClickCollectApiCheckedIdentity, *http.Response, error)
	OrdersStatusInfo(ctx context.Context, request ClickCollectApiOrdersRequestV2) (*ClickCollectApiOrderStatusesV2, *http.Response, error)
	ConfirmOrders(ctx context.Context, request ClickCollectApiOrdersRequestV2) (*ClickCollectApiStatusSetResponses, *http.Response, error)
	PrepareOrders(ctx context.Context, request ClickCollectApiOrdersRequestV2) (*ClickCollectApiMetaDetailsResponse, *http.Response, error)
	ReceiveOrders(ctx context.Context, request ClickCollectApiOrdersRequestV2) (*ClickCollectApiStatusSetResponses, *http.Response, error)
	RejectOrders(ctx context.Context, request ClickCollectApiOrdersRequestV2) (*ClickCollectApiStatusSetResponses, *http.Response, error)
	CancelOrders(ctx context.Context, request ClickCollectApiOrdersRequestV2) (*ClickCollectApiStatusSetResponses, *http.Response, error)
	MetadataInfo(ctx context.Context, request ClickCollectApiOrdersRequestV2) (*ClickCollectApiOrdersMetaResponse, *http.Response, error)
	MetadataDetails(ctx context.Context, request ClickCollectApiOrdersRequestV2) (*ClickCollectApiOrdersMetaDetailsResponse, *http.Response, error)
	DeleteMetadata(ctx context.Context, request ClickCollectApiOrdersMetaDeleteRequest) (*ClickCollectApiOrdersResponses, *http.Response, error)
	SetSGTINMetadata(ctx context.Context, request ClickCollectApiOrdersSGTINsSetRequest) (*ClickCollectApiMetaSetResponses, *http.Response, error)
	SetUINMetadata(ctx context.Context, request ClickCollectApiOrdersUINSetRequest) (*ClickCollectApiMetaSetResponses, *http.Response, error)
	SetIMEIMetadata(ctx context.Context, request ClickCollectApiOrdersIMEISetRequest) (*ClickCollectApiMetaSetResponses, *http.Response, error)
	SetGTINMetadata(ctx context.Context, request ClickCollectApiOrdersGTINSetRequest) (*ClickCollectApiMetaSetResponses, *http.Response, error)
}

type clickCollectService struct {
	api *ClickCollectAPIClient
}

func (s *clickCollectService) NewOrders(ctx context.Context) (*ClickCollectApiNewOrders, *http.Response, error) {
	resp, httpResp, err := s.api.DefaultApi.ApiV3ClickCollectOrdersNewGet(ctx).Execute()
	if err != nil {
		return nil, httpResp, transport.WrapResponseError(httpResp, err)
	}
	return resp, httpResp, nil
}

func (s *clickCollectService) Orders(ctx context.Context, query ClickCollectOrdersQuery) (*ClickCollectApiOrders, *http.Response, error) {
	resp, httpResp, err := s.api.DefaultApi.
		ApiV3ClickCollectOrdersGet(ctx).
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

func (s *clickCollectService) BuyerInfo(ctx context.Context, request ClickCollectApiOrdersRequest) (*ClickCollectApiOrderClientInfoResp, *http.Response, error) {
	resp, httpResp, err := s.api.DefaultApi.
		ApiV3ClickCollectOrdersClientPost(ctx).
		ApiOrdersRequest(request).
		Execute()
	if err != nil {
		return nil, httpResp, transport.WrapResponseError(httpResp, err)
	}
	return resp, httpResp, nil
}

func (s *clickCollectService) CheckBuyerIdentity(ctx context.Context, request ClickCollectApiCheckIdentityRequest) (*ClickCollectApiCheckedIdentity, *http.Response, error) {
	resp, httpResp, err := s.api.DefaultApi.
		ApiV3ClickCollectOrdersClientIdentityPost(ctx).
		ApiCheckIdentityRequest(request).
		Execute()
	if err != nil {
		return nil, httpResp, transport.WrapResponseError(httpResp, err)
	}
	return resp, httpResp, nil
}

func (s *clickCollectService) OrdersStatusInfo(ctx context.Context, request ClickCollectApiOrdersRequestV2) (*ClickCollectApiOrderStatusesV2, *http.Response, error) {
	resp, httpResp, err := s.api.DefaultApi.
		ApiMarketplaceV3ClickCollectOrdersStatusInfoPost(ctx).
		ApiOrdersRequestV2(request).
		Execute()
	if err != nil {
		return nil, httpResp, transport.WrapResponseError(httpResp, err)
	}
	return resp, httpResp, nil
}

func (s *clickCollectService) ConfirmOrders(ctx context.Context, request ClickCollectApiOrdersRequestV2) (*ClickCollectApiStatusSetResponses, *http.Response, error) {
	resp, httpResp, err := s.api.DefaultApi.
		ApiMarketplaceV3ClickCollectOrdersStatusConfirmPost(ctx).
		ApiOrdersRequestV2(request).
		Execute()
	if err != nil {
		return nil, httpResp, transport.WrapResponseError(httpResp, err)
	}
	return resp, httpResp, nil
}

func (s *clickCollectService) PrepareOrders(ctx context.Context, request ClickCollectApiOrdersRequestV2) (*ClickCollectApiMetaDetailsResponse, *http.Response, error) {
	resp, httpResp, err := s.api.DefaultApi.
		ApiMarketplaceV3ClickCollectOrdersStatusPreparePost(ctx).
		ApiOrdersRequestV2(request).
		Execute()
	if err != nil {
		return nil, httpResp, transport.WrapResponseError(httpResp, err)
	}
	return resp, httpResp, nil
}

func (s *clickCollectService) ReceiveOrders(ctx context.Context, request ClickCollectApiOrdersRequestV2) (*ClickCollectApiStatusSetResponses, *http.Response, error) {
	resp, httpResp, err := s.api.DefaultApi.
		ApiMarketplaceV3ClickCollectOrdersStatusReceivePost(ctx).
		ApiOrdersRequestV2(request).
		Execute()
	if err != nil {
		return nil, httpResp, transport.WrapResponseError(httpResp, err)
	}
	return resp, httpResp, nil
}

func (s *clickCollectService) RejectOrders(ctx context.Context, request ClickCollectApiOrdersRequestV2) (*ClickCollectApiStatusSetResponses, *http.Response, error) {
	resp, httpResp, err := s.api.DefaultApi.
		ApiMarketplaceV3ClickCollectOrdersStatusRejectPost(ctx).
		ApiOrdersRequestV2(request).
		Execute()
	if err != nil {
		return nil, httpResp, transport.WrapResponseError(httpResp, err)
	}
	return resp, httpResp, nil
}

func (s *clickCollectService) CancelOrders(ctx context.Context, request ClickCollectApiOrdersRequestV2) (*ClickCollectApiStatusSetResponses, *http.Response, error) {
	resp, httpResp, err := s.api.DefaultApi.
		ApiMarketplaceV3ClickCollectOrdersStatusCancelPost(ctx).
		ApiOrdersRequestV2(request).
		Execute()
	if err != nil {
		return nil, httpResp, transport.WrapResponseError(httpResp, err)
	}
	return resp, httpResp, nil
}

func (s *clickCollectService) MetadataInfo(ctx context.Context, request ClickCollectApiOrdersRequestV2) (*ClickCollectApiOrdersMetaResponse, *http.Response, error) {
	resp, httpResp, err := s.api.DefaultApi.
		ApiMarketplaceV3ClickCollectOrdersMetaInfoPost(ctx).
		ApiOrdersRequestV2(request).
		Execute()
	if err != nil {
		return nil, httpResp, transport.WrapResponseError(httpResp, err)
	}
	return resp, httpResp, nil
}

func (s *clickCollectService) MetadataDetails(ctx context.Context, request ClickCollectApiOrdersRequestV2) (*ClickCollectApiOrdersMetaDetailsResponse, *http.Response, error) {
	resp, httpResp, err := s.api.DefaultApi.
		ApiMarketplaceV3ClickCollectOrdersMetaDetailsPost(ctx).
		ApiOrdersRequestV2(request).
		Execute()
	if err != nil {
		return nil, httpResp, transport.WrapResponseError(httpResp, err)
	}
	return resp, httpResp, nil
}

func (s *clickCollectService) DeleteMetadata(ctx context.Context, request ClickCollectApiOrdersMetaDeleteRequest) (*ClickCollectApiOrdersResponses, *http.Response, error) {
	resp, httpResp, err := s.api.DefaultApi.
		ApiMarketplaceV3ClickCollectOrdersMetaDeletePost(ctx).
		ApiOrdersMetaDeleteRequest(request).
		Execute()
	if err != nil {
		return nil, httpResp, transport.WrapResponseError(httpResp, err)
	}
	return resp, httpResp, nil
}

func (s *clickCollectService) SetSGTINMetadata(ctx context.Context, request ClickCollectApiOrdersSGTINsSetRequest) (*ClickCollectApiMetaSetResponses, *http.Response, error) {
	resp, httpResp, err := s.api.DefaultApi.
		ApiMarketplaceV3ClickCollectOrdersMetaSgtinPost(ctx).
		ApiOrdersSGTINsSetRequest(request).
		Execute()
	if err != nil {
		return nil, httpResp, transport.WrapResponseError(httpResp, err)
	}
	return resp, httpResp, nil
}

func (s *clickCollectService) SetUINMetadata(ctx context.Context, request ClickCollectApiOrdersUINSetRequest) (*ClickCollectApiMetaSetResponses, *http.Response, error) {
	resp, httpResp, err := s.api.DefaultApi.
		ApiMarketplaceV3ClickCollectOrdersMetaUinPost(ctx).
		ApiOrdersUINSetRequest(request).
		Execute()
	if err != nil {
		return nil, httpResp, transport.WrapResponseError(httpResp, err)
	}
	return resp, httpResp, nil
}

func (s *clickCollectService) SetIMEIMetadata(ctx context.Context, request ClickCollectApiOrdersIMEISetRequest) (*ClickCollectApiMetaSetResponses, *http.Response, error) {
	resp, httpResp, err := s.api.DefaultApi.
		ApiMarketplaceV3ClickCollectOrdersMetaImeiPost(ctx).
		ApiOrdersIMEISetRequest(request).
		Execute()
	if err != nil {
		return nil, httpResp, transport.WrapResponseError(httpResp, err)
	}
	return resp, httpResp, nil
}

func (s *clickCollectService) SetGTINMetadata(ctx context.Context, request ClickCollectApiOrdersGTINSetRequest) (*ClickCollectApiMetaSetResponses, *http.Response, error) {
	resp, httpResp, err := s.api.DefaultApi.
		ApiMarketplaceV3ClickCollectOrdersMetaGtinPost(ctx).
		ApiOrdersGTINSetRequest(request).
		Execute()
	if err != nil {
		return nil, httpResp, transport.WrapResponseError(httpResp, err)
	}
	return resp, httpResp, nil
}
