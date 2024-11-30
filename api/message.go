package api

import (
	"net/http"

	"github.com/codebulletin/AQMFluxAPI/db"
	"github.com/codebulletin/AQMFluxAPI/services/apiservice"
)

type MessageAPI struct {
	m *apiservice.MessageService
}

func NewMessageAPI(db db.DB) *MessageAPI {
	return &MessageAPI{
		m: apiservice.NewMessageService(db),
	}
}

func (m *MessageAPI) Router() http.Handler {
	router := http.NewServeMux()

	router.HandleFunc("POST /", m.m.CreateMessage)
	router.HandleFunc("GET /all/", m.m.GetMessages)
	router.HandleFunc("GET /{id}/", m.m.GetMessage)
	router.HandleFunc("PUT /", m.m.UpdateMessage)
	router.HandleFunc("DELETE /{id}/", m.m.DeleteMessage)

	return router
}