package service

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/KernelGamut32/golang-microservices/demos/inventory/internal/inventory"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

var inventoryService *InventoryService

type reorderResult struct {
	Message string
}

func getReorderConfig() string {
	dir, _ := os.Getwd()
	viper.SetConfigName("app")
	viper.AddConfigPath(dir + "/../configs")
	viper.AutomaticEnv()

	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	return viper.GetString("REORDER_ENDPOINT")
}

func Get() *InventoryService {
	if inventoryService == nil {
		inventoryService = &InventoryService{DB: GetInventoryDataStore()}
		return inventoryService
	}
	return inventoryService
}

type InventoryService struct {
	DB      inventory.InventoryDatastore
}

func (is *InventoryService) SetInitial(w http.ResponseWriter, r *http.Request) {
	inventory := &inventory.Inventory{}
	json.NewDecoder(r.Body).Decode(inventory)
	if err := is.DB.CreateInventory(inventory); err != nil {
		log.Print("error occured when creating new inventory ", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(inventory)
}

func (is *InventoryService) UpdateCurrent(w http.ResponseWriter, r *http.Request) {
	inventory := inventory.Inventory{}
	params := mux.Vars(r)
	var id = params["id"]

	json.NewDecoder(r.Body).Decode(&inventory)

	if err := is.DB.UpdateInventory(id, inventory); err != nil {
		log.Print("error occured when updating inventory ", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	message := ""
	if inventory.Quantity <= 15 {
		if rr, err := processLowInventory(w, r, inventory.ProductNumber); err != nil {
			log.Print("error occurred on attempts to reorder dynamically ", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		} else {
			message = rr.Message
		}
	}
	var resp = map[string]interface{}{"inventory": inventory, "message": message}
	json.NewEncoder(w).Encode(&resp)
}

func (is *InventoryService) GetCurrent(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var prod_num = params["productNumber"]

	inventory, err := is.DB.GetInventory(prod_num)

	if err != nil {
		log.Print("error occured when getting inventory ", err.Error())
		w.WriteHeader(http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(&inventory)
}

func processLowInventory(w http.ResponseWriter, r *http.Request, prod_num string) (reorderResult, error) {
	reorderEndpoint := getReorderConfig() + "/" + prod_num

	client := &http.Client{
		Timeout: time.Second * 10,
	}
	req, err := http.NewRequest("POST", reorderEndpoint, nil)
	if err != nil {
		log.Print("error occurred requesting reorder ", err.Error())
		return reorderResult{}, err
	}
	req.Header.Set("x-access-token", r.Header.Get("x-access-token"))
	response, err := client.Do(req)

	if response.StatusCode != http.StatusOK || err != nil {
		log.Print("error occurred requesting reorder ", err.Error())
		return reorderResult{}, err
	}
	defer response.Body.Close()

	rr := reorderResult{}
	json.NewDecoder(response.Body).Decode(&rr)
	return rr, nil
}