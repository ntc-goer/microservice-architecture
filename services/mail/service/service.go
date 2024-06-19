package service

import "github.com/ntc-goer/microservice-examples/mailservice/config"

type Service struct {
	Config *config.Config
}

func NewService(cfg *config.Config) *Service {
	return &Service{
		Config: cfg,
	}
}

func (s *Service) Handle(data string) {

}
