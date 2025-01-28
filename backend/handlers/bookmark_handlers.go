package handlers

import (
 "database/sql"
     "encoding/json"
      "net/http"
      "strconv"
        "github.com/gorilla/mux"
        "github.com/Iksolot21/URL-Shortener-Bookmarker/models"
         "github.com/Iksolot21/URL-Shortener-Bookmarker/services"
         "github.com/Iksolot21/URL-Shortener-Bookmarker/errors"
    )

func GetBookmarks(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
      user, ok := r.Context().Value("user").(models.User)
        if !ok{
             errors.RespondWithError(w, http.StatusUnauthorized, "Not authorized")
             return
        }
      bookmarks, err := services.GetBookmarks(db, user.ID)
     if err != nil {
        errors.RespondWithError(w, http.StatusInternalServerError, "Could not get bookmarks")
         return
     }
         w.Header().Set("Content-Type", "application/json")
         json.NewEncoder(w).Encode(bookmarks)


    }
}

 func CreateBookmark(db *sql.DB) http.HandlerFunc {
      return func(w http.ResponseWriter, r *http.Request) {
          user, ok := r.Context().Value("user").(models.User)
           if !ok{
                errors.RespondWithError(w, http.StatusUnauthorized, "Not authorized")
              return
          }
          var bookmark models.Bookmark
         if err := json.NewDecoder(r.Body).Decode(&bookmark); err != nil {
            errors.RespondWithError(w, http.StatusBadRequest, "Invalid request body")
             return
          }
          err := services.CreateBookmark(db, &bookmark, user.ID)
           if err != nil {
              errors.RespondWithError(w, http.StatusInternalServerError, "Could not create bookmark")
               return
          }
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusCreated)
        json.NewEncoder(w).Encode(map[string]string{"message": "Bookmark created successfully"})

     }
 }

 func GetBookmarkById(db *sql.DB) http.HandlerFunc {
      return func(w http.ResponseWriter, r *http.Request) {
          user, ok := r.Context().Value("user").(models.User)
           if !ok{
              errors.RespondWithError(w, http.StatusUnauthorized, "Not authorized")
               return
          }
          vars := mux.Vars(r)
          id, err := strconv.Atoi(vars["id"])
           if err != nil {
             errors.RespondWithError(w, http.StatusBadRequest, "Invalid bookmark id")
             return
         }
         bookmark, err := services.GetBookmarkById(db, id, user.ID)
          if err != nil {
              errors.RespondWithError(w, http.StatusNotFound, "Bookmark not found")
              return
           }
         w.Header().Set("Content-Type", "application/json")
          json.NewEncoder(w).Encode(bookmark)
      }
 }


   func PatchBookmarkById(db *sql.DB) http.HandlerFunc {
       return func(w http.ResponseWriter, r *http.Request) {
            user, ok := r.Context().Value("user").(models.User)
             if !ok{
                  errors.RespondWithError(w, http.StatusUnauthorized, "Not authorized")
                  return
              }
           vars := mux.Vars(r)
            id, err := strconv.Atoi(vars["id"])
           if err != nil {
                errors.RespondWithError(w, http.StatusBadRequest, "Invalid bookmark id")
                return
            }
            var bookmark models.Bookmark
           if err := json.NewDecoder(r.Body).Decode(&bookmark); err != nil {
              errors.RespondWithError(w, http.StatusBadRequest, "Invalid request body")
              return
          }
            err = services.PatchBookmarkById(db, id, user.ID, &bookmark)
             if err != nil {
                errors.RespondWithError(w, http.StatusInternalServerError, "Could not update bookmark")
                 return
             }
           w.Header().Set("Content-Type", "application/json")
           json.NewEncoder(w).Encode(map[string]string{"message": "Bookmark updated successfully"})
      }
   }

 func DeleteBookmarkById(db *sql.DB) http.HandlerFunc {
      return func(w http.ResponseWriter, r *http.Request) {
          user, ok := r.Context().Value("user").(models.User)
          if !ok{
             errors.RespondWithError(w, http.StatusUnauthorized, "Not authorized")
              return
          }
          vars := mux.Vars(r)
           id, err := strconv.Atoi(vars["id"])
          if err != nil {
            errors.RespondWithError(w, http.StatusBadRequest, "Invalid bookmark id")
               return
          }
          err = services.DeleteBookmarkById(db, id, user.ID)
            if err != nil {
              errors.RespondWithError(w, http.StatusInternalServerError, "Could not delete bookmark")
               return
           }
         w.Header().Set("Content-Type", "application/json")
         json.NewEncoder(w).Encode(map[string]string{"message": "Bookmark deleted successfully"})
      }
 }
 func SearchBookmarks(db *sql.DB) http.HandlerFunc {
     return func(w http.ResponseWriter, r *http.Request) {
          user, ok := r.Context().Value("user").(models.User)
            if !ok{
                 errors.RespondWithError(w, http.StatusUnauthorized, "Not authorized")
                 return
             }
         query := r.URL.Query().Get("q")
         if query == "" {
                errors.RespondWithError(w, http.StatusBadRequest, "Query is required")
             return
         }
         bookmarks, err := services.SearchBookmarks(db, query, user.ID)
          if err != nil {
              errors.RespondWithError(w, http.StatusInternalServerError, "Could not search bookmark")
              return
          }
            w.Header().Set("Content-Type", "application/json")
          json.NewEncoder(w).Encode(bookmarks)

     }
 }