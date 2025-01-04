package jwt

import (
	"time"

	"github.com/BraianMarinP/todo-backend/config"
	"github.com/BraianMarinP/todo-backend/models"
	jwt "github.com/dgrijalva/jwt-go"
)

// GeneratesJsonWebToken generates a JWT using the user's info.
func GeneratesJsonWebToken(user models.User) (string, error) {
	// Private password.
	privatePassword := []byte(config.GetEnvironmentVariable("JWT_SIGNED_STRING"))
	/*
		Claims are pieces of information about an entity that
		are included in the token.
	*/
	payload := jwt.MapClaims{
		"id":       user.ID,
		"userName": user.UserName,
		"email":    user.Email,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // it returns the data in a long format
	}

	// Generate the token.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenString, err := token.SignedString(privatePassword)
	if err != nil {
		return tokenString, err
	}
	return tokenString, err
}
