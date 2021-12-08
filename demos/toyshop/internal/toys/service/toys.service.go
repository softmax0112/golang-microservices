package service

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/KernelGamut32/golang-microservices/demos/toyshop/internal/toys"
	"github.com/gorilla/mux"
)

var toysService *ToysService

func Get() *ToysService {
	if toysService == nil {
		toysService = &ToysService{DB: GetToysDataStore()}
		return toysService
	}
	return toysService
}

type ToysService struct {
	DB      toys.ToysDatastore
}

func (ts *ToysService) CreateToy(w http.ResponseWriter, r *http.Request) {
	toy := &toys.Toy{}
	json.NewDecoder(r.Body).Decode(toy)
	log.Print(toy)
	if err := ts.DB.CreateToy(toy); err != nil {
		log.Print("error occured when creating new toy ", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(toy)
}

func (ts *ToysService) GetAllToys(w http.ResponseWriter, r *http.Request) {
	theToys, err := ts.DB.GetAllToys()
	if err != nil {
		log.Print("error occured when getting all toys ", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(theToys)
}

func (ts *ToysService) UpdateToy(w http.ResponseWriter, r *http.Request) {
	toy := toys.Toy{}
	params := mux.Vars(r)
	var id = params["id"]

	json.NewDecoder(r.Body).Decode(&toy)

	if err := ts.DB.UpdateToy(id, toy); err != nil {
		log.Print("error occured when updating toy ", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(&toy)
}

func (ts *ToysService) DeleteToy(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var id = params["id"]

	if err := ts.DB.DeleteToy(id); err != nil {
		log.Print("error occured when deleting toy ", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode("toy deleted")
}

func (ts *ToysService) GetToy(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var id = params["id"]

	toy, err := ts.DB.GetToy(id)

	if err != nil {
		log.Print("error occured when getting toy ", err.Error())
		w.WriteHeader(http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(&toy)
}
