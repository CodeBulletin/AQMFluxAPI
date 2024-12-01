package notificationservice

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/codebulletin/AQMFluxAPI/db"
	"github.com/codebulletin/AQMFluxAPI/logger"
	"github.com/codebulletin/AQMFluxAPI/ntfy"
	"github.com/codebulletin/AQMFluxAPI/repo"
	"github.com/codebulletin/AQMFluxAPI/types"
	"github.com/codebulletin/AQMFluxAPI/utils"
)

type NotificationService struct {
	ntfy *ntfy.NTFY
	logger logger.Logger
	db db.DB
}

func NewNotificationService(n *ntfy.NTFY, db db.DB) *NotificationService {
	logger := logger.GetLogger()
	return &NotificationService{
		ntfy: n,
		logger: logger,
		db: db,
	}
}

func (n *NotificationService) NotifyTrigger(data string) {
	defer func() {
		if r := recover(); r != nil {
			n.logger.Fatal("Recovered in NotifyTrigger: %v", r)
		}
	}()

	ctx := context.Background()
	timout_ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var alert types.AlertNotificationIncoming
	err := utils.ParseJSON(data, &alert)

	if err != nil {
		n.logger.Error("Error parsing data: %v", err)
	}

	query := repo.New(n.db)

	notification, err := query.GetMessages(timout_ctx, alert.ID)

	if err != nil {
		n.logger.Error("Error getting message: %v", err)
	}

	var tags string
	err = notification.Tags.Scan(&tags);

	if err != nil {
		tags = ""
	}

	t := time.Now()

	values := map[string]string{
		"$0": fmt.Sprintf("%v", alert.V0),
		"$1": fmt.Sprintf("%v", alert.V1),
		"$2": fmt.Sprintf("%v", alert.V2),
		"$Active": fmt.Sprintf("%v", alert.V0 == 1),
		"$Value": fmt.Sprintf("%v", alert.V0),
		"$Min": fmt.Sprintf("%v", alert.V1),
		"$Max": fmt.Sprintf("%v", alert.V2),
		"$Attribute": alert.AttrName,
		"$Device": alert.DevName,
		"$Sensor": alert.SenName,
		"$Operator": alert.OP,
		"$Alert": alert.AlertName,
		"$Time": fmt.Sprintf("%v", alert.Time),
		"$DD": fmt.Sprintf("%02d", t.Day()),
		"$MM": fmt.Sprintf("%02d", t.Month()),
		"$YYYY": fmt.Sprintf("%04d", t.Year()),
		"$hh": fmt.Sprintf("%02d", t.Hour()),
		"$mm": fmt.Sprintf("%02d", t.Minute()),
		"$ss": fmt.Sprintf("%02d", t.Second()),
		"$Location": alert.Location,
		"$Unit": alert.Unit,
	}

	replacerArgs := []string{}
	for key, val := range values {
		replacerArgs = append(replacerArgs, key, val)
	}

	replacer := strings.NewReplacer(replacerArgs...)

	var message types.Message
	message.Title = replacer.Replace(notification.Title)
	message.Topic = notification.Topic
	message.Payload = replacer.Replace(notification.Payload)
	message.Tags = tags
	message.Priority = int(notification.Messagepriority)

	n.ntfy.Send(timout_ctx, message)
}

func (n *NotificationService) ChangeFreq(freqChan chan int32) func(data string) {
	return func(data string) {
		defer func() {
			if r := recover(); r != nil {
				n.logger.Fatal("Recovered in ChangeFreq: %v", r)
			}
		}()

		freq, err := strconv.Atoi(data) 
		if err != nil {
			n.logger.Error("Error parsing data: %v", err)
		}

		freqChan <- int32(freq)
	}
}