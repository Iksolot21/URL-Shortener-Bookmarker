package db

  import (
  	"context"
  	"fmt"
  	"log"
  	"github.com/go-redis/redis/v8"

  )

  func OpenRedis(redisUrl string) (*redis.Client, error) {
  	opt, err := redis.ParseURL(redisUrl)
  	if err != nil {
  		return nil,  fmt.Errorf("failed to parse redis url: %w", err)
  	}


  	redisClient := redis.NewClient(opt)
   	err = redisClient.Ping(context.Background()).Err()
     if err != nil {
          return nil,  fmt.Errorf("unable to reach redis: %w", err)
      }
  	log.Println("Redis opened")
     return redisClient, nil
  }