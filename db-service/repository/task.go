package repository

import (
	"database/sql"

	"github.com/redis/go-redis/v9"
)

type TaskRepository struct {
	db    *sql.DB
	redis *redis.Client
}

func NewTaskRepository(db *sql.DB, redis *redis.Client) *TaskRepository {
	return &TaskRepository{
		db:    db,
		redis: redis,
	}
}
