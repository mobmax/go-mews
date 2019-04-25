package orders

import (
	"github.com/mobmax/go-mews/json"
)

const (
	endpointAdd = "orders/add"
)

// List all products
func (s *Service) Add(requestBody *AddRequest) (*AddResponse, error) {
	// @TODO: create wrapper?
	if err := s.Client.CheckTokens(); err != nil {
		return nil, err
	}

	apiURL, err := s.Client.GetApiURL(endpointAdd)
	if err != nil {
		return nil, err
	}

	responseBody := &AddResponse{}
	httpReq, err := s.Client.NewRequest(apiURL, requestBody)
	if err != nil {
		return nil, err
	}

	_, err = s.Client.Do(httpReq, responseBody)
	return responseBody, err
}

func (s *Service) NewAddRequest() *AddRequest {
	return &AddRequest{}
}

type AddRequest struct {
	json.BaseRequest
	CustomerID	 	string    	`json:"CustomerId"`
	ServiceID	 	string    	`json:"ServiceId"`
	ConsumptionUtc 	string  	`json:"ConsumptionUtc,omitempty"`
	Notes        	string    	`json:"Notes,omitempty"`
	ProductOrders 	Products 	`json:"ProductOrders,omitempty"`
	Items  		 	Items 		`json:"Items,omitempty"`
}

type Products []Product

type Product struct {
	ProductId 	string 	`json:"ProductId"`
	Count		int		`json:"Count,omitempty"`
	UnitCost	Cost 	`json:"UnitCost,omitempty"`
}

type Items []Item

type Item struct {
	ItemName 				string	`json:"Name"`
	UnitCount				int		`json:"UnitCount"`
	UnitCost				Cost	`json:"UnitCost"`
	AccountingCategoryId	string	`json:"AccountingCategoryId,omitempty"`
}

type Cost struct {
	Amount		float64	`json:"Amount"`
	Currency	string	`json:"Currency"`
	Tax			float64	`json:"Tax"`
}

type AddResponse struct {
	OrderId		string	`json:"OrderId"`
}
