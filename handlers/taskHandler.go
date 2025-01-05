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
	w.Write([]byte("Task successfully created."))
}

// DeleteTask deletes a record task from the database.
func DeleteTask(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Header", "Content-Type")

	// Fetch the task info to delete.
	var task models.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, "Error in the received data: "+err.Error(), http.StatusBadRequest)
		return
	}

	var deleted bool
	deleted, err = db.DeleteTask(ctx, task.ID, UserID)
	if err != nil {
		http.Error(w, "Error while trying to delete the task: "+err.Error(), http.StatusInternalServerError)
		return
	}
	if !deleted {
		http.Error(w, "Couldn't delete the task. Task not found.", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Task successfully deleted."))
}

// UndoTask sets the task state to incomplete.
func UndoTask(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	w.Header().Set("Content-type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Header", "Content-Type")

	// Fetch the task to undo state.
	var task models.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, "Error in the received data: "+err.Error(), http.StatusBadRequest)
		return
	}

	var updated bool
	updated, err = db.UndoTask(ctx, task.ID, UserID)
	if err != nil {
		http.Error(w, "Error while trying to undo the task: "+err.Error(), http.StatusInternalServerError)
		return
	}
	if !updated {
		http.Error(w, "Couldn't undo the task. Task not found.", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Task successfully updated."))
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

// DeleteAllTasks deletes all user task records from the database.
func DeleteAllTasks(w http.ResponseWriter, r *http.Request) {

}
