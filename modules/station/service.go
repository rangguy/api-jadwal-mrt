package station

import (
	"encoding/json"
	"github.com/rangguy/api-jadwal-mrt.git/common/client"
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
	byteResponse, err := client.DoRequest(s.client, url)
	if err != nil {
		return
	}

	var stations []Station
	err = json.Unmarshal(byteResponse, &stations)

	// return response
	for _, station := range stations {
		response = append(response, StationResponse{
			Id:   station.Id,
			Name: station.Name,
		})
	}

	return response, err
}
