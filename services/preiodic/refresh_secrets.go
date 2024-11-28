package preiodic

import (
	"context"
	"database/sql"
	"time"

	"github.com/codebulletin/AQMFluxAPI/db"
	"github.com/codebulletin/AQMFluxAPI/logger"
	"github.com/codebulletin/AQMFluxAPI/repo"
	"github.com/codebulletin/AQMFluxAPI/utils"
)

type RefreshSecrets struct {
	db db.DB
	logger logger.Logger
	exitChan chan bool
}

func NewRefreshSecrets(db db.DB) *RefreshSecrets {
	logger := logger.GetLogger()
	return &RefreshSecrets{
		db: db,
		logger: logger,
		exitChan: make(chan bool),
	}
}

func  (r *RefreshSecrets) updateSecrets(ctx context.Context, expireAt time.Time, newSecret string) error {
	query := repo.New(r.db)
	defer query.Close()

	secrets, err := query.GetExpiredSecrets(ctx)
	if err != nil {
		return err
	}

	for _, secret := range secrets {
		err := query.UpdateSecret(ctx, repo.UpdateSecretParams{
			Name: secret.Name,
			ExpiresAt: sql.NullTime{
				Time:  expireAt,
			},
			Value: newSecret,
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *RefreshSecrets) refresh() {
	defer func ()  {
		if recover := recover(); recover != nil {
			r.logger.Fatal("Recovered in refresh: %v", recover)
		}
	}()

	expireAt := time.Now().Add(24 * time.Hour)
	newSecret, err := utils.GenerateSecrets(32)
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()
	if err != nil {
		r.logger.Error("error generating new secret: %v", err)
	}
	err = r.updateSecrets(ctx, expireAt, newSecret)
	if err != nil {
		r.logger.Error("error updating secrets: %v", err)
	}

	r.logger.Status("secrets updated successfully")
}

func (r *RefreshSecrets) Stop() {
	r.exitChan <- true
}

func (r *RefreshSecrets) Start() {
	r.refresh()
	ticker := time.NewTicker(1 * time.Hour)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			go r.refresh()
		case <-r.exitChan:
			r.logger.Status("Refresh Secrets Stopped")
			return
		}
	}
}
