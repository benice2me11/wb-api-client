package client

import (
	"context"
	"net/http"
)

type OrdersFBWAcceptanceOptionsQuery struct {
	WarehouseID *int32
}

type OrdersFBWSuppliesQuery struct {
	Filters ModelsSuppliesFiltersRequest
	Limit   *int32
	Offset  *int32
}

type OrdersFBWSupplyQuery struct {
	IsPreorderID *bool
}

type OrdersFBWSupplyGoodsQuery struct {
	Limit        *int32
	Offset       *int32
	IsPreorderID *bool
}

// OrdersFBWService provides ergonomic access to Orders FBW operations.
type OrdersFBWService interface {
	AcceptanceOptions(ctx context.Context, goods []ModelsGood, query *OrdersFBWAcceptanceOptionsQuery) (*ModelsOptionsResultModel, *http.Response, error)
	Warehouses(ctx context.Context) ([]ModelsWarehousesResultItems, *http.Response, error)
	TransitTariffs(ctx context.Context) ([]ModelsTransitTariff, *http.Response, error)
	Supplies(ctx context.Context, query OrdersFBWSuppliesQuery) ([]ModelsSupply, *http.Response, error)
	Supply(ctx context.Context, id int32, query *OrdersFBWSupplyQuery) (*ModelsSupplyDetails, *http.Response, error)
	SupplyGoods(ctx context.Context, id int32, query *OrdersFBWSupplyGoodsQuery) ([]ModelsGoodInSupply, *http.Response, error)
	SupplyPackage(ctx context.Context, id int32) ([]ModelsBox, *http.Response, error)
}

type ordersFBWService struct {
	api *OrdersFBWAPIClient
}

func (s *ordersFBWService) AcceptanceOptions(ctx context.Context, goods []ModelsGood, query *OrdersFBWAcceptanceOptionsQuery) (*ModelsOptionsResultModel, *http.Response, error) {
	req := s.api.DefaultApi.ApiV1AcceptanceOptionsPost(ctx).ModelsGood(goods)
	if query != nil && query.WarehouseID != nil {
		req = req.WarehouseID(*query.WarehouseID)
	}
	return wrapFacadeResult(req.Execute())
}

func (s *ordersFBWService) Warehouses(ctx context.Context) ([]ModelsWarehousesResultItems, *http.Response, error) {
	return wrapFacadeResult(s.api.DefaultApi.ApiV1WarehousesGet(ctx).Execute())
}

func (s *ordersFBWService) TransitTariffs(ctx context.Context) ([]ModelsTransitTariff, *http.Response, error) {
	return wrapFacadeResult(s.api.DefaultApi.ApiV1TransitTariffsGet(ctx).Execute())
}

func (s *ordersFBWService) Supplies(ctx context.Context, query OrdersFBWSuppliesQuery) ([]ModelsSupply, *http.Response, error) {
	req := s.api.DefaultApi.ApiV1SuppliesPost(ctx).ModelsSuppliesFiltersRequest(query.Filters)
	if query.Limit != nil {
		req = req.Limit(*query.Limit)
	}
	if query.Offset != nil {
		req = req.Offset(*query.Offset)
	}
	return wrapFacadeResult(req.Execute())
}

func (s *ordersFBWService) Supply(ctx context.Context, id int32, query *OrdersFBWSupplyQuery) (*ModelsSupplyDetails, *http.Response, error) {
	req := s.api.DefaultApi.ApiV1SuppliesIDGet(ctx, id)
	if query != nil && query.IsPreorderID != nil {
		req = req.IsPreorderID(*query.IsPreorderID)
	}
	return wrapFacadeResult(req.Execute())
}

func (s *ordersFBWService) SupplyGoods(ctx context.Context, id int32, query *OrdersFBWSupplyGoodsQuery) ([]ModelsGoodInSupply, *http.Response, error) {
	req := s.api.DefaultApi.ApiV1SuppliesIDGoodsGet(ctx, id)
	if query != nil {
		if query.Limit != nil {
			req = req.Limit(*query.Limit)
		}
		if query.Offset != nil {
			req = req.Offset(*query.Offset)
		}
		if query.IsPreorderID != nil {
			req = req.IsPreorderID(*query.IsPreorderID)
		}
	}
	return wrapFacadeResult(req.Execute())
}

func (s *ordersFBWService) SupplyPackage(ctx context.Context, id int32) ([]ModelsBox, *http.Response, error) {
	return wrapFacadeResult(s.api.DefaultApi.ApiV1SuppliesIDPackageGet(ctx, id).Execute())
}
