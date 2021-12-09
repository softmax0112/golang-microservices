package service

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/KernelGamut32/golang-microservices/demos/toyshop/internal/inventory"
	"github.com/KernelGamut32/golang-microservices/demos/toyshop/internal/toys"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
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

func getInvConfig() string {
	dir, _ := os.Getwd()
	viper.SetConfigName("app")
	viper.AddConfigPath(dir + "/../configs")
	viper.AutomaticEnv()

	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	return viper.GetString("INV_ENDPOINT")
}

func (ts *ToysService) CreateToy(w http.ResponseWriter, r *http.Request) {
	toy := &toys.Toy{}
	json.NewDecoder(r.Body).Decode(toy)
	if err := ts.DB.CreateToy(toy); err != nil {
		log.Print("error occured when creating new toy ", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := processNewInventory(w, r, toy.ProductNumber); err != nil {
		log.Print("error occurred on attempts to add inventory dynamically ", err.Error())
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

func processNewInventory(w http.ResponseWriter, r *http.Request, prod_num string) error {
	invEndpoint := getInvConfig()
	inventory := &inventory.Inventory{ID: 0, ProductNumber: prod_num, Quantity: 50}

	requestBody, err := json.Marshal(inventory)
	if err != nil {
		log.Print("error occurred marshaling inventory object ", err.Error())
		return err
	}
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	req, err := http.NewRequest("POST", invEndpoint, bytes.NewBuffer(requestBody))
	if err != nil {
		log.Print("error occurred requesting new inventory ", err.Error())
		return err
	}
	req.Header.Set("x-access-token", r.Header.Get("x-access-token"))
	response, err := client.Do(req)

	if response.StatusCode != http.StatusCreated || err != nil {
		log.Print("error occurred requesting new inventory ", err.Error())
		return err
	}
	defer response.Body.Close()

	return nil
}