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

	// Executes the prepare statement with the provided context and user details.
	var result sql.Result
	result, err = preparedStatement.ExecContext(ctx, user.Username, user.Email, user.Password, user.Avatar)
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

// CheckUserExistsByUsername checks if a user name already exists.
func CheckUserExistsByUsername(ctx context.Context, username string) (bool, error) {
	// SQL query to check if a user exists using a user name.
	query := "SELECT id FROM user WHERE username = ? LIMIT 1"
	return findID(ctx, username, query)
}

// CheckUserExistsByEmail checks if a email already exists.
func CheckUserExistsByEmail(ctx context.Context, email string) (bool, error) {
	// SQL query to check if a user exists using a user name.
	query := "SELECT id FROM user WHERE email = ? LIMIT 1"
	return findID(ctx, email, query)
}

// findID fetches a user's ID using their email or username.
func findID(ctx context.Context, usernameOrEmail string, query string) (bool, error) {
	// Prepare the statment.
	preparedStatement, err := databaseConnection.PrepareContext(ctx, query)
	if err != nil {
		return false, fmt.Errorf("failed to prepare query: %w", err)
	}
	defer preparedStatement.Close()

	// Execute the query and scan the user id.
	var id int
	err = preparedStatement.QueryRowContext(ctx, usernameOrEmail).Scan(&id)

	if err == sql.ErrNoRows {
		return false, nil //Means that user doesn't exists.
	} else if err != nil {
		return false, fmt.Errorf("error checking user existence: %w", err)
	}
	return true, nil //User exists.
}

// getUser fetches a user from the database.
func getUser(ctx context.Context, userOrEmail string) (models.User, error) {
	// SQL query to fetch user details.
	query := "SELECT id, userName, email, password, avatar FROM user WHERE email = ? OR userName = ? "

	// Prepare the statment.
	preparedStatement, err := databaseConnection.PrepareContext(ctx, query)
	if err != nil {
		return models.User{}, fmt.Errorf("failed to prepare query: %w", err)
	}
	defer preparedStatement.Close()

	// Create an empty User struct to hold the result.
	var user models.User

	// Execute the query and scan the results into the struct.
	err = preparedStatement.QueryRowContext(ctx, userOrEmail, userOrEmail).Scan(
		&user.ID,
		&user.Username,
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
