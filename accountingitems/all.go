package accountingitems

import (
	"time"

	"github.com/mobmax/go-mews/json"
)

const (
	endpointAll = "accountingItems/getAll"

	ServiceRenue      AccountingItemType = "ServiceRevenue"
	ProductRevenue    AccountingItemType = "ProductRevenue"
	AdditionalRevenue AccountingItemType = "AdditionalRevenue"
	Payment           AccountingItemType = "Payment"
)

// List all products
func (s *Service) All(requestBody *AllRequest) (*AllResponse, error) {
	// @TODO: create wrapper?
	if err := s.Client.CheckTokens(); err != nil {
		return nil, err
	}

	apiURL, err := s.Client.GetApiURL(endpointAll)
	if err != nil {
		return nil, err
	}

	responseBody := &AllResponse{}
	httpReq, err := s.Client.NewRequest(apiURL, requestBody)
	if err != nil {
		return nil, err
	}

	_, err = s.Client.Do(httpReq, responseBody)
	return responseBody, err
}

type AllResponse struct {
	AccountingItems        []AccountingItem
	CreditCardTransactions CreditCardTransactions
}

func (s *Service) NewAllRequest() *AllRequest {
	return &AllRequest{
		Extent: AccountingItemsExtent{

			AccountingItems:        true,
			CreditCardTransactions: false,
		},
	}
}

type AllRequest struct {
	json.BaseRequest
	StartUTC   *time.Time                `json:"StartUtc,omitempty"`
	EndUTC     *time.Time                `json:"EndUtc,omitempty"`
	TimeFilter AccountingItemsTimeFilter `json:"TimeFilter,omitempty"`
	Currency   string                    `json:"Currency,omitempty"`
	Extent     AccountingItemsExtent     `json:"Extent,omitempty"`
}

type AccountingItemsTimeFilter string

const (
	TimeFilterClosed  AccountingItemsTimeFilter = "Closed"
	TimeFilterUpdated AccountingItemsTimeFilter = "Updated"
)

// 	"AccountingCategoryId": "4ac8ce68-5732-4f1d-bf0d-e557072c926f",
// 	"Amount": {
// 		"Currency": "GBP",
// 		"Tax": 0.42,
// 		"TaxRate": 0.2,
// 		"Value": 2.5
// 	},
// 	"BillId": null,
// 	"ConsumptionUtc": "2016-07-27T12:48:39Z",
// 	"Id": "89b93f7c-5c63-4de2-bd17-ec5fee5e3120",
// 	"Name": "Caramel, Pepper & Chilli Popcorn",
// 	"Notes": null,
// 	"OrderId": "810b8c3a-d358-4378-84a9-534c830016fc",
// 	"ProductId": null,
// 	"Type": "ServiceRevenue"
// }

type AccountingItems []AccountingItem

type AccountingItem struct {
	ID                   string             `json:"Id"`                     // Unique identifier of the item.
	CustomerID           string             `json:"CustomerId"`             // Unique identifier of the Customer whose account the item belongs to.
	ProductID            string             `json:"ProductId"`              // Unique identifier of the Product.
	ServiceID            string             `json:"ServiceId"`              // Unique identifier of the Service the item belongs to.
	OrderID              string             `json:"OrderId"`                // Unique identifier of the order (or Reservation) the item belongs to.
	BillID               string             `json:"BillId"`                 // Unique identifier of the bill the item is assigned to.
	CreditCardID         string             `json:"CreditCardId,omitempty"` // Unique identifier of the Credit card the item is associated to.
	InvoiceID            string             `json:"InvoiceId"`              // Unique identifier of the invoiced Bill the item is receivable for.
	AccountingCategoryID string             `json:"AccountingCategoryId"`   // Unique identifier of the Accounting Category the item belongs to.
	Amount               Amount             `json:"Amount"`                 // Amount the item costs, negative amount represents either rebate or a payment.
	Type                 AccountingItemType `json:"Type"`                   // Type of the item.
	Name                 string             `json:"Name"`                   // Name of the item.
	Notes                string             `json:"Notes"`                  // Additional notes.
	ConsumptionUTC       time.Time          `json:"ConsumptionUtc"`         // Date and time of the item consumption in UTC timezone in ISO 8601 format.
	ClosedUTC            time.Time          `json:"ClosedUtc"`              // Date and time of the item bill closure in UTC timezone in ISO 8601 format.
}

type CreditCardTransactions []CreditCardTransaction

type CreditCardTransaction struct {
	ID            string    `json:"Id"`
	PaymentID     string    `json:"PaymentId"`
	ChargedAmount Cost      `json:"ChargedAmount"`
	SettledAmount Cost      `json:"SettledAmount"`
	Fee           Cost      `json:"Fee"`
	SettlementID  string    `json:"SettlementId"`
	SettledUTC    time.Time `json:"SettledUtc"`
}

type Amount struct {
	Currency string   `json:"Currency"` // ISO-4217 code of the Currency.
	Net      float64  `json:"Net"`      // Net value in case the item is taxed.
	Tax      float64  `json:"Tax"`      // Tax value in case the item is taxed.
	TaxRate  *float64 `json:"TaxRate"`  // Tax rate in case the item is taxed (e.g. 0.21).
	Value    float64  `json:"Value"`    // Amount in the currency (including tax if taxed).
}

type AccountingItemType string

type AccountingItemsExtent struct {
	AccountingItems        bool `json:"AccountingItems"`
	CreditCardTransactions bool `json:"CreditCardTransactions"`
}

type Cost struct {
	Currency string   `json:"Currency"` // ISO-4217 code of the Currency.
	Net      float64  `json:"Net"`      // Net value in case the item is taxed.
	Tax      float64  `json:"Tax"`      // Tax value in case the item is taxed.
	TaxRate  *float64 `json:"TaxRate"`  // Tax rate in case the item is taxed (e.g. 0.21).
	Value    float64  `json:"Value"`    // Amount in the currency (including tax if taxed).
}
