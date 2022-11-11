package handlers

import (
	"api-go/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func Create(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo

	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		log.Println("Erro decode JSON: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	id, err := models.Insert(todo)
	var resp map[string]any

	if err != nil {
		resp = map[string]any{
			"Error": true,
			"Message": fmt.Sprintf("Erro ao inserir: %v", err),
		}
	} else {
		resp = map[string]any{
			"Error": false,
			"Message": fmt.Sprintf("Todo inserido com sucesso | ID: %d", id)
		}
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}