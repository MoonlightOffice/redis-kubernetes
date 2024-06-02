package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func GetClusterClient() *redis.ClusterClient {
	return redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{
			"redis-cluster-0.redis:7777",
			"redis-cluster-1.redis:7777",
			"redis-cluster-2.redis:7777",
		},
	})
}

func PopulateCluster() {
	ctx := context.Background()

	rdb := GetClusterClient()

	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("key%d", i)
		val := fmt.Sprintf("val%d", i)
		fmt.Println(rdb.Set(ctx, key, val, 0).Result())
	}

	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("key%d", i)
		fmt.Println(rdb.Get(ctx, key).Result())
	}
}
