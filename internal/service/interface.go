package service

import "Todo-Verba/internal/entity"

type Service interface {
	CreateTask(task *entity.Task) error
	GetTasks() ([]*entity.Task, error)
	GetTaskById(taskID int64) (*entity.Task, error)
	UpdateTask(task *entity.Task) error
	DeleteTask(taskID int64) error
}
