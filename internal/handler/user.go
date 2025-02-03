package handler

import (
	"encoding/json"
	"net/http"

	"github.com/ShekleinAleksey/goTickets/internal/entity"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user entity.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	//Создание пользователя

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// Пополнить баланс
func Replenish(w http.ResponseWriter, r *http.Request) {

}
