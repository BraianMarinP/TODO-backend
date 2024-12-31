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

func Router() {
	router := mux.NewRouter()
	router.HandleFunc("/api/users/sign-up", middlewares.CreateUser(handlers.CreateUser)).Methods("POST", "OPTIONS")
	/*
		router.HandleFunc("/api/createTask", middlewares.CreateTask(handlers.CreateTask)).Methods("POST", "OPTIONS")
		router.HandleFunc("/api/getAllTasks", middlewares.GetAllTasks(handlers.GetAllTasks)).Methods("GET", "OPTIONS")
		router.HandleFunc("/api/completeTask", middlewares.CompleteTask(handlers.CompleteTask)).Methods("PUT", "OPTIONS")
		router.HandleFunc("/api/undoTask", middlewares.UndoTask(handlers.UndoTask)).Methods("PUT", "OPTIONS")
		router.HandleFunc("/api/deleteTask", middlewares.DeleteTask(handlers.DeleteTask)).Methods("DELETE", "OPTIONS")
		router.HandleFunc("/api/deleteAllTask", middlewares.DeleteAllTask(handlers.DeleteAllTask)).Methods("DELETE", "OPTIONS")
	*/
	// Grant permissions to everyone to allow access without problems
	handler := cors.AllowAll().Handler(router)
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = config.GetEnviromentVariable("SERVER_PORT")
	}
	fmt.Println("Starting the server on port " + PORT + "...")
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
