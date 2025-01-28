package tests

import (
	"net/http"
	 "net/http/httptest"
	 "strings"
	  "testing"
	  "github.com/gorilla/mux"

	  "github.com/Iksolot21/URL-Shortener-Bookmarker/backend/db"
	  "github.com/Iksolot21/URL-Shortener-Bookmarker/backend/config"
	  "github.com/Iksolot21/URL-Shortener-Bookmarker/backend/handlers"

	  _ "github.com/lib/pq"
	  "encoding/json"
)

 func TestShortenURLHandler(t *testing.T) {
   cfg, err := config.LoadConfig()
	 if err != nil {
		  t.Fatalf("Failed to load config: %v", err)
	 }
	  database, err := db.OpenDB(cfg.DatabaseURL)
	 if err != nil {
		  t.Fatalf("Failed to connect to db: %v", err)
	 }
   redisClient, err := db.OpenRedis(cfg.RedisURL)
	  if err != nil {
		  t.Fatalf("Failed to connect to redis: %v", err)
	  }

	 r := mux.NewRouter()
	r.HandleFunc("/shorten", handlers.ShortenURL(database, redisClient)).Methods("POST")

	 reqBody := `{"original_url":"https://www.test.com"}`
	req, err := http.NewRequest("POST", "/shorten", strings.NewReader(reqBody))
	if err != nil {
		 t.Fatalf("Failed to create request: %v", err)
	 }

	rr := httptest.NewRecorder()
	 r.ServeHTTP(rr, req)
   if status := rr.Code; status != http.StatusCreated {
	  t.Errorf("Handler return wrong status code: got %v, want %v", status, http.StatusCreated )
	}
	 var resp map[string]string
   err = json.NewDecoder(rr.Body).Decode(&resp)
  if err != nil {
		t.Fatalf("Failed to decode response: %v", err)
   }
	if _, ok := resp["short_url"]; !ok {
		t.Error("Handler did not return short_url")
	 }
 }