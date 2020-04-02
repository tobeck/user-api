package handlers

import (
	"log"
	"net/http"

	"github.com/tobeck/user-api/data"
)

// Users logging struct
type Users struct {
	l *log.Logger
}

// NewUsers dependency injection func
func NewUsers(l *log.Logger) *Users {
	return &Users{l}
}

// GetUsers func return Users to the ResponseWriter
func (u *Users) GetUsers(rw http.ResponseWriter, r *http.Request) {
	u.l.Println("GET Users")

	lu := data.GetUsers()

	// Use ToJSON encoder instead of marshal, to avoid having to buffer into memory
	err := lu.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}