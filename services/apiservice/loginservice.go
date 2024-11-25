package apiservice

import (
	"net/http"
	"time"

	"github.com/codebulletin/AQMFluxAPI/config"
	"github.com/codebulletin/AQMFluxAPI/db"
	"github.com/codebulletin/AQMFluxAPI/db/repo"
	"github.com/codebulletin/AQMFluxAPI/types"
	"github.com/codebulletin/AQMFluxAPI/utils"
)

type LoginService struct {
	db db.DB
}

func NewLoginService(db db.DB) *LoginService {
	return &LoginService{db: db}
}

func (l *LoginService) Login(w http.ResponseWriter, r *http.Request) {
	config.GetAUTHConfig().Load()
	query := repo.New(l.db)
	defer query.Close()

	// Get the username and password from the request body
	var data types.LoginRequest

	err := utils.ParseRequest(r, &data)

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// Get the user from the database
	user, err := query.GetConfigByKey(r.Context(), "USERNAME")

	if err != nil {
		utils.WriteError(w, http.StatusUnauthorized, utils.ErrUnauthorized)
		return
	}

	// Check if the user exists
	if !user.Cvalue.Valid || user.Cvalue.String != data.Username {
		utils.WriteError(w, http.StatusUnauthorized, utils.ErrUnauthorized)
		return
	}

	// Check if the password is correct
	password, err := query.GetConfigByKey(r.Context(), "PASSWORD")

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	if !password.Cvalue.Valid || password.Cvalue.String != data.Password {
		utils.WriteError(w, http.StatusUnauthorized, utils.ErrUnauthorized)
		return
	}

	jwt_secret, err := query.GetSecretByName(r.Context(), "jwt_secret")
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	refresh_secret, err := query.GetSecretByName(r.Context(), "jwt_refresh_secret")
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	// Generate a token, refresh token and return them
	jwt_token, refresh_token, err := utils.GenerateTokens(data.Username, jwt_secret.Value, refresh_secret.Value, time.Duration(config.GetAUTHConfig().TokenDuration()) * time.Minute, time.Duration(config.GetAUTHConfig().RefreshTokenDuration()) * time.Minute)

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	// w.Header().Set("Authorization", jwt_token)
	// w.Header().Set("Refresh-Token", refresh_token)

	utils.SetCookie(w, "Authorization", jwt_token, 24 * time.Hour, true)
	utils.SetCookie(w, "Refresh-Token", refresh_token, 24 * time.Hour, true)

	err = utils.WriteJSON(w, http.StatusOK, types.OkJsonMessage{Message: "Login successful"})

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
}