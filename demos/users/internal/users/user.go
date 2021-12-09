package users

import "net/http"

// This should actually be pulled into a separate package
// since used from multiple locations
type User struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserDatastore interface {
	CreateUser(user *User) error
	FindUser(email, password string) (*User, error)
}

type UserAuth interface {
	IsTokenExists(r *http.Request) (bool, string)
	IsUserTokenValid(token string) bool
	UserFromToken(tokenString string) (*User, error)
	GetTokenForUser(user *User) (string, error)
}
