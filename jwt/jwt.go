package jwt

import (
	"github.com/BraianMarinP/todo-backend/models"
	jwt "github.com/dgrijalva/jwt-go"
)

func GeneratesJsonWebToken(user models.User) (string, error) {
	// Private password
	privatePassword := []byte("ThisIsAnExampleOf_A_PrivatePasssWordToCreateTheJsonWebToken")
	/*
		Claims are pieces of information about an entity that
		are included in the token.
	*/
	payload := jwt.MapClaims{
		"id":       user.ID,
		"userName": user.UserName,
		"email":    user.Email,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenString, err := token.SignedString(privatePassword)
	if err != nil {
		return tokenString, err
	}
	return tokenString, err
}
