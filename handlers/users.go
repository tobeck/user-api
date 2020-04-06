package handlers

import (
	"context"
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

// GetUsers function to retrieve users
func (u *Users) GetUsers(rw http.ResponseWriter, r *http.Request) {
	u.l.Println("GET Users")
	usr := data.GetUsers()

	err := usr.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

// AddUser function to add new users
func (u *Users) AddUser(rw http.ResponseWriter, r *http.Request) {
	u.l.Println("POST Users")
	usr := r.Context().Value(KeyUser{}).(data.User)
	data.AddUser(&usr)
}

// KeyUser struct
type KeyUser struct {
}

// MiddlewareValidateUser func
func (u Users) MiddlewareValidateUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		usr := data.User{}

		err := usr.FromJSON(r.Body)
		if err != nil {
			http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
			return
		}
		ctx := context.WithValue(r.Context(), KeyUser{}, usr)
		req := r.WithContext(ctx)

		next.ServeHTTP(rw, req)
	})
}
