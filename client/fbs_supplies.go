package client

import (
	"context"
	"net/http"

	"github.com/benice2me11/wb-api-client/transport"
)

func (s *fbsService) Supplies(ctx context.Context, query FBSSuppliesQuery) (*ApiV3SuppliesGet200Response, *http.Response, error) {
	resp, httpResp, err := s.api.FBSSuppliesAPI.
		ApiV3SuppliesGet(ctx).
		Limit(query.Limit).
		Next(query.Next).
		Execute()
	if err != nil {
		return nil, httpResp, transport.WrapResponseError(httpResp, err)
	}
	return resp, httpResp, nil
}

func (s *fbsService) CreateSupply(ctx context.Context, request ApiV3SuppliesPostRequest) (*ApiV3SuppliesPost201Response, *http.Response, error) {
	resp, httpResp, err := s.api.FBSSuppliesAPI.
		ApiV3SuppliesPost(ctx).
		ApiV3SuppliesPostRequest(request).
		Execute()
	if err != nil {
		return nil, httpResp, transport.WrapResponseError(httpResp, err)
	}
	return resp, httpResp, nil
}

func (s *fbsService) Supply(ctx context.Context, supplyID string) (*Supply, *http.Response, error) {
	resp, httpResp, err := s.api.FBSSuppliesAPI.
		ApiV3SuppliesSupplyIdGet(ctx, supplyID).
		Execute()
	if err != nil {
		return nil, httpResp, transport.WrapResponseError(httpResp, err)
	}
	return resp, httpResp, nil
}

func (s *fbsService) DeleteSupply(ctx context.Context, supplyID string) (*http.Response, error) {
	httpResp, err := s.api.FBSSuppliesAPI.
		ApiV3SuppliesSupplyIdDelete(ctx, supplyID).
		Execute()
	if err != nil {
		return httpResp, transport.WrapResponseError(httpResp, err)
	}
	return httpResp, nil
}

func (s *fbsService) DeliverSupply(ctx context.Context, supplyID string) (*http.Response, error) {
	httpResp, err := s.api.FBSSuppliesAPI.
		ApiV3SuppliesSupplyIdDeliverPatch(ctx, supplyID).
		Execute()
	if err != nil {
		return httpResp, transport.WrapResponseError(httpResp, err)
	}
	return httpResp, nil
}

func (s *fbsService) SupplyBarcode(ctx context.Context, supplyID string, typ *string) (*ApiV3SuppliesSupplyIdBarcodeGet200Response, *http.Response, error) {
	req := s.api.FBSSuppliesAPI.ApiV3SuppliesSupplyIdBarcodeGet(ctx, supplyID)
	if typ != nil {
		req = req.Type_(*typ)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, transport.WrapResponseError(httpResp, err)
	}
	return resp, httpResp, nil
}

func (s *fbsService) SupplyOrderIDs(ctx context.Context, supplyID string) (*V3SupplyOrderIDsAPI, *http.Response, error) {
	resp, httpResp, err := s.api.FBSSuppliesAPI.
		ApiMarketplaceV3SuppliesSupplyIdOrderIdsGet(ctx, supplyID).
		Execute()
	if err != nil {
		return nil, httpResp, transport.WrapResponseError(httpResp, err)
	}
	return resp, httpResp, nil
}

func (s *fbsService) AddOrdersToSupply(ctx context.Context, supplyID string, request OrdersRequestAPI) (*http.Response, error) {
	httpResp, err := s.api.FBSSuppliesAPI.
		ApiMarketplaceV3SuppliesSupplyIdOrdersPatch(ctx, supplyID).
		ApiMarketplaceV3SuppliesSupplyIdOrdersPatchRequest(ApiMarketplaceV3SuppliesSupplyIdOrdersPatchRequest{
			Orders: request.GetOrders(),
		}).
		Execute()
	if err != nil {
		return httpResp, transport.WrapResponseError(httpResp, err)
	}
	return httpResp, nil
}
