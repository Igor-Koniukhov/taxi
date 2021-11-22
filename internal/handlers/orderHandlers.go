package handlers

import (
	"net/http"
	"taxi/configs"
	"taxi/internal/repository"
)

type Order interface {
	GetAllOrders(w http.ResponseWriter, r *http.Request)
	GetOrderByID(w http.ResponseWriter, r *http.Request)
	UpdateOrder(w http.ResponseWriter, r *http.Request)
	DeleteOrder(w http.ResponseWriter, r *http.Request)
}
type OrderStruct struct {
	App *configs.AppConfig
	repo repository.OrderRepository
}

func NewOrder(app *configs.AppConfig, repo repository.OrderRepository) *OrderStruct {
	return &OrderStruct{App: app, repo: repo}
}
func (o *OrderStruct) GetAllOrders(w http.ResponseWriter, r *http.Request) {

}

func (o *OrderStruct) GetOrderByID(w http.ResponseWriter, r *http.Request)  {

}

func (o*OrderStruct) UpdateOrder(w http.ResponseWriter, r *http.Request) {

}

func (o *OrderStruct) DeleteOrder(w http.ResponseWriter, r *http.Request) {

}


