package main

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"taxi/configs"
	"taxi/internal/handlers"
	"taxi/internal/repository"
)

var db *sql.DB

func Run(app *configs.AppConfig) {
	var port = "8080"
	repo := repository.NewRepository(app, db)
	www := handlers.NewHandlerStruct(app, repo)

	r := mux.NewRouter()
	r.HandleFunc("/", www.User.Home).Methods("GET")
	r.HandleFunc("/users", www.User.GetAllUsers).Methods("GET")
	r.HandleFunc("/user/{id:[0-9]+}", www.User.GetUserByID).Methods("GET")
	r.HandleFunc("/user-email/{key}", www.User.GetUserByEmail).Methods("GET")
	r.HandleFunc("/update-user/{id:[0-9]+}", www.User.UpdateUser).Methods("PUT")
	r.HandleFunc("/delete-user/{{id:[0-9]+}", www.User.DeleteUser).Methods("DELETE")

	r.HandleFunc("/orders", www.Order.GetAllOrders).Methods("GET")
	r.HandleFunc("/order/{id:[0-9]+}", www.Order.GetOrderByID).Methods("GET")
	r.HandleFunc("/order-update/", www.Order.UpdateOrder).Methods("PUT")
	r.HandleFunc("/order-delete/{id:[0-9]+}", www.Order.DeleteOrder).Methods("DELETE")

	fmt.Printf("Server strar on port %v\n", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatalln(err)
	}

}
