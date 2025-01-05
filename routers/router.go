package routers

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/BraianMarinP/todo-backend/config"
	"github.com/BraianMarinP/todo-backend/handlers"
	"github.com/BraianMarinP/todo-backend/middlewares"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// Router sets up the HTTP server with specific routes for handling API requests.
func Router() {
	router := mux.NewRouter()
	router.HandleFunc("/api/users/sign-up", handlers.CreateUser).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/users/login", handlers.Login).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/task/create", middlewares.ValidateJWT(handlers.CreateTask)).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/task/delete", middlewares.ValidateJWT(handlers.DeleteTask)).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/api/task/complete", middlewares.ValidateJWT(handlers.CompleteTask)).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/task/undo", middlewares.ValidateJWT(handlers.UndoTask)).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/task/deleteAllTasks", middlewares.ValidateJWT(handlers.DeleteAllTasks)).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/api/task/update", middlewares.ValidateJWT(handlers.UpdateTask)).Methods("PUT", "OPTIONS")
	/*
		router.HandleFunc("/api/getAllTasks", middlewares.GetAllTasks(handlers.GetAllTasks)).Methods("GET", "OPTIONS")
	*/
	// Grant permissions to everyone to allow access without problems
	handler := cors.AllowAll().Handler(router)
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = config.GetEnvironmentVariable("SERVER_PORT")
	}
	fmt.Println("Starting the server on port " + PORT + "...")
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
