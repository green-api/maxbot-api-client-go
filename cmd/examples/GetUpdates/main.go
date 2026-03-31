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

	response, err := bot.Subscriptions.GetUpdates(context.Background(), &models.GetUpdatesReq{
		Limit:   10,
		Timeout: 10,
		Types: []models.UpdateType{
			"bot_added",
			"bot_removed",
			"bot_started",
			"bot_stopped",
		},
	})
	if err != nil {
		log.Error().Msgf("GetUpdates error: %v", err)
	}
	log.Info().Msgf("New update received: %v", response)
}
