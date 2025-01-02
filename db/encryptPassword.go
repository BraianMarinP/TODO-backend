package db

import "golang.org/x/crypto/bcrypt"

// EncryptPassword this function encrypts a user's password.
func EncryptPassword(password string) (string, error) {
	// For normal users, it is common to use a cost of 6.
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 6)
	return string(bytes), err
}
