package client

import (
	"context"
	"net/http"

	wbgeneral "github.com/benice2me11/wb-api-client/internal/generated/general"
	"github.com/benice2me11/wb-api-client/transport"
)

// GeneralService provides ergonomic access to General category operations.
type GeneralService interface {
	Raw() *wbgeneral.APIClient
	Ping(ctx context.Context) (*wbgeneral.PingGet200Response, *http.Response, error)
	SellerInfo(ctx context.Context) (*wbgeneral.ApiV1SellerInfoGet200Response, *http.Response, error)
}

type generalService struct {
	api *wbgeneral.APIClient
}

func (s *generalService) Raw() *wbgeneral.APIClient {
	return s.api
}

func (s *generalService) Ping(ctx context.Context) (*wbgeneral.PingGet200Response, *http.Response, error) {
	resp, httpResp, err := s.api.WBAPIConnectionCheckAPI.PingGet(ctx).Execute()
	if err != nil {
		return nil, httpResp, transport.WrapResponseError(httpResp, err)
	}
	return resp, httpResp, nil
}

func (s *generalService) SellerInfo(ctx context.Context) (*wbgeneral.ApiV1SellerInfoGet200Response, *http.Response, error) {
	resp, httpResp, err := s.api.SellerInformationAPI.ApiV1SellerInfoGet(ctx).Execute()
	if err != nil {
		return nil, httpResp, transport.WrapResponseError(httpResp, err)
	}
	return resp, httpResp, nil
}
