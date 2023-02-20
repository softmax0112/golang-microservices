package service

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"time"

	"<go module name>/internal/users"
	database "<go module name>/internal/db"

	"golang.org/x/crypto/bcrypt"
)

type UsersDB struct {
	*sql.DB
}

func GetUsersDataStore() users.UserDatastore {
	return &UsersDB{database.Get()}
}

func (db *UsersDB) CreateUser(user *users.User) error {
	if user.Email == "" || user.Password == "" || user.Name == "" {
		return errors.New("user service repo - cannot have empty fields")
	}

	pass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		return errors.New("user service repo - password encryption failed")
	}
	user.Password = string(pass)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	result, err := db.ExecContext(ctx, "insert into users (name, email, password) values (?, ?, ?)",
		user.Name, user.Email, user.Password)

	if err != nil {
		return err
	}

	id, e := result.LastInsertId()
	if e != nil {
		return e
	}

	user.ID = uint(id)

	return nil
}

func (db *UsersDB) FindUser(email, password string) (*users.User, error) {
	user := &users.User{}

	if email == "" || password == "" {
		return nil, errors.New("user service repo - cannot have empty email or password")
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	row := db.QueryRowContext(ctx, "select id, name, email, password from users where email = ?", email)

	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password)

	if err == sql.ErrNoRows {
		return nil, err
	}

	errf := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if errf != nil {
		return nil, errors.New("user service repo - invalid login credentials; please try again")
	}

	return user, nil
}
