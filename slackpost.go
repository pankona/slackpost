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
	Text     string `json:"text"`     //投稿内容
	Username string `json:"username"` //投稿者名 or Bot名（存在しなくてOK）
	Channel  string `json:"channel"`  //#部屋名
}

func main() {
	webhookUrl := getEnvVar("SLACK_WEBHOOK_URL")
	if webhookUrl == "" {
		fmt.Println("SLACK_WEBHOOK_URL is not specified.")
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
			"red",
			"#alone"})

	resp, _ := http.PostForm(
		webhookUrl,
		url.Values{"payload": {string(params)}},
	)

	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	fmt.Println(string(body))
}
