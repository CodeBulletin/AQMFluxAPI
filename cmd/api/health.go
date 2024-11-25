package api

import (
	"net/http"

	"github.com/codebulletin/AQMFluxAPI/services/apiservice"
)

type HealthAPI struct {
	healthService *apiservice.HealthService
}

func NewHealthAPI() *HealthAPI {
	return &HealthAPI{
		healthService: apiservice.NewHealthService(),
	}
}

func (h *HealthAPI) Router() http.Handler {
	router := http.NewServeMux()

	router.HandleFunc("/", h.healthService.Check)

	return router
}
