package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/chulipinho/person-api/data"
	"github.com/gorilla/mux"
)

type PersonHandler struct {
	db data.PersonDAO
	l  *log.Logger
}

func NewPersonHandler(db data.PersonDAO, l *log.Logger) *PersonHandler {
	return &PersonHandler{
		db, l,
	}
}

type GenericError struct {
	Message string `json:"message"`
}

func getId(r *http.Request) int {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	return id
}
