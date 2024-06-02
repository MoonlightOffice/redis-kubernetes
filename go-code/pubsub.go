package main

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

func PubSub() {
	ctx := context.Background()

	rdb := GetClusterClient()

	go subscribe(ctx, rdb, "mychannel")
	time.Sleep(time.Second)

	for i := 0; i < 5; i++ {
		message := fmt.Sprintf("Message %d", i)
		err := publish(ctx, rdb, "mychannel", message)
		if err != nil {
			fmt.Println("Publish error:", err)
		}
		time.Sleep(time.Second)
	}
}

func publish(ctx context.Context, rdb *redis.ClusterClient, channel, message string) error {
	err := rdb.Publish(ctx, channel, message).Err()
	if err != nil {
		return err
	}

	return nil
}

func subscribe(ctx context.Context, rdb *redis.ClusterClient, channel string) {
	sub := rdb.Subscribe(ctx, channel)
	defer sub.Unsubscribe(ctx, channel)

	ch := sub.Channel()
	fmt.Printf("Listening %s...\n", channel)

	for {
		select {
		case msg, ok := <-ch:
			if !ok {
				fmt.Println("Channel closed")
				return
			}
			fmt.Printf("Received message from %s: Message: %s\n", msg.Channel, msg.Payload)
		case <-ctx.Done():
			fmt.Println("Check:", ctx.Err().Error())
			return
			//default:
			//	time.Sleep(time.Second)
			//	_, err := rdb.Ping(ctx).Result()
			//	if err != nil {
			//		fmt.Println("Err:", err.Error())
			//		return
			//	}
		}
	}
}
