package client

import (
	"context"
	"net/http"

	"github.com/benice2me11/wb-api-client/transport"
)

func (s *fbsService) Passes(ctx context.Context) ([]Pass, *http.Response, error) {
	resp, httpResp, err := s.api.FBSPassesAPI.ApiV3PassesGet(ctx).Execute()
	if err != nil {
		return nil, httpResp, transport.WrapResponseError(httpResp, err)
	}
	return resp, httpResp, nil
}

func (s *fbsService) PassOffices(ctx context.Context) ([]PassOffice, *http.Response, error) {
	resp, httpResp, err := s.api.FBSPassesAPI.ApiV3PassesOfficesGet(ctx).Execute()
	if err != nil {
		return nil, httpResp, transport.WrapResponseError(httpResp, err)
	}
	return resp, httpResp, nil
}

func (s *fbsService) CreatePass(ctx context.Context, request ApiV3PassesPostRequest) (*ApiV3PassesPost201Response, *http.Response, error) {
	resp, httpResp, err := s.api.FBSPassesAPI.
		ApiV3PassesPost(ctx).
		ApiV3PassesPostRequest(request).
		Execute()
	if err != nil {
		return nil, httpResp, transport.WrapResponseError(httpResp, err)
	}
	return resp, httpResp, nil
}

func (s *fbsService) UpdatePass(ctx context.Context, passID int64, request ApiV3PassesPostRequest) (*http.Response, error) {
	httpResp, err := s.api.FBSPassesAPI.
		ApiV3PassesPassIdPut(ctx, passID).
		ApiV3PassesPostRequest(request).
		Execute()
	if err != nil {
		return httpResp, transport.WrapResponseError(httpResp, err)
	}
	return httpResp, nil
}

func (s *fbsService) DeletePass(ctx context.Context, passID int64) (*http.Response, error) {
	httpResp, err := s.api.FBSPassesAPI.
		ApiV3PassesPassIdDelete(ctx, passID).
		Execute()
	if err != nil {
		return httpResp, transport.WrapResponseError(httpResp, err)
	}
	return httpResp, nil
}
