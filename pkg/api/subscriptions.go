package api

import (
	"context"
	"net/http"

	cl "github.com/green-api/maxbot-api-client-go/pkg/client"
	m "github.com/green-api/maxbot-api-client-go/pkg/models"
)

type subscriptions struct {
	client *cl.Client
}

/*
GetSubscriptions returns a list of all subscriptions if your bot receives data via a Webhook

Example:

	subscriptions, err := api.Updates.GetSubscriptions(ctx)
*/
func (api *subscriptions) GetSubscriptions(ctx context.Context) (*m.GetSubscriptionsResp, error) {
	return decode[m.GetSubscriptionsResp](ctx, api.client, http.MethodGet, m.PathSubscriptions)
}

/*
Subscribe configures the delivery of bot events via Webhook

Example:

	success, err := api.Updates.Subscribe(ctx, Subscribe{
		Url: https:webhook.site,
	})
*/
func (api *subscriptions) Subscribe(ctx context.Context, req m.SubscribeReq) (*m.SimpleQueryResult, error) {
	return decode[m.SimpleQueryResult](ctx, api.client, http.MethodPost, m.PathSubscriptions, cl.WithBody(req))
}

/*
Unsubscribe unsubscribes the bot from receiving updates via Webhook

Example:

	success, err := api.Updates.Unsubscribe(ctx, UnsubscribeReq{
		Url: https:webhook.site,
	})
*/
func (api *subscriptions) Unsubscribe(ctx context.Context, req m.UnsubscribeReq) (*m.SimpleQueryResult, error) {
	return decode[m.SimpleQueryResult](ctx, api.client, http.MethodDelete, m.PathSubscriptions, cl.WithQuery(req))
}

/*
GetUpdates fetches new events (incoming messages, bot additions, etc.) from the server
Use this method for long-polling
Provide an Marker to acknowledge previous updates and fetch only new ones

Example:

	resp, err := api.Updates.GetUpdates(ctx, GetUpdatesReq{
		Marker:  123456789,
		Timeout: 30,  seconds to wait for new updates
	})
*/
func (api *subscriptions) GetUpdates(ctx context.Context, req *m.GetUpdatesReq) (*m.GetUpdatesResp, error) {
	return decode[m.GetUpdatesResp](ctx, api.client, http.MethodGet, m.PathUpdates, cl.WithQuery(req))
}
