package service

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func (s *UserService) CreateToken(ID int) (string, error) {
	// Validate the ID
	if ID <= 0 {
		return "", errors.New("ID cannot be negative or zero")
	}
	// Check if JWT_SECRET is set
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return "", errors.New("JWT_SECRET environment variable is not set")
	}

	// Create the JWT claims, including the ID and expiration time
	claims := jwt.MapClaims{
		"id":  ID,
		"exp": time.Now().AddDate(0, 1, 0).Unix(), // Adds 1 month to the current time
	}

	// Create the JWT token with the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", errors.New("failed to sign the token")
	}

	return tokenString, nil
}
