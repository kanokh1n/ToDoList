package repository

import (
	"context"
	"database/sql"
	"db-service/models"
	"errors"
	"time"
)

func (r *TaskRepository) Create(ctx context.Context, task *models.Task) error {
	task.IsCompleted = false
	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()

	query := `INSERT INTO tasks (title, description, is_completed, created_at, updated_at) 
			  VALUES ($1, $2, $3, $4, $5)`
	_, err := r.db.Exec(query, task.Title, task.Description, task.IsCompleted, task.CreatedAt, task.UpdatedAt)
	if err != nil {
		return err
	}

	r.InvalidateAllTasksCache(ctx)

	return nil
}

func (r *TaskRepository) GetAllTasks(ctx context.Context) ([]models.Task, error) {

	tasks, err := r.GetAllTasksFromCache(ctx)
	if err == nil && tasks != nil {
		return tasks, nil
	}

	query := `SELECT id, title, description, is_completed, created_at, updated_at FROM tasks`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tasks = []models.Task{}

	for rows.Next() {
		task := models.Task{}
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.IsCompleted, &task.CreatedAt, &task.UpdatedAt)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	r.SetAllTasksCache(ctx, tasks)

	return tasks, nil
}

func (r *TaskRepository) GetTaskById(ctx context.Context, taskId int) (*models.Task, error) {

	tasks, err := r.GetTaskByIdFromCache(ctx, taskId)
	if err == nil && tasks != nil {
		return tasks, nil
	}

	query := `SELECT id, title, description, is_completed, created_at, updated_at FROM tasks
			  WHERE id = $1`

	task := models.Task{}
	err = r.db.QueryRow(query, taskId).Scan(&task.ID, &task.Title, &task.Description, &task.IsCompleted, &task.CreatedAt, &task.UpdatedAt)

	if err == sql.ErrNoRows {
		return nil, errors.New("task not found")
	}

	if err != nil {
		return nil, err
	}

	r.SetTaskCache(ctx, &task)

	return &task, nil
}

func (r *TaskRepository) DeleteTask(ctx context.Context, taskId int) error {

	query := `DELETE FROM tasks WHERE id = $1`
	_, err := r.db.Exec(query, taskId)
	if err != nil {
		return err
	}

	r.InvalidateTaskCache(ctx, taskId)

	return nil
}

func (r *TaskRepository) CompleteTask(ctx context.Context, taskId int) error {
	updateTime := time.Now()
	query := `UPDATE tasks SET is_completed = true, updated_at = $1 WHERE id = $2`

	_, err := r.db.Exec(query, updateTime, taskId)
	if err != nil {
		return err
	}

	r.InvalidateTaskCaches(ctx, taskId)

	return nil
}
