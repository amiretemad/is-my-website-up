package MessageChannels

import (
	"bytes"
	"encoding/json"
	"main/Structs"
	"net/http"
	"os"
)

type slackMessage struct {
	Text string `json:"text"`
}

type SlackChannel struct {
	L Structs.LinkStruct
}

func (sc *SlackChannel) SendMessage() bool {

	if os.Getenv("SLACK_HOOK_URL") == "" {
		return false
	}

	message := slackMessage{Text: "*" + sc.L.Description + "* Is Down!"}
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
