package api

import (
	"context"
	"net/http"

	cl "github.com/green-api/maxbot-api-client-go/pkg/client"
	"github.com/green-api/maxbot-api-client-go/pkg/models"
)

type bots struct {
	client *cl.Client
}

/*
BotInfo retrieves information about the current bot, such as its user ID, name, and description

Example:
info, err := api.Bots.BotInfo(ctx)

	if err == nil {
		fmt.Println("Connected as:", info.Name)
	}
*/
func (api *bots) GetBot(ctx context.Context) (*models.BotInfo, error) {
	return decode[models.BotInfo](ctx, api.client, http.MethodGet, models.PathMe)
}

/*
PatchBot edits current bot info
Fill only the fields you want to update - all remaining fields will stay untouched

Example:

	newinfo, err := api.Bots.PatchBot(ctx, &models.BotPatch{
			Name: "New Name",
			Description: "New description",
	})
*/
func (api *bots) PatchBot(ctx context.Context, req *models.BotPatch) (*models.BotInfo, error) {
	return decode[models.BotInfo](ctx, api.client, http.MethodPatch, models.PathMe, cl.WithBody(req))
}
