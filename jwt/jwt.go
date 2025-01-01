package jwt

import (
	"time"

	"github.com/BraianMarinP/todo-backend/models"
	jwt "github.com/dgrijalva/jwt-go"
)

func GeneratesJsonWebToken(user models.User) (string, error) {
	// Private password.
	privatePassword := []byte("ThisIsAnExampleOf_A_PrivatePasssWordToCreateTheJsonWebToken")
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
