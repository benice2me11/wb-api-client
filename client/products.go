package client

import (
	"context"
	"net/http"

	"github.com/benice2me11/wb-api-client/transport"
)

// PriceTaskDetailsQuery describes processed upload details query.
type PriceTaskDetailsQuery struct {
	Limit    int32
	UploadID int32
	Offset   *int32
}

// ListGoodsFilterQuery describes prices list query by optional WB article filter.
type ListGoodsFilterQuery struct {
	Limit      int32
	Offset     *int32
	FilterNmID *int32
}

// ListGoodsSizeByNmQuery describes query for size prices of one WB article.
type ListGoodsSizeByNmQuery struct {
	Limit  int32
	NmID   int32
	Offset *int32
}

// ProductsService provides ergonomic access to Product Management operations.
type ProductsService interface {
	CreateCards(ctx context.Context, cards []ContentV2CardsUploadPostRequestInner) (*ResponseCardCreate, *http.Response, error)
	AddCards(ctx context.Context, request ContentV2CardsUploadAddPostRequest) (*ResponseCardCreate, *http.Response, error)
	UpdateCards(ctx context.Context, cards []ContentV2CardsUpdatePostRequestInner) (*ResponseCardCreate, *http.Response, error)
	GetCardsList(ctx context.Context, request ContentV2GetCardsListPostRequest) (*ContentV2GetCardsListPost200Response, *http.Response, error)
	SetPrices(ctx context.Context, request ApiV2UploadTaskPostRequest) (*TaskCreated, *http.Response, error)
	SetSizePrices(ctx context.Context, request ApiV2UploadTaskSizePostRequest) (*TaskCreated, *http.Response, error)
	PriceTaskDetails(ctx context.Context, query PriceTaskDetailsQuery) (*ApiV2HistoryGoodsTaskGet200Response, *http.Response, error)
	ListGoodsFilter(ctx context.Context, query ListGoodsFilterQuery) (*ApiV2ListGoodsFilterGet200Response, *http.Response, error)
	ListGoodsFilterBulk(ctx context.Context, request ApiV2ListGoodsFilterPostRequest) (*ApiV2ListGoodsFilterGet200Response, *http.Response, error)
	ListGoodsSizeByNm(ctx context.Context, query ListGoodsSizeByNmQuery) (*ApiV2ListGoodsSizeNmGet200Response, *http.Response, error)
	GetInventory(ctx context.Context, warehouseID int64, request ApiV3StocksWarehouseIdPostRequest) (*ApiV3StocksWarehouseIdPost200Response, *http.Response, error)
	UpdateInventory(ctx context.Context, warehouseID int64, request ApiV3StocksWarehouseIdPutRequest) (*http.Response, error)
	Warehouses(ctx context.Context) ([]Warehouse, *http.Response, error)
}

type productsService struct {
	api *ProductsAPIClient
}

func (s *productsService) CreateCards(ctx context.Context, cards []ContentV2CardsUploadPostRequestInner) (*ResponseCardCreate, *http.Response, error) {
	resp, httpResp, err := s.api.CreatingProductCardsAPI.
		ContentV2CardsUploadPost(ctx).
		ContentV2CardsUploadPostRequestInner(cards).
		Execute()
	if err != nil {
		return nil, httpResp, transport.WrapResponseError(httpResp, err)
	}
	return resp, httpResp, nil
}

func (s *productsService) AddCards(ctx context.Context, request ContentV2CardsUploadAddPostRequest) (*ResponseCardCreate, *http.Response, error) {
	resp, httpResp, err := s.api.CreatingProductCardsAPI.
		ContentV2CardsUploadAddPost(ctx).
		ContentV2CardsUploadAddPostRequest(request).
		Execute()
	if err != nil {
		return nil, httpResp, transport.WrapResponseError(httpResp, err)
	}
	return resp, httpResp, nil
}

func (s *productsService) UpdateCards(ctx context.Context, cards []ContentV2CardsUpdatePostRequestInner) (*ResponseCardCreate, *http.Response, error) {
	resp, httpResp, err := s.api.ProductCardsAPI.
		ContentV2CardsUpdatePost(ctx).
		ContentV2CardsUpdatePostRequestInner(cards).
		Execute()
	if err != nil {
		return nil, httpResp, transport.WrapResponseError(httpResp, err)
	}
	return resp, httpResp, nil
}

