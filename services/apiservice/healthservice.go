package apiservice

import (
	"net/http"
)

type HealthService struct {
}

func NewHealthService() *HealthService {
	return &HealthService{}
}

func (h *HealthService) Check(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
