package repository

import (
	"database/sql"
	"taxi/configs"
)

type Repository struct {	
	DriverRepository
	OrderRepository
	UserRepository
}

func NewRepository(app *configs.AppConfig, db *sql.DB) *Repository {
	return &Repository{
		DriverRepository: NewDriverRepo(app, db),
		OrderRepository:  NewOrderRepo(app, db),
		UserRepository:   NewUserRepo(app, db),
	}
}

