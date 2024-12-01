package ntfy

import (
	"context"
	"fmt"

	"github.com/codebulletin/AQMFluxAPI/config"
	"github.com/codebulletin/AQMFluxAPI/logger"
	"github.com/codebulletin/AQMFluxAPI/types"
	"github.com/codebulletin/AQMFluxAPI/utils"
)

type NTFY struct {
	NTFYConfig *config.NTFYConfig
	logger     logger.Logger
}

func New() *NTFY {
	ntfyconfig := config.GetNTFYConfig()

	ntfyconfig.Load()

	logger := logger.GetLogger()

	return &NTFY{
		NTFYConfig: ntfyconfig,
		logger:     logger,
	}
}

func (n *NTFY) Send(ctx context.Context, message types.Message) {
	auth := fmt.Sprintf("Bearer %s", n.NTFYConfig.Token)

	err := utils.HTTPPost(
		fmt.Sprintf("%s/%s",
			n.NTFYConfig.Host,
			message.Topic,
		),
		[]byte(message.Payload),
		map[string]string{
			"priority":      fmt.Sprintf("%d", message.Priority),
			"title":         message.Title,
			"tags":          message.Tags,
			"Authorization": auth,
		},
		ctx,
		n.logger,
	)

	if err != nil {
		n.logger.Error("Error sending message: %v", err)
	}

	n.logger.Info("Message sent: %v", message)
}
