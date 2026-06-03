package client

import (
	"context"
	"net/http"
	"time"
)

type ReportsDateRangeQuery struct {
	DateFrom string
	DateTo   string
}

type ReportsSupplierStocksQuery struct {
	DateFrom time.Time
}

type ReportsSupplierOrdersQuery struct {
	DateFrom time.Time
	Flag     *int32
}

type ReportsWarehouseRemainsQuery struct {
	Locale         *string
	GroupByBrand   *bool
	GroupBySubject *bool
	GroupBySa      *bool
	GroupByNm      *bool
	GroupByBarcode *bool
	GroupBySize    *bool
	FilterPics     *int32
	FilterVolume   *int32
}

type ReportsBannedProductsQuery struct {
	Sort  *string
	Order *string
}

type ReportsBrandShareQuery struct {
	ParentID int32
	Brand    string
	DateFrom string
	DateTo   string
}

type ReportsBrandShareParentSubjectsQuery struct {
	Brand    string
	DateFrom string
	DateTo   string
	Locale   *string
}

type ReportsDeductionsQuery struct {
	DateTo   time.Time
	Limit    int32
	DateFrom *time.Time
	Sort     *string
	Order    *string
	Offset   *int32
}

type ReportsPagedTimeQuery struct {
	DateTo   time.Time
	Limit    int32
	DateFrom *time.Time
	Offset   *int32
}

// ReportsService provides ergonomic access to Reports operations.
type ReportsService interface {
	CreateWarehouseRemainsReport(ctx context.Context, query *ReportsWarehouseRemainsQuery) (*CreateTaskResponse, *http.Response, error)
	WarehouseRemainsReportStatus(ctx context.Context, taskID string) (*GetTasksResponse, *http.Response, error)
	DownloadWarehouseRemainsReport(ctx context.Context, taskID string) ([]ApiV1WarehouseRemainsTasksTaskIdDownloadGet200ResponseInner, *http.Response, error)
	SupplierStocks(ctx context.Context, query ReportsSupplierStocksQuery) ([]StocksItem, *http.Response, error)
	SupplierOrders(ctx context.Context, query ReportsSupplierOrdersQuery) ([]OrdersItem, *http.Response, error)
	SupplierSales(ctx context.Context, query ReportsSupplierOrdersQuery) ([]SalesItem, *http.Response, error)
	ExciseReport(ctx context.Context, query ReportsDateRangeQuery, request ExciseReportRequest) (*ExciseReportResponse, *http.Response, error)
	CreateAcceptanceReport(ctx context.Context, query ReportsDateRangeQuery) (*CreateTaskResponse, *http.Response, error)
	AcceptanceReportStatus(ctx context.Context, taskID string) (*GetTasksResponse, *http.Response, error)
	DownloadAcceptanceReport(ctx context.Context, taskID string) ([]ApiV1AcceptanceReportTasksTaskIdDownloadGet200ResponseInner, *http.Response, error)
	CreatePaidStorageReport(ctx context.Context, query ReportsDateRangeQuery) (*CreateTaskResponse, *http.Response, error)
	PaidStorageReportStatus(ctx context.Context, taskID string) (*GetTasksResponse, *http.Response, error)
	DownloadPaidStorageReport(ctx context.Context, taskID string) ([]ResponsePaidStorageInner, *http.Response, error)
	AntifraudDetails(ctx context.Context, date string) (*ApiV1AnalyticsAntifraudDetailsGet200Response, *http.Response, error)
	BannedProductsBlocked(ctx context.Context, query *ReportsBannedProductsQuery) (*ApiV1AnalyticsBannedProductsBlockedGet200Response, *http.Response, error)
	BannedProductsShadowed(ctx context.Context, query *ReportsBannedProductsQuery) (*ApiV1AnalyticsBannedProductsShadowedGet200Response, *http.Response, error)
	BrandShareBrands(ctx context.Context) (*ApiV1AnalyticsBrandShareBrandsGet200Response, *http.Response, error)
	BrandShare(ctx context.Context, query ReportsBrandShareQuery) (*ApiV1AnalyticsBrandShareGet200Response, *http.Response, error)
	BrandShareParentSubjects(ctx context.Context, query ReportsBrandShareParentSubjectsQuery) (*ApiV1AnalyticsBrandShareParentSubjectsGet200Response, *http.Response, error)
	GoodsLabeling(ctx context.Context, query ReportsDateRangeQuery) (*ApiV1AnalyticsGoodsLabelingGet200Response, *http.Response, error)
	GoodsReturn(ctx context.Context, query ReportsDateRangeQuery) (*ApiV1AnalyticsGoodsReturnGet200Response, *http.Response, error)
	RegionSale(ctx context.Context, query ReportsDateRangeQuery) (*ApiV1AnalyticsRegionSaleGet200Response, *http.Response, error)
	Deductions(ctx context.Context, query ReportsDeductionsQuery) (*GetDeductions200Response, *http.Response, error)
	MeasurementPenalties(ctx context.Context, query ReportsPagedTimeQuery) (*MeasurementPenalties, *http.Response, error)
	WarehouseMeasurements(ctx context.Context, query ReportsPagedTimeQuery) (*WHM, *http.Response, error)
}

