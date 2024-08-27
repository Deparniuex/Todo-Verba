package handler

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) InitRouter() *gin.Engine {
	router := gin.Default()
	tasks := router.Group("/tasks")

	tasks.POST("/")
	tasks.GET("/")
	tasks.GET("/:id")
	tasks.PUT("/update/:id")
	tasks.DELETE("/delete/:id")
	return router
}
