package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/BraianMarinP/todo-backend/models"
	"golang.org/x/crypto/bcrypt"
)

// CreateUser this function creates a new user record in the database.
func CreateUser(ctx context.Context, user models.User) (int, error) {
	// Ecrypts the user's password.
	var err error
	user.Password, err = EncryptPassword(user.Password)
	if err != nil {
		return 0, fmt.Errorf("failed to ecrypt password: %w", err)
	}

	// Prepares a SQL statement with the provided context
	// and query for creating a user.
	var preparedStatement *sql.Stmt
	query := "INSERT INTO user (username, email, password, avatar) VALUES (?, ?, ?, ?)"
	preparedStatement, err = databaseConnection.PrepareContext(ctx, query)
	if err != nil {
		return 0, fmt.Errorf("failed to prepare query: %w", err)
	}
	defer preparedStatement.Close()

	//Executes the prepare statement with the provided context and user details.
	var result sql.Result
	result, err = preparedStatement.ExecContext(ctx, user.UserName, user.Email, user.Password, user.Avatar)
	if err != nil {
		return 0, fmt.Errorf("failed to execute prepared insert statement: %w", err)
	}

	// Retrieves the ID of the last inserted record and return it.
	var lastID int64
	lastID, err = result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("failed to retrieve user ID: %w", err)
	}

	return int(lastID), nil
}

// CheckUserExists checks if a user name or email already exists.
func CheckUserExists(username, email string) (bool, error) {
	query := "SELECT id FROM user WHERE userName = ? OR email = ? LIMIT 1"
	var id int
	err := databaseConnection.QueryRow(query, username, email).Scan(&id)
	if err == sql.ErrNoRows {
		return false, nil //Means that user doesn't exists.
	} else if err != nil {
		return false, fmt.Errorf("error cheching user existence: %w", err)
	}
	return true, nil //User exists.
}

// getUser fetches a user from the database.
func getUser(ctx context.Context, userOrEmail string) (models.User, error) {
	// SQL query to fetch user details.
	query := "SELECT id, userName, email, password, avatar FROM user WHERE email = ? OR userName = ? "

	//Prepare the statment.
	preparedStatement, err := databaseConnection.PrepareContext(ctx, query)
	if err != nil {
		return models.User{}, fmt.Errorf("failed to prepare query: %w", err)
	}
	defer preparedStatement.Close()

	// Create an empty User struct to hold the result.
	var user models.User

	//Execute the query and scan the results into the struct.
	err = preparedStatement.QueryRowContext(ctx, userOrEmail, userOrEmail).Scan(
		&user.ID,
		&user.UserName,
		&user.Email,
		&user.Password,
		&user.Avatar,
	)
	if err != nil {
		return models.User{}, fmt.Errorf("failed to execute the query or scan result: %w", err)
	}

	// Return the populated User struct.
	return user, nil

}

// AttemptLogin tries to perform a login attempt.
func AttemptLogin(ctx context.Context, userOrEmail string, password string) (models.User, bool, error) {

	user, err := getUser(ctx, userOrEmail)
	//	Checks if an error has been ocurred.
	if err != nil {
		return models.User{}, false, err
	}

	// Compare the passwords.
	passwordBytes := []byte(password)
	fecthUserPassword := []byte(user.Password)
	err = bcrypt.CompareHashAndPassword(fecthUserPassword, passwordBytes)
	if err != nil {
		return models.User{}, false, nil //The passwords are not the same.
	}

	return user, true, nil // Successful login.
}
