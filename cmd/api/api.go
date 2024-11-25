package api

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
}

func NewServer(addr string, db db.DB, logger logger.Logger) *Server {
	return &Server{addr: addr, db: db, logger: logger}
}

func (s *Server) V1Router() http.Handler {
	router := http.NewServeMux()

	healthRouter := NewHealthAPI().Router()
	router.Handle("/health/", http.StripPrefix("/health", healthRouter))

	loginRouter := NewLoginAPI(s.db).Router()
	router.Handle("/login/", http.StripPrefix("/login", loginRouter))

	configRouter := NewConfigAPI(s.db).Router()
	router.Handle("/config/", middleware.Authentication(s.db)(http.StripPrefix("/config", configRouter)))

	weatherRouter := NewWeatherAPI(s.db).Router()
	router.Handle("/weather/", middleware.Authentication(s.db)(http.StripPrefix("/weather", weatherRouter)))

	attributeRouter := NewAttributeAPI(s.db).Router()
	router.Handle("/attribute/", middleware.Authentication(s.db)(http.StripPrefix("/attribute", attributeRouter)))

	locationRouter := NewLocationAPI(s.db).Router()
	router.Handle("/location/", middleware.Authentication(s.db)(http.StripPrefix("/location", locationRouter)))
	
	sensorRouter := NewSensorsAPI(s.db).Router()
	router.Handle("/sensor/", middleware.Authentication(s.db)(http.StripPrefix("/sensor", sensorRouter)))

	deviceRouter := NewDeviceAPI(s.db).Router()
	router.Handle("/device/", middleware.Authentication(s.db)(http.StripPrefix("/device", deviceRouter)))

	return router
}

func (s *Server) Start() {
	v1 := http.NewServeMux()
	v1.Handle("/api/v1/", http.StripPrefix("/api/v1", s.V1Router()))
	
	s.logger.Status("Server started on %s", s.addr)

	httpserver := http.Server{
		Addr:    s.addr,
		Handler: middleware.Chain(middleware.Logger(s.logger), middleware.AllowCors(), middleware.Preflight, middleware.Slow)(v1),
	}

	httpserver.ListenAndServe()
}
