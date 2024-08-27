package handler

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) InitRouter() *gin.Engine {
	router := gin.Default()
	tasks := router.Group("/tasks")

	tasks.POST("/", h.createTask)
	tasks.GET("/", h.getTasks)
	tasks.GET("/:id", h.getTaskByID)
	tasks.PUT("/update/:id", h.updateTask)
	tasks.DELETE("/delete/:id", h.deleteTask)
	return router
}
