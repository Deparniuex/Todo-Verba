package api

import "Todo-Verba/internal/entity"

type CreateTaskRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	DueDate     string `json:"due_date" db:"due_date" binding:"required"`
	CreatedAt   string `json:"created_at" db:"created_at" binding:"required"`
	UpdatedAt   string `json:"updated_at" db:"updated_at" binding:"required"`
}

type GetTaskByIDResponse struct {
	Code    int          `json:"code"`
	Message string       `json:"message"`
	Body    *entity.Task `json:"body"`
}

type GetTasksResponse struct {
	Code    int            `json:"code"`
	Message string         `json:"message"`
	Body    []*entity.Task `json:"body"`
}

type UpdateTaskRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	DueDate     string `json:"due_date" db:"due_date" binding:"required"`
	CreatedAt   string `json:"created_at" db:"created_at" binding:"required"`
	UpdatedAt   string `json:"updated_at" db:"updated_at" binding:"required"`
}
