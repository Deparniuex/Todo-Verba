package pgrepo

import (
	"Todo-Verba/internal/entity"
	"database/sql"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
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
	query := fmt.Sprintf(`
	SELECT * FROM %s
	ORDER BY id
`, tasksTable)
	rows, err := p.DB.Query(query)
	if err != nil {
		return nil, err
	}
	var tasks []*entity.Task
	for rows.Next() {
		var task entity.Task
		rows.Scan(&task.ID, &task.Title, &task.Description, &task.DueDate, &task.CreatedAt, &task.UpdatedAt)
		tasks = append(tasks, &task)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return tasks, nil
}

func (p Postgres) GetTaskById(taskID int64) (*entity.Task, error) {
	query := fmt.Sprintf(`
	SELECT * FROM %s WHERE id = $1;
`, tasksTable)
	logrus.Debugf("taskID: %d", taskID)
	task := &entity.Task{}
	err := p.DB.QueryRow(query, taskID).Scan(&task.ID, &task.Title, &task.Description, &task.DueDate, &task.CreatedAt, &task.UpdatedAt)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}
	return task, nil
}

func (p Postgres) UpdateTask(task *entity.Task) error {
	query := fmt.Sprintf(`
	UPDATE %s 
	SET 
	    title = $1,
		description = $2,
		due_date = $3,
		created_at = $4,
		updated_at = $5
	WHERE id = $6
`, tasksTable)
	tag, err := p.DB.Exec(query, task.Title, task.Description, task.DueDate, task.CreatedAt, task.UpdatedAt, task.ID)
	if err != nil {
		return err
	}

	rowsSum, _ := tag.RowsAffected()
	if rowsSum == 0 {
		return ErrRecordNotFound
	}
	return nil
}

func (p Postgres) DeleteTask(taskID int64) error {
	query := fmt.Sprintf(`
		DELETE FROM %s
		WHERE id = $1;
`, tasksTable)
	tag, err := p.DB.Exec(query, taskID)
	if err != nil {
		return err
	}
	rowsSum, _ := tag.RowsAffected()
	if rowsSum == 0 {
		return ErrRecordNotFound
	}
	return nil
}
