package handler

import (
	"Todo-Verba/internal/entity"
	"Todo-Verba/internal/handler/api"
	"Todo-Verba/internal/repository/pgrepo"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func (h *Handler) createTask(ctx *gin.Context) {
	var req api.CreateTaskRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &api.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}
	err = h.Service.CreateTask(&entity.Task{
		ID:          req.Value,
		Title:       req.Title,
		Description: req.Description,
		DueDate:     req.DueDate,
		CreatedAt:   req.CreatedAt,
		UpdatedAt:   req.UpdatedAt,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &api.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, &api.DefaultResponse{
		Code:    http.StatusCreated,
		Message: "ok",
	})
}

func (h *Handler) getTaskByID(ctx *gin.Context) {
	var req api.ID
	err := ctx.ShouldBindUri(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &api.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}
	logrus.Debugf("req ID: %d", req.Value)
	task, err := h.Service.GetTaskById(req.Value)
	if err != nil {
		switch {
		case errors.Is(err, pgrepo.ErrRecordNotFound):
			{
				ctx.JSON(http.StatusNotFound, &api.ErrorResponse{
					Code:    http.StatusNotFound,
					Message: err.Error(),
				})
				return
			}
		default:
			ctx.JSON(http.StatusInternalServerError, &api.ErrorResponse{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			})
			return
		}
	}
	ctx.JSON(http.StatusOK, &api.GetTaskByIDResponse{
		Code:    http.StatusOK,
		Message: "ok",
		Body:    task,
	})
}

func (h *Handler) getTasks(ctx *gin.Context) {

}

func (h *Handler) updateTask(ctx *gin.Context) {

}

func (h *Handler) deleteTask(ctx *gin.Context) {

}
