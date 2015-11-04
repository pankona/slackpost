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

func main() {
	slackWebhookUrl := getEnvVar("SLACK_WEBHOOK_URL")
	if slackWebhookUrl == "" {
		fmt.Println("SLACK_WEBHOOK_URL is not specified.")
		os.Exit(1)
	}

	slackUserName := getEnvVar("SLACK_BOT_USERNAME")
	if slackUserName == "" {
		fmt.Println("SLACK_BOT_USERNAME is not specified.")
		os.Exit(1)
	}

	slackChannelToPost := getEnvVar("SLACK_BOT_CHANNEL_TO_POST")
	if slackChannelToPost == "" {
		fmt.Println("SLACK_BOT_CHANNEL_TO_POST is not specified.")
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
			slackUserName,
			slackChannelToPost})

	resp, _ := http.PostForm(
		slackWebhookUrl,
		url.Values{"payload": {string(params)}},
	)

	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	fmt.Println(string(body))
}
