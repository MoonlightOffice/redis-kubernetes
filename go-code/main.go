package main

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

func main() {
	Subscribe2()
}

func GetClient() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:7777",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return rdb
}

func SwitchDatabase(rdb *redis.Client, db int) error {
	return rdb.Do(context.TODO(), "SELECT", db).Err()
}

func Transaction() {
	ctx := context.Background()
	rdb := GetClient()

	// Transaction: WATCH, MULTI, EXEC, DISCARD
	rdb.Watch(ctx, func(tx *redis.Tx) error {
		tx.Pipelined(ctx, func(p redis.Pipeliner) error {
			return nil
		})

		return nil
	}, "", "")

}

func ExampleClient() {
	ctx := context.Background()

	rdb := GetClient()

	err := rdb.Set(ctx, "key", "value", 3*time.Second).Err()
	if err != nil {
		fmt.Println(err)
		return
	}

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := rdb.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
	// Output: key value
	// key2 does not exist
}
