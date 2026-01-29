package repository

import (
	"context"
	"db-service/models"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

const (
	tasksAllKey   = "tasks:all"
	taskKeyPrefix = "task:"
	cacheTTL      = 10 * time.Minute
)

func (r *TaskRepository) GetAllTasksFromCache(ctx context.Context) ([]models.Task, error) {
	data, err := r.redis.Get(ctx, tasksAllKey).Result()
	if err == redis.Nil {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	var tasks []models.Task
	if err := json.Unmarshal([]byte(data), &tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *TaskRepository) SetAllTasksCache(ctx context.Context, tasks []models.Task) error {
	data, err := json.Marshal(tasks)
	if err != nil {
		return err
	}

	return r.redis.Set(ctx, tasksAllKey, data, cacheTTL).Err()
}

func (r *TaskRepository) GetTaskByIdFromCache(ctx context.Context, taskId int) (*models.Task, error) {
	key := taskKey(taskId)

	data, err := r.redis.Get(ctx, key).Result()
	if err == redis.Nil {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	var task models.Task
	if err := json.Unmarshal([]byte(data), &task); err != nil {
		return nil, err
	}

	return &task, nil
}

func (r *TaskRepository) SetTaskCache(ctx context.Context, task *models.Task) error {
	data, err := json.Marshal(task)
	if err != nil {
		return err
	}

	key := taskKey(task.ID)
	return r.redis.Set(ctx, key, data, cacheTTL).Err()
}

func (r *TaskRepository) InvalidateAllTasksCache(ctx context.Context) error {
	return r.redis.Del(ctx, tasksAllKey).Err()
}

func (r *TaskRepository) InvalidateTaskCache(ctx context.Context, taskId int) error {
	key := taskKey(taskId)
	return r.redis.Del(ctx, key).Err()
}

func (r *TaskRepository) InvalidateTaskCaches(ctx context.Context, taskId int) error {
	key := taskKey(taskId)
	return r.redis.Del(ctx, key, tasksAllKey).Err()
}

func taskKey(taskId int) string {
	return fmt.Sprintf("%s%d", taskKeyPrefix, taskId)
}
