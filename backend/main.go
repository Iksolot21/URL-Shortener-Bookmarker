package main

import (
    "fmt"
    "log"
    "net/http"
    "os"

    "github.com/gorilla/mux"
    "github.com/joho/godotenv"

    "github.com/Iksolot21/URL-Shortener-Bookmarker/backend/config"
    "github.com/Iksolot21/URL-Shortener-Bookmarker/backend/db"
    "github.com/Iksolot21/URL-Shortener-Bookmarker/backend/handlers"
     "github.com/Iksolot21/URL-Shortener-Bookmarker/backend/middleware"
     "github.com/Iksolot21/URL-Shortener-Bookmarker/backend/logger"
   "github.com/Iksolot21/URL-Shortener-Bookmarker/backend/migrations"

)

func main() {
    err := godotenv.Load()
    if err != nil {
         logger.Error("Error loading .env file", err)
    }


    cfg, err := config.LoadConfig()
    if err != nil {
         logger.Error("Error loading config", err)
    }

     database, err := db.OpenDB(cfg.DatabaseURL)
    if err != nil {
          logger.Error("Error opening database", err)
    }
    defer database.Close()

    redisClient, err := db.OpenRedis(cfg.RedisURL)
    if err != nil {
         logger.Error("Error opening redis", err)
    }
    defer redisClient.Close()


    err = migrations.RunMigrations(database)
     if err != nil {
          logger.Error("Error running migrations", err)
      }


    r := mux.NewRouter()
    
     r.Use(middleware.CORSMiddleware(cfg.FrontendURL))

    // Auth Routes
    r.HandleFunc("/auth/register", handlers.RegisterUser(database)).Methods("POST")
    r.HandleFunc("/auth/login", handlers.LoginUser(database, cfg.JWTSecret)).Methods("POST")
    r.HandleFunc("/me", middleware.AuthMiddleware(handlers.GetCurrentUser(database), cfg.JWTSecret)).Methods("GET")

    // URL Routes
    r.HandleFunc("/shorten", handlers.ShortenURL(database, redisClient)).Methods("POST")
    r.HandleFunc("/{short_url}", handlers.RedirectURL(database, redisClient)).Methods("GET")

    // Bookmark Routes
    r.HandleFunc("/bookmarks", middleware.AuthMiddleware(handlers.GetBookmarks(database), cfg.JWTSecret)).Methods("GET")
    r.HandleFunc("/bookmarks", middleware.AuthMiddleware(handlers.CreateBookmark(database), cfg.JWTSecret)).Methods("POST")
    r.HandleFunc("/bookmarks/{id}", middleware.AuthMiddleware(handlers.GetBookmarkById(database), cfg.JWTSecret)).Methods("GET")
    r.HandleFunc("/bookmarks/{id}", middleware.AuthMiddleware(handlers.PatchBookmarkById(database), cfg.JWTSecret)).Methods("PATCH")
    r.HandleFunc("/bookmarks/{id}", middleware.AuthMiddleware(handlers.DeleteBookmarkById(database), cfg.JWTSecret)).Methods("DELETE")
    r.HandleFunc("/search", middleware.AuthMiddleware(handlers.SearchBookmarks(database), cfg.JWTSecret)).Methods("GET")

   log.Println(fmt.Sprintf("server starting at %s:%s", cfg.ServerHost, cfg.ServerPort ))
    err = http.ListenAndServe(fmt.Sprintf("%s:%s", cfg.ServerHost, cfg.ServerPort), r)
     if err != nil {
        logger.Error("Error starting server", err)
         os.Exit(1)
    }


}