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

	_, err = bot.Messages.SendMessage(context.Background(), models.SendMessageReq{
		UserID: exampleUserID,
		Text:   "Hello world!",
	})
	if err != nil {
		log.Error().Msgf("SendMessage error: %v", err)
	}
	log.Info().Msgf("SendMessage success!")
}
