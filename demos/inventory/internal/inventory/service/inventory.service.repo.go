package service

import (
	"context"
	"database/sql"
	"log"
	"time"

	database "github.com/KernelGamut32/golang-microservices/demos/inventory/internal/db"
	"github.com/KernelGamut32/golang-microservices/demos/inventory/internal/inventory"
)

type InventoryDB struct {
	*sql.DB
}

func GetInventoryDataStore() inventory.InventoryDatastore {
	return &InventoryDB{database.Get()}
}

func (db *InventoryDB) CreateInventory(inventory *inventory.Inventory) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result, err := db.ExecContext(ctx, "insert into inventory (product_number, quantity) values (?, ?)",
		inventory.ProductNumber, inventory.Quantity)
	if err != nil {
		return err
	}
	id, e := result.LastInsertId()
	if e != nil {
		return e
	}

	inventory.ID = uint(id)

	return nil
}

func (db *InventoryDB) UpdateInventory(id string, inventory inventory.Inventory) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result, err := db.ExecContext(ctx, "update inventory set quantity = ? where id = ?",
		inventory.Quantity, id)
	if err != nil {
		log.Print("error occured when updating inventory ", err.Error())
		return err
	}
	num, err := result.RowsAffected()
	if err != nil {
		log.Fatal("could not update database ", err.Error())
		return err
	}

	log.Println("number of rows affected is ", num)
	return nil
}

func (db *InventoryDB) GetInventory(productNumber string) (inventory.Inventory, error) {
	var inv inventory.Inventory
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	row := db.QueryRowContext(ctx, "select id, product_number, quantity from inventory where product_number=?", productNumber)
	err := row.Scan(&inv.ID, &inv.ProductNumber, &inv.Quantity)
	if err == sql.ErrNoRows {
		return inventory.Inventory{}, err
	}
	return inv, nil
}
