package api

import (
	"net/http"

	"github.com/codebulletin/AQMFluxAPI/db"
	"github.com/codebulletin/AQMFluxAPI/services/apiservice"
)

type AlertAPI struct {
	a *apiservice.AlertService
}

func NewAlertAPI(db db.DB) *AlertAPI {
	return &AlertAPI{
		a: apiservice.NewAlertService(db),
	}
}

func (a *AlertAPI) Router() http.Handler {
	router := http.NewServeMux()

	router.HandleFunc("POST /", a.a.CreateAlert)

	return router
}