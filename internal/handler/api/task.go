package api

type CreateTaskRequest struct {
	ID
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	DueDate     string `json:"due_date" db:"due_date" binding:"required"`
	CreatedAt   string `json:"created_at" db:"created_at" binding:"required"`
	UpdatedAt   string `json:"updated_at" db:"updated_at" binding:"required"`
}
