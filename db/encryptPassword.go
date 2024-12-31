package db

import "golang.org/x/crypto/bcrypt"

/*
EncryptPassword this function encrypts a user's password
*/
func EncryptPassword(password string) (string, error) {
	//Normal users 6
	cost := 6
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	return string(bytes), err
}
