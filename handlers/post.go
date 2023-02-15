package handlers

import (
	"net/http"

	"github.com/chulipinho/person-api/data"
)

func (h *PersonHandler) Post(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("[DEBUG] Posting new person")
	var person data.Person

	data.FromJSON(&person, r.Body)
	err := h.db.Post(person)
	if err != nil {
		h.l.Println("[ERROR] Getting all people: ", err.Error())
		data.ToJSON(&GenericError{err.Error()}, rw)
		return
	}
}
