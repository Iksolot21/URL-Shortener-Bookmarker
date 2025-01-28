package utils

  import (
    "fmt"
    "time"
    "github.com/golang-jwt/jwt/v5"
      "github.com/Iksolot21/URL-Shortener-Bookmarker/backend/models"
    "errors"
    "log"
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
    log.Println(err)
    return "", fmt.Errorf("failed to generate token: %w", err)
    }

   return tokenString, nil
  }
  func ValidateJWT(tokenString, jwtSecret string) (models.User, error) {
     var user models.User
  	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
         if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
        return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
      }

      return []byte(jwtSecret), nil
  	})

     if err != nil {
         return user, fmt.Errorf("failed to parse token: %w", err)
     }
     if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
         id, ok := claims["userId"].(float64)
         if !ok {
           return user, errors.New("invalid token claims")
       }
           username, ok := claims["username"].(string)
        if !ok {
         return user, errors.New("invalid token claims")
         }
          email, ok := claims["email"].(string)
       if !ok {
            return user, errors.New("invalid token claims")
         }
    user.ID = int(id)
    user.Username = username
     user.Email = email
        return user, nil
      }

  	return user, errors.New("invalid token")
  }