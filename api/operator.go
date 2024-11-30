package api

import (
	"net/http"

	"github.com/codebulletin/AQMFluxAPI/db"
	"github.com/codebulletin/AQMFluxAPI/services/apiservice"
)

type OperatorAPI struct {
	a *apiservice.OperatorService
}

func NewOperatorAPI(db db.DB) *OperatorAPI {
	return &OperatorAPI{
		a: apiservice.NewOperatorService(db),
	}
}

func (a *OperatorAPI) Router() http.Handler {
	router := http.NewServeMux()

	router.HandleFunc("GET /", a.a.GetOperators)

	return router
}