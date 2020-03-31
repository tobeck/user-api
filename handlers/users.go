package handlers

import (
	"log"
	"net/http"
	"tobeck.github.com/user-api/data"
	"encoding/json"
)

// Users struct
type Users struct {
	l *log.Logger
}

// NewUsers func
func NewUsers(l*log.Logger) *Users {
	return &Users{l}
}

// ServeHTTP func
func (u *Users) ServeHTTP(rw http.ResponseWriter, h *http.Request) {
	lu := data.GetUsers()
	d, err := json.Marshal(lu)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
	rw.Write(d)
}