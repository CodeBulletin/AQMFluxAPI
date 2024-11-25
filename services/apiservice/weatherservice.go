package apiservice

import (
	"fmt"
	"net/http"
	"time"

	"github.com/codebulletin/AQMFluxAPI/db"
	"github.com/codebulletin/AQMFluxAPI/db/repo"
	"github.com/codebulletin/AQMFluxAPI/logger"
	"github.com/codebulletin/AQMFluxAPI/types"
	"github.com/codebulletin/AQMFluxAPI/utils"
)

type WeatherService struct {
	db db.DB
}

func NewWeatherService(db db.DB) *WeatherService {
	return &WeatherService{
		db: db,
	}
}

func getOpenWeatherLocation(apiKey string, queryParams *map[string]string, limit string, data *[]types.OpenWeatherLocation) (int, error) {
	// Create a network request to the OpenWeatherMap API
	// Use the apiKey to authenticate the request
	// start := time.Now()
	// resp, err := http.Get(fmt.Sprintf("http://api.openweathermap.org/geo/1.0/direct?q=%s&limit=%s&appid=%s", location, limit, apiKey))

	logger := logger.GetLogger()

	resp, err := utils.HTTPGet("https://api.openweathermap.org/geo/1.0/direct", map[string]string{
		"q":     fmt.Sprintf("%s,%s,%s", (*queryParams)["city"], (*queryParams)["state"], (*queryParams)["country"]),
		"limit": limit,
		"appid": apiKey,
	}, 5*time.Second, logger)

	if resp.StatusCode == 401 {
		return resp.StatusCode, fmt.Errorf("error getting data from OpenWeatherMap API: Unauthorized")
	}

	if err != nil {
		return resp.StatusCode, fmt.Errorf("error getting data from OpenWeatherMap API: %v", err)
	}

	defer resp.Body.Close()

	err = utils.ParseResponse(resp, data)

	if err != nil {
		return resp.StatusCode, fmt.Errorf("error parsing response body: %v", err)
	}

	return resp.StatusCode, nil
}

func (h *WeatherService) GetOpenWeatherLocation(w http.ResponseWriter, r *http.Request) {
	queryParams, err := utils.ParseQueryParams(r, types.Param{
		Name:          "city",
		Optional:      false,
		DefaultValue:  "",
	}, types.Param{
		Name:          "state",
		Optional:      true,
		DefaultValue:  "",
	}, types.Param{
		Name:          "country",
		Optional:      true,
		DefaultValue:  "",
	}, types.Param{
		Name:          "limit",
		Optional:      true,
		DefaultValue:  "5",
	})

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("error parsing query parameters: %v", err))
		return
	}

	query := repo.New(h.db)
	defer query.Close()

	keyval, err := query.GetConfigByKey(r.Context(), "OPEN WEATHER MAP API KEY")

	var api_key string

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("error loading config data"))
		return
	}

	if keyval.Cvalue.Valid {
		api_key = keyval.Cvalue.String
	} else {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("error loading config data - OpenWeatherMap api key not found"))
		return
	}

	var data []types.OpenWeatherLocation

	_, err = getOpenWeatherLocation(api_key, &queryParams, queryParams["limit"], &data)

	if err != nil {
		if err.Error() == "error getting data from OpenWeatherMap API: Unauthorized" {
			utils.WriteError(w, http.StatusForbidden, fmt.Errorf("unauthorized access to OpenWeatherMap api - check your api key"))
			return
		}
		utils.WriteError(w, http.StatusForbidden, fmt.Errorf("error getting OpenWeatherMap location: %v", err))
		return
	}

	err = utils.WriteJSON(w, http.StatusOK, data)

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("error writing response: %v", err))
		return
	}
}
