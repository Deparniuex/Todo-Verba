package service

import "Todo-Verba/internal/entity"

func (m Manager) CreateTask(task *entity.Task) error {
	return m.Repository.CreateTask(task)
}

func (m Manager) GetTasks() ([]*entity.Task, error) {
	return m.Repository.GetTasks()
}

func (m Manager) GetTaskById(taskID int64) (*entity.Task, error) {
	return m.Repository.GetTaskById(taskID)
}

func (m Manager) UpdateTask(task *entity.Task) error {
	return m.Repository.UpdateTask(task)
}

func (m Manager) DeleteTask(taskID int64) error {
	return m.Repository.DeleteTask(taskID)
}
