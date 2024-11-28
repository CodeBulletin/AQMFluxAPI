package server

import (
	"net/http"

	"github.com/codebulletin/AQMFluxAPI/db"
	"github.com/codebulletin/AQMFluxAPI/logger"
	"github.com/codebulletin/AQMFluxAPI/middleware"
)

type Server struct {
	addr 	string
	db   	db.DB
	logger  logger.Logger
	router  http.Handler
	*http.Server
}


func NewServer(addr string, db db.DB, logger logger.Logger, router http.Handler) *Server {
	return &Server{addr: addr, db: db, logger: logger, router: router}
}

func (s *Server) Close() {
	s.logger.Status("Server closed on %s", s.addr)

	s.Server.Close()
}

func (s *Server) Start() {
	s.logger.Status("Server started on %s", s.addr)

	httpserver := http.Server{
		Addr:    s.addr,
		Handler: middleware.Logger(s.logger)(s.router),
	}

	httpserver.ListenAndServe()

	s.Server = &httpserver
}
