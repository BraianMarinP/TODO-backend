package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/BraianMarinP/todo-backend/db"
	"github.com/BraianMarinP/todo-backend/models"
)

/*
 */
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

	finded, _ := db.CheckUserExistss(user.UserName, user.Email)
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
