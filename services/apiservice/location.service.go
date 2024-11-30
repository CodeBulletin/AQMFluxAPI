package apiservice

import (
	"fmt"
	"net/http"

	"github.com/codebulletin/AQMFluxAPI/db"
	"github.com/codebulletin/AQMFluxAPI/logger"
	"github.com/codebulletin/AQMFluxAPI/repo"
	"github.com/codebulletin/AQMFluxAPI/types"
	"github.com/codebulletin/AQMFluxAPI/utils"
)

type LocationService struct {
	db db.DB
	logger logger.Logger
}

func NewLocationService(db db.DB) *LocationService {
	logger := logger.GetLogger()
	return &LocationService{
		db: db,
		logger: logger,
	}
}

func (a *LocationService) AddLocation(w http.ResponseWriter, r *http.Request) {
	query := repo.New(a.db)
	defer query.Close()

	var loct types.Location

	err := utils.ParseRequest(r, &loct)

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("error parsing location"))
		return
	}

	err = query.CreateLocation(r.Context(), repo.CreateLocationParams{
		LocationName: loct.Name,
		LocationID:   loct.Id,
		LocationDesc: loct.Desc,
	})

	if err != nil {
		a.logger.Error(err.Error())
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("error adding location"))
		return
	}

	err = utils.WriteJSON(w, http.StatusCreated, types.OkCreatedJsonMessage{
		Message: "location added successfully",
		Data:    loct,
	})

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("error writing response"))
		return
	}
}

func (a *LocationService) UpdateLocation(w http.ResponseWriter, r *http.Request) {
	query := repo.New(a.db)
	defer query.Close()

	var loct types.Location

	err := utils.ParseRequest(r, &loct)

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("error parsing location"))
		return
	}

	err = query.UpdateLocation(r.Context(), repo.UpdateLocationParams{
		LocationName: loct.Name,
		LocationID:   loct.Id,
		LocationDesc: loct.Desc,
	})

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("error updating location"))
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

func (a *LocationService) GetLocation(w http.ResponseWriter, r *http.Request) {
	query := repo.New(a.db)
	defer query.Close()

	locs, err := query.GetLocations(r.Context())

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("error getting locations"))
		return
	}

	locations := make([]types.Location, len(locs))

	for i, loc := range locs {
		locations[i] = types.Location{
			Id:   loc.LocationID,
			Name: loc.LocationName,
			Desc: loc.LocationDesc,
		}
	}

	err = utils.WriteJSON(w, http.StatusOK, locations)

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("error writing response"))
		return
	}
}