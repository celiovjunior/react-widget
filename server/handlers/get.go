package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/celiovjunior/goserverwidget/models"
)

func Get(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		log.Printf("Erro ao fazer conversão: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}

	feedback, err := models.Get(int64(id))
	if err != nil {
		log.Printf("Erro ao executar método no registro: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(feedback)
}
