package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func main() {
	// Create a new Redis client
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis server address
		Password: "",               // Redis password
		DB:       0,                // Redis database index
	})

	// Set a namespace/prefix for the keys
	namespace := "myapp:"

	// Set a key-value pair with namespace
	err := rdb.Set(context.Background(), namespace+"greeting", "Hello, Redis with Namespace!", 0).Err()
	if err != nil {
		fmt.Println("Failed to set key-value in Redis:", err)
		return
	}
	fmt.Println("Key-value set successfully!")

	// Get the value for a key with namespace from Redis
	value, err := rdb.Get(context.Background(), namespace+"greeting").Result()
	if err != nil {
		fmt.Println("Failed to get value from Redis:", err)
		return
	}
	fmt.Println("Value from Redis:", value)

	// Close the Redis client connection
	err = rdb.Close()
	if err != nil {
		fmt.Println("Failed to close Redis connection:", err)
		return
	}
	fmt.Println("Redis connection closed successfully!")
}
