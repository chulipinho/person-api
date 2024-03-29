package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/chulipinho/person-api/data"
	"github.com/gorilla/mux"
)

type PersonHandler struct {
	db data.IPersonDAO
	l  *log.Logger
}

func NewPersonHandler(db data.IPersonDAO, l *log.Logger) *PersonHandler {
	return &PersonHandler{
		db, l,
	}
}

func getId(r *http.Request) int {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	return id
}
