package client

import (
	"net/http"
	"strings"

	wbanalytics "github.com/benice2me11/wb-api-client/internal/generated/analytics"
	wbdbs "github.com/benice2me11/wb-api-client/internal/generated/dbs"
	wbdbw "github.com/benice2me11/wb-api-client/internal/generated/dbw"
	wbfbs "github.com/benice2me11/wb-api-client/internal/generated/fbs"
	wbgeneral "github.com/benice2me11/wb-api-client/internal/generated/general"
	wbordersfbw "github.com/benice2me11/wb-api-client/internal/generated/ordersfbw"
	wbproducts "github.com/benice2me11/wb-api-client/internal/generated/products"
	wbreports "github.com/benice2me11/wb-api-client/internal/generated/reports"
	"github.com/benice2me11/wb-api-client/transport"
)

// Client is the public SDK facade over generated category clients.
type Client struct {
	cfg       Config
	general   GeneralService
	products  ProductsService
	fbs       FBSService
	dbw       DBWService
	dbs       DBSService
	reports   ReportsService
	analytics AnalyticsService
	ordersFBW OrdersFBWService
}

// NewClient builds a client with generated SDKs + transport middleware.
func NewClient(opts ...Option) *Client {
	cfg := defaultConfig()
	for _, opt := range opts {
		opt(&cfg)
	}

	httpClient := buildHTTPClient(cfg)

	generalCfg := wbgeneral.NewConfiguration()
	generalCfg.HTTPClient = httpClient
	if cfg.overrideGeneralBaseURL {
		overrideGeneralServers(generalCfg, cfg.BaseURLs.General)
	}

	productsCfg := wbproducts.NewConfiguration()
	productsCfg.HTTPClient = httpClient
	if cfg.overrideProductsBaseURL {
		overrideProductsServers(productsCfg, cfg.BaseURLs.Products)
	}

	fbsCfg := wbfbs.NewConfiguration()
	fbsCfg.HTTPClient = httpClient
	if cfg.overrideFBSBaseURL {
		overrideFBSServers(fbsCfg, cfg.BaseURLs.FBS)
	}

	dbwCfg := wbdbw.NewConfiguration()
	dbwCfg.HTTPClient = httpClient
	if cfg.overrideDBWBaseURL {
		overrideDBWServers(dbwCfg, cfg.BaseURLs.DBW)
	}

	dbsCfg := wbdbs.NewConfiguration()
	dbsCfg.HTTPClient = httpClient
	if cfg.overrideDBSBaseURL {
		overrideDBSServers(dbsCfg, cfg.BaseURLs.DBS)
	}

	reportsCfg := wbreports.NewConfiguration()
	reportsCfg.HTTPClient = httpClient
	if cfg.overrideReportsBaseURL {
		overrideReportsServers(reportsCfg, cfg.BaseURLs.Reports)
	}

	analyticsCfg := wbanalytics.NewConfiguration()
	analyticsCfg.HTTPClient = httpClient
	if cfg.overrideAnalyticsBaseURL {
		overrideAnalyticsServers(analyticsCfg, cfg.BaseURLs.Analytics)
	}

	ordersFBWCfg := wbordersfbw.NewConfiguration()
	ordersFBWCfg.HTTPClient = httpClient
	if cfg.overrideOrdersFBWBaseURL {
		overrideOrdersFBWServers(ordersFBWCfg, cfg.BaseURLs.OrdersFBW)
	}

	generalAPI := wbgeneral.NewAPIClient(generalCfg)
	productsAPI := wbproducts.NewAPIClient(productsCfg)
	fbsAPI := wbfbs.NewAPIClient(fbsCfg)
	dbwAPI := wbdbw.NewAPIClient(dbwCfg)
	dbsAPI := wbdbs.NewAPIClient(dbsCfg)
	reportsAPI := wbreports.NewAPIClient(reportsCfg)
	analyticsAPI := wbanalytics.NewAPIClient(analyticsCfg)
	ordersFBWAPI := wbordersfbw.NewAPIClient(ordersFBWCfg)

	return &Client{
		cfg:       cfg,
		general:   &generalService{api: generalAPI},
		products:  &productsService{api: productsAPI},
		fbs:       &fbsService{api: fbsAPI},
		dbw:       &dbwService{api: dbwAPI},
		dbs:       &dbsService{api: dbsAPI},
		reports:   &reportsService{api: reportsAPI},
		analytics: &analyticsService{api: analyticsAPI},
		ordersFBW: &ordersFBWService{api: ordersFBWAPI},
	}
}

