package middlewares

import (
	"net/http"
)

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
