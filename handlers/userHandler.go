package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"
	"unicode"

	"github.com/BraianMarinP/todo-backend/db"
	"github.com/BraianMarinP/todo-backend/jwt"
	"github.com/BraianMarinP/todo-backend/models"
	"github.com/asaskevich/govalidator"
)

// CreateUser records a new user in the database.
func CreateUser(w http.ResponseWriter, r *http.Request) {

	// Control the timeout of the operation.
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// Parse the user values from the request body.
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Error in the received data: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Validates the entered email address.
	if !govalidator.IsEmail(user.Email) {
		http.Error(w, "Invalid email.", http.StatusBadRequest)
		return
	}

	// Validates the entered password.
	validPassword, passErr := validatePassword(user.Password)
	if !validPassword {
		http.Error(w, passErr.Error(), http.StatusBadRequest)
		return
	}

	// Check if the user doesn't exist to allow it to be created.
	exists, err := db.CheckUserExistsByUsername(ctx, user.Username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if exists {
		http.Error(w, "There is already a user with that username.", http.StatusBadRequest)
		return
	}

	// Check if the user doesn't exist to allow it to be created.
	exists, err = db.CheckUserExistsByEmail(ctx, user.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if exists {
		http.Error(w, "There is already a user with that email.", http.StatusBadRequest)
		return
	}

	// Record the user in the database.
	ID, err := db.CreateUser(ctx, user)
	if err != nil {
		http.Error(w, "An error has ocurred while trying to register the user. "+err.Error(), http.StatusInternalServerError)
		return
	}

	// If the returned ID equals 0, an error has been ocurred while recording.
	if ID == 0 {
		http.Error(w, "The user registration has not been completed.", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// validatePassword validates the password's strength.
func validatePassword(password string) (bool, error) {

	// Validates password length.
	if len(password) < 8 {
		return false, errors.New("password must be at least 8 characters long")
	}

	// Check if the password has at least one lowercase letter, one uppercase letter,
	// one digit, and one special character.
	hasLower, hasUpper, hasDigit, hasSpecial := false, false, false, false

	// Evaluate each character of the password.
	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsDigit(char):
			hasDigit = true
		case unicode.IsSymbol(char) || unicode.IsPunct(char):
			hasSpecial = true
		}
	}

	// Verify if the password fulfills the conditions.
	switch {
	case !hasLower:
		return false, errors.New("the password needs at least one lowercase character")
	case !hasUpper:
		return false, errors.New("the password needs at least one uppercase character")
	case !hasDigit:
		return false, errors.New("the password needs at least one digit character")
	case !hasSpecial:
		return false, errors.New("the password needs at least one special character")
	}

	// The password is valid.
	return true, nil
}

// Login authenticates the user by verifying their credentials.
func Login(w http.ResponseWriter, r *http.Request) {

	// Indicate that the response body will be in JSON format.
	w.Header().Set("Context-Type", "application/json")
	// Control the timeout of the operation.
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

	// Generate the token.
	var token string
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
