package commands

import (
	"errors"
	"time"
)

const (
	endpointAllActive = "commands/getAllActive"
)

var (
	ErrNoToken = errors.New("No token specified")
)

// List all products
func (s *Service) AllActive(requestBody *AllActiveRequest) (*AllActiveResponse, error) {
	// @TODO: create wrapper?
	// Set request token
	requestBody.AccessToken = s.Client.AccessToken

	if s.Client.AccessToken == "" {
		return nil, ErrNoToken
	}

	apiURL, err := s.Client.GetApiURL(endpointAllActive)
	if err != nil {
		return nil, err
	}

	responseBody := &AllActiveResponse{}
	httpReq, err := s.Client.NewRequest(apiURL, requestBody)
	if err != nil {
		return nil, err
	}

	_, err = s.Client.Do(httpReq, responseBody)
	return responseBody, err
}

func (s *Service) NewAllActiveRequest() *AllActiveRequest {
	return &AllActiveRequest{}
}

type AllActiveRequest struct {
	AccessToken string `json:"AccessToken"`
}

type AllActiveResponse struct {
	Commands Commands `json:"Commands"` // The closed bills.
}

type Commands []Command

type Command struct {
	ID         string       `json:"Id"`         // Unique identifier of the command.
	CreatedUTC time.Time    `json:"CreatedUtc"` // Date and time of the command was created in UTC timezone in ISO 8601 format.
	Data       Data         `json:"Data"`       // Data send with the command from MEWS
	Device     Device       `json:"Device"`     // Device information
	State      CommandState `json:"State"`      // State of the command.
}

type CommandState string

const (
	CommandStatePending    CommandState = "Pending"
	CommandStateReceived   CommandState = "Received"
	CommandStateProcessing CommandState = "Processing"
	CommandStateProcessed  CommandState = "Processed"
	CommandStateCancelled  CommandState = "Cancelled"
	CommandStateError      CommandState = "Error"
)

type Data struct {
	__type          string `json:"__type,omitempty"`           // Type of command.
	Bill            Bill   `json:"Bill, omitempty"`            // If available add Bill informaion
	FiscalMachineID string `json:"FiscalMachineId, omitempty"` // Unique identifier of the Fiscal Machine.
	TaxIdentifier   string `json:"TaxIdentifier, omitempty"`   //Tax Identifier number
}

type Device struct {
	ID   string     `json:"Id"`   // Unique identifier of the Device to which the command is send
	Name string     `json:"Name"` // Name of the Device to which the command is send
	Type DeviceType `json:"Type"` //Type of Device
}

type DeviceType string

const (
	DevicePrinter         DeviceType = "Printer"
	DevicePaymentTerminal DeviceType = "PaymentTerminal"
	DevicePassportScanner DeviceType = "PassportScanner"
	DeviceFiscalMachine   DeviceType = "FiscalMachine"
	DeviceKeyCutter       DeviceType = "KeyCutter"
	DeviceVisiKeyCutter   DeviceType = "VisiOnlineKeyCutter"
)

type Bill struct {
	ID         string    `json:"Id"`         // Unique identifier of the bill.
	CustomerID string    `json:"CustomerId"` // Unique identifier of the Customer the bill is issued to.
	CompanyID  string    `json:"CompanyId"`  // Unique identifier of the Company the bill is issued to.
	Type       BillType  `json:"Type"`       // Type of the bill.
	Number     string    `json:"Number"`     // Number of the bill.
	IssuedUTC  time.Time `json:"IssuedUtc"`  // Date and time of the bill issuance in UTC timezone in ISO 8601 format.
	DueUTC     time.Time `json:"DueUtc"`     // Bill due date and time in UTC timezone in ISO 8601 format.
	Notes      string    `json:"Notes"`      // Additional notes.
	Revenue    Revenue   `json:"Revenue"`    // The revenue items on the bill.
	Payments   Payments  `json:"Payments"`   // The payments on the bill.
	State      BillState `json:"State"`      // State of the bill.
}

type BillType string

const (
	BillTypeReceipt BillType = "Receipt"
	BillTypeInvoice BillType = "Invoice"
)

type BillState string

const (
	BillStateOpen   BillState = "open"
	BillStateClosed BillState = "closed"
)

type Revenue []AccountingItem

type AccountingItem struct {
	ID                   string             `json:"Id"`                   // Unique identifier of the item.
	CustomerID           string             `json:"CustomerId"`           // Unique identifier of the Customer whose account the item belongs to.
	ProductID            string             `json:"ProductId"`            // Unique identifier of the Product.
	ServiceID            string             `json:"ServiceId"`            // Unique identifier of the Service the item belongs to.
	OrderID              string             `json:"OrderId"`              // Unique identifier of the order (or Reservation) the item belongs to.
	BillID               string             `json:"BillId"`               // Unique identifier of the bill the item is assigned to.
	AccountingCategoryID string             `json:"AccountingCategoryId"` // Unique identifier of the Accounting Category the item belongs to.
	Amount               Amount             `json:"Amount"`               // Amount the item costs, negative amount represents either rebate or a payment.
	Type                 AccountingItemType `json:"Type"`                 // Type of the item.
	Name                 string             `json:"Name"`                 // Name of the item.
	InvoiceID            string             `json:"InvoiceId"`            // Unique identifier of the invoiced Bill the item is receivable for.
	Notes                string             `json:"Notes"`                // Additional notes.
	ConsumptionUTC       time.Time          `json:"ConsumptionUtc"`       // Date and time of the item consumption in UTC timezone in ISO 8601 format.
}

type AccountingItemType string

const (
	AccountingItemTypeServiceRevenue    AccountingItemType = "ServiceRevenue"
	AccountingItemTypeProductRevenue    AccountingItemType = "ProductRevenue"
	AccountingItemTypeAdditionalRevenue AccountingItemType = "AdditionalRevenue"
	AccountingItemTypePayment           AccountingItemType = "Payment"
)

type Payments []AccountingItem

type Amount struct {
	Currency string  `json:"Currency"` // ISO-4217 currency code, e.g. EUR or USD.
	Net      float64 `json:"Net"`      // Net value in case the item is taxed.
	Tax      float64 `json:"Tax"`      // Tax value in case the item is taxed.
	TaxRate  float64 `json:"TaxRate"`  // Tax rate in case the item is taxed (e.g. 0.21).
	Value    float64 `json:"Value"`    // Amount in the currency (including tax if taxed).
}
