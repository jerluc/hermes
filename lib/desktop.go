package hermes

import (
	"github.com/0xAX/notificator"
)

type DesktopNotifier struct {
	notifier *notificator.Notificator
}

// Creates a new desktop-based notification using any one of the
// available desktop notification libraries available for your system
// (e.g. Growl on Windows, notify-send on Linux, etc.). To configure
// this notifier, you should use the following configuration:
//
// notifier:
//   type: desktop
//
func NewDesktopNotifier(config Config) Notifier {
	n := notificator.New(notificator.Options{
		AppName: "Hermes Notification",
	})
	return &DesktopNotifier{n}
}

func (d DesktopNotifier) Notify(cmd *Command) error {
	if cmd.Successful() {
		d.notifier.Push(
			cmd.CmdLine(),
			cmd.PeekStdout(1),
			"",
			notificator.UR_NORMAL,
		)
	} else {
		d.notifier.Push(
			cmd.CmdLine(),
			cmd.PeekStderr(1),
			"",
			notificator.UR_CRITICAL,
		)
	}
	return nil
}
