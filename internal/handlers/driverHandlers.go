package handlers

import (
	"net/http"
	"taxi/configs"
	"taxi/internal/repository"
)

type Driver interface {
	GetAllDrivers(w http.Response, r *http.Request)
	GetDriver(w http.Response, r *http.Request)
	RemoveDriver(w http.Response, r *http.Request)
	UpdateDriver(w http.Response, r *http.Request)
}
type DriverStruct struct {
	App *configs.AppConfig
	repo repository.DriverRepository
}
func NewDriver(app *configs.AppConfig, repo repository.DriverRepository) *DriverStruct {
	return &DriverStruct{App: app, repo: repo}
}
func (d DriverStruct) GetAllDrivers(w http.Response, r *http.Request) {

}

func (d DriverStruct) GetDriver(w http.Response, r *http.Request) {

}

func (d DriverStruct) RemoveDriver(w http.Response, r *http.Request) {

}

func (d DriverStruct) UpdateDriver(w http.Response, r *http.Request) {

}





