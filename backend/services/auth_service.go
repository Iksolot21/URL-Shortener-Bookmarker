package services

import (
	"database/sql"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
	"github.com/Iksolot21/URL-Shortener-Bookmarker/backend/models"
)

// RegisterUser registers a new user
func RegisterUser(db *sql.DB, user *models.User) error {
	// 1. Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}
	user.Password = string(hashedPassword) // Store the hashed password

	// 2. Insert the user into the database
	query := `INSERT INTO users (username, password, email, created_at) VALUES ($1, $2, $3, $4)`
	_, err = db.Exec(query, user.Username, user.Password, user.Email, time.Now())
	if err != nil {
		return fmt.Errorf("failed to insert user into database: %w", err)
	}

	return nil
}

// LoginUser logs in an existing user
func LoginUser(db *sql.DB, username, password string) (models.User, error) {
	var user models.User

	// 1. Fetch the user from the database by username
	query := `SELECT id, username, password, email FROM users WHERE username = $1`
	row := db.QueryRow(query, username)
	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Email)
	if err != nil {
		return user, fmt.Errorf("user not found: %w", err)
	}

	// 2. Compare the stored hash password with the password from request
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, fmt.Errorf("invalid password: %w", err)
	}

	return user, nil
}