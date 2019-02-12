package service

import (
	m "go-template/model"
)

type Service struct {
}

type GotemplateService interface {
        Health() (*m.Health, error)
}

func NewService() *Service {
	return &Service{
	}
}

func (s *Service) Health() (*m.Health, error) {
	return s.Health()
}
