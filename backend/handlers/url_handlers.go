package handlers

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"

    "github.com/go-redis/redis/v8"
    "github.com/gorilla/mux"
   "github.com/Iksolot21/URL-Shortener-Bookmarker/backend/utils"
   "github.com/Iksolot21/URL-Shortener-Bookmarker/backend/internal/errors"
   "database/sql"

)

func ShortenURL(db *sql.DB, redisClient *redis.Client) http.HandlerFunc {
     return func(w http.ResponseWriter, r *http.Request) {
       var req struct {
        OriginalURL string `json:"original_url"`
    }

    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        errors.RespondWithError(w, http.StatusBadRequest, "Invalid request body")
        return
    }

    shortURL, err := utils.GenerateShortURL()
     if err != nil {
         errors.RespondWithError(w, http.StatusInternalServerError, "Failed to generate short url")
         return
      }

    query := `INSERT INTO urls (short_url, original_url) VALUES ($1, $2)`
    _, err = db.Exec(query, shortURL, req.OriginalURL)
    if err != nil {
      log.Println(err)
       errors.RespondWithError(w, http.StatusInternalServerError, "Failed to save url")
        return
    }

        err = utils.CacheURL(r.Context(),redisClient, shortURL, req.OriginalURL)
         if err != nil {
             errors.RespondWithError(w, http.StatusInternalServerError, "Failed to save url to cache")
            return
         }
     w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(map[string]string{"short_url": fmt.Sprintf("http://%s:%s/%s", r.Host,r.URL.Port(),shortURL) })


    }
}


func RedirectURL(db *sql.DB, redisClient *redis.Client) http.HandlerFunc {
     return func(w http.ResponseWriter, r *http.Request) {
        vars := mux.Vars(r)
        shortURL := vars["short_url"]

         originalURL, err := utils.GetOriginalURL(r.Context(), redisClient, db, shortURL)
        if err != nil {
           errors.RespondWithError(w, http.StatusNotFound, "URL not found")
            return
        }


         http.Redirect(w, r, originalURL, http.StatusMovedPermanently)

     }
}