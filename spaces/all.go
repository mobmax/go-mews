package spaces

import "github.com/mobmax/go-mews/json"

const (
	endpointAll = "spaces/getAll"
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
	Spaces                  Spaces                  `json:"Spaces"`                  // The spaces of the enterprise
	SpaceCategories         SpaceCategories         `json:"SpaceCategories"`         // Categories of spaces in the enterprise.
	SpaceFeatures           SpaceFeatures           `json:"SpaceFeatures"`           // Features of spaces in the enterprise.
	SpaceFeatureAssignments SpaceFeatureAssignments `json:"SpaceFeatureAssignments"` // Assignments of space features to spaces.
}

type Spaces []Space
type SpaceCategories []SpaceCategory
type SpaceFeatures []SpaceFeatures
type SpaceFeatureAssignments []SpaceFeatureAssignment

func (s *Service) NewAllRequest() *AllRequest {
	return &AllRequest{}
}

type AllRequest struct {
	json.BaseRequest
	Extent AllExtent `json:"Extent,omitempty"`
}

type AllExtent struct {
	Spaces          bool `json:"Spaces"`          // Whether the response should contain spaces.
	SpaceCategories bool `json:"SpaceCategories"` // Whether the response should contain space categories.
	SpaceFeatures   bool `json:"SpaceFeatures"`   // Whether the response should contain space features and their assignments.
	Inactive        bool `json:"Inactive"`        // Whether the response should contain inactive entities.
}

type Space struct {
	ID            string     `json:"Id"`            // Unique identifier of the reservation.
	Type          SpaceType  `json:"Type"`          // 	Type of the space.
	Number        string     `json:"number"`        // Number of the space (e.g. room number).
	ParentSpaceID string     `json:"ParentSpaceId"` // dentifier of the parent Space (e.g. room of a bed).
	CategoryID    string     `json:"CategoryId"`    // Identifier of the Space Category assigned to the space.
	State         SpaceState `json:"State"`         // State of the room.
}

type SpaceCategory struct {
	ID          string   `json:"Id"`                    // Unique identifier of the category.
	IsActive    bool     `json:"IsActive"`              // Whether the space category is still active.
	Name        string   `json:"Name"`                  // Name of the category.
	ShortName   string   `json:"ShortName,omitempty"`   // Short name (e.g. code) of the category.
	Description string   `json:"Description,omitempty"` // Description of the category.
	Ordering    int      `json:"Ordering"`              // Ordering of the category, lower number corresponds to lower category (note that uniqueness nor continuous sequence is guaranteed).
	UnitCount   int      `json:"UnitCount"`             // Count of units that can be accommodated (e.g. bed count).
	ImageIDs    []string `json:"ImageIds"`              // Unique identifiers of the space category images.
}

type SpaceFeature struct {
	ID          string `json:"Id"`          // Unique identifier of the feature.
	Name        string `json:"Name"`        // Name of the feature.
	Description string `json:"Description"` // Description of the feature.
}

type SpaceFeatureAssignment struct {
	SpaceID        string `json:"SpaceId"`        // Unique identifier Space.
	SpaceFeatureID string `json:"SpaceFeatureId"` // Unique identifier Space Feature.
}

type SpaceType string

const (
	SpaceTypeRoom SpaceType = "Room"
	SpaceTypeDorm SpaceType = "Dorm"
	SpaceTypeBed  SpaceType = "Bed"
)

type SpaceState string

const (
	SpaceStateDirty        SpaceState = "Dirty"
	SpaceStateClean        SpaceState = "Clean"
	SpaceStateInspected    SpaceState = "Inspected"
	SpaceStateOutOfService SpaceState = "OutOfService"
	SpaceStateOutOfOrder   SpaceState = "OutOfOrder"
)
