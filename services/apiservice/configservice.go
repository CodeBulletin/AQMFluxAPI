package apiservice

import (
	"fmt"
	"net/http"

	"github.com/codebulletin/AQMFluxAPI/db"
	"github.com/codebulletin/AQMFluxAPI/db/repo"
	"github.com/codebulletin/AQMFluxAPI/logger"
	"github.com/codebulletin/AQMFluxAPI/types"
	"github.com/codebulletin/AQMFluxAPI/utils"
)

type ConfigService struct {
	db db.DB
	logger logger.Logger
}

func NewConfigService(db db.DB) *ConfigService {
	logger := logger.GetLogger()
	return &ConfigService{
		db: db,
		logger: logger,
	}
}

func (h *ConfigService) GetConfigs(w http.ResponseWriter, r *http.Request) {
	query := repo.New(h.db)
	defer query.Close()

	queryParams, err := utils.ParseQueryParams(r, types.Param{
		Name:		     "keys",
		Optional:        false,
		DefaultValue:    "",
	})

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("error parsing query parameters"))
		return
	}

	strKeys := utils.SplitString(queryParams["keys"], ',');

	configs, err := query.LoadConfigData(r.Context(), strKeys)

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("error loading config data"))
		return
	}

	err = utils.WriteJSON(w, http.StatusOK, configs)

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("error writing response"))
		return
	}
}

func (h *ConfigService) UpdateConfigByKeyValuePairs(w http.ResponseWriter, r *http.Request) {
	query := repo.New(h.db)
	defer query.Close()

	// Get All the key value pairs of the form [{key: value}, {key: value}] from the request body
	var data []types.KeyValuePair
	err := utils.ParseRequest(r, &data)

	// Convert the key value pairs to a [(key1, key2), (value1, value2)] format
	var keys   []string
	var values []string
	for _, kv := range data {
		keys = append(keys, kv.Key)
		values = append(values, kv.Value)
	}

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("error parsing request body"))
		return
	}

	err = query.SetConfigBykey(r.Context(), repo.SetConfigBykeyParams{
		Column1: keys,
		Column2: values,
	})
	
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("error updating config data"))
		return
	}

	err = utils.WriteJSON(w, http.StatusOK, types.OkJsonMessage{Message: "Config data updated successfully"})

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("error writing response"))
		return
	}
}
