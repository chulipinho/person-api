package handlers

import (
	"net/http"

	"github.com/chulipinho/person-api/data"
)

func (h *PersonHandler) Post(rw http.ResponseWriter, r *http.Request) {
	var person data.Person

	data.FromJSON(&person, r.Body)
	h.db.Post(person)
}
