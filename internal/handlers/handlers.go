package handlers

import (
	"taxi/configs"
	"taxi/internal/repository"
)

type HandlerStruct struct {
	Driver
	User
	Order
}

func NewHandlerStruct(app *configs.AppConfig, repo *repository.Repository) *HandlerStruct {
	return &HandlerStruct{
		Driver: NewDriver(app, repo.DriverRepository),
		User:         NewUser(app, repo.UserRepository),
		Order:        NewOrder(app, repo.OrderRepository),
	}
}
