package handlers

import (
	"log"

	"github.com/chulipinho/person-api/data"
)

type PersonHandler struct {
	db data.PersonDAO
	l  log.Logger
}

func New(db data.PersonDAO, l log.Logger) *PersonHandler {
	return &PersonHandler{
		db, l,
	}
}
