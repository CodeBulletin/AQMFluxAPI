package api

import (
	"net/http"

	"github.com/codebulletin/AQMFluxAPI/db"
	"github.com/codebulletin/AQMFluxAPI/services/apiservice"
)

type LoginAPI struct {
	LoginAPI *apiservice.LoginService
}

func NewLoginAPI(db db.DB) *LoginAPI {
	return &LoginAPI{
		LoginAPI: apiservice.NewLoginService(db),
	}
}

func (l *LoginAPI) Router() http.Handler {
	router := http.NewServeMux()

	router.HandleFunc("POST /", l.LoginAPI.Login)

	return router
}