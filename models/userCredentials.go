package models

// UserCredentials contain the user's credentials to perform a login attempt.
type UserCredentials struct {
	User     string `json:"user"`
	Password string `json:"password"`
}
