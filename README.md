# これは何をするかと言うと

* パイプで渡された標準出力をSlackにポストします。

# Install

* `go get github.com/pankona/slackpost`

# Usage

* 以下の環境変数を設定します。

  * SLACK_WEBHOOK_URL
    * SlackのWebhook URL

  * 以下のコマンドを実行すると、hoge.txt の内容をSlackにポストします。

    * `$ cat hoge.txt | slackpost`

# License

* MIT

# Contribution

* Any contribution is welcome!
