package db

import (
	"context"
	"database/sql"
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

// DeleteTask deletes a user's task using task and user ID.
func DeleteTask(ctx context.Context, taskID int, userID int) (bool, error) {
	// Creates a SQL statement with the provided context
	// and query for deleting a task.
	query := "DELETE FROM task WHERE id = ? and user_id = ?"
	preparedStatement, err := databaseConnection.PrepareContext(ctx, query)
	if err != nil {
		return false, fmt.Errorf("failed to prepare query: %w", err)
	}
	defer preparedStatement.Close()

	// Executes the prepared statement.
	var result sql.Result
	result, err = preparedStatement.ExecContext(ctx, taskID, userID)
	if err != nil {
		return false, fmt.Errorf("failed to execute prepared delete statement: %w", err)
	}

	// Check if any rows were affected.
	var rowsAffected int64
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		return false, fmt.Errorf("error while fetching rows affected: %w", err)
	}

	// No rows were deleted.
	if rowsAffected == 0 {
		return false, nil
	}

	return true, nil
}

// UndoTask updates a task state to incomplete.
func UndoTask(ctx context.Context, taskID int, userID int) (bool, error) {

	// Create a SQL statement with the provided context
	// and query for undoing a task.
	query := "UPDATE task SET state = false WHERE id = ? and user_id = ?"
	preparedStatement, err := databaseConnection.PrepareContext(ctx, query)
	if err != nil {
		return false, fmt.Errorf("failed to prepare query: %w", err)
	}

	// Execute the prepared statement.
	var result sql.Result
	result, err = preparedStatement.ExecContext(ctx, taskID, userID)
	if err != nil {
		return false, fmt.Errorf("failed to execute prepared delete statement: %w", err)
	}

	return detectRowsAffected(result)
}

// UndoTask updates a task state to completed.
func CompleteTask(ctx context.Context, taskID int, userID int) (bool, error) {

	// Create a SQL statement with the provided context
	// and query for undoing a task.
	query := "UPDATE task SET state = true WHERE id = ? and user_id = ?"
	preparedStatement, err := databaseConnection.PrepareContext(ctx, query)
	if err != nil {
		return false, fmt.Errorf("failed to prepare query: %w", err)
	}

	// Execute the prepared statement.
	var result sql.Result
	result, err = preparedStatement.ExecContext(ctx, taskID, userID)
	if err != nil {
		return false, fmt.Errorf("failed to execute prepared delete statement: %w", err)
	}

	return detectRowsAffected(result)
}

// detectRowsAffected detects if the updates were successfully performed.
func detectRowsAffected(result sql.Result) (bool, error) {
	// Check if any rows were affected.
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return false, fmt.Errorf("error while fetching rows affected: %w", err)
	}

	// No rows were deleted.
	if rowsAffected == 0 {
		return false, nil
	}

	return true, nil
}
