package api

import (
	"context"
	"encoding/json"
	"fmt"

	cl "github.com/green-api/maxbot-api-client-go/pkg/client"
)

type API struct {
	Bots          *bots
	Chats         *chats
	Helpers       *helpers
	Messages      *messages
	Subscriptions *subscriptions
	Uploads       *uploads

	Client *cl.Client
}

/*
New initializes a new API client instance with the provided configuration
*/
func New(cfg cl.Config) (*API, error) {
	client, err := cl.NewClient(cfg)
	if err != nil {
		return nil, err
	}

	return &API{
		Bots:          &bots{client},
		Chats:         &chats{client},
		Helpers:       &helpers{client},
		Messages:      &messages{client},
		Subscriptions: &subscriptions{client},
		Uploads:       &uploads{client},

		Client: client,
	}, nil
}

/*
decode is a generic internal helper function for executing HTTP requests
*/
func decode[T any](ctx context.Context, client *cl.Client, method, path string, opts ...cl.RequestOption) (*T, error) {
	data, err := client.Request(ctx, method, path, opts...)
	if err != nil {
		return nil, err
	}

	var res T
	if err := json.Unmarshal(data, &res); err != nil {
		return nil, fmt.Errorf("decode json error: %w", err)
	}
	if len(data) == 0 {
		return &res, nil
	}

	return &res, nil
}
