package handlers

import (
	"net/http"

	"github.com/chulipinho/person-api/data"
)

func (h *PersonHandler) Put(rw http.ResponseWriter, r *http.Request) {
	id := getId(r)
	h.l.Printf("[DEBUG] Updating person ID: %d\n", id)

	var person data.Person
	data.FromJSON(&person, r.Body)

	err := h.db.Put(id, person)
	if err != nil {
		h.l.Println("[ERROR] updating person: " + err.Error())
		data.ToJSON(&GenericError{err.Error()}, rw)
	}
}
