package repository

import (
	"database/sql"
	"taxi/configs"
	"taxi/internal/models"
)

type DriverRepository interface {
	GetAllDrivers()([]models.Driver, error)
	GetDriver(id int) models.Driver
	RemoveDriver(id int) error
	UpdateDriver(id int) error
}
type DriverRepo struct {
	App *configs.AppConfig
	DB *sql.DB
}

func NewDriverRepo(app *configs.AppConfig, db *sql.DB) *DriverRepo {
	return &DriverRepo{App: app, DB: db}
}

func (d *DriverRepo) GetAllDrivers() ([]models.Driver, error) {
	return nil, nil
}

func (d *DriverRepo) GetDriver(id int) models.Driver {
	return models.Driver{}
}

func (d *DriverRepo) RemoveDriver(id int) error {
	return nil
}

func (d *DriverRepo) UpdateDriver(id int) error {
	return nil
}




