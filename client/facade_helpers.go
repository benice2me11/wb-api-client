package client

import (
	"net/http"

	"github.com/benice2me11/wb-api-client/transport"
)

func wrapFacadeResult[T any](value T, httpResp *http.Response, err error) (T, *http.Response, error) {
	if err != nil {
		var zero T
		return zero, httpResp, transport.WrapResponseError(httpResp, err)
	}
	return value, httpResp, nil
}
