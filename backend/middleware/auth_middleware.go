package middleware
import (
    "context"
    "database/sql"
    "log"
    "net/http"
    "strings"
   "github.com/Iksolot21/URL-Shortener-Bookmarker/backend/models"
    "github.com/Iksolot21/URL-Shortener-Bookmarker/backend/utils"
    "github.com/Iksolot21/URL-Shortener-Bookmarker/backend/internal/errors"
)

func AuthMiddleware(next http.HandlerFunc, jwtSecret string) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
         authHeader := r.Header.Get("Authorization")
        if authHeader == "" {
            errors.RespondWithError(w, http.StatusUnauthorized, "Authorization header is required")
            return
        }

         parts := strings.Split(authHeader, " ")
        if len(parts) != 2 || parts[0] != "Bearer" {
            errors.RespondWithError(w, http.StatusUnauthorized, "Invalid token format")
             return
        }
        token := parts[1]

       user, err := utils.ValidateJWT(token, jwtSecret)
       if err != nil {
          log.Println(err)
           errors.RespondWithError(w, http.StatusUnauthorized, "Invalid token")
             return
         }

         ctx := context.WithValue(r.Context(), "user", user)
        next.ServeHTTP(w, r.WithContext(ctx))


    }
}