type reportsService struct {
	api *ReportsAPIClient
}

func (s *reportsService) CreateWarehouseRemainsReport(ctx context.Context, query *ReportsWarehouseRemainsQuery) (*CreateTaskResponse, *http.Response, error) {
	req := s.api.DefaultApi.ApiV1WarehouseRemainsGet(ctx)
	if query != nil {
		if query.Locale != nil {
			req = req.Locale(*query.Locale)
		}
		if query.GroupByBrand != nil {
			req = req.GroupByBrand(*query.GroupByBrand)
		}
		if query.GroupBySubject != nil {
			req = req.GroupBySubject(*query.GroupBySubject)
		}
		if query.GroupBySa != nil {
			req = req.GroupBySa(*query.GroupBySa)
		}
		if query.GroupByNm != nil {
			req = req.GroupByNm(*query.GroupByNm)
		}
		if query.GroupByBarcode != nil {
			req = req.GroupByBarcode(*query.GroupByBarcode)
		}
		if query.GroupBySize != nil {
			req = req.GroupBySize(*query.GroupBySize)
		}
		if query.FilterPics != nil {
			req = req.FilterPics(*query.FilterPics)
		}
		if query.FilterVolume != nil {
			req = req.FilterVolume(*query.FilterVolume)
		}
	}
	return wrapFacadeResult(req.Execute())
}

func (s *reportsService) WarehouseRemainsReportStatus(ctx context.Context, taskID string) (*GetTasksResponse, *http.Response, error) {
	return wrapFacadeResult(s.api.DefaultApi.ApiV1WarehouseRemainsTasksTaskIdStatusGet(ctx, taskID).Execute())
}

func (s *reportsService) DownloadWarehouseRemainsReport(ctx context.Context, taskID string) ([]ApiV1WarehouseRemainsTasksTaskIdDownloadGet200ResponseInner, *http.Response, error) {
	return wrapFacadeResult(s.api.DefaultApi.ApiV1WarehouseRemainsTasksTaskIdDownloadGet(ctx, taskID).Execute())
}

func (s *reportsService) SupplierStocks(ctx context.Context, query ReportsSupplierStocksQuery) ([]StocksItem, *http.Response, error) {
	return wrapFacadeResult(s.api.DefaultApi.ApiV1SupplierStocksGet(ctx).DateFrom(query.DateFrom).Execute())
}

func (s *reportsService) SupplierOrders(ctx context.Context, query ReportsSupplierOrdersQuery) ([]OrdersItem, *http.Response, error) {
	req := s.api.DefaultApi.ApiV1SupplierOrdersGet(ctx).DateFrom(query.DateFrom)
	if query.Flag != nil {
		req = req.Flag(*query.Flag)
	}
	return wrapFacadeResult(req.Execute())
}

func (s *reportsService) SupplierSales(ctx context.Context, query ReportsSupplierOrdersQuery) ([]SalesItem, *http.Response, error) {
	req := s.api.DefaultApi.ApiV1SupplierSalesGet(ctx).DateFrom(query.DateFrom)
	if query.Flag != nil {
		req = req.Flag(*query.Flag)
	}
	return wrapFacadeResult(req.Execute())
}

