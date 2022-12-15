// https://redis.uptrace.dev/guide/go-redis-pubsub.html
package main

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
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

		// process items from the queue
		for {
			result, err := client.BRPop(ctx, 0, "work-queue").Result()
			if err != nil {
				panic(err)
			}

			// process the item from the queue
			item := result[1]
			log.Println("Processing item:", item)
		}
	}(ctx, rdb)

	wg.Wait()
}
