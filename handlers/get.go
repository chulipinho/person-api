package handlers

import (
	"net/http"

	"github.com/chulipinho/person-api/data"
)

func (h *PersonHandler) GetAll(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("[DEBUG] Get all people")

	p, err := h.db.Get()
	if err != nil {
		h.l.Println("[ERROR] Getting all people: ", err.Error())
		data.ToJSON(&GenericError{err.Error()}, rw)
		return
	}

	data.ToJSON(p, rw)
}

func (h *PersonHandler) GetById(rw http.ResponseWriter, r *http.Request) {
	id := getId(r)
	h.l.Printf("[DEBUG] Getting person ID: %d\n", id)

	p, err := h.db.GetById(id)
	if err != nil {
		h.l.Println("[ERROR] Getting person: ", err.Error())
		data.ToJSON(&GenericError{err.Error()}, rw)
		return
	}

	data.ToJSON(p, rw)
}
