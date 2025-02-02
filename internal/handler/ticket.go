package handler

import (
	"net/http"
)

func GetTicket(w http.ResponseWriter, r *http.Request) {
	// idStr := r.URL.Query().Get("id")
	// id, err := strconv.Atoi(idStr)
	// if err != nil {
	// 	http.Error(w, "Invalid ticket ID", http.StatusBadRequest)
	// 	return
	// }

	// json.NewEncoder(w).Encode(ticket)
}

func BuyTicket(w http.ResponseWriter, r *http.Request) {

}

func GetUserTickets(w http.ResponseWriter, r *http.Request) {

}
