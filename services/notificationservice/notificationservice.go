package notificationservice

import (
	"strconv"

	"github.com/codebulletin/AQMFluxAPI/logger"
	"github.com/codebulletin/AQMFluxAPI/ntfy"
)

type NotificationService struct {
	ntfy *ntfy.NTFY
	logger logger.Logger
}

func NewNotificationService(n *ntfy.NTFY) *NotificationService {
	logger := logger.GetLogger()
	return &NotificationService{
		ntfy: n,
		logger: logger,
	}
}

func (n *NotificationService) NotifyTrigger(data string) {
	defer func() {
		if r := recover(); r != nil {
			n.logger.Fatal("Recovered in NotifyTrigger: %v", r)
		}
	}()
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