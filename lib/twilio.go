package hermes

import (
	"fmt"
	"strings"

	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

type TwilioNotifier struct {
	client    *twilio.RestClient
	recipient string
	sender    string
}

func NewTwilioNotifier(config Config) Notifier {
	accountSID := config.Get("accountSID")
	authToken := config.Get("authToken")
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSID,
		Password: authToken,
	})

	return &TwilioNotifier{client, config.Get("recipient"), config.Get("sender")}
}

func (t *TwilioNotifier) Notify(cmd *Command) error {
	msg := strings.Join([]string{
		"Command:", cmd.CmdLine(),
		"",
		"Exit code:", fmt.Sprintf("%v", cmd.ExitCode()),
		"",
		"Stdout:",
		cmd.PeekStdout(10),
		"",
		"Stderr:",
		cmd.PeekStderr(10),
		"",
	}, "\n")

	params := &openapi.CreateMessageParams{}
	params.SetTo(t.recipient)
	params.SetFrom(t.sender)
	params.SetBody(msg)

	_, err := t.client.Api.CreateMessage(params)
	return err
}
