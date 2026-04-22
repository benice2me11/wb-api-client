package client

import (
	"context"
	"net/http"

	"github.com/benice2me11/wb-api-client/transport"
)

// GeneralService provides ergonomic access to General category operations.
type GeneralService interface {
	Ping(ctx context.Context) (*PingGet200Response, *http.Response, error)
	SellerInfo(ctx context.Context) (*ApiV1SellerInfoGet200Response, *http.Response, error)
}

type generalService struct {
	api *GeneralAPIClient
}

func (s *generalService) Ping(ctx context.Context) (*PingGet200Response, *http.Response, error) {
	resp, httpResp, err := s.api.WBAPIConnectionCheckAPI.PingGet(ctx).Execute()
	if err != nil {
		return nil, httpResp, transport.WrapResponseError(httpResp, err)
	}
	return resp, httpResp, nil
}

func (s *generalService) SellerInfo(ctx context.Context) (*ApiV1SellerInfoGet200Response, *http.Response, error) {
	resp, httpResp, err := s.api.SellerInformationAPI.ApiV1SellerInfoGet(ctx).Execute()
	if err != nil {
		return nil, httpResp, transport.WrapResponseError(httpResp, err)
	}
	return resp, httpResp, nil
}
