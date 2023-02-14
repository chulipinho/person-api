package handlers

import (
	"net/http"

	"github.com/chulipinho/person-api/data"
)

func (h *PersonHandler) Delete(rw http.ResponseWriter, r *http.Request) {
	id := getId(r)
	err := h.db.Delete(id)
	if err != nil {
		h.l.Println("[ERROR] Deleting person: ", err.Error())
		data.ToJSON(&GenericError{err.Error()}, rw)
	}
}
