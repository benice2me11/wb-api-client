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
	var _ client.RequestPublicViewerPublicErrorsTableListV2
	var _ client.ResponsePublicViewerPublicErrorsTableListV2
	var _ client.SwaggerPublicErrorsCursorInput
	var _ client.SwaggerPublicErrorsOrderV2
	var _ client.ModelsErrorTableListPublicRespV2
	var _ client.ModelsErrorTableListPublicRespV2Item
	var _ client.ViewerContractPublicErrorsCursorOutput
}
