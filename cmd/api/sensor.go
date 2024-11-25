package api

import (
	"net/http"

	"github.com/codebulletin/AQMFluxAPI/db"
	"github.com/codebulletin/AQMFluxAPI/services/apiservice"
)

type SensorsAPI struct {
	s *apiservice.SensorService
}

func NewSensorsAPI(db db.DB) *SensorsAPI {
	return &SensorsAPI{
		s: apiservice.NewSensorService(db),
	}
}

func (s *SensorsAPI) Router() http.Handler {
	router := http.NewServeMux()

	router.HandleFunc("POST /", s.s.AddSensor)
	router.HandleFunc("PUT /", s.s.UpdateSensor)
	router.HandleFunc("GET /", s.s.GetSensor)

	return router
}