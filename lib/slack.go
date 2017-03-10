package hermes

import (
	// "bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type slackMessage struct {
	Sender      string       `json:"username"`
	Recipient   string       `json:"channel"`
	Markdown    bool         `json:"mrkdown"`
	Text        string       `json:"text,omitempty"`
	Attachments []attachment `json:"attachments,omitempty"`
}

type attachment struct {
	Fields     []field  `json:"fields,omitempty"`
	Color      string   `json:"color,omitempty"`
	Title      string   `json:"title,omitempty"`
	Text       string   `json:"text,omitempty"`
	MarkdownIn []string `json:"mrkdwn_in,omitempty"`
}

type field struct {
	Title string `json:"title"`
	Value string `json:"value"`
	Short bool   `json:"short"`
}

type SlackNotifier struct {
	webHookUrl string
	sender     string
	recipient  string
}

// Creates a new Slack-enabled notifier. This expects a minimal Hermes
// configuration that looks like:
//
// notifier:
//   type: slack
//   webHookUrl: <SLACK_WEB_HOOK_URL>
//   sender: <SLACK_SENDER_NAME>
//   recipient: <SLACK_RECIPIENT_NAME>
//
func NewSlackNotifier(config Config) Notifier {
	return &SlackNotifier{config.Get("webHookUrl"), config.Get("sender", "hermes"), config.Get("recipient")}
}

func codeOrEmpty(code string) string {
	if len(code) > 0 {
		return fmt.Sprintf("```%s```", code)
	}
	return "_<empty>_"
}

func (s SlackNotifier) Success(cmd *Command) {
	msg := slackMessage{
		Sender:    s.sender,
		Recipient: s.recipient,
		Markdown:  true,
		Text:      "*Hermes Command Completed!*",
		Attachments: []attachment{
			attachment{
				Title:      "Command",
				Text:       fmt.Sprintf("```%s```", strings.Join(cmd.Cmd.Args, " ")),
				MarkdownIn: []string{"text"},
			},
			attachment{
				Title: "Timing",
				Fields: []field{
					field{Title: "CPU", Value: cmd.Cmd.ProcessState.SystemTime().String()},
					field{Title: "User", Value: cmd.Cmd.ProcessState.UserTime().String()},
				},
				MarkdownIn: []string{"text"},
			},
			attachment{
				Title:      "Standard out",
				Text:       codeOrEmpty(cmd.Stdout.String()),
				MarkdownIn: []string{"text"},
			},
			attachment{
				Title:      "Standard error",
				Text:       codeOrEmpty(cmd.Stderr.String()),
				MarkdownIn: []string{"text"},
			},
		},
	}

	s.postToSlack(msg)
}

func (s SlackNotifier) Failure(cmd *Command, err error) {
	// NOT IMPLEMENTED!
}

func (s SlackNotifier) postToSlack(msg slackMessage) error {
	dataBuffer, _ := json.Marshal(&msg)
	resp, respErr := http.Post(s.webHookUrl, "application/json", strings.NewReader(string(dataBuffer)))
	if resp != nil {
		defer resp.Body.Close()
	}
	return respErr
}
