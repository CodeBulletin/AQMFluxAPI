package api

import (
	"net/http"

	"github.com/codebulletin/AQMFluxAPI/db"
	"github.com/codebulletin/AQMFluxAPI/services/apiservice"
)

type LocationAPI struct {
	l *apiservice.LocationService
}

func NewLocationAPI(db db.DB) *LocationAPI {
	return &LocationAPI{
		l: apiservice.NewLocationService(db),
	}
}

func (l *LocationAPI) Router() http.Handler {
	router := http.NewServeMux()

	router.HandleFunc("POST /", l.l.AddLocation)
	router.HandleFunc("PUT /", l.l.UpdateLocation)
	router.HandleFunc("GET /", l.l.GetLocation)

	return router
}