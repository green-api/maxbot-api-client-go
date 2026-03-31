package main

import (
	"context"

	"github.com/rs/zerolog/log"

	"github.com/green-api/maxbot-api-client-go/pkg/api"
	"github.com/green-api/maxbot-api-client-go/pkg/client"
)

func main() {
	bot, err := api.New(client.Config{
		BaseURL: "https://platform-api.max.ru",
		Token:   "YOUR_BOT_TOKEN",
	})
	if err != nil {
		log.Fatal().Err(err).Msg("failed to init MAX API")
	}

	response, err := bot.Bots.GetBot(context.Background())
	if err != nil {
		log.Error().Err(err).Msg("GetBot error")
		return
	}
	log.Info().Interface("info", response).Msg("Bot info received")
}