func (s *reportsService) ExciseReport(ctx context.Context, query ReportsDateRangeQuery, request ExciseReportRequest) (*ExciseReportResponse, *http.Response, error) {
	req := s.api.CAPI.ApiV1AnalyticsExciseReportPost(ctx).
		DateFrom(query.DateFrom).
		DateTo(query.DateTo).
		ExciseReportRequest(request)
	return wrapFacadeResult(req.Execute())
}

func (s *reportsService) CreateAcceptanceReport(ctx context.Context, query ReportsDateRangeQuery) (*CreateTaskResponse, *http.Response, error) {
	req := s.api.DefaultApi.ApiV1AcceptanceReportGet(ctx).DateFrom(query.DateFrom).DateTo(query.DateTo)
	return wrapFacadeResult(req.Execute())
}

func (s *reportsService) AcceptanceReportStatus(ctx context.Context, taskID string) (*GetTasksResponse, *http.Response, error) {
	return wrapFacadeResult(s.api.DefaultApi.ApiV1AcceptanceReportTasksTaskIdStatusGet(ctx, taskID).Execute())
}

func (s *reportsService) DownloadAcceptanceReport(ctx context.Context, taskID string) ([]ApiV1AcceptanceReportTasksTaskIdDownloadGet200ResponseInner, *http.Response, error) {
	return wrapFacadeResult(s.api.DefaultApi.ApiV1AcceptanceReportTasksTaskIdDownloadGet(ctx, taskID).Execute())
}

func (s *reportsService) CreatePaidStorageReport(ctx context.Context, query ReportsDateRangeQuery) (*CreateTaskResponse, *http.Response, error) {
	req := s.api.DefaultApi.ApiV1PaidStorageGet(ctx).DateFrom(query.DateFrom).DateTo(query.DateTo)
	return wrapFacadeResult(req.Execute())
}

func (s *reportsService) PaidStorageReportStatus(ctx context.Context, taskID string) (*GetTasksResponse, *http.Response, error) {
	return wrapFacadeResult(s.api.DefaultApi.ApiV1PaidStorageTasksTaskIdStatusGet(ctx, taskID).Execute())
}

func (s *reportsService) DownloadPaidStorageReport(ctx context.Context, taskID string) ([]ResponsePaidStorageInner, *http.Response, error) {
	return wrapFacadeResult(s.api.DefaultApi.ApiV1PaidStorageTasksTaskIdDownloadGet(ctx, taskID).Execute())
}

func (s *reportsService) AntifraudDetails(ctx context.Context, date string) (*ApiV1AnalyticsAntifraudDetailsGet200Response, *http.Response, error) {
	return wrapFacadeResult(s.api.DefaultApi.ApiV1AnalyticsAntifraudDetailsGet(ctx).Date(date).Execute())
}

func (s *reportsService) BannedProductsBlocked(ctx context.Context, query *ReportsBannedProductsQuery) (*ApiV1AnalyticsBannedProductsBlockedGet200Response, *http.Response, error) {
	req := s.api.DefaultApi.ApiV1AnalyticsBannedProductsBlockedGet(ctx)
	if query != nil {
		if query.Sort != nil {
			req = req.Sort(*query.Sort)
		}
		if query.Order != nil {
			req = req.Order(*query.Order)
		}
	}
	return wrapFacadeResult(req.Execute())
}

func (s *reportsService) BannedProductsShadowed(ctx context.Context, query *ReportsBannedProductsQuery) (*ApiV1AnalyticsBannedProductsShadowedGet200Response, *http.Response, error) {
	req := s.api.DefaultApi.ApiV1AnalyticsBannedProductsShadowedGet(ctx)
	if query != nil {
		if query.Sort != nil {
			req = req.Sort(*query.Sort)
		}
		if query.Order != nil {
			req = req.Order(*query.Order)
		}
	}
	return wrapFacadeResult(req.Execute())
}

