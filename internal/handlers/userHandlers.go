package handlers

import (
	"fmt"
	"net/http"
	"taxi/configs"
	"taxi/internal/repository"
)

type User interface {
	Home(w http.ResponseWriter, r *http.Request)
	GetAllUsers(w http.ResponseWriter, r *http.Request)
	GetUserByID(w http.ResponseWriter, r *http.Request)
	GetUserByEmail(w http.ResponseWriter, r *http.Request)
	UpdateUser(w http.ResponseWriter, r *http.Request)
	DeleteUser(w http.ResponseWriter, r *http.Request)
}
type UserStruct struct {
	App  *configs.AppConfig
	repo repository.UserRepository
}

func (u *UserStruct) Home(w http.ResponseWriter, r *http.Request) {
	_, err :=fmt.Fprintf(w, "Hello! This is home page!")
	fmt.Println(err)
}
func NewUser(app *configs.AppConfig, repo repository.UserRepository) *UserStruct {
	return &UserStruct{App: app, repo: repo}
}

func (u UserStruct) GetAllUsers(w http.ResponseWriter, r *http.Request) {

}

func (u UserStruct) GetUserByID(w http.ResponseWriter, r *http.Request) {

}

func (u UserStruct) GetUserByEmail(w http.ResponseWriter, r *http.Request) {

}

func (u UserStruct) UpdateUser(w http.ResponseWriter, r *http.Request) {

}

func (u UserStruct) DeleteUser(w http.ResponseWriter, r *http.Request) {

}
