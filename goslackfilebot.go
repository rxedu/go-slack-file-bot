package goslackfilebot

import (
	"fmt"

	"github.com/slack-go/slack"
)

func StartBot(botToken string, channelID string, fileName string) {
	client := slack.New(botToken)
	uploadFiles(client, channelID, []string{fileName})
}

func uploadFiles(client *slack.Client, channelID string, files []string) {
	for i := 0; i < len(files); i++ {
		params := slack.FileUploadParameters{
			Channels: []string{channelID},
			File:     files[i],
		}
		file, err := client.UploadFile(params)
		if err != nil {
			fmt.Printf("%s\n", err)
		}
		fmt.Printf(file.Name)
	}
}
