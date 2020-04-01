package handlers

import (
	"log"
	"net/http"
	"github.com/tobeck/user-api/data"
)

// Users struct
type Users struct {
	l *log.Logger
}

// NewUsers func to write json
func NewUsers(l*log.Logger) *Users {
	return &Users{l}
}

// ServeHTTP
func (u*Users) ServeHTTP(rw http.ResponseWriter, r*http.Request) {
	if r.Method == http.MethodGet {
		u.getUsers(rw, r)
		return
	}

	// Catch all for not yet implemented
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

// getUsers func return Users to the ResponseWriter
func (u *Users) getUsers(rw http.ResponseWriter, h *http.Request) {
	lu := data.GetUsers()
	// Use ToJSON encoder instead of marshal, to avoid having to buffer into memory
	err := lu.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}