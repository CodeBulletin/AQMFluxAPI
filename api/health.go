package api

import (
	"net/http"

	"github.com/codebulletin/AQMFluxAPI/services/apiservice"
)

type HealthAPI struct {
	h *apiservice.HealthService
}

func NewHealthAPI() *HealthAPI {
	return &HealthAPI{
		h: apiservice.NewHealthService(),
	}
}

func (h *HealthAPI) Router() http.Handler {
	router := http.NewServeMux()

	router.HandleFunc("/", h.h.Check)

	return router
}
