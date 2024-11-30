package api

import (
	"net/http"

	"github.com/codebulletin/AQMFluxAPI/db"
	"github.com/codebulletin/AQMFluxAPI/services/apiservice"
)

type LoginAPI struct {
	l *apiservice.LoginService
}

func NewLoginAPI(db db.DB) *LoginAPI {
	return &LoginAPI{
		l: apiservice.NewLoginService(db),
	}
}

func (l *LoginAPI) Router() http.Handler {
	router := http.NewServeMux()

	router.HandleFunc("POST /", l.l.Login)

	return router
}