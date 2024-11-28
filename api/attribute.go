package api

import (
	"net/http"

	"github.com/codebulletin/AQMFluxAPI/db"
	"github.com/codebulletin/AQMFluxAPI/services/apiservice"
)

type AttributeAPI struct {
	a *apiservice.AttributeService
}

func NewAttributeAPI(db db.DB) *AttributeAPI {
	return &AttributeAPI{
		a: apiservice.NewAttributeService(db),
	}
}

func (a *AttributeAPI) Router() http.Handler {
	router := http.NewServeMux()

	router.HandleFunc("POST /", a.a.AddAttribute)
	router.HandleFunc("PUT /", a.a.UpdateAttribute)
	router.HandleFunc("GET /", a.a.GetAttributes)

	return router
}