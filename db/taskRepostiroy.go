package db

import (
	"context"
	"fmt"

	"github.com/BraianMarinP/todo-backend/models"
)

// CreateTask this function records a new task in the database.
func CreateTask(ctx context.Context, task models.Task) (bool, error) {

	// Creates a SQL statement with the provided context
	// and query for creating a task.
	query := "INSERT INTO task (tittle, description, state, user_id) VALUES (?, ?, ?, ?)"
	preparedStatement, err := databaseConnection.PrepareContext(ctx, query)
	if err != nil {
		return false, fmt.Errorf("failed to prepare query: %w", err)
	}
	defer preparedStatement.Close()

	// Executes the prepared statement.
	_, err = preparedStatement.ExecContext(ctx, task.Title, task.Description, false, task.UserID)
	if err != nil {
		return false, fmt.Errorf("failed to execute prepared insert statement: %w", err)
	}

	return true, nil
}
