package transport

import "net/http"

// AuthTransport injects a static auth header into outbound requests.
type AuthTransport struct {
	next   http.RoundTripper
	header string
	value  string
}

// NewAuthTransport creates a header-injecting round tripper.
func NewAuthTransport(next http.RoundTripper, header, value string) *AuthTransport {
	if next == nil {
		next = http.DefaultTransport
	}
	return &AuthTransport{next: next, header: header, value: value}
}

func (t *AuthTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if req == nil {
		return t.next.RoundTrip(req)
	}
	if t.header == "" || t.value == "" {
		return t.next.RoundTrip(req)
	}

	clone := req.Clone(req.Context())
	if clone.Header.Get(t.header) == "" {
		clone.Header.Set(t.header, t.value)
	}
	return t.next.RoundTrip(clone)
}
