package client_test

import (
	"testing"

	"github.com/benice2me11/wb-api-client/client"
)

func TestOrderPublicTypeAliases(t *testing.T) {
	t.Parallel()

	var _ client.OrdersRequestAPI
	var _ client.ApiV3OrdersStickersPostRequest
	var _ client.ApiV3SuppliesPostRequest
	var _ client.ApiV3PassesPostRequest
	var _ client.V3GetMetaMultiRequest
	var _ client.ApiV3OrdersOrderIdMetaSgtinPutRequest

	var _ client.DBWOrdersRequestAPI
	var _ client.DBWClientInfoResp
	var _ client.OrderCourierInfoResp
	var _ client.DBWDeliveryDatesRequest
	var _ client.ApiV3DbwOrdersStickersPostRequest
	var _ client.ApiV3DbwOrdersOrderIdMetaSgtinPutRequest

	var _ client.DBSOrdersRequestAPI
	var _ client.ApiOrdersCodeRequest
	var _ client.ApiOrderGroupsRequest
	var _ client.DbsOnlyClientInfoResp
	var _ client.DeliveryDatesRequest
	var _ client.ApiMarketplaceV3DbsOrdersStickersPostRequest
	var _ client.ApiOrdersSGTINsSetRequest

	var _ client.ClickCollectApiOrdersRequest
	var _ client.ClickCollectApiOrdersRequestV2
	var _ client.ClickCollectApiOrderClientInfoResp
	var _ client.ClickCollectApiCheckIdentityRequest
	var _ client.ClickCollectApiStatusSetResponses
	var _ client.ClickCollectApiOrdersMetaDetailsResponse
	var _ client.ClickCollectApiOrdersSGTINsSetRequest
}
