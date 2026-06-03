package client

import (
	"context"
	"net/http"
	"os"
)

// AnalyticsService provides ergonomic access to Analytics operations.
type AnalyticsService interface {
	CSVReports(ctx context.Context, filterDownloadIDs []string) (*NmReportGetReportsResponse, *http.Response, error)
	CreateCSVReport(ctx context.Context, request ApiV2NmReportDownloadsPostRequest) (*NmReportCreateReportResponse, *http.Response, error)
	RetryCSVReport(ctx context.Context, request NmReportRetryReportRequest) (*NmReportRetryReportResponse, *http.Response, error)
	DownloadCSVReport(ctx context.Context, downloadID string) (*os.File, *http.Response, error)
	SalesFunnelProducts(ctx context.Context, request ProductsRequest) (*PostSalesFunnelProducts200Response, *http.Response, error)
	SalesFunnelProductsHistory(ctx context.Context, request ProductHistoryRequest) ([]ProductHistoryResponseInner, *http.Response, error)
	SalesFunnelGroupedHistory(ctx context.Context, request GroupedHistoryRequest) (*PostSalesFunnelGroupedHistory200Response, *http.Response, error)
	SearchReport(ctx context.Context, request MainRequest) (*ApiV2SearchReportReportPost200Response, *http.Response, error)
	SearchReportTableGroups(ctx context.Context, request TableGroupRequest) (*ApiV2SearchReportTableGroupsPost200Response, *http.Response, error)
	SearchReportTableDetails(ctx context.Context, request TableDetailsRequest) (*ApiV2SearchReportTableDetailsPost200Response, *http.Response, error)
	SearchReportProductSearchTexts(ctx context.Context, request ProductSearchTextsRequest) (*ApiV2SearchReportProductSearchTextsPost200Response, *http.Response, error)
	SearchReportProductOrders(ctx context.Context, request ProductOrdersRequest) (*ApiV2SearchReportProductOrdersPost200Response, *http.Response, error)
	StocksReportWBWarehouses(ctx context.Context, request InventoryRequest) (*PostV1StocksReportWbWarehouses200Response, *http.Response, error)
	StocksReportProductGroups(ctx context.Context, request TableGroupRequestSt) (*ApiV2StocksReportProductsGroupsPost200Response, *http.Response, error)
	StocksReportProducts(ctx context.Context, request TableProductRequest) (*ApiV2StocksReportProductsProductsPost200Response, *http.Response, error)
	StocksReportProductSizes(ctx context.Context, request CommonSizeFilters) (*ApiV2StocksReportProductsSizesPost200Response, *http.Response, error)
	StocksReportOffices(ctx context.Context, request CommonShippingOfficeFilters) (*ApiV2StocksReportOfficesPost200Response, *http.Response, error)
	ItemRating(ctx context.Context, request ItemRatingRequest) (*PostV1ItemRating200Response, *http.Response, error)
}

type analyticsService struct {
	api *AnalyticsAPIClient
}

func (s *analyticsService) CSVReports(ctx context.Context, filterDownloadIDs []string) (*NmReportGetReportsResponse, *http.Response, error) {
	req := s.api.CSVAPI.ApiV2NmReportDownloadsGet(ctx)
	if filterDownloadIDs != nil {
		req = req.FilterDownloadIds(filterDownloadIDs)
	}
	return wrapFacadeResult(req.Execute())
}

func (s *analyticsService) CreateCSVReport(ctx context.Context, request ApiV2NmReportDownloadsPostRequest) (*NmReportCreateReportResponse, *http.Response, error) {
	req := s.api.CSVAPI.ApiV2NmReportDownloadsPost(ctx).ApiV2NmReportDownloadsPostRequest(request)
	return wrapFacadeResult(req.Execute())
}

func (s *analyticsService) RetryCSVReport(ctx context.Context, request NmReportRetryReportRequest) (*NmReportRetryReportResponse, *http.Response, error) {
	req := s.api.CSVAPI.ApiV2NmReportDownloadsRetryPost(ctx).NmReportRetryReportRequest(request)
	return wrapFacadeResult(req.Execute())
}

func (s *analyticsService) DownloadCSVReport(ctx context.Context, downloadID string) (*os.File, *http.Response, error) {
	return wrapFacadeResult(s.api.CSVAPI.ApiV2NmReportDownloadsFileDownloadIdGet(ctx, downloadID).Execute())
}

func (s *analyticsService) SalesFunnelProducts(ctx context.Context, request ProductsRequest) (*PostSalesFunnelProducts200Response, *http.Response, error) {
	req := s.api.DefaultApi.PostSalesFunnelProducts(ctx).ProductsRequest(request)
	return wrapFacadeResult(req.Execute())
}

