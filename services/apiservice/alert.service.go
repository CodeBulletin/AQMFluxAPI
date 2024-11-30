package apiservice

import (
	"database/sql"
	"fmt"
	"net/http"

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