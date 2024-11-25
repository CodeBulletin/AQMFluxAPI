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

type SensorService struct {
	db db.DB
	logger logger.Logger
}

func NewSensorService(db db.DB) *SensorService {
	logger := logger.GetLogger()
	return &SensorService{
		db: db,
		logger: logger,
	}
}

func (a *SensorService) AddSensor(w http.ResponseWriter, r *http.Request) {
	query := repo.New(a.db)
	defer query.Close()

	var senr types.Sensor

	err := utils.ParseRequest(r, &senr)

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("error parsing sensor"))
		return
	}

	err = query.CreateSensors(r.Context(), repo.CreateSensorsParams{
		SensorName: senr.Name,
		SensorID:   senr.Id,
		SensorDesc: senr.Desc,
	})

	if err != nil {
		a.logger.Error(err.Error())
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("error adding sensor"))
		return
	}

	err = utils.WriteJSON(w, http.StatusCreated, types.OkCreatedJsonMessage{
		Message: "sensor added successfully",
		Data:    senr,
	})

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("error writing response"))
		return
	}
}

func (a *SensorService) UpdateSensor(w http.ResponseWriter, r *http.Request) {
	query := repo.New(a.db)
	defer query.Close()

	var senr types.Sensor

	err := utils.ParseRequest(r, &senr)

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("error parsing sensor"))
		return
	}

	err = query.UpdateSensors(r.Context(), repo.UpdateSensorsParams{
		SensorName: senr.Name,
		SensorID:   senr.Id,
		SensorDesc: senr.Desc,
	})

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("error updating sensor"))
		return
	}

	err = utils.WriteJSON(w, http.StatusOK, types.OkJsonMessage{
		Message: "Attribute updated successfully",
	})

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("error writing response"))
		return
	}
}

func (a *SensorService) GetSensor(w http.ResponseWriter, r *http.Request) {
	query := repo.New(a.db)
	defer query.Close()

	sens, err := query.GetSensors(r.Context())

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("error getting sensors"))
		return
	}

	sensors := make([]types.Sensor, len(sens))

	for i, sen := range sens {
		sensors[i] = types.Sensor{
			Id:   sen.SensorID,
			Name: sen.SensorName,
			Desc: sen.SensorDesc,
		}
	}

	err = utils.WriteJSON(w, http.StatusOK, sensors)

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("error writing response"))
		return
	}
}