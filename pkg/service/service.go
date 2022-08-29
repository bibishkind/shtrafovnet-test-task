package service

import (
	"net/http"

	"github.com/bibishkind/shtrafovnet-test-task/models"
)

type Service interface {
	GetCompanyByINN(inn string) (*models.Company, error)
}

func NewService() Service {
	return &service{httpClient: &http.Client{Transport: &http.Transport{
		MaxIdleConnsPerHost: 10,
	}}}
}

type service struct {
	httpClient *http.Client
}
