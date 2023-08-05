package helper

import (
	"freelance/clinic_queue/models"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Function to generate a JWT token
func GenerateToken(username string) (string, error) {
	// Create the claims for the token
	claims := &models.Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(), // Token expires in 48 hours
			IssuedAt:  time.Now().Unix(),
		},
	}

	// Create the token with the claims and sign it with a secret key
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("you can not found me bro ))!..?"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}