package router

import (
	"github.com/DracFiendMG/go-microservice/handler"
	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/users", handler.GetUsers).Methods("GET")
	r.HandleFunc("/api/users", handler.CreateUser).Methods("POST")

	return r
}