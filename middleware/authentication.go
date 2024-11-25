package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/codebulletin/AQMFluxAPI/config"
	"github.com/codebulletin/AQMFluxAPI/db"
	"github.com/codebulletin/AQMFluxAPI/db/repo"
	"github.com/codebulletin/AQMFluxAPI/utils"
)

func Authentication(db db.DB) func(http.Handler) http.Handler {
	config.GetAUTHConfig().Load()
	return func (next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			cookie, err := r.Cookie("Authorization")

			if err != nil {
				utils.WriteError(w, http.StatusUnauthorized, utils.ErrNoToken)
				return
			}

			token := cookie.Value

			cookie, err = r.Cookie("Refresh-Token")

			if err != nil {
				utils.WriteError(w, http.StatusUnauthorized, utils.ErrNoToken)
				return
			}

			refresh_token := cookie.Value

			if token == "" {
				utils.WriteError(w, http.StatusUnauthorized, utils.ErrNoToken)
				return
			}

			query := repo.New(db)
			jwt_secret, err := query.GetSecretByName(r.Context(), "jwt_secret")
			if err != nil {
				utils.WriteError(w, http.StatusInternalServerError, err)
				return
			}

			// fmt.Println(jwt_secret)

			username, err := query.GetConfigByKey(r.Context(), "USERNAME")

			if err != nil {
				utils.WriteError(w, http.StatusInternalServerError, err)
				return
			}

			if !username.Cvalue.Valid {
				utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("username not set"))
				return
			}

			valid, err := utils.ValidateToken(token, jwt_secret.Value, username.Cvalue.String)

			if err != nil {
				// fmt.Printf("Error validating token, %s", err.Error())
				utils.WriteError(w, http.StatusUnauthorized, utils.ErrUnauthorized)
				return
			}

			if !valid {
				// fmt.Println("Valid token is false")
				utils.WriteError(w, http.StatusUnauthorized, utils.ErrUnauthorized)
				return
			}

			_, err = query.GetSecretByName(r.Context(), "jwt_refresh_secret")

			if err != nil {
				utils.WriteError(w, http.StatusInternalServerError, err)
				return
			}

			// fmt.Println(refresh_token)

			// Generate a new token
			if refresh_token != "" {
				valid, err := utils.ValidateToken(refresh_token, jwt_secret.Value, username.Cvalue.String)
				if err != nil {
					// fmt.Println(err)
					utils.WriteError(w, http.StatusUnauthorized, utils.ErrUnauthorized)
					return
				}

				if !valid {
					utils.WriteError(w, http.StatusUnauthorized, utils.ErrUnauthorized)
					return
				}

				jwt_token, err := utils.RefreshToken(token, jwt_secret.Value, time.Duration(config.GetAUTHConfig().TokenDuration()) * time.Minute)

				if err != nil {
					// fmt.Println(err)
					utils.WriteError(w, http.StatusInternalServerError, err)
					return
				}

				utils.SetCookie(w, "Authorization", jwt_token, 24 * time.Hour, true)
			}


			next.ServeHTTP(w, r)
		})
	}
}