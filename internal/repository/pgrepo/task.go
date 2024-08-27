package pgrepo

import "Todo-Verba/internal/entity"

func (p Postgres) CreateTask(task *entity.Task) error {
	//TODO implement me
	panic("implement me")
}

func (p Postgres) GetTasks() ([]*entity.Task, error) {
	//TODO implement me
	panic("implement me")
}

func (p Postgres) GetTaskById(taskID int64) (*entity.Task, error) {
	//TODO implement me
	panic("implement me")
}

func (p Postgres) UpdateTask(task *entity.Task) error {
	//TODO implement me
	panic("implement me")
}

func (p Postgres) DeleteTask(taskID int64) error {
	//TODO implement me
	panic("implement me")
}
