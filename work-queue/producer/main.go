// https://redis.uptrace.dev/guide/go-redis-pubsub.html
package main

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
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
			err := client.LPush(ctx, "work-queue", i).Err()
			if err != nil {
				log.Fatal(err)
			}
			i++
			time.Sleep(time.Second)
		}
	}(ctx, rdb)

	wg.Wait()
}
