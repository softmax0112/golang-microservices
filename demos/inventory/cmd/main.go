package main

import (
	"log"
	"net/http"

	"github.com/KernelGamut32/golang-microservices/demos/inventory/internal/routes"
)

func main() {
	r := routes.Handlers()

	err := http.ListenAndServe(":5100", r)
	if err != nil {
		log.Fatal(err)
	}
}
