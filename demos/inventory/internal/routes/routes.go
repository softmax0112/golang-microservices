package routes

import (
	"net/http"

	inventoryService "github.com/KernelGamut32/golang-microservices/demos/inventory/internal/inventory/service"
	"github.com/gorilla/mux"
)

func Handlers() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)

	r.Use(CommonMiddleware)

	is := inventoryService.Get()

	r.HandleFunc("/inventory/{productNumber}", is.GetCurrent).Methods("GET")
	r.HandleFunc("/inventory", is.SetInitial).Methods("POST")
	r.HandleFunc("/inventory/{id}", is.UpdateCurrent).Methods("PUT")

	return r
}

func CommonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")
		next.ServeHTTP(w, r)
	})
}
