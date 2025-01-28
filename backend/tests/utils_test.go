package tests

 import (
    "testing"
   "github.com/Iksolot21/URL-Shortener-Bookmarker/backend/utils"
)

 func TestGenerateShortUrl(t *testing.T) {
   shortUrl, err := utils.GenerateShortURL()
    if err != nil {
       t.Errorf("Failed to generate short url, %v", err)
    }

     if len(shortUrl) == 0 {
        t.Error("Generated short url is empty")
     }
 }