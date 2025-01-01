package models

// Credentials contain the user's credentials to perform a login attempt.
type Credentials struct {
	User     string `json:"user"`
	Password string `json:"password"`
}
