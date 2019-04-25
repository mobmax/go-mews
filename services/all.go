package services

import (
	"github.com/mobmax/go-mews/json"
)

const (
	endpointAll = "services/getAll"
)

// List all products
func (s *APIService) All(requestBody *AllRequest) (*AllResponse, error) {
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
	Services          Services                        `json:"Services"`          // Services that have been reserved.
}

type Services []Service

type Service struct {
	ID         string      `json:"Id"`         // Unique identifier of the service.
	IsActive   bool        `json:"IsActive"`   // Whether the service is still active.
	Name       string      `json:"Name"`       // Whether the service is still active.
	StartTime  string      `json:"StartTime"`  // Default start time of the service orders in ISO 8601 duration format.
	EndTime    string      `json:"EndTime"`    // Default end time of the service orders in ISO 8601 duration format.
	Promotions Promotions  `json:"Promotions"` // Promotions of the service.
	Type       ServiceType `json:"Type"`       // Type of the service.
}

func (s *APIService) NewAllRequest() *AllRequest {
	return &AllRequest{}
}

type AllRequest struct {
	json.BaseRequest
}

type Promotions struct {
	BeforeCheckIn  bool `json:"BeforeCheckIn"`  // Whether it can be promoted before check-in.
	AfterCheckIn   bool `json:"AfterCheckIn"`   // Whether it can be promoted after check-in.
	DuringStay     bool `json:"DuringStay"`     // Whether it can be promoted during stay.
	BeforeCheckOut bool `json:"BeforeCheckout"` // Whether it can be promoted before check-out.
	AfterCheckOut  bool `json:"AfterCheckout"`  // Whether it can be promoted after check-out.
}

type ServiceType string
