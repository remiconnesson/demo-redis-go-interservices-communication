// https://redis.uptrace.dev/guide/go-redis-pubsub.html
package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"sync"
)

func main() {
	ctx := context.Background()

	var wg sync.WaitGroup
	wg.Add(1)

	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// we declare the function anonymously and start executing
	// it right away
	go func(ctx context.Context, client *redis.Client) {
		defer wg.Done()
		pubsub := client.Subscribe(ctx, "chat-room")
		defer pubsub.Close()

		ch := pubsub.Channel()

		for msg := range ch {
			fmt.Println(msg.Channel, msg.Payload)
		}
	}(ctx, rdb)

	wg.Wait()
}