func (s *reportsService) BrandShareBrands(ctx context.Context) (*ApiV1AnalyticsBrandShareBrandsGet200Response, *http.Response, error) {
	return wrapFacadeResult(s.api.DefaultApi.ApiV1AnalyticsBrandShareBrandsGet(ctx).Execute())
}

func (s *reportsService) BrandShare(ctx context.Context, query ReportsBrandShareQuery) (*ApiV1AnalyticsBrandShareGet200Response, *http.Response, error) {
	req := s.api.DefaultApi.ApiV1AnalyticsBrandShareGet(ctx).
		ParentId(query.ParentID).
		Brand(query.Brand).
		DateFrom(query.DateFrom).
		DateTo(query.DateTo)
	return wrapFacadeResult(req.Execute())
}

func (s *reportsService) BrandShareParentSubjects(ctx context.Context, query ReportsBrandShareParentSubjectsQuery) (*ApiV1AnalyticsBrandShareParentSubjectsGet200Response, *http.Response, error) {
	req := s.api.DefaultApi.ApiV1AnalyticsBrandShareParentSubjectsGet(ctx).
		Brand(query.Brand).
		DateFrom(query.DateFrom).
		DateTo(query.DateTo)
	if query.Locale != nil {
		req = req.Locale(*query.Locale)
	}
	return wrapFacadeResult(req.Execute())
}

func (s *reportsService) GoodsLabeling(ctx context.Context, query ReportsDateRangeQuery) (*ApiV1AnalyticsGoodsLabelingGet200Response, *http.Response, error) {
	return wrapFacadeResult(s.api.DefaultApi.ApiV1AnalyticsGoodsLabelingGet(ctx).DateFrom(query.DateFrom).DateTo(query.DateTo).Execute())
}

func (s *reportsService) GoodsReturn(ctx context.Context, query ReportsDateRangeQuery) (*ApiV1AnalyticsGoodsReturnGet200Response, *http.Response, error) {
	return wrapFacadeResult(s.api.DefaultApi.ApiV1AnalyticsGoodsReturnGet(ctx).DateFrom(query.DateFrom).DateTo(query.DateTo).Execute())
}

func (s *reportsService) RegionSale(ctx context.Context, query ReportsDateRangeQuery) (*ApiV1AnalyticsRegionSaleGet200Response, *http.Response, error) {
	return wrapFacadeResult(s.api.DefaultApi.ApiV1AnalyticsRegionSaleGet(ctx).DateFrom(query.DateFrom).DateTo(query.DateTo).Execute())
}

func (s *reportsService) Deductions(ctx context.Context, query ReportsDeductionsQuery) (*GetDeductions200Response, *http.Response, error) {
	req := s.api.DefaultApi.GetDeductions(ctx).DateTo(query.DateTo).Limit(query.Limit)
	if query.DateFrom != nil {
		req = req.DateFrom(*query.DateFrom)
	}
	if query.Sort != nil {
		req = req.Sort(*query.Sort)
	}
	if query.Order != nil {
		req = req.Order(*query.Order)
	}
	if query.Offset != nil {
		req = req.Offset(*query.Offset)
	}
	return wrapFacadeResult(req.Execute())
}

func (s *reportsService) MeasurementPenalties(ctx context.Context, query ReportsPagedTimeQuery) (*MeasurementPenalties, *http.Response, error) {
	req := s.api.DefaultApi.GetMeasurementPenalties(ctx).DateTo(query.DateTo).Limit(query.Limit)
	if query.DateFrom != nil {
		req = req.DateFrom(*query.DateFrom)
	}
	if query.Offset != nil {
		req = req.Offset(*query.Offset)
	}
	return wrapFacadeResult(req.Execute())
}

func (s *reportsService) WarehouseMeasurements(ctx context.Context, query ReportsPagedTimeQuery) (*WHM, *http.Response, error) {
	req := s.api.DefaultApi.GetWarehouseMeasurements(ctx).DateTo(query.DateTo).Limit(query.Limit)
	if query.DateFrom != nil {
		req = req.DateFrom(*query.DateFrom)
	}
	if query.Offset != nil {
		req = req.Offset(*query.Offset)
	}
	return wrapFacadeResult(req.Execute())
}
