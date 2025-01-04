package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/BraianMarinP/todo-backend/db"
	"github.com/BraianMarinP/todo-backend/models"
)

// CreateTask this function records a new task.
func CreateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Header", "Content-Type")

	// Control the timeout of the operation.
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// Parse the task values from the request body.
	var task models.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, "Error in the received data: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Get the user id.
	// userID := mux.Vars(r)["user_id"]
	// task.UserID, err = strconv.Atoi(userID)

	// if err != nil {
	// 	http.Error(w, "Error fetching user id: "+err.Error(), http.StatusBadRequest)
	// 	return
	// }
	task.UserID = UserID
	task.State = false

	// Record the task in the database.
	created, err := db.CreateTask(ctx, task)
	if err != nil {
		http.Error(w, "Error while recording the task: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Checks if the task was created successfully.
	if !created {
		http.Error(w, "The task couldn't be created.", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusCreated)
}

// GetAllTasks this function retrieves all tasks of a user from the database
func GetAllTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// payload := getAllTasks()
	// json.NewEncoder(w).Encode(payload)
}

// CompleteTask sets the state of a task to completed.
func CompleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Header", "Content-Type")
}

// UndoTask sets the task state to incomplete.
func UndoTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Header", "Content-Type")

	//taskID := r.PathValue("id")

}

// DeleteTask deletes a record task from the database.
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Header", "Content-Type")
}

// DeleteAllTasks deletes all user task records from the database.
func DeleteAllTasks(w http.ResponseWriter, r *http.Request) {

}
