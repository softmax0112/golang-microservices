package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

// Use `docker run -p 3306:3306 -p 33060:33060 --name mysqldb -v ~/mysql:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=pass -d mysql` to create DB

var db *sql.DB

func initDB() {

	db = connectDB()

	db.SetMaxIdleConns(4)
	db.SetMaxOpenConns(4)
	db.SetConnMaxLifetime(time.Second * 15)
}

func Get() *sql.DB {
	if db == nil {
		initDB()
	}
	return db
}

func getDBConfig() (username string, password string,
	databasename string, databaseHost string) {
	dir, _ := os.Getwd()
	viper.SetConfigName("app")
	viper.AddConfigPath(dir + "/../configs")
	viper.AutomaticEnv()

	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	databasename = viper.GetString("MYSQL_DATABASE")
	databaseHost = viper.GetString("MYSQL_SERVICE_HOST")
	username = viper.GetString("MYSQL_USERNAME")
	password = viper.GetString("MYSQL_PASSWORD")

	return
}

func createAndOpen(name string, dbURI string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dbURI)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + name)
	if err != nil {
		return nil, err
	}

	db, err = sql.Open("mysql", dbURI+name)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
		id int NOT NULL AUTO_INCREMENT,
		name varchar(100) NOT NULL,
		email varchar(100) NOT NULL,
		password varchar(100) NOT NULL,
		PRIMARY KEY (id)
	);`)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func connectDB() *sql.DB {
	username, password, databasename, databaseHost := getDBConfig()
	dbURI := fmt.Sprintf("%s:%s@(%s)/", username, password, databaseHost)

	db, err := createAndOpen(databasename, dbURI)
	if err != nil {
		fmt.Println("error", err)
		log.Fatalf(err.Error())
	}
	err = db.Ping()

	if err != nil {
		fmt.Println("error", err)
		log.Fatalf(err.Error())
	}

	log.Println("Successfully connected to users db!")
	return db
}
