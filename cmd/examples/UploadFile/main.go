package main

import (
	"context"

	"github.com/rs/zerolog/log"

	"github.com/green-api/maxbot-api-client-go/pkg/api"
	"github.com/green-api/maxbot-api-client-go/pkg/client"
	"github.com/green-api/maxbot-api-client-go/pkg/models"
)

func main() {
	bot, err := api.New(client.Config{
		BaseURL: "https://platform-api.max.ru",
		Token:   "YOUR_BOT_TOKEN",
	})
	if err != nil {
		log.Fatal().Err(err).Msg("failed to init MAX API")
	}

	const exampleUserID int64 = 123456789 // recipient user ID

	file, err := bot.Uploads.UploadFile(context.Background(), models.UploadFileReq{
		Type:     "image",
		FilePath: "cmd/examples/assets/file.png",
	})
	if err != nil {
		log.Error().Msgf("UploadFile error: %v", err)
		return
	}

	if file.Token == "" {
		log.Warn().Msg("upload success, but file token is empty")
		return
	}

	attachment := models.AttachImage(file.Token, "")
	_, err = bot.Messages.SendMessage(context.Background(), models.SendMessageReq{
		UserID:      exampleUserID,
		Attachments: []models.Attachment{attachment},
	})
	if err != nil {
		log.Error().Msgf("Send File Message error: %v", err)
	}
	log.Info().Msg("File successfully sent to chat!")
}
