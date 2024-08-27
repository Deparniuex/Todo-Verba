package handler

import "Todo-Verba/internal/service"

type Handler struct {
	Service service.Service
}

func New(service service.Service) *Handler {
	return &Handler{
		Service: service,
	}
}
