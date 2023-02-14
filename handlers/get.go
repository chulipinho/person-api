package handlers

import (
	"net/http"

	"github.com/chulipinho/person-api/data"
)

func (h *PersonHandler) GetAll(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("[DEBUG] Get all people")

	p, err := h.db.Get()
	if err != nil {
		h.l.Println("[ERROR] ", err.Error())
		data.ToJSON(err, rw)
	}

	data.ToJSON(p, rw)
}

func (h *PersonHandler) GetById(rw http.ResponseWriter, r *http.Request) {
	id := getId(r)
	h.l.Printf("[DEBUG] Getting person id %d\n", id)

	p, err := h.db.GetById(id)
	if err != nil {
		h.l.Println("[ERROR] ", err.Error())
		data.ToJSON(err, rw)
	}

	data.ToJSON(p, rw)
}
