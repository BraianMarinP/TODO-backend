package middlewares

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/BraianMarinP/todo-backend/config"
	"github.com/BraianMarinP/todo-backend/db"
	"github.com/BraianMarinP/todo-backend/handlers"
	"github.com/BraianMarinP/todo-backend/models"
	jwtgo "github.com/dgrijalva/jwt-go"
)

// ValidateJWT is a middleware that validates the client's JSON Web Token.
func ValidateJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()

		splitToken := strings.Split(r.Header.Get("Authorization"), "Bearer")
		// Validate token format.
		if len(splitToken) != 2 {
			http.Error(w, "invalid token format", http.StatusBadRequest)
			return
		}
		userToken := splitToken[1]

		// Fetch the JWT signed string.
		jwtSignedString := []byte(config.GetEnvironmentVariable("JWT_SIGNED_STRING"))
		claims := &models.JWTClaims{}
		token, err := jwtgo.ParseWithClaims(
			userToken,
			claims,
			func(t *jwtgo.Token) (interface{}, error) {
				return jwtSignedString, nil
			},
		)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var exists bool
		exists, err = db.CheckUserExistsByUserName(ctx, claims.UserName)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if !exists {
			http.Error(w, "user doesn't exist", http.StatusNotFound)
			return
		}
		// If the user exists, set the user variables for the all endpoints.
		handlers.UserID = claims.UserID
		handlers.UserName = claims.UserName
		handlers.UserEmail = claims.Email

		// Check if the token is valid
		if !token.Valid {
			http.Error(w, "invalid token", http.StatusBadRequest)
			return
		}

		next.ServeHTTP(w, r)
	}
}

/*
 */
func CreateUser(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		next.ServeHTTP(w, r)
	}
}

/*
 */
func GetAllTasks(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		next.ServeHTTP(w, r)
	}
}

/*
 */
func CreateTask(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		next.ServeHTTP(w, r)
	}
}

/*
 */
func CompleteTask(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		next.ServeHTTP(w, r)
	}
}

/*
 */
func UndoTask(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		next.ServeHTTP(w, r)
	}
}

/*
 */
func DeleteTask(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		next.ServeHTTP(w, r)
	}
}

/*
 */
func DeleteAllTask(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		next.ServeHTTP(w, r)
	}
}

func DatabaseCheck(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		next.ServeHTTP(w, r)
	}
}
