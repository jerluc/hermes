package hermes

import (
	"fmt"
)

type Notifier interface {
	Notify(cmd *Command) error
}

// Creates a new Hermes notifier instance from the parsed configuration
// file. The configuration's YAML structure should look something like:
//
// notifier:
//   type: <NOTIFIER_TYPE>
//   ...notifier-specific configuration goes here...
//
func GetNotifier(config *Config) Notifier {
	notifierConfig := config.GetConfig("notifier")
	notifierType := notifierConfig.Get("type")
	switch notifierType {
	case "desktop":
		return NewDesktopNotifier(notifierConfig)
	case "twilio":
		return NewTwilioNotifier(notifierConfig)
	case "multi":
		return NewMultiNotifier(notifierConfig)
	default:
		panic(fmt.Sprintf("Unknown notifier type: %s", notifierType))
	}
}
