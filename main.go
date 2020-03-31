package main

import (
	"tobeck.github.com/user-api/handlers"
	"log"
	"net/http"
	"os"
	"time"
	"context"
	"os/signal"
)

func main() {
	l := log.New(os.Stdout, "user-api", log.LstdFlags)

	// User handler
	uh := handlers.NewUsers(l)

	// Server mux
	sm := http.NewServeMux()
	sm.Handle("/", uh)

	// Server configuration
	s := &http.Server{
		Addr: ":8080",
		Handler: sm,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// Shutdown the server gracefully
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <- sigChan
	l.Println("Recieved terminate, graceful shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}
