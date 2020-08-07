package MessageChannels

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"main/Structs"
	"net/http"
	"os"
)

type mattermostMessage struct {
	Text string `json:"text"`
}

type MatterMostChannel struct {
	L Structs.LinkStruct
}

func (sc *MatterMostChannel) SendMessage() bool {

	if os.Getenv("MATTERMOST_HOOK_URL") == "" {
		return false
	}

	message := mattermostMessage{Text: "*" + sc.L.Description + "* Is Down!"}
	slackMessageJson, _ := json.Marshal(message)

	request, _ := http.NewRequest("POST",
		os.Getenv("MATTERMOST_HOOK_URL"),
		bytes.NewBuffer(slackMessageJson))

	request.Header.Set("Content-Type", "application/json")

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	_, err := http.DefaultClient.Do(request)

	if err != nil {
		fmt.Println(err)
		return false
	}

	return true
}
