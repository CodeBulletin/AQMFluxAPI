package ntfy

import (
	"github.com/codebulletin/AQMFluxAPI/config"
	"github.com/codebulletin/AQMFluxAPI/types"
)

type NTFY struct {
	NTFYConfig *config.NTFYConfig
}

func New() *NTFY {
	ntfyconfig := config.GetNTFYConfig()

	ntfyconfig.Load()

	return &NTFY{
		NTFYConfig: ntfyconfig,
	}
}

func (n *NTFY) Send(message types.Message) {
	// auth := fmt.Sprintf("Bearer %s", n.NTFYConfig.Token)

	// err := utils.HTTPPost(
	// 	fmt.Sprintf("%s/%s",
	// 		n.NTFYConfig.Host,
	// 		message.Topic,
	// 	),
	// 	[]byte(message.Payload),
	// 	map[string]string{
	// 		"priority":      fmt.Sprintf("%d", message.Priority),
	// 		"title":         message.Title,
	// 		"tags":          utils.JoinStrings(message.Tags, ","),
	// 		"Authorization": auth,
	// 	},
	// 	5*time.Second,
	// )

	// if err != nil {
	// 	log.Printf(
	// 		"Error sending message to %s: %s\n",
	// 		fmt.Sprintf("%s/%s",
	// 			n.NTFYConfig.Host,
	// 			message.Topic,
	// 		),
	// 		err,
	// 	)
	// }

	// log.Printf(
	// 	"Message sent to %s\n",
	// 	fmt.Sprintf("%s/%s",
	// 		n.NTFYConfig.Host,
	// 		message.Topic,
	// 	),
	// )
}
