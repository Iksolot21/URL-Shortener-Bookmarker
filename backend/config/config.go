package config

import (
    "fmt"
    "os"
)

type Config struct {
    DatabaseURL string
    RedisURL    string
    ServerPort  string
    ServerHost string
    JWTSecret   string
    FrontendURL string
}

func LoadConfig() (Config, error) {
    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")
    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbName := os.Getenv("DB_NAME")

    redisHost := os.Getenv("REDIS_HOST")
    redisPort := os.Getenv("REDIS_PORT")

    serverPort := os.Getenv("SERVER_PORT")
     serverHost := os.Getenv("SERVER_HOST")


     jwtSecret := os.Getenv("JWT_SECRET")

     frontendURL := os.Getenv("FRONTEND_URL")


    databaseURL := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        dbHost, dbPort, dbUser, dbPassword, dbName)

    redisURL := fmt.Sprintf("%s:%s", redisHost, redisPort)

    return Config{
        DatabaseURL: databaseURL,
        RedisURL:    redisURL,
        ServerPort: serverPort,
        ServerHost: serverHost,
         JWTSecret:  jwtSecret,
         FrontendURL: frontendURL,

    }, nil
}