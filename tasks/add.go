package tasks

import (
	"time"

	"github.com/mobmax/go-mews/json"
)

const (
	endpointAdd = "tasks/add"
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
	DepartmentID string    `json:"DepartmentId,omitempty"`
	Name         string    `json:"Name"`
	Description  string    `json:"Description,omitempty"`
	DeadlineUTC  time.Time `json:"DeadlineUtc"`
}

type AddResponse struct {
}
