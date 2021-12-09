package users

// This should actually be pulled into a separate package
// since used from multiple locations
type User struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
