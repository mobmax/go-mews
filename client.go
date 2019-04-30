package mews

import (
	"net/http"
	"net/url"

	"github.com/mobmax/go-mews/accountingcategories"
	"github.com/mobmax/go-mews/accountingitems"
	"github.com/mobmax/go-mews/bills"
	"github.com/mobmax/go-mews/businesssegments"
	"github.com/mobmax/go-mews/commands"
	"github.com/mobmax/go-mews/companies"
	"github.com/mobmax/go-mews/configuration"
	"github.com/mobmax/go-mews/customers"
	"github.com/mobmax/go-mews/json"
	"github.com/mobmax/go-mews/orders"
	"github.com/mobmax/go-mews/outletitems"
	"github.com/mobmax/go-mews/reservations"
	"github.com/mobmax/go-mews/services"
	"github.com/mobmax/go-mews/spaceblocks"
	"github.com/mobmax/go-mews/spaces"
	"github.com/mobmax/go-mews/tasks"
)

const (
	libraryVersion = "0.0.1"
	userAgent      = "go-mews/" + libraryVersion
)

var (
	BaseURL = &url.URL{
		Scheme: "https",
		Host:   "www.mews.li",
		Path:   "/api/connector/v1/",
	}
	BaseURLDemo = &url.URL{
		Scheme: "https",
		Host:   "demo.mews.li",
		Path:   "/api/connector/v1/",
	}
)

// NewClient returns a new MEWS API client
func NewClient(httpClient *http.Client, accessToken string, clientToken string) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	jsonClient := json.NewClient(httpClient, accessToken, clientToken)
	jsonClient.UserAgent = userAgent
	jsonClient.AccessToken = accessToken
	jsonClient.ClientToken = clientToken
	jsonClient.Debug = false

	c := &Client{
		client: jsonClient,
	}

	c.SetBaseURL(BaseURL)

	// Services
	c.AccountingItems = accountingitems.NewService()
	c.AccountingItems.Client = c.client
	c.OutletItems = outletitems.NewService()
	c.OutletItems.Client = c.client
	c.AccountingCategories = accountingcategories.NewService()
	c.AccountingCategories.Client = c.client
	c.Companies = companies.NewService()
	c.Companies.Client = c.client
	c.Customers = customers.NewService()
	c.Customers.Client = c.client
	c.Reservations = reservations.NewAPIService()
	c.Reservations.Client = c.client
	c.Spaces = spaces.NewService()
	c.Spaces.Client = c.client
	c.SpaceBlocks = spaceblocks.NewService()
	c.SpaceBlocks.Client = c.client
	c.Bills = bills.NewService()
	c.Bills.Client = c.client
	c.Commands = commands.NewService()
	c.Commands.Client = c.client
	c.Configuration = configuration.NewService()
	c.Configuration.Client = c.client
	c.BusinessSegments = businesssegments.NewService()
	c.BusinessSegments.Client = c.client
	c.Tasks = tasks.NewService()
	c.Tasks.Client = c.client
	c.Services = services.NewAPIService()
	c.Services.Client = c.client
	c.Orders = orders.NewService()
	c.Orders.Client = c.client

	return c
}

// Client manages communication with MEWS API
type Client struct {
	// HTTP client used to communicate with the API.
	client *json.Client

	// Services used for communicating with the API
	AccountingItems      *accountingitems.Service
	OutletItems          *outletitems.Service
	AccountingCategories *accountingcategories.Service
	Companies            *companies.Service
	Customers            *customers.Service
	Reservations         *reservations.APIService
	Spaces               *spaces.Service
	SpaceBlocks          *spaceblocks.Service
	Bills                *bills.Service
	Commands             *commands.Service
	Configuration        *configuration.Service
	BusinessSegments     *businesssegments.Service
	Tasks                *tasks.Service
	Services			 *services.APIService
	Orders				 *orders.Service
}

func (c *Client) SetDebug(debug bool) {
	c.client.Debug = debug
}

func (c *Client) SetBaseURL(baseURL *url.URL) {
	c.client.BaseURL = baseURL
}

func (c *Client) SetDisallowUnknownFields(disallowUnknownFields bool) {
	c.client.DisallowUnknownFields = disallowUnknownFields
}

func (c *Client) SetLanguageCode(code string) {
	c.client.SetLanguageCode(code)
}

func (c *Client) SetCultureCode(code string) {
	c.client.SetCultureCode(code)
}

func (c *Client) GetHost() string {
	return c.client.BaseURL.Host
}
