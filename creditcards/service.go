package creditcards

import "github.com/mobmax/go-mews/json"

type Service struct {
	Client *json.Client
}

func NewService() *Service {
	return &Service{}
}
