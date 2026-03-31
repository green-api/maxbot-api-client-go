package main

import (
	"context"

	"github.com/green-api/maxbot-api-client-go/pkg/api"
	"github.com/green-api/maxbot-api-client-go/pkg/client"
	"github.com/green-api/maxbot-api-client-go/pkg/models"
	"github.com/rs/zerolog/log"
)

func main() {
	bot, err := api.New(client.Config{
		BaseURL: "https://platform-api.max.ru",
		Token:   "YOUR_BOT_TOKEN",
	})
	if err != nil {
		log.Fatal().Err(err).Msg("failed to init MAX API")
	}

	const exampleChatID int64 = 123456789 // recipient user ID

	_, err = bot.Helpers.SendFile(context.Background(), models.SendFileReq{
		ChatID:     exampleChatID,
		FileSource: "cmd/examples/assets/file.txt",
	})
	if err != nil {
		log.Error().Err(err).Msg("failed to send local file")
		return
	}

	_, err = bot.Helpers.SendFile(context.Background(), models.SendFileReq{
		ChatID:     exampleChatID,
		FileSource: "https://storage.yandexcloud.net/sw-prod-03-test/ChatBot/corgi.jpg",
	})
	if err != nil {
		log.Error().Err(err).Msg("failed to send file by URL")
	}
}