func buildHTTPClient(cfg Config) *http.Client {
	baseClient := &http.Client{}
	if cfg.HTTPClient != nil {
		copyClient := *cfg.HTTPClient
		baseClient = &copyClient
	}

	baseTransport := baseClient.Transport
	if baseTransport == nil {
		baseTransport = http.DefaultTransport
	}

	authValue := formatAuthValue(cfg.TokenPrefix, cfg.Token)
	authTransport := transport.NewAuthTransport(baseTransport, cfg.AuthHeader, authValue)
	retryTransport := transport.NewRetryTransport(authTransport, cfg.RetryPolicy)
	baseClient.Transport = retryTransport
	return baseClient
}

func formatAuthValue(prefix, token string) string {
	token = strings.TrimSpace(token)
	if token == "" {
		return ""
	}
	prefix = strings.TrimSpace(prefix)
	if prefix == "" {
		return token
	}
	return prefix + " " + token
}

// General returns General category service facade.
func (c *Client) General() GeneralService {
	return c.general
}

// Products returns Product Management category service facade.
func (c *Client) Products() ProductsService {
	return c.products
}

// FBS returns FBS Orders category service facade.
func (c *Client) FBS() FBSService {
	return c.fbs
}

// DBW returns DBW Orders category service facade.
func (c *Client) DBW() DBWService {
	return c.dbw
}

// DBS returns DBS Orders category service facade.
func (c *Client) DBS() DBSService {
	return c.dbs
}

// Reports returns Reports category service facade.
func (c *Client) Reports() ReportsService {
	return c.reports
}

// Analytics returns Analytics category service facade.
func (c *Client) Analytics() AnalyticsService {
	return c.analytics
}

// OrdersFBW returns Orders FBW category service facade.
func (c *Client) OrdersFBW() OrdersFBWService {
	return c.ordersFBW
}

func overrideGeneralServers(cfg *wbgeneral.Configuration, baseURL string) {
	server := wbgeneral.ServerConfiguration{
		URL:         baseURL,
		Description: "overridden",
	}
	cfg.Servers = wbgeneral.ServerConfigurations{server}
	for key := range cfg.OperationServers {
		cfg.OperationServers[key] = wbgeneral.ServerConfigurations{server}
	}
}

func overrideProductsServers(cfg *wbproducts.Configuration, baseURL string) {
	server := wbproducts.ServerConfiguration{
		URL:         baseURL,
		Description: "overridden",
	}
	cfg.Servers = wbproducts.ServerConfigurations{server}
	for key := range cfg.OperationServers {
		cfg.OperationServers[key] = wbproducts.ServerConfigurations{server}
	}
}

func overrideFBSServers(cfg *wbfbs.Configuration, baseURL string) {
	server := wbfbs.ServerConfiguration{
		URL:         baseURL,
		Description: "overridden",
	}
	cfg.Servers = wbfbs.ServerConfigurations{server}
	for key := range cfg.OperationServers {
		cfg.OperationServers[key] = wbfbs.ServerConfigurations{server}
	}
}

func overrideDBWServers(cfg *wbdbw.Configuration, baseURL string) {
	server := wbdbw.ServerConfiguration{
		URL:         baseURL,
		Description: "overridden",
	}
	cfg.Servers = wbdbw.ServerConfigurations{server}
	for key := range cfg.OperationServers {
		cfg.OperationServers[key] = wbdbw.ServerConfigurations{server}
	}
}

func overrideDBSServers(cfg *wbdbs.Configuration, baseURL string) {
	server := wbdbs.ServerConfiguration{
		URL:         baseURL,
		Description: "overridden",
	}
	cfg.Servers = wbdbs.ServerConfigurations{server}
	for key := range cfg.OperationServers {
		cfg.OperationServers[key] = wbdbs.ServerConfigurations{server}
	}
}

func overrideReportsServers(cfg *wbreports.Configuration, baseURL string) {
	server := wbreports.ServerConfiguration{
		URL:         baseURL,
		Description: "overridden",
	}
	cfg.Servers = wbreports.ServerConfigurations{server}
	for key := range cfg.OperationServers {
		cfg.OperationServers[key] = wbreports.ServerConfigurations{server}
	}
}

func overrideAnalyticsServers(cfg *wbanalytics.Configuration, baseURL string) {
	server := wbanalytics.ServerConfiguration{
		URL:         baseURL,
		Description: "overridden",
	}
	cfg.Servers = wbanalytics.ServerConfigurations{server}
	for key := range cfg.OperationServers {
		cfg.OperationServers[key] = wbanalytics.ServerConfigurations{server}
	}
}

func overrideOrdersFBWServers(cfg *wbordersfbw.Configuration, baseURL string) {
	server := wbordersfbw.ServerConfiguration{
		URL:         baseURL,
		Description: "overridden",
	}
	cfg.Servers = wbordersfbw.ServerConfigurations{server}
	for key := range cfg.OperationServers {
		cfg.OperationServers[key] = wbordersfbw.ServerConfigurations{server}
	}
}
