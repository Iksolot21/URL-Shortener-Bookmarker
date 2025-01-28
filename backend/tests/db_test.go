package tests

import (
 	"testing"
    "github.com/Iksolot21/URL-Shortener-Bookmarker/backend/db"
      "github.com/Iksolot21/URL-Shortener-Bookmarker/backend/config"
   _ "github.com/lib/pq"
  )

  func TestOpenDB(t *testing.T) {
   cfg, err := config.LoadConfig()
    if err != nil {
       t.Fatalf("Failed to load config: %v", err)
   }
	database, err := db.OpenDB(cfg.DatabaseURL)
    if err != nil {
       t.Fatalf("Failed to connect to db: %v", err)
    }
   defer database.Close()
 }
  func TestOpenRedis(t *testing.T) {
      cfg, err := config.LoadConfig()
   if err != nil {
          t.Fatalf("Failed to load config: %v", err)
     }
	redisClient, err := db.OpenRedis(cfg.RedisURL)
    if err != nil {
      t.Fatalf("Failed to connect to redis: %v", err)
  }
     defer redisClient.Close()
 }