package configuration

import (
	"time"

	"github.com/mobmax/go-mews/json"
)

const (
	endpointGet = "configuration/get"
)

// Returns configuration of the enterprise and the client.
func (s *Service) Get(requestBody *GetRequest) (*GetResponse, error) {
	// @TODO: create wrapper?
	if err := s.Client.CheckTokens(); err != nil {
		return nil, err
	}

	apiURL, err := s.Client.GetApiURL(endpointGet)
	if err != nil {
		return nil, err
	}

	responseBody := &GetResponse{}
	httpReq, err := s.Client.NewRequest(apiURL, requestBody)
	if err != nil {
		return nil, err
	}

	_, err = s.Client.Do(httpReq, responseBody)
	return responseBody, err
}

func (s *Service) NewGetRequest() *GetRequest {
	return &GetRequest{}
}

type GetRequest struct {
	json.BaseRequest
}

type GetResponse struct {
	Enterprise Enterprise `json:"Enterprise"`
	NowUtc     time.Time  `json:"NowUtc"` // NEW
}

type Enterprise struct {
	ID                      string     `json:"Id"`                      // Unique identifier of the enterprise.
	CreatedUTC              time.Time  `json:"CreatedUtc"`              // Creation date and time of the enterprise in UTC timezone in ISO 8601 format.
	Name                    string     `json:"Name"`                    // Name of the enterprise.
	WebsiteURL              string     `json:"WebsiteUrl"`              // URL of the enterprise website.
	Email                   string     `json:"Email"`                   // Email address of the enterprise.
	Phone                   string     `json:"Phone"`                   // Phone number of the enterprise.
	TimeZoneIdentifier      string     `json:"TimeZoneIdentifier"`      // IANA timezone identifier of the enterprise.
	EditableHistoryInterval string     `json:"EditableHistoryInterval"` // Editable history interval in ISO 8601 duration format.
	Address                 Address    `json:"Address"`                 // Address of the enterprise.
	Currencies              Currencies `json:"Currencies"`              // Currencies accepted by the enterprise.
	ChainID                 string     `json:"ChainId"`                 // NEW
	CoverImageId            string     `json:"CoverImageId"`            // NEW
	DefaultLanguageCode     string     `json:"DefaultLanguageCode"`     // NEW
	LegalEnvironmentCode    string     `json:"LegalEnvironmentCode"`    // NEW
	LogoImageID             string     `json:"LogoImageId"`             // NEW
}

type Currencies []Currency

type Currency struct {
	Currency  string `json:"Currency"`
	IsDefault bool   `json:"IsDefault"`
	IsEnabled bool   `json:"IsEnabled"`
}

type Address struct {
	Line1                  string `json:"Line1"`                            // First line of the address.
	Line2                  string `json:"Line2"`                            // Second line of the address.
	City                   string `json:"City"`                             // The city.
	PostalCode             string `json:"PostalCode"`                       // Postal code.
	CountryCode            string `json:"CountryCode"`                      // ISO 3166-1 alpha-2 country code (two letter country code).
	CountrySubdivisionCode string `json:"CountrySubdivisionCode,omitempty"` // ISO 3166-2 code of the administrative division, e.g. DE-BW.
}

type LocalizedText map[string]string