func (s *analyticsService) SalesFunnelProductsHistory(ctx context.Context, request ProductHistoryRequest) ([]ProductHistoryResponseInner, *http.Response, error) {
	req := s.api.DefaultApi.PostSalesFunnelProductsHistory(ctx).ProductHistoryRequest(request)
	return wrapFacadeResult(req.Execute())
}

func (s *analyticsService) SalesFunnelGroupedHistory(ctx context.Context, request GroupedHistoryRequest) (*PostSalesFunnelGroupedHistory200Response, *http.Response, error) {
	req := s.api.DefaultApi.PostSalesFunnelGroupedHistory(ctx).GroupedHistoryRequest(request)
	return wrapFacadeResult(req.Execute())
}

func (s *analyticsService) SearchReport(ctx context.Context, request MainRequest) (*ApiV2SearchReportReportPost200Response, *http.Response, error) {
	req := s.api.DefaultApi.ApiV2SearchReportReportPost(ctx).MainRequest(request)
	return wrapFacadeResult(req.Execute())
}

func (s *analyticsService) SearchReportTableGroups(ctx context.Context, request TableGroupRequest) (*ApiV2SearchReportTableGroupsPost200Response, *http.Response, error) {
	req := s.api.DefaultApi.ApiV2SearchReportTableGroupsPost(ctx).TableGroupRequest(request)
	return wrapFacadeResult(req.Execute())
}

func (s *analyticsService) SearchReportTableDetails(ctx context.Context, request TableDetailsRequest) (*ApiV2SearchReportTableDetailsPost200Response, *http.Response, error) {
	req := s.api.DefaultApi.ApiV2SearchReportTableDetailsPost(ctx).TableDetailsRequest(request)
	return wrapFacadeResult(req.Execute())
}

func (s *analyticsService) SearchReportProductSearchTexts(ctx context.Context, request ProductSearchTextsRequest) (*ApiV2SearchReportProductSearchTextsPost200Response, *http.Response, error) {
	req := s.api.DefaultApi.ApiV2SearchReportProductSearchTextsPost(ctx).ProductSearchTextsRequest(request)
	return wrapFacadeResult(req.Execute())
}

func (s *analyticsService) SearchReportProductOrders(ctx context.Context, request ProductOrdersRequest) (*ApiV2SearchReportProductOrdersPost200Response, *http.Response, error) {
	req := s.api.DefaultApi.ApiV2SearchReportProductOrdersPost(ctx).ProductOrdersRequest(request)
	return wrapFacadeResult(req.Execute())
}

func (s *analyticsService) StocksReportWBWarehouses(ctx context.Context, request InventoryRequest) (*PostV1StocksReportWbWarehouses200Response, *http.Response, error) {
	req := s.api.DefaultApi.PostV1StocksReportWbWarehouses(ctx).InventoryRequest(request)
	return wrapFacadeResult(req.Execute())
}

func (s *analyticsService) StocksReportProductGroups(ctx context.Context, request TableGroupRequestSt) (*ApiV2StocksReportProductsGroupsPost200Response, *http.Response, error) {
	req := s.api.DefaultApi.ApiV2StocksReportProductsGroupsPost(ctx).TableGroupRequestSt(request)
	return wrapFacadeResult(req.Execute())
}

func (s *analyticsService) StocksReportProducts(ctx context.Context, request TableProductRequest) (*ApiV2StocksReportProductsProductsPost200Response, *http.Response, error) {
	req := s.api.DefaultApi.ApiV2StocksReportProductsProductsPost(ctx).TableProductRequest(request)
	return wrapFacadeResult(req.Execute())
}

func (s *analyticsService) StocksReportProductSizes(ctx context.Context, request CommonSizeFilters) (*ApiV2StocksReportProductsSizesPost200Response, *http.Response, error) {
	req := s.api.DefaultApi.ApiV2StocksReportProductsSizesPost(ctx).Body(request)
	return wrapFacadeResult(req.Execute())
}

func (s *analyticsService) StocksReportOffices(ctx context.Context, request CommonShippingOfficeFilters) (*ApiV2StocksReportOfficesPost200Response, *http.Response, error) {
	req := s.api.DefaultApi.ApiV2StocksReportOfficesPost(ctx).Body(request)
	return wrapFacadeResult(req.Execute())
}

func (s *analyticsService) ItemRating(ctx context.Context, request ItemRatingRequest) (*PostV1ItemRating200Response, *http.Response, error) {
	req := s.api.DefaultApi.PostV1ItemRating(ctx).ItemRatingRequest(request)
	return wrapFacadeResult(req.Execute())
}
