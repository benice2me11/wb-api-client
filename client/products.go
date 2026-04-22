package client

import (
	"context"
	"net/http"

	wbproducts "github.com/benice2me11/wb-api-client/internal/generated/products"
	"github.com/benice2me11/wb-api-client/transport"
)

// PriceTaskDetailsQuery describes processed upload details query.
type PriceTaskDetailsQuery struct {
	Limit    int32
	UploadID int32
	Offset   *int32
}

// ProductsService provides ergonomic access to Product Management operations.
type ProductsService interface {
	Raw() *wbproducts.APIClient
	CreateCards(ctx context.Context, cards []wbproducts.ContentV2CardsUploadPostRequestInner) (*wbproducts.ResponseCardCreate, *http.Response, error)
	AddCards(ctx context.Context, request wbproducts.ContentV2CardsUploadAddPostRequest) (*wbproducts.ResponseCardCreate, *http.Response, error)
	UpdateCards(ctx context.Context, cards []wbproducts.ContentV2CardsUpdatePostRequestInner) (*wbproducts.ResponseCardCreate, *http.Response, error)
	GetCardsList(ctx context.Context, request wbproducts.ContentV2GetCardsListPostRequest) (*wbproducts.ContentV2GetCardsListPost200Response, *http.Response, error)
	SetPrices(ctx context.Context, request wbproducts.ApiV2UploadTaskPostRequest) (*wbproducts.TaskCreated, *http.Response, error)
	SetSizePrices(ctx context.Context, request wbproducts.ApiV2UploadTaskSizePostRequest) (*wbproducts.TaskCreated, *http.Response, error)
	PriceTaskDetails(ctx context.Context, query PriceTaskDetailsQuery) (*wbproducts.ApiV2HistoryGoodsTaskGet200Response, *http.Response, error)
	GetInventory(ctx context.Context, warehouseID int64, request wbproducts.ApiV3StocksWarehouseIdPostRequest) (*wbproducts.ApiV3StocksWarehouseIdPost200Response, *http.Response, error)
	UpdateInventory(ctx context.Context, warehouseID int64, request wbproducts.ApiV3StocksWarehouseIdPutRequest) (*http.Response, error)
	Warehouses(ctx context.Context) ([]wbproducts.Warehouse, *http.Response, error)
}

type productsService struct {
	api *wbproducts.APIClient
}

func (s *productsService) Raw() *wbproducts.APIClient {
	return s.api
}

func (s *productsService) CreateCards(ctx context.Context, cards []wbproducts.ContentV2CardsUploadPostRequestInner) (*wbproducts.ResponseCardCreate, *http.Response, error) {
	resp, httpResp, err := s.api.CreatingProductCardsAPI.
		ContentV2CardsUploadPost(ctx).
		ContentV2CardsUploadPostRequestInner(cards).
		Execute()
	if err != nil {
		return nil, httpResp, transport.WrapResponseError(httpResp, err)
	}
	return resp, httpResp, nil
}

func (s *productsService) AddCards(ctx context.Context, request wbproducts.ContentV2CardsUploadAddPostRequest) (*wbproducts.ResponseCardCreate, *http.Response, error) {
	resp, httpResp, err := s.api.CreatingProductCardsAPI.
		ContentV2CardsUploadAddPost(ctx).
		ContentV2CardsUploadAddPostRequest(request).
		Execute()
	if err != nil {
		return nil, httpResp, transport.WrapResponseError(httpResp, err)
	}
	return resp, httpResp, nil
}

func (s *productsService) UpdateCards(ctx context.Context, cards []wbproducts.ContentV2CardsUpdatePostRequestInner) (*wbproducts.ResponseCardCreate, *http.Response, error) {
	resp, httpResp, err := s.api.ProductCardsAPI.
		ContentV2CardsUpdatePost(ctx).
		ContentV2CardsUpdatePostRequestInner(cards).
		Execute()
	if err != nil {
		return nil, httpResp, transport.WrapResponseError(httpResp, err)
	}
	return resp, httpResp, nil
}

func (s *productsService) GetCardsList(ctx context.Context, request wbproducts.ContentV2GetCardsListPostRequest) (*wbproducts.ContentV2GetCardsListPost200Response, *http.Response, error) {
	resp, httpResp, err := s.api.ProductCardsAPI.
		ContentV2GetCardsListPost(ctx).
		ContentV2GetCardsListPostRequest(request).
		Execute()
	if err != nil {
		return nil, httpResp, transport.WrapResponseError(httpResp, err)
	}
	return resp, httpResp, nil
}

func (s *productsService) SetPrices(ctx context.Context, request wbproducts.ApiV2UploadTaskPostRequest) (*wbproducts.TaskCreated, *http.Response, error) {
	resp, httpResp, err := s.api.PricesAndDiscountsAPI.
		ApiV2UploadTaskPost(ctx).
		ApiV2UploadTaskPostRequest(request).
		Execute()
	if err != nil {
		return nil, httpResp, transport.WrapResponseError(httpResp, err)
	}
	return resp, httpResp, nil
}

func (s *productsService) SetSizePrices(ctx context.Context, request wbproducts.ApiV2UploadTaskSizePostRequest) (*wbproducts.TaskCreated, *http.Response, error) {
	resp, httpResp, err := s.api.PricesAndDiscountsAPI.
		ApiV2UploadTaskSizePost(ctx).
		ApiV2UploadTaskSizePostRequest(request).
		Execute()
	if err != nil {
		return nil, httpResp, transport.WrapResponseError(httpResp, err)
	}
	return resp, httpResp, nil
}

func (s *productsService) PriceTaskDetails(ctx context.Context, query PriceTaskDetailsQuery) (*wbproducts.ApiV2HistoryGoodsTaskGet200Response, *http.Response, error) {
	req := s.api.PricesAndDiscountsAPI.
		ApiV2HistoryGoodsTaskGet(ctx).
		Limit(query.Limit).
		UploadID(query.UploadID)
	if query.Offset != nil {
		req = req.Offset(*query.Offset)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, transport.WrapResponseError(httpResp, err)
	}
	return resp, httpResp, nil
}

func (s *productsService) GetInventory(ctx context.Context, warehouseID int64, request wbproducts.ApiV3StocksWarehouseIdPostRequest) (*wbproducts.ApiV3StocksWarehouseIdPost200Response, *http.Response, error) {
	resp, httpResp, err := s.api.SellerWarehousesInventoryAPI.
		ApiV3StocksWarehouseIdPost(ctx, warehouseID).
		ApiV3StocksWarehouseIdPostRequest(request).
		Execute()
	if err != nil {
		return nil, httpResp, transport.WrapResponseError(httpResp, err)
	}
	return resp, httpResp, nil
}

func (s *productsService) UpdateInventory(ctx context.Context, warehouseID int64, request wbproducts.ApiV3StocksWarehouseIdPutRequest) (*http.Response, error) {
	httpResp, err := s.api.SellerWarehousesInventoryAPI.
		ApiV3StocksWarehouseIdPut(ctx, warehouseID).
		ApiV3StocksWarehouseIdPutRequest(request).
		Execute()
	if err != nil {
		return httpResp, transport.WrapResponseError(httpResp, err)
	}
	return httpResp, nil
}

func (s *productsService) Warehouses(ctx context.Context) ([]wbproducts.Warehouse, *http.Response, error) {
	resp, httpResp, err := s.api.SellerWarehousesAPI.ApiV3WarehousesGet(ctx).Execute()
	if err != nil {
		return nil, httpResp, transport.WrapResponseError(httpResp, err)
	}
	return resp, httpResp, nil
}
