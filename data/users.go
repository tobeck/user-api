package data

import (
	"encoding/json"
	"io"
)

// User struct that defines a user entry
type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Gender    string `json:"gender"`
	IPAddress string `json:"ip_address"`
}

// Users type
type Users []*User

// ToJSON func to encode to json
func (u *Users) ToJSON(w io.Writer) error {
	// Create new encoder to use instead of marshal
	e := json.NewEncoder(w)
	return e.Encode(u)
}

// GetUsers GET users
func GetUsers() Users {
	return userList
}

// TODO: Remove once reading from testData file
var userList = []*User{
	&User{
		ID:        1,
		FirstName: "Gabrielle",
		LastName:  "Kear",
		Email:     "gkear0@trellian.com",
		Gender:    "Female",
		IPAddress: "8.79.142.95",
	},
	&User{
		ID:        2,
		FirstName: "Fabian",
		LastName:  "Banfill",
		Email:     "fbanfill1@woothemes.com",
		Gender:    "Male",
		IPAddress: "174.192.211.88",
	},
}
