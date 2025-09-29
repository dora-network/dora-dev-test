package service

import (
	"net/http"
)

type Service struct {
}

func NewService() Service {
	return Service{}
}

func (s Service) HealthCheck(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func (s Service) GetTicks(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func (s Service) GetCandles(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}
