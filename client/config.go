package client

import (
	"net/http"
	"strings"

	"github.com/benice2me11/wb-api-client/transport"
)

const (
	defaultGeneralBaseURL  = "https://common-api.wildberries.ru"
	defaultProductsBaseURL = "https://content-api.wildberries.ru"
	defaultFBSBaseURL      = "https://marketplace-api.wildberries.ru"
	defaultDBWBaseURL      = "https://marketplace-api.wildberries.ru"
	defaultDBSBaseURL      = "https://marketplace-api.wildberries.ru"
)

// BaseURLs holds category-specific API hosts.
type BaseURLs struct {
	General  string
	Products string
	FBS      string
	DBW      string
	DBS      string
}

// Config configures API client behavior.
type Config struct {
	BaseURLs    BaseURLs
	Token       string
	TokenPrefix string
	AuthHeader  string
	HTTPClient  *http.Client
	RetryPolicy transport.RetryPolicy

	overrideGeneralBaseURL  bool
	overrideProductsBaseURL bool
	overrideFBSBaseURL      bool
	overrideDBWBaseURL      bool
	overrideDBSBaseURL      bool
}

// Option mutates client configuration.
type Option func(*Config)

func defaultConfig() Config {
	return Config{
		BaseURLs: BaseURLs{
			General:  defaultGeneralBaseURL,
			Products: defaultProductsBaseURL,
			FBS:      defaultFBSBaseURL,
			DBW:      defaultDBWBaseURL,
			DBS:      defaultDBSBaseURL,
		},
		TokenPrefix: "Bearer",
		AuthHeader:  "Authorization",
		RetryPolicy: transport.DefaultRetryPolicy(),
	}
}

// WithToken configures auth token value.
func WithToken(token string) Option {
	return func(c *Config) {
		c.Token = strings.TrimSpace(token)
	}
}

// WithTokenPrefix configures token prefix (e.g. Bearer).
func WithTokenPrefix(prefix string) Option {
	return func(c *Config) {
		c.TokenPrefix = strings.TrimSpace(prefix)
	}
}

// WithAuthHeader configures outbound auth header name.
func WithAuthHeader(header string) Option {
	return func(c *Config) {
		c.AuthHeader = strings.TrimSpace(header)
	}
}

// WithBaseURLs overrides all category base URLs.
func WithBaseURLs(baseURLs BaseURLs) Option {
	return func(c *Config) {
		if baseURLs.General != "" {
			c.BaseURLs.General = strings.TrimRight(baseURLs.General, "/")
			c.overrideGeneralBaseURL = true
		}
		if baseURLs.Products != "" {
			c.BaseURLs.Products = strings.TrimRight(baseURLs.Products, "/")
			c.overrideProductsBaseURL = true
		}
		if baseURLs.FBS != "" {
			c.BaseURLs.FBS = strings.TrimRight(baseURLs.FBS, "/")
			c.overrideFBSBaseURL = true
		}
		if baseURLs.DBW != "" {
			c.BaseURLs.DBW = strings.TrimRight(baseURLs.DBW, "/")
			c.overrideDBWBaseURL = true
		}
		if baseURLs.DBS != "" {
			c.BaseURLs.DBS = strings.TrimRight(baseURLs.DBS, "/")
			c.overrideDBSBaseURL = true
		}
	}
}

func WithGeneralBaseURL(url string) Option {
	return func(c *Config) {
		if url != "" {
			c.BaseURLs.General = strings.TrimRight(url, "/")
			c.overrideGeneralBaseURL = true
		}
	}
}

func WithProductsBaseURL(url string) Option {
	return func(c *Config) {
		if url != "" {
			c.BaseURLs.Products = strings.TrimRight(url, "/")
			c.overrideProductsBaseURL = true
		}
	}
}

func WithFBSBaseURL(url string) Option {
	return func(c *Config) {
		if url != "" {
			c.BaseURLs.FBS = strings.TrimRight(url, "/")
			c.overrideFBSBaseURL = true
		}
	}
}

func WithDBWBaseURL(url string) Option {
	return func(c *Config) {
		if url != "" {
			c.BaseURLs.DBW = strings.TrimRight(url, "/")
			c.overrideDBWBaseURL = true
		}
	}
}

func WithDBSBaseURL(url string) Option {
	return func(c *Config) {
		if url != "" {
			c.BaseURLs.DBS = strings.TrimRight(url, "/")
			c.overrideDBSBaseURL = true
		}
	}
}

// WithHTTPClient provides a custom HTTP client.
func WithHTTPClient(httpClient *http.Client) Option {
	return func(c *Config) {
		c.HTTPClient = httpClient
	}
}

// WithRetryPolicy configures retry policy.
func WithRetryPolicy(policy transport.RetryPolicy) Option {
	return func(c *Config) {
		c.RetryPolicy = policy
	}
}
