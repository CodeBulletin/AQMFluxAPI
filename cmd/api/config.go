package api

import (
	"net/http"

	"github.com/codebulletin/AQMFluxAPI/db"
	"github.com/codebulletin/AQMFluxAPI/services/apiservice"
)

type ConfigAPI struct {
	a *apiservice.ConfigService
}

func NewConfigAPI(db db.DB) *ConfigAPI {
	return &ConfigAPI{
		a: apiservice.NewConfigService(db),
	}
}

func (c *ConfigAPI) Router() http.Handler {
	router := http.NewServeMux()

	router.HandleFunc("POST /keys", c.a.GetConfigs)
	router.HandleFunc("PUT /", c.a.UpdateConfigByKeyValuePairs)

	adminRouter := http.NewServeMux()
	adminRouter.Handle("/", router)

	return adminRouter
}
