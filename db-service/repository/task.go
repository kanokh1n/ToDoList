package repository

import (
	"database/sql"
	"db-service/models"
	"errors"
	"time"
)

type TaskRepository struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (r *TaskRepository) Create(task *models.Task) error {

	task.IsCompleted = false
	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()

	query := `INSERT INTO tasks (title, description, is_completed, created_at, updated_at) 
			  VALUES ($1, $2, $3, $4, $5)`
	_, err := r.db.Exec(query, task.Title, task.Description, task.IsCompleted, task.CreatedAt, task.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (r *TaskRepository) GetAllTasks() ([]models.Task, error) {
	query := `SELECT id, title, description, is_completed, created_at, updated_at FROM tasks`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tasks := []models.Task{}

	for rows.Next() {
		task := models.Task{}
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.IsCompleted, &task.CreatedAt, &task.UpdatedAt)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (r *TaskRepository) GetTaskByTitle(title string) (*models.Task, error) {
	query := `SELECT id, title, description, is_completed, created_at, updated_at FROM tasks
			  WHERE title = $1`

	task := models.Task{}
	err := r.db.QueryRow(query, title).Scan(&task.ID, &task.Title, &task.Description, &task.IsCompleted, &task.CreatedAt, &task.UpdatedAt)

	if err == sql.ErrNoRows {
		return nil, errors.New("task not found")
	}

	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (r *TaskRepository) DeleteTask(title string) error {
	query := `DELETE FROM tasks WHERE title = $1`
	_, err := r.db.Exec(query, title)
	if err != nil {
		return err
	}
	return nil
}

func (r *TaskRepository) CompleteTask(title string) error {
	updateTime := time.Now()
	query := `UPDATE tasks SET is_completed = true, updated_at = $1 WHERE title = $2`

	_, err := r.db.Exec(query, updateTime, title)
	if err != nil {
		return err
	}
	return nil
}
