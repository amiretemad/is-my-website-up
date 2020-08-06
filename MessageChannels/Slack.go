package MessageChannels

import (
	"bytes"
	"encoding/json"
	"main/Structs"
	"net/http"
	"os"
)

type SlackChannel struct {
	l Structs.LinkStruct
}

type slackMessage struct {
	Text string `json:"text"`
}

func NewSlackChannel(l Structs.LinkStruct) SlackChannel {
	return SlackChannel{l}
}

func (sc *SlackChannel) SendMessage() bool {

	message := slackMessage{Text: "*" + sc.l.Description + "* Is Down!"}
	slackMessageJson, _ := json.Marshal(message)

	request, _ := http.NewRequest("POST",
		os.Getenv("SLACK_HOOK_URL"),
		bytes.NewBuffer(slackMessageJson))

	request.Header.Set("Content-Type", "application/json")

	_, err := http.DefaultClient.Do(request)

	if err != nil {
		return false
	}

	return true
}
