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

type AttributeService struct {
	db db.DB
	logger logger.Logger
}

func NewAttributeService(db db.DB) *AttributeService {
	logger := logger.GetLogger()
	return &AttributeService{
		db: db,
		logger: logger,
	}
}

// AddAttribute adds a new attribute to the database
func (a *AttributeService) AddAttribute(w http.ResponseWriter, r *http.Request) {
	query := repo.New(a.db)
	defer query.Close()

	var attr types.Attribute

	err := utils.ParseRequest(r, &attr)

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("error parsing attribute"))
		return
	}

	err = query.CreateAttribute(r.Context(), repo.CreateAttributeParams{
		AttributeName: attr.Name,
		AttributeID:   attr.Id,
		AttributeDesc: attr.Desc,
	})

	if err != nil {
		a.logger.Error(err.Error())
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("error adding attribute"))
		return
	}

	err = utils.WriteJSON(w, http.StatusCreated, types.OkCreatedJsonMessage{
		Message: "Attribute added successfully",
		Data:    attr,
	})

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("error writing response"))
		return
	}
}

// UpdateAttribute updates an existing attribute in the database
func (a *AttributeService) UpdateAttribute(w http.ResponseWriter, r *http.Request) {
	query := repo.New(a.db)
	defer query.Close()

	var attr types.Attribute

	err := utils.ParseRequest(r, &attr)

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("error parsing attribute"))
		return
	}

	err = query.UpdateAttribute(r.Context(), repo.UpdateAttributeParams{
		AttributeName: attr.Name,
		AttributeID:   attr.Id,
		AttributeDesc: attr.Desc,
	})

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("error updating attribute"))
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

// Get ALL attributes from the database
func (a *AttributeService) GetAttributes(w http.ResponseWriter, r *http.Request) {
	query := repo.New(a.db)
	defer query.Close()

	attrs, err := query.GetAllAttributes(r.Context())

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("error getting attributes"))
		return
	}

	attributes := make([]types.Attribute, len(attrs))

	for i, attr := range attrs {
		attributes[i] = types.Attribute{
			Id:   attr.AttributeID,
			Name: attr.AttributeName,
			Desc: attr.AttributeDesc,
		}
	}

	err = utils.WriteJSON(w, http.StatusOK, attributes)

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("error writing response"))
		return
	}
}