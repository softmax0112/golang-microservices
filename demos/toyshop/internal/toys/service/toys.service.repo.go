package service

import (
	"context"
	"database/sql"
	"log"
	"time"

	database "github.com/KernelGamut32/golang-microservices/demos/toyshop/internal/db"
	"github.com/KernelGamut32/golang-microservices/demos/toyshop/internal/toys"
)

type ToysDB struct {
	*sql.DB
}

func GetToysDataStore() toys.ToysDatastore {
	return &ToysDB{database.Get()}
}

func (db *ToysDB) CreateToy(toy *toys.Toy) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result, err := db.ExecContext(ctx, "insert into toys (product_number, name, description, unit_cost) values (?, ?,?,?)",
		toy.ProductNumber, toy.Name, toy.Description, toy.UnitCost)
	if err != nil {
		return err
	}
	id, e := result.LastInsertId()
	if e != nil {
		return e
	}

	toy.ID = uint(id)

	return nil
}

func (db *ToysDB) GetAllToys() ([]toys.Toy, error) {
	var theToys []toys.Toy

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	rows, err := db.QueryContext(ctx, "select id, product_number, name, description, unit_cost from toys")
	if err != nil {
		log.Print("error occured when getting all toys ", err.Error())
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var toy toys.Toy
		rows.Scan(&toy.ID, &toy.ProductNumber, &toy.Name, &toy.Description, &toy.UnitCost)
		theToys = append(theToys, toy)
	}
	return theToys, nil
}

func (db *ToysDB) UpdateToy(id string, toy toys.Toy) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result, err := db.ExecContext(ctx, "update toys set product_number = ?, name = ?, description = ?, unit_cost = ? where id = ?",
		toy.ProductNumber, toy.Name, toy.Description, toy.UnitCost, id)
	if err != nil {
		log.Print("error occured when updating toy ", err.Error())
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

func (db *ToysDB) DeleteToy(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result, err := db.ExecContext(ctx, "delete from toys where id = ?", id)
	if err != nil {
		log.Print("error occured when deleting toy ", err.Error())
		return err
	}
	_, err = result.RowsAffected()
	if err != nil {
		log.Fatal("could not update database ", err.Error())
		return err
	}
	return nil
}

func (db *ToysDB) GetToy(id string) (toys.Toy, error) {
	var toy toys.Toy

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	row := db.QueryRowContext(ctx, "select id, product_number, name, description, unit_cost from toys where id=?", id)
	err := row.Scan(&toy.ID, &toy.ProductNumber, &toy.Name, &toy.Description, &toy.UnitCost)

	if err == sql.ErrNoRows {
		return toys.Toy{}, err
	}
	return toy, nil
}
