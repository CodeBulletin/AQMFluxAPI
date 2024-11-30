package apiservice

import (
	"fmt"
	"net/http"

	"github.com/codebulletin/AQMFluxAPI/db"
	"github.com/codebulletin/AQMFluxAPI/logger"
	"github.com/codebulletin/AQMFluxAPI/repo"
	"github.com/codebulletin/AQMFluxAPI/utils"
)

type OperatorService struct {
	db db.DB
	logger logger.Logger
}

func NewOperatorService(db db.DB) *OperatorService {
	logger := logger.GetLogger()
	return &OperatorService{
		db: db,
		logger: logger,
	}
}

func (a *OperatorService) GetOperators(w http.ResponseWriter, r *http.Request) {
	query := repo.New(a.db)
	defer query.Close()

	operators, err := query.GetOperators(r.Context())

	if err != nil {
		a.logger.Error(err.Error())
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("error getting operators"))
		return
	}

	err = utils.WriteJSON(w, http.StatusOK, operators)

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("error writing response"))
		return
	}
}