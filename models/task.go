package models

// Task is the structure that represents a task.
type Task struct {
	ID          int    `json:"id,omitempty"`
	Title       string `json:"tittle"`
	Description string `json:"description"`
	State       bool   `json:"state"`
	UserID      int    `json:"userID"`
}
