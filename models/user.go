package models

// User is the struct that represents an user in the database.
type User struct {
	ID       int    `json:"id,omitempty"`
	UserName string `json:"userName"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Avatar   string `json:"avatar"`
}
