package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/chulipinho/person-api/data"
	"github.com/chulipinho/person-api/handlers"
	"github.com/gorilla/mux"
)

func main() {
	l := log.New(os.Stdout, "person-api", log.LstdFlags)
	db := data.NewMock()

	ph := handlers.NewPersonHandler(db, l)

	r := mux.NewRouter()
	r.Use(handlers.PersonMiddleware)

	getR := r.Methods(http.MethodGet).PathPrefix("/person/").Subrouter()
	getR.HandleFunc("/", ph.GetAll)
	getR.HandleFunc("/{id}", ph.GetById)

	postR := r.Methods(http.MethodPost).PathPrefix("/person/").Subrouter()
	postR.HandleFunc("/", ph.Post)

	deleteR := r.Methods(http.MethodDelete).PathPrefix("/person/").Subrouter()
	deleteR.HandleFunc("/{id}", ph.Delete)

	putR := r.Methods(http.MethodPut).PathPrefix("/person/").Subrouter()
	putR.HandleFunc("/{id}", ph.Put)

	server := http.Server{
		Addr:         ":1234",           // configure the bind address
		Handler:      r,                 // set the default handler
		ErrorLog:     l,                 // set the logger for the server
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}

	go func() {
		l.Println("Starting server on port 1234")

		err := server.ListenAndServe()
		if err != nil {
			l.Fatalf("Error starting server: %s\n", err)
		}
	}()

	ch := make(chan os.Signal)

	signal.Notify(ch, os.Interrupt)
	signal.Notify(ch, os.Kill)

	sig := <-ch
	l.Println("Exiting application. Signal: ", sig)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	server.Shutdown(ctx)
}
