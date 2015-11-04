# これは何をするかと言うと

* パイプで渡された標準出力をSlackにポストします。

# Install

* `go get github.com/pankona/slackpost`

# Usage

* 以下の環境変数を設定します。

  * SLACK_WEBHOOK_URL
    * SlackのWebhook URL
  * SLACK_BOT_USERNAME
    * 誰名義でのポストにするか
  * SLACK_BOT_CHANNELTOPOST
    * どのチャンネルにポストするか

  * 以下のコマンドを実行すると、hoge.txt の内容をSlackにポストします。
    * `$ cat hoge.txt | slackpost`

  * 以下のコマンドを実行すると、hoge とSlackにポストします。
    * `$ echo "hoge" | slackpost`

# License

* MIT

# Contribution

* Any contribution is welcome!
