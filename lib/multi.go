package hermes

import (
	"context"
	"golang.org/x/sync/errgroup"
)

type MultiNotifier struct {
	Notifiers []Notifier
}

func NewMultiNotifier(config Config) Notifier {
	var notifiers []Notifier
	for _, n := range config.GetMultiple("notifiers") {
		subconfig := n.(Config)
		notifier := GetNotifier(&subconfig)
		notifiers = append(notifiers, notifier)
	}
	return &MultiNotifier{notifiers}
}

func (m *MultiNotifier) Notify(cmd *Command) error {
	eg, _ := errgroup.WithContext(context.Background())
	for _, notifier := range m.Notifiers {
		n := notifier
		eg.Go(func() error {
			return n.Notify(cmd)
		})
	}
	return eg.Wait()
}
