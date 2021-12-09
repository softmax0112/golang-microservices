package service

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/KernelGamut32/golang-microservices/demos/users/internal/users"
	"github.com/KernelGamut32/golang-microservices/demos/users/internal/users/auth"
)

var usersService *UsersService

func Get() *UsersService {
	if usersService == nil {
		usersService = &UsersService{DB: GetUsersDataStore(), JwtAuth: auth.GetAuthenticator()}
		return usersService
	}
	return usersService
}

type UsersService struct {
	DB      users.UserDatastore
	JwtAuth users.UserAuth
}

func (us *UsersService) Login(w http.ResponseWriter, r *http.Request) {
	user := &users.User{}
	err := json.NewDecoder(r.Body).Decode(user)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	currUser, err := us.DB.FindUser(user.Email, user.Password)

	if err != nil {
		log.Print("error occued FindUser ", err.Error())
		w.WriteHeader(http.StatusNotFound)
		return
	}

	tokenString, _ := us.JwtAuth.GetTokenForUser(currUser)

	http.SetCookie(w, &http.Cookie{
		Name:       auth.TokenName,
		Value:      tokenString,
		Path:       "/",
		RawExpires: "0",
	})

	var resp = map[string]interface{}{"status": true, "access-token": tokenString, "user": currUser}
	json.NewEncoder(w).Encode(resp)
}

func (us *UsersService) CreateUser(w http.ResponseWriter, r *http.Request) {

	user := &users.User{}
	json.NewDecoder(r.Body).Decode(user)

	_, err := us.DB.FindUser(user.Email, user.Password)

	if err == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := us.DB.CreateUser(user); err != nil {
		log.Print("error occued CreateUser ", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)

	var resp = map[string]interface{}{"status": true, "user": user}
	json.NewEncoder(w).Encode(resp)
}

func (us *UsersService) VerifyAuth(w http.ResponseWriter, r *http.Request) {
	exist, token := us.JwtAuth.IsTokenExists(r)
	if !exist {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	usr, err := us.JwtAuth.UserFromToken(token)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(usr)
}
