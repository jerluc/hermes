package hermes

type Notifier interface {
	Success(cmd *Command)
	Failure(cmd *Command, err error)
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
	case "slack":
		return NewSlackNotifier(notifierConfig)
	default:
		return nil
	}
}
