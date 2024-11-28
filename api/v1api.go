package api

import (
	"net/http"

	"github.com/codebulletin/AQMFluxAPI/db"
	"github.com/codebulletin/AQMFluxAPI/middleware"
)

type V1API struct {
	db 		db.DB
}

func NewV1API(db db.DB) *V1API {
	return &V1API{db: db}
}

func (v *V1API) Router() http.Handler {
	router := http.NewServeMux()

	healthRouter := NewHealthAPI().Router()
	router.Handle("/health/", http.StripPrefix("/health", healthRouter))

	loginRouter := NewLoginAPI(v.db).Router()
	router.Handle("/login/", http.StripPrefix("/login", loginRouter))

	configRouter := NewConfigAPI(v.db).Router()
	router.Handle("/config/", middleware.Authentication(v.db)(http.StripPrefix("/config", configRouter)))

	weatherRouter := NewWeatherAPI(v.db).Router()
	router.Handle("/weather/", middleware.Authentication(v.db)(http.StripPrefix("/weather", weatherRouter)))

	attributeRouter := NewAttributeAPI(v.db).Router()
	router.Handle("/attribute/", middleware.Authentication(v.db)(http.StripPrefix("/attribute", attributeRouter)))

	locationRouter := NewLocationAPI(v.db).Router()
	router.Handle("/location/", middleware.Authentication(v.db)(http.StripPrefix("/location", locationRouter)))
	
	sensorRouter := NewSensorsAPI(v.db).Router()
	router.Handle("/sensor/", middleware.Authentication(v.db)(http.StripPrefix("/sensor", sensorRouter)))

	deviceRouter := NewDeviceAPI(v.db).Router()
	router.Handle("/device/", middleware.Authentication(v.db)(http.StripPrefix("/device", deviceRouter)))


	v1 := http.NewServeMux()
	v1.Handle("/v1/", http.StripPrefix("/v1", middleware.Chain(middleware.AllowCors(), middleware.Preflight)(router)))

	return v1
}