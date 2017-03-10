package hermes

import (
	"strings"
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

func (d DesktopNotifier) Failure(cmd *Command, err error) {
	d.notifier.Push(
		strings.Join(cmd.Cmd.Args, " "),
		cmd.Stderr.String(),
		"",
		notificator.UR_CRITICAL,
	)
}

func (d DesktopNotifier) Success(cmd *Command) {
	d.notifier.Push(
		strings.Join(cmd.Cmd.Args, " "),
		cmd.Stdout.String(),
		"",
		notificator.UR_NORMAL,
	)
}
