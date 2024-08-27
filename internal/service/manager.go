package service

import "Todo-Verba/internal/repository"

type Manager struct {
	Repository repository.Repository
}

func New(repository repository.Repository) *Manager {
	return &Manager{Repository: repository}
}
