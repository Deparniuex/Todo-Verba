package pgrepo

import (
	"Todo-Verba/internal/entity"
	"fmt"
)

func (p Postgres) CreateTask(task *entity.Task) error {
	query := fmt.Sprintf(`
		INSERT INTO %s (
		                title,
		                description,
		                due_date,
		                created_at,
		                updated_at
		)
		VALUES ($1, $2, $3, $4, $5)
`, tasksTable)
	_, err := p.DB.Exec(query, task.Title, task.Description,
		task.DueDate, task.CreatedAt,
		task.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
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
