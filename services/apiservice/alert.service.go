package apiservice

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/codebulletin/AQMFluxAPI/db"
	"github.com/codebulletin/AQMFluxAPI/logger"
	"github.com/codebulletin/AQMFluxAPI/repo"
	"github.com/codebulletin/AQMFluxAPI/types"
	"github.com/codebulletin/AQMFluxAPI/utils"
)

type AlertService struct {
	db db.DB
	logger logger.Logger
}

func NewAlertService(db db.DB) *AlertService {
	logger := logger.GetLogger()
	return &AlertService{db: db, logger: logger}
}

func (a *AlertService) CreateAlert(w http.ResponseWriter, r *http.Request) {
	query := repo.New(a.db)
	defer query.Close()

	var data types.Alert

	err := utils.ParseRequest(r, &data)

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("error parsing alert %v", err))
		return
	}

	alert, err := query.CreateThreshold(r.Context(), repo.CreateThresholdParams{
		SensorID: data.SensorId,
		DeviceID: data.DeviceId,
		AttributeID: data.AttributeId,
		MessageID: data.MessageId,
		OperatorID: data.OperatorId,
		Triggerenabled: data.Enabled,
		Triggername: data.Name,
		Frequency: data.Frequency,
		Value1: data.Value1,
		Value2: sql.NullFloat64{
			Valid: data.Value2.Valid,
			Float64: data.Value2.Float64,
		},
	})

	if err != nil {
		a.logger.Error(err.Error())
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("error creating alert"))
		return
	}

	err = utils.WriteJSON(w, http.StatusCreated, types.OkCreatedJsonMessage{
		Message: "Alert created successfully",
		Data:    alert,
	})

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("error writing response"))
		return
	}
}

func (a *AlertService) GetAlerts(w http.ResponseWriter, r *http.Request) {
	query := repo.New(a.db)
	defer query.Close()

	data, err := query.GetThresholds(r.Context())

	if err != nil {
		a.logger.Error(err.Error())
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("error getting alerts"))
		return
	}

	alerts := make([]types.Threshold, len(data))

	for i, alert := range data {
		alerts[i] = types.Threshold{
			Id: alert.ID,
			SensorId: alert.SensorID,
			DeviceId: alert.DeviceID,
			AttributeId: alert.AttributeID,
			MessageId: alert.MessageID,
			OperatorId: alert.OperatorID,
			Enabled: alert.Triggerenabled,
			Name: alert.Triggername,
			Frequency: alert.Frequency,
			Value1: alert.Value1,
			Value2: sql.NullFloat64{
				Valid: alert.Value2.Valid,
				Float64: alert.Value2.Float64,
			},
		}
	}

	err = utils.WriteJSON(w, http.StatusOK, alerts)

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("error writing response"))
		return
	}
}

func (a *AlertService) UpdateAlert(w http.ResponseWriter, r *http.Request) {
	query := repo.New(a.db)
	defer query.Close()

	var data types.Threshold

	err := utils.ParseRequest(r, &data)

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("error parsing alert %v", err))
		return
	}

	alert, err := query.UpdateThreshold(r.Context(), repo.UpdateThresholdParams{
		ID: data.Id,
		SensorID: data.SensorId,
		DeviceID: data.DeviceId,
		AttributeID: data.AttributeId,
		MessageID: data.MessageId,
		OperatorID: data.OperatorId,
		Triggerenabled: data.Enabled,
		Triggername: data.Name,
		Frequency: data.Frequency,
		Value1: data.Value1,
		Value2: sql.NullFloat64{
			Valid: data.Value2.Valid,
			Float64: data.Value2.Float64,
		},
	})

	if err != nil {
		a.logger.Error(err.Error())
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("error updating alert"))
		return
	}

	err = utils.WriteJSON(w, http.StatusOK, types.OkCreatedJsonMessage{
		Message: "Alert updated successfully",
		Data:    alert,
	})

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("error writing response"))
		return
	}
}

func (a *AlertService) DeleteAlert(w http.ResponseWriter, r *http.Request) {
	query := repo.New(a.db)
	defer query.Close()

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	data, err := query.DeleteThreshold(r.Context(), int32(id))

	if err != nil {
		a.logger.Error(err.Error())
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("error deleting alert"))
		return
	}

	err = utils.WriteJSON(w, http.StatusOK, types.OkCreatedJsonMessage{
		Message: "Alert deleted successfully",
		Data:   data,
	})

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("error writing response"))
		return
	}
}