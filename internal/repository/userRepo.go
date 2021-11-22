package repository

import (
	"database/sql"
	"taxi/configs"
	"taxi/internal/models"
)

type UserRepository interface {
	GetAllUsers()[]models.User
	GetUserByEmail(email string) (models.User, error)
	GetUserByID(id int)(models.User, error)	
}
type UserRepo struct {
	App *configs.AppConfig
	DB *sql.DB
}

func NewUserRepo(app *configs.AppConfig, db *sql.DB) *UserRepo {
	return &UserRepo{App: app, DB: db}
}

func (u *UserRepo) GetAllUsers() []models.User {
	return nil
}

func (u *UserRepo) GetUserByEmail(email string) (models.User, error) {
	return models.User{}, nil
}

func (u *UserRepo) GetUserByID(id int) (models.User, error) {
	return models.User{}, nil
}


