package handler

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) InitRouter() *gin.Engine {
	router := gin.Default()
	return router
}
