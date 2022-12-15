// https://redis.uptrace.dev/guide/go-redis-pubsub.html
package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"sync"
	"time"
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

	// launch anonymous function
	go func(ctx context.Context, client *redis.Client) {
		defer wg.Done()

		i := 0
		for {
			message := fmt.Sprintf("Hello World #%d", i)
			err := client.Publish(ctx, "chat-room", message).Err()
			if err != nil {
				panic(err)
			}
			time.Sleep(time.Second)
			i++
		}
	}(ctx, rdb)

	wg.Wait()
}
