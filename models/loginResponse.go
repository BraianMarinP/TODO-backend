package models

// LoginResponse contain the token returned during the login.
type LoginResponse struct {
	Token string `json:"token"`
}
