package handlers

import (
 "database/sql"
  "encoding/json"
 "log"
  "net/http"
	 "github.com/Iksolot21/URL-Shortener-Bookmarker/models"
	 "github.com/Iksolot21/URL-Shortener-Bookmarker/services"
	  "github.com/Iksolot21/URL-Shortener-Bookmarker/errors"
	  "github.com/Iksolot21/URL-Shortener-Bookmarker/utils"
)


func RegisterUser(db *sql.DB) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
	 var user models.User
	 if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		  errors.RespondWithError(w, http.StatusBadRequest, "Invalid request body")
	   return
	  }
	 if err := services.RegisterUser(db, &user); err != nil {
	   errors.RespondWithError(w, http.StatusBadRequest, err.Error())
		  return
	 }

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	 json.NewEncoder(w).Encode(map[string]string{"message": "User registered successfully"})
 }
}

func LoginUser(db *sql.DB, jwtSecret string) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
	  var user models.User
	  if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		  errors.RespondWithError(w, http.StatusBadRequest, "Invalid request body")
		   return
	   }
	 
	  storedUser, err := services.LoginUser(db, user.Username, user.Password)
	  if err != nil {
		  log.Println(err)
		  errors.RespondWithError(w, http.StatusUnauthorized, "Invalid credentials")
		  return
	  }

	   token, err := utils.GenerateJWT(storedUser, jwtSecret)
	  if err != nil {
		  log.Println(err)
		   errors.RespondWithError(w, http.StatusInternalServerError, "Could not generate JWT")
		  return
	  }

	  w.Header().Set("Content-Type", "application/json")
	   json.NewEncoder(w).Encode(map[string]string{"token": token})
  }
}
func GetCurrentUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, ok := r.Context().Value("user").(models.User)
		 if !ok {
			   errors.RespondWithError(w, http.StatusUnauthorized, "Not authorized")
			 return
		  }
		 w.Header().Set("Content-Type", "application/json")
		 json.NewEncoder(w).Encode(user)
	}
}