package service

import "Todo-Verba/internal/entity"

func (m Manager) CreateTask(task *entity.Task) error {
	return m.Repository.CreateTask(task)
}

func (m Manager) GetTasks() ([]*entity.Task, error) {
	//TODO implement me
	panic("implement me")
}

func (m Manager) GetTaskById(taskID int64) (*entity.Task, error) {
	//TODO implement me
	panic("implement me")
}

func (m Manager) UpdateTask(task *entity.Task) error {
	//TODO implement me
	panic("implement me")
}

func (m Manager) DeleteTask(taskID int64) error {
	//TODO implement me
	panic("implement me")
}
