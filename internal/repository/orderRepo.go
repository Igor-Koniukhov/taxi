package repository

import (
	"database/sql"
	"taxi/configs"
	"taxi/internal/models"
)

type OrderRepository interface {
	GetAllOrders() []models.Order
	GetOrderByDriverID(driverID int) (models.Order, error)
	CreateOrder() (models.Order, error)
	UpdateOrder(driverID int) (models.Order, error)
	DeleteOrder(driverID int) error
}
type OrderRepo struct {
	App *configs.AppConfig
	DB *sql.DB
}

func NewOrderRepo(app *configs.AppConfig, db *sql.DB) *OrderRepo {
	return &OrderRepo{App: app, DB: db}
}
func (o *OrderRepo) GetAllOrders() []models.Order {
	return nil
}

func (o *OrderRepo) GetOrderByDriverID(driverID int) (models.Order, error) {
	return models.Order{}, nil
}

func (o *OrderRepo) CreateOrder() (models.Order, error) {
	return models.Order{}, nil
}

func (o *OrderRepo) UpdateOrder(driverID int) (models.Order, error) {
	return models.Order{}, nil
}

func (o *OrderRepo) DeleteOrder(driverID int) error {
	return nil
}
