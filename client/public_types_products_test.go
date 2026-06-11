package client_test

import (
	"testing"

	"github.com/benice2me11/wb-api-client/client"
)

func TestProductsPublicTypeAliases(t *testing.T) {
	t.Parallel()

	var _ client.ContentV3MediaSavePostRequest
	var _ client.ContentV3MediaFilePost200Response
	var _ client.ContentV2CardsLimitsGet200Response
	var _ client.ContentV2CardsLimitsGet200ResponseData
	var _ client.ContentV2CardsDeleteTrashPostRequest
	var _ client.ContentV2CardsDeleteTrashPost200Response
	var _ client.ContentV2CardsRecoverPost200Response
	var _ client.RequestPublicViewerPublicErrorsTableListV2
	var _ client.ResponsePublicViewerPublicErrorsTableListV2
	var _ client.SwaggerPublicErrorsCursorInput
	var _ client.SwaggerPublicErrorsOrderV2
	var _ client.ModelsErrorTableListPublicRespV2
	var _ client.ModelsErrorTableListPublicRespV2Item
	var _ client.ViewerContractPublicErrorsCursorOutput
	var _ client.ApiV3StocksWarehouseIdDeleteRequest
	var _ client.ApiV3StocksWarehouseIdPutRequestStocksInner
	var _ client.ApiV3StocksWarehouseIdPost200ResponseStocksInner
	var _ client.StocksWarehouseErrorInner
	var _ client.StocksWarehouseErrorInnerDataInner
	var _ client.ApiV3StocksWarehouseIdPut406Response
	var _ client.Office
	var _ client.ApiV3WarehousesPostRequest
	var _ client.ApiV3WarehousesPost201Response
	var _ client.ApiV3WarehousesWarehouseIdPutRequest
	var _ client.StoreContactRequestBody
	var _ client.StoreContactRequestBodyContactsInner
	var _ client.ApiV3DbwWarehousesWarehouseIdContactsGet200Response
	var _ client.ApiV3DbwWarehousesWarehouseIdContactsGet200ResponseContactsInner
}
