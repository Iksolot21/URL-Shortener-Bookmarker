package jwt
  import (
  	"fmt"
     "time"
       "github.com/golang-jwt/jwt/v5"
     "github.com/Iksolot21/URL-Shortener-Bookmarker/backend/models"
   )
  func GenerateJWT(user models.User, jwtSecret string) (string, error) {
  	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
  		"userId": user.ID,
        "username": user.Username,
         "email": user.Email,
  		"exp": time.Now().Add(time.Hour * 24).Unix(),
  	})

  	tokenString, err := token.SignedString([]byte(jwtSecret))
   if err != nil {
     return "", fmt.Errorf("failed to generate token: %w", err)
    }
   return tokenString, nil
  }
