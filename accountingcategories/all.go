package accountingcategories

import "github.com/mobmax/go-mews/json"

const (
	endpointAll = "accountingCategories/getAll"
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
	AccountingCategories []AccountingCategory
}

func (s *Service) NewAllRequest() *AllRequest {
	return &AllRequest{}
}

type AllRequest struct {
	json.BaseRequest
}

type AccountingCategory struct {
	ID                 string `json:"ID"`                 // Unique identifier of the category.
	IsActive           bool   `json:"IsActive"`           // Whether the accounting category is still active.
	Name               string `json:"Name"`               // Name of the category.
	Code               string `json:"Code"`               // Code of the category within Mews.
	Classification     string `json:"Classification"`     // Classification of the accounting category allowing cross-enterprise reporting.
	ExternalCode       string `json:"ExternalCode"`       // Code of the category in external systems.
	LedgerAccountCode  string `json:"LedgerAccountCode"`  // Code of the ledger account (double entry accounting).
	PostingAccountCode string `json:"PostingAccountCode"` // Code of the posting account (double entry accounting).
	CostCenterCode     string `json:"CostCenterCode"`     // Code of cost center.
}
