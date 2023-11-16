package graphql

import (
	"context"
	"time"
)

type subscriptionResolver struct {
	*rootResolver
}

func (r *subscriptionResolver) Ping(ctx context.Context) (<-chan string, error) {
	var (
		ch = make(chan string)
	)

	go func() {
		time.Sleep(3 * time.Second)

		ch <- "pong"
	}()

	return ch, nil
}

func (r *rootResolver) Subscription() SubscriptionResolver { return &subscriptionResolver{r} }
