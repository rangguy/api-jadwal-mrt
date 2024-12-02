package station

import (
	"net/http"
	"time"
)

type Service interface {
	GetAllStations() (response []StationResponse, err error)
}

type service struct {
	client *http.Client
}

func NewService() Service {
	return &service{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (s *service) GetAllStations() (response []StationResponse, err error) {
	url := "https://jakartamrt.co.id/id/val/stasiuns"

	// hit url

	// return response

	return
}
