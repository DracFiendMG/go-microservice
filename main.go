package main

import (
	"log"
	"net/http"

	"github.com/DracFiendMG/go-microservice/router"
)

func main() {
	r := router.SetupRouter()
	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}