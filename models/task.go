package models

// Task is the structure that represents a task.
type Task struct {
	ID     int    `json:"id"`
	Task   string `json:"task"`
	Status bool   `json:"status"`
}
