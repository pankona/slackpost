package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func getEnvVar(varName string) (result string) {
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		if pair[0] == varName {
			return pair[1]
		}
	}
	return ""
}

type Slack struct {
	Text     string `json:"text"`
	Username string `json:"username"`
	Channel  string `json:"channel"`
}

const ENVKEY_SLACKPOST_WEBHOOK_URL = "SLACKPOST_WEBHOOK_URL"
const ENVKEY_SLACKPOST_USERNAME = "SLACKPOST_USERNAME"
const ENVKEY_SLACKPOST_CHANNEL_TO_POST = "SLACKPOST_CHANNEL_TO_POST"

func main() {
	slackpostWebhookUrl := getEnvVar(ENVKEY_SLACKPOST_WEBHOOK_URL)
	if slackpostWebhookUrl == "" {
		fmt.Println(ENVKEY_SLACKPOST_WEBHOOK_URL, "is not specified.")
		os.Exit(1)
	}

	slackpostUserName := getEnvVar(ENVKEY_SLACKPOST_USERNAME)
	if slackpostUserName == "" {
		fmt.Println(ENVKEY_SLACKPOST_USERNAME, "is not specified.")
		os.Exit(1)
	}

	slackpostChannelToPost := getEnvVar(ENVKEY_SLACKPOST_CHANNEL_TO_POST)
	if slackpostChannelToPost == "" {
		fmt.Println(ENVKEY_SLACKPOST_CHANNEL_TO_POST, "is not specified.")
		os.Exit(1)
	}

	in := os.Stdin
	var buf string
	reader := bufio.NewReaderSize(in, 4096)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("failed to read from stdin. err =", err)
			os.Exit(1)
		}
		buf += string(line) + "\n"
	}

	params, _ := json.Marshal(
		Slack{
			buf,
			slackpostUserName,
			slackpostChannelToPost})

	resp, _ := http.PostForm(
		slackpostWebhookUrl,
		url.Values{"payload": {string(params)}},
	)

	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	fmt.Println(string(body))
}
