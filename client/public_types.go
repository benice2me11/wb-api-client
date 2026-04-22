package client

import (
	wbdbs "github.com/benice2me11/wb-api-client/internal/generated/dbs"
	wbdbw "github.com/benice2me11/wb-api-client/internal/generated/dbw"
	wbfbs "github.com/benice2me11/wb-api-client/internal/generated/fbs"
	wbgeneral "github.com/benice2me11/wb-api-client/internal/generated/general"
	wbproducts "github.com/benice2me11/wb-api-client/internal/generated/products"
)

// Generated API clients (for internal facade wiring).
type GeneralAPIClient = wbgeneral.APIClient
type ProductsAPIClient = wbproducts.APIClient
type FBSAPIClient = wbfbs.APIClient
type DBWAPIClient = wbdbw.APIClient
type DBSAPIClient = wbdbs.APIClient

// General responses.
type PingGet200Response = wbgeneral.PingGet200Response
type ApiV1SellerInfoGet200Response = wbgeneral.ApiV1SellerInfoGet200Response

// Product cards requests/responses.
type ContentV2CardsUploadPostRequestInner = wbproducts.ContentV2CardsUploadPostRequestInner
type ContentV2CardsUploadAddPostRequest = wbproducts.ContentV2CardsUploadAddPostRequest
type ContentV2CardsUpdatePostRequestInner = wbproducts.ContentV2CardsUpdatePostRequestInner
type ContentV2GetCardsListPostRequest = wbproducts.ContentV2GetCardsListPostRequest
type ContentV2GetCardsListPostRequestSettings = wbproducts.ContentV2GetCardsListPostRequestSettings
type ContentV2GetCardsListPostRequestSettingsCursor = wbproducts.ContentV2GetCardsListPostRequestSettingsCursor
type ContentV2GetCardsListPostRequestSettingsFilter = wbproducts.ContentV2GetCardsListPostRequestSettingsFilter
type ContentV2GetCardsListPost200Response = wbproducts.ContentV2GetCardsListPost200Response
type ContentV2GetCardsListPost200ResponseCardsInner = wbproducts.ContentV2GetCardsListPost200ResponseCardsInner
type ContentV2GetCardsListPost200ResponseCardsInnerCharacteristicsInner = wbproducts.ContentV2GetCardsListPost200ResponseCardsInnerCharacteristicsInner
type ContentV2GetCardsListPost200ResponseCardsInnerPhotosInner = wbproducts.ContentV2GetCardsListPost200ResponseCardsInnerPhotosInner
type ContentV2GetCardsListPost200ResponseCardsInnerSizesInner = wbproducts.ContentV2GetCardsListPost200ResponseCardsInnerSizesInner
type ContentV2GetCardsListPost200ResponseCardsInnerTagsInner = wbproducts.ContentV2GetCardsListPost200ResponseCardsInnerTagsInner
type ResponseCardCreate = wbproducts.ResponseCardCreate

// Prices/discounts requests/responses.
type ApiV2UploadTaskPostRequest = wbproducts.ApiV2UploadTaskPostRequest
type ApiV2UploadTaskSizePostRequest = wbproducts.ApiV2UploadTaskSizePostRequest
type ApiV2HistoryGoodsTaskGet200Response = wbproducts.ApiV2HistoryGoodsTaskGet200Response
type ApiV2ListGoodsFilterGet200Response = wbproducts.ApiV2ListGoodsFilterGet200Response
type ApiV2ListGoodsFilterPostRequest = wbproducts.ApiV2ListGoodsFilterPostRequest
type ApiV2ListGoodsSizeNmGet200Response = wbproducts.ApiV2ListGoodsSizeNmGet200Response
type GoodsList = wbproducts.GoodsList
type TaskCreated = wbproducts.TaskCreated

// Warehouses/inventory requests/responses.
type ApiV3StocksWarehouseIdPostRequest = wbproducts.ApiV3StocksWarehouseIdPostRequest
type ApiV3StocksWarehouseIdPost200Response = wbproducts.ApiV3StocksWarehouseIdPost200Response
type ApiV3StocksWarehouseIdPutRequest = wbproducts.ApiV3StocksWarehouseIdPutRequest
type Warehouse = wbproducts.Warehouse

// FBS requests/responses.
type ApiV3OrdersGet200Response = wbfbs.ApiV3OrdersGet200Response
type ApiV3OrdersNewGet200Response = wbfbs.ApiV3OrdersNewGet200Response
type ApiV3OrdersStatusPostRequest = wbfbs.ApiV3OrdersStatusPostRequest
type ApiV3OrdersStatusPost200Response = wbfbs.ApiV3OrdersStatusPost200Response
type ApiV3OrdersStatusHistoryPostRequest = wbfbs.ApiV3OrdersStatusHistoryPostRequest
type ApiV3OrdersStatusHistoryPost200Response = wbfbs.ApiV3OrdersStatusHistoryPost200Response

// DBW requests/responses.
type ApiV3DbwOrdersGet200Response = wbdbw.ApiV3DbwOrdersGet200Response
type ApiV3DbwOrdersNewGet200Response = wbdbw.ApiV3DbwOrdersNewGet200Response
type ApiV3DbwOrdersStatusPostRequest = wbdbw.ApiV3DbwOrdersStatusPostRequest
type ApiV3DbwOrdersStatusPost200Response = wbdbw.ApiV3DbwOrdersStatusPost200Response

// DBS requests/responses.
type ApiV3DbsOrdersGet200Response = wbdbs.ApiV3DbsOrdersGet200Response
type ApiV3DbsOrdersNewGet200Response = wbdbs.ApiV3DbsOrdersNewGet200Response
type ApiOrdersRequestV2 = wbdbs.ApiOrdersRequestV2
type ApiOrderStatusesV2 = wbdbs.ApiOrderStatusesV2
type DBSCode = wbdbs.Code

// NewDBSCode builds DBS confirmation/rejection code value from a plain string.
func NewDBSCode(value string) DBSCode {
	code := wbdbs.NewCode()
	code.SetCode(value)
	return *code
}
