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

// DBWStickerOptions defines optional sticker output query parameters.
type DBWStickerOptions struct {
	Type   *string
	Width  *int32
	Height *int32
}

// DBWService provides ergonomic access to DBW Orders operations.
type DBWService interface {
	Orders(ctx context.Context, query DBWOrdersQuery) (*ApiV3DbwOrdersGet200Response, *http.Response, error)
	NewOrders(ctx context.Context) (*ApiV3DbwOrdersNewGet200Response, *http.Response, error)
	OrdersByStatus(ctx context.Context, request ApiV3DbwOrdersStatusPostRequest) (*ApiV3DbwOrdersStatusPost200Response, *http.Response, error)
	BuyerInfo(ctx context.Context, request DBWOrdersRequestAPI) (*DBWClientInfoResp, *http.Response, error)
	CourierInfo(ctx context.Context, request DBWOrdersRequestAPI) (*OrderCourierInfoResp, *http.Response, error)
	DeliveryDate(ctx context.Context, request DBWDeliveryDatesRequest) (*DBWDeliveryDatesInfoResp, *http.Response, error)
	Stickers(ctx context.Context, request ApiV3DbwOrdersStickersPostRequest, options *DBWStickerOptions) (*ApiV3DbwOrdersStickersPost200Response, *http.Response, error)
	Metadata(ctx context.Context, orderID int64) (*ApiV3DbwOrdersOrderIdMetaGet200Response, *http.Response, error)
	DeleteMetadata(ctx context.Context, orderID int64, key string) (*http.Response, error)
	SetSGTINMetadata(ctx context.Context, orderID int64, request ApiV3DbwOrdersOrderIdMetaSgtinPutRequest) (*http.Response, error)
	SetUINMetadata(ctx context.Context, orderID int64, request ApiV3DbwOrdersOrderIdMetaUinPutRequest) (*http.Response, error)
	SetIMEIMetadata(ctx context.Context, orderID int64, request ApiV3DbwOrdersOrderIdMetaImeiPutRequest) (*http.Response, error)
	SetGTINMetadata(ctx context.Context, orderID int64, request ApiV3DbwOrdersOrderIdMetaGtinPutRequest) (*http.Response, error)
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

func (s *dbwService) BuyerInfo(ctx context.Context, request DBWOrdersRequestAPI) (*DBWClientInfoResp, *http.Response, error) {
	resp, httpResp, err := s.api.DBWAssemblyOrdersAPI.
		ApiMarketplaceV3DbwOrdersClientPost(ctx).
		OrdersRequestAPI(request).
		Execute()
	if err != nil {
		return nil, httpResp, transport.WrapResponseError(httpResp, err)
	}
	return resp, httpResp, nil
}

func (s *dbwService) CourierInfo(ctx context.Context, request DBWOrdersRequestAPI) (*OrderCourierInfoResp, *http.Response, error) {
	resp, httpResp, err := s.api.DBWAssemblyOrdersAPI.
		ApiV3DbwOrdersCourierPost(ctx).
		OrdersRequestAPI(request).
		Execute()
	if err != nil {
		return nil, httpResp, transport.WrapResponseError(httpResp, err)
	}
	return resp, httpResp, nil
}

func (s *dbwService) DeliveryDate(ctx context.Context, request DBWDeliveryDatesRequest) (*DBWDeliveryDatesInfoResp, *http.Response, error) {
	resp, httpResp, err := s.api.DBWAssemblyOrdersAPI.
		ApiV3DbwOrdersDeliveryDatePost(ctx).
		DeliveryDatesRequest(request).
		Execute()
	if err != nil {
		return nil, httpResp, transport.WrapResponseError(httpResp, err)
	}
	return resp, httpResp, nil
}

func (s *dbwService) Stickers(ctx context.Context, request ApiV3DbwOrdersStickersPostRequest, options *DBWStickerOptions) (*ApiV3DbwOrdersStickersPost200Response, *http.Response, error) {
	req := s.api.DBWAssemblyOrdersAPI.
		ApiV3DbwOrdersStickersPost(ctx).
		ApiV3DbwOrdersStickersPostRequest(request)
	if options != nil {
		if options.Type != nil {
			req = req.Type_(*options.Type)
		}
		if options.Width != nil {
			req = req.Width(*options.Width)
		}
		if options.Height != nil {
			req = req.Height(*options.Height)
		}
	}

	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, httpResp, transport.WrapResponseError(httpResp, err)
	}
	return resp, httpResp, nil
}

func (s *dbwService) Metadata(ctx context.Context, orderID int64) (*ApiV3DbwOrdersOrderIdMetaGet200Response, *http.Response, error) {
	resp, httpResp, err := s.api.DBWMetadataAPI.
		ApiV3DbwOrdersOrderIdMetaGet(ctx, orderID).
		Execute()
	if err != nil {
		return nil, httpResp, transport.WrapResponseError(httpResp, err)
	}
	return resp, httpResp, nil
}

func (s *dbwService) DeleteMetadata(ctx context.Context, orderID int64, key string) (*http.Response, error) {
	httpResp, err := s.api.DBWMetadataAPI.
		ApiV3DbwOrdersOrderIdMetaDelete(ctx, orderID).
		Key(key).
		Execute()
	if err != nil {
		return httpResp, transport.WrapResponseError(httpResp, err)
	}
	return httpResp, nil
}

func (s *dbwService) SetSGTINMetadata(ctx context.Context, orderID int64, request ApiV3DbwOrdersOrderIdMetaSgtinPutRequest) (*http.Response, error) {
	httpResp, err := s.api.DBWMetadataAPI.
		ApiV3DbwOrdersOrderIdMetaSgtinPut(ctx, orderID).
		ApiV3DbwOrdersOrderIdMetaSgtinPutRequest(request).
		Execute()
	if err != nil {
		return httpResp, transport.WrapResponseError(httpResp, err)
	}
	return httpResp, nil
}

func (s *dbwService) SetUINMetadata(ctx context.Context, orderID int64, request ApiV3DbwOrdersOrderIdMetaUinPutRequest) (*http.Response, error) {
	httpResp, err := s.api.DBWMetadataAPI.
		ApiV3DbwOrdersOrderIdMetaUinPut(ctx, orderID).
		ApiV3DbwOrdersOrderIdMetaUinPutRequest(request).
		Execute()
	if err != nil {
		return httpResp, transport.WrapResponseError(httpResp, err)
	}
	return httpResp, nil
}

func (s *dbwService) SetIMEIMetadata(ctx context.Context, orderID int64, request ApiV3DbwOrdersOrderIdMetaImeiPutRequest) (*http.Response, error) {
	httpResp, err := s.api.DBWMetadataAPI.
		ApiV3DbwOrdersOrderIdMetaImeiPut(ctx, orderID).
		ApiV3DbwOrdersOrderIdMetaImeiPutRequest(request).
		Execute()
	if err != nil {
		return httpResp, transport.WrapResponseError(httpResp, err)
	}
	return httpResp, nil
}

func (s *dbwService) SetGTINMetadata(ctx context.Context, orderID int64, request ApiV3DbwOrdersOrderIdMetaGtinPutRequest) (*http.Response, error) {
	httpResp, err := s.api.DBWMetadataAPI.
		ApiV3DbwOrdersOrderIdMetaGtinPut(ctx, orderID).
		ApiV3DbwOrdersOrderIdMetaGtinPutRequest(request).
		Execute()
	if err != nil {
		return httpResp, transport.WrapResponseError(httpResp, err)
	}
	return httpResp, nil
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
