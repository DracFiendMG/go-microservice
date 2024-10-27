package handler

import (
	"encoding/json"
	"net/http"

	"github.com/DracFiendMG/go-microservice/model"
)

var users = []model.User{
	{ID: 1, Name: "Alice"},
	{ID: 2, Name: "Bob"},
}

func GetUsers(writer http.ResponseWriter, r *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(users)
}

func CreateUser(writer http.ResponseWriter, r *http.Request) {
	var user model.User
	json.NewDecoder(r.Body).Decode(&user)
	user.ID = len(users) + 1
	users = append(users, user)
	json.NewEncoder(writer).Encode(user)
}