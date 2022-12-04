package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/celiovjunior/goserverwidget/models"
)

func Create(w http.ResponseWriter, r *http.Request) {
	var feedback models.Feedback

	err := json.NewDecoder(r.Body).Decode(&feedback)

	if err != nil {
		log.Printf("Erro ao fazer decode do json: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	id, err := models.Insert(feedback)

	var resp map[string]any
	if err != nil {
		resp = map[string]any{
			"Error":   true,
			"Message": fmt.Sprintf("Ocorreu um erro ao tentar inserir: %v", err),
		}
	} else {
		resp = map[string]any{
			"Error":   false,
			"Message": fmt.Sprintf("Feedback inserido. ID: %d", id),
		}
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
