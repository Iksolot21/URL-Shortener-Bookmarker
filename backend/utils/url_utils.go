package utils

import (
    "context"
    "crypto/rand"
    "encoding/base64"
    "fmt"
     "database/sql"
      "github.com/go-redis/redis/v8"
    "log"
)

func GenerateShortURL() (string, error) {
    b := make([]byte, 6)
    _, err := rand.Read(b)
    if err != nil {
        return "", err
    }
    shortURL := base64.RawURLEncoding.EncodeToString(b)
    return shortURL, nil
}

func CacheURL(ctx context.Context, client *redis.Client, shortURL, originalURL string) error {
   err := client.Set(ctx, shortURL, originalURL, 0).Err()
	if err != nil {
	  	 log.Println(err)
     	  return fmt.Errorf("failed to save url to cache: %w", err)
   	}
	   return nil

}
func GetOriginalURL(ctx context.Context, client *redis.Client, db *sql.DB, shortURL string) (string, error) {

	originalURL, err := client.Get(ctx, shortURL).Result()
  if err == nil{
    return originalURL, nil
  }
   if err != redis.Nil{
    log.Println(err)
     return "", fmt.Errorf("failed to get from cache: %w", err)
   }
    query := `SELECT original_url FROM urls WHERE short_url = $1`
     row := db.QueryRow(query, shortURL)
    err = row.Scan(&originalURL)
     if err != nil {
         return "", fmt.Errorf("failed to get from database: %w", err)
     }

   return originalURL, nil
 }