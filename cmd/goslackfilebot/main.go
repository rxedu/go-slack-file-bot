package main

import (
	"os"

	"github.com/rxedu/go-slack-file-bot"
)

func main() {
	goslackfilebot.StartBot(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_CHANNEL_ID"))
}
