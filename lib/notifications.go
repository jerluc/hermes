package hermes

type Notifier interface {
	Success(cmd *Command)
	Failure(cmd *Command, err error)
}

func GetNotifier(config *Config) Notifier {
	notifierConfig := config.GetConfig("notifier")
	notifierType := notifierConfig.Get("type")
	switch notifierType {
	case "slack":
		return NewSlackNotifier(notifierConfig)
	default:
		return nil
	}
}
