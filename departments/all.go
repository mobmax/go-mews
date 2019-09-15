package departments

import "github.com/mobmax/go-mews/json"

const (
	endpointAll = "departments/getAll"
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
	Departments []Department `json:"Departments"`
}

func (s *Service) NewAllRequest() *AllRequest {
	return &AllRequest{}
}

type AllRequest struct {
	json.BaseRequest
}

type Department struct {
	Id       string `json:"Id"`       // Unique identifier of the department.
	IsActive bool   `json:"IsActive"` // Whether the accounting category is still active.
	Name     string `json:"Name"`     // Name of the category.
}
