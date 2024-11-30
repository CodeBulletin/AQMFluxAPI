package api

import (
	"net/http"

	"github.com/codebulletin/AQMFluxAPI/db"
	"github.com/codebulletin/AQMFluxAPI/services/apiservice"
)

type WeatherAPI struct {
	w *apiservice.WeatherService
}

func NewWeatherAPI(db db.DB) *WeatherAPI {
	return &WeatherAPI{
		w: apiservice.NewWeatherService(db),
	}
}

func (c *WeatherAPI) Router() http.Handler {
	router := http.NewServeMux()

	router.HandleFunc("POST /Location", c.w.GetOpenWeatherLocation)

	adminRouter := http.NewServeMux()
	adminRouter.Handle("/", router)

	return adminRouter
}
