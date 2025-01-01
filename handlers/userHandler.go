package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/BraianMarinP/todo-backend/db"
	"github.com/BraianMarinP/todo-backend/jwt"
	"github.com/BraianMarinP/todo-backend/models"
)

// CreateUser records a new user in the database.
func CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Error in the received data: "+err.Error(), http.StatusBadRequest)
		return
	}

	if len(user.Email) == 0 {
		http.Error(w, "The email is required.", 400)
		return
	}

	if len(user.Password) < 6 {
		http.Error(w, "The password length needs to be at least 6 characters.", 400)
		return
	}

	finded, _ := db.CheckUserExists(user.UserName, user.Email)
	if finded {
		http.Error(w, "There is already a user with that email.", 400)
		return
	}

	ID, err := db.CreateUser(ctx, user)
	if err != nil {
		http.Error(w, "An error has ocurred while trying to register the user. "+err.Error(), 400)
		return
	}

	if ID == 0 {
		http.Error(w, "The user registration has not been completed.", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// Login authenticates the user by verifying their credentials.
func Login(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Context-Type", "application/json")
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	var userCredentials models.Credentials

	// Fetch the credentials.
	err := json.NewDecoder(r.Body).Decode(&userCredentials)
	if err != nil {
		http.Error(w, "Invalid user or password.", http.StatusUnauthorized)
		return
	}

	var user models.User
	var authenticated bool
	// Perform a login attempt.
	user, authenticated, err = db.AttemptLogin(ctx, userCredentials.User, userCredentials.Password)
	if err != nil {
		// If authenticated is false, the password is incorrect.
		if !authenticated {
			http.Error(w, "Incorrect password.", http.StatusUnauthorized)
			return
		} else {
			// If the user is authenticated, it was an internal server error.
			http.Error(w, "An error has ocurred: "+err.Error(), http.StatusInternalServerError)
			return
		}
	}

	var token string
	// Generate the token.
	token, err = jwt.GeneratesJsonWebToken(user)
	if err != nil {
		http.Error(w, "Error generating the JSON Web Token: "+err.Error(), http.StatusInternalServerError)
		return
	}

	responseToken := models.LoginResponse{
		Token: token,
	}

	// Sets the HTTP status code for the response.
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(responseToken) // Return the response token to the client.
}
