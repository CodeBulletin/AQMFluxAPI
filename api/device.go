package api

import (
	"net/http"

	"github.com/codebulletin/AQMFluxAPI/db"
	"github.com/codebulletin/AQMFluxAPI/services/apiservice"
)

type DeviceAPI struct {
	a *apiservice.DeviceService
}

func NewDeviceAPI(db db.DB) *DeviceAPI {
	return &DeviceAPI{
		a: apiservice.NewDeviceService(db),
	}
}

func (d *DeviceAPI) Router() http.Handler {
	router := http.NewServeMux()

	router.HandleFunc("POST /", d.a.CreateDevice)
	router.HandleFunc("PUT /", d.a.UpdateDevice)
	router.HandleFunc("GET /", d.a.GetDevice)
	router.HandleFunc("GET /all/", d.a.GetDeviceList)

	return router
}
