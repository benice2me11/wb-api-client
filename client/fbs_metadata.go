package client

import (
	"context"
	"net/http"

	"github.com/benice2me11/wb-api-client/transport"
)

func (s *fbsService) Metadata(ctx context.Context, request V3GetMetaMultiRequest) (*V3OrdersMetaAPI, *http.Response, error) {
	resp, httpResp, err := s.api.FBSMetadataAPI.
		ApiMarketplaceV3OrdersMetaPost(ctx).
		V3GetMetaMultiRequest(request).
		Execute()
	if err != nil {
		return nil, httpResp, transport.WrapResponseError(httpResp, err)
	}
	return resp, httpResp, nil
}

func (s *fbsService) DeleteMetadata(ctx context.Context, orderID int64, key string) (*http.Response, error) {
	httpResp, err := s.api.FBSMetadataAPI.
		ApiV3OrdersOrderIdMetaDelete(ctx, orderID).
		Key(key).
		Execute()
	if err != nil {
		return httpResp, transport.WrapResponseError(httpResp, err)
	}
	return httpResp, nil
}

func (s *fbsService) SetSGTINMetadata(ctx context.Context, orderID int64, request ApiV3OrdersOrderIdMetaSgtinPutRequest) (*http.Response, error) {
	httpResp, err := s.api.FBSMetadataAPI.
		ApiV3OrdersOrderIdMetaSgtinPut(ctx, orderID).
		ApiV3OrdersOrderIdMetaSgtinPutRequest(request).
		Execute()
	if err != nil {
		return httpResp, transport.WrapResponseError(httpResp, err)
	}
	return httpResp, nil
}

func (s *fbsService) SetUINMetadata(ctx context.Context, orderID int64, request ApiV3OrdersOrderIdMetaUinPutRequest) (*http.Response, error) {
	httpResp, err := s.api.FBSMetadataAPI.
		ApiV3OrdersOrderIdMetaUinPut(ctx, orderID).
		ApiV3OrdersOrderIdMetaUinPutRequest(request).
		Execute()
	if err != nil {
		return httpResp, transport.WrapResponseError(httpResp, err)
	}
	return httpResp, nil
}

func (s *fbsService) SetIMEIMetadata(ctx context.Context, orderID int64, request ApiV3OrdersOrderIdMetaImeiPutRequest) (*http.Response, error) {
	httpResp, err := s.api.FBSMetadataAPI.
		ApiV3OrdersOrderIdMetaImeiPut(ctx, orderID).
		ApiV3OrdersOrderIdMetaImeiPutRequest(request).
		Execute()
	if err != nil {
		return httpResp, transport.WrapResponseError(httpResp, err)
	}
	return httpResp, nil
}

func (s *fbsService) SetGTINMetadata(ctx context.Context, orderID int64, request ApiV3OrdersOrderIdMetaGtinPutRequest) (*http.Response, error) {
	httpResp, err := s.api.FBSMetadataAPI.
		ApiV3OrdersOrderIdMetaGtinPut(ctx, orderID).
		ApiV3OrdersOrderIdMetaGtinPutRequest(request).
		Execute()
	if err != nil {
		return httpResp, transport.WrapResponseError(httpResp, err)
	}
	return httpResp, nil
}

func (s *fbsService) SetExpirationMetadata(ctx context.Context, orderID int64, request ApiV3OrdersOrderIdMetaExpirationPutRequest) (*http.Response, error) {
	httpResp, err := s.api.FBSMetadataAPI.
		ApiV3OrdersOrderIdMetaExpirationPut(ctx, orderID).
		ApiV3OrdersOrderIdMetaExpirationPutRequest(request).
		Execute()
	if err != nil {
		return httpResp, transport.WrapResponseError(httpResp, err)
	}
	return httpResp, nil
}

func (s *fbsService) SetCustomsDeclarationMetadata(ctx context.Context, orderID int64, request ApiMarketplaceV3OrdersOrderIdMetaCustomsDeclarationPutRequest) (*http.Response, error) {
	httpResp, err := s.api.FBSMetadataAPI.
		ApiMarketplaceV3OrdersOrderIdMetaCustomsDeclarationPut(ctx, orderID).
		ApiMarketplaceV3OrdersOrderIdMetaCustomsDeclarationPutRequest(request).
		Execute()
	if err != nil {
		return httpResp, transport.WrapResponseError(httpResp, err)
	}
	return httpResp, nil
}