func (s *productsService) GetCardsList(ctx context.Context, request ContentV2GetCardsListPostRequest) (*ContentV2GetCardsListPost200Response, *http.Response, error) {
	resp, httpResp, err := s.api.ProductCardsAPI.
		ContentV2GetCardsListPost(ctx).
		ContentV2GetCardsListPostRequest(request).
		Execute()
	if err != nil {
		return nil, httpResp, transport.WrapResponseError(httpResp, err)
	}
	return resp, httpResp, nil
}

func (s *productsService) SetPrices(ctx context.Context, request ApiV2UploadTaskPostRequest) (*TaskCreated, *http.Response, error) {
	resp, httpResp, err := s.api.PricesAndDiscountsAPI.
		ApiV2UploadTaskPost(ctx).
		ApiV2UploadTaskPostRequest(request).
		Execute()
	if err != nil {
		return nil, httpResp, transport.WrapResponseError(httpResp, err)
	}
	return resp, httpResp, nil
}

func (s *productsService) SetSizePrices(ctx context.Context, request ApiV2UploadTaskSizePostRequest) (*TaskCreated, *http.Response, error) {
	resp, httpResp, err := s.api.PricesAndDiscountsAPI.
		ApiV2UploadTaskSizePost(ctx).
		ApiV2UploadTaskSizePostRequest(request).
		Execute()
	if err != nil {
		return nil, httpResp, transport.WrapResponseError(httpResp, err)
	}
	return resp, httpResp, nil
}

func (s *productsService) PriceTaskDetails(ctx context.Context, query PriceTaskDetailsQuery) (*ApiV2HistoryGoodsTaskGet200Response, *http.Response, error) {
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

func (s *productsService) ListGoodsFilter(ctx context.Context, query ListGoodsFilterQuery) (*ApiV2ListGoodsFilterGet200Response, *http.Response, error) {
	req := s.api.PricesAndDiscountsAPI.
		ApiV2ListGoodsFilterGet(ctx).
		Limit(query.Limit)
	if query.Offset != nil {
		req = req.Offset(*query.Offset)
	}
	if query.FilterNmID != nil {
		req = req.FilterNmID(*query.FilterNmID)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, transport.WrapResponseError(httpResp, err)
	}
	return resp, httpResp, nil
}

func (s *productsService) ListGoodsFilterBulk(ctx context.Context, request ApiV2ListGoodsFilterPostRequest) (*ApiV2ListGoodsFilterGet200Response, *http.Response, error) {
	resp, httpResp, err := s.api.PricesAndDiscountsAPI.
		ApiV2ListGoodsFilterPost(ctx).
		ApiV2ListGoodsFilterPostRequest(request).
		Execute()
	if err != nil {
		return nil, httpResp, transport.WrapResponseError(httpResp, err)
	}
	return resp, httpResp, nil
}

func (s *productsService) ListGoodsSizeByNm(ctx context.Context, query ListGoodsSizeByNmQuery) (*ApiV2ListGoodsSizeNmGet200Response, *http.Response, error) {
	req := s.api.PricesAndDiscountsAPI.
		ApiV2ListGoodsSizeNmGet(ctx).
		Limit(query.Limit).
		NmID(query.NmID)
	if query.Offset != nil {
		req = req.Offset(*query.Offset)
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, transport.WrapResponseError(httpResp, err)
	}
	return resp, httpResp, nil
}

func (s *productsService) GetInventory(ctx context.Context, warehouseID int64, request ApiV3StocksWarehouseIdPostRequest) (*ApiV3StocksWarehouseIdPost200Response, *http.Response, error) {
	resp, httpResp, err := s.api.SellerWarehousesInventoryAPI.
		ApiV3StocksWarehouseIdPost(ctx, warehouseID).
		ApiV3StocksWarehouseIdPostRequest(request).
		Execute()
	if err != nil {
		return nil, httpResp, transport.WrapResponseError(httpResp, err)
	}
	return resp, httpResp, nil
}

func (s *productsService) UpdateInventory(ctx context.Context, warehouseID int64, request ApiV3StocksWarehouseIdPutRequest) (*http.Response, error) {
	httpResp, err := s.api.SellerWarehousesInventoryAPI.
		ApiV3StocksWarehouseIdPut(ctx, warehouseID).
		ApiV3StocksWarehouseIdPutRequest(request).
		Execute()
	if err != nil {
		return httpResp, transport.WrapResponseError(httpResp, err)
	}
	return httpResp, nil
}

func (s *productsService) Warehouses(ctx context.Context) ([]Warehouse, *http.Response, error) {
	resp, httpResp, err := s.api.SellerWarehousesAPI.ApiV3WarehousesGet(ctx).Execute()
	if err != nil {
		return nil, httpResp, transport.WrapResponseError(httpResp, err)
	}
	return resp, httpResp, nil
}
