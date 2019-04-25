package services

import "github.com/mobmax/go-mews/json"

type APIService struct {
	Client *json.Client
}

func NewAPIService() *APIService {
	return &APIService{}
}